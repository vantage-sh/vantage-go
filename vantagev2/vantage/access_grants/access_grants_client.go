// Code generated by go-swagger; DO NOT EDIT.

package access_grants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new access grants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access grants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAccessGrant(params *CreateAccessGrantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccessGrantCreated, error)

	DeleteAccessGrant(params *DeleteAccessGrantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAccessGrantNoContent, error)

	GetAccessGrant(params *GetAccessGrantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccessGrantOK, error)

	GetAccessGrants(params *GetAccessGrantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccessGrantsOK, error)

	UpdateAccessGrant(params *UpdateAccessGrantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccessGrantOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAccessGrant Create an Access Grant.
*/
func (a *Client) CreateAccessGrant(params *CreateAccessGrantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccessGrantCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessGrantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAccessGrant",
		Method:             "POST",
		PathPattern:        "/access_grants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAccessGrantReader{formats: a.formats},
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
	success, ok := result.(*CreateAccessGrantCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccessGrant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAccessGrant Delete an Access Grant.
*/
func (a *Client) DeleteAccessGrant(params *DeleteAccessGrantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAccessGrantNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccessGrantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAccessGrant",
		Method:             "DELETE",
		PathPattern:        "/access_grants/{resource_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccessGrantReader{formats: a.formats},
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
	success, ok := result.(*DeleteAccessGrantNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAccessGrant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessGrant Return a specific Access Grant.
*/
func (a *Client) GetAccessGrant(params *GetAccessGrantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccessGrantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessGrantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccessGrant",
		Method:             "GET",
		PathPattern:        "/access_grants/{access_grant_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessGrantReader{formats: a.formats},
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
	success, ok := result.(*GetAccessGrantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccessGrant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessGrants Return all Access Grants that the current API token has access to.
*/
func (a *Client) GetAccessGrants(params *GetAccessGrantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccessGrantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessGrantsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccessGrants",
		Method:             "GET",
		PathPattern:        "/access_grants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessGrantsReader{formats: a.formats},
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
	success, ok := result.(*GetAccessGrantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccessGrants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAccessGrant Update an AccessGrant.
*/
func (a *Client) UpdateAccessGrant(params *UpdateAccessGrantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAccessGrantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccessGrantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAccessGrant",
		Method:             "PUT",
		PathPattern:        "/access_grants/{resource_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAccessGrantReader{formats: a.formats},
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
	success, ok := result.(*UpdateAccessGrantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAccessGrant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
