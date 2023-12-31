// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Costs Costs model
//
// swagger:model Costs
type Costs struct {

	// costs
	Costs []string `json:"costs"`

	// The currency both the total and cost amount values are represented in. e.g. USD
	Currency string `json:"currency,omitempty"`

	// links
	Links interface{} `json:"links,omitempty"`

	// The sum of all amounts of costs for the entire report for the requested period.
	Total string `json:"total,omitempty"`
}

// Validate validates this costs
func (m *Costs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this costs based on context it is used
func (m *Costs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Costs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Costs) UnmarshalBinary(b []byte) error {
	var res Costs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
