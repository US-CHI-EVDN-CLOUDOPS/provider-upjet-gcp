// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AgentPoolInitParameters struct {

	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	BandwidthLimit *BandwidthLimitInitParameters `json:"bandwidthLimit,omitempty" tf:"bandwidth_limit,omitempty"`

	// Specifies the client-specified AgentPool description.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type AgentPoolObservation struct {

	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	BandwidthLimit *BandwidthLimitObservation `json:"bandwidthLimit,omitempty" tf:"bandwidth_limit,omitempty"`

	// Specifies the client-specified AgentPool description.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/agentPools/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies the state of the AgentPool.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type AgentPoolParameters struct {

	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BandwidthLimit *BandwidthLimitParameters `json:"bandwidthLimit,omitempty" tf:"bandwidth_limit,omitempty"`

	// Specifies the client-specified AgentPool description.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BandwidthLimitInitParameters struct {

	// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
	LimitMbps *string `json:"limitMbps,omitempty" tf:"limit_mbps,omitempty"`
}

type BandwidthLimitObservation struct {

	// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
	LimitMbps *string `json:"limitMbps,omitempty" tf:"limit_mbps,omitempty"`
}

type BandwidthLimitParameters struct {

	// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
	// +kubebuilder:validation:Optional
	LimitMbps *string `json:"limitMbps" tf:"limit_mbps,omitempty"`
}

// AgentPoolSpec defines the desired state of AgentPool
type AgentPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgentPoolParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AgentPoolInitParameters `json:"initProvider,omitempty"`
}

// AgentPoolStatus defines the observed state of AgentPool.
type AgentPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgentPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AgentPool is the Schema for the AgentPools API. Represents an On-Premises Agent pool.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AgentPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AgentPoolSpec   `json:"spec"`
	Status            AgentPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentPoolList contains a list of AgentPools
type AgentPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentPool `json:"items"`
}

// Repository type metadata.
var (
	AgentPool_Kind             = "AgentPool"
	AgentPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AgentPool_Kind}.String()
	AgentPool_KindAPIVersion   = AgentPool_Kind + "." + CRDGroupVersion.String()
	AgentPool_GroupVersionKind = CRDGroupVersion.WithKind(AgentPool_Kind)
)

func init() {
	SchemeBuilder.Register(&AgentPool{}, &AgentPoolList{})
}
