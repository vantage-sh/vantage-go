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

// CreateBillingRule Create a BillingRule.
//
// swagger:model createBillingRule
type CreateBillingRule struct {

	// The amount for the BillingRule. Example value: 300
	// Required: true
	Amount *float64 `json:"amount"`

	// Determines if the BillingRule applies to all current and future managed accounts.
	ApplyToAll bool `json:"apply_to_all,omitempty"`

	// The category of the BillingRule.
	// Required: true
	Category *string `json:"category"`

	// The charge type of the BillingRule.
	// Required: true
	ChargeType *string `json:"charge_type"`

	// The end date of the BillingRule. ISO 8601 formatted.
	EndDate string `json:"end_date,omitempty"`

	// The percentage of the cost shown. Example value: 75.0
	// Required: true
	Percentage *float64 `json:"percentage"`

	// The service of the BillingRule.
	// Required: true
	Service *string `json:"service"`

	// The start date of the BillingRule. ISO 8601 formatted.
	StartDate string `json:"start_date,omitempty"`

	// The start period of the BillingRule.
	// Required: true
	StartPeriod *string `json:"start_period"`

	// The subcategory of the BillingRule.
	// Required: true
	SubCategory *string `json:"sub_category"`

	// The title of the BillingRule.
	// Required: true
	Title *string `json:"title"`

	// The type of the BillingRule. Note: the values are case insensitive.
	// Required: true
	// Enum: ["exclusion","adjustment","credit","charge"]
	Type *string `json:"type"`
}

// Validate validates this create billing rule
func (m *CreateBillingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBillingRule) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *CreateBillingRule) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *CreateBillingRule) validateChargeType(formats strfmt.Registry) error {

	if err := validate.Required("charge_type", "body", m.ChargeType); err != nil {
		return err
	}

	return nil
}

func (m *CreateBillingRule) validatePercentage(formats strfmt.Registry) error {

	if err := validate.Required("percentage", "body", m.Percentage); err != nil {
		return err
	}

	return nil
}

func (m *CreateBillingRule) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	return nil
}

func (m *CreateBillingRule) validateStartPeriod(formats strfmt.Registry) error {

	if err := validate.Required("start_period", "body", m.StartPeriod); err != nil {
		return err
	}

	return nil
}

func (m *CreateBillingRule) validateSubCategory(formats strfmt.Registry) error {

	if err := validate.Required("sub_category", "body", m.SubCategory); err != nil {
		return err
	}

	return nil
}

func (m *CreateBillingRule) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

var createBillingRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["exclusion","adjustment","credit","charge"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createBillingRuleTypeTypePropEnum = append(createBillingRuleTypeTypePropEnum, v)
	}
}

const (

	// CreateBillingRuleTypeExclusion captures enum value "exclusion"
	CreateBillingRuleTypeExclusion string = "exclusion"

	// CreateBillingRuleTypeAdjustment captures enum value "adjustment"
	CreateBillingRuleTypeAdjustment string = "adjustment"

	// CreateBillingRuleTypeCredit captures enum value "credit"
	CreateBillingRuleTypeCredit string = "credit"

	// CreateBillingRuleTypeCharge captures enum value "charge"
	CreateBillingRuleTypeCharge string = "charge"
)

// prop value enum
func (m *CreateBillingRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createBillingRuleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateBillingRule) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create billing rule based on context it is used
func (m *CreateBillingRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateBillingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBillingRule) UnmarshalBinary(b []byte) error {
	var res CreateBillingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
