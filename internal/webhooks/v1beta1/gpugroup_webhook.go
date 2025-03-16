/*Copyright 2025.

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
var gpugrouplog = logf.Log.WithName("gpugroup-resource")

// SetupGPUGroupWebhookWithManager registers the webhook for GPUGroup in the manager.
func (webhook *GPUGroup) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&vellunv1beta1.GPUGroup{}).
		WithValidator(webhook).
		WithDefaulter(webhook).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-vellun-io-v1beta1-gpugroup,mutating=true,failurePolicy=fail,sideEffects=None,groups=vellun.io,resources=gpugroups,verbs=create;update,versions=v1beta1,name=mgpugroup-v1beta1.kb.io,admissionReviewVersions=v1
// +kubebuilder:webhook:path=/validate-vellun-io-v1beta1-gpugroup,mutating=false,failurePolicy=fail,sideEffects=None,groups=vellun.io,resources=gpugroups,verbs=create;update,versions=v1beta1,name=vgpugroup-v1beta1.kb.io,admissionReviewVersions=v1

// GPUGroup implements a defaulting and validation webhook for GPUGroup
type GPUGroup struct{}

var _ webhook.CustomDefaulter = &GPUGroup{}
var _ webhook.CustomValidator = &GPUGroup{}

// Default implements webhook.CusotmDefaulter so a webhook will be registered for the Kind GPUGroup.
func (d *GPUGroup) Default(ctx context.Context, obj runtime.Object) error {
	gpugroup, ok := obj.(*vellunv1beta1.GPUGroup)

	if !ok {
		return fmt.Errorf("expected an GPUGroup object but got %T", obj)
	}
	gpugrouplog.Info("Defaulting for GPUGroup", "name", gpugroup.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// ValidateCreate implements webhoolhk.CustomValidator so a webhook will be registered for the type GPUGroup
func (v *GPUGroup) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	gpugroup, ok := obj.(*vellunv1beta1.GPUGroup)
	if !ok {
		return nil, fmt.Errorf("expected a GPUGroup object but got %T", obj)
	}
	gpugrouplog.Info("Validation for GPUGroup upon creation", "name", gpugroup.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type GPUGroup.
func (v *GPUGroup) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	gpugroup, ok := newObj.(*vellunv1beta1.GPUGroup)
	if !ok {
		return nil, fmt.Errorf("expected a GPUGroup object for the newObj but got %T", newObj)
	}
	gpugrouplog.Info("Validation for GPUGroup upon update", "name", gpugroup.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type GPUGroup.
func (v *GPUGroup) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	gpugroup, ok := obj.(*vellunv1beta1.GPUGroup)
	if !ok {
		return nil, fmt.Errorf("expected a GPUGroup object but got %T", obj)
	}
	gpugrouplog.Info("Validation for GPUGroup upon deletion", "name", gpugroup.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
