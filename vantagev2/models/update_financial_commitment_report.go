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

// UpdateFinancialCommitmentReport Update a FinancialCommitmentReport.
//
// swagger:model updateFinancialCommitmentReport
type UpdateFinancialCommitmentReport struct {

	// The date bucket of the FinancialCommitmentReport.
	// Enum: ["day","week","month"]
	DateBucket string `json:"date_bucket,omitempty"`

	// The date interval of the FinancialCommitmentReport. Incompatible with 'start_date' and 'end_date' parameters. Defaults to 'this_month' if start_date and end_date are not provided.
	// Enum: ["this_month","last_7_days","last_30_days","last_month","last_3_months","last_6_months","custom","last_12_months","last_24_months","last_36_months","year_to_date","last_3_days"]
	DateInterval string `json:"date_interval,omitempty"`

	// The end date of the FinancialCommitmentReport. YYYY-MM-DD formatted. Incompatible with 'date_interval' parameter.
	// Example: 2024-03-01
	// Format: date
	EndDate strfmt.Date `json:"end_date,omitempty"`

	// The filter query language to apply to the FinancialCommitmentReport. Additional documentation available at https://docs.vantage.sh/vql.
	Filter string `json:"filter,omitempty"`

	// Grouping values for aggregating costs on the FinancialCommitmentReport. Valid groupings: cost_type, commitment_type, service, resource_account_id, provider_account_id, region, cost_category, cost_sub_category, instance_type, tag, label:<label_name>.
	Groupings []string `json:"groupings"`

	// The scope for the costs. Possible values: discountable, all.
	// Enum: ["discountable","all"]
	OnDemandCostsScope string `json:"on_demand_costs_scope,omitempty"`

	// The start date of the FinancialCommitmentReport. YYYY-MM-DD formatted. Incompatible with 'date_interval' parameter.
	// Example: 2024-03-01
	// Format: date
	StartDate strfmt.Date `json:"start_date,omitempty"`

	// The title of the FinancialCommitmentReport.
	Title string `json:"title,omitempty"`
}

// Validate validates this update financial commitment report
func (m *UpdateFinancialCommitmentReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnDemandCostsScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateFinancialCommitmentReportTypeDateBucketPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["day","week","month"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateFinancialCommitmentReportTypeDateBucketPropEnum = append(updateFinancialCommitmentReportTypeDateBucketPropEnum, v)
	}
}

const (

	// UpdateFinancialCommitmentReportDateBucketDay captures enum value "day"
	UpdateFinancialCommitmentReportDateBucketDay string = "day"

	// UpdateFinancialCommitmentReportDateBucketWeek captures enum value "week"
	UpdateFinancialCommitmentReportDateBucketWeek string = "week"

	// UpdateFinancialCommitmentReportDateBucketMonth captures enum value "month"
	UpdateFinancialCommitmentReportDateBucketMonth string = "month"
)

// prop value enum
func (m *UpdateFinancialCommitmentReport) validateDateBucketEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateFinancialCommitmentReportTypeDateBucketPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateFinancialCommitmentReport) validateDateBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.DateBucket) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateBucketEnum("date_bucket", "body", m.DateBucket); err != nil {
		return err
	}

	return nil
}

var updateFinancialCommitmentReportTypeDateIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["this_month","last_7_days","last_30_days","last_month","last_3_months","last_6_months","custom","last_12_months","last_24_months","last_36_months","year_to_date","last_3_days"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateFinancialCommitmentReportTypeDateIntervalPropEnum = append(updateFinancialCommitmentReportTypeDateIntervalPropEnum, v)
	}
}

const (

	// UpdateFinancialCommitmentReportDateIntervalThisMonth captures enum value "this_month"
	UpdateFinancialCommitmentReportDateIntervalThisMonth string = "this_month"

	// UpdateFinancialCommitmentReportDateIntervalLast7Days captures enum value "last_7_days"
	UpdateFinancialCommitmentReportDateIntervalLast7Days string = "last_7_days"

	// UpdateFinancialCommitmentReportDateIntervalLast30Days captures enum value "last_30_days"
	UpdateFinancialCommitmentReportDateIntervalLast30Days string = "last_30_days"

	// UpdateFinancialCommitmentReportDateIntervalLastMonth captures enum value "last_month"
	UpdateFinancialCommitmentReportDateIntervalLastMonth string = "last_month"

	// UpdateFinancialCommitmentReportDateIntervalLast3Months captures enum value "last_3_months"
	UpdateFinancialCommitmentReportDateIntervalLast3Months string = "last_3_months"

	// UpdateFinancialCommitmentReportDateIntervalLast6Months captures enum value "last_6_months"
	UpdateFinancialCommitmentReportDateIntervalLast6Months string = "last_6_months"

	// UpdateFinancialCommitmentReportDateIntervalCustom captures enum value "custom"
	UpdateFinancialCommitmentReportDateIntervalCustom string = "custom"

	// UpdateFinancialCommitmentReportDateIntervalLast12Months captures enum value "last_12_months"
	UpdateFinancialCommitmentReportDateIntervalLast12Months string = "last_12_months"

	// UpdateFinancialCommitmentReportDateIntervalLast24Months captures enum value "last_24_months"
	UpdateFinancialCommitmentReportDateIntervalLast24Months string = "last_24_months"

	// UpdateFinancialCommitmentReportDateIntervalLast36Months captures enum value "last_36_months"
	UpdateFinancialCommitmentReportDateIntervalLast36Months string = "last_36_months"

	// UpdateFinancialCommitmentReportDateIntervalYearToDate captures enum value "year_to_date"
	UpdateFinancialCommitmentReportDateIntervalYearToDate string = "year_to_date"

	// UpdateFinancialCommitmentReportDateIntervalLast3Days captures enum value "last_3_days"
	UpdateFinancialCommitmentReportDateIntervalLast3Days string = "last_3_days"
)

// prop value enum
func (m *UpdateFinancialCommitmentReport) validateDateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateFinancialCommitmentReportTypeDateIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateFinancialCommitmentReport) validateDateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.DateInterval) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateIntervalEnum("date_interval", "body", m.DateInterval); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFinancialCommitmentReport) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("end_date", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var updateFinancialCommitmentReportTypeOnDemandCostsScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["discountable","all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateFinancialCommitmentReportTypeOnDemandCostsScopePropEnum = append(updateFinancialCommitmentReportTypeOnDemandCostsScopePropEnum, v)
	}
}

const (

	// UpdateFinancialCommitmentReportOnDemandCostsScopeDiscountable captures enum value "discountable"
	UpdateFinancialCommitmentReportOnDemandCostsScopeDiscountable string = "discountable"

	// UpdateFinancialCommitmentReportOnDemandCostsScopeAll captures enum value "all"
	UpdateFinancialCommitmentReportOnDemandCostsScopeAll string = "all"
)

// prop value enum
func (m *UpdateFinancialCommitmentReport) validateOnDemandCostsScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateFinancialCommitmentReportTypeOnDemandCostsScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateFinancialCommitmentReport) validateOnDemandCostsScope(formats strfmt.Registry) error {
	if swag.IsZero(m.OnDemandCostsScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateOnDemandCostsScopeEnum("on_demand_costs_scope", "body", m.OnDemandCostsScope); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFinancialCommitmentReport) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update financial commitment report based on context it is used
func (m *UpdateFinancialCommitmentReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateFinancialCommitmentReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateFinancialCommitmentReport) UnmarshalBinary(b []byte) error {
	var res UpdateFinancialCommitmentReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
