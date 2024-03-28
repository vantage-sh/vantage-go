// Code generated by go-swagger; DO NOT EDIT.

package business_metrics

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
)

// NewUpdateBusinessMetricValuesCSVParams creates a new UpdateBusinessMetricValuesCSVParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBusinessMetricValuesCSVParams() *UpdateBusinessMetricValuesCSVParams {
	return &UpdateBusinessMetricValuesCSVParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBusinessMetricValuesCSVParamsWithTimeout creates a new UpdateBusinessMetricValuesCSVParams object
// with the ability to set a timeout on a request.
func NewUpdateBusinessMetricValuesCSVParamsWithTimeout(timeout time.Duration) *UpdateBusinessMetricValuesCSVParams {
	return &UpdateBusinessMetricValuesCSVParams{
		timeout: timeout,
	}
}

// NewUpdateBusinessMetricValuesCSVParamsWithContext creates a new UpdateBusinessMetricValuesCSVParams object
// with the ability to set a context for a request.
func NewUpdateBusinessMetricValuesCSVParamsWithContext(ctx context.Context) *UpdateBusinessMetricValuesCSVParams {
	return &UpdateBusinessMetricValuesCSVParams{
		Context: ctx,
	}
}

// NewUpdateBusinessMetricValuesCSVParamsWithHTTPClient creates a new UpdateBusinessMetricValuesCSVParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBusinessMetricValuesCSVParamsWithHTTPClient(client *http.Client) *UpdateBusinessMetricValuesCSVParams {
	return &UpdateBusinessMetricValuesCSVParams{
		HTTPClient: client,
	}
}

/*
UpdateBusinessMetricValuesCSVParams contains all the parameters to send to the API endpoint

	for the update business metric values c s v operation.

	Typically these are written to a http.Request.
*/
type UpdateBusinessMetricValuesCSVParams struct {

	// BusinessMetricToken.
	BusinessMetricToken string

	/* Csv.

	   CSV file containing BusinessMetric dates and amounts
	*/
	Csv runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update business metric values c s v params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBusinessMetricValuesCSVParams) WithDefaults() *UpdateBusinessMetricValuesCSVParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update business metric values c s v params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBusinessMetricValuesCSVParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) WithTimeout(timeout time.Duration) *UpdateBusinessMetricValuesCSVParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) WithContext(ctx context.Context) *UpdateBusinessMetricValuesCSVParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) WithHTTPClient(client *http.Client) *UpdateBusinessMetricValuesCSVParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessMetricToken adds the businessMetricToken to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) WithBusinessMetricToken(businessMetricToken string) *UpdateBusinessMetricValuesCSVParams {
	o.SetBusinessMetricToken(businessMetricToken)
	return o
}

// SetBusinessMetricToken adds the businessMetricToken to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) SetBusinessMetricToken(businessMetricToken string) {
	o.BusinessMetricToken = businessMetricToken
}

// WithCsv adds the csv to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) WithCsv(csv runtime.NamedReadCloser) *UpdateBusinessMetricValuesCSVParams {
	o.SetCsv(csv)
	return o
}

// SetCsv adds the csv to the update business metric values c s v params
func (o *UpdateBusinessMetricValuesCSVParams) SetCsv(csv runtime.NamedReadCloser) {
	o.Csv = csv
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBusinessMetricValuesCSVParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param business_metric_token
	if err := r.SetPathParam("business_metric_token", o.BusinessMetricToken); err != nil {
		return err
	}
	// form file param csv
	if err := r.SetFileParam("csv", o.Csv); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
