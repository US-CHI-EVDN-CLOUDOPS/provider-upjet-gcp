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

type AuthInfoObservation struct {
}

type AuthInfoParameters struct {

	// The password to authenticate.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username to authenticate.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type ContentMatchersObservation struct {
}

type ContentMatchersParameters struct {

	// String or regex content to match (max 1024 bytes)
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
	// Default value is CONTAINS_STRING.
	// Possible values are CONTAINS_STRING, NOT_CONTAINS_STRING, MATCHES_REGEX, and NOT_MATCHES_REGEX.
	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`
}

type HTTPCheckObservation struct {
}

type HTTPCheckParameters struct {

	// The authentication information. Optional when creating an HTTP check; defaults to empty.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AuthInfo []AuthInfoParameters `json:"authInfo,omitempty" tf:"auth_info,omitempty"`

	// The request body associated with the HTTP POST request. If contentType is URL_ENCODED, the body passed in must be URL-encoded. Users can provide a Content-Length header via the headers field or the API will do so. If the requestMethod is GET and body is not empty, the API will return an error. The maximum byte size is 1 megabyte. Note - As with all bytes fields JSON representations are base64 encoded. e.g. "foo=bar" in URL-encoded form is "foo%3Dbar" and in base64 encoding is "Zm9vJTI1M0RiYXI=".
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The content type to use for the check.
	// Possible values are TYPE_UNSPECIFIED and URL_ENCODED.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
	// +kubebuilder:validation:Optional
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.
	// +kubebuilder:validation:Optional
	MaskHeaders *bool `json:"maskHeaders,omitempty" tf:"mask_headers,omitempty"`

	// The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. If the provided path does not begin with "/", a "/" will be prepended automatically. Optional (defaults to "/").
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) and path to construct the full URL. Optional (defaults to 80 without SSL, or 443 with SSL).
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The HTTP request method to use for the check. If set to METHOD_UNSPECIFIED then requestMethod defaults to GET.
	// Default value is GET.
	// Possible values are METHOD_UNSPECIFIED, GET, and POST.
	// +kubebuilder:validation:Optional
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method,omitempty"`

	// If true, use HTTPS instead of HTTP to run the check.
	// +kubebuilder:validation:Optional
	UseSSL *bool `json:"useSsl,omitempty" tf:"use_ssl,omitempty"`

	// Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.
	// +kubebuilder:validation:Optional
	ValidateSSL *bool `json:"validateSsl,omitempty" tf:"validate_ssl,omitempty"`
}

type MonitoredResourceObservation struct {
}

type MonitoredResourceParameters struct {

	// Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels "project_id", "instance_id", and "zone".
	// +kubebuilder:validation:Required
	Labels map[string]*string `json:"labels" tf:"labels,omitempty"`

	// The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ResourceGroupObservation struct {
}

type ResourceGroupParameters struct {

	// The group of resources being monitored. Should be the name of a group
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The resource type of the group members.
	// Possible values are RESOURCE_TYPE_UNSPECIFIED, INSTANCE, and AWS_ELB_LOAD_BALANCER.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type TCPCheckObservation struct {
}

type TCPCheckParameters struct {

	// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`
}

type UptimeCheckConfigObservation struct {

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique resource name for this UptimeCheckConfig. The format is projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the uptime check
	UptimeCheckID *string `json:"uptimeCheckId,omitempty" tf:"uptime_check_id,omitempty"`
}

type UptimeCheckConfigParameters struct {

	// The checker type to use for the check. If the monitored resource type is servicedirectory_service, checkerType must be set to VPC_CHECKERS.
	// Possible values are STATIC_IP_CHECKERS and VPC_CHECKERS.
	// +kubebuilder:validation:Optional
	CheckerType *string `json:"checkerType,omitempty" tf:"checker_type,omitempty"`

	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ContentMatchers []ContentMatchersParameters `json:"contentMatchers,omitempty" tf:"content_matchers,omitempty"`

	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPCheck []HTTPCheckParameters `json:"httpCheck,omitempty" tf:"http_check,omitempty"`

	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance  aws_elb_load_balancer
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MonitoredResource []MonitoredResourceParameters `json:"monitoredResource,omitempty" tf:"monitored_resource,omitempty"`

	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	// +kubebuilder:validation:Optional
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The group resource associated with the configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ResourceGroup []ResourceGroupParameters `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	// +kubebuilder:validation:Optional
	SelectedRegions []*string `json:"selectedRegions,omitempty" tf:"selected_regions,omitempty"`

	// Contains information needed to make a TCP check.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TCPCheck []TCPCheckParameters `json:"tcpCheck,omitempty" tf:"tcp_check,omitempty"`

	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	// +kubebuilder:validation:Required
	Timeout *string `json:"timeout" tf:"timeout,omitempty"`
}

// UptimeCheckConfigSpec defines the desired state of UptimeCheckConfig
type UptimeCheckConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UptimeCheckConfigParameters `json:"forProvider"`
}

// UptimeCheckConfigStatus defines the observed state of UptimeCheckConfig.
type UptimeCheckConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UptimeCheckConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UptimeCheckConfig is the Schema for the UptimeCheckConfigs API. This message configures which resources and services to monitor for availability.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type UptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UptimeCheckConfigSpec   `json:"spec"`
	Status            UptimeCheckConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UptimeCheckConfigList contains a list of UptimeCheckConfigs
type UptimeCheckConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UptimeCheckConfig `json:"items"`
}

// Repository type metadata.
var (
	UptimeCheckConfig_Kind             = "UptimeCheckConfig"
	UptimeCheckConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UptimeCheckConfig_Kind}.String()
	UptimeCheckConfig_KindAPIVersion   = UptimeCheckConfig_Kind + "." + CRDGroupVersion.String()
	UptimeCheckConfig_GroupVersionKind = CRDGroupVersion.WithKind(UptimeCheckConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&UptimeCheckConfig{}, &UptimeCheckConfigList{})
}
