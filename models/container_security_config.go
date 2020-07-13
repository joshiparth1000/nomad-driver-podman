// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerSecurityConfig ContainerSecurityConfig is a container's security features, including
// SELinux, Apparmor, and Seccomp.
//
// swagger:model ContainerSecurityConfig
type ContainerSecurityConfig struct {

	// ApparmorProfile is the name of the Apparmor profile the container
	// will use.
	// Optional.
	ApparmorProfile string `json:"apparmor_profile,omitempty"`

	// CapAdd are capabilities which will be added to the container.
	// Conflicts with Privileged.
	// Optional.
	CapAdd []string `json:"cap_add"`

	// CapDrop are capabilities which will be removed from the container.
	// Conflicts with Privileged.
	// Optional.
	CapDrop []string `json:"cap_drop"`

	// Groups are a list of supplemental groups the container's user will
	// be granted access to.
	// Optional.
	Groups []string `json:"groups"`

	// NoNewPrivileges is whether the container will set the no new
	// privileges flag on create, which disables gaining additional
	// privileges (e.g. via setuid) in the container.
	NoNewPrivileges bool `json:"no_new_privileges,omitempty"`

	// Privileged is whether the container is privileged.
	// Privileged does the following:
	// Adds all devices on the system to the container.
	// Adds all capabilities to the container.
	// Disables Seccomp, SELinux, and Apparmor confinement.
	// (Though SELinux can be manually re-enabled).
	// TODO: this conflicts with things.
	// TODO: this does more.
	Privileged bool `json:"privileged,omitempty"`

	// ReadOnlyFilesystem indicates that everything will be mounted
	// as read-only
	ReadOnlyFilesystem bool `json:"read_only_filesystem,omitempty"`

	// SeccompPolicy determines which seccomp profile gets applied
	// the container. valid values: empty,default,image
	SeccompPolicy string `json:"seccomp_policy,omitempty"`

	// SeccompProfilePath is the path to a JSON file containing the
	// container's Seccomp profile.
	// If not specified, no Seccomp profile will be used.
	// Optional.
	SeccompProfilePath string `json:"seccomp_profile_path,omitempty"`

	// SelinuxProcessLabel is the process label the container will use.
	// If SELinux is enabled and this is not specified, a label will be
	// automatically generated if not specified.
	// Optional.
	SelinuxOpts []string `json:"selinux_opts"`

	// User is the user the container will be run as.
	// Can be given as a UID or a username; if a username, it will be
	// resolved within the container, using the container's /etc/passwd.
	// If unset, the container will be run as root.
	// Optional.
	User string `json:"user,omitempty"`

	// idmappings
	Idmappings *IDMappingOptions `json:"idmappings,omitempty"`

	// userns
	Userns *Namespace `json:"userns,omitempty"`
}

// Validate validates this container security config
func (m *ContainerSecurityConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdmappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerSecurityConfig) validateIdmappings(formats strfmt.Registry) error {

	if swag.IsZero(m.Idmappings) { // not required
		return nil
	}

	if m.Idmappings != nil {
		if err := m.Idmappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idmappings")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerSecurityConfig) validateUserns(formats strfmt.Registry) error {

	if swag.IsZero(m.Userns) { // not required
		return nil
	}

	if m.Userns != nil {
		if err := m.Userns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userns")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerSecurityConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerSecurityConfig) UnmarshalBinary(b []byte) error {
	var res ContainerSecurityConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
