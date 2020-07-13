// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InspectContainerState InspectContainerState provides a detailed record of a container's current
// state. It is returned as part of InspectContainerData.
// As with InspectContainerData, many portions of this struct are matched to
// Docker, but here we see more fields that are unused (nonsensical in the
// context of Libpod).
//
// swagger:model InspectContainerState
type InspectContainerState struct {

	// conmon pid
	ConmonPid int64 `json:"ConmonPid,omitempty"`

	// dead
	Dead bool `json:"Dead,omitempty"`

	// error
	Error string `json:"Error,omitempty"`

	// exit code
	ExitCode int32 `json:"ExitCode,omitempty"`

	// finished at
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"FinishedAt,omitempty"`

	// healthcheck
	Healthcheck *HealthCheckResults `json:"Healthcheck,omitempty"`

	// o o m killed
	OOMKilled bool `json:"OOMKilled,omitempty"`

	// oci version
	OciVersion string `json:"OciVersion,omitempty"`

	// paused
	Paused bool `json:"Paused,omitempty"`

	// pid
	Pid int64 `json:"Pid,omitempty"`

	// restarting
	Restarting bool `json:"Restarting,omitempty"`

	// running
	Running bool `json:"Running,omitempty"`

	// started at
	// Format: date-time
	StartedAt strfmt.DateTime `json:"StartedAt,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this inspect container state
func (m *InspectContainerState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthcheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InspectContainerState) validateFinishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("FinishedAt", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InspectContainerState) validateHealthcheck(formats strfmt.Registry) error {

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

func (m *InspectContainerState) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("StartedAt", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InspectContainerState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectContainerState) UnmarshalBinary(b []byte) error {
	var res InspectContainerState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
