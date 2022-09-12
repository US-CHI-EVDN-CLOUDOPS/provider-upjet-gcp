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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AlternativeNameServerConfigObservation struct {
}

type AlternativeNameServerConfigParameters struct {

	// Sets an alternative name server for the associated networks. When specified,
	// all DNS queries are forwarded to a name server that you choose. Names such as .internal
	// are not available when an alternative name server is specified.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	TargetNameServers []AlternativeNameServerConfigTargetNameServersParameters `json:"targetNameServers" tf:"target_name_servers,omitempty"`
}

type AlternativeNameServerConfigTargetNameServersObservation struct {
}

type AlternativeNameServerConfigTargetNameServersParameters struct {

	// Forwarding path for this TargetNameServer. If unset or default Cloud DNS will make forwarding
	// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	// to the Internet. When set to private, Cloud DNS will always send queries through VPC for this target
	// Possible values are default and private.
	// +kubebuilder:validation:Optional
	ForwardingPath *string `json:"forwardingPath,omitempty" tf:"forwarding_path,omitempty"`

	// IPv4 address to forward to.
	// +kubebuilder:validation:Required
	IPv4Address *string `json:"ipv4Address" tf:"ipv4_address,omitempty"`
}

type PolicyNetworksObservation struct {
}

type PolicyNetworksParameters struct {

	// Reference to a Network in compute to populate networkUrl.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate networkUrl.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The id or fully qualified URL of the VPC network to forward queries to.
	// This should be formatted like projects/{project}/global/networks/{network} or
	// https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-gcp/config/common.ExtractResourceID()
	// +crossplane:generate:reference:refFieldName=NetworkRef
	// +crossplane:generate:reference:selectorFieldName=NetworkSelector
	// +kubebuilder:validation:Optional
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`
}

type PolicyObservation struct {

	// an identifier for the resource with format projects/{{project}}/policies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyParameters struct {

	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AlternativeNameServerConfig []AlternativeNameServerConfigParameters `json:"alternativeNameServerConfig,omitempty" tf:"alternative_name_server_config,omitempty"`

	// A textual description field.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	// +kubebuilder:validation:Optional
	EnableInboundForwarding *bool `json:"enableInboundForwarding,omitempty" tf:"enable_inbound_forwarding,omitempty"`

	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	// +kubebuilder:validation:Optional
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Networks []PolicyNetworksParameters `json:"networks,omitempty" tf:"networks,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. A policy is a collection of DNS rules applied to one or more Virtual Private Cloud resources.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
