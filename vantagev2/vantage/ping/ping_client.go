// Code generated by go-swagger; DO NOT EDIT.

package ping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ping API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ping API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Ping(params *PingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Ping This is a health check endpoint that can be used to determine Vantage API healthiness. It will return 200 if everything is running smoothly.
*/
func (a *Client) Ping(params *PingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ping",
		Method:             "GET",
		PathPattern:        "/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PingReader{formats: a.formats},
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
	success, ok := result.(*PingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
