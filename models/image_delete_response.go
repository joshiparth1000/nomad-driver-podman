// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageDeleteResponse ImageDeleteResponse is the response for removing an image from storage and containers
// what was untagged vs actually removed
//
// swagger:model ImageDeleteResponse
type ImageDeleteResponse struct {

	// deleted
	Deleted string `json:"deleted,omitempty"`

	// untagged
	Untagged []string `json:"untagged"`
}

// Validate validates this image delete response
func (m *ImageDeleteResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageDeleteResponse) UnmarshalBinary(b []byte) error {
	var res ImageDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
