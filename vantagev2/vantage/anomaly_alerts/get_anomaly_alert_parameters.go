// Code generated by go-swagger; DO NOT EDIT.

package anomaly_alerts

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

// NewGetAnomalyAlertParams creates a new GetAnomalyAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAnomalyAlertParams() *GetAnomalyAlertParams {
	return &GetAnomalyAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnomalyAlertParamsWithTimeout creates a new GetAnomalyAlertParams object
// with the ability to set a timeout on a request.
func NewGetAnomalyAlertParamsWithTimeout(timeout time.Duration) *GetAnomalyAlertParams {
	return &GetAnomalyAlertParams{
		timeout: timeout,
	}
}

// NewGetAnomalyAlertParamsWithContext creates a new GetAnomalyAlertParams object
// with the ability to set a context for a request.
func NewGetAnomalyAlertParamsWithContext(ctx context.Context) *GetAnomalyAlertParams {
	return &GetAnomalyAlertParams{
		Context: ctx,
	}
}

// NewGetAnomalyAlertParamsWithHTTPClient creates a new GetAnomalyAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAnomalyAlertParamsWithHTTPClient(client *http.Client) *GetAnomalyAlertParams {
	return &GetAnomalyAlertParams{
		HTTPClient: client,
	}
}

/*
GetAnomalyAlertParams contains all the parameters to send to the API endpoint

	for the get anomaly alert operation.

	Typically these are written to a http.Request.
*/
type GetAnomalyAlertParams struct {

	// AnomalyAlertToken.
	AnomalyAlertToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get anomaly alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnomalyAlertParams) WithDefaults() *GetAnomalyAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get anomaly alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnomalyAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get anomaly alert params
func (o *GetAnomalyAlertParams) WithTimeout(timeout time.Duration) *GetAnomalyAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get anomaly alert params
func (o *GetAnomalyAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get anomaly alert params
func (o *GetAnomalyAlertParams) WithContext(ctx context.Context) *GetAnomalyAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get anomaly alert params
func (o *GetAnomalyAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get anomaly alert params
func (o *GetAnomalyAlertParams) WithHTTPClient(client *http.Client) *GetAnomalyAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get anomaly alert params
func (o *GetAnomalyAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnomalyAlertToken adds the anomalyAlertToken to the get anomaly alert params
func (o *GetAnomalyAlertParams) WithAnomalyAlertToken(anomalyAlertToken string) *GetAnomalyAlertParams {
	o.SetAnomalyAlertToken(anomalyAlertToken)
	return o
}

// SetAnomalyAlertToken adds the anomalyAlertToken to the get anomaly alert params
func (o *GetAnomalyAlertParams) SetAnomalyAlertToken(anomalyAlertToken string) {
	o.AnomalyAlertToken = anomalyAlertToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnomalyAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param anomaly_alert_token
	if err := r.SetPathParam("anomaly_alert_token", o.AnomalyAlertToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
