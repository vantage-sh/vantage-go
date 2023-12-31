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

// NewGetPricesParams creates a new GetPricesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPricesParams() *GetPricesParams {
	return &GetPricesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPricesParamsWithTimeout creates a new GetPricesParams object
// with the ability to set a timeout on a request.
func NewGetPricesParamsWithTimeout(timeout time.Duration) *GetPricesParams {
	return &GetPricesParams{
		timeout: timeout,
	}
}

// NewGetPricesParamsWithContext creates a new GetPricesParams object
// with the ability to set a context for a request.
func NewGetPricesParamsWithContext(ctx context.Context) *GetPricesParams {
	return &GetPricesParams{
		Context: ctx,
	}
}

// NewGetPricesParamsWithHTTPClient creates a new GetPricesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPricesParamsWithHTTPClient(client *http.Client) *GetPricesParams {
	return &GetPricesParams{
		HTTPClient: client,
	}
}

/*
GetPricesParams contains all the parameters to send to the API endpoint

	for the get prices operation.

	Typically these are written to a http.Request.
*/
type GetPricesParams struct {

	/* Limit.

	   The amount of results to return. The maximum is 1000

	   Format: int32
	*/
	Limit *int32

	/* Page.

	   The page of results to return.

	   Format: int32
	*/
	Page *int32

	// ProductID.
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get prices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPricesParams) WithDefaults() *GetPricesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get prices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPricesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get prices params
func (o *GetPricesParams) WithTimeout(timeout time.Duration) *GetPricesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get prices params
func (o *GetPricesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get prices params
func (o *GetPricesParams) WithContext(ctx context.Context) *GetPricesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get prices params
func (o *GetPricesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get prices params
func (o *GetPricesParams) WithHTTPClient(client *http.Client) *GetPricesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get prices params
func (o *GetPricesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get prices params
func (o *GetPricesParams) WithLimit(limit *int32) *GetPricesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get prices params
func (o *GetPricesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get prices params
func (o *GetPricesParams) WithPage(page *int32) *GetPricesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get prices params
func (o *GetPricesParams) SetPage(page *int32) {
	o.Page = page
}

// WithProductID adds the productID to the get prices params
func (o *GetPricesParams) WithProductID(productID string) *GetPricesParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get prices params
func (o *GetPricesParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPricesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param product_id
	if err := r.SetPathParam("product_id", o.ProductID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
