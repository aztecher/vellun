package flags

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/metrics/filters"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func TestGetManagerOptions(t *testing.T) {
	tests := []struct {
		name            string
		mgrOpts         ManagerOptions
		wantErr         bool
		wantWebhookOpts *webhook.Options
		wantMetricsOpts *metricsserver.Options
	}{
		{
			name:            "valid, no options, use defaults",
			mgrOpts:         NewManagerOptions(),
			wantErr:         false,
			wantWebhookOpts: &webhook.Options{TLSOpts: nil},
			wantMetricsOpts: &metricsserver.Options{TLSOpts: nil},
		},
		{
			name: "invalid, webhook certs not found",
			mgrOpts: ManagerOptions{
				Webhook: &WebhookOptions{
					CertPath: ".",
					CertName: "webhook.notfound.crt",
					CertKey:  "webhook.notfound.key",
				},
				Metrics: &MetricsOptions{},
			},
			wantErr: true,
		},
		{
			name: "invalid, metrics certs not found",
			mgrOpts: ManagerOptions{
				Webhook: &WebhookOptions{},
				Metrics: &MetricsOptions{
					CertPath: ".",
					CertName: "metrics.notfound.crt",
					CertKey:  "metrics.notfound.key",
				},
			},
			wantErr: true,
		},
		{
			name: "valid, secure metrics, metrics address",
			mgrOpts: ManagerOptions{
				Webhook: &WebhookOptions{},
				Metrics: &MetricsOptions{
					BindAddress:   ":8443",
					SecureServing: true,
				},
			},
			wantErr:         false,
			wantWebhookOpts: &webhook.Options{TLSOpts: nil},
			wantMetricsOpts: &metricsserver.Options{
				BindAddress:    ":8443",
				SecureServing:  true,
				TLSOpts:        nil,
				FilterProvider: filters.WithAuthenticationAndAuthorization,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			webhookOpts, metricsOpts, err := GetManagerOptions(tt.mgrOpts)
			if tt.wantErr {
				g.Expect(err).To(HaveOccurred())
				return
			}
			g.Expect(err).NotTo(HaveOccurred())

			g.Expect(webhookOpts).To(Equal(tt.wantWebhookOpts))
			// There is no way we cn compare the FilterProvider
			g.Expect(metricsOpts.FilterProvider == nil).To(Equal(tt.wantMetricsOpts.FilterProvider == nil))
			metricsOpts.FilterProvider = nil
			tt.wantMetricsOpts.FilterProvider = nil

			g.Expect(metricsOpts).To(Equal(tt.wantMetricsOpts))
		})
	}
}
