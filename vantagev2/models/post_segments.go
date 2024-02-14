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

// PostSegments Create a Segment.
//
// swagger:model postSegments
type PostSegments struct {

	// The description of the Segment.
	Description string `json:"description,omitempty"`

	// The filter query language to apply to the Segment. Additional documentation available at https://docs.vantage.sh/vql.
	Filter string `json:"filter,omitempty"`

	// The token of the parent Segment this new Segment belongs to. Determines the Workspace the segment is assigned to.
	ParentSegmentToken string `json:"parent_segment_token,omitempty"`

	// The priority of the Segment.
	Priority int32 `json:"priority,omitempty"`

	// report settings
	ReportSettings *PostSegmentsReportSettings `json:"report_settings,omitempty"`

	// The title of the Segment.
	// Required: true
	Title *string `json:"title"`

	// Track Unallocated Costs which are not assigned to any of the created Segments.
	TrackUnallocated *bool `json:"track_unallocated,omitempty"`

	// The token of the Workspace to add the Segment to. Ignored if 'segment_token' is set. Required if the API token is associated with multiple Workspaces.
	WorkspaceToken string `json:"workspace_token,omitempty"`
}

// Validate validates this post segments
func (m *PostSegments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostSegments) validateReportSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.ReportSettings) { // not required
		return nil
	}

	if m.ReportSettings != nil {
		if err := m.ReportSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("report_settings")
			}
			return err
		}
	}

	return nil
}

func (m *PostSegments) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post segments based on the context it is used
func (m *PostSegments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReportSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostSegments) contextValidateReportSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.ReportSettings != nil {

		if swag.IsZero(m.ReportSettings) { // not required
			return nil
		}

		if err := m.ReportSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("report_settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostSegments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSegments) UnmarshalBinary(b []byte) error {
	var res PostSegments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostSegmentsReportSettings Report settings configurable on top-level Segments.
//
// swagger:model PostSegmentsReportSettings
type PostSegmentsReportSettings struct {

	// Reports created under this Segment will amortize.
	Amortize bool `json:"amortize,omitempty"`

	// Reports created under this Segment will include credits.
	IncludeCredits bool `json:"include_credits,omitempty"`

	// Reports created under this Segment will include discounts.
	IncludeDiscounts bool `json:"include_discounts,omitempty"`

	// Reports created under this Segment will include refunds.
	IncludeRefunds bool `json:"include_refunds,omitempty"`

	// Reports created under this Segment will include tax.
	IncludeTax bool `json:"include_tax,omitempty"`
}

// Validate validates this post segments report settings
func (m *PostSegmentsReportSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post segments report settings based on context it is used
func (m *PostSegmentsReportSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostSegmentsReportSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSegmentsReportSettings) UnmarshalBinary(b []byte) error {
	var res PostSegmentsReportSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
