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

type CronTriggerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CronTriggerParameters struct {

	// The account identifier to target for the resource.
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Required
	Schedules []*string `json:"schedules" tf:"schedules,omitempty"`

	// +kubebuilder:validation:Required
	ScriptName *string `json:"scriptName" tf:"script_name,omitempty"`
}

// CronTriggerSpec defines the desired state of CronTrigger
type CronTriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CronTriggerParameters `json:"forProvider"`
}

// CronTriggerStatus defines the observed state of CronTrigger.
type CronTriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CronTriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CronTrigger is the Schema for the CronTriggers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflarejet}
type CronTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CronTriggerSpec   `json:"spec"`
	Status            CronTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronTriggerList contains a list of CronTriggers
type CronTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronTrigger `json:"items"`
}

// Repository type metadata.
var (
	CronTrigger_Kind             = "CronTrigger"
	CronTrigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CronTrigger_Kind}.String()
	CronTrigger_KindAPIVersion   = CronTrigger_Kind + "." + CRDGroupVersion.String()
	CronTrigger_GroupVersionKind = CRDGroupVersion.WithKind(CronTrigger_Kind)
)

func init() {
	SchemeBuilder.Register(&CronTrigger{}, &CronTriggerList{})
}
