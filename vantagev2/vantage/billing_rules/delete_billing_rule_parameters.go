// Code generated by go-swagger; DO NOT EDIT.

package billing_rules

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

// NewDeleteBillingRuleParams creates a new DeleteBillingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBillingRuleParams() *DeleteBillingRuleParams {
	return &DeleteBillingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBillingRuleParamsWithTimeout creates a new DeleteBillingRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteBillingRuleParamsWithTimeout(timeout time.Duration) *DeleteBillingRuleParams {
	return &DeleteBillingRuleParams{
		timeout: timeout,
	}
}

// NewDeleteBillingRuleParamsWithContext creates a new DeleteBillingRuleParams object
// with the ability to set a context for a request.
func NewDeleteBillingRuleParamsWithContext(ctx context.Context) *DeleteBillingRuleParams {
	return &DeleteBillingRuleParams{
		Context: ctx,
	}
}

// NewDeleteBillingRuleParamsWithHTTPClient creates a new DeleteBillingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBillingRuleParamsWithHTTPClient(client *http.Client) *DeleteBillingRuleParams {
	return &DeleteBillingRuleParams{
		HTTPClient: client,
	}
}

/*
DeleteBillingRuleParams contains all the parameters to send to the API endpoint

	for the delete billing rule operation.

	Typically these are written to a http.Request.
*/
type DeleteBillingRuleParams struct {

	// BillingRuleToken.
	BillingRuleToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete billing rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBillingRuleParams) WithDefaults() *DeleteBillingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete billing rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBillingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete billing rule params
func (o *DeleteBillingRuleParams) WithTimeout(timeout time.Duration) *DeleteBillingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete billing rule params
func (o *DeleteBillingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete billing rule params
func (o *DeleteBillingRuleParams) WithContext(ctx context.Context) *DeleteBillingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete billing rule params
func (o *DeleteBillingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete billing rule params
func (o *DeleteBillingRuleParams) WithHTTPClient(client *http.Client) *DeleteBillingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete billing rule params
func (o *DeleteBillingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingRuleToken adds the billingRuleToken to the delete billing rule params
func (o *DeleteBillingRuleParams) WithBillingRuleToken(billingRuleToken string) *DeleteBillingRuleParams {
	o.SetBillingRuleToken(billingRuleToken)
	return o
}

// SetBillingRuleToken adds the billingRuleToken to the delete billing rule params
func (o *DeleteBillingRuleParams) SetBillingRuleToken(billingRuleToken string) {
	o.BillingRuleToken = billingRuleToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBillingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billing_rule_token
	if err := r.SetPathParam("billing_rule_token", o.BillingRuleToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
