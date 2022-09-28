// +kubebuilder:object:generate=true
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PXIssue is a resentation of a PX issue
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=pxi
// +groupName=autopilot.libopenstorage.org
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.reason`,description="The reason for this issue"
// +kubebuilder:printcolumn:name="Resolution",type=string,JSONPath=`.resolution`,description="The resolution for this issue"
type PXIssue struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Namespace defines the namespace for this issue
	Namespace string `json:"namespace" protobuf:"bytes,2,opt,name=namespace"`

	// Type defines the type for this issue
	Type string `json:"type" protobuf:"bytes,2,opt,name=type"`

	// State defines the state for this issue
	State string `json:"state" protobuf:"bytes,2,opt,name=state"`

	// Reason defines the reason for this issue
	Reason string `json:"reason" protobuf:"bytes,2,opt,name=reason"`

	// Docs defines the docs for this issue
	Docs string `json:"docs" protobuf:"bytes,2,opt,name=docs"`

	// Resolution defines the resolution for this issue
	// +optional
	Resolution string `json:"resolution" protobuf:"bytes,3,opt,name=resolution"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PXIssueList is a list of PXIssue objects
type PXIssueList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of PXIssueList
	Items []PXIssue `json:"items" protobuf:"bytes,2,rep,name=items"`
}
