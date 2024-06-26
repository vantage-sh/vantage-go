// Code generated by go-swagger; DO NOT EDIT.

package budget_alerts

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

// NewGetBudgetAlertsParams creates a new GetBudgetAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBudgetAlertsParams() *GetBudgetAlertsParams {
	return &GetBudgetAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBudgetAlertsParamsWithTimeout creates a new GetBudgetAlertsParams object
// with the ability to set a timeout on a request.
func NewGetBudgetAlertsParamsWithTimeout(timeout time.Duration) *GetBudgetAlertsParams {
	return &GetBudgetAlertsParams{
		timeout: timeout,
	}
}

// NewGetBudgetAlertsParamsWithContext creates a new GetBudgetAlertsParams object
// with the ability to set a context for a request.
func NewGetBudgetAlertsParamsWithContext(ctx context.Context) *GetBudgetAlertsParams {
	return &GetBudgetAlertsParams{
		Context: ctx,
	}
}

// NewGetBudgetAlertsParamsWithHTTPClient creates a new GetBudgetAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBudgetAlertsParamsWithHTTPClient(client *http.Client) *GetBudgetAlertsParams {
	return &GetBudgetAlertsParams{
		HTTPClient: client,
	}
}

/*
GetBudgetAlertsParams contains all the parameters to send to the API endpoint

	for the get budget alerts operation.

	Typically these are written to a http.Request.
*/
type GetBudgetAlertsParams struct {

	/* Limit.

	   The number of results to return. The maximum is 1000.

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

// WithDefaults hydrates default values in the get budget alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBudgetAlertsParams) WithDefaults() *GetBudgetAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get budget alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBudgetAlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get budget alerts params
func (o *GetBudgetAlertsParams) WithTimeout(timeout time.Duration) *GetBudgetAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get budget alerts params
func (o *GetBudgetAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get budget alerts params
func (o *GetBudgetAlertsParams) WithContext(ctx context.Context) *GetBudgetAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get budget alerts params
func (o *GetBudgetAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get budget alerts params
func (o *GetBudgetAlertsParams) WithHTTPClient(client *http.Client) *GetBudgetAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get budget alerts params
func (o *GetBudgetAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get budget alerts params
func (o *GetBudgetAlertsParams) WithLimit(limit *int32) *GetBudgetAlertsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get budget alerts params
func (o *GetBudgetAlertsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get budget alerts params
func (o *GetBudgetAlertsParams) WithPage(page *int32) *GetBudgetAlertsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get budget alerts params
func (o *GetBudgetAlertsParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetBudgetAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
