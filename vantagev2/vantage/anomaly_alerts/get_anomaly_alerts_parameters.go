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
	"github.com/go-openapi/swag"
)

// NewGetAnomalyAlertsParams creates a new GetAnomalyAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAnomalyAlertsParams() *GetAnomalyAlertsParams {
	return &GetAnomalyAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnomalyAlertsParamsWithTimeout creates a new GetAnomalyAlertsParams object
// with the ability to set a timeout on a request.
func NewGetAnomalyAlertsParamsWithTimeout(timeout time.Duration) *GetAnomalyAlertsParams {
	return &GetAnomalyAlertsParams{
		timeout: timeout,
	}
}

// NewGetAnomalyAlertsParamsWithContext creates a new GetAnomalyAlertsParams object
// with the ability to set a context for a request.
func NewGetAnomalyAlertsParamsWithContext(ctx context.Context) *GetAnomalyAlertsParams {
	return &GetAnomalyAlertsParams{
		Context: ctx,
	}
}

// NewGetAnomalyAlertsParamsWithHTTPClient creates a new GetAnomalyAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAnomalyAlertsParamsWithHTTPClient(client *http.Client) *GetAnomalyAlertsParams {
	return &GetAnomalyAlertsParams{
		HTTPClient: client,
	}
}

/*
GetAnomalyAlertsParams contains all the parameters to send to the API endpoint

	for the get anomaly alerts operation.

	Typically these are written to a http.Request.
*/
type GetAnomalyAlertsParams struct {

	/* Limit.

	   The amount of results to return. The maximum is 1000.

	   Format: int32
	*/
	Limit *int32

	/* Page.

	   The page of results to return.

	   Format: int32
	*/
	Page *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get anomaly alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnomalyAlertsParams) WithDefaults() *GetAnomalyAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get anomaly alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnomalyAlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) WithTimeout(timeout time.Duration) *GetAnomalyAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) WithContext(ctx context.Context) *GetAnomalyAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) WithHTTPClient(client *http.Client) *GetAnomalyAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) WithLimit(limit *int32) *GetAnomalyAlertsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) WithPage(page *int32) *GetAnomalyAlertsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get anomaly alerts params
func (o *GetAnomalyAlertsParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnomalyAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
