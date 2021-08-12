/*
Copyright 2021 The Kruise Authors.

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

package v1alpha1

import (
	csv1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FederatedCloneSetSpec defines the desired state of FederatedCloneSet
type FederatedCloneSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Total number of pods desired across federated clusters.
	// Replicas specified in the spec for target deployment template or replicaset
	// template will be discarded/overridden when scheduling preferences are
	// specified.
	TotalReplicas int32 `json:"totalReplicas"`

	// If set to true then already scheduled and running replicas may be moved to other clusters
	// in order to match current state to the specified preferences. Otherwise, if set to false,
	// up and running replicas will not be moved.
	// +optional
	Rebalance bool `json:"rebalance,omitempty"`

	// If set to true, the placement of target kind will be determined using the instersection
	// of RSP placement scheduling result and the clusterSelector (spec.placement.clusterSelector)
	// specified on the target kind.
	// If set to false or not defined, RSP placement scheduling result overwrites the clusters
	// list in the spec.placement.clusters of the target resource.
	// +optional
	IntersectWithClusterSelector bool `json:"intersectWithClusterSelector"`

	// A mapping between cluster names and preferences regarding a local workload object (dep, rs, .. ) in
	// these clusters.
	// "*" (if provided) applies to all clusters if an explicit mapping is not provided.
	// If omitted, clusters without explicit preferences should not have any replicas scheduled.
	// +optional
	Clusters map[string]ClusterPreferences `json:"clusters,omitempty"`

	// To tell CloneSet in the clusters what to do
	// +optional
	CloneSetSpec csv1.CloneSetSpec `json:"cloneSetSpec"`
}

type ClusterPreferences struct {
	// Minimum number of replicas that should be assigned to this cluster workload object. 0 by default.
	// +optional
	MinReplicas int64 `json:"minReplicas,omitempty"`

	// Maximum number of replicas that should be assigned to this cluster workload object.
	// Unbounded if no value provided (default).
	// +optional
	MaxReplicas *int64 `json:"maxReplicas,omitempty"`

	// A number expressing the preference to put an additional replica to this cluster workload object.
	// 0 by default.
	Weight int64 `json:"weight,omitempty"`
}

// FederatedCloneSetStatus defines the observed state of FederatedCloneSet
type FederatedCloneSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FederatedCloneSet is the Schema for the federatedclonesets API
type FederatedCloneSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedCloneSetSpec   `json:"spec,omitempty"`
	Status FederatedCloneSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FederatedCloneSetList contains a list of FederatedCloneSet
type FederatedCloneSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedCloneSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FederatedCloneSet{}, &FederatedCloneSetList{})
}
