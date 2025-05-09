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

// UpdateDashboard Update a Dashboard.
//
// swagger:model updateDashboard
type UpdateDashboard struct {

	// Determines how to group costs in the Dashboard.
	// Enum: ["cumulative","day","week","month"]
	DateBin string `json:"date_bin,omitempty"`

	// Determines the date range in the Dashboard. Incompatible with 'start_date' and 'end_date' parameters.
	// Enum: ["this_month","last_7_days","last_30_days","last_month","last_3_months","last_6_months","custom","last_12_months","last_24_months","last_36_months","next_month","next_3_months","next_6_months","next_12_months","year_to_date","last_3_days",""]
	DateInterval string `json:"date_interval"`

	// The end date for the date range for costs in the Dashboard. ISO 8601 Formatted. Incompatible with 'date_interval' parameter.
	EndDate string `json:"end_date,omitempty"`

	// The tokens of the Saved Filters used in the Dashboard.
	SavedFilterTokens []string `json:"saved_filter_tokens"`

	// The start date for the date range for costs in the Dashboard. ISO 8601 Formatted. Incompatible with 'date_interval' parameter.
	StartDate string `json:"start_date,omitempty"`

	// The title of the Dashboard.
	Title string `json:"title,omitempty"`

	// The widgets to add to the Dashboard. Currently supports CostReport, ResourceReport, KubernetesEfficiencyReport, and FinancialCommitmentReport.
	Widgets []*UpdateDashboardWidgetsItems0 `json:"widgets"`

	// The token of the Workspace the Dashboard belongs to. Required if the API token is associated with multiple Workspaces.
	WorkspaceToken string `json:"workspace_token,omitempty"`
}

// Validate validates this update dashboard
func (m *UpdateDashboard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateBin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidgets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateDashboardTypeDateBinPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cumulative","day","week","month"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDashboardTypeDateBinPropEnum = append(updateDashboardTypeDateBinPropEnum, v)
	}
}

const (

	// UpdateDashboardDateBinCumulative captures enum value "cumulative"
	UpdateDashboardDateBinCumulative string = "cumulative"

	// UpdateDashboardDateBinDay captures enum value "day"
	UpdateDashboardDateBinDay string = "day"

	// UpdateDashboardDateBinWeek captures enum value "week"
	UpdateDashboardDateBinWeek string = "week"

	// UpdateDashboardDateBinMonth captures enum value "month"
	UpdateDashboardDateBinMonth string = "month"
)

// prop value enum
func (m *UpdateDashboard) validateDateBinEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateDashboardTypeDateBinPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateDashboard) validateDateBin(formats strfmt.Registry) error {
	if swag.IsZero(m.DateBin) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateBinEnum("date_bin", "body", m.DateBin); err != nil {
		return err
	}

	return nil
}

var updateDashboardTypeDateIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["this_month","last_7_days","last_30_days","last_month","last_3_months","last_6_months","custom","last_12_months","last_24_months","last_36_months","next_month","next_3_months","next_6_months","next_12_months","year_to_date","last_3_days",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDashboardTypeDateIntervalPropEnum = append(updateDashboardTypeDateIntervalPropEnum, v)
	}
}

const (

	// UpdateDashboardDateIntervalThisMonth captures enum value "this_month"
	UpdateDashboardDateIntervalThisMonth string = "this_month"

	// UpdateDashboardDateIntervalLast7Days captures enum value "last_7_days"
	UpdateDashboardDateIntervalLast7Days string = "last_7_days"

	// UpdateDashboardDateIntervalLast30Days captures enum value "last_30_days"
	UpdateDashboardDateIntervalLast30Days string = "last_30_days"

	// UpdateDashboardDateIntervalLastMonth captures enum value "last_month"
	UpdateDashboardDateIntervalLastMonth string = "last_month"

	// UpdateDashboardDateIntervalLast3Months captures enum value "last_3_months"
	UpdateDashboardDateIntervalLast3Months string = "last_3_months"

	// UpdateDashboardDateIntervalLast6Months captures enum value "last_6_months"
	UpdateDashboardDateIntervalLast6Months string = "last_6_months"

	// UpdateDashboardDateIntervalCustom captures enum value "custom"
	UpdateDashboardDateIntervalCustom string = "custom"

	// UpdateDashboardDateIntervalLast12Months captures enum value "last_12_months"
	UpdateDashboardDateIntervalLast12Months string = "last_12_months"

	// UpdateDashboardDateIntervalLast24Months captures enum value "last_24_months"
	UpdateDashboardDateIntervalLast24Months string = "last_24_months"

	// UpdateDashboardDateIntervalLast36Months captures enum value "last_36_months"
	UpdateDashboardDateIntervalLast36Months string = "last_36_months"

	// UpdateDashboardDateIntervalNextMonth captures enum value "next_month"
	UpdateDashboardDateIntervalNextMonth string = "next_month"

	// UpdateDashboardDateIntervalNext3Months captures enum value "next_3_months"
	UpdateDashboardDateIntervalNext3Months string = "next_3_months"

	// UpdateDashboardDateIntervalNext6Months captures enum value "next_6_months"
	UpdateDashboardDateIntervalNext6Months string = "next_6_months"

	// UpdateDashboardDateIntervalNext12Months captures enum value "next_12_months"
	UpdateDashboardDateIntervalNext12Months string = "next_12_months"

	// UpdateDashboardDateIntervalYearToDate captures enum value "year_to_date"
	UpdateDashboardDateIntervalYearToDate string = "year_to_date"

	// UpdateDashboardDateIntervalLast3Days captures enum value "last_3_days"
	UpdateDashboardDateIntervalLast3Days string = "last_3_days"

	// UpdateDashboardDateIntervalEmpty captures enum value ""
	UpdateDashboardDateIntervalEmpty string = ""
)

// prop value enum
func (m *UpdateDashboard) validateDateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateDashboardTypeDateIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateDashboard) validateDateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.DateInterval) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateIntervalEnum("date_interval", "body", m.DateInterval); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDashboard) validateWidgets(formats strfmt.Registry) error {
	if swag.IsZero(m.Widgets) { // not required
		return nil
	}

	for i := 0; i < len(m.Widgets); i++ {
		if swag.IsZero(m.Widgets[i]) { // not required
			continue
		}

		if m.Widgets[i] != nil {
			if err := m.Widgets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("widgets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("widgets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update dashboard based on the context it is used
func (m *UpdateDashboard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWidgets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDashboard) contextValidateWidgets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Widgets); i++ {

		if m.Widgets[i] != nil {

			if swag.IsZero(m.Widgets[i]) { // not required
				return nil
			}

			if err := m.Widgets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("widgets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("widgets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDashboard) UnmarshalBinary(b []byte) error {
	var res UpdateDashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdateDashboardWidgetsItems0 update dashboard widgets items0
//
// swagger:model UpdateDashboardWidgetsItems0
type UpdateDashboardWidgetsItems0 struct {

	// settings
	Settings *UpdateDashboardWidgetsItems0Settings `json:"settings,omitempty"`

	// The title of the Widget (defaults to the title of the Resource).
	Title string `json:"title,omitempty"`

	// The token of the represented Resource.
	// Required: true
	WidgetableToken *string `json:"widgetable_token"`
}

// Validate validates this update dashboard widgets items0
func (m *UpdateDashboardWidgetsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidgetableToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDashboardWidgetsItems0) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateDashboardWidgetsItems0) validateWidgetableToken(formats strfmt.Registry) error {

	if err := validate.Required("widgetable_token", "body", m.WidgetableToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update dashboard widgets items0 based on the context it is used
func (m *UpdateDashboardWidgetsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDashboardWidgetsItems0) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {

		if swag.IsZero(m.Settings) { // not required
			return nil
		}

		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDashboardWidgetsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDashboardWidgetsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateDashboardWidgetsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdateDashboardWidgetsItems0Settings The settings for the DashboardWidget.
//
// swagger:model UpdateDashboardWidgetsItems0Settings
type UpdateDashboardWidgetsItems0Settings struct {

	// display type
	// Required: true
	// Enum: ["table","chart"]
	DisplayType *string `json:"display_type"`
}

// Validate validates this update dashboard widgets items0 settings
func (m *UpdateDashboardWidgetsItems0Settings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateDashboardWidgetsItems0SettingsTypeDisplayTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["table","chart"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDashboardWidgetsItems0SettingsTypeDisplayTypePropEnum = append(updateDashboardWidgetsItems0SettingsTypeDisplayTypePropEnum, v)
	}
}

const (

	// UpdateDashboardWidgetsItems0SettingsDisplayTypeTable captures enum value "table"
	UpdateDashboardWidgetsItems0SettingsDisplayTypeTable string = "table"

	// UpdateDashboardWidgetsItems0SettingsDisplayTypeChart captures enum value "chart"
	UpdateDashboardWidgetsItems0SettingsDisplayTypeChart string = "chart"
)

// prop value enum
func (m *UpdateDashboardWidgetsItems0Settings) validateDisplayTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateDashboardWidgetsItems0SettingsTypeDisplayTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateDashboardWidgetsItems0Settings) validateDisplayType(formats strfmt.Registry) error {

	if err := validate.Required("settings"+"."+"display_type", "body", m.DisplayType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDisplayTypeEnum("settings"+"."+"display_type", "body", *m.DisplayType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update dashboard widgets items0 settings based on context it is used
func (m *UpdateDashboardWidgetsItems0Settings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDashboardWidgetsItems0Settings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDashboardWidgetsItems0Settings) UnmarshalBinary(b []byte) error {
	var res UpdateDashboardWidgetsItems0Settings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
