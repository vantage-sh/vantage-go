// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Provider provider
//
// swagger:model Provider
type Provider struct {

	// The full descriptive name of the provider.
	// Example: Amazon Web Services
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// The common name of the provider.
	// Example: AWS
	Name string `json:"name,omitempty"`
}

// Validate validates this provider
func (m *Provider) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this provider based on context it is used
func (m *Provider) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Provider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Provider) UnmarshalBinary(b []byte) error {
	var res Provider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}