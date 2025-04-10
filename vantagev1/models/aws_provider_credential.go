// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsProviderCredential AwsProviderCredential model
//
// swagger:model AwsProviderCredential
type AwsProviderCredential struct {

	// external id
	ExternalID string `json:"external_id,omitempty"`

	// iam role arn
	IamRoleArn string `json:"iam_role_arn,omitempty"`

	// iam role id
	IamRoleID string `json:"iam_role_id,omitempty"`

	// policies
	Policies *AwsIamPolicies `json:"policies,omitempty"`

	// vantage id
	VantageID string `json:"vantage_id,omitempty"`
}

// Validate validates this aws provider credential
func (m *AwsProviderCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsProviderCredential) validatePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	if m.Policies != nil {
		if err := m.Policies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policies")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this aws provider credential based on the context it is used
func (m *AwsProviderCredential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsProviderCredential) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	if m.Policies != nil {

		if swag.IsZero(m.Policies) { // not required
			return nil
		}

		if err := m.Policies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policies")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsProviderCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsProviderCredential) UnmarshalBinary(b []byte) error {
	var res AwsProviderCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
