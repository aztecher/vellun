package flags

import (
	"crypto/tls"
	"path/filepath"

	"github.com/aztecher/vellun/util/lazy"
	"github.com/spf13/pflag"
	"sigs.k8s.io/controller-runtime/pkg/certwatcher"
	"sigs.k8s.io/controller-runtime/pkg/metrics/filters"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

// MetricsOptions provides command line flags for manager option
type MetricsOptions struct {
	// Metrics Options
	// There are used to configure the metrics server

	// BindAddress is the field that stores the value of the --metrics-server-address flag.
	BindAddress string
	// SecureServing is the field that stores the value of the --metrics-secure flag.
	SecureServing bool
	// CertPath
	CertPath string
	// CertName
	CertName string
	// CertKey
	CertKey string
}

// addMetricsOptions adds the metrics options flags to the flag set.
func addMetricsOptions(fs *pflag.FlagSet, options *MetricsOptions) {
	fs.StringVar(&options.BindAddress, "metrics-bind-address", "0", "The address the metrics endpoint binds to. "+
		"Use :8443 for HTTPS or :8080 for HTTP, or leave as 0 to disable the metrics service.")

	fs.BoolVar(&options.SecureServing, "metrics-secure-serving", true,
		"If set, the metrics endpoint is served securely via HTTPS. Use --metrics-secure=false to use HTTP instead.")

	fs.StringVar(&options.CertPath, "metrics-cert-path", "",
		"The directory that contains the metrics server certificate.")
	fs.StringVar(&options.CertName, "metrics-cert-name", "tls.crt", "The name of the metrics server certificate file.")
	fs.StringVar(&options.CertKey, "metrics-cert-key", "tls.key", "The name of the metrics server key file.")
}

// getMetricsOptions
func getMetricsOptions(options MetricsOptions, tlsOptions []func(*tls.Config)) (*metricsserver.Options, error) {
	var metricsOptions *metricsserver.Options
	var certWatcher *certwatcher.CertWatcher
	var err error

	metricsOptions = &metricsserver.Options{
		BindAddress:   options.BindAddress,
		SecureServing: options.SecureServing,
		TLSOpts:       tlsOptions,
	}

	if options.SecureServing {
		metricsOptions.FilterProvider = filters.WithAuthenticationAndAuthorization
	}

	if len(options.CertPath) > 0 {
		certWatcher, err = certwatcher.New(
			filepath.Join(options.CertPath, options.CertName),
			filepath.Join(options.CertPath, options.CertKey),
		)
		if err != nil {
			return metricsOptions, err
		}
		metricsOptions.TLSOpts = append(
			metricsOptions.TLSOpts,
			func(config *tls.Config) {
				config.GetCertificate = certWatcher.GetCertificate
			},
		)
		lazy.RegisterMetricsCertWatcher(certWatcher)
	}

	return metricsOptions, nil
}
