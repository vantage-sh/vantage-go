// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetIntegrationsParams creates a new GetIntegrationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationsParams() *GetIntegrationsParams {
	return &GetIntegrationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsParamsWithTimeout creates a new GetIntegrationsParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationsParamsWithTimeout(timeout time.Duration) *GetIntegrationsParams {
	return &GetIntegrationsParams{
		timeout: timeout,
	}
}

// NewGetIntegrationsParamsWithContext creates a new GetIntegrationsParams object
// with the ability to set a context for a request.
func NewGetIntegrationsParamsWithContext(ctx context.Context) *GetIntegrationsParams {
	return &GetIntegrationsParams{
		Context: ctx,
	}
}

// NewGetIntegrationsParamsWithHTTPClient creates a new GetIntegrationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationsParamsWithHTTPClient(client *http.Client) *GetIntegrationsParams {
	return &GetIntegrationsParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationsParams contains all the parameters to send to the API endpoint

	for the get integrations operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationsParams struct {

	/* AccountIdentifier.

	   Query by account identifier to list all Integrations that match a specific account. For Azure, this is the subscription ID. Must include provider when using this parameter.
	*/
	AccountIdentifier *string

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

	/* Provider.

	   Query by provider name to list all Integrations for a specific provider.
	*/
	Provider *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integrations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsParams) WithDefaults() *GetIntegrationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integrations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integrations params
func (o *GetIntegrationsParams) WithTimeout(timeout time.Duration) *GetIntegrationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations params
func (o *GetIntegrationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations params
func (o *GetIntegrationsParams) WithContext(ctx context.Context) *GetIntegrationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations params
func (o *GetIntegrationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations params
func (o *GetIntegrationsParams) WithHTTPClient(client *http.Client) *GetIntegrationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations params
func (o *GetIntegrationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountIdentifier adds the accountIdentifier to the get integrations params
func (o *GetIntegrationsParams) WithAccountIdentifier(accountIdentifier *string) *GetIntegrationsParams {
	o.SetAccountIdentifier(accountIdentifier)
	return o
}

// SetAccountIdentifier adds the accountIdentifier to the get integrations params
func (o *GetIntegrationsParams) SetAccountIdentifier(accountIdentifier *string) {
	o.AccountIdentifier = accountIdentifier
}

// WithLimit adds the limit to the get integrations params
func (o *GetIntegrationsParams) WithLimit(limit *int32) *GetIntegrationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get integrations params
func (o *GetIntegrationsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get integrations params
func (o *GetIntegrationsParams) WithPage(page *int32) *GetIntegrationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get integrations params
func (o *GetIntegrationsParams) SetPage(page *int32) {
	o.Page = page
}

// WithProvider adds the provider to the get integrations params
func (o *GetIntegrationsParams) WithProvider(provider *string) *GetIntegrationsParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get integrations params
func (o *GetIntegrationsParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountIdentifier != nil {

		// query param account_identifier
		var qrAccountIdentifier string

		if o.AccountIdentifier != nil {
			qrAccountIdentifier = *o.AccountIdentifier
		}
		qAccountIdentifier := qrAccountIdentifier
		if qAccountIdentifier != "" {

			if err := r.SetQueryParam("account_identifier", qAccountIdentifier); err != nil {
				return err
			}
		}
	}

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

	if o.Provider != nil {

		// query param provider
		var qrProvider string

		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {

			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
