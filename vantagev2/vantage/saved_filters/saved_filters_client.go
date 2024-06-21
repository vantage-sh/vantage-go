// Code generated by go-swagger; DO NOT EDIT.

package saved_filters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new saved filters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new saved filters API client with basic auth credentials.
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

// New creates a new saved filters API client with a bearer token for authentication.
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
Client for saved filters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSavedFilter(params *CreateSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSavedFilterCreated, error)

	DeleteSavedFilter(params *DeleteSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSavedFilterNoContent, error)

	GetSavedFilter(params *GetSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSavedFilterOK, error)

	GetSavedFilters(params *GetSavedFiltersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSavedFiltersOK, error)

	UpdateSavedFilter(params *UpdateSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSavedFilterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSavedFilter Create a SavedFilter for CostReports.
*/
func (a *Client) CreateSavedFilter(params *CreateSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSavedFilterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSavedFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSavedFilter",
		Method:             "POST",
		PathPattern:        "/saved_filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSavedFilterReader{formats: a.formats},
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
	success, ok := result.(*CreateSavedFilterCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSavedFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteSavedFilter Delete a SavedFilter for CostReports.
*/
func (a *Client) DeleteSavedFilter(params *DeleteSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSavedFilterNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSavedFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSavedFilter",
		Method:             "DELETE",
		PathPattern:        "/saved_filters/{saved_filter_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSavedFilterReader{formats: a.formats},
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
	success, ok := result.(*DeleteSavedFilterNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSavedFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSavedFilter Return a specific SavedFilter.
*/
func (a *Client) GetSavedFilter(params *GetSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSavedFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSavedFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSavedFilter",
		Method:             "GET",
		PathPattern:        "/saved_filters/{saved_filter_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSavedFilterReader{formats: a.formats},
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
	success, ok := result.(*GetSavedFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSavedFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSavedFilters Return all SavedFilters that can be applied to a CostReport.
*/
func (a *Client) GetSavedFilters(params *GetSavedFiltersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSavedFiltersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSavedFiltersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSavedFilters",
		Method:             "GET",
		PathPattern:        "/saved_filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSavedFiltersReader{formats: a.formats},
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
	success, ok := result.(*GetSavedFiltersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSavedFilters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSavedFilter Update a SavedFilter for CostReports.
*/
func (a *Client) UpdateSavedFilter(params *UpdateSavedFilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSavedFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSavedFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSavedFilter",
		Method:             "PUT",
		PathPattern:        "/saved_filters/{saved_filter_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSavedFilterReader{formats: a.formats},
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
	success, ok := result.(*UpdateSavedFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSavedFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
