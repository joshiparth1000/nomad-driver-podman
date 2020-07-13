// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InspectContainerConfig InspectContainerConfig holds further data about how a container was initially
// configured.
//
// swagger:model InspectContainerConfig
type InspectContainerConfig struct {

	// Container annotations
	Annotations map[string]string `json:"Annotations,omitempty"`

	// Unused, at present
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// Unused, at present
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Unused, at present
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Container command
	Cmd []string `json:"Cmd"`

	// CreateCommand is the full command plus arguments of the process the
	// container has been created with.
	CreateCommand []string `json:"CreateCommand"`

	// Container domain name - unused at present
	DomainName string `json:"Domainname,omitempty"`

	// Container entrypoint
	Entrypoint string `json:"Entrypoint,omitempty"`

	// Container environment variables
	Env []string `json:"Env"`

	// healthcheck
	Healthcheck *Schema2HealthConfig `json:"Healthcheck,omitempty"`

	// Container hostname
	Hostname string `json:"Hostname,omitempty"`

	// Container image
	Image string `json:"Image,omitempty"`

	// Container labels
	Labels map[string]string `json:"Labels,omitempty"`

	// On-build arguments - presently unused. More of Buildah's domain.
	OnBuild string `json:"OnBuild,omitempty"`

	// Whether the container leaves STDIN open
	OpenStdin bool `json:"OpenStdin,omitempty"`

	// Whether STDIN is only left open once.
	// Presently not supported by Podman, unused.
	StdinOnce bool `json:"StdinOnce,omitempty"`

	// Container stop signal
	StopSignal uint64 `json:"StopSignal,omitempty"`

	// Timezone is the timezone inside the container.
	// Local means it has the same timezone as the host machine
	Timezone string `json:"Timezone,omitempty"`

	// Whether the container creates a TTY
	Tty bool `json:"Tty,omitempty"`

	// User the container was launched with
	User string `json:"User,omitempty"`

	// Unused, at present. I've never seen this field populated.
	Volumes map[string]interface{} `json:"Volumes,omitempty"`

	// Container working directory
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this inspect container config
func (m *InspectContainerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthcheck(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InspectContainerConfig) validateHealthcheck(formats strfmt.Registry) error {

	if swag.IsZero(m.Healthcheck) { // not required
		return nil
	}

	if m.Healthcheck != nil {
		if err := m.Healthcheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Healthcheck")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InspectContainerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectContainerConfig) UnmarshalBinary(b []byte) error {
	var res InspectContainerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
