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

// UpdateTeam Update a Team.
//
// swagger:model updateTeam
type UpdateTeam struct {

	// The description of the Team.
	Description string `json:"description,omitempty"`

	// The name of the Team.
	Name string `json:"name,omitempty"`

	// The role to assign to the provided Users. Defaults to 'editor' which has editor permissions.
	// Enum: ["owner","editor","viewer"]
	Role string `json:"role,omitempty"`

	// The User emails to associate to the Team.
	UserEmails []string `json:"user_emails"`

	// The User tokens to associate to the Team.
	UserTokens []string `json:"user_tokens"`

	// The Workspace tokens to associate to the Team.
	WorkspaceTokens []string `json:"workspace_tokens"`
}

// Validate validates this update team
func (m *UpdateTeam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateTeamTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["owner","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateTeamTypeRolePropEnum = append(updateTeamTypeRolePropEnum, v)
	}
}

const (

	// UpdateTeamRoleOwner captures enum value "owner"
	UpdateTeamRoleOwner string = "owner"

	// UpdateTeamRoleEditor captures enum value "editor"
	UpdateTeamRoleEditor string = "editor"

	// UpdateTeamRoleViewer captures enum value "viewer"
	UpdateTeamRoleViewer string = "viewer"
)

// prop value enum
func (m *UpdateTeam) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateTeamTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateTeam) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update team based on context it is used
func (m *UpdateTeam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTeam) UnmarshalBinary(b []byte) error {
	var res UpdateTeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
