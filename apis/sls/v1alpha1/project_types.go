/*


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
	"reflect"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SLSProjectKind is the kind of SLSProject
	SLSProjectKind = reflect.TypeOf(SLSProject{}).Name()
	// SLSProjectGroupKind is the group and kind of SLSProject
	SLSProjectGroupKind = schema.GroupKind{Group: GroupVersion.Group, Kind: SLSProjectKind}.String()
	// SLSProjectKindAPIVersion is the kind, APIVersion of SLSProject
	SLSProjectKindAPIVersion = SLSProjectKind + "." + GroupVersion.Version
	// SLSProjectGroupVersionKind is the group, version and kind of SLSProject
	SLSProjectGroupVersionKind = GroupVersion.WithKind(SLSProjectKind)
)

// SLS Project states.
const (
	// The instance is healthy and available
	SLSProjectStateRunning = "Running"
	// The instance is being created. The instance is inaccessible while it is being created.
	SLSProjectStateCreating = "Creating"
	// The instance is being deleted.
	SLSProjectStateDeleting = "Deleting"
)

// SLSProjectSpec defines the desired state of SLS Project
type SLSProjectSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SLSProjectParameters `json:"forProvider"`
}

// SLSProjectObservation is the representation of the current state that is observed.
type SLSProjectObservation struct {
	// Name specifies the DB instance ID.
	Name   string `json:"name"`
	Status string `json:"status"`
}

// SLSProjectStatus defines the observed state of SLS Project
type SLSProjectStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SLSProjectObservation `json:"atProvider,omitempty"`
}

// SLSProjectParameters define the desired state of an SLS project.
type SLSProjectParameters struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// +kubebuilder:object:root=true

// SLSProject is the Schema for the SLS Projects API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibaba}
type SLSProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SLSProjectSpec   `json:"spec"`
	Status SLSProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SLSProjectList contains a list of SLSProject
type SLSProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SLSProject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SLSProject{}, &SLSProjectList{})
}
