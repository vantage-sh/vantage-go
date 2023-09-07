// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Price Price model
//
// swagger:model Price
type Price struct {

	// The amount of money this specific product price costs.
	// Example: 1.324
	Amount float64 `json:"amount,omitempty"`

	// The currency of the amount.
	// Example: USD
	Currency string `json:"currency,omitempty"`

	// Service specific metadata.
	// Example: {"lifecycle":"on-demand","platform":"linux-enterprise"}
	Details interface{} `json:"details,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// The part of the product the price applies to. (compute, transfer, etc..)
	// Example: compute
	RateType string `json:"rate_type,omitempty"`

	// The region the price is specific to.
	// Example: us-east-1
	Region string `json:"region,omitempty"`

	// The unit in which the amount is billed.
	// Example: hour
	Unit string `json:"unit,omitempty"`
}

// Validate validates this price
func (m *Price) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this price based on context it is used
func (m *Price) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Price) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Price) UnmarshalBinary(b []byte) error {
	var res Price
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}