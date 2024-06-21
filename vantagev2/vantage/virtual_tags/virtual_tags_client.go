// Code generated by go-swagger; DO NOT EDIT.

package virtual_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new virtual tags API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new virtual tags API client with basic auth credentials.
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

// New creates a new virtual tags API client with a bearer token for authentication.
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
Client for virtual tags API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVirtualTagConfig(params *CreateVirtualTagConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVirtualTagConfigCreated, error)

	DeleteVirtualTagConfig(params *DeleteVirtualTagConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVirtualTagConfigNoContent, error)

	GetVirtualTagConfig(params *GetVirtualTagConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVirtualTagConfigOK, error)

	GetVirtualTagConfigs(params *GetVirtualTagConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVirtualTagConfigsOK, error)

	UpdateVirtualTagConfig(params *UpdateVirtualTagConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVirtualTagConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateVirtualTagConfig Create a new VirtualTagConfig.
*/
func (a *Client) CreateVirtualTagConfig(params *CreateVirtualTagConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVirtualTagConfigCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVirtualTagConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createVirtualTagConfig",
		Method:             "POST",
		PathPattern:        "/virtual_tag_configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateVirtualTagConfigReader{formats: a.formats},
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
	success, ok := result.(*CreateVirtualTagConfigCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createVirtualTagConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteVirtualTagConfig Deletes an existing VirtualTagConfig.
*/
func (a *Client) DeleteVirtualTagConfig(params *DeleteVirtualTagConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVirtualTagConfigNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVirtualTagConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVirtualTagConfig",
		Method:             "DELETE",
		PathPattern:        "/virtual_tag_configs/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVirtualTagConfigReader{formats: a.formats},
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
	success, ok := result.(*DeleteVirtualTagConfigNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVirtualTagConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVirtualTagConfig Return a specific VirtualTagConfig.
*/
func (a *Client) GetVirtualTagConfig(params *GetVirtualTagConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVirtualTagConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVirtualTagConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVirtualTagConfig",
		Method:             "GET",
		PathPattern:        "/virtual_tag_configs/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVirtualTagConfigReader{formats: a.formats},
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
	success, ok := result.(*GetVirtualTagConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVirtualTagConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVirtualTagConfigs Return all VirtualTagConfigs that the current API token has access to.
*/
func (a *Client) GetVirtualTagConfigs(params *GetVirtualTagConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVirtualTagConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVirtualTagConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVirtualTagConfigs",
		Method:             "GET",
		PathPattern:        "/virtual_tag_configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVirtualTagConfigsReader{formats: a.formats},
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
	success, ok := result.(*GetVirtualTagConfigsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVirtualTagConfigs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateVirtualTagConfig Updates an existing VirtualTagConfig.
*/
func (a *Client) UpdateVirtualTagConfig(params *UpdateVirtualTagConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVirtualTagConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVirtualTagConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateVirtualTagConfig",
		Method:             "PUT",
		PathPattern:        "/virtual_tag_configs/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVirtualTagConfigReader{formats: a.formats},
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
	success, ok := result.(*UpdateVirtualTagConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVirtualTagConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
