// Code generated by go-swagger; DO NOT EDIT.

package prices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new prices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for prices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPrice(params *GetPriceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPriceOK, error)

	GetPrices(params *GetPricesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPricesOK, error)

	GetProduct(params *GetProductParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductOK, error)

	GetProducts(params *GetProductsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsOK, error)

	GetProviders(params *GetProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProvidersOK, error)

	GetServices(params *GetServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServicesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetPrice Returns a price
*/
func (a *Client) GetPrice(params *GetPriceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPriceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPriceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPrice",
		Method:             "GET",
		PathPattern:        "/products/{product_id}/prices/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPriceReader{formats: a.formats},
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
	success, ok := result.(*GetPriceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPrice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrices Return available Prices across all Regions for a Product.
*/
func (a *Client) GetPrices(params *GetPricesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPricesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPrices",
		Method:             "GET",
		PathPattern:        "/products/{product_id}/prices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPricesReader{formats: a.formats},
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
	success, ok := result.(*GetPricesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPrices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetProduct Return a product
*/
func (a *Client) GetProduct(params *GetProductParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getProduct",
		Method:             "GET",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductReader{formats: a.formats},
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
	success, ok := result.(*GetProductOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetProducts Return available Products for a Service. For example, with a Provider of AWS and a Service of EC2, Products will be a list of all EC2 Instances. By default, this endpoint returns all Products across all Services and Providers but has optional query parameters for filtering listed below.
*/
func (a *Client) GetProducts(params *GetProductsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getProducts",
		Method:             "GET",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsReader{formats: a.formats},
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
	success, ok := result.(*GetProductsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProducts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetProviders Providers are cloud infrastructure and service providers from which all cloud prices are derived. You can think of example Providers as being AWS, GCP, Cloudflare or Datadog. Currently, Vantage only supports a single provider of AWS but over time more will be added. Use this endpoint to retrieve a provider id for other API calls.
*/
func (a *Client) GetProviders(params *GetProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProvidersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getProviders",
		Method:             "GET",
		PathPattern:        "/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProvidersReader{formats: a.formats},
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
	success, ok := result.(*GetProvidersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProviders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetServices Return all Services. Examples of Services are EC2 for AWS. This endpoint will return all Services by default but you have the ability to filter Services by Provider using the optional query parameter documented below.
*/
func (a *Client) GetServices(params *GetServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServices",
		Method:             "GET",
		PathPattern:        "/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServicesReader{formats: a.formats},
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
	success, ok := result.(*GetServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}