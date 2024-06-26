// Code generated by go-swagger; DO NOT EDIT.

package report_notifications

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

// NewUpdateReportNotificationParams creates a new UpdateReportNotificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateReportNotificationParams() *UpdateReportNotificationParams {
	return &UpdateReportNotificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReportNotificationParamsWithTimeout creates a new UpdateReportNotificationParams object
// with the ability to set a timeout on a request.
func NewUpdateReportNotificationParamsWithTimeout(timeout time.Duration) *UpdateReportNotificationParams {
	return &UpdateReportNotificationParams{
		timeout: timeout,
	}
}

// NewUpdateReportNotificationParamsWithContext creates a new UpdateReportNotificationParams object
// with the ability to set a context for a request.
func NewUpdateReportNotificationParamsWithContext(ctx context.Context) *UpdateReportNotificationParams {
	return &UpdateReportNotificationParams{
		Context: ctx,
	}
}

// NewUpdateReportNotificationParamsWithHTTPClient creates a new UpdateReportNotificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateReportNotificationParamsWithHTTPClient(client *http.Client) *UpdateReportNotificationParams {
	return &UpdateReportNotificationParams{
		HTTPClient: client,
	}
}

/*
UpdateReportNotificationParams contains all the parameters to send to the API endpoint

	for the update report notification operation.

	Typically these are written to a http.Request.
*/
type UpdateReportNotificationParams struct {

	// ReportNotificationToken.
	ReportNotificationToken string

	// UpdateReportNotification.
	UpdateReportNotification *models.UpdateReportNotification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update report notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateReportNotificationParams) WithDefaults() *UpdateReportNotificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update report notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateReportNotificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update report notification params
func (o *UpdateReportNotificationParams) WithTimeout(timeout time.Duration) *UpdateReportNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update report notification params
func (o *UpdateReportNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update report notification params
func (o *UpdateReportNotificationParams) WithContext(ctx context.Context) *UpdateReportNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update report notification params
func (o *UpdateReportNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update report notification params
func (o *UpdateReportNotificationParams) WithHTTPClient(client *http.Client) *UpdateReportNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update report notification params
func (o *UpdateReportNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportNotificationToken adds the reportNotificationToken to the update report notification params
func (o *UpdateReportNotificationParams) WithReportNotificationToken(reportNotificationToken string) *UpdateReportNotificationParams {
	o.SetReportNotificationToken(reportNotificationToken)
	return o
}

// SetReportNotificationToken adds the reportNotificationToken to the update report notification params
func (o *UpdateReportNotificationParams) SetReportNotificationToken(reportNotificationToken string) {
	o.ReportNotificationToken = reportNotificationToken
}

// WithUpdateReportNotification adds the updateReportNotification to the update report notification params
func (o *UpdateReportNotificationParams) WithUpdateReportNotification(updateReportNotification *models.UpdateReportNotification) *UpdateReportNotificationParams {
	o.SetUpdateReportNotification(updateReportNotification)
	return o
}

// SetUpdateReportNotification adds the updateReportNotification to the update report notification params
func (o *UpdateReportNotificationParams) SetUpdateReportNotification(updateReportNotification *models.UpdateReportNotification) {
	o.UpdateReportNotification = updateReportNotification
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReportNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param report_notification_token
	if err := r.SetPathParam("report_notification_token", o.ReportNotificationToken); err != nil {
		return err
	}
	if o.UpdateReportNotification != nil {
		if err := r.SetBodyParam(o.UpdateReportNotification); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
