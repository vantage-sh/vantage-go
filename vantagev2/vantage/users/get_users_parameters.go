// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUsersParams creates a new GetUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersParams() *GetUsersParams {
	return &GetUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersParamsWithTimeout creates a new GetUsersParams object
// with the ability to set a timeout on a request.
func NewGetUsersParamsWithTimeout(timeout time.Duration) *GetUsersParams {
	return &GetUsersParams{
		timeout: timeout,
	}
}

// NewGetUsersParamsWithContext creates a new GetUsersParams object
// with the ability to set a context for a request.
func NewGetUsersParamsWithContext(ctx context.Context) *GetUsersParams {
	return &GetUsersParams{
		Context: ctx,
	}
}

// NewGetUsersParamsWithHTTPClient creates a new GetUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersParamsWithHTTPClient(client *http.Client) *GetUsersParams {
	return &GetUsersParams{
		HTTPClient: client,
	}
}

/*
GetUsersParams contains all the parameters to send to the API endpoint

	for the get users operation.

	Typically these are written to a http.Request.
*/
type GetUsersParams struct {

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

// WithDefaults hydrates default values in the get users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersParams) WithDefaults() *GetUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get users params
func (o *GetUsersParams) WithTimeout(timeout time.Duration) *GetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users params
func (o *GetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users params
func (o *GetUsersParams) WithContext(ctx context.Context) *GetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users params
func (o *GetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) WithHTTPClient(client *http.Client) *GetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get users params
func (o *GetUsersParams) WithLimit(limit *int32) *GetUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get users params
func (o *GetUsersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get users params
func (o *GetUsersParams) WithPage(page *int32) *GetUsersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get users params
func (o *GetUsersParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
