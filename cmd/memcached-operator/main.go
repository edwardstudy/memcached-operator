package main

import (
	"context"
	"runtime"
	"time"

	"github.com/edwardstudy/memcached-operator/pkg/stub"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"

	"github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()

	sdk.ExposeMetricsPort()
	metrics, err := stub.RegisterOperatorMetrics()
	if err != nil {
		logrus.Errorf("Failed to register operator specific metrics: %v", err)
	}
	h := stub.NewHandler(metrics)

	resource := "cache.example.com/v1alpha1"
	kind := "Memcached"
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		logrus.Fatalf("Failed to get watch namespace: %v", err)
	}

	resyncPeriod := time.Duration(5) * time.Second
	logrus.Infof("Watching %s, %s, %s, %d", resource, kind, namespace, resyncPeriod)
	sdk.Watch(resource, kind, namespace, resyncPeriod)
	sdk.Handle(h)
	sdk.Run(context.TODO())

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Fatalf("Failed to get config for talking to the apiserver: %v", err)
	}

	// Create a new manager
	mgr, err := manager.New(cfg, manager.Options{Namespace: namespace})
	if err != nil {
		logrus.Fatalf("Failed to create a new manager: %v", err)
	}

	logrus.Infof("Display api server host", mgr.GetConfig().Host)


	// Setup Scheme for all resources

	// Setup all Controllers

	// Start the Cmd
}
