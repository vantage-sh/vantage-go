// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Tag Tag model
//
// swagger:model Tag
type Tag struct {

	// Whether the Tag has been hidden from the Vantage UI.
	Hidden bool `json:"hidden,omitempty"`

	// The unique providers that are covered by the Tag key.
	Providers []string `json:"providers"`

	// The Tag key.
	// Example: aws:createdBy
	TagKey string `json:"tag_key,omitempty"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tag based on context it is used
func (m *Tag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Tag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag) UnmarshalBinary(b []byte) error {
	var res Tag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
