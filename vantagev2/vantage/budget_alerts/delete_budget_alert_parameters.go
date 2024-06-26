// Code generated by go-swagger; DO NOT EDIT.

package budget_alerts

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

// NewDeleteBudgetAlertParams creates a new DeleteBudgetAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBudgetAlertParams() *DeleteBudgetAlertParams {
	return &DeleteBudgetAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBudgetAlertParamsWithTimeout creates a new DeleteBudgetAlertParams object
// with the ability to set a timeout on a request.
func NewDeleteBudgetAlertParamsWithTimeout(timeout time.Duration) *DeleteBudgetAlertParams {
	return &DeleteBudgetAlertParams{
		timeout: timeout,
	}
}

// NewDeleteBudgetAlertParamsWithContext creates a new DeleteBudgetAlertParams object
// with the ability to set a context for a request.
func NewDeleteBudgetAlertParamsWithContext(ctx context.Context) *DeleteBudgetAlertParams {
	return &DeleteBudgetAlertParams{
		Context: ctx,
	}
}

// NewDeleteBudgetAlertParamsWithHTTPClient creates a new DeleteBudgetAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBudgetAlertParamsWithHTTPClient(client *http.Client) *DeleteBudgetAlertParams {
	return &DeleteBudgetAlertParams{
		HTTPClient: client,
	}
}

/*
DeleteBudgetAlertParams contains all the parameters to send to the API endpoint

	for the delete budget alert operation.

	Typically these are written to a http.Request.
*/
type DeleteBudgetAlertParams struct {

	// BudgetAlertToken.
	BudgetAlertToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete budget alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBudgetAlertParams) WithDefaults() *DeleteBudgetAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete budget alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBudgetAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete budget alert params
func (o *DeleteBudgetAlertParams) WithTimeout(timeout time.Duration) *DeleteBudgetAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete budget alert params
func (o *DeleteBudgetAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete budget alert params
func (o *DeleteBudgetAlertParams) WithContext(ctx context.Context) *DeleteBudgetAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete budget alert params
func (o *DeleteBudgetAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete budget alert params
func (o *DeleteBudgetAlertParams) WithHTTPClient(client *http.Client) *DeleteBudgetAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete budget alert params
func (o *DeleteBudgetAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBudgetAlertToken adds the budgetAlertToken to the delete budget alert params
func (o *DeleteBudgetAlertParams) WithBudgetAlertToken(budgetAlertToken string) *DeleteBudgetAlertParams {
	o.SetBudgetAlertToken(budgetAlertToken)
	return o
}

// SetBudgetAlertToken adds the budgetAlertToken to the delete budget alert params
func (o *DeleteBudgetAlertParams) SetBudgetAlertToken(budgetAlertToken string) {
	o.BudgetAlertToken = budgetAlertToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBudgetAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param budget_alert_token
	if err := r.SetPathParam("budget_alert_token", o.BudgetAlertToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
