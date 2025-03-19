package lazy

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/certwatcher"
)

var (
	metricsCertWatcher *certwatcher.CertWatcher
	webhookCertWatcher *certwatcher.CertWatcher
)

func RegisterMetricsCertWatcher(watcher *certwatcher.CertWatcher) {
	metricsCertWatcher = watcher
}

func RegisterWebhookCertWatcher(watcher *certwatcher.CertWatcher) {
	webhookCertWatcher = watcher
}

func BindMetricsCertWatcher(mgr ctrl.Manager) error {
	if metricsCertWatcher != nil {
		if err := mgr.Add(metricsCertWatcher); err != nil {
			return err
		}
	}
	return nil
}

func BindWebhookCertWatcher(mgr ctrl.Manager) error {
	if webhookCertWatcher != nil {
		if err := mgr.Add(webhookCertWatcher); err != nil {
			return err
		}
	}
	return nil

}
