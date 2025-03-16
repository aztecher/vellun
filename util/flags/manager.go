package flags

import (
	"crypto/tls"

	"github.com/spf13/pflag"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// ManagerOptions provides command line flags for manager options.
type ManagerOptions struct {
	Webhook *WebhookOptions
	Metrics *MetricsOptions

	EnableLeaderElection        bool
	ProveAddr                   string
	GPUGroupConcurrency         int
	GPUNetworkPolicyConcurrency int
}

// NewManagerOptions initialize the ManagerOptions
func NewManagerOptions() ManagerOptions {
	return ManagerOptions{
		Webhook: &WebhookOptions{},
		Metrics: &MetricsOptions{},
	}
}

// AddManagerOptions adds the manager options flags to the flag set.
func AddManagerOptions(fs *pflag.FlagSet, options *ManagerOptions) {
	fs.StringVar(&options.ProveAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	fs.BoolVar(&options.EnableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	//
	fs.IntVar(&options.GPUGroupConcurrency, "gpugroup-concurrency", 10, "Number of GPUGroup to process simultaneously")
	fs.IntVar(&options.GPUNetworkPolicyConcurrency, "gpunetworkpolicy-concurrency", 10, "Number of GPUNetworkPolicy to proccess simultaneously")

	addWebhookOptions(fs, options.Webhook)
	addMetricsOptions(fs, options.Metrics)
}

// GetManagerOptions returns options which can be used to configure a Manager.
func GetManagerOptions(options ManagerOptions) (*webhook.Options, *metricsserver.Options, error) {
	var tlsOptions []func(config *tls.Config)
	var webhookOptions *webhook.Options
	var metricsOptions *metricsserver.Options
	var err error
	// Configure TLSOptions if needed
	webhookOptions, err = getWebhookOptions(*options.Webhook, tlsOptions)
	if err != nil {
		return webhookOptions, metricsOptions, err
	}

	metricsOptions, err = getMetricsOptions(*options.Metrics, tlsOptions)
	if err != nil {
		return webhookOptions, metricsOptions, err
	}

	return webhookOptions, metricsOptions, nil
}
