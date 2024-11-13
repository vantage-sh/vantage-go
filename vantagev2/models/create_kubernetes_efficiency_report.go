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

// CreateKubernetesEfficiencyReport Create a KubernetesEfficiencyReport.
//
// swagger:model createKubernetesEfficiencyReport
type CreateKubernetesEfficiencyReport struct {

	// The column by which the costs are aggregated.
	// Enum: ["idle_cost","amount","cost_efficiency"]
	AggregatedBy string `json:"aggregated_by,omitempty"`

	// The date bin of the KubernetesEfficiencyReport.
	// Enum: ["cumulative","day","week","month"]
	DateBin *string `json:"date_bin,omitempty"`

	// The date interval of the KubernetesEfficiencyReport. Incompatible with 'start_date' and 'end_date' parameters. Defaults to 'this_month' if start_date and end_date are not provided.
	// Enum: ["this_month","last_7_days","last_30_days","last_month","last_3_months","last_6_months","custom","last_12_months","last_24_months","last_36_months","next_month","next_3_months","next_6_months","next_12_months","year_to_date"]
	DateInterval string `json:"date_interval,omitempty"`

	// The end date of the KubernetesEfficiencyReport. ISO 8601 Formatted. Incompatible with 'date_interval' parameter.
	// Required: true
	// Format: date
	EndDate *strfmt.Date `json:"end_date"`

	// The filter query language to apply to the KubernetesEfficiencyReport. Additional documentation available at https://docs.vantage.sh/vql.
	Filter string `json:"filter,omitempty"`

	// Grouping values for aggregating costs on the KubernetesEfficiencyReport. Valid groupings: cluster_id, namespace, labeled, category, label, label:<label_name>.
	Groupings []string `json:"groupings"`

	// The start date of the KubernetesEfficiencyReport. ISO 8601 Formatted. Incompatible with 'date_interval' parameter.
	// Required: true
	// Format: date
	StartDate *strfmt.Date `json:"start_date"`

	// The title of the KubernetesEfficiencyReport.
	// Required: true
	Title *string `json:"title"`

	// The Workspace in which the KubernetesEfficiencyReport will be created.
	// Required: true
	WorkspaceToken *string `json:"workspace_token"`
}

// Validate validates this create kubernetes efficiency report
func (m *CreateKubernetesEfficiencyReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateBin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaceToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createKubernetesEfficiencyReportTypeAggregatedByPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["idle_cost","amount","cost_efficiency"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createKubernetesEfficiencyReportTypeAggregatedByPropEnum = append(createKubernetesEfficiencyReportTypeAggregatedByPropEnum, v)
	}
}

const (

	// CreateKubernetesEfficiencyReportAggregatedByIdleCost captures enum value "idle_cost"
	CreateKubernetesEfficiencyReportAggregatedByIdleCost string = "idle_cost"

	// CreateKubernetesEfficiencyReportAggregatedByAmount captures enum value "amount"
	CreateKubernetesEfficiencyReportAggregatedByAmount string = "amount"

	// CreateKubernetesEfficiencyReportAggregatedByCostEfficiency captures enum value "cost_efficiency"
	CreateKubernetesEfficiencyReportAggregatedByCostEfficiency string = "cost_efficiency"
)

// prop value enum
func (m *CreateKubernetesEfficiencyReport) validateAggregatedByEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createKubernetesEfficiencyReportTypeAggregatedByPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateKubernetesEfficiencyReport) validateAggregatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.AggregatedBy) { // not required
		return nil
	}

	// value enum
	if err := m.validateAggregatedByEnum("aggregated_by", "body", m.AggregatedBy); err != nil {
		return err
	}

	return nil
}

var createKubernetesEfficiencyReportTypeDateBinPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cumulative","day","week","month"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createKubernetesEfficiencyReportTypeDateBinPropEnum = append(createKubernetesEfficiencyReportTypeDateBinPropEnum, v)
	}
}

const (

	// CreateKubernetesEfficiencyReportDateBinCumulative captures enum value "cumulative"
	CreateKubernetesEfficiencyReportDateBinCumulative string = "cumulative"

	// CreateKubernetesEfficiencyReportDateBinDay captures enum value "day"
	CreateKubernetesEfficiencyReportDateBinDay string = "day"

	// CreateKubernetesEfficiencyReportDateBinWeek captures enum value "week"
	CreateKubernetesEfficiencyReportDateBinWeek string = "week"

	// CreateKubernetesEfficiencyReportDateBinMonth captures enum value "month"
	CreateKubernetesEfficiencyReportDateBinMonth string = "month"
)

// prop value enum
func (m *CreateKubernetesEfficiencyReport) validateDateBinEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createKubernetesEfficiencyReportTypeDateBinPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateKubernetesEfficiencyReport) validateDateBin(formats strfmt.Registry) error {
	if swag.IsZero(m.DateBin) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateBinEnum("date_bin", "body", *m.DateBin); err != nil {
		return err
	}

	return nil
}

var createKubernetesEfficiencyReportTypeDateIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["this_month","last_7_days","last_30_days","last_month","last_3_months","last_6_months","custom","last_12_months","last_24_months","last_36_months","next_month","next_3_months","next_6_months","next_12_months","year_to_date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createKubernetesEfficiencyReportTypeDateIntervalPropEnum = append(createKubernetesEfficiencyReportTypeDateIntervalPropEnum, v)
	}
}

const (

	// CreateKubernetesEfficiencyReportDateIntervalThisMonth captures enum value "this_month"
	CreateKubernetesEfficiencyReportDateIntervalThisMonth string = "this_month"

	// CreateKubernetesEfficiencyReportDateIntervalLast7Days captures enum value "last_7_days"
	CreateKubernetesEfficiencyReportDateIntervalLast7Days string = "last_7_days"

	// CreateKubernetesEfficiencyReportDateIntervalLast30Days captures enum value "last_30_days"
	CreateKubernetesEfficiencyReportDateIntervalLast30Days string = "last_30_days"

	// CreateKubernetesEfficiencyReportDateIntervalLastMonth captures enum value "last_month"
	CreateKubernetesEfficiencyReportDateIntervalLastMonth string = "last_month"

	// CreateKubernetesEfficiencyReportDateIntervalLast3Months captures enum value "last_3_months"
	CreateKubernetesEfficiencyReportDateIntervalLast3Months string = "last_3_months"

	// CreateKubernetesEfficiencyReportDateIntervalLast6Months captures enum value "last_6_months"
	CreateKubernetesEfficiencyReportDateIntervalLast6Months string = "last_6_months"

	// CreateKubernetesEfficiencyReportDateIntervalCustom captures enum value "custom"
	CreateKubernetesEfficiencyReportDateIntervalCustom string = "custom"

	// CreateKubernetesEfficiencyReportDateIntervalLast12Months captures enum value "last_12_months"
	CreateKubernetesEfficiencyReportDateIntervalLast12Months string = "last_12_months"

	// CreateKubernetesEfficiencyReportDateIntervalLast24Months captures enum value "last_24_months"
	CreateKubernetesEfficiencyReportDateIntervalLast24Months string = "last_24_months"

	// CreateKubernetesEfficiencyReportDateIntervalLast36Months captures enum value "last_36_months"
	CreateKubernetesEfficiencyReportDateIntervalLast36Months string = "last_36_months"

	// CreateKubernetesEfficiencyReportDateIntervalNextMonth captures enum value "next_month"
	CreateKubernetesEfficiencyReportDateIntervalNextMonth string = "next_month"

	// CreateKubernetesEfficiencyReportDateIntervalNext3Months captures enum value "next_3_months"
	CreateKubernetesEfficiencyReportDateIntervalNext3Months string = "next_3_months"

	// CreateKubernetesEfficiencyReportDateIntervalNext6Months captures enum value "next_6_months"
	CreateKubernetesEfficiencyReportDateIntervalNext6Months string = "next_6_months"

	// CreateKubernetesEfficiencyReportDateIntervalNext12Months captures enum value "next_12_months"
	CreateKubernetesEfficiencyReportDateIntervalNext12Months string = "next_12_months"

	// CreateKubernetesEfficiencyReportDateIntervalYearToDate captures enum value "year_to_date"
	CreateKubernetesEfficiencyReportDateIntervalYearToDate string = "year_to_date"
)

// prop value enum
func (m *CreateKubernetesEfficiencyReport) validateDateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createKubernetesEfficiencyReportTypeDateIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateKubernetesEfficiencyReport) validateDateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.DateInterval) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateIntervalEnum("date_interval", "body", m.DateInterval); err != nil {
		return err
	}

	return nil
}

func (m *CreateKubernetesEfficiencyReport) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("end_date", "body", m.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("end_date", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateKubernetesEfficiencyReport) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("start_date", "body", m.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("start_date", "body", "date", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateKubernetesEfficiencyReport) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *CreateKubernetesEfficiencyReport) validateWorkspaceToken(formats strfmt.Registry) error {

	if err := validate.Required("workspace_token", "body", m.WorkspaceToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create kubernetes efficiency report based on context it is used
func (m *CreateKubernetesEfficiencyReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateKubernetesEfficiencyReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateKubernetesEfficiencyReport) UnmarshalBinary(b []byte) error {
	var res CreateKubernetesEfficiencyReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
