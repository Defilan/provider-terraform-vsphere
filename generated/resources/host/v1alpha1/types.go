/*
	Copyright 2019 The Crossplane Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// +kubebuilder:object:root=true

// Host is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Host struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostSpec   `json:"spec"`
	Status HostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Host contains a list of HostList
type HostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Host `json:"items"`
}

// A HostSpec defines the desired state of a Host
type HostSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HostParameters `json:"forProvider"`
}

// A HostParameters defines the desired state of a Host
type HostParameters struct {
	Cluster        string `json:"cluster"`
	ClusterManaged bool   `json:"cluster_managed"`
	Connected      bool   `json:"connected"`
	Datacenter     string `json:"datacenter"`
	Force          bool   `json:"force"`
	Hostname       string `json:"hostname"`
	License        string `json:"license"`
	Lockdown       string `json:"lockdown"`
	Maintenance    bool   `json:"maintenance"`
	Password       string `json:"password"`
	Thumbprint     string `json:"thumbprint"`
	Username       string `json:"username"`
}

// A HostStatus defines the observed state of a Host
type HostStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HostObservation `json:"atProvider"`
}

// A HostObservation records the observed state of a Host
type HostObservation struct{}
