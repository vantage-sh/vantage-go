// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TagValue tag value
//
// swagger:model TagValue
type TagValue struct {

	// The unique providers that are covered by the TagValue.
	Providers []string `json:"providers"`

	// The TagValue.
	// Example: vantage
	TagValue string `json:"tag_value,omitempty"`
}

// Validate validates this tag value
func (m *TagValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tag value based on context it is used
func (m *TagValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagValue) UnmarshalBinary(b []byte) error {
	var res TagValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
