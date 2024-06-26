// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_efficiency_reports

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

// NewDeleteKubernetesEfficiencyReportParams creates a new DeleteKubernetesEfficiencyReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKubernetesEfficiencyReportParams() *DeleteKubernetesEfficiencyReportParams {
	return &DeleteKubernetesEfficiencyReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKubernetesEfficiencyReportParamsWithTimeout creates a new DeleteKubernetesEfficiencyReportParams object
// with the ability to set a timeout on a request.
func NewDeleteKubernetesEfficiencyReportParamsWithTimeout(timeout time.Duration) *DeleteKubernetesEfficiencyReportParams {
	return &DeleteKubernetesEfficiencyReportParams{
		timeout: timeout,
	}
}

// NewDeleteKubernetesEfficiencyReportParamsWithContext creates a new DeleteKubernetesEfficiencyReportParams object
// with the ability to set a context for a request.
func NewDeleteKubernetesEfficiencyReportParamsWithContext(ctx context.Context) *DeleteKubernetesEfficiencyReportParams {
	return &DeleteKubernetesEfficiencyReportParams{
		Context: ctx,
	}
}

// NewDeleteKubernetesEfficiencyReportParamsWithHTTPClient creates a new DeleteKubernetesEfficiencyReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKubernetesEfficiencyReportParamsWithHTTPClient(client *http.Client) *DeleteKubernetesEfficiencyReportParams {
	return &DeleteKubernetesEfficiencyReportParams{
		HTTPClient: client,
	}
}

/*
DeleteKubernetesEfficiencyReportParams contains all the parameters to send to the API endpoint

	for the delete kubernetes efficiency report operation.

	Typically these are written to a http.Request.
*/
type DeleteKubernetesEfficiencyReportParams struct {

	// KubernetesEfficiencyReportToken.
	KubernetesEfficiencyReportToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete kubernetes efficiency report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKubernetesEfficiencyReportParams) WithDefaults() *DeleteKubernetesEfficiencyReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete kubernetes efficiency report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKubernetesEfficiencyReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete kubernetes efficiency report params
func (o *DeleteKubernetesEfficiencyReportParams) WithTimeout(timeout time.Duration) *DeleteKubernetesEfficiencyReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete kubernetes efficiency report params
func (o *DeleteKubernetesEfficiencyReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete kubernetes efficiency report params
func (o *DeleteKubernetesEfficiencyReportParams) WithContext(ctx context.Context) *DeleteKubernetesEfficiencyReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete kubernetes efficiency report params
func (o *DeleteKubernetesEfficiencyReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete kubernetes efficiency report params
func (o *DeleteKubernetesEfficiencyReportParams) WithHTTPClient(client *http.Client) *DeleteKubernetesEfficiencyReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete kubernetes efficiency report params
func (o *DeleteKubernetesEfficiencyReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKubernetesEfficiencyReportToken adds the kubernetesEfficiencyReportToken to the delete kubernetes efficiency report params
func (o *DeleteKubernetesEfficiencyReportParams) WithKubernetesEfficiencyReportToken(kubernetesEfficiencyReportToken string) *DeleteKubernetesEfficiencyReportParams {
	o.SetKubernetesEfficiencyReportToken(kubernetesEfficiencyReportToken)
	return o
}

// SetKubernetesEfficiencyReportToken adds the kubernetesEfficiencyReportToken to the delete kubernetes efficiency report params
func (o *DeleteKubernetesEfficiencyReportParams) SetKubernetesEfficiencyReportToken(kubernetesEfficiencyReportToken string) {
	o.KubernetesEfficiencyReportToken = kubernetesEfficiencyReportToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKubernetesEfficiencyReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kubernetes_efficiency_report_token
	if err := r.SetPathParam("kubernetes_efficiency_report_token", o.KubernetesEfficiencyReportToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
