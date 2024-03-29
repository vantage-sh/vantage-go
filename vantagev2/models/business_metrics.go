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

// BusinessMetrics BusinessMetrics model
//
// swagger:model BusinessMetrics
type BusinessMetrics struct {

	// business metrics
	BusinessMetrics []*BusinessMetric `json:"business_metrics"`
}

// Validate validates this business metrics
func (m *BusinessMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BusinessMetrics) validateBusinessMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.BusinessMetrics) { // not required
		return nil
	}

	for i := 0; i < len(m.BusinessMetrics); i++ {
		if swag.IsZero(m.BusinessMetrics[i]) { // not required
			continue
		}

		if m.BusinessMetrics[i] != nil {
			if err := m.BusinessMetrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("business_metrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("business_metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this business metrics based on the context it is used
func (m *BusinessMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBusinessMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BusinessMetrics) contextValidateBusinessMetrics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BusinessMetrics); i++ {

		if m.BusinessMetrics[i] != nil {

			if swag.IsZero(m.BusinessMetrics[i]) { // not required
				return nil
			}

			if err := m.BusinessMetrics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("business_metrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("business_metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BusinessMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BusinessMetrics) UnmarshalBinary(b []byte) error {
	var res BusinessMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
