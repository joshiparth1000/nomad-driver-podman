// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodPauseReport pod pause report
//
// swagger:model PodPauseReport
type PodPauseReport struct {

	// errs
	Errs []string `json:"Errs"`

	// Id
	ID string `json:"Id,omitempty"`
}

// Validate validates this pod pause report
func (m *PodPauseReport) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodPauseReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodPauseReport) UnmarshalBinary(b []byte) error {
	var res PodPauseReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
