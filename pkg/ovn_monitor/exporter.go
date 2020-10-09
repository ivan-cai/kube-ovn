package ovn_monitor

import (
	"os"
	"sync"
	"time"

	"github.com/greenpau/ovsdb"
	"github.com/prometheus/common/log"
)

const (
	MetricNamespace = "kube_ovn"
)

var (
	appName          = "ovn-monitor"
	isClusterEnabled = true
	tryConnectCnt    = 0
)

// Exporter collects OVN data from the given server and exports them using
// the prometheus metrics package.
type Exporter struct {
	sync.RWMutex
	Client       *ovsdb.OvnClient
	timeout      int
	pollInterval int64
	errors       int64
	errorsLocker sync.RWMutex
}

type Options struct {
	Timeout int
}

// ClusterPeer contains information about a cluster peer.
type ClusterPeer struct {
	ID         string
	Address    string
	NextIndex  uint64
	MatchIndex uint64
	Connection struct {
		Inbound  int
		Outbound int
	}
}

type OVNDBClusterStatus struct {
	cid             string
	sid             string
	status          string
	role            string
	leader          string
	vote            string
	term            float64
	electionTimer   float64
	logIndexStart   float64
	logIndexNext    float64
	logNotCommitted float64
	logNotApplied   float64
	connIn          float64
	connOut         float64
	connInErr       float64
	connOutErr      float64

	peers       map[string]*ClusterPeer
	connections struct {
		inbound  int
		outbound int
	}
}

// NewExporter returns an initialized Exporter.
func NewExporter(opts Options) *Exporter {
	e := Exporter{
		timeout: opts.Timeout,
	}
	client := ovsdb.NewOvnClient()
	client.Timeout = opts.Timeout
	e.Client = client

	err := e.Client.GetSystemID()
	if err != nil {
		log.Errorf("%s failed to get system id: %s", appName, err)
	}

	return &e
}

// StartConnection connect to database socket
func (e *Exporter) StartConnection() error {
	log.Debugf("%s: exporter calls Connect()", e.Client.System.ID)
	if err := e.Client.Connect(); err != nil {
		return err
	}

	if err := e.Client.GetSystemInfo(); err != nil {
		return err
	}
	log.Infof("%s: exporter connect successfully", e.Client.System.Hostname)

	return nil
}

func (e *Exporter) TryClientConnection() {
	for {
		if tryConnectCnt > 5 {
			log.Errorf("%s: ovn-monitor failed to reconnect db socket finally", e.Client.System.Hostname)
			break
		}

		err := e.StartConnection()
		if err != nil {
			tryConnectCnt++
			log.Errorf("%s: ovn-monitor failed to reconnect db socket %v times", e.Client.System.Hostname, tryConnectCnt)
		} else {
			_ = e.Client.GetSystemID()
			log.Infof("%s: ovn-monitor reconnect db successfully", e.Client.System.ID)
			break
		}

		time.Sleep(5 * time.Second)
	}
}

var registerOvnMetricsOnce sync.Once

// StartOvnMetrics register and start to update ovn metrics
func (e *Exporter) StartOvnMetrics() {
	registerOvnMetricsOnce.Do(func() {
		registerOvnMetrics()

		// OVN metrics updater
		go e.ovnMetricsUpdate()
	})
}

// ovnMetricsUpdate updates the ovn metrics for every 30 sec
func (e *Exporter) ovnMetricsUpdate() {
	for {
		e.exportOvnStatusGauge()
		e.exportOvnInfoGauge()
		e.exportOvnLogFileSizeGauge()
		e.exportOvnDBFileSizeGauge()
		e.exportOvnRequestErrorGauge()

		e.exportOvnChassisGauge()
		e.exportLogicalSwitchGauge()
		e.exportLogicalSwitchPortGauge()

		e.exportOvnClusterEnableGauge()
		if isClusterEnabled {
			e.exportOvnClusterInfoGauge()
		}

		time.Sleep(30 * time.Second)
	}
}

// GetExporterName returns exporter name.
func GetExporterName() string {
	return appName
}

// SetPollInterval sets exporter's polling interval.
func (e *Exporter) SetPollInterval(i int64) {
	e.pollInterval = i
}

func (e *Exporter) exportOvnStatusGauge() {
	if up, _ := e.getOvnStatus(); up {
		metricOvnHealthyStatus.Set(1)
	} else {
		metricOvnHealthyStatus.Set(0)
	}
}

func (e *Exporter) exportOvnInfoGauge() {
	// All parameters should be filled in GetSystemInfo when calls NewExporter
	metricOvnInfo.WithLabelValues(e.Client.System.ID, e.Client.System.RunDir, e.Client.System.Hostname,
		e.Client.System.Type, e.Client.System.Version, e.Client.Database.Vswitch.Version,
		e.Client.Database.Vswitch.Schema.Version).Set(1)
}

func (e *Exporter) exportOvnLogFileSizeGauge() {
	components := []string{
		"ovsdb-server-southbound",
		"ovsdb-server-northbound",
		"ovn-northd",
	}
	for _, component := range components {
		file, err := e.Client.GetLogFileInfo(component)
		if err != nil {
			log.Errorf("%s: log-file-%v", component, err)
			e.IncrementErrorCounter()
			continue
		}
		log.Debugf("%s: getOvnLogFileSize() completed GetLogFileInfo(%s)", e.Client.System.ID, component)
		metricLogFileSize.WithLabelValues(e.Client.System.Hostname, file.Component, file.Path).Set(float64(file.Info.Size()))
	}
}

func (e *Exporter) exportOvnDBFileSizeGauge() {
	nbPath := e.Client.Database.Northbound.File.Data.Path
	sbPath := e.Client.Database.Southbound.File.Data.Path
	dirDbMap := map[string]string{
		nbPath: "OVN_Northbound",
		sbPath: "OVN_Southbound",
	}
	for dbFile, database := range dirDbMap {
		fileInfo, err := os.Stat(dbFile)
		if err != nil {
			log.Errorf("Failed to get the DB size for database %s: %v", database, err)
			return
		}
		metricDBFileSize.WithLabelValues(e.Client.System.Hostname, database).Set(float64(fileInfo.Size()))
	}
}

func (e *Exporter) exportOvnRequestErrorGauge() {
	metricRequestErrorNums.WithLabelValues(e.Client.System.Hostname).Set(float64(e.errors))
}

func (e *Exporter) exportOvnChassisGauge() {
	if vteps, err := e.Client.GetChassis(); err != nil {
		log.Errorf("%s: %v", e.Client.Database.Southbound.Name, err)
		e.IncrementErrorCounter()
	} else {
		for _, vtep := range vteps {
			metricChassisInfo.WithLabelValues(e.Client.System.Hostname, vtep.UUID, vtep.Name, vtep.IPAddress.String()).Set(float64(vtep.Up))
		}
	}
}

func (e *Exporter) exportLogicalSwitchGauge() {
	e.setLogicalSwitchInfoMetric()
}

func (e *Exporter) exportLogicalSwitchPortGauge() {
	e.setLogicalSwitchPortInfoMetric()
}

func (e *Exporter) exportOvnClusterEnableGauge() {
	isClusterEnabled, err := getClusterEnableState(e.Client.Database.Northbound.File.Data.Path)
	if err != nil {
		log.Debugf("failed to get output of cluster status: %v", err)
	}
	if isClusterEnabled {
		metricClusterEnabled.WithLabelValues(e.Client.System.Hostname, e.Client.Database.Northbound.File.Data.Path).Set(1)
	} else {
		metricClusterEnabled.WithLabelValues(e.Client.System.Hostname, e.Client.Database.Northbound.File.Data.Path).Set(0)
	}
}

func (e *Exporter) exportOvnClusterInfoGauge() {
	dirDbMap := map[string]string{
		"nb": "OVN_Northbound",
		"sb": "OVN_Southbound",
	}
	for direction, database := range dirDbMap {
		clusterStatus, err := getClusterInfo(direction, database)
		if err != nil {
			log.Errorf("Failed to get Cluster Info for database %s: %v", database, err)
			return
		}
		e.setOvnClusterInfoMetric(clusterStatus, database)
	}
}
