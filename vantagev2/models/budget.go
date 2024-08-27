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

// Budget Budget model
//
// swagger:model Budget
type Budget struct {

	// The tokens of the BudgetAlerts associated with the Budget.
	BudgetAlertTokens []string `json:"budget_alert_tokens"`

	// The token of the Report associated with the Budget.
	CostReportToken string `json:"cost_report_token,omitempty"`

	// The date and time, in UTC, the Budget was created. ISO 8601 Formatted.
	// Example: 2024-03-19T00:00:00Z
	CreatedAt string `json:"created_at,omitempty"`

	// The token of the Creator of the Budget.
	CreatedByToken string `json:"created_by_token,omitempty"`

	// The name of the Budget.
	// Example: Acme123 Budget
	Name string `json:"name,omitempty"`

	// The historical performance of the Budget.
	Performance []*BudgetPerformance `json:"performance"`

	// The budget periods associated with the Budget.
	Periods []*BudgetPeriod `json:"periods"`

	// token
	Token string `json:"token,omitempty"`

	// The token for the User who created this Budget.
	UserToken string `json:"user_token,omitempty"`

	// The token for the Workspace the Budget is a part of.
	WorkspaceToken string `json:"workspace_token,omitempty"`
}

// Validate validates this budget
func (m *Budget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePerformance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Budget) validatePerformance(formats strfmt.Registry) error {
	if swag.IsZero(m.Performance) { // not required
		return nil
	}

	for i := 0; i < len(m.Performance); i++ {
		if swag.IsZero(m.Performance[i]) { // not required
			continue
		}

		if m.Performance[i] != nil {
			if err := m.Performance[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("performance" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("performance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Budget) validatePeriods(formats strfmt.Registry) error {
	if swag.IsZero(m.Periods) { // not required
		return nil
	}

	for i := 0; i < len(m.Periods); i++ {
		if swag.IsZero(m.Periods[i]) { // not required
			continue
		}

		if m.Periods[i] != nil {
			if err := m.Periods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("periods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("periods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this budget based on the context it is used
func (m *Budget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePerformance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeriods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Budget) contextValidatePerformance(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Performance); i++ {

		if m.Performance[i] != nil {

			if swag.IsZero(m.Performance[i]) { // not required
				return nil
			}

			if err := m.Performance[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("performance" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("performance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Budget) contextValidatePeriods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Periods); i++ {

		if m.Periods[i] != nil {

			if swag.IsZero(m.Periods[i]) { // not required
				return nil
			}

			if err := m.Periods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("periods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("periods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Budget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Budget) UnmarshalBinary(b []byte) error {
	var res Budget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
