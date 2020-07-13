// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodBasicConfig PodBasicConfig contains basic configuration options for pods.
//
// swagger:model PodBasicConfig
type PodBasicConfig struct {

	// Hostname is the pod's hostname. If not set, the name of the pod will
	// be used (if a name was not provided here, the name auto-generated for
	// the pod will be used). This will be used by the infra container and
	// all containers in the pod as long as the UTS namespace is shared.
	// Optional.
	Hostname string `json:"hostname,omitempty"`

	// InfraCommand sets the command that will be used to start the infra
	// container.
	// If not set, the default set in the Libpod configuration file will be
	// used.
	// Conflicts with NoInfra=true.
	// Optional.
	InfraCommand []string `json:"infra_command"`

	// InfraConmonPidFile is a custom path to store the infra container's
	// conmon PID.
	InfraConmonPidFile string `json:"infra_conmon_pid_file,omitempty"`

	// InfraImage is the image that will be used for the infra container.
	// If not set, the default set in the Libpod configuration file will be
	// used.
	// Conflicts with NoInfra=true.
	// Optional.
	InfraImage string `json:"infra_image,omitempty"`

	// Labels are key-value pairs that are used to add metadata to pods.
	// Optional.
	Labels map[string]string `json:"labels,omitempty"`

	// Name is the name of the pod.
	// If not provided, a name will be generated when the pod is created.
	// Optional.
	Name string `json:"name,omitempty"`

	// NoInfra tells the pod not to create an infra container. If this is
	// done, many networking-related options will become unavailable.
	// Conflicts with setting any options in PodNetworkConfig, and the
	// InfraCommand and InfraImages in this struct.
	// Optional.
	NoInfra bool `json:"no_infra,omitempty"`

	// PodCreateCommand is the command used to create this pod.
	// This will be shown in the output of Inspect() on the pod, and may
	// also be used by some tools that wish to recreate the pod
	// (e.g. `podman generate systemd --new`).
	// Optional.
	PodCreateCommand []string `json:"pod_create_command"`

	// SharedNamespaces instructs the pod to share a set of namespaces.
	// Shared namespaces will be joined (by default) by every container
	// which joins the pod.
	// If not set and NoInfra is false, the pod will set a default set of
	// namespaces to share.
	// Conflicts with NoInfra=true.
	// Optional.
	SharedNamespaces []string `json:"shared_namespaces"`
}

// Validate validates this pod basic config
func (m *PodBasicConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodBasicConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodBasicConfig) UnmarshalBinary(b []byte) error {
	var res PodBasicConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
