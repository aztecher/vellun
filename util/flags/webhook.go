package flags

import (
	"crypto/tls"
	"path/filepath"

	"github.com/aztecher/vellun/util/lazy"
	"github.com/spf13/pflag"
	"sigs.k8s.io/controller-runtime/pkg/certwatcher"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// WebhookOptions provides command line flags for webhool option
type WebhookOptions struct {
	// Webhook Options
	// There are used to configure the webhook

	// CertPath
	CertPath string
	// CertName
	CertName string
	// CertKey
	CertKey string
}

// addWebhookOptions adds the webhook options flags to the flag set.
func addWebhookOptions(fs *pflag.FlagSet, options *WebhookOptions) {
	fs.StringVar(&options.CertPath, "webhook-cert-path", "", "The directory that contains the webhook certificate.")
	fs.StringVar(&options.CertName, "webhook-cert-name", "tls.crt", "The name of the webhook certificate file.")
	fs.StringVar(&options.CertKey, "webhook-cert-key", "tls.key", "The name of the webhook key file.")
}

// getWebhookOptions
func getWebhookOptions(options WebhookOptions, tlsOptions []func(*tls.Config)) (*webhook.Options, error) {
	var webhookOptions *webhook.Options
	var certWatcher *certwatcher.CertWatcher
	var err error

	webhookOptions = &webhook.Options{TLSOpts: tlsOptions}
	if len(options.CertPath) > 0 {
		certWatcher, err = certwatcher.New(
			filepath.Join(options.CertPath, options.CertName),
			filepath.Join(options.CertPath, options.CertKey),
		)
		if err != nil {
			return webhookOptions, err
		}
		webhookOptions.TLSOpts = append(
			webhookOptions.TLSOpts,
			func(config *tls.Config) {
				config.GetCertificate = certWatcher.GetCertificate
			},
		)
		lazy.RegisterWebhookCertWatcher(certWatcher)
	}

	return webhookOptions, nil
}
