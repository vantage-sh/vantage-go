// Code generated by go-swagger; DO NOT EDIT.

package business_metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new business metrics API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new business metrics API client with basic auth credentials.
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

// New creates a new business metrics API client with a bearer token for authentication.
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
Client for business metrics API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeApplicationxWwwFormUrlencoded sets the Content-Type header to "application/x-www-form-urlencoded".
func WithContentTypeApplicationxWwwFormUrlencoded(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
}

// WithContentTypeMultipartFormData sets the Content-Type header to "multipart/form-data".
func WithContentTypeMultipartFormData(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"multipart/form-data"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBusinessMetric(params *CreateBusinessMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBusinessMetricCreated, error)

	DeleteBusinessMetric(params *DeleteBusinessMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBusinessMetricNoContent, error)

	GetBusinessMetric(params *GetBusinessMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBusinessMetricOK, error)

	GetBusinessMetrics(params *GetBusinessMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBusinessMetricsOK, error)

	UpdateBusinessMetric(params *UpdateBusinessMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBusinessMetricOK, error)

	UpdateBusinessMetricValuesCSV(params *UpdateBusinessMetricValuesCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBusinessMetricValuesCSVCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateBusinessMetric Create a new BusinessMetric.
*/
func (a *Client) CreateBusinessMetric(params *CreateBusinessMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBusinessMetricCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBusinessMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBusinessMetric",
		Method:             "POST",
		PathPattern:        "/business_metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBusinessMetricReader{formats: a.formats},
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
	success, ok := result.(*CreateBusinessMetricCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBusinessMetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteBusinessMetric Deletes an existing BusinessMetric.
*/
func (a *Client) DeleteBusinessMetric(params *DeleteBusinessMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBusinessMetricNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBusinessMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteBusinessMetric",
		Method:             "DELETE",
		PathPattern:        "/business_metrics/{business_metric_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBusinessMetricReader{formats: a.formats},
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
	success, ok := result.(*DeleteBusinessMetricNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBusinessMetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBusinessMetric Return a specific BusinessMetrics.
*/
func (a *Client) GetBusinessMetric(params *GetBusinessMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBusinessMetricOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBusinessMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBusinessMetric",
		Method:             "GET",
		PathPattern:        "/business_metrics/{business_metric_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBusinessMetricReader{formats: a.formats},
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
	success, ok := result.(*GetBusinessMetricOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBusinessMetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBusinessMetrics Return all BusinessMetrics that the current API token has access to.
*/
func (a *Client) GetBusinessMetrics(params *GetBusinessMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBusinessMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBusinessMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBusinessMetrics",
		Method:             "GET",
		PathPattern:        "/business_metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBusinessMetricsReader{formats: a.formats},
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
	success, ok := result.(*GetBusinessMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBusinessMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateBusinessMetric Updates an existing BusinessMetric.
*/
func (a *Client) UpdateBusinessMetric(params *UpdateBusinessMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBusinessMetricOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBusinessMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBusinessMetric",
		Method:             "PUT",
		PathPattern:        "/business_metrics/{business_metric_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBusinessMetricReader{formats: a.formats},
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
	success, ok := result.(*UpdateBusinessMetricOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBusinessMetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateBusinessMetricValuesCSV Updates the values for an existing BusinessMetric from a CSV file.
*/
func (a *Client) UpdateBusinessMetricValuesCSV(params *UpdateBusinessMetricValuesCSVParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBusinessMetricValuesCSVCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBusinessMetricValuesCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBusinessMetricValuesCSV",
		Method:             "PUT",
		PathPattern:        "/business_metrics/{business_metric_token}/values.csv",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded", "multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBusinessMetricValuesCSVReader{formats: a.formats},
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
	success, ok := result.(*UpdateBusinessMetricValuesCSVCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBusinessMetricValuesCSV: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
