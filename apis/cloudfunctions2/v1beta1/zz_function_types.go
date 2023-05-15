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

type BuildConfigObservation struct {

	// (Output)
	// The Cloud Build name of the latest successful
	// deployment of the function.
	Build *string `json:"build,omitempty" tf:"build,omitempty"`

	// User managed repository created in Artifact Registry optionally with a customer managed encryption key.
	DockerRepository *string `json:"dockerRepository,omitempty" tf:"docker_repository,omitempty"`

	// The name of the function (as defined in source code) that will be executed.
	// Defaults to the resource name suffix, if not specified. For backward
	// compatibility, if function with given name is not found, then the system
	// will try to use function named "function". For Node.js this is name of a
	// function exported by the module specified in source_location.
	EntryPoint *string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// User-provided build-time environment variables for the function.
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// The runtime in which to run the function. Required when deploying a new
	// function, optional when updating an existing function.
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// The location of the function source code.
	// Structure is documented below.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// Name of the Cloud Build Custom Worker Pool that should be used to build the function.
	WorkerPool *string `json:"workerPool,omitempty" tf:"worker_pool,omitempty"`
}

type BuildConfigParameters struct {

	// User managed repository created in Artifact Registry optionally with a customer managed encryption key.
	// +kubebuilder:validation:Optional
	DockerRepository *string `json:"dockerRepository,omitempty" tf:"docker_repository,omitempty"`

	// The name of the function (as defined in source code) that will be executed.
	// Defaults to the resource name suffix, if not specified. For backward
	// compatibility, if function with given name is not found, then the system
	// will try to use function named "function". For Node.js this is name of a
	// function exported by the module specified in source_location.
	// +kubebuilder:validation:Optional
	EntryPoint *string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// User-provided build-time environment variables for the function.
	// +kubebuilder:validation:Optional
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// The runtime in which to run the function. Required when deploying a new
	// function, optional when updating an existing function.
	// +kubebuilder:validation:Optional
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// The location of the function source code.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// Name of the Cloud Build Custom Worker Pool that should be used to build the function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudbuild/v1beta1.WorkerPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkerPool *string `json:"workerPool,omitempty" tf:"worker_pool,omitempty"`

	// Reference to a WorkerPool in cloudbuild to populate workerPool.
	// +kubebuilder:validation:Optional
	WorkerPoolRef *v1.Reference `json:"workerPoolRef,omitempty" tf:"-"`

	// Selector for a WorkerPool in cloudbuild to populate workerPool.
	// +kubebuilder:validation:Optional
	WorkerPoolSelector *v1.Selector `json:"workerPoolSelector,omitempty" tf:"-"`
}

type EventFiltersObservation struct {

	// 'Required. The name of a CloudEvents attribute.
	// Currently, only a subset of attributes are supported for filtering. Use the gcloud eventarc providers describe command to learn more about events and their attributes.
	// Do not filter for the 'type' attribute here, as this is already achieved by the resource's event_type attribute.
	Attribute *string `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// Optional. The operator used for matching the events with the value of
	// the filter. If not specified, only events that have an exact key-value
	// pair specified in the filter are matched.
	// The only allowed value is match-path-pattern.
	// See documentation on path patterns here'
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Required. The value for the attribute.
	// If the operator field is set as match-path-pattern, this value can be a path pattern instead of an exact value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EventFiltersParameters struct {

	// 'Required. The name of a CloudEvents attribute.
	// Currently, only a subset of attributes are supported for filtering. Use the gcloud eventarc providers describe command to learn more about events and their attributes.
	// Do not filter for the 'type' attribute here, as this is already achieved by the resource's event_type attribute.
	// +kubebuilder:validation:Required
	Attribute *string `json:"attribute" tf:"attribute,omitempty"`

	// Optional. The operator used for matching the events with the value of
	// the filter. If not specified, only events that have an exact key-value
	// pair specified in the filter are matched.
	// The only allowed value is match-path-pattern.
	// See documentation on path patterns here'
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Required. The value for the attribute.
	// If the operator field is set as match-path-pattern, this value can be a path pattern instead of an exact value.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Reference to a Bucket in storage to populate value.
	// +kubebuilder:validation:Optional
	ValueRef *v1.Reference `json:"valueRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate value.
	// +kubebuilder:validation:Optional
	ValueSelector *v1.Selector `json:"valueSelector,omitempty" tf:"-"`
}

type EventTriggerObservation struct {

	// Criteria used to filter events.
	// Structure is documented below.
	EventFilters []EventFiltersObservation `json:"eventFilters,omitempty" tf:"event_filters,omitempty"`

	// Required. The type of event to observe.
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// The name of a Pub/Sub topic in the same project that will be used
	// as the transport topic for the event delivery.
	PubsubTopic *string `json:"pubsubTopic,omitempty" tf:"pubsub_topic,omitempty"`

	// Describes the retry policy in case of function's execution failure.
	// Retried execution is charged as any other execution.
	// Possible values are: RETRY_POLICY_UNSPECIFIED, RETRY_POLICY_DO_NOT_RETRY, RETRY_POLICY_RETRY.
	RetryPolicy *string `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// The email of the service account for this function.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// (Output)
	// Output only. The resource name of the Eventarc trigger.
	Trigger *string `json:"trigger,omitempty" tf:"trigger,omitempty"`

	// The region that the trigger will be in. The trigger will only receive
	// events originating in this region. It can be the same
	// region as the function, a different region or multi-region, or the global
	// region. If not provided, defaults to the same region as the function.
	TriggerRegion *string `json:"triggerRegion,omitempty" tf:"trigger_region,omitempty"`
}

type EventTriggerParameters struct {

	// Criteria used to filter events.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EventFilters []EventFiltersParameters `json:"eventFilters,omitempty" tf:"event_filters,omitempty"`

	// Required. The type of event to observe.
	// +kubebuilder:validation:Optional
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// The name of a Pub/Sub topic in the same project that will be used
	// as the transport topic for the event delivery.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/pubsub/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PubsubTopic *string `json:"pubsubTopic,omitempty" tf:"pubsub_topic,omitempty"`

	// Reference to a Topic in pubsub to populate pubsubTopic.
	// +kubebuilder:validation:Optional
	PubsubTopicRef *v1.Reference `json:"pubsubTopicRef,omitempty" tf:"-"`

	// Selector for a Topic in pubsub to populate pubsubTopic.
	// +kubebuilder:validation:Optional
	PubsubTopicSelector *v1.Selector `json:"pubsubTopicSelector,omitempty" tf:"-"`

	// Describes the retry policy in case of function's execution failure.
	// Retried execution is charged as any other execution.
	// Possible values are: RETRY_POLICY_UNSPECIFIED, RETRY_POLICY_DO_NOT_RETRY, RETRY_POLICY_RETRY.
	// +kubebuilder:validation:Optional
	RetryPolicy *string `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// The email of the service account for this function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailRef *v1.Reference `json:"serviceAccountEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailSelector *v1.Selector `json:"serviceAccountEmailSelector,omitempty" tf:"-"`

	// The region that the trigger will be in. The trigger will only receive
	// events originating in this region. It can be the same
	// region as the function, a different region or multi-region, or the global
	// region. If not provided, defaults to the same region as the function.
	// +kubebuilder:validation:Optional
	TriggerRegion *string `json:"triggerRegion,omitempty" tf:"trigger_region,omitempty"`
}

type FunctionObservation struct {

	// Describes the Build step of the function that builds a container
	// from the given source.
	// Structure is documented below.
	BuildConfig []BuildConfigObservation `json:"buildConfig,omitempty" tf:"build_config,omitempty"`

	// User-provided description of a function.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The environment the function is hosted on.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// An Eventarc trigger managed by Google Cloud Functions that fires events in
	// response to a condition in another service.
	// Structure is documented below.
	EventTrigger []EventTriggerObservation `json:"eventTrigger,omitempty" tf:"event_trigger,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/functions/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs associated with this Cloud Function.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location of this cloud function.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Describes the Service being deployed.
	// Structure is documented below.
	ServiceConfig []ServiceConfigObservation `json:"serviceConfig,omitempty" tf:"service_config,omitempty"`

	// Describes the current state of the function.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The last update timestamp of a Cloud Function.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type FunctionParameters struct {

	// Describes the Build step of the function that builds a container
	// from the given source.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BuildConfig []BuildConfigParameters `json:"buildConfig,omitempty" tf:"build_config,omitempty"`

	// User-provided description of a function.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An Eventarc trigger managed by Google Cloud Functions that fires events in
	// response to a condition in another service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EventTrigger []EventTriggerParameters `json:"eventTrigger,omitempty" tf:"event_trigger,omitempty"`

	// A set of key/value label pairs associated with this Cloud Function.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location of this cloud function.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Describes the Service being deployed.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ServiceConfig []ServiceConfigParameters `json:"serviceConfig,omitempty" tf:"service_config,omitempty"`
}

type RepoSourceObservation struct {

	// Regex matching branches to build.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Regex matching tags to build.
	CommitSha *string `json:"commitSha,omitempty" tf:"commit_sha,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	Dir *string `json:"dir,omitempty" tf:"dir,omitempty"`

	// Only trigger a build if the revision regex does
	// NOT match the revision regex.
	InvertRegex *bool `json:"invertRegex,omitempty" tf:"invert_regex,omitempty"`

	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Name of the Cloud Source Repository.
	RepoName *string `json:"repoName,omitempty" tf:"repo_name,omitempty"`

	// Regex matching tags to build.
	TagName *string `json:"tagName,omitempty" tf:"tag_name,omitempty"`
}

type RepoSourceParameters struct {

	// Regex matching branches to build.
	// +kubebuilder:validation:Optional
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Regex matching tags to build.
	// +kubebuilder:validation:Optional
	CommitSha *string `json:"commitSha,omitempty" tf:"commit_sha,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	// +kubebuilder:validation:Optional
	Dir *string `json:"dir,omitempty" tf:"dir,omitempty"`

	// Only trigger a build if the revision regex does
	// NOT match the revision regex.
	// +kubebuilder:validation:Optional
	InvertRegex *bool `json:"invertRegex,omitempty" tf:"invert_regex,omitempty"`

	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Name of the Cloud Source Repository.
	// +kubebuilder:validation:Optional
	RepoName *string `json:"repoName,omitempty" tf:"repo_name,omitempty"`

	// Regex matching tags to build.
	// +kubebuilder:validation:Optional
	TagName *string `json:"tagName,omitempty" tf:"tag_name,omitempty"`
}

type SecretEnvironmentVariablesObservation struct {

	// Name of the environment variable.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Version of the secret (version number or the string 'latest'). It is preferable to use latest version with secret volumes as secret value changes are reflected immediately.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SecretEnvironmentVariablesParameters struct {

	// Name of the environment variable.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/secretmanager/v1beta1.Secret
	// +kubebuilder:validation:Optional
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Reference to a Secret in secretmanager to populate secret.
	// +kubebuilder:validation:Optional
	SecretRef *v1.Reference `json:"secretRef,omitempty" tf:"-"`

	// Selector for a Secret in secretmanager to populate secret.
	// +kubebuilder:validation:Optional
	SecretSelector *v1.Selector `json:"secretSelector,omitempty" tf:"-"`

	// Version of the secret (version number or the string 'latest'). It is preferable to use latest version with secret volumes as secret value changes are reflected immediately.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type SecretVolumesObservation struct {

	// The path within the container to mount the secret volume. For example, setting the mountPath as /etc/secrets would mount the secret value files under the /etc/secrets directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount path: /etc/secrets
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// List of secret versions to mount for this secret. If empty, the latest version of the secret will be made available in a file named after the secret under the mount point.'
	// Structure is documented below.
	Versions []VersionsObservation `json:"versions,omitempty" tf:"versions,omitempty"`
}

type SecretVolumesParameters struct {

	// The path within the container to mount the secret volume. For example, setting the mountPath as /etc/secrets would mount the secret value files under the /etc/secrets directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount path: /etc/secrets
	// +kubebuilder:validation:Required
	MountPath *string `json:"mountPath" tf:"mount_path,omitempty"`

	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/secretmanager/v1beta1.Secret
	// +kubebuilder:validation:Optional
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Reference to a Secret in secretmanager to populate secret.
	// +kubebuilder:validation:Optional
	SecretRef *v1.Reference `json:"secretRef,omitempty" tf:"-"`

	// Selector for a Secret in secretmanager to populate secret.
	// +kubebuilder:validation:Optional
	SecretSelector *v1.Selector `json:"secretSelector,omitempty" tf:"-"`

	// List of secret versions to mount for this secret. If empty, the latest version of the secret will be made available in a file named after the secret under the mount point.'
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Versions []VersionsParameters `json:"versions,omitempty" tf:"versions,omitempty"`
}

type ServiceConfigObservation struct {

	// Whether 100% of traffic is routed to the latest revision. Defaults to true.
	AllTrafficOnLatestRevision *bool `json:"allTrafficOnLatestRevision,omitempty" tf:"all_traffic_on_latest_revision,omitempty"`

	// The number of CPUs used in a single container instance. Default value is calculated from available memory.
	AvailableCPU *string `json:"availableCpu,omitempty" tf:"available_cpu,omitempty"`

	// The amount of memory available for a function.
	// Defaults to 256M. Supported units are k, M, G, Mi, Gi. If no unit is
	// supplied the value is interpreted as bytes.
	AvailableMemory *string `json:"availableMemory,omitempty" tf:"available_memory,omitempty"`

	// Environment variables that shall be available during function execution.
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// (Output)
	// URIs of the Service deployed
	GcfURI *string `json:"gcfUri,omitempty" tf:"gcf_uri,omitempty"`

	// Available ingress settings. Defaults to "ALLOW_ALL" if unspecified.
	// Default value is ALLOW_ALL.
	// Possible values are: ALLOW_ALL, ALLOW_INTERNAL_ONLY, ALLOW_INTERNAL_AND_GCLB.
	IngressSettings *string `json:"ingressSettings,omitempty" tf:"ingress_settings,omitempty"`

	// The limit on the maximum number of function instances that may coexist at a
	// given time.
	MaxInstanceCount *float64 `json:"maxInstanceCount,omitempty" tf:"max_instance_count,omitempty"`

	// Sets the maximum number of concurrent requests that each instance can receive. Defaults to 1.
	MaxInstanceRequestConcurrency *float64 `json:"maxInstanceRequestConcurrency,omitempty" tf:"max_instance_request_concurrency,omitempty"`

	// The limit on the minimum number of function instances that may coexist at a
	// given time.
	MinInstanceCount *float64 `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`

	// Secret environment variables configuration.
	// Structure is documented below.
	SecretEnvironmentVariables []SecretEnvironmentVariablesObservation `json:"secretEnvironmentVariables,omitempty" tf:"secret_environment_variables,omitempty"`

	// Secret volumes configuration.
	// Structure is documented below.
	SecretVolumes []SecretVolumesObservation `json:"secretVolumes,omitempty" tf:"secret_volumes,omitempty"`

	// Name of the service associated with a Function.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// The email of the service account for this function.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// The function execution timeout. Execution is considered failed and
	// can be terminated if the function is not completed at the end of the
	// timeout period. Defaults to 60 seconds.
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`

	// (Output)
	// URI of the Service deployed.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// The Serverless VPC Access connector that this cloud function can connect to.
	VPCConnector *string `json:"vpcConnector,omitempty" tf:"vpc_connector,omitempty"`

	// Available egress settings.
	// Possible values are: VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED, PRIVATE_RANGES_ONLY, ALL_TRAFFIC.
	VPCConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty" tf:"vpc_connector_egress_settings,omitempty"`
}

type ServiceConfigParameters struct {

	// Whether 100% of traffic is routed to the latest revision. Defaults to true.
	// +kubebuilder:validation:Optional
	AllTrafficOnLatestRevision *bool `json:"allTrafficOnLatestRevision,omitempty" tf:"all_traffic_on_latest_revision,omitempty"`

	// The number of CPUs used in a single container instance. Default value is calculated from available memory.
	// +kubebuilder:validation:Optional
	AvailableCPU *string `json:"availableCpu,omitempty" tf:"available_cpu,omitempty"`

	// The amount of memory available for a function.
	// Defaults to 256M. Supported units are k, M, G, Mi, Gi. If no unit is
	// supplied the value is interpreted as bytes.
	// +kubebuilder:validation:Optional
	AvailableMemory *string `json:"availableMemory,omitempty" tf:"available_memory,omitempty"`

	// Environment variables that shall be available during function execution.
	// +kubebuilder:validation:Optional
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Available ingress settings. Defaults to "ALLOW_ALL" if unspecified.
	// Default value is ALLOW_ALL.
	// Possible values are: ALLOW_ALL, ALLOW_INTERNAL_ONLY, ALLOW_INTERNAL_AND_GCLB.
	// +kubebuilder:validation:Optional
	IngressSettings *string `json:"ingressSettings,omitempty" tf:"ingress_settings,omitempty"`

	// The limit on the maximum number of function instances that may coexist at a
	// given time.
	// +kubebuilder:validation:Optional
	MaxInstanceCount *float64 `json:"maxInstanceCount,omitempty" tf:"max_instance_count,omitempty"`

	// Sets the maximum number of concurrent requests that each instance can receive. Defaults to 1.
	// +kubebuilder:validation:Optional
	MaxInstanceRequestConcurrency *float64 `json:"maxInstanceRequestConcurrency,omitempty" tf:"max_instance_request_concurrency,omitempty"`

	// The limit on the minimum number of function instances that may coexist at a
	// given time.
	// +kubebuilder:validation:Optional
	MinInstanceCount *float64 `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`

	// Secret environment variables configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SecretEnvironmentVariables []SecretEnvironmentVariablesParameters `json:"secretEnvironmentVariables,omitempty" tf:"secret_environment_variables,omitempty"`

	// Secret volumes configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SecretVolumes []SecretVolumesParameters `json:"secretVolumes,omitempty" tf:"secret_volumes,omitempty"`

	// Name of the service associated with a Function.
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// The email of the service account for this function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailRef *v1.Reference `json:"serviceAccountEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailSelector *v1.Selector `json:"serviceAccountEmailSelector,omitempty" tf:"-"`

	// The function execution timeout. Execution is considered failed and
	// can be terminated if the function is not completed at the end of the
	// timeout period. Defaults to 60 seconds.
	// +kubebuilder:validation:Optional
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`

	// The Serverless VPC Access connector that this cloud function can connect to.
	// +kubebuilder:validation:Optional
	VPCConnector *string `json:"vpcConnector,omitempty" tf:"vpc_connector,omitempty"`

	// Available egress settings.
	// Possible values are: VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED, PRIVATE_RANGES_ONLY, ALL_TRAFFIC.
	// +kubebuilder:validation:Optional
	VPCConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty" tf:"vpc_connector_egress_settings,omitempty"`
}

type SourceObservation struct {

	// If provided, get the source from this location in a Cloud Source Repository.
	// Structure is documented below.
	RepoSource []RepoSourceObservation `json:"repoSource,omitempty" tf:"repo_source,omitempty"`

	// If provided, get the source from this location in Google Cloud Storage.
	// Structure is documented below.
	StorageSource []StorageSourceObservation `json:"storageSource,omitempty" tf:"storage_source,omitempty"`
}

type SourceParameters struct {

	// If provided, get the source from this location in a Cloud Source Repository.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RepoSource []RepoSourceParameters `json:"repoSource,omitempty" tf:"repo_source,omitempty"`

	// If provided, get the source from this location in Google Cloud Storage.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	StorageSource []StorageSourceParameters `json:"storageSource,omitempty" tf:"storage_source,omitempty"`
}

type StorageSourceObservation struct {

	// Google Cloud Storage bucket containing the source
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Google Cloud Storage generation for the object. If the generation
	// is omitted, the latest generation will be used.
	Generation *float64 `json:"generation,omitempty" tf:"generation,omitempty"`

	// Google Cloud Storage object containing the source.
	Object *string `json:"object,omitempty" tf:"object,omitempty"`
}

type StorageSourceParameters struct {

	// Google Cloud Storage bucket containing the source
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Google Cloud Storage generation for the object. If the generation
	// is omitted, the latest generation will be used.
	// +kubebuilder:validation:Optional
	Generation *float64 `json:"generation,omitempty" tf:"generation,omitempty"`

	// Google Cloud Storage object containing the source.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.BucketObject
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// Reference to a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectRef *v1.Reference `json:"objectRef,omitempty" tf:"-"`

	// Selector for a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectSelector *v1.Selector `json:"objectSelector,omitempty" tf:"-"`
}

type VersionsObservation struct {

	// Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mountPath as '/etc/secrets' and path as secret_foo would mount the secret value file at /etc/secrets/secret_foo.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Version of the secret (version number or the string 'latest'). It is preferable to use latest version with secret volumes as secret value changes are reflected immediately.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type VersionsParameters struct {

	// Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mountPath as '/etc/secrets' and path as secret_foo would mount the secret value file at /etc/secrets/secret_foo.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Version of the secret (version number or the string 'latest'). It is preferable to use latest version with secret volumes as secret value changes are reflected immediately.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API. A Cloud Function that contains user computation executed in response to an event.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
