// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineFinalizer allows ReconcileMetalMachine to clean up resources before removing it from the apiserver.
	MachineFinalizer = "metalmachine.infrastructure.cluster.x-k8s.io"
)

// TODO(rsmitty): These should probably come from a secret

// BMC defines data about how to talk to the node via ipmitool
type BMC struct {
	Endpoint string `json:"endpoint"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
}

// MetalMachineSpec defines the desired state of MetalMachine
type MetalMachineSpec struct {
	BMC BMC `json:"bmc,omitempty"`
}

// MetalMachineStatus defines the observed state of MetalMachine
type MetalMachineStatus struct {
	Ready bool `json:"ready"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=metalmachines,scope=Namespaced,categories=cluster-api
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// MetalMachine is the Schema for the metalmachines API
type MetalMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetalMachineSpec   `json:"spec,omitempty"`
	Status MetalMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetalMachineList contains a list of MetalMachine
type MetalMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetalMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetalMachine{}, &MetalMachineList{})
}
