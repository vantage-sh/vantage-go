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
	"github.com/go-openapi/swag"
)

// NewGetResourceReportsParams creates a new GetResourceReportsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceReportsParams() *GetResourceReportsParams {
	return &GetResourceReportsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceReportsParamsWithTimeout creates a new GetResourceReportsParams object
// with the ability to set a timeout on a request.
func NewGetResourceReportsParamsWithTimeout(timeout time.Duration) *GetResourceReportsParams {
	return &GetResourceReportsParams{
		timeout: timeout,
	}
}

// NewGetResourceReportsParamsWithContext creates a new GetResourceReportsParams object
// with the ability to set a context for a request.
func NewGetResourceReportsParamsWithContext(ctx context.Context) *GetResourceReportsParams {
	return &GetResourceReportsParams{
		Context: ctx,
	}
}

// NewGetResourceReportsParamsWithHTTPClient creates a new GetResourceReportsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceReportsParamsWithHTTPClient(client *http.Client) *GetResourceReportsParams {
	return &GetResourceReportsParams{
		HTTPClient: client,
	}
}

/*
GetResourceReportsParams contains all the parameters to send to the API endpoint

	for the get resource reports operation.

	Typically these are written to a http.Request.
*/
type GetResourceReportsParams struct {

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

// WithDefaults hydrates default values in the get resource reports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceReportsParams) WithDefaults() *GetResourceReportsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource reports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceReportsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource reports params
func (o *GetResourceReportsParams) WithTimeout(timeout time.Duration) *GetResourceReportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource reports params
func (o *GetResourceReportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource reports params
func (o *GetResourceReportsParams) WithContext(ctx context.Context) *GetResourceReportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource reports params
func (o *GetResourceReportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource reports params
func (o *GetResourceReportsParams) WithHTTPClient(client *http.Client) *GetResourceReportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource reports params
func (o *GetResourceReportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get resource reports params
func (o *GetResourceReportsParams) WithLimit(limit *int32) *GetResourceReportsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get resource reports params
func (o *GetResourceReportsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get resource reports params
func (o *GetResourceReportsParams) WithPage(page *int32) *GetResourceReportsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get resource reports params
func (o *GetResourceReportsParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceReportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
