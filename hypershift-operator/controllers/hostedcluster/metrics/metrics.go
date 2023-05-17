package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	DeletionDurationMetricName    = "hypershift_cluster_deletion_duration_seconds"
	HostedClusterDeletionDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Help: "Time in seconds it took from HostedCluster having a deletion timestamp to all hypershift finalizers being removed",
		Name: DeletionDurationMetricName,
	}, []string{"namespace", "name"})

	GuestCloudResourcesDeletionDurationMetricName    = "hypershift_cluster_guest_cloud_resources_deletion_duration_seconds"
	HostedClusterGuestCloudResourcesDeletionDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Help: "Time in seconds it took from HostedCluster having a deletion timestamp to the CloudResourcesDestroyed being true",
		Name: GuestCloudResourcesDeletionDurationMetricName,
	}, []string{"namespace", "name"})

	AvailableDurationName          = "hypershift_cluster_available_duration_seconds"
	HostedClusterAvailableDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Help: "Time in seconds it took from initial cluster creation to HostedClusterAvailable condition becoming true",
		Name: AvailableDurationName,
	}, []string{"namespace", "name"})

	InitialRolloutDurationName          = "hypershift_cluster_initial_rollout_duration_seconds"
	HostedClusterInitialRolloutDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Help: "Time in seconds it took from initial cluster creation and rollout of initial version",
		Name: InitialRolloutDurationName,
	}, []string{"namespace", "name"})

	LimitedSupportEnabledName = "hypershift_cluster_limited_support_enabled"
	LimitedSupportEnabled     = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Help: "Indicates if limited support is enabled for each cluster",
		Name: LimitedSupportEnabledName,
	}, []string{"namespace", "name", "_id"})

	SilenceAlertsName = "hypershift_cluster_silence_alerts"
	SilenceAlerts     = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Help: "Indicates if alerts are silenced for each cluster",
		Name: SilenceAlertsName,
	}, []string{"namespace", "name", "_id"})

	ProxyName   = "hypershift_cluster_proxy"
	ProxyConfig = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Help: "Indicates cluster proxy state for each cluster",
		Name: ProxyName,
	}, []string{"namespace", "name", "proxy_http", "proxy_https", "proxy_trusted_ca"})

	SkippedCloudResourcesDeletionName = "hypershift_cluster_skipped_cloud_resources_deletion"
	SkippedCloudResourcesDeletion     = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Help: "Indicates the operator will skip the aws resources deletion",
		Name: SkippedCloudResourcesDeletionName,
	}, []string{"namespace", "name"})
)

func init() {
	metrics.Registry.MustRegister(
		HostedClusterDeletionDuration,
		HostedClusterGuestCloudResourcesDeletionDuration,
		HostedClusterAvailableDuration,
		HostedClusterInitialRolloutDuration,
		LimitedSupportEnabled,
		SilenceAlerts,
		ProxyConfig,
		SkippedCloudResourcesDeletion,
	)
}
