package ovn_monitor

import "github.com/prometheus/client_golang/prometheus"

var (
	// OVN basic info
	metricOvnHealthyStatus = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "ovn_status",
			Help:      "OVN Health Status. The values are: health(1), unhealth(0).",
		})

	metricOvnInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "ovn_info",
			Help:      "This metric provides basic information about OVN. It is always set to 1.",
		},
		[]string{
			"system_id",
			"rundir",
			"hostname",
			"system_type",
			"system_version",
			"ovs_version",
			"db_version",
		})

	metricRequestErrorNums = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "failed_req_count",
			Help:      "The number of failed requests to OVN stack.",
		},
		[]string{
			"hostname",
		})

	metricLogFileSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "log_file_size",
			Help:      "The size of a log file associated with an OVN component.",
		},
		[]string{
			"hostname",
			"component",
			"filename",
		})

	metricDBFileSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "db_file_size",
			Help:      "The size of a database file associated with an OVN component.",
		},
		[]string{
			"hostname",
			"db_name",
		})

	// OVN Chassis metrics
	metricChassisInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "chassis_info",
			Help:      "Whether the OVN chassis is up (1) or down (0), together with additional information about the chassis.",
		},
		[]string{
			"hostname",
			"uuid",
			"name",
			"ip",
		})

	metricLogicalSwitchInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "logical_switch_info",
			Help:      "The information about OVN logical switch. This metric is always up (1).",
		},
		[]string{
			"hostname",
			"uuid",
			"name",
		})

	metricLogicalSwitchExternalIDs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "logical_switch_external_id",
			Help:      "Provides the external IDs and values associated with OVN logical switches. This metric is always up (1).",
		},
		[]string{
			"hostname",
			"uuid",
			"key",
			"value",
		})

	metricLogicalSwitchPortBinding = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "logical_switch_port_binding",
			Help:      "Provides the association between a logical switch and a logical switch port. This metric is always up (1).",
		},
		[]string{
			"hostname",
			"uuid",
			"port",
		})

	metricLogicalSwitchTunnelKey = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "logical_switch_tunnel_key",
			Help:      "The value of the tunnel key associated with the logical switch.",
		},
		[]string{
			"hostname",
			"uuid",
		})

	metricLogicalSwitchPortsNum = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "logical_switch_ports_num",
			Help:      "The number of logical switch ports connected to the OVN logical switch.",
		},
		[]string{
			"hostname",
			"uuid",
		})

	metricLogicalSwitchPortInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "logical_switch_port_info",
			Help:      "The information about OVN logical switch port. This metric is always up (1).",
		},
		[]string{
			"hostname",
			"uuid",
			"name",
			"chassis",
			"logical_switch",
			"datapath",
			"port_binding",
			"mac_address",
			"ip_address",
		})

	metricLogicalSwitchPortTunnelKey = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "logical_switch_port_tunnel_key",
			Help:      "The value of the tunnel key associated with the logical switch port.",
		},
		[]string{
			"hostname",
			"uuid",
		})

	// OVN Cluster basic info metrics
	metricClusterEnabled = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_enabled",
			Help:      "Is OVN clustering enabled (1) or not (0).",
		},
		[]string{
			"hostname",
			"db_name",
		})

	metricClusterRole = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_role",
			Help:      "A metric with a constant '1' value labeled by server role.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
			"server_role",
		})

	metricClusterStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_status",
			Help:      "A metric with a constant '1' value labeled by server status.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
			"server_status",
		})

	metricClusterTerm = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_term",
			Help:      "The current raft term known by this server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterLeaderSelf = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_leader_self",
			Help:      "Is this server consider itself a leader (1) or not (0).",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterVoteSelf = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_vote_self",
			Help:      "Is this server voted itself as a leader (1) or not (0).",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterElectionTimer = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_election_timer",
			Help:      "The current election timer value.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterNotCommittedEntryCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_log_not_committed",
			Help:      "The number of log entries not yet committed by this server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterNotAppliedEntryCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_log_not_applied",
			Help:      "The number of log entries not yet applied by this server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterLogIndexStart = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_log_index_start",
			Help:      "The log entry index start value associated with this server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterLogIndexNext = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_log_index_next",
			Help:      "The log entry index next value associated with this server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterInConnTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_inbound_connections_total",
			Help:      "The total number of inbound connections to the server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterOutConnTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_outbound_connections_total",
			Help:      "The total number of outbound connections from the server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterInConnErrTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_inbound_connections_error_total",
			Help:      "The total number of failed inbound connections to the server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterOutConnErrTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_outbound_connections_error_total",
			Help:      "The total number of failed outbound connections from the server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	// Todo: The metrics downside are to be implemented
	metricClusterPeerInConnInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_inbound_peer_connected",
			Help:      "This metric appears when a cluster peer is connected to this server. This metric is always 1.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
			"peer_id",
			"peer_address",
		})

	metricClusterPeerOutConnInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_outbound_peer_connected",
			Help:      "This metric appears when this server connects to a cluster peer. This metric is always 1.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
			"peer_id",
			"peer_address",
		})

	metricClusterPeerCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_peer_count",
			Help:      "The total number of peers in this server's cluster.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterPeerNextIndex = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_peer_next_index",
			Help:      "The raft's next index associated with this cluster peer.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
			"peer_id",
		})

	metricClusterPeerMatchIndex = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_peer_match_index",
			Help:      "The raft's match index associated with this cluster peer.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
			"peer_id",
		})

	metricClusterNextIndex = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_next_index",
			Help:      "The raft's next index associated with this server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})

	metricClusterMatchIndex = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: MetricNamespace,
			Name:      "cluster_match_index",
			Help:      "The raft's match index associated with this server.",
		},
		[]string{
			"hostname",
			"db_name",
			"server_id",
			"cluster_id",
		})
)

func registerOvnMetrics() {
	// ovn status metrics
	prometheus.MustRegister(metricOvnHealthyStatus)
	prometheus.MustRegister(metricOvnInfo)
	prometheus.MustRegister(metricRequestErrorNums)
	prometheus.MustRegister(metricLogFileSize)
	prometheus.MustRegister(metricDBFileSize)

	// ovn chassis metrics
	prometheus.MustRegister(metricChassisInfo)
	prometheus.MustRegister(metricLogicalSwitchInfo)
	prometheus.MustRegister(metricLogicalSwitchExternalIDs)
	prometheus.MustRegister(metricLogicalSwitchPortBinding)
	prometheus.MustRegister(metricLogicalSwitchTunnelKey)
	prometheus.MustRegister(metricLogicalSwitchPortsNum)
	prometheus.MustRegister(metricLogicalSwitchPortInfo)
	prometheus.MustRegister(metricLogicalSwitchPortTunnelKey)

	// OVN Cluster basic info metrics
	prometheus.MustRegister(metricClusterEnabled)
	prometheus.MustRegister(metricClusterRole)
	prometheus.MustRegister(metricClusterStatus)
	prometheus.MustRegister(metricClusterTerm)

	prometheus.MustRegister(metricClusterLeaderSelf)
	prometheus.MustRegister(metricClusterVoteSelf)
	prometheus.MustRegister(metricClusterElectionTimer)
	prometheus.MustRegister(metricClusterNotCommittedEntryCount)
	prometheus.MustRegister(metricClusterNotAppliedEntryCount)

	prometheus.MustRegister(metricClusterLogIndexStart)
	prometheus.MustRegister(metricClusterLogIndexNext)
	prometheus.MustRegister(metricClusterInConnTotal)
	prometheus.MustRegister(metricClusterOutConnTotal)
	prometheus.MustRegister(metricClusterInConnErrTotal)
	prometheus.MustRegister(metricClusterOutConnErrTotal)

	// to be implemented
	prometheus.MustRegister(metricClusterPeerNextIndex)
	prometheus.MustRegister(metricClusterPeerMatchIndex)
	prometheus.MustRegister(metricClusterNextIndex)
	prometheus.MustRegister(metricClusterMatchIndex)
	prometheus.MustRegister(metricClusterPeerInConnInfo)
	prometheus.MustRegister(metricClusterPeerOutConnInfo)
	prometheus.MustRegister(metricClusterPeerCount)
}
