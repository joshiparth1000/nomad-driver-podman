// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NamedVolume NamedVolume holds information about a named volume that will be mounted into
// the container.
//
// swagger:model NamedVolume
type NamedVolume struct {

	// Destination to mount the named volume within the container. Must be
	// an absolute path. Path will be created if it does not exist.
	Dest string `json:"Dest,omitempty"`

	// Name is the name of the named volume to be mounted. May be empty.
	// If empty, a new named volume with a pseudorandomly generated name
	// will be mounted at the given destination.
	Name string `json:"Name,omitempty"`

	// Options are options that the named volume will be mounted with.
	Options []string `json:"Options"`
}

// Validate validates this named volume
func (m *NamedVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamedVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamedVolume) UnmarshalBinary(b []byte) error {
	var res NamedVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
