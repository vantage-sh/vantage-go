// Code generated by go-swagger; DO NOT EDIT.

package financial_commitment_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new financial commitment reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new financial commitment reports API client with basic auth credentials.
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

// New creates a new financial commitment reports API client with a bearer token for authentication.
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
Client for financial commitment reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateFinancialCommitmentReport(params *CreateFinancialCommitmentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFinancialCommitmentReportCreated, error)

	DeleteFinancialCommitmentReport(params *DeleteFinancialCommitmentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFinancialCommitmentReportNoContent, error)

	GetFinancialCommitmentReport(params *GetFinancialCommitmentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFinancialCommitmentReportOK, error)

	GetFinancialCommitmentReports(params *GetFinancialCommitmentReportsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFinancialCommitmentReportsOK, error)

	UpdateFinancialCommitmentReport(params *UpdateFinancialCommitmentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFinancialCommitmentReportOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateFinancialCommitmentReport Create a FinancialCommitmentReport.
*/
func (a *Client) CreateFinancialCommitmentReport(params *CreateFinancialCommitmentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFinancialCommitmentReportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFinancialCommitmentReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFinancialCommitmentReport",
		Method:             "POST",
		PathPattern:        "/financial_commitment_reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFinancialCommitmentReportReader{formats: a.formats},
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
	success, ok := result.(*CreateFinancialCommitmentReportCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFinancialCommitmentReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteFinancialCommitmentReport Delete a FinancialCommitmentReport.
*/
func (a *Client) DeleteFinancialCommitmentReport(params *DeleteFinancialCommitmentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFinancialCommitmentReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFinancialCommitmentReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteFinancialCommitmentReport",
		Method:             "DELETE",
		PathPattern:        "/financial_commitment_reports/{financial_commitment_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFinancialCommitmentReportReader{formats: a.formats},
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
	success, ok := result.(*DeleteFinancialCommitmentReportNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteFinancialCommitmentReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFinancialCommitmentReport Return a FinancialCommitmentReport.
*/
func (a *Client) GetFinancialCommitmentReport(params *GetFinancialCommitmentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFinancialCommitmentReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFinancialCommitmentReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFinancialCommitmentReport",
		Method:             "GET",
		PathPattern:        "/financial_commitment_reports/{financial_commitment_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFinancialCommitmentReportReader{formats: a.formats},
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
	success, ok := result.(*GetFinancialCommitmentReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFinancialCommitmentReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFinancialCommitmentReports Return all FinancialCommitmentReports.
*/
func (a *Client) GetFinancialCommitmentReports(params *GetFinancialCommitmentReportsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFinancialCommitmentReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFinancialCommitmentReportsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFinancialCommitmentReports",
		Method:             "GET",
		PathPattern:        "/financial_commitment_reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFinancialCommitmentReportsReader{formats: a.formats},
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
	success, ok := result.(*GetFinancialCommitmentReportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFinancialCommitmentReports: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateFinancialCommitmentReport Update a FinancialCommitmentReport.
*/
func (a *Client) UpdateFinancialCommitmentReport(params *UpdateFinancialCommitmentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFinancialCommitmentReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFinancialCommitmentReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateFinancialCommitmentReport",
		Method:             "PUT",
		PathPattern:        "/financial_commitment_reports/{financial_commitment_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFinancialCommitmentReportReader{formats: a.formats},
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
	success, ok := result.(*UpdateFinancialCommitmentReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFinancialCommitmentReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
