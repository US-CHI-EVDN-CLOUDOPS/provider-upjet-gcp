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

type AppEngineRoutingOverrideObservation struct {

	// (Output)
	// The host that the task is sent to.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// App instance.
	// By default, the task is sent to an instance which is available when the task is attempted.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// App service.
	// By default, the task is sent to the service which is the default service when the task is attempted.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// App version.
	// By default, the task is sent to the version which is the default version when the task is attempted.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AppEngineRoutingOverrideParameters struct {

	// App instance.
	// By default, the task is sent to an instance which is available when the task is attempted.
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// App service.
	// By default, the task is sent to the service which is the default service when the task is attempted.
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// App version.
	// By default, the task is sent to the version which is the default version when the task is attempted.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type QueueObservation struct {

	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride []AppEngineRoutingOverrideObservation `json:"appEngineRoutingOverride,omitempty" tf:"app_engine_routing_override,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/queues/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the queue
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	RateLimits []RateLimitsObservation `json:"rateLimits,omitempty" tf:"rate_limits,omitempty"`

	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig []RetryConfigObservation `json:"retryConfig,omitempty" tf:"retry_config,omitempty"`

	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig []StackdriverLoggingConfigObservation `json:"stackdriverLoggingConfig,omitempty" tf:"stackdriver_logging_config,omitempty"`
}

type QueueParameters struct {

	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AppEngineRoutingOverride []AppEngineRoutingOverrideParameters `json:"appEngineRoutingOverride,omitempty" tf:"app_engine_routing_override,omitempty"`

	// The location of the queue
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// +kubebuilder:validation:Optional
	RateLimits []RateLimitsParameters `json:"rateLimits,omitempty" tf:"rate_limits,omitempty"`

	// Settings that determine the retry behavior.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RetryConfig []RetryConfigParameters `json:"retryConfig,omitempty" tf:"retry_config,omitempty"`

	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	StackdriverLoggingConfig []StackdriverLoggingConfigParameters `json:"stackdriverLoggingConfig,omitempty" tf:"stackdriver_logging_config,omitempty"`
}

type RateLimitsObservation struct {

	// (Output)
	// The max burst size.
	// Max burst size limits how fast tasks in queue are processed when many tasks are
	// in the queue and the rate is high. This field allows the queue to have a high
	// rate so processing starts shortly after a task is enqueued, but still limits
	// resource usage when many tasks are enqueued in a short period of time.
	MaxBurstSize *float64 `json:"maxBurstSize,omitempty" tf:"max_burst_size,omitempty"`

	// The maximum number of concurrent tasks that Cloud Tasks allows to
	// be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	MaxConcurrentDispatches *float64 `json:"maxConcurrentDispatches,omitempty" tf:"max_concurrent_dispatches,omitempty"`

	// The maximum rate at which tasks are dispatched from this queue.
	// If unspecified when the queue is created, Cloud Tasks will pick the default.
	MaxDispatchesPerSecond *float64 `json:"maxDispatchesPerSecond,omitempty" tf:"max_dispatches_per_second,omitempty"`
}

type RateLimitsParameters struct {

	// The maximum number of concurrent tasks that Cloud Tasks allows to
	// be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	// +kubebuilder:validation:Optional
	MaxConcurrentDispatches *float64 `json:"maxConcurrentDispatches,omitempty" tf:"max_concurrent_dispatches,omitempty"`

	// The maximum rate at which tasks are dispatched from this queue.
	// If unspecified when the queue is created, Cloud Tasks will pick the default.
	// +kubebuilder:validation:Optional
	MaxDispatchesPerSecond *float64 `json:"maxDispatchesPerSecond,omitempty" tf:"max_dispatches_per_second,omitempty"`
}

type RetryConfigObservation struct {

	// Number of attempts per task.
	// Cloud Tasks will attempt the task maxAttempts times (that is, if
	// the first attempt fails, then there will be maxAttempts - 1
	// retries). Must be >= -1.
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the default.
	// -1 indicates unlimited attempts.
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`

	// A task will be scheduled for retry between minBackoff and
	// maxBackoff duration after it fails, if the queue's RetryConfig
	// specifies that the task should be retried.
	MaxBackoff *string `json:"maxBackoff,omitempty" tf:"max_backoff,omitempty"`

	// The time between retries will double maxDoublings times.
	// A task's retry interval starts at minBackoff, then doubles maxDoublings times,
	// then increases linearly, and finally retries retries at intervals of maxBackoff
	// up to maxAttempts times.
	MaxDoublings *float64 `json:"maxDoublings,omitempty" tf:"max_doublings,omitempty"`

	// If positive, maxRetryDuration specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once maxRetryDuration time has passed and the task has
	// been attempted maxAttempts times, no further attempts will be
	// made and the task will be deleted.
	// If zero, then the task age is unlimited.
	MaxRetryDuration *string `json:"maxRetryDuration,omitempty" tf:"max_retry_duration,omitempty"`

	// A task will be scheduled for retry between minBackoff and
	// maxBackoff duration after it fails, if the queue's RetryConfig
	// specifies that the task should be retried.
	MinBackoff *string `json:"minBackoff,omitempty" tf:"min_backoff,omitempty"`
}

type RetryConfigParameters struct {

	// Number of attempts per task.
	// Cloud Tasks will attempt the task maxAttempts times (that is, if
	// the first attempt fails, then there will be maxAttempts - 1
	// retries). Must be >= -1.
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the default.
	// -1 indicates unlimited attempts.
	// +kubebuilder:validation:Optional
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`

	// A task will be scheduled for retry between minBackoff and
	// maxBackoff duration after it fails, if the queue's RetryConfig
	// specifies that the task should be retried.
	// +kubebuilder:validation:Optional
	MaxBackoff *string `json:"maxBackoff,omitempty" tf:"max_backoff,omitempty"`

	// The time between retries will double maxDoublings times.
	// A task's retry interval starts at minBackoff, then doubles maxDoublings times,
	// then increases linearly, and finally retries retries at intervals of maxBackoff
	// up to maxAttempts times.
	// +kubebuilder:validation:Optional
	MaxDoublings *float64 `json:"maxDoublings,omitempty" tf:"max_doublings,omitempty"`

	// If positive, maxRetryDuration specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once maxRetryDuration time has passed and the task has
	// been attempted maxAttempts times, no further attempts will be
	// made and the task will be deleted.
	// If zero, then the task age is unlimited.
	// +kubebuilder:validation:Optional
	MaxRetryDuration *string `json:"maxRetryDuration,omitempty" tf:"max_retry_duration,omitempty"`

	// A task will be scheduled for retry between minBackoff and
	// maxBackoff duration after it fails, if the queue's RetryConfig
	// specifies that the task should be retried.
	// +kubebuilder:validation:Optional
	MinBackoff *string `json:"minBackoff,omitempty" tf:"min_backoff,omitempty"`
}

type StackdriverLoggingConfigObservation struct {

	// Specifies the fraction of operations to write to Stackdriver Logging.
	// This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
	// default and means that no operations are logged.
	SamplingRatio *float64 `json:"samplingRatio,omitempty" tf:"sampling_ratio,omitempty"`
}

type StackdriverLoggingConfigParameters struct {

	// Specifies the fraction of operations to write to Stackdriver Logging.
	// This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
	// default and means that no operations are logged.
	// +kubebuilder:validation:Required
	SamplingRatio *float64 `json:"samplingRatio" tf:"sampling_ratio,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Queue is the Schema for the Queues API. A named resource to which messages are sent by publishers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec"`
	Status            QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
