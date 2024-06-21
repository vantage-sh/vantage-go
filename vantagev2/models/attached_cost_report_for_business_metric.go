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

// AttachedCostReportForBusinessMetric attached cost report for business metric
//
// swagger:model AttachedCostReportForBusinessMetric
type AttachedCostReportForBusinessMetric struct {

	// The token of the CostReport the BusinessMetric is attached to.
	// Example: rprt_1234
	CostReportToken string `json:"cost_report_token,omitempty"`

	// The labels that the BusinessMetric is filtered by within a particular CostReport.
	// Example: ["label_1","label_2"]
	LabelFilter string `json:"label_filter,omitempty"`

	// Determines the scale of the BusinessMetric's values within a particular CostReport.
	// Example: per_hundred
	// Enum: ["per_unit","per_hundred","per_thousand","per_million","per_billion"]
	UnitScale string `json:"unit_scale,omitempty"`
}

// Validate validates this attached cost report for business metric
func (m *AttachedCostReportForBusinessMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnitScale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var attachedCostReportForBusinessMetricTypeUnitScalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["per_unit","per_hundred","per_thousand","per_million","per_billion"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attachedCostReportForBusinessMetricTypeUnitScalePropEnum = append(attachedCostReportForBusinessMetricTypeUnitScalePropEnum, v)
	}
}

const (

	// AttachedCostReportForBusinessMetricUnitScalePerUnit captures enum value "per_unit"
	AttachedCostReportForBusinessMetricUnitScalePerUnit string = "per_unit"

	// AttachedCostReportForBusinessMetricUnitScalePerHundred captures enum value "per_hundred"
	AttachedCostReportForBusinessMetricUnitScalePerHundred string = "per_hundred"

	// AttachedCostReportForBusinessMetricUnitScalePerThousand captures enum value "per_thousand"
	AttachedCostReportForBusinessMetricUnitScalePerThousand string = "per_thousand"

	// AttachedCostReportForBusinessMetricUnitScalePerMillion captures enum value "per_million"
	AttachedCostReportForBusinessMetricUnitScalePerMillion string = "per_million"

	// AttachedCostReportForBusinessMetricUnitScalePerBillion captures enum value "per_billion"
	AttachedCostReportForBusinessMetricUnitScalePerBillion string = "per_billion"
)

// prop value enum
func (m *AttachedCostReportForBusinessMetric) validateUnitScaleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, attachedCostReportForBusinessMetricTypeUnitScalePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttachedCostReportForBusinessMetric) validateUnitScale(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitScale) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitScaleEnum("unit_scale", "body", m.UnitScale); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this attached cost report for business metric based on context it is used
func (m *AttachedCostReportForBusinessMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttachedCostReportForBusinessMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachedCostReportForBusinessMetric) UnmarshalBinary(b []byte) error {
	var res AttachedCostReportForBusinessMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
