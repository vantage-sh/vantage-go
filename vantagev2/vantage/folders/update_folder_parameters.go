// Code generated by go-swagger; DO NOT EDIT.

package folders

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

// NewUpdateFolderParams creates a new UpdateFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateFolderParams() *UpdateFolderParams {
	return &UpdateFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFolderParamsWithTimeout creates a new UpdateFolderParams object
// with the ability to set a timeout on a request.
func NewUpdateFolderParamsWithTimeout(timeout time.Duration) *UpdateFolderParams {
	return &UpdateFolderParams{
		timeout: timeout,
	}
}

// NewUpdateFolderParamsWithContext creates a new UpdateFolderParams object
// with the ability to set a context for a request.
func NewUpdateFolderParamsWithContext(ctx context.Context) *UpdateFolderParams {
	return &UpdateFolderParams{
		Context: ctx,
	}
}

// NewUpdateFolderParamsWithHTTPClient creates a new UpdateFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateFolderParamsWithHTTPClient(client *http.Client) *UpdateFolderParams {
	return &UpdateFolderParams{
		HTTPClient: client,
	}
}

/*
UpdateFolderParams contains all the parameters to send to the API endpoint

	for the update folder operation.

	Typically these are written to a http.Request.
*/
type UpdateFolderParams struct {

	// FolderToken.
	FolderToken string

	// UpdateFolder.
	UpdateFolder *models.UpdateFolder

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFolderParams) WithDefaults() *UpdateFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFolderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update folder params
func (o *UpdateFolderParams) WithTimeout(timeout time.Duration) *UpdateFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update folder params
func (o *UpdateFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update folder params
func (o *UpdateFolderParams) WithContext(ctx context.Context) *UpdateFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update folder params
func (o *UpdateFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update folder params
func (o *UpdateFolderParams) WithHTTPClient(client *http.Client) *UpdateFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update folder params
func (o *UpdateFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolderToken adds the folderToken to the update folder params
func (o *UpdateFolderParams) WithFolderToken(folderToken string) *UpdateFolderParams {
	o.SetFolderToken(folderToken)
	return o
}

// SetFolderToken adds the folderToken to the update folder params
func (o *UpdateFolderParams) SetFolderToken(folderToken string) {
	o.FolderToken = folderToken
}

// WithUpdateFolder adds the updateFolder to the update folder params
func (o *UpdateFolderParams) WithUpdateFolder(updateFolder *models.UpdateFolder) *UpdateFolderParams {
	o.SetUpdateFolder(updateFolder)
	return o
}

// SetUpdateFolder adds the updateFolder to the update folder params
func (o *UpdateFolderParams) SetUpdateFolder(updateFolder *models.UpdateFolder) {
	o.UpdateFolder = updateFolder
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param folder_token
	if err := r.SetPathParam("folder_token", o.FolderToken); err != nil {
		return err
	}
	if o.UpdateFolder != nil {
		if err := r.SetBodyParam(o.UpdateFolder); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
