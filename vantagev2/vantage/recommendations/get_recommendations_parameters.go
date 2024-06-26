// Code generated by go-swagger; DO NOT EDIT.

package recommendations

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

// NewGetRecommendationsParams creates a new GetRecommendationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRecommendationsParams() *GetRecommendationsParams {
	return &GetRecommendationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecommendationsParamsWithTimeout creates a new GetRecommendationsParams object
// with the ability to set a timeout on a request.
func NewGetRecommendationsParamsWithTimeout(timeout time.Duration) *GetRecommendationsParams {
	return &GetRecommendationsParams{
		timeout: timeout,
	}
}

// NewGetRecommendationsParamsWithContext creates a new GetRecommendationsParams object
// with the ability to set a context for a request.
func NewGetRecommendationsParamsWithContext(ctx context.Context) *GetRecommendationsParams {
	return &GetRecommendationsParams{
		Context: ctx,
	}
}

// NewGetRecommendationsParamsWithHTTPClient creates a new GetRecommendationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRecommendationsParamsWithHTTPClient(client *http.Client) *GetRecommendationsParams {
	return &GetRecommendationsParams{
		HTTPClient: client,
	}
}

/*
GetRecommendationsParams contains all the parameters to send to the API endpoint

	for the get recommendations operation.

	Typically these are written to a http.Request.
*/
type GetRecommendationsParams struct {

	/* Category.

	   Filter by category.
	*/
	Category *string

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

	   Filter by provider.
	*/
	Provider *string

	/* ProviderAccountID.

	   Filter by provider account id (AWS account, Azure subscription id, etc).
	*/
	ProviderAccountID *string

	/* WorkspaceToken.

	   Filter by workspace.
	*/
	WorkspaceToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get recommendations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRecommendationsParams) WithDefaults() *GetRecommendationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get recommendations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRecommendationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get recommendations params
func (o *GetRecommendationsParams) WithTimeout(timeout time.Duration) *GetRecommendationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recommendations params
func (o *GetRecommendationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recommendations params
func (o *GetRecommendationsParams) WithContext(ctx context.Context) *GetRecommendationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recommendations params
func (o *GetRecommendationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recommendations params
func (o *GetRecommendationsParams) WithHTTPClient(client *http.Client) *GetRecommendationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recommendations params
func (o *GetRecommendationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get recommendations params
func (o *GetRecommendationsParams) WithCategory(category *string) *GetRecommendationsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get recommendations params
func (o *GetRecommendationsParams) SetCategory(category *string) {
	o.Category = category
}

// WithLimit adds the limit to the get recommendations params
func (o *GetRecommendationsParams) WithLimit(limit *int32) *GetRecommendationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get recommendations params
func (o *GetRecommendationsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get recommendations params
func (o *GetRecommendationsParams) WithPage(page *int32) *GetRecommendationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get recommendations params
func (o *GetRecommendationsParams) SetPage(page *int32) {
	o.Page = page
}

// WithProvider adds the provider to the get recommendations params
func (o *GetRecommendationsParams) WithProvider(provider *string) *GetRecommendationsParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get recommendations params
func (o *GetRecommendationsParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithProviderAccountID adds the providerAccountID to the get recommendations params
func (o *GetRecommendationsParams) WithProviderAccountID(providerAccountID *string) *GetRecommendationsParams {
	o.SetProviderAccountID(providerAccountID)
	return o
}

// SetProviderAccountID adds the providerAccountId to the get recommendations params
func (o *GetRecommendationsParams) SetProviderAccountID(providerAccountID *string) {
	o.ProviderAccountID = providerAccountID
}

// WithWorkspaceToken adds the workspaceToken to the get recommendations params
func (o *GetRecommendationsParams) WithWorkspaceToken(workspaceToken *string) *GetRecommendationsParams {
	o.SetWorkspaceToken(workspaceToken)
	return o
}

// SetWorkspaceToken adds the workspaceToken to the get recommendations params
func (o *GetRecommendationsParams) SetWorkspaceToken(workspaceToken *string) {
	o.WorkspaceToken = workspaceToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecommendationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// query param category
		var qrCategory string

		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := qrCategory
		if qCategory != "" {

			if err := r.SetQueryParam("category", qCategory); err != nil {
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

	if o.ProviderAccountID != nil {

		// query param provider_account_id
		var qrProviderAccountID string

		if o.ProviderAccountID != nil {
			qrProviderAccountID = *o.ProviderAccountID
		}
		qProviderAccountID := qrProviderAccountID
		if qProviderAccountID != "" {

			if err := r.SetQueryParam("provider_account_id", qProviderAccountID); err != nil {
				return err
			}
		}
	}

	if o.WorkspaceToken != nil {

		// query param workspace_token
		var qrWorkspaceToken string

		if o.WorkspaceToken != nil {
			qrWorkspaceToken = *o.WorkspaceToken
		}
		qWorkspaceToken := qrWorkspaceToken
		if qWorkspaceToken != "" {

			if err := r.SetQueryParam("workspace_token", qWorkspaceToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
