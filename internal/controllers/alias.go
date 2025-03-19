package controllers

import (
	"context"

	gpugroupcontroller "github.com/aztecher/vellun/internal/controllers/gpugroup"
	gpunetworkpolicycontroller "github.com/aztecher/vellun/internal/controllers/gpunetworkpolicy"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

// GPUGroupReconciler reconciles a GPUGroup object
type GPUGroupReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *GPUGroupReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	return (&gpugroupcontroller.Reconciler{
		Client: r.Client,
		Scheme: r.Scheme,
	}).SetupWithManager(mgr)

}

// GPUNetworkPolicyReconciler reconciles a GPUNetworkPolicy object
type GPUNetworkPolicyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *GPUNetworkPolicyReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	return (&gpunetworkpolicycontroller.Reconciler{
		Client: r.Client,
		Scheme: r.Scheme,
	}).SetupWithManager(mgr)
}
