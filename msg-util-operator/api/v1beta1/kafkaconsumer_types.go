/*
Copyright © 2025 Dexter Dixon dexterdixon561@gmail.com
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KafkaConsumerSpec defines the desired state of KafkaConsumer
type KafkaConsumerSpec struct {
	// Name of the consumer
	// +kubebuilder:validation:MinLength=0
	// +required
	Name string `json:"name"`

	// A list of bootstrap servers for the Kafka cluster
	// +kubebuilder:validation:MinLength=0
	// +required
	BootstrapServers []string `json:"bootstrapServers"`

	// List of topics to consume messages from
	// +kubebuilder:validation:MinLength=0
	// +required
	Topics []string `json:"topics"`

	// Name of the namespace where the secret containing TLS certificates is located
	// +kubebuilder:validation:MinLength=0
	// +required
	SecretNamespace string `json:"secretNamespace"`

	// Consumer Group ID
	// +kubebuilder:validation:MinLength=0
	// +required
	GroupID string `json:"groupID"`
}

// KafkaConsumerStatus defines the observed state of KafkaConsumer.
type KafkaConsumerStatus struct {
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KafkaConsumer is the Schema for the kafkaconsumers API
type KafkaConsumer struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of KafkaConsumer
	// +required
	Spec KafkaConsumerSpec `json:"spec"`

	// status defines the observed state of KafkaConsumer
	// +optional
	Status KafkaConsumerStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// KafkaConsumerList contains a list of KafkaConsumer
type KafkaConsumerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaConsumer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KafkaConsumer{}, &KafkaConsumerList{})
}
