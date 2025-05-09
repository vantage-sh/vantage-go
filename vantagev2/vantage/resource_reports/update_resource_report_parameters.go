// Code generated by go-swagger; DO NOT EDIT.

package resource_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// NewUpdateResourceReportParams creates a new UpdateResourceReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateResourceReportParams() *UpdateResourceReportParams {
	return &UpdateResourceReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateResourceReportParamsWithTimeout creates a new UpdateResourceReportParams object
// with the ability to set a timeout on a request.
func NewUpdateResourceReportParamsWithTimeout(timeout time.Duration) *UpdateResourceReportParams {
	return &UpdateResourceReportParams{
		timeout: timeout,
	}
}

// NewUpdateResourceReportParamsWithContext creates a new UpdateResourceReportParams object
// with the ability to set a context for a request.
func NewUpdateResourceReportParamsWithContext(ctx context.Context) *UpdateResourceReportParams {
	return &UpdateResourceReportParams{
		Context: ctx,
	}
}

// NewUpdateResourceReportParamsWithHTTPClient creates a new UpdateResourceReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateResourceReportParamsWithHTTPClient(client *http.Client) *UpdateResourceReportParams {
	return &UpdateResourceReportParams{
		HTTPClient: client,
	}
}

/*
UpdateResourceReportParams contains all the parameters to send to the API endpoint

	for the update resource report operation.

	Typically these are written to a http.Request.
*/
type UpdateResourceReportParams struct {

	// ResourceReportToken.
	ResourceReportToken string

	// UpdateResourceReport.
	UpdateResourceReport *models.UpdateResourceReport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update resource report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateResourceReportParams) WithDefaults() *UpdateResourceReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update resource report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateResourceReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update resource report params
func (o *UpdateResourceReportParams) WithTimeout(timeout time.Duration) *UpdateResourceReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update resource report params
func (o *UpdateResourceReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update resource report params
func (o *UpdateResourceReportParams) WithContext(ctx context.Context) *UpdateResourceReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update resource report params
func (o *UpdateResourceReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update resource report params
func (o *UpdateResourceReportParams) WithHTTPClient(client *http.Client) *UpdateResourceReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update resource report params
func (o *UpdateResourceReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceReportToken adds the resourceReportToken to the update resource report params
func (o *UpdateResourceReportParams) WithResourceReportToken(resourceReportToken string) *UpdateResourceReportParams {
	o.SetResourceReportToken(resourceReportToken)
	return o
}

// SetResourceReportToken adds the resourceReportToken to the update resource report params
func (o *UpdateResourceReportParams) SetResourceReportToken(resourceReportToken string) {
	o.ResourceReportToken = resourceReportToken
}

// WithUpdateResourceReport adds the updateResourceReport to the update resource report params
func (o *UpdateResourceReportParams) WithUpdateResourceReport(updateResourceReport *models.UpdateResourceReport) *UpdateResourceReportParams {
	o.SetUpdateResourceReport(updateResourceReport)
	return o
}

// SetUpdateResourceReport adds the updateResourceReport to the update resource report params
func (o *UpdateResourceReportParams) SetUpdateResourceReport(updateResourceReport *models.UpdateResourceReport) {
	o.UpdateResourceReport = updateResourceReport
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateResourceReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_report_token
	if err := r.SetPathParam("resource_report_token", o.ResourceReportToken); err != nil {
		return err
	}
	if o.UpdateResourceReport != nil {
		if err := r.SetBodyParam(o.UpdateResourceReport); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
