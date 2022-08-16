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

type NetworkPeeringObservation struct {

	// an identifier for the resource with format {{network}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// State for the peering, either ACTIVE or INACTIVE. The peering is
	// ACTIVE when there's a matching configuration in the peer network.
	// State for the peering, either ACTIVE or INACTIVE. The peering is ACTIVE when there's a matching configuration in the peer network.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Details about the current state of the peering.
	// Details about the current state of the peering.
	StateDetails *string `json:"stateDetails,omitempty" tf:"state_details,omitempty"`
}

type NetworkPeeringParameters struct {

	// Whether to export the custom routes to the peer network. Defaults to false.
	// Whether to export the custom routes to the peer network. Defaults to false.
	// +kubebuilder:validation:Optional
	ExportCustomRoutes *bool `json:"exportCustomRoutes,omitempty" tf:"export_custom_routes,omitempty"`

	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges  are always exported to peers and are not controlled by this field.
	// +kubebuilder:validation:Optional
	ExportSubnetRoutesWithPublicIP *bool `json:"exportSubnetRoutesWithPublicIp,omitempty" tf:"export_subnet_routes_with_public_ip,omitempty"`

	// Whether to import the custom routes from the peer network. Defaults to false.
	// Whether to export the custom routes from the peer network. Defaults to false.
	// +kubebuilder:validation:Optional
	ImportCustomRoutes *bool `json:"importCustomRoutes,omitempty" tf:"import_custom_routes,omitempty"`

	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges  are always imported from peers and are not controlled by this field.
	// +kubebuilder:validation:Optional
	ImportSubnetRoutesWithPublicIP *bool `json:"importSubnetRoutesWithPublicIp,omitempty" tf:"import_subnet_routes_with_public_ip,omitempty"`

	// The primary network of the peering.
	// The primary network of the peering.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The peer network in the peering. The peer network
	// may belong to a different project.
	// The peer network in the peering. The peer network may belong to a different project.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	PeerNetwork *string `json:"peerNetwork,omitempty" tf:"peer_network,omitempty"`

	// +kubebuilder:validation:Optional
	PeerNetworkRef *v1.Reference `json:"peerNetworkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PeerNetworkSelector *v1.Selector `json:"peerNetworkSelector,omitempty" tf:"-"`
}

// NetworkPeeringSpec defines the desired state of NetworkPeering
type NetworkPeeringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkPeeringParameters `json:"forProvider"`
}

// NetworkPeeringStatus defines the observed state of NetworkPeering.
type NetworkPeeringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkPeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPeering is the Schema for the NetworkPeerings API. Manages a network peering within GCE.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPeeringSpec   `json:"spec"`
	Status            NetworkPeeringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPeeringList contains a list of NetworkPeerings
type NetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkPeering `json:"items"`
}

// Repository type metadata.
var (
	NetworkPeering_Kind             = "NetworkPeering"
	NetworkPeering_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkPeering_Kind}.String()
	NetworkPeering_KindAPIVersion   = NetworkPeering_Kind + "." + CRDGroupVersion.String()
	NetworkPeering_GroupVersionKind = CRDGroupVersion.WithKind(NetworkPeering_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkPeering{}, &NetworkPeeringList{})
}
