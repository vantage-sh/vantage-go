// Code generated by go-swagger; DO NOT EDIT.

package folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new folders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new folders API client with basic auth credentials.
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

// New creates a new folders API client with a bearer token for authentication.
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
Client for folders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateFolder(params *CreateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFolderCreated, error)

	DeleteFolder(params *DeleteFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFolderNoContent, error)

	GetFolder(params *GetFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderOK, error)

	GetFolders(params *GetFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFoldersOK, error)

	UpdateFolder(params *UpdateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFolderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateFolder Create a Folder for CostReports.
*/
func (a *Client) CreateFolder(params *CreateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFolderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFolder",
		Method:             "POST",
		PathPattern:        "/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFolderReader{formats: a.formats},
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
	success, ok := result.(*CreateFolderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteFolder Delete a Folder for CostReports.
*/
func (a *Client) DeleteFolder(params *DeleteFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFolderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteFolder",
		Method:             "DELETE",
		PathPattern:        "/folders/{folder_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFolderReader{formats: a.formats},
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
	success, ok := result.(*DeleteFolderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFolder Return a specific Folder for CostReports.
*/
func (a *Client) GetFolder(params *GetFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolder",
		Method:             "GET",
		PathPattern:        "/folders/{folder_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFolderReader{formats: a.formats},
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
	success, ok := result.(*GetFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFolders Return all Folders for CostReports.
*/
func (a *Client) GetFolders(params *GetFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoldersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolders",
		Method:             "GET",
		PathPattern:        "/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoldersReader{formats: a.formats},
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
	success, ok := result.(*GetFoldersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateFolder Update a Folder for CostReports.
*/
func (a *Client) UpdateFolder(params *UpdateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateFolder",
		Method:             "PUT",
		PathPattern:        "/folders/{folder_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFolderReader{formats: a.formats},
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
	success, ok := result.(*UpdateFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
