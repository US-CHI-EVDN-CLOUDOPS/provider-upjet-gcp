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

type ConditionsObservation struct {

	// Human readable message indicating details about the current status.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// One-word CamelCase reason for the condition's current status.
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Resource record type. Example: AAAA.
	// Possible values are A, AAAA, and CNAME.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConditionsParameters struct {
}

type DomainMappingObservation struct {

	// an identifier for the resource with format locations/{{location}}/namespaces/{{project}}/domainmappings/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type DomainMappingParameters struct {

	// The location of the cloud run instance. eg us-central1
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Metadata []MetadataParameters `json:"metadata" tf:"metadata,omitempty"`

	// Name should be a verified domain
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The spec for this DomainMapping.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Spec []SpecParameters `json:"spec" tf:"spec,omitempty"`
}

type MetadataObservation struct {

	// A sequence number representing a specific generation of the desired state.
	Generation *float64 `json:"generation,omitempty" tf:"generation,omitempty"`

	// An opaque value that represents the internal version of this object that
	// can be used by clients to determine when objects have changed. May be used
	// for optimistic concurrency, change detection, and the watch operation on a
	// resource or set of resources. They may only be valid for a
	// particular resource or set of resources.
	// More info:
	// https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version,omitempty"`

	// SelfLink is a URL representing this object.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// UID is a unique id generated by the server on successful creation of a resource and is not
	// allowed to change on PUT operations.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type MetadataParameters struct {

	// Annotations is a key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata. More
	// info: http://kubernetes.io/docs/user-guide/annotations
	// Note: The Cloud Run API may add additional annotations that were not provided in your config.ignore_changes rule to the metadata.0.annotations field.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and routes.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// In Cloud Run the namespace must be equal to either the
	// project ID or project number.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Reference to a Project in cloudplatform to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceRef *v1.Reference `json:"namespaceRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceSelector *v1.Selector `json:"namespaceSelector,omitempty" tf:"-"`
}

type ResourceRecordsObservation struct {

	// Relative name of the object affected by this record. Only applicable for
	// CNAME records. Example: 'www'.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Data for this record. Values vary by record type, as defined in RFC 1035
	// (section 5) and RFC 1034 (section 3.6.1).
	Rrdata *string `json:"rrdata,omitempty" tf:"rrdata,omitempty"`

	// Resource record type. Example: AAAA.
	// Possible values are A, AAAA, and CNAME.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceRecordsParameters struct {
}

type SpecObservation struct {
}

type SpecParameters struct {

	// The mode of the certificate.
	// Default value is AUTOMATIC.
	// Possible values are NONE and AUTOMATIC.
	// +kubebuilder:validation:Optional
	CertificateMode *string `json:"certificateMode,omitempty" tf:"certificate_mode,omitempty"`

	// If set, the mapping will override any mapping set before this spec was set.
	// It is recommended that the user leaves this empty to receive an error
	// warning about a potential conflict and only set it once the respective UI
	// has given such a warning.
	// +kubebuilder:validation:Optional
	ForceOverride *bool `json:"forceOverride,omitempty" tf:"force_override,omitempty"`

	// The name of the Cloud Run Service that this DomainMapping applies to.
	// The route must exist.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	RouteName *string `json:"routeName,omitempty" tf:"route_name,omitempty"`

	// Reference to a Service to populate routeName.
	// +kubebuilder:validation:Optional
	RouteNameRef *v1.Reference `json:"routeNameRef,omitempty" tf:"-"`

	// Selector for a Service to populate routeName.
	// +kubebuilder:validation:Optional
	RouteNameSelector *v1.Selector `json:"routeNameSelector,omitempty" tf:"-"`
}

type StatusObservation struct {

	// Array of observed DomainMappingConditions, indicating the current state
	// of the DomainMapping.
	// Structure is documented below.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Name should be a verified domain
	MappedRouteName *string `json:"mappedRouteName,omitempty" tf:"mapped_route_name,omitempty"`

	// ObservedGeneration is the 'Generation' of the DomainMapping that
	// was last processed by the controller.
	ObservedGeneration *float64 `json:"observedGeneration,omitempty" tf:"observed_generation,omitempty"`

	// The resource records required to configure this domain mapping. These
	// records must be added to the domain's DNS configuration in order to
	// serve the application via this domain mapping.
	// Structure is documented below.
	ResourceRecords []ResourceRecordsObservation `json:"resourceRecords,omitempty" tf:"resource_records,omitempty"`
}

type StatusParameters struct {
}

// DomainMappingSpec defines the desired state of DomainMapping
type DomainMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainMappingParameters `json:"forProvider"`
}

// DomainMappingStatus defines the observed state of DomainMapping.
type DomainMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMapping is the Schema for the DomainMappings API. Resource to hold the state and status of a user's domain mapping.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DomainMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainMappingSpec   `json:"spec"`
	Status            DomainMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMappingList contains a list of DomainMappings
type DomainMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainMapping `json:"items"`
}

// Repository type metadata.
var (
	DomainMapping_Kind             = "DomainMapping"
	DomainMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainMapping_Kind}.String()
	DomainMapping_KindAPIVersion   = DomainMapping_Kind + "." + CRDGroupVersion.String()
	DomainMapping_GroupVersionKind = CRDGroupVersion.WithKind(DomainMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainMapping{}, &DomainMappingList{})
}
