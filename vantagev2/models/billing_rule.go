// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BillingRule BillingRule model
//
// swagger:model BillingRule
type BillingRule struct {

	// The amount for the BillingRule (Charge).
	// Example: 5000.25
	Amount string `json:"amount,omitempty"`

	// Whether the BillingRule applies to all future managed accounts.
	// Example: true
	ApplyToAll bool `json:"apply_to_all,omitempty"`

	// The category for the BillingRule (Charge).
	// Example: MSP Fee
	Category string `json:"category,omitempty"`

	// The charge type for the BillingRule.
	// Example: RIFee
	ChargeType string `json:"charge_type,omitempty"`

	// The date and time, in UTC, the BillingRule was created. ISO 8601 Formatted.
	// Example: 2024-06-28T00:00:00Z
	CreatedAt string `json:"created_at,omitempty"`

	// The token of the Creator of the BillingRule.
	// Example: usr_1234
	CreatedByToken string `json:"created_by_token,omitempty"`

	// The end date of the BillingRule.
	// Example: 2024-06-28T00:00:00Z
	EndDate string `json:"end_date,omitempty"`

	// The percentage of the cost shown for the BillingRule (Adjustment).
	// Example: 75.0
	Percentage string `json:"percentage,omitempty"`

	// The service for the BillingRule (Charge).
	// Example: AWS Cloudfront
	Service string `json:"service,omitempty"`

	// The SQL query for the BillingRule (Custom).
	// Example: UPDATE costs SET costs.amount = costs.amount * 0.95
	SQLQuery string `json:"sql_query,omitempty"`

	// The start date of the BillingRule.
	// Example: 2024-06-28T00:00:00Z
	StartDate string `json:"start_date,omitempty"`

	// The start period for the BillingRule (Charge).
	// Example: 2024-05-01
	StartPeriod string `json:"start_period,omitempty"`

	// The subcategory for the BillingRule (Charge).
	// Example: One-time
	SubCategory string `json:"sub_category,omitempty"`

	// The title of the BillingRule.
	// Example: Credit for Unused EC2 Instances
	Title string `json:"title,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// The type of the BillingRule.
	// Example: credit
	Type string `json:"type,omitempty"`
}

// Validate validates this billing rule
func (m *BillingRule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this billing rule based on context it is used
func (m *BillingRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BillingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingRule) UnmarshalBinary(b []byte) error {
	var res BillingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
