// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceReport ResourceReport model
//
// swagger:model ResourceReport
type ResourceReport struct {

	// The date and time, in UTC, the report was created. ISO 8601 Formatted.
	// Example: 2024-03-19T00:00:00Z
	CreatedAt string `json:"created_at,omitempty"`

	// The token for the User or Team who created this ResourceReport.
	CreatedByToken string `json:"created_by_token,omitempty"`

	// Indicates whether the ResourceReport is the default report.
	Default bool `json:"default,omitempty"`

	// The filter applied to the ResourceReport. Additional documentation available at https://docs.vantage.sh/vql.
	Filter string `json:"filter,omitempty"`

	// The token for the Segment the ResourceReport is a part of.
	SegmentToken string `json:"segment_token,omitempty"`

	// The title of the ResourceReport.
	// Example: Acme123 Active Resources
	Title string `json:"title,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// The token for the User who created this ResourceReport.
	UserToken string `json:"user_token,omitempty"`

	// The token for the Workspace the ResourceReport is a part of.
	WorkspaceToken string `json:"workspace_token,omitempty"`
}

// Validate validates this resource report
func (m *ResourceReport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this resource report based on context it is used
func (m *ResourceReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceReport) UnmarshalBinary(b []byte) error {
	var res ResourceReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
