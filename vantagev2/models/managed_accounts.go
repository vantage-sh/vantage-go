// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManagedAccounts ManagedAccounts model
//
// swagger:model ManagedAccounts
type ManagedAccounts struct {

	// links
	Links interface{} `json:"links,omitempty"`

	// managed accounts
	ManagedAccounts []*ManagedAccount `json:"managed_accounts"`
}

// Validate validates this managed accounts
func (m *ManagedAccounts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagedAccounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagedAccounts) validateManagedAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagedAccounts) { // not required
		return nil
	}

	for i := 0; i < len(m.ManagedAccounts); i++ {
		if swag.IsZero(m.ManagedAccounts[i]) { // not required
			continue
		}

		if m.ManagedAccounts[i] != nil {
			if err := m.ManagedAccounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("managed_accounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("managed_accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this managed accounts based on the context it is used
func (m *ManagedAccounts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManagedAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagedAccounts) contextValidateManagedAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ManagedAccounts); i++ {

		if m.ManagedAccounts[i] != nil {

			if swag.IsZero(m.ManagedAccounts[i]) { // not required
				return nil
			}

			if err := m.ManagedAccounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("managed_accounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("managed_accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagedAccounts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagedAccounts) UnmarshalBinary(b []byte) error {
	var res ManagedAccounts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
