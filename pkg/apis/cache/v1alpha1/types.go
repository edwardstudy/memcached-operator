package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Memcached `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              MemcachedSpec   `json:"spec"`
	Status            MemcachedStatus `json:"status,omitempty"`
}

// After modifying the *_types.go file always run the following command to update the generated code for that resource type:
//
// $ operator-sdk generate k8s

type MemcachedSpec struct {
	// Size is the size of the memcached deployment
	Size int32 `json:"size"`
}
type MemcachedStatus struct {
	// Nodes are the names of the memcached pods
	Nodes []string `json:"nodes"`
}
