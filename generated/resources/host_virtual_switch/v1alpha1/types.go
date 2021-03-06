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

// HostVirtualSwitch is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type HostVirtualSwitch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostVirtualSwitchSpec   `json:"spec"`
	Status HostVirtualSwitchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostVirtualSwitch contains a list of HostVirtualSwitchList
type HostVirtualSwitchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostVirtualSwitch `json:"items"`
}

// A HostVirtualSwitchSpec defines the desired state of a HostVirtualSwitch
type HostVirtualSwitchSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HostVirtualSwitchParameters `json:"forProvider"`
}

// A HostVirtualSwitchParameters defines the desired state of a HostVirtualSwitch
type HostVirtualSwitchParameters struct {
	ActiveNics              []string `json:"active_nics,omitempty"`
	AllowForgedTransmits    bool     `json:"allow_forged_transmits"`
	AllowMacChanges         bool     `json:"allow_mac_changes"`
	AllowPromiscuous        bool     `json:"allow_promiscuous"`
	BeaconInterval          int64    `json:"beacon_interval"`
	CheckBeacon             bool     `json:"check_beacon"`
	Failback                bool     `json:"failback"`
	HostSystemId            string   `json:"host_system_id"`
	LinkDiscoveryOperation  string   `json:"link_discovery_operation"`
	LinkDiscoveryProtocol   string   `json:"link_discovery_protocol"`
	Mtu                     int64    `json:"mtu"`
	Name                    string   `json:"name"`
	NetworkAdapters         []string `json:"network_adapters,omitempty"`
	NotifySwitches          bool     `json:"notify_switches"`
	NumberOfPorts           int64    `json:"number_of_ports"`
	ShapingAverageBandwidth int64    `json:"shaping_average_bandwidth"`
	ShapingBurstSize        int64    `json:"shaping_burst_size"`
	ShapingEnabled          bool     `json:"shaping_enabled"`
	ShapingPeakBandwidth    int64    `json:"shaping_peak_bandwidth"`
	StandbyNics             []string `json:"standby_nics,omitempty"`
	TeamingPolicy           string   `json:"teaming_policy"`
}

// A HostVirtualSwitchStatus defines the observed state of a HostVirtualSwitch
type HostVirtualSwitchStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HostVirtualSwitchObservation `json:"atProvider"`
}

// A HostVirtualSwitchObservation records the observed state of a HostVirtualSwitch
type HostVirtualSwitchObservation struct{}
