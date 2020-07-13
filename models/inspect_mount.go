// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InspectMount InspectMount provides a record of a single mount in a container. It contains
// fields for both named and normal volumes. Only user-specified volumes will be
// included, and tmpfs volumes are not included even if the user specified them.
//
// swagger:model InspectMount
type InspectMount struct {

	// The destination directory for the volume. Specified as a path within
	// the container, as it would be passed into the OCI runtime.
	Destination string `json:"Destination,omitempty"`

	// The driver used for the named volume. Empty for bind mounts.
	Driver string `json:"Driver,omitempty"`

	// Contains SELinux :z/:Z mount options. Unclear what, if anything, else
	// goes in here.
	Mode string `json:"Mode,omitempty"`

	// The name of the volume. Empty for bind mounts.
	Name string `json:"Name,omitempty"`

	// All remaining mount options. Additional data, not present in the
	// original output.
	Options []string `json:"Options"`

	// Mount propagation for the mount. Can be empty if not specified, but
	// is always printed - no omitempty.
	Propagation string `json:"Propagation,omitempty"`

	// Whether the volume is read-write
	RW bool `json:"RW,omitempty"`

	// The source directory for the volume.
	Source string `json:"Source,omitempty"`

	// Whether the mount is a volume or bind mount. Allowed values are
	// "volume" and "bind".
	Type string `json:"Type,omitempty"`
}

// Validate validates this inspect mount
func (m *InspectMount) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InspectMount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectMount) UnmarshalBinary(b []byte) error {
	var res InspectMount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
