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

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	vellunv1beta1 "github.com/aztecher/vellun/api/v1beta1"
	"github.com/aztecher/vellun/internal/controllers"
	"github.com/aztecher/vellun/pkg/version"
	"github.com/aztecher/vellun/util/flags"
	"github.com/aztecher/vellun/util/lazy"
	"github.com/aztecher/vellun/webhooks"
	"github.com/spf13/pflag"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")

	managerOptions = flags.NewManagerOptions()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(vellunv1beta1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

// InitFlags initializes the flags.
func InitFlags(fs *pflag.FlagSet) {
	flags.AddManagerOptions(fs, &managerOptions)
}

// nolint:gocyclo
func main() {
	setupLog.Info(fmt.Sprintf("Version: %+v", version.Version))
	InitFlags(pflag.CommandLine)

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	webhookOptions, metricsOptions, err := flags.GetManagerOptions(managerOptions)
	if err != nil {
		setupLog.Error(err, "unable to construct manager options")
		os.Exit(1)
	}

	// Configure restConfig
	restConfig := ctrl.GetConfigOrDie()

	// Configure ctrl.Options
	ctrlOptions := ctrl.Options{
		Scheme:                 scheme,
		HealthProbeBindAddress: managerOptions.ProveAddr,
		LeaderElection:         managerOptions.EnableLeaderElection,
		LeaderElectionID:       "controller-leader-election-vellun",
		Metrics:                *metricsOptions,
		WebhookServer:          webhook.NewServer(*webhookOptions),
	}

	mgr, err := ctrl.NewManager(restConfig, ctrlOptions)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// Setup the context that's going to be used in controllers and for the manager.
	ctx := ctrl.SetupSignalHandler()
	setupChecks(mgr)
	// TODO: setupIndex
	setupReconcilers(ctx, mgr, managerOptions)
	setupWebhooks(mgr)
	setupCertWatchers(mgr)

	setupLog.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func setupChecks(mgr ctrl.Manager) {
	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}
}

func setupWebhooks(mgr ctrl.Manager) {
	if err := (&webhooks.GPUGroup{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "Unable to create webhook", "webhook", "GPUGroup")
		os.Exit(1)
	}
	if err := (&webhooks.GPUNetworkPolicy{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "GPUNetworkPolicy")
		os.Exit(1)
	}
}

func setupCertWatchers(mgr ctrl.Manager) {
	if err := lazy.BindMetricsCertWatcher(mgr); err != nil {
		setupLog.Error(err, "unable to add metrics certificate watcher to manager")
		os.Exit(1)
	}

	if err := lazy.BindWebhookCertWatcher(mgr); err != nil {
		setupLog.Error(err, "unable to add webhook certificate watcher to manager")
		os.Exit(1)
	}
}

func setupReconcilers(ctx context.Context, mgr ctrl.Manager, options flags.ManagerOptions) {
	var err error
	if err = (&controllers.GPUGroupReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(ctx, mgr, concurrency(options.GPUGroupConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "GPUGroup")
		os.Exit(1)
	}
	if err = (&controllers.GPUNetworkPolicyReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(ctx, mgr, concurrency(options.GPUNetworkPolicyConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "GPUNetworkPolicy")
		os.Exit(1)
	}
}

func concurrency(c int) controller.Options {
	return controller.Options{MaxConcurrentReconciles: c}
}
