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

// Segments Segments model
//
// swagger:model Segments
type Segments struct {

	// links
	Links interface{} `json:"links,omitempty"`

	// segments
	Segments []*Segment `json:"segments"`
}

// Validate validates this segments
func (m *Segments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSegments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Segments) validateSegments(formats strfmt.Registry) error {
	if swag.IsZero(m.Segments) { // not required
		return nil
	}

	for i := 0; i < len(m.Segments); i++ {
		if swag.IsZero(m.Segments[i]) { // not required
			continue
		}

		if m.Segments[i] != nil {
			if err := m.Segments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this segments based on the context it is used
func (m *Segments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSegments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Segments) contextValidateSegments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Segments); i++ {

		if m.Segments[i] != nil {

			if swag.IsZero(m.Segments[i]) { // not required
				return nil
			}

			if err := m.Segments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Segments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Segments) UnmarshalBinary(b []byte) error {
	var res Segments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
