// Code generated by go-swagger; DO NOT EDIT.

package access_grants

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

// NewDeleteAccessGrantParams creates a new DeleteAccessGrantParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAccessGrantParams() *DeleteAccessGrantParams {
	return &DeleteAccessGrantParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccessGrantParamsWithTimeout creates a new DeleteAccessGrantParams object
// with the ability to set a timeout on a request.
func NewDeleteAccessGrantParamsWithTimeout(timeout time.Duration) *DeleteAccessGrantParams {
	return &DeleteAccessGrantParams{
		timeout: timeout,
	}
}

// NewDeleteAccessGrantParamsWithContext creates a new DeleteAccessGrantParams object
// with the ability to set a context for a request.
func NewDeleteAccessGrantParamsWithContext(ctx context.Context) *DeleteAccessGrantParams {
	return &DeleteAccessGrantParams{
		Context: ctx,
	}
}

// NewDeleteAccessGrantParamsWithHTTPClient creates a new DeleteAccessGrantParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAccessGrantParamsWithHTTPClient(client *http.Client) *DeleteAccessGrantParams {
	return &DeleteAccessGrantParams{
		HTTPClient: client,
	}
}

/*
DeleteAccessGrantParams contains all the parameters to send to the API endpoint

	for the delete access grant operation.

	Typically these are written to a http.Request.
*/
type DeleteAccessGrantParams struct {

	// AccessGrantToken.
	AccessGrantToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete access grant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAccessGrantParams) WithDefaults() *DeleteAccessGrantParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete access grant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAccessGrantParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete access grant params
func (o *DeleteAccessGrantParams) WithTimeout(timeout time.Duration) *DeleteAccessGrantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete access grant params
func (o *DeleteAccessGrantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete access grant params
func (o *DeleteAccessGrantParams) WithContext(ctx context.Context) *DeleteAccessGrantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete access grant params
func (o *DeleteAccessGrantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete access grant params
func (o *DeleteAccessGrantParams) WithHTTPClient(client *http.Client) *DeleteAccessGrantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete access grant params
func (o *DeleteAccessGrantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessGrantToken adds the accessGrantToken to the delete access grant params
func (o *DeleteAccessGrantParams) WithAccessGrantToken(accessGrantToken string) *DeleteAccessGrantParams {
	o.SetAccessGrantToken(accessGrantToken)
	return o
}

// SetAccessGrantToken adds the accessGrantToken to the delete access grant params
func (o *DeleteAccessGrantParams) SetAccessGrantToken(accessGrantToken string) {
	o.AccessGrantToken = accessGrantToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccessGrantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param access_grant_token
	if err := r.SetPathParam("access_grant_token", o.AccessGrantToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
