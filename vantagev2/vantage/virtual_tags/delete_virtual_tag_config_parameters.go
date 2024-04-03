// Code generated by go-swagger; DO NOT EDIT.

package virtual_tags

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

// NewDeleteVirtualTagConfigParams creates a new DeleteVirtualTagConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVirtualTagConfigParams() *DeleteVirtualTagConfigParams {
	return &DeleteVirtualTagConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVirtualTagConfigParamsWithTimeout creates a new DeleteVirtualTagConfigParams object
// with the ability to set a timeout on a request.
func NewDeleteVirtualTagConfigParamsWithTimeout(timeout time.Duration) *DeleteVirtualTagConfigParams {
	return &DeleteVirtualTagConfigParams{
		timeout: timeout,
	}
}

// NewDeleteVirtualTagConfigParamsWithContext creates a new DeleteVirtualTagConfigParams object
// with the ability to set a context for a request.
func NewDeleteVirtualTagConfigParamsWithContext(ctx context.Context) *DeleteVirtualTagConfigParams {
	return &DeleteVirtualTagConfigParams{
		Context: ctx,
	}
}

// NewDeleteVirtualTagConfigParamsWithHTTPClient creates a new DeleteVirtualTagConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVirtualTagConfigParamsWithHTTPClient(client *http.Client) *DeleteVirtualTagConfigParams {
	return &DeleteVirtualTagConfigParams{
		HTTPClient: client,
	}
}

/*
DeleteVirtualTagConfigParams contains all the parameters to send to the API endpoint

	for the delete virtual tag config operation.

	Typically these are written to a http.Request.
*/
type DeleteVirtualTagConfigParams struct {

	// VirtualTagConfigToken.
	VirtualTagConfigToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete virtual tag config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualTagConfigParams) WithDefaults() *DeleteVirtualTagConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete virtual tag config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualTagConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete virtual tag config params
func (o *DeleteVirtualTagConfigParams) WithTimeout(timeout time.Duration) *DeleteVirtualTagConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete virtual tag config params
func (o *DeleteVirtualTagConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete virtual tag config params
func (o *DeleteVirtualTagConfigParams) WithContext(ctx context.Context) *DeleteVirtualTagConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete virtual tag config params
func (o *DeleteVirtualTagConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete virtual tag config params
func (o *DeleteVirtualTagConfigParams) WithHTTPClient(client *http.Client) *DeleteVirtualTagConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete virtual tag config params
func (o *DeleteVirtualTagConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualTagConfigToken adds the virtualTagConfigToken to the delete virtual tag config params
func (o *DeleteVirtualTagConfigParams) WithVirtualTagConfigToken(virtualTagConfigToken string) *DeleteVirtualTagConfigParams {
	o.SetVirtualTagConfigToken(virtualTagConfigToken)
	return o
}

// SetVirtualTagConfigToken adds the virtualTagConfigToken to the delete virtual tag config params
func (o *DeleteVirtualTagConfigParams) SetVirtualTagConfigToken(virtualTagConfigToken string) {
	o.VirtualTagConfigToken = virtualTagConfigToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVirtualTagConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param virtual_tag_config_token
	if err := r.SetPathParam("virtual_tag_config_token", o.VirtualTagConfigToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
