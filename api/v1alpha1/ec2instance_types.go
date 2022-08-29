package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EC2InstanceSpec defines the desired state of EC2Instance
type EC2InstanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of EC2Instance. Edit ec2instance_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// EC2InstanceStatus defines the observed state of EC2Instance
type EC2InstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// EC2Instance is the Schema for the ec2instances API
type EC2Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EC2InstanceSpec   `json:"spec,omitempty"`
	Status EC2InstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EC2InstanceList contains a list of EC2Instance
type EC2InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EC2Instance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EC2Instance{}, &EC2InstanceList{})
}
