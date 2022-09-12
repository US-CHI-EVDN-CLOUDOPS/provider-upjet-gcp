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

type CertificateTemplateIdentityConstraintsObservation struct {
}

type CertificateTemplateIdentityConstraintsParameters struct {

	// Required. If this is true, the SubjectAltNames extension may be copied from a certificate request into the signed certificate. Otherwise, the requested SubjectAltNames will be discarded.
	// +kubebuilder:validation:Required
	AllowSubjectAltNamesPassthrough *bool `json:"allowSubjectAltNamesPassthrough" tf:"allow_subject_alt_names_passthrough,omitempty"`

	// Required. If this is true, the Subject field may be copied from a certificate request into the signed certificate. Otherwise, the requested Subject will be discarded.
	// +kubebuilder:validation:Required
	AllowSubjectPassthrough *bool `json:"allowSubjectPassthrough" tf:"allow_subject_passthrough,omitempty"`

	// Optional. A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a certificate is signed. To see the full allowed syntax and some examples, see https://cloud.google.com/certificate-authority-service/docs/using-cel
	// +kubebuilder:validation:Optional
	CelExpression []IdentityConstraintsCelExpressionParameters `json:"celExpression,omitempty" tf:"cel_expression,omitempty"`
}

type CertificateTemplateObservation struct {

	// Output only. The time at which this CertificateTemplate was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Output only. The time at which this CertificateTemplate was updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CertificateTemplateParameters struct {

	// Optional. A human-readable description of scenarios this template is intended for.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	// +kubebuilder:validation:Optional
	IdentityConstraints []CertificateTemplateIdentityConstraintsParameters `json:"identityConstraints,omitempty" tf:"identity_constraints,omitempty"`

	// Optional. Labels with user-defined metadata.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	// +kubebuilder:validation:Optional
	PassthroughExtensions []PassthroughExtensionsParameters `json:"passthroughExtensions,omitempty" tf:"passthrough_extensions,omitempty"`

	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
	// +kubebuilder:validation:Optional
	PredefinedValues []PredefinedValuesParameters `json:"predefinedValues,omitempty" tf:"predefined_values,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type IdentityConstraintsCelExpressionObservation struct {
}

type IdentityConstraintsCelExpressionParameters struct {

	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type PassthroughExtensionsAdditionalExtensionsObservation struct {
}

type PassthroughExtensionsAdditionalExtensionsParameters struct {

	// Required. The parts of an OID path. The most significant parts of the path come first.
	// +kubebuilder:validation:Required
	ObjectIDPath []*float64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

type PassthroughExtensionsObservation struct {
}

type PassthroughExtensionsParameters struct {

	// Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions.
	// +kubebuilder:validation:Optional
	AdditionalExtensions []PassthroughExtensionsAdditionalExtensionsParameters `json:"additionalExtensions,omitempty" tf:"additional_extensions,omitempty"`

	// Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
	// +kubebuilder:validation:Optional
	KnownExtensions []*string `json:"knownExtensions,omitempty" tf:"known_extensions,omitempty"`
}

type PredefinedValuesAdditionalExtensionsObjectIDObservation struct {
}

type PredefinedValuesAdditionalExtensionsObjectIDParameters struct {

	// Required. The parts of an OID path. The most significant parts of the path come first.
	// +kubebuilder:validation:Required
	ObjectIDPath []*float64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

type PredefinedValuesAdditionalExtensionsObservation struct {
}

type PredefinedValuesAdditionalExtensionsParameters struct {

	// Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
	// +kubebuilder:validation:Optional
	Critical *bool `json:"critical,omitempty" tf:"critical,omitempty"`

	// Required. The OID for this X.509 extension.
	// +kubebuilder:validation:Required
	ObjectID []PredefinedValuesAdditionalExtensionsObjectIDParameters `json:"objectId" tf:"object_id,omitempty"`

	// Required. The value of this X.509 extension.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PredefinedValuesCAOptionsObservation struct {
}

type PredefinedValuesCAOptionsParameters struct {

	// Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate.
	// +kubebuilder:validation:Optional
	IsCA *bool `json:"isCa,omitempty" tf:"is_ca,omitempty"`

	// Optional. Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate.
	// +kubebuilder:validation:Optional
	MaxIssuerPathLength *float64 `json:"maxIssuerPathLength,omitempty" tf:"max_issuer_path_length,omitempty"`
}

type PredefinedValuesKeyUsageBaseKeyUsageObservation struct {
}

type PredefinedValuesKeyUsageBaseKeyUsageParameters struct {

	// The key may be used to sign certificates.
	// +kubebuilder:validation:Optional
	CertSign *bool `json:"certSign,omitempty" tf:"cert_sign,omitempty"`

	// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	// +kubebuilder:validation:Optional
	ContentCommitment *bool `json:"contentCommitment,omitempty" tf:"content_commitment,omitempty"`

	// The key may be used sign certificate revocation lists.
	// +kubebuilder:validation:Optional
	CrlSign *bool `json:"crlSign,omitempty" tf:"crl_sign,omitempty"`

	// The key may be used to encipher data.
	// +kubebuilder:validation:Optional
	DataEncipherment *bool `json:"dataEncipherment,omitempty" tf:"data_encipherment,omitempty"`

	// The key may be used to decipher only.
	// +kubebuilder:validation:Optional
	DecipherOnly *bool `json:"decipherOnly,omitempty" tf:"decipher_only,omitempty"`

	// The key may be used for digital signatures.
	// +kubebuilder:validation:Optional
	DigitalSignature *bool `json:"digitalSignature,omitempty" tf:"digital_signature,omitempty"`

	// The key may be used to encipher only.
	// +kubebuilder:validation:Optional
	EncipherOnly *bool `json:"encipherOnly,omitempty" tf:"encipher_only,omitempty"`

	// The key may be used in a key agreement protocol.
	// +kubebuilder:validation:Optional
	KeyAgreement *bool `json:"keyAgreement,omitempty" tf:"key_agreement,omitempty"`

	// The key may be used to encipher other keys.
	// +kubebuilder:validation:Optional
	KeyEncipherment *bool `json:"keyEncipherment,omitempty" tf:"key_encipherment,omitempty"`
}

type PredefinedValuesKeyUsageExtendedKeyUsageObservation struct {
}

type PredefinedValuesKeyUsageExtendedKeyUsageParameters struct {

	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	// +kubebuilder:validation:Optional
	ClientAuth *bool `json:"clientAuth,omitempty" tf:"client_auth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	// +kubebuilder:validation:Optional
	CodeSigning *bool `json:"codeSigning,omitempty" tf:"code_signing,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	// +kubebuilder:validation:Optional
	EmailProtection *bool `json:"emailProtection,omitempty" tf:"email_protection,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	// +kubebuilder:validation:Optional
	OcspSigning *bool `json:"ocspSigning,omitempty" tf:"ocsp_signing,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	// +kubebuilder:validation:Optional
	ServerAuth *bool `json:"serverAuth,omitempty" tf:"server_auth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	// +kubebuilder:validation:Optional
	TimeStamping *bool `json:"timeStamping,omitempty" tf:"time_stamping,omitempty"`
}

type PredefinedValuesKeyUsageObservation struct {
}

type PredefinedValuesKeyUsageParameters struct {

	// Describes high-level ways in which a key may be used.
	// +kubebuilder:validation:Optional
	BaseKeyUsage []PredefinedValuesKeyUsageBaseKeyUsageParameters `json:"baseKeyUsage,omitempty" tf:"base_key_usage,omitempty"`

	// Detailed scenarios in which a key may be used.
	// +kubebuilder:validation:Optional
	ExtendedKeyUsage []PredefinedValuesKeyUsageExtendedKeyUsageParameters `json:"extendedKeyUsage,omitempty" tf:"extended_key_usage,omitempty"`

	// Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message.
	// +kubebuilder:validation:Optional
	UnknownExtendedKeyUsages []PredefinedValuesKeyUsageUnknownExtendedKeyUsagesParameters `json:"unknownExtendedKeyUsages,omitempty" tf:"unknown_extended_key_usages,omitempty"`
}

type PredefinedValuesKeyUsageUnknownExtendedKeyUsagesObservation struct {
}

type PredefinedValuesKeyUsageUnknownExtendedKeyUsagesParameters struct {

	// Required. The parts of an OID path. The most significant parts of the path come first.
	// +kubebuilder:validation:Required
	ObjectIDPath []*float64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

type PredefinedValuesObservation struct {
}

type PredefinedValuesParameters struct {

	// Optional. Describes custom X.509 extensions.
	// +kubebuilder:validation:Optional
	AdditionalExtensions []PredefinedValuesAdditionalExtensionsParameters `json:"additionalExtensions,omitempty" tf:"additional_extensions,omitempty"`

	// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	// +kubebuilder:validation:Optional
	AiaOcspServers []*string `json:"aiaOcspServers,omitempty" tf:"aia_ocsp_servers,omitempty"`

	// Optional. Describes options in this X509Parameters that are relevant in a CA certificate.
	// +kubebuilder:validation:Optional
	CAOptions []PredefinedValuesCAOptionsParameters `json:"caOptions,omitempty" tf:"ca_options,omitempty"`

	// Optional. Indicates the intended use for keys that correspond to a certificate.
	// +kubebuilder:validation:Optional
	KeyUsage []PredefinedValuesKeyUsageParameters `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +kubebuilder:validation:Optional
	PolicyIds []PredefinedValuesPolicyIdsParameters `json:"policyIds,omitempty" tf:"policy_ids,omitempty"`
}

type PredefinedValuesPolicyIdsObservation struct {
}

type PredefinedValuesPolicyIdsParameters struct {

	// Required. The parts of an OID path. The most significant parts of the path come first.
	// +kubebuilder:validation:Required
	ObjectIDPath []*float64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

// CertificateTemplateSpec defines the desired state of CertificateTemplate
type CertificateTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateTemplateParameters `json:"forProvider"`
}

// CertificateTemplateStatus defines the observed state of CertificateTemplate.
type CertificateTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateTemplate is the Schema for the CertificateTemplates API. Certificate Authority Service provides reusable and parameterized templates that you can use for common certificate issuance scenarios. A certificate template represents a relatively static and well-defined certificate issuance schema within an organization.  A certificate template can essentially become a full-fledged vertical certificate issuance framework.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CertificateTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateTemplateSpec   `json:"spec"`
	Status            CertificateTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateTemplateList contains a list of CertificateTemplates
type CertificateTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateTemplate `json:"items"`
}

// Repository type metadata.
var (
	CertificateTemplate_Kind             = "CertificateTemplate"
	CertificateTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateTemplate_Kind}.String()
	CertificateTemplate_KindAPIVersion   = CertificateTemplate_Kind + "." + CRDGroupVersion.String()
	CertificateTemplate_GroupVersionKind = CRDGroupVersion.WithKind(CertificateTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateTemplate{}, &CertificateTemplateList{})
}
