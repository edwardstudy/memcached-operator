package controller

import (
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/edwardstudy/memcached-operator/pkg/client/clientset/versioned"
)

type Controller struct {
	Config
}

type Config struct {
	NameSpace      string
	MemcachedCRCli versioned.Interface
	KubeCli        kubernetes.Interface
	KubeExtCli     apiextensionsclient.Interface
}

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager) error

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager) error {
	for _, f := range AddToManagerFuncs {
		if err := f(m); err != nil {
			return err
		}
	}
	return nil
}
