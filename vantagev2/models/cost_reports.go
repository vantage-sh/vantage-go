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

// CostReports Cost Reports model
//
// swagger:model Cost Reports
type CostReports struct {

	// cost reports
	CostReports []*CostReport `json:"cost_reports"`

	// links
	Links interface{} `json:"links,omitempty"`
}

// Validate validates this cost reports
func (m *CostReports) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostReports(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CostReports) validateCostReports(formats strfmt.Registry) error {
	if swag.IsZero(m.CostReports) { // not required
		return nil
	}

	for i := 0; i < len(m.CostReports); i++ {
		if swag.IsZero(m.CostReports[i]) { // not required
			continue
		}

		if m.CostReports[i] != nil {
			if err := m.CostReports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_reports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cost_reports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cost reports based on the context it is used
func (m *CostReports) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCostReports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CostReports) contextValidateCostReports(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CostReports); i++ {

		if m.CostReports[i] != nil {

			if swag.IsZero(m.CostReports[i]) { // not required
				return nil
			}

			if err := m.CostReports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_reports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cost_reports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CostReports) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostReports) UnmarshalBinary(b []byte) error {
	var res CostReports
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
