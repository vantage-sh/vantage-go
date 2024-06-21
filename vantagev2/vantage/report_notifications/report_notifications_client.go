// Code generated by go-swagger; DO NOT EDIT.

package report_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new report notifications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new report notifications API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new report notifications API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for report notifications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateReportNotification(params *CreateReportNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateReportNotificationCreated, error)

	DeleteReportNotification(params *DeleteReportNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteReportNotificationNoContent, error)

	GetReportNotification(params *GetReportNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportNotificationOK, error)

	GetReportNotifications(params *GetReportNotificationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportNotificationsOK, error)

	UpdateReportNotification(params *UpdateReportNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateReportNotificationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateReportNotification Create a ReportNotification.
*/
func (a *Client) CreateReportNotification(params *CreateReportNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateReportNotificationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateReportNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createReportNotification",
		Method:             "POST",
		PathPattern:        "/report_notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateReportNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateReportNotificationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createReportNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteReportNotification Delete a ReportNotification.
*/
func (a *Client) DeleteReportNotification(params *DeleteReportNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteReportNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReportNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteReportNotification",
		Method:             "DELETE",
		PathPattern:        "/report_notifications/{report_notification_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteReportNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteReportNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteReportNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetReportNotification Return a ReportNotification.
*/
func (a *Client) GetReportNotification(params *GetReportNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReportNotification",
		Method:             "GET",
		PathPattern:        "/report_notifications/{report_notification_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReportNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReportNotificationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReportNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetReportNotifications Return all ReportNotifications.
*/
func (a *Client) GetReportNotifications(params *GetReportNotificationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportNotificationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReportNotifications",
		Method:             "GET",
		PathPattern:        "/report_notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReportNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReportNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReportNotifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateReportNotification Update a ReportNotification.
*/
func (a *Client) UpdateReportNotification(params *UpdateReportNotificationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateReportNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateReportNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateReportNotification",
		Method:             "PUT",
		PathPattern:        "/report_notifications/{report_notification_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateReportNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateReportNotificationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateReportNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
