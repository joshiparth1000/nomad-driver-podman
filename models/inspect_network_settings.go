// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InspectNetworkSettings InspectNetworkSettings holds information about the network settings of the
// container.
// Many fields are maintained only for compatibility with `docker inspect` and
// are unused within Libpod.
//
// swagger:model InspectNetworkSettings
type InspectNetworkSettings struct {

	// AdditionalMacAddresses is a set of additional MAC Addresses beyond
	// the first. CNI may configure more than one interface for a single
	// network, which can cause this.
	AdditionalMacAddresses []string `json:"AdditionalMACAddresses"`

	// bridge
	Bridge string `json:"Bridge,omitempty"`

	// EndpointID is unused, maintained exclusively for compatibility.
	EndpointID string `json:"EndpointID,omitempty"`

	// Gateway is the IP address of the gateway this network will use.
	Gateway string `json:"Gateway,omitempty"`

	// GlobalIPv6Address is the global-scope IPv6 Address for this network.
	GlobalIPV6Address string `json:"GlobalIPv6Address,omitempty"`

	// GlobalIPv6PrefixLen is the length of the subnet mask of this network.
	GlobalIPV6PrefixLen int64 `json:"GlobalIPv6PrefixLen,omitempty"`

	// hairpin mode
	HairpinMode bool `json:"HairpinMode,omitempty"`

	// IPAddress is the IP address for this network.
	IPAddress string `json:"IPAddress,omitempty"`

	// IPPrefixLen is the length of the subnet mask of this network.
	IPPrefixLen int64 `json:"IPPrefixLen,omitempty"`

	// IPv6Gateway is the IPv6 gateway this network will use.
	IPV6Gateway string `json:"IPv6Gateway,omitempty"`

	// link local IPv6 address
	LinkLocalIPV6Address string `json:"LinkLocalIPv6Address,omitempty"`

	// link local IPv6 prefix len
	LinkLocalIPV6PrefixLen int64 `json:"LinkLocalIPv6PrefixLen,omitempty"`

	// MacAddress is the MAC address for the interface in this network.
	MacAddress string `json:"MacAddress,omitempty"`

	// Networks contains information on non-default CNI networks this
	// container has joined.
	// It is a map of network name to network information.
	Networks map[string]InspectAdditionalNetwork `json:"Networks,omitempty"`

	// ports
	Ports map[string][]InspectHostPort `json:"Ports,omitempty"`

	// sandbox ID
	SandboxID string `json:"SandboxID,omitempty"`

	// sandbox key
	SandboxKey string `json:"SandboxKey,omitempty"`

	// SecondaryIPAddresses is a list of extra IP Addresses that the
	// container has been assigned in this network.
	SecondaryIPAddresses []string `json:"SecondaryIPAddresses"`

	// SecondaryIPv6Addresses is a list of extra IPv6 Addresses that the
	// container has been assigned in this networ.
	SecondaryIPV6Addresses []string `json:"SecondaryIPv6Addresses"`
}

// Validate validates this inspect network settings
func (m *InspectNetworkSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InspectNetworkSettings) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for k := range m.Networks {

		if err := validate.Required("Networks"+"."+k, "body", m.Networks[k]); err != nil {
			return err
		}
		if val, ok := m.Networks[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *InspectNetworkSettings) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for k := range m.Ports {

		if err := validate.Required("Ports"+"."+k, "body", m.Ports[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Ports[k]); i++ {

			if err := m.Ports[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Ports" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InspectNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectNetworkSettings) UnmarshalBinary(b []byte) error {
	var res InspectNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
