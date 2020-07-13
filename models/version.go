// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Version Version is an output struct for varlink
//
// swagger:model Version
type Version struct {

	// API version
	APIVersion int64 `json:"APIVersion,omitempty"`

	// built
	Built int64 `json:"Built,omitempty"`

	// built time
	BuiltTime string `json:"BuiltTime,omitempty"`

	// git commit
	GitCommit string `json:"GitCommit,omitempty"`

	// go version
	GoVersion string `json:"GoVersion,omitempty"`

	// os arch
	OsArch string `json:"OsArch,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this version
func (m *Version) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Version) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Version) UnmarshalBinary(b []byte) error {
	var res Version
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
