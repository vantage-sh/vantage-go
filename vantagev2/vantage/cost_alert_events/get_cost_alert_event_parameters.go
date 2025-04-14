// Code generated by go-swagger; DO NOT EDIT.

package cost_alert_events

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

// NewGetCostAlertEventParams creates a new GetCostAlertEventParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCostAlertEventParams() *GetCostAlertEventParams {
	return &GetCostAlertEventParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCostAlertEventParamsWithTimeout creates a new GetCostAlertEventParams object
// with the ability to set a timeout on a request.
func NewGetCostAlertEventParamsWithTimeout(timeout time.Duration) *GetCostAlertEventParams {
	return &GetCostAlertEventParams{
		timeout: timeout,
	}
}

// NewGetCostAlertEventParamsWithContext creates a new GetCostAlertEventParams object
// with the ability to set a context for a request.
func NewGetCostAlertEventParamsWithContext(ctx context.Context) *GetCostAlertEventParams {
	return &GetCostAlertEventParams{
		Context: ctx,
	}
}

// NewGetCostAlertEventParamsWithHTTPClient creates a new GetCostAlertEventParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCostAlertEventParamsWithHTTPClient(client *http.Client) *GetCostAlertEventParams {
	return &GetCostAlertEventParams{
		HTTPClient: client,
	}
}

/*
GetCostAlertEventParams contains all the parameters to send to the API endpoint

	for the get cost alert event operation.

	Typically these are written to a http.Request.
*/
type GetCostAlertEventParams struct {

	// CostAlertToken.
	CostAlertToken string

	// EventToken.
	EventToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cost alert event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostAlertEventParams) WithDefaults() *GetCostAlertEventParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cost alert event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostAlertEventParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cost alert event params
func (o *GetCostAlertEventParams) WithTimeout(timeout time.Duration) *GetCostAlertEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cost alert event params
func (o *GetCostAlertEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cost alert event params
func (o *GetCostAlertEventParams) WithContext(ctx context.Context) *GetCostAlertEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cost alert event params
func (o *GetCostAlertEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cost alert event params
func (o *GetCostAlertEventParams) WithHTTPClient(client *http.Client) *GetCostAlertEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cost alert event params
func (o *GetCostAlertEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCostAlertToken adds the costAlertToken to the get cost alert event params
func (o *GetCostAlertEventParams) WithCostAlertToken(costAlertToken string) *GetCostAlertEventParams {
	o.SetCostAlertToken(costAlertToken)
	return o
}

// SetCostAlertToken adds the costAlertToken to the get cost alert event params
func (o *GetCostAlertEventParams) SetCostAlertToken(costAlertToken string) {
	o.CostAlertToken = costAlertToken
}

// WithEventToken adds the eventToken to the get cost alert event params
func (o *GetCostAlertEventParams) WithEventToken(eventToken string) *GetCostAlertEventParams {
	o.SetEventToken(eventToken)
	return o
}

// SetEventToken adds the eventToken to the get cost alert event params
func (o *GetCostAlertEventParams) SetEventToken(eventToken string) {
	o.EventToken = eventToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetCostAlertEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cost_alert_token
	if err := r.SetPathParam("cost_alert_token", o.CostAlertToken); err != nil {
		return err
	}

	// path param event_token
	if err := r.SetPathParam("event_token", o.EventToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
