/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type KvObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type KvParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceID *string `json:"namespaceId" tf:"namespace_id,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// KvSpec defines the desired state of Kv
type KvSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KvParameters `json:"forProvider"`
}

// KvStatus defines the observed state of Kv.
type KvStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KvObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Kv is the Schema for the Kvs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflarejet}
type Kv struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KvSpec   `json:"spec"`
	Status            KvStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KvList contains a list of Kvs
type KvList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kv `json:"items"`
}

// Repository type metadata.
var (
	Kv_Kind             = "Kv"
	Kv_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Kv_Kind}.String()
	Kv_KindAPIVersion   = Kv_Kind + "." + CRDGroupVersion.String()
	Kv_GroupVersionKind = CRDGroupVersion.WithKind(Kv_Kind)
)

func init() {
	SchemeBuilder.Register(&Kv{}, &KvList{})
}
