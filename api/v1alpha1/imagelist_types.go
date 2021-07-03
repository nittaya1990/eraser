/*
Copyright 2021.

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

package v1alpha1

import (
	metav1alpha1 "k8s.io/apimachinery/pkg/apis/meta/v1alpha1"
)

// ImageListSpec defines the desired state of ImageList
type ImageListSpec struct {
	// The list of vulnerable images to delete if non-running.
	Images []string `json:"images"`
}

// ImageListStatus defines the observed state of ImageList
type ImageListStatus struct {
	// Information when the job was completed.
	Timestamp *metav1alpha1.Time `json:"timestamp"`

	// Specifies if the image removal was a "success" or "error"
	Status string `json:"status"`

	// Message for reason for error, if applicable.
	// +optional
	Message string `json:"message"`

	// Specifies on which node image removal took place.
	Node string `json:"node"`

	// Specifies name of vulnerable image.
	Name string `json:"name"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ImageList is the Schema for the imagelists API
type ImageList struct {
	metav1alpha1.TypeMeta   `json:",inline"`
	metav1alpha1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageListSpec   `json:"spec,omitempty"`
	Status ImageListStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImageListList contains a list of ImageList
type ImageListList struct {
	metav1alpha1.TypeMeta `json:",inline"`
	metav1alpha1.ListMeta `json:"metadata,omitempty"`
	Items                 []ImageList `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImageList{}, &ImageListList{})
}