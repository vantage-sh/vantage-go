// Code generated by go-swagger; DO NOT EDIT.

package costs

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

// NewGetCostReportParams creates a new GetCostReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCostReportParams() *GetCostReportParams {
	return &GetCostReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCostReportParamsWithTimeout creates a new GetCostReportParams object
// with the ability to set a timeout on a request.
func NewGetCostReportParamsWithTimeout(timeout time.Duration) *GetCostReportParams {
	return &GetCostReportParams{
		timeout: timeout,
	}
}

// NewGetCostReportParamsWithContext creates a new GetCostReportParams object
// with the ability to set a context for a request.
func NewGetCostReportParamsWithContext(ctx context.Context) *GetCostReportParams {
	return &GetCostReportParams{
		Context: ctx,
	}
}

// NewGetCostReportParamsWithHTTPClient creates a new GetCostReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCostReportParamsWithHTTPClient(client *http.Client) *GetCostReportParams {
	return &GetCostReportParams{
		HTTPClient: client,
	}
}

/*
GetCostReportParams contains all the parameters to send to the API endpoint

	for the get cost report operation.

	Typically these are written to a http.Request.
*/
type GetCostReportParams struct {

	// CostReportToken.
	CostReportToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cost report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostReportParams) WithDefaults() *GetCostReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cost report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cost report params
func (o *GetCostReportParams) WithTimeout(timeout time.Duration) *GetCostReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cost report params
func (o *GetCostReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cost report params
func (o *GetCostReportParams) WithContext(ctx context.Context) *GetCostReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cost report params
func (o *GetCostReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cost report params
func (o *GetCostReportParams) WithHTTPClient(client *http.Client) *GetCostReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cost report params
func (o *GetCostReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCostReportToken adds the costReportToken to the get cost report params
func (o *GetCostReportParams) WithCostReportToken(costReportToken string) *GetCostReportParams {
	o.SetCostReportToken(costReportToken)
	return o
}

// SetCostReportToken adds the costReportToken to the get cost report params
func (o *GetCostReportParams) SetCostReportToken(costReportToken string) {
	o.CostReportToken = costReportToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetCostReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cost_report_token
	if err := r.SetPathParam("cost_report_token", o.CostReportToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}