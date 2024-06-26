// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateGCPIntegration Create a GCP Integration
//
// swagger:model createGCPIntegration
type CreateGCPIntegration struct {

	// GCP billing account ID.
	// Required: true
	BillingAccountID *string `json:"billing_account_id"`

	// BigQuery dataset name.
	// Required: true
	DatasetName *string `json:"dataset_name"`

	// GCP project ID.
	// Required: true
	ProjectID *string `json:"project_id"`
}

// Validate validates this create g c p integration
func (m *CreateGCPIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatasetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGCPIntegration) validateBillingAccountID(formats strfmt.Registry) error {

	if err := validate.Required("billing_account_id", "body", m.BillingAccountID); err != nil {
		return err
	}

	return nil
}

func (m *CreateGCPIntegration) validateDatasetName(formats strfmt.Registry) error {

	if err := validate.Required("dataset_name", "body", m.DatasetName); err != nil {
		return err
	}

	return nil
}

func (m *CreateGCPIntegration) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("project_id", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create g c p integration based on context it is used
func (m *CreateGCPIntegration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateGCPIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGCPIntegration) UnmarshalBinary(b []byte) error {
	var res CreateGCPIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
