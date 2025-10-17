package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OnboardUserSpec defines the desired state of OnboardUser
type OnboardUserSpec struct {
	// user to onboard
	// +kubebuilder:validation:MinLength=0
	// +required
	Username string `json:"username"`

	// role of user to onboard
	// +kubebuilder:validation:MinLength=0
	// +required
	Role string `json:"role"`

	// prefered text editor of the user
	// +kubebuilder:validation:MinLength=0
	// +required
	TextEditor string `json:"textEditor"`
}

// Sets the role of the user to be onboarded the default one
// is Developer.
// +kubebuilder:validation:Enum=developer;maintainer;guest
// +required
type Role string

const (
	// Sets role to developer
	Developer Role = "Developer"

	// Sets role to maintainter
	Maintainer Role = "Maintainer"

	// Sets role to guest
	Guest Role = "Guest"
)

// OnboardUser defines the observed state of Onbarding process.
type OnboardUserStatus struct {
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

// OnboardUser is the Schema for the OnboardUsers API
type OnboardUser struct {
	/*
	 */
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of OnboardUser
	// +required
	Spec OnboardUserSpec `json:"spec"`

	// status defines the observed state of OnboardUser
	// +optional
	Status OnboardUserStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// OnboardUserList contains a list of OnboardUsers
type OnboardUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OnboardUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OnboardUser{}, &OnboardUserList{})
}

// +kubebuilder:docs-gen:collapse=Root Object Definitions
