/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	vellunv1beta1 "github.com/aztecher/vellun/api/v1beta1"
)

// TODO: Need to change or remove this.
var gpunetworkpolicylog = logf.Log.WithName("gpunetworkpolicy-resource")

// SetupGPUNetworkPolicyWebhookWithManager registers the webhook for GPUNetworkPolicy in the manager.
func (webhook *GPUNetworkPolicy) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&vellunv1beta1.GPUNetworkPolicy{}).
		WithValidator(webhook).
		WithDefaulter(webhook).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-vellun-io-v1beta1-gpunetworkpolicy,mutating=true,failurePolicy=fail,sideEffects=None,groups=vellun.io,resources=gpunetworkpolicies,verbs=create;update,versions=v1beta1,name=mgpunetworkpolicy-v1beta1.kb.io,admissionReviewVersions=v1
// +kubebuilder:webhook:path=/validate-vellun-io-v1beta1-gpunetworkpolicy,mutating=false,failurePolicy=fail,sideEffects=None,groups=vellun.io,resources=gpunetworkpolicies,verbs=create;update,versions=v1beta1,name=vgpunetworkpolicy-v1beta1.kb.io,admissionReviewVersions=v1

// GPUNetworkPolicy implements a defaulting and validation webhoook forR GPUNetworkPolicy
type GPUNetworkPolicy struct{}

var _ webhook.CustomDefaulter = &GPUNetworkPolicy{}
var _ webhook.CustomValidator = &GPUNetworkPolicy{}

// Default implements webhook. so a webhook will be registered for the Kind GPUNetworkPolicy.
func (d *GPUNetworkPolicy) Default(ctx context.Context, obj runtime.Object) error {
	gpunetworkpolicy, ok := obj.(*vellunv1beta1.GPUNetworkPolicy)

	if !ok {
		return fmt.Errorf("expected an GPUNetworkPolicy object but got %T", obj)
	}
	gpunetworkpolicylog.Info("Defaulting for GPUNetworkPolicy", "name", gpunetworkpolicy.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type GPUNetworkPolicy.
func (v *GPUNetworkPolicy) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	gpunetworkpolicy, ok := obj.(*vellunv1beta1.GPUNetworkPolicy)
	if !ok {
		return nil, fmt.Errorf("expected a GPUNetworkPolicy object but got %T", obj)
	}
	gpunetworkpolicylog.Info("Validation for GPUNetworkPolicy upon creation", "name", gpunetworkpolicy.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type GPUNetworkPolicy.
func (v *GPUNetworkPolicy) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	gpunetworkpolicy, ok := newObj.(*vellunv1beta1.GPUNetworkPolicy)
	if !ok {
		return nil, fmt.Errorf("expected a GPUNetworkPolicy object for the newObj but got %T", newObj)
	}
	gpunetworkpolicylog.Info("Validation for GPUNetworkPolicy upon update", "name", gpunetworkpolicy.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type GPUNetworkPolicy.
func (v *GPUNetworkPolicy) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	gpunetworkpolicy, ok := obj.(*vellunv1beta1.GPUNetworkPolicy)
	if !ok {
		return nil, fmt.Errorf("expected a GPUNetworkPolicy object but got %T", obj)
	}
	gpunetworkpolicylog.Info("Validation for GPUNetworkPolicy upon deletion", "name", gpunetworkpolicy.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
