// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BackendBucketSignedURLKeyInitParameters struct {

	// The backend bucket this signed URL key belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta2.BackendBucket
	BackendBucket *string `json:"backendBucket,omitempty" tf:"backend_bucket,omitempty"`

	// Reference to a BackendBucket in compute to populate backendBucket.
	// +kubebuilder:validation:Optional
	BackendBucketRef *v1.Reference `json:"backendBucketRef,omitempty" tf:"-"`

	// Selector for a BackendBucket in compute to populate backendBucket.
	// +kubebuilder:validation:Optional
	BackendBucketSelector *v1.Selector `json:"backendBucketSelector,omitempty" tf:"-"`

	// Name of the signed URL key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BackendBucketSignedURLKeyObservation struct {

	// The backend bucket this signed URL key belongs.
	BackendBucket *string `json:"backendBucket,omitempty" tf:"backend_bucket,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/backendBuckets/{{backend_bucket}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the signed URL key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BackendBucketSignedURLKeyParameters struct {

	// The backend bucket this signed URL key belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta2.BackendBucket
	// +kubebuilder:validation:Optional
	BackendBucket *string `json:"backendBucket,omitempty" tf:"backend_bucket,omitempty"`

	// Reference to a BackendBucket in compute to populate backendBucket.
	// +kubebuilder:validation:Optional
	BackendBucketRef *v1.Reference `json:"backendBucketRef,omitempty" tf:"-"`

	// Selector for a BackendBucket in compute to populate backendBucket.
	// +kubebuilder:validation:Optional
	BackendBucketSelector *v1.Selector `json:"backendBucketSelector,omitempty" tf:"-"`

	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	KeyValueSecretRef v1.SecretKeySelector `json:"keyValueSecretRef" tf:"-"`

	// Name of the signed URL key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// BackendBucketSignedURLKeySpec defines the desired state of BackendBucketSignedURLKey
type BackendBucketSignedURLKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendBucketSignedURLKeyParameters `json:"forProvider"`
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
	InitProvider BackendBucketSignedURLKeyInitParameters `json:"initProvider,omitempty"`
}

// BackendBucketSignedURLKeyStatus defines the observed state of BackendBucketSignedURLKey.
type BackendBucketSignedURLKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendBucketSignedURLKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackendBucketSignedURLKey is the Schema for the BackendBucketSignedURLKeys API. A key for signing Cloud CDN signed URLs for BackendBuckets.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type BackendBucketSignedURLKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keyValueSecretRef)",message="spec.forProvider.keyValueSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   BackendBucketSignedURLKeySpec   `json:"spec"`
	Status BackendBucketSignedURLKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendBucketSignedURLKeyList contains a list of BackendBucketSignedURLKeys
type BackendBucketSignedURLKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendBucketSignedURLKey `json:"items"`
}

// Repository type metadata.
var (
	BackendBucketSignedURLKey_Kind             = "BackendBucketSignedURLKey"
	BackendBucketSignedURLKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendBucketSignedURLKey_Kind}.String()
	BackendBucketSignedURLKey_KindAPIVersion   = BackendBucketSignedURLKey_Kind + "." + CRDGroupVersion.String()
	BackendBucketSignedURLKey_GroupVersionKind = CRDGroupVersion.WithKind(BackendBucketSignedURLKey_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendBucketSignedURLKey{}, &BackendBucketSignedURLKeyList{})
}
