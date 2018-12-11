package controller

import (
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

func Start(){
	source := cache.NewListWatchFromClient(
		c.Config.EtcdCRCli.EtcdV1beta2().RESTClient(),
		api.EtcdClusterResourcePlural,
		ns,
		fields.Everything())
}