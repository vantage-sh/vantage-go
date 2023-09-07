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

// NewDeleteCostReportParams creates a new DeleteCostReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCostReportParams() *DeleteCostReportParams {
	return &DeleteCostReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCostReportParamsWithTimeout creates a new DeleteCostReportParams object
// with the ability to set a timeout on a request.
func NewDeleteCostReportParamsWithTimeout(timeout time.Duration) *DeleteCostReportParams {
	return &DeleteCostReportParams{
		timeout: timeout,
	}
}

// NewDeleteCostReportParamsWithContext creates a new DeleteCostReportParams object
// with the ability to set a context for a request.
func NewDeleteCostReportParamsWithContext(ctx context.Context) *DeleteCostReportParams {
	return &DeleteCostReportParams{
		Context: ctx,
	}
}

// NewDeleteCostReportParamsWithHTTPClient creates a new DeleteCostReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCostReportParamsWithHTTPClient(client *http.Client) *DeleteCostReportParams {
	return &DeleteCostReportParams{
		HTTPClient: client,
	}
}

/*
DeleteCostReportParams contains all the parameters to send to the API endpoint

	for the delete cost report operation.

	Typically these are written to a http.Request.
*/
type DeleteCostReportParams struct {

	// CostReportToken.
	CostReportToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cost report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCostReportParams) WithDefaults() *DeleteCostReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cost report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCostReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cost report params
func (o *DeleteCostReportParams) WithTimeout(timeout time.Duration) *DeleteCostReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cost report params
func (o *DeleteCostReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cost report params
func (o *DeleteCostReportParams) WithContext(ctx context.Context) *DeleteCostReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cost report params
func (o *DeleteCostReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cost report params
func (o *DeleteCostReportParams) WithHTTPClient(client *http.Client) *DeleteCostReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cost report params
func (o *DeleteCostReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCostReportToken adds the costReportToken to the delete cost report params
func (o *DeleteCostReportParams) WithCostReportToken(costReportToken string) *DeleteCostReportParams {
	o.SetCostReportToken(costReportToken)
	return o
}

// SetCostReportToken adds the costReportToken to the delete cost report params
func (o *DeleteCostReportParams) SetCostReportToken(costReportToken string) {
	o.CostReportToken = costReportToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCostReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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