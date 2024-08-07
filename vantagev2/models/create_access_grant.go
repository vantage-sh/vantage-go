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

// CreateAccessGrant Create an Access Grant.
//
// swagger:model createAccessGrant
type CreateAccessGrant struct {

	// The access level you want to grant. Defaults to 'allowed'.
	// Enum: ["denied","allowed"]
	Access string `json:"access,omitempty"`

	// The token of the resource for which you are granting access.
	// Required: true
	ResourceToken *string `json:"resource_token"`

	// The token of the Team you want to grant access to.
	// Required: true
	TeamToken *string `json:"team_token"`
}

// Validate validates this create access grant
func (m *CreateAccessGrant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createAccessGrantTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["denied","allowed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAccessGrantTypeAccessPropEnum = append(createAccessGrantTypeAccessPropEnum, v)
	}
}

const (

	// CreateAccessGrantAccessDenied captures enum value "denied"
	CreateAccessGrantAccessDenied string = "denied"

	// CreateAccessGrantAccessAllowed captures enum value "allowed"
	CreateAccessGrantAccessAllowed string = "allowed"
)

// prop value enum
func (m *CreateAccessGrant) validateAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAccessGrantTypeAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateAccessGrant) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessEnum("access", "body", m.Access); err != nil {
		return err
	}

	return nil
}

func (m *CreateAccessGrant) validateResourceToken(formats strfmt.Registry) error {

	if err := validate.Required("resource_token", "body", m.ResourceToken); err != nil {
		return err
	}

	return nil
}

func (m *CreateAccessGrant) validateTeamToken(formats strfmt.Registry) error {

	if err := validate.Required("team_token", "body", m.TeamToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create access grant based on context it is used
func (m *CreateAccessGrant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAccessGrant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAccessGrant) UnmarshalBinary(b []byte) error {
	var res CreateAccessGrant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
