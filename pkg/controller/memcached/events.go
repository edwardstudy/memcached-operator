package memcached

import (
	"fmt"

	kwatch "k8s.io/apimachinery/pkg/watch"

	cachev1alpha1 "github.com/edwardstudy/memcached-operator/pkg/apis/cache/v1alpha1"
)

type Event struct {
	Type   kwatch.EventType
	Object *cachev1alpha1.Memcached
}

func handleMemcachedEvent(event *Event) (bool, error) {
	memcached := event.Object

	if len(memcached.Status.Nodes) == 0 {
		return false, fmt.Errorf("ignore failed memcahche (%s). Please delete its CR", memcached.Name)
	}

	switch event.Type {
	case kwatch.Added:
		// TODO create

	case kwatch.Modified:
		// TODO update

	case kwatch.Deleted:
		// TODO delete
	}
	return false, nil
}
