package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NotifierSpec defines the desired state of Notifier
type NotifierSpec struct {
	Message string `json:"message"`
	Target  string `json:"target"`
	Type    string `json:"type"`
}

// NotifierStatus defines the observed state of Notifier
type NotifierStatus struct {
	Notified          bool        `json:"notified"`
	LastSentTimestamp metav1.Time `json:"lastSentTimestamp,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Notifier is the Schema for the notifiers API
type Notifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotifierSpec   `json:"spec,omitempty"`
	Status NotifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotifierList contains a list of Notifier
type NotifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notifier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Notifier{}, &NotifierList{})
}
