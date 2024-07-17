// Code generated by go-swagger; DO NOT EDIT.

package managed_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new managed accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new managed accounts API client with basic auth credentials.
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

// New creates a new managed accounts API client with a bearer token for authentication.
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
Client for managed accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateManagedAccount(params *CreateManagedAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateManagedAccountCreated, error)

	DeleteManagedAccount(params *DeleteManagedAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteManagedAccountNoContent, error)

	GetManagedAccount(params *GetManagedAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetManagedAccountOK, error)

	GetManagedAccounts(params *GetManagedAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetManagedAccountsOK, error)

	UpdateManagedAccount(params *UpdateManagedAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateManagedAccountOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateManagedAccount Create a Managed Account.
*/
func (a *Client) CreateManagedAccount(params *CreateManagedAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateManagedAccountCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateManagedAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createManagedAccount",
		Method:             "POST",
		PathPattern:        "/managed_accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateManagedAccountReader{formats: a.formats},
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
	success, ok := result.(*CreateManagedAccountCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createManagedAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteManagedAccount Delete a Managed Account.
*/
func (a *Client) DeleteManagedAccount(params *DeleteManagedAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteManagedAccountNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteManagedAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteManagedAccount",
		Method:             "DELETE",
		PathPattern:        "/managed_accounts/{managed_account_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteManagedAccountReader{formats: a.formats},
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
	success, ok := result.(*DeleteManagedAccountNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteManagedAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetManagedAccount Return a Managed Account.
*/
func (a *Client) GetManagedAccount(params *GetManagedAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetManagedAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagedAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getManagedAccount",
		Method:             "GET",
		PathPattern:        "/managed_accounts/{managed_account_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetManagedAccountReader{formats: a.formats},
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
	success, ok := result.(*GetManagedAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getManagedAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetManagedAccounts Returns a list of managed accounts.
*/
func (a *Client) GetManagedAccounts(params *GetManagedAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetManagedAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagedAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getManagedAccounts",
		Method:             "GET",
		PathPattern:        "/managed_accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetManagedAccountsReader{formats: a.formats},
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
	success, ok := result.(*GetManagedAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getManagedAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateManagedAccount Update a Managed Account.
*/
func (a *Client) UpdateManagedAccount(params *UpdateManagedAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateManagedAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateManagedAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateManagedAccount",
		Method:             "PUT",
		PathPattern:        "/managed_accounts/{managed_account_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateManagedAccountReader{formats: a.formats},
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
	success, ok := result.(*UpdateManagedAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateManagedAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
