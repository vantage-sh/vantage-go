// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewDeleteTeamParams creates a new DeleteTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTeamParams() *DeleteTeamParams {
	return &DeleteTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamParamsWithTimeout creates a new DeleteTeamParams object
// with the ability to set a timeout on a request.
func NewDeleteTeamParamsWithTimeout(timeout time.Duration) *DeleteTeamParams {
	return &DeleteTeamParams{
		timeout: timeout,
	}
}

// NewDeleteTeamParamsWithContext creates a new DeleteTeamParams object
// with the ability to set a context for a request.
func NewDeleteTeamParamsWithContext(ctx context.Context) *DeleteTeamParams {
	return &DeleteTeamParams{
		Context: ctx,
	}
}

// NewDeleteTeamParamsWithHTTPClient creates a new DeleteTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTeamParamsWithHTTPClient(client *http.Client) *DeleteTeamParams {
	return &DeleteTeamParams{
		HTTPClient: client,
	}
}

/*
DeleteTeamParams contains all the parameters to send to the API endpoint

	for the delete team operation.

	Typically these are written to a http.Request.
*/
type DeleteTeamParams struct {

	// TeamToken.
	TeamToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeamParams) WithDefaults() *DeleteTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete team params
func (o *DeleteTeamParams) WithTimeout(timeout time.Duration) *DeleteTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete team params
func (o *DeleteTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete team params
func (o *DeleteTeamParams) WithContext(ctx context.Context) *DeleteTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete team params
func (o *DeleteTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete team params
func (o *DeleteTeamParams) WithHTTPClient(client *http.Client) *DeleteTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete team params
func (o *DeleteTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamToken adds the teamToken to the delete team params
func (o *DeleteTeamParams) WithTeamToken(teamToken string) *DeleteTeamParams {
	o.SetTeamToken(teamToken)
	return o
}

// SetTeamToken adds the teamToken to the delete team params
func (o *DeleteTeamParams) SetTeamToken(teamToken string) {
	o.TeamToken = teamToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team_token
	if err := r.SetPathParam("team_token", o.TeamToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
