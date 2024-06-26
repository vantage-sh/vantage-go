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
	"github.com/go-openapi/swag"
)

// NewUpdateBudgetAlertParams creates a new UpdateBudgetAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBudgetAlertParams() *UpdateBudgetAlertParams {
	return &UpdateBudgetAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBudgetAlertParamsWithTimeout creates a new UpdateBudgetAlertParams object
// with the ability to set a timeout on a request.
func NewUpdateBudgetAlertParamsWithTimeout(timeout time.Duration) *UpdateBudgetAlertParams {
	return &UpdateBudgetAlertParams{
		timeout: timeout,
	}
}

// NewUpdateBudgetAlertParamsWithContext creates a new UpdateBudgetAlertParams object
// with the ability to set a context for a request.
func NewUpdateBudgetAlertParamsWithContext(ctx context.Context) *UpdateBudgetAlertParams {
	return &UpdateBudgetAlertParams{
		Context: ctx,
	}
}

// NewUpdateBudgetAlertParamsWithHTTPClient creates a new UpdateBudgetAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBudgetAlertParamsWithHTTPClient(client *http.Client) *UpdateBudgetAlertParams {
	return &UpdateBudgetAlertParams{
		HTTPClient: client,
	}
}

/*
UpdateBudgetAlertParams contains all the parameters to send to the API endpoint

	for the update budget alert operation.

	Typically these are written to a http.Request.
*/
type UpdateBudgetAlertParams struct {

	// BudgetAlertToken.
	BudgetAlertToken string

	/* BudgetTokens.

	   The tokens of the Budget that has the alert.
	*/
	BudgetTokens []string

	/* DurationInDays.

	   The number of days from the start or end of the month to trigger the alert if the threshold is reached. For the full month, pass an empty value.

	   Format: int32
	*/
	DurationInDays *int32

	/* PeriodToTrack.

	   The period tracked on the alert. Used with duration_in_days to determine the time window of the alert. Defaults to start_of_the_month if not passed. Possible values: start_of_the_month, end_of_the_month.
	*/
	PeriodToTrack *string

	/* RecipientChannels.

	   The channels receiving the alerts. Requires an integration provider to be connected.
	*/
	RecipientChannels []string

	/* Threshold.

	   The threshold amount that must be met for the alert to fire.

	   Format: int32
	*/
	Threshold *int32

	/* UserTokens.

	   The tokens of the users that receive the alert.
	*/
	UserTokens []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update budget alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBudgetAlertParams) WithDefaults() *UpdateBudgetAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update budget alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBudgetAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update budget alert params
func (o *UpdateBudgetAlertParams) WithTimeout(timeout time.Duration) *UpdateBudgetAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update budget alert params
func (o *UpdateBudgetAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update budget alert params
func (o *UpdateBudgetAlertParams) WithContext(ctx context.Context) *UpdateBudgetAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update budget alert params
func (o *UpdateBudgetAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update budget alert params
func (o *UpdateBudgetAlertParams) WithHTTPClient(client *http.Client) *UpdateBudgetAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update budget alert params
func (o *UpdateBudgetAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBudgetAlertToken adds the budgetAlertToken to the update budget alert params
func (o *UpdateBudgetAlertParams) WithBudgetAlertToken(budgetAlertToken string) *UpdateBudgetAlertParams {
	o.SetBudgetAlertToken(budgetAlertToken)
	return o
}

// SetBudgetAlertToken adds the budgetAlertToken to the update budget alert params
func (o *UpdateBudgetAlertParams) SetBudgetAlertToken(budgetAlertToken string) {
	o.BudgetAlertToken = budgetAlertToken
}

// WithBudgetTokens adds the budgetTokens to the update budget alert params
func (o *UpdateBudgetAlertParams) WithBudgetTokens(budgetTokens []string) *UpdateBudgetAlertParams {
	o.SetBudgetTokens(budgetTokens)
	return o
}

// SetBudgetTokens adds the budgetTokens to the update budget alert params
func (o *UpdateBudgetAlertParams) SetBudgetTokens(budgetTokens []string) {
	o.BudgetTokens = budgetTokens
}

// WithDurationInDays adds the durationInDays to the update budget alert params
func (o *UpdateBudgetAlertParams) WithDurationInDays(durationInDays *int32) *UpdateBudgetAlertParams {
	o.SetDurationInDays(durationInDays)
	return o
}

// SetDurationInDays adds the durationInDays to the update budget alert params
func (o *UpdateBudgetAlertParams) SetDurationInDays(durationInDays *int32) {
	o.DurationInDays = durationInDays
}

// WithPeriodToTrack adds the periodToTrack to the update budget alert params
func (o *UpdateBudgetAlertParams) WithPeriodToTrack(periodToTrack *string) *UpdateBudgetAlertParams {
	o.SetPeriodToTrack(periodToTrack)
	return o
}

// SetPeriodToTrack adds the periodToTrack to the update budget alert params
func (o *UpdateBudgetAlertParams) SetPeriodToTrack(periodToTrack *string) {
	o.PeriodToTrack = periodToTrack
}

// WithRecipientChannels adds the recipientChannels to the update budget alert params
func (o *UpdateBudgetAlertParams) WithRecipientChannels(recipientChannels []string) *UpdateBudgetAlertParams {
	o.SetRecipientChannels(recipientChannels)
	return o
}

// SetRecipientChannels adds the recipientChannels to the update budget alert params
func (o *UpdateBudgetAlertParams) SetRecipientChannels(recipientChannels []string) {
	o.RecipientChannels = recipientChannels
}

// WithThreshold adds the threshold to the update budget alert params
func (o *UpdateBudgetAlertParams) WithThreshold(threshold *int32) *UpdateBudgetAlertParams {
	o.SetThreshold(threshold)
	return o
}

// SetThreshold adds the threshold to the update budget alert params
func (o *UpdateBudgetAlertParams) SetThreshold(threshold *int32) {
	o.Threshold = threshold
}

// WithUserTokens adds the userTokens to the update budget alert params
func (o *UpdateBudgetAlertParams) WithUserTokens(userTokens []string) *UpdateBudgetAlertParams {
	o.SetUserTokens(userTokens)
	return o
}

// SetUserTokens adds the userTokens to the update budget alert params
func (o *UpdateBudgetAlertParams) SetUserTokens(userTokens []string) {
	o.UserTokens = userTokens
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBudgetAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param budget_alert_token
	if err := r.SetPathParam("budget_alert_token", o.BudgetAlertToken); err != nil {
		return err
	}

	if o.BudgetTokens != nil {

		// binding items for budget_tokens
		joinedBudgetTokens := o.bindParamBudgetTokens(reg)

		// form array param budget_tokens
		if err := r.SetFormParam("budget_tokens", joinedBudgetTokens...); err != nil {
			return err
		}
	}

	if o.DurationInDays != nil {

		// form param duration_in_days
		var frDurationInDays int32
		if o.DurationInDays != nil {
			frDurationInDays = *o.DurationInDays
		}
		fDurationInDays := swag.FormatInt32(frDurationInDays)
		if fDurationInDays != "" {
			if err := r.SetFormParam("duration_in_days", fDurationInDays); err != nil {
				return err
			}
		}
	}

	if o.PeriodToTrack != nil {

		// form param period_to_track
		var frPeriodToTrack string
		if o.PeriodToTrack != nil {
			frPeriodToTrack = *o.PeriodToTrack
		}
		fPeriodToTrack := frPeriodToTrack
		if fPeriodToTrack != "" {
			if err := r.SetFormParam("period_to_track", fPeriodToTrack); err != nil {
				return err
			}
		}
	}

	if o.RecipientChannels != nil {

		// binding items for recipient_channels
		joinedRecipientChannels := o.bindParamRecipientChannels(reg)

		// form array param recipient_channels
		if err := r.SetFormParam("recipient_channels", joinedRecipientChannels...); err != nil {
			return err
		}
	}

	if o.Threshold != nil {

		// form param threshold
		var frThreshold int32
		if o.Threshold != nil {
			frThreshold = *o.Threshold
		}
		fThreshold := swag.FormatInt32(frThreshold)
		if fThreshold != "" {
			if err := r.SetFormParam("threshold", fThreshold); err != nil {
				return err
			}
		}
	}

	if o.UserTokens != nil {

		// binding items for user_tokens
		joinedUserTokens := o.bindParamUserTokens(reg)

		// form array param user_tokens
		if err := r.SetFormParam("user_tokens", joinedUserTokens...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUpdateBudgetAlert binds the parameter budget_tokens
func (o *UpdateBudgetAlertParams) bindParamBudgetTokens(formats strfmt.Registry) []string {
	budgetTokensIR := o.BudgetTokens

	var budgetTokensIC []string
	for _, budgetTokensIIR := range budgetTokensIR { // explode []string

		budgetTokensIIV := budgetTokensIIR // string as string
		budgetTokensIC = append(budgetTokensIC, budgetTokensIIV)
	}

	// items.CollectionFormat: ""
	budgetTokensIS := swag.JoinByFormat(budgetTokensIC, "")

	return budgetTokensIS
}

// bindParamUpdateBudgetAlert binds the parameter recipient_channels
func (o *UpdateBudgetAlertParams) bindParamRecipientChannels(formats strfmt.Registry) []string {
	recipientChannelsIR := o.RecipientChannels

	var recipientChannelsIC []string
	for _, recipientChannelsIIR := range recipientChannelsIR { // explode []string

		recipientChannelsIIV := recipientChannelsIIR // string as string
		recipientChannelsIC = append(recipientChannelsIC, recipientChannelsIIV)
	}

	// items.CollectionFormat: ""
	recipientChannelsIS := swag.JoinByFormat(recipientChannelsIC, "")

	return recipientChannelsIS
}

// bindParamUpdateBudgetAlert binds the parameter user_tokens
func (o *UpdateBudgetAlertParams) bindParamUserTokens(formats strfmt.Registry) []string {
	userTokensIR := o.UserTokens

	var userTokensIC []string
	for _, userTokensIIR := range userTokensIR { // explode []string

		userTokensIIV := userTokensIIR // string as string
		userTokensIC = append(userTokensIC, userTokensIIV)
	}

	// items.CollectionFormat: ""
	userTokensIS := swag.JoinByFormat(userTokensIC, "")

	return userTokensIS
}
