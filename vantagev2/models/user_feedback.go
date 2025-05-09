// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserFeedback UserFeedback model
//
// swagger:model UserFeedback
type UserFeedback struct {

	// Feedback creation timestamp
	// Example: 2023-01-01T00:00:00Z
	CreatedAt string `json:"created_at,omitempty"`

	// Token of the creator of the feedback
	CreatedByToken string `json:"created_by_token,omitempty"`

	// User feedback message
	Message string `json:"message,omitempty"`

	// Token of the feedback
	Token string `json:"token,omitempty"`
}

// Validate validates this user feedback
func (m *UserFeedback) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user feedback based on context it is used
func (m *UserFeedback) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserFeedback) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserFeedback) UnmarshalBinary(b []byte) error {
	var res UserFeedback
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
