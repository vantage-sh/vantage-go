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

// BudgetAlerts BudgetAlerts model
//
// swagger:model BudgetAlerts
type BudgetAlerts struct {

	// budget alerts
	BudgetAlerts []*BudgetAlert `json:"budget_alerts"`

	// links
	Links interface{} `json:"links,omitempty"`
}

// Validate validates this budget alerts
func (m *BudgetAlerts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBudgetAlerts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BudgetAlerts) validateBudgetAlerts(formats strfmt.Registry) error {
	if swag.IsZero(m.BudgetAlerts) { // not required
		return nil
	}

	for i := 0; i < len(m.BudgetAlerts); i++ {
		if swag.IsZero(m.BudgetAlerts[i]) { // not required
			continue
		}

		if m.BudgetAlerts[i] != nil {
			if err := m.BudgetAlerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("budget_alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("budget_alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this budget alerts based on the context it is used
func (m *BudgetAlerts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBudgetAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BudgetAlerts) contextValidateBudgetAlerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BudgetAlerts); i++ {

		if m.BudgetAlerts[i] != nil {

			if swag.IsZero(m.BudgetAlerts[i]) { // not required
				return nil
			}

			if err := m.BudgetAlerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("budget_alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("budget_alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BudgetAlerts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BudgetAlerts) UnmarshalBinary(b []byte) error {
	var res BudgetAlerts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
