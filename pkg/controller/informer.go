package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"

	"github.com/edwardstudy/memcached-operator/pkg/apis/cache/v1alpha1"
)

func Start() {
	// TODO Init CRD

	// TODO Run controller
}

func (c *Controller) run() {
	listWatch := cache.NewListWatchFromClient(
		c.Config.MemcachedCRCli.Cache(),
		"",
		c.Config.NameSpace,
		fields.Everything(),
	)

	_, informer := cache.NewIndexerInformer(
		listWatch,
		v1alpha1.Memcached{},
		0, // DO not reenqueue
		cache.ResourceEventHandlerFuncs{},
		cache.Indexers{},
	)

	informer.Run(context.TODO().Done())
	// TODO Use workqueue
}
