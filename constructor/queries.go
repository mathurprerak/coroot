package constructor

var QUERIES = map[string]string{
	"up": `up`,

	"node_info":                   `node_info`,
	"node_cloud_info":             `node_cloud_info`,
	"node_cpu_cores":              `node_resources_cpu_logical_cores`,
	"node_cpu_usage_percent":      `sum(rate(node_resources_cpu_usage_seconds_total{mode!="idle"}[$RANGE])) without(mode) /sum(rate(node_resources_cpu_usage_seconds_total[$RANGE])) without(mode)*100`,
	"node_cpu_usage_by_mode":      `rate(node_resources_cpu_usage_seconds_total{mode!="idle"}[$RANGE]) / ignoring(mode) group_left sum(rate(node_resources_cpu_usage_seconds_total[$RANGE])) without(mode)*100`,
	"node_memory_total_bytes":     `node_resources_memory_total_bytes`,
	"node_memory_available_bytes": `node_resources_memory_available_bytes`,
	"node_memory_free_bytes":      `node_resources_memory_free_bytes`,
	"node_memory_cached_bytes":    `node_resources_memory_cached_bytes`,
	"node_disk_read_time":         `rate(node_resources_disk_read_time_seconds_total[$RANGE])`,
	"node_disk_write_time":        `rate(node_resources_disk_write_time_seconds_total[$RANGE])`,
	"node_disk_reads":             `rate(node_resources_disk_reads_total[$RANGE])`,
	"node_disk_writes":            `rate(node_resources_disk_writes_total[$RANGE])`,
	"node_disk_read_bytes":        `rate(node_resources_disk_read_bytes_total[$RANGE])`,
	"node_disk_written_bytes":     `rate(node_resources_disk_written_bytes_total[$RANGE])`,
	"node_disk_io_time":           `rate(node_resources_disk_io_time_seconds_total[$RANGE])`,
	"node_net_up":                 `node_net_interface_up`,
	"node_net_ip":                 `node_net_interface_ip`,
	"node_net_rx_bytes":           `rate(node_net_received_bytes_total[$RANGE])`,
	"node_net_tx_bytes":           `rate(node_net_transmitted_bytes_total[$RANGE])`,

	"kube_service_info": `kube_service_info`,

	"kube_pod_info":             `kube_pod_info`,
	"kube_pod_labels":           `kube_pod_labels`,
	"kube_pod_status_phase":     `kube_pod_status_phase`,
	"kube_pod_status_ready":     `kube_pod_status_ready{condition="true"}`,
	"kube_pod_status_scheduled": `kube_pod_status_scheduled{condition="true"} > 0`,

	"container_net_latency":                 `container_net_latency_seconds`,
	"container_net_tcp_successful_connects": `rate(container_net_tcp_successful_connects_total[$RANGE])`,
	"container_net_tcp_active_connections":  `container_net_tcp_active_connections`,
	"container_net_tcp_listen_info":         `container_net_tcp_listen_info`,
	"container_log_messages":                `container_log_messages_total`,
	"container_application_type":            `container_application_type`,
	"container_cpu_limit":                   `container_resources_cpu_limit_cores`,
	"container_cpu_usage":                   `rate(container_resources_cpu_usage_seconds_total[$RANGE])`,
	"container_cpu_delay":                   `rate(container_resources_cpu_delay_seconds_total[$RANGE])`,
	"container_throttled_time":              `rate(container_resources_cpu_throttled_seconds_total[$RANGE])`,
	"container_memory_rss":                  `container_resources_memory_rss_bytes`,
	"container_memory_cache":                `container_resources_memory_cache_bytes`,
	"container_memory_limit":                `container_resources_memory_limit_bytes`,
	"container_oom_kills_total":             `container_oom_kills_total`,
	"container_restarts":                    `container_restarts_total`,
	"container_volume_size":                 `container_resources_disk_size_bytes`,
	"container_volume_used":                 `container_resources_disk_used_bytes`,

	"kube_pod_init_container_info":                     `kube_pod_init_container_info`,
	"kube_pod_container_status_ready":                  `kube_pod_container_status_ready > 0`,
	"kube_pod_container_status_waiting":                `kube_pod_container_status_waiting > 0`,
	"kube_pod_container_status_running":                `kube_pod_container_status_running > 0 `,
	"kube_pod_container_status_terminated":             `kube_pod_container_status_terminated > 0`,
	"kube_pod_container_status_waiting_reason":         `kube_pod_container_status_waiting_reason > 0`,
	"kube_pod_container_status_last_terminated_reason": `kube_pod_container_status_last_terminated_reason`,
	"kube_deployment_spec_replicas":                    `kube_deployment_spec_replicas`,
	"kube_daemonset_status_desired_number_scheduled":   `kube_daemonset_status_desired_number_scheduled`,
	"kube_statefulset_replicas":                        `kube_statefulset_replicas`,

	"pg_connections":                  `pg_connections{db!="postgres"}`,
	"pg_up":                           `pg_up`,
	"pg_info":                         `pg_info`,
	"pg_setting":                      `pg_setting`,
	"pg_lock_awaiting_queries":        `pg_lock_awaiting_queries`,
	"pg_latency_seconds":              `pg_latency_seconds`,
	"pg_top_query_calls_per_second":   `pg_top_query_calls_per_second`,
	"pg_top_query_time_per_second":    `pg_top_query_time_per_second`,
	"pg_top_query_io_time_per_second": `pg_top_query_io_time_per_second`,
	"pg_db_queries_per_second":        `pg_db_queries_per_second`,
	"pg_wal_current_lsn":              `pg_wal_current_lsn`,
	"pg_wal_receive_lsn":              `pg_wal_receive_lsn`,
	"pg_wal_reply_lsn":                `pg_wal_reply_lsn`,

	"redis_up":                              `redis_up`,
	"redis_instance_info":                   `redis_instance_info`,
	"redis_commands_duration_seconds_total": `rate(redis_commands_duration_seconds_total[$RANGE])`,
	"redis_commands_total":                  `rate(redis_commands_total[$RANGE])`,
}
