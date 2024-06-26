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

// Budgets Budgets model
//
// swagger:model Budgets
type Budgets struct {

	// budgets
	Budgets []*Budget `json:"budgets"`

	// links
	Links interface{} `json:"links,omitempty"`
}

// Validate validates this budgets
func (m *Budgets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBudgets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Budgets) validateBudgets(formats strfmt.Registry) error {
	if swag.IsZero(m.Budgets) { // not required
		return nil
	}

	for i := 0; i < len(m.Budgets); i++ {
		if swag.IsZero(m.Budgets[i]) { // not required
			continue
		}

		if m.Budgets[i] != nil {
			if err := m.Budgets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("budgets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("budgets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this budgets based on the context it is used
func (m *Budgets) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBudgets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Budgets) contextValidateBudgets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Budgets); i++ {

		if m.Budgets[i] != nil {

			if swag.IsZero(m.Budgets[i]) { // not required
				return nil
			}

			if err := m.Budgets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("budgets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("budgets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Budgets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Budgets) UnmarshalBinary(b []byte) error {
	var res Budgets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
