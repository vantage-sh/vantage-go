// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutBusinessMetrics Updates an existing BusinessMetric.
//
// swagger:model putBusinessMetrics
type PutBusinessMetrics struct {

	// The tokens for any CostReports that use the BusinessMetric, and the unit scale.
	CostReportTokensWithMetadata []*PutBusinessMetricsCostReportTokensWithMetadataItems0 `json:"cost_report_tokens_with_metadata"`

	// The title of the BusinessMetric.
	Title string `json:"title,omitempty"`

	// The dates and amounts for the BusinessMetric.
	Values []*PutBusinessMetricsValuesItems0 `json:"values"`
}

// Validate validates this put business metrics
func (m *PutBusinessMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostReportTokensWithMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutBusinessMetrics) validateCostReportTokensWithMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.CostReportTokensWithMetadata) { // not required
		return nil
	}

	for i := 0; i < len(m.CostReportTokensWithMetadata); i++ {
		if swag.IsZero(m.CostReportTokensWithMetadata[i]) { // not required
			continue
		}

		if m.CostReportTokensWithMetadata[i] != nil {
			if err := m.CostReportTokensWithMetadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_report_tokens_with_metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cost_report_tokens_with_metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutBusinessMetrics) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this put business metrics based on the context it is used
func (m *PutBusinessMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCostReportTokensWithMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutBusinessMetrics) contextValidateCostReportTokensWithMetadata(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CostReportTokensWithMetadata); i++ {

		if m.CostReportTokensWithMetadata[i] != nil {

			if swag.IsZero(m.CostReportTokensWithMetadata[i]) { // not required
				return nil
			}

			if err := m.CostReportTokensWithMetadata[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_report_tokens_with_metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cost_report_tokens_with_metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutBusinessMetrics) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Values); i++ {

		if m.Values[i] != nil {

			if swag.IsZero(m.Values[i]) { // not required
				return nil
			}

			if err := m.Values[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutBusinessMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutBusinessMetrics) UnmarshalBinary(b []byte) error {
	var res PutBusinessMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutBusinessMetricsCostReportTokensWithMetadataItems0 put business metrics cost report tokens with metadata items0
//
// swagger:model PutBusinessMetricsCostReportTokensWithMetadataItems0
type PutBusinessMetricsCostReportTokensWithMetadataItems0 struct {

	// The token of the CostReport the BusinessMetric is attached to.
	// Required: true
	CostReportToken *string `json:"cost_report_token"`

	// Determines the scale of the BusinessMetric's values within the CostReport.
	// Enum: [per_unit per_hundred per_thousand per_million per_billion]
	UnitScale *string `json:"unit_scale,omitempty"`
}

// Validate validates this put business metrics cost report tokens with metadata items0
func (m *PutBusinessMetricsCostReportTokensWithMetadataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostReportToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitScale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutBusinessMetricsCostReportTokensWithMetadataItems0) validateCostReportToken(formats strfmt.Registry) error {

	if err := validate.Required("cost_report_token", "body", m.CostReportToken); err != nil {
		return err
	}

	return nil
}

var putBusinessMetricsCostReportTokensWithMetadataItems0TypeUnitScalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["per_unit","per_hundred","per_thousand","per_million","per_billion"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putBusinessMetricsCostReportTokensWithMetadataItems0TypeUnitScalePropEnum = append(putBusinessMetricsCostReportTokensWithMetadataItems0TypeUnitScalePropEnum, v)
	}
}

const (

	// PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerUnit captures enum value "per_unit"
	PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerUnit string = "per_unit"

	// PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerHundred captures enum value "per_hundred"
	PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerHundred string = "per_hundred"

	// PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerThousand captures enum value "per_thousand"
	PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerThousand string = "per_thousand"

	// PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerMillion captures enum value "per_million"
	PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerMillion string = "per_million"

	// PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerBillion captures enum value "per_billion"
	PutBusinessMetricsCostReportTokensWithMetadataItems0UnitScalePerBillion string = "per_billion"
)

// prop value enum
func (m *PutBusinessMetricsCostReportTokensWithMetadataItems0) validateUnitScaleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, putBusinessMetricsCostReportTokensWithMetadataItems0TypeUnitScalePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PutBusinessMetricsCostReportTokensWithMetadataItems0) validateUnitScale(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitScale) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitScaleEnum("unit_scale", "body", *m.UnitScale); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put business metrics cost report tokens with metadata items0 based on context it is used
func (m *PutBusinessMetricsCostReportTokensWithMetadataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutBusinessMetricsCostReportTokensWithMetadataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutBusinessMetricsCostReportTokensWithMetadataItems0) UnmarshalBinary(b []byte) error {
	var res PutBusinessMetricsCostReportTokensWithMetadataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutBusinessMetricsValuesItems0 put business metrics values items0
//
// swagger:model PutBusinessMetricsValuesItems0
type PutBusinessMetricsValuesItems0 struct {

	// amount
	// Required: true
	Amount *float64 `json:"amount"`

	// date
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`
}

// Validate validates this put business metrics values items0
func (m *PutBusinessMetricsValuesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutBusinessMetricsValuesItems0) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *PutBusinessMetricsValuesItems0) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put business metrics values items0 based on context it is used
func (m *PutBusinessMetricsValuesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutBusinessMetricsValuesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutBusinessMetricsValuesItems0) UnmarshalBinary(b []byte) error {
	var res PutBusinessMetricsValuesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
