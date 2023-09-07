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
	"github.com/go-openapi/swag"
)

// NewGetCostReportFoldersParams creates a new GetCostReportFoldersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCostReportFoldersParams() *GetCostReportFoldersParams {
	return &GetCostReportFoldersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCostReportFoldersParamsWithTimeout creates a new GetCostReportFoldersParams object
// with the ability to set a timeout on a request.
func NewGetCostReportFoldersParamsWithTimeout(timeout time.Duration) *GetCostReportFoldersParams {
	return &GetCostReportFoldersParams{
		timeout: timeout,
	}
}

// NewGetCostReportFoldersParamsWithContext creates a new GetCostReportFoldersParams object
// with the ability to set a context for a request.
func NewGetCostReportFoldersParamsWithContext(ctx context.Context) *GetCostReportFoldersParams {
	return &GetCostReportFoldersParams{
		Context: ctx,
	}
}

// NewGetCostReportFoldersParamsWithHTTPClient creates a new GetCostReportFoldersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCostReportFoldersParamsWithHTTPClient(client *http.Client) *GetCostReportFoldersParams {
	return &GetCostReportFoldersParams{
		HTTPClient: client,
	}
}

/*
GetCostReportFoldersParams contains all the parameters to send to the API endpoint

	for the get cost report folders operation.

	Typically these are written to a http.Request.
*/
type GetCostReportFoldersParams struct {

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

// WithDefaults hydrates default values in the get cost report folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostReportFoldersParams) WithDefaults() *GetCostReportFoldersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cost report folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostReportFoldersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cost report folders params
func (o *GetCostReportFoldersParams) WithTimeout(timeout time.Duration) *GetCostReportFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cost report folders params
func (o *GetCostReportFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cost report folders params
func (o *GetCostReportFoldersParams) WithContext(ctx context.Context) *GetCostReportFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cost report folders params
func (o *GetCostReportFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cost report folders params
func (o *GetCostReportFoldersParams) WithHTTPClient(client *http.Client) *GetCostReportFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cost report folders params
func (o *GetCostReportFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get cost report folders params
func (o *GetCostReportFoldersParams) WithLimit(limit *int32) *GetCostReportFoldersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get cost report folders params
func (o *GetCostReportFoldersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get cost report folders params
func (o *GetCostReportFoldersParams) WithPage(page *int32) *GetCostReportFoldersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get cost report folders params
func (o *GetCostReportFoldersParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetCostReportFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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