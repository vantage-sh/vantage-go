// Code generated by go-swagger; DO NOT EDIT.

package segments

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

// NewGetSegmentParams creates a new GetSegmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSegmentParams() *GetSegmentParams {
	return &GetSegmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSegmentParamsWithTimeout creates a new GetSegmentParams object
// with the ability to set a timeout on a request.
func NewGetSegmentParamsWithTimeout(timeout time.Duration) *GetSegmentParams {
	return &GetSegmentParams{
		timeout: timeout,
	}
}

// NewGetSegmentParamsWithContext creates a new GetSegmentParams object
// with the ability to set a context for a request.
func NewGetSegmentParamsWithContext(ctx context.Context) *GetSegmentParams {
	return &GetSegmentParams{
		Context: ctx,
	}
}

// NewGetSegmentParamsWithHTTPClient creates a new GetSegmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSegmentParamsWithHTTPClient(client *http.Client) *GetSegmentParams {
	return &GetSegmentParams{
		HTTPClient: client,
	}
}

/*
GetSegmentParams contains all the parameters to send to the API endpoint

	for the get segment operation.

	Typically these are written to a http.Request.
*/
type GetSegmentParams struct {

	// SegmentToken.
	SegmentToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get segment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSegmentParams) WithDefaults() *GetSegmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get segment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSegmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get segment params
func (o *GetSegmentParams) WithTimeout(timeout time.Duration) *GetSegmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get segment params
func (o *GetSegmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get segment params
func (o *GetSegmentParams) WithContext(ctx context.Context) *GetSegmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get segment params
func (o *GetSegmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get segment params
func (o *GetSegmentParams) WithHTTPClient(client *http.Client) *GetSegmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get segment params
func (o *GetSegmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSegmentToken adds the segmentToken to the get segment params
func (o *GetSegmentParams) WithSegmentToken(segmentToken string) *GetSegmentParams {
	o.SetSegmentToken(segmentToken)
	return o
}

// SetSegmentToken adds the segmentToken to the get segment params
func (o *GetSegmentParams) SetSegmentToken(segmentToken string) {
	o.SegmentToken = segmentToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetSegmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param segment_token
	if err := r.SetPathParam("segment_token", o.SegmentToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
