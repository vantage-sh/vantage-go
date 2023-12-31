// Code generated by go-swagger; DO NOT EDIT.

package prices

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

// NewGetProductsParams creates a new GetProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProductsParams() *GetProductsParams {
	return &GetProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsParamsWithTimeout creates a new GetProductsParams object
// with the ability to set a timeout on a request.
func NewGetProductsParamsWithTimeout(timeout time.Duration) *GetProductsParams {
	return &GetProductsParams{
		timeout: timeout,
	}
}

// NewGetProductsParamsWithContext creates a new GetProductsParams object
// with the ability to set a context for a request.
func NewGetProductsParamsWithContext(ctx context.Context) *GetProductsParams {
	return &GetProductsParams{
		Context: ctx,
	}
}

// NewGetProductsParamsWithHTTPClient creates a new GetProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProductsParamsWithHTTPClient(client *http.Client) *GetProductsParams {
	return &GetProductsParams{
		HTTPClient: client,
	}
}

/*
GetProductsParams contains all the parameters to send to the API endpoint

	for the get products operation.

	Typically these are written to a http.Request.
*/
type GetProductsParams struct {

	/* Limit.

	   The amount of results to return. The maximum is 1000

	   Format: int32
	*/
	Limit *int32

	/* Name.

	   Query by name of the Product to see a list of products which match that name. e.g. m5a.16xlarge
	*/
	Name *string

	/* Page.

	   The page of results to return.

	   Format: int32
	*/
	Page *int32

	/* ProviderID.

	   Query by Provider to list all Products across all Services for a Provider. e.g. aws
	*/
	ProviderID *string

	/* ServiceID.

	   Query by Service to list all Products for a specific provider service. e.g. aws-ec2
	*/
	ServiceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductsParams) WithDefaults() *GetProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get products params
func (o *GetProductsParams) WithTimeout(timeout time.Duration) *GetProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products params
func (o *GetProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products params
func (o *GetProductsParams) WithContext(ctx context.Context) *GetProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products params
func (o *GetProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products params
func (o *GetProductsParams) WithHTTPClient(client *http.Client) *GetProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products params
func (o *GetProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get products params
func (o *GetProductsParams) WithLimit(limit *int32) *GetProductsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get products params
func (o *GetProductsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithName adds the name to the get products params
func (o *GetProductsParams) WithName(name *string) *GetProductsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get products params
func (o *GetProductsParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the get products params
func (o *GetProductsParams) WithPage(page *int32) *GetProductsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get products params
func (o *GetProductsParams) SetPage(page *int32) {
	o.Page = page
}

// WithProviderID adds the providerID to the get products params
func (o *GetProductsParams) WithProviderID(providerID *string) *GetProductsParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the get products params
func (o *GetProductsParams) SetProviderID(providerID *string) {
	o.ProviderID = providerID
}

// WithServiceID adds the serviceID to the get products params
func (o *GetProductsParams) WithServiceID(serviceID *string) *GetProductsParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the get products params
func (o *GetProductsParams) SetServiceID(serviceID *string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.ProviderID != nil {

		// query param provider_id
		var qrProviderID string

		if o.ProviderID != nil {
			qrProviderID = *o.ProviderID
		}
		qProviderID := qrProviderID
		if qProviderID != "" {

			if err := r.SetQueryParam("provider_id", qProviderID); err != nil {
				return err
			}
		}
	}

	if o.ServiceID != nil {

		// query param service_id
		var qrServiceID string

		if o.ServiceID != nil {
			qrServiceID = *o.ServiceID
		}
		qServiceID := qrServiceID
		if qServiceID != "" {

			if err := r.SetQueryParam("service_id", qServiceID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
