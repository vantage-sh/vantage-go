// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Integration Integration model
//
// swagger:model Integration
type Integration struct {

	// The account identifier. For GCP this is the billing Account ID, for Azure this is the account ID
	// Example: 011389-EF4C3E-3ED7AE
	AccountIdentifier string `json:"account_identifier,omitempty"`

	// The date and time, in UTC, the Integration was created. ISO 8601 Formatted.
	// Example: 2023-08-04T00:00:00Z
	CreatedAt string `json:"created_at,omitempty"`

	// The name of the Integration.
	// Example: AWS
	Provider string `json:"provider,omitempty"`

	// The status of the Integration. Can be 'connected' or 'disconnected'.
	// Example: connected
	// Enum: [connected disconnected]
	Status string `json:"status,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// The tokens for any Workspaces that the account belongs to.
	WorkspaceTokens []string `json:"workspace_tokens"`
}

// Validate validates this integration
func (m *Integration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var integrationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","disconnected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationTypeStatusPropEnum = append(integrationTypeStatusPropEnum, v)
	}
}

const (

	// IntegrationStatusConnected captures enum value "connected"
	IntegrationStatusConnected string = "connected"

	// IntegrationStatusDisconnected captures enum value "disconnected"
	IntegrationStatusDisconnected string = "disconnected"
)

// prop value enum
func (m *Integration) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, integrationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Integration) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this integration based on context it is used
func (m *Integration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Integration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Integration) UnmarshalBinary(b []byte) error {
	var res Integration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
