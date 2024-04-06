// Code generated by go-swagger; DO NOT EDIT.

package filters

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

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// NewCreateSavedFilterParams creates a new CreateSavedFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSavedFilterParams() *CreateSavedFilterParams {
	return &CreateSavedFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSavedFilterParamsWithTimeout creates a new CreateSavedFilterParams object
// with the ability to set a timeout on a request.
func NewCreateSavedFilterParamsWithTimeout(timeout time.Duration) *CreateSavedFilterParams {
	return &CreateSavedFilterParams{
		timeout: timeout,
	}
}

// NewCreateSavedFilterParamsWithContext creates a new CreateSavedFilterParams object
// with the ability to set a context for a request.
func NewCreateSavedFilterParamsWithContext(ctx context.Context) *CreateSavedFilterParams {
	return &CreateSavedFilterParams{
		Context: ctx,
	}
}

// NewCreateSavedFilterParamsWithHTTPClient creates a new CreateSavedFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSavedFilterParamsWithHTTPClient(client *http.Client) *CreateSavedFilterParams {
	return &CreateSavedFilterParams{
		HTTPClient: client,
	}
}

/*
CreateSavedFilterParams contains all the parameters to send to the API endpoint

	for the create saved filter operation.

	Typically these are written to a http.Request.
*/
type CreateSavedFilterParams struct {

	// CreateSavedFilter.
	CreateSavedFilter *models.CreateSavedFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create saved filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSavedFilterParams) WithDefaults() *CreateSavedFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create saved filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSavedFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create saved filter params
func (o *CreateSavedFilterParams) WithTimeout(timeout time.Duration) *CreateSavedFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create saved filter params
func (o *CreateSavedFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create saved filter params
func (o *CreateSavedFilterParams) WithContext(ctx context.Context) *CreateSavedFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create saved filter params
func (o *CreateSavedFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create saved filter params
func (o *CreateSavedFilterParams) WithHTTPClient(client *http.Client) *CreateSavedFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create saved filter params
func (o *CreateSavedFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateSavedFilter adds the createSavedFilter to the create saved filter params
func (o *CreateSavedFilterParams) WithCreateSavedFilter(createSavedFilter *models.CreateSavedFilter) *CreateSavedFilterParams {
	o.SetCreateSavedFilter(createSavedFilter)
	return o
}

// SetCreateSavedFilter adds the createSavedFilter to the create saved filter params
func (o *CreateSavedFilterParams) SetCreateSavedFilter(createSavedFilter *models.CreateSavedFilter) {
	o.CreateSavedFilter = createSavedFilter
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSavedFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateSavedFilter != nil {
		if err := r.SetBodyParam(o.CreateSavedFilter); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
