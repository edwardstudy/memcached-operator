package client

import "github.com/coreos/etcd-operator/pkg/util/k8sutil"

func MustNewInCluster() versioned.Interface {
	cfg, err := k8sutil.InClusterConfig()
	if err != nil {
		panic(err)
	}
	return MustNew(cfg)
}