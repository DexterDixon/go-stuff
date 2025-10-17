/*
Copyright © 2025 Dexter Dixon dexterdixon561@gmail.com
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KafkaProducerSpec defines the desired state of KafkaProducer
type KafkaProducerSpec struct {
	// Name of the producer
	// +kubebuilder:validation:MinLength=0
	// +required
	Name string `json:"name"`

	// A list of bootstrap servers for the Kafka cluster
	// +kubebuilder:validation:MinLength=0
	// +required
	BootstrapServers []string `json:"bootstrapServers"`

	// Name of the topic to send messages to
	// +kubebuilder:validation:MinLength=0
	// +required
	Topic string `json:"topic"`

	// Name of the namespace where the secret containing TLS certificates is located
	// +kubebuilder:validation:MinLength=0
	// +required
	SecretNamespace string `json:"secretNamespace"`
}

// KafkaProducerStatus defines the observed state of KafkaProducer.
type KafkaProducerStatus struct {
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KafkaProducer is the Schema for the kafkaproducers API
type KafkaProducer struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of KafkaProducer
	// +required
	Spec KafkaProducerSpec `json:"spec"`

	// status defines the observed state of KafkaProducer
	// +optional
	Status KafkaProducerStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// KafkaProducerList contains a list of KafkaProducer
type KafkaProducerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaProducer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KafkaProducer{}, &KafkaProducerList{})
}
