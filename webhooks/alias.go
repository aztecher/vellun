package webhooks

import (
	webhooksv1beta1 "github.com/aztecher/vellun/internal/webhooks/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// GPUGroup implements a defaulting and validating webhook for GPUGroup
type GPUGroup struct{}

// SetupWebhookWithManager sets up GPUGroup webhooks.
func (webhook *GPUGroup) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return (&webhooksv1beta1.GPUGroup{}).SetupWebhookWithManager(mgr)
}

// GPUNetworkPolicy implementss a defaulting and validating webhook for GPUNetworkPolicy
type GPUNetworkPolicy struct{}

// SetupWebhookWithManager sets up GPUNetworkPolicy webhooks.
func (webhook *GPUNetworkPolicy) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return (&webhooksv1beta1.GPUNetworkPolicy{}).SetupWebhookWithManager(mgr)
}
