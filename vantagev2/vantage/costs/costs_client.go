// Code generated by go-swagger; DO NOT EDIT.

package costs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new costs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new costs API client with basic auth credentials.
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

// New creates a new costs API client with a bearer token for authentication.
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
Client for costs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCostExport(params *CreateCostExportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCostExportAccepted, error)

	CreateCostReport(params *CreateCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCostReportCreated, error)

	DeleteCostReport(params *DeleteCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCostReportNoContent, error)

	GetCostReport(params *GetCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostReportOK, error)

	GetCostReports(params *GetCostReportsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostReportsOK, error)

	GetCosts(params *GetCostsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsOK, error)

	GetForecastedCosts(params *GetForecastedCostsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetForecastedCostsOK, error)

	UpdateCostReport(params *UpdateCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCostReportOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateCostExport Generate a DataExport of costs.
*/
func (a *Client) CreateCostExport(params *CreateCostExportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCostExportAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCostExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCostExport",
		Method:             "POST",
		PathPattern:        "/costs/data_exports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCostExportReader{formats: a.formats},
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
	success, ok := result.(*CreateCostExportAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCostExport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateCostReport Create a CostReport.
*/
func (a *Client) CreateCostReport(params *CreateCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCostReportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCostReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCostReport",
		Method:             "POST",
		PathPattern:        "/cost_reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCostReportReader{formats: a.formats},
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
	success, ok := result.(*CreateCostReportCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCostReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteCostReport Delete a CostReport.
*/
func (a *Client) DeleteCostReport(params *DeleteCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCostReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCostReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCostReport",
		Method:             "DELETE",
		PathPattern:        "/cost_reports/{cost_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCostReportReader{formats: a.formats},
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
	success, ok := result.(*DeleteCostReportNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCostReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCostReport Return a CostReport.
*/
func (a *Client) GetCostReport(params *GetCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCostReport",
		Method:             "GET",
		PathPattern:        "/cost_reports/{cost_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostReportReader{formats: a.formats},
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
	success, ok := result.(*GetCostReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCostReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCostReports Return all CostReports.
*/
func (a *Client) GetCostReports(params *GetCostReportsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostReportsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCostReports",
		Method:             "GET",
		PathPattern:        "/cost_reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostReportsReader{formats: a.formats},
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
	success, ok := result.(*GetCostReportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCostReports: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCosts Return all Costs for a CostReport or VQL filter.
*/
func (a *Client) GetCosts(params *GetCostsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCosts",
		Method:             "GET",
		PathPattern:        "/costs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostsReader{formats: a.formats},
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
	success, ok := result.(*GetCostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetForecastedCosts Return all ForecastedCosts.
*/
func (a *Client) GetForecastedCosts(params *GetForecastedCostsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetForecastedCostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetForecastedCostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getForecastedCosts",
		Method:             "GET",
		PathPattern:        "/cost_reports/{cost_report_token}/forecasted_costs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetForecastedCostsReader{formats: a.formats},
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
	success, ok := result.(*GetForecastedCostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getForecastedCosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCostReport Update a CostReport.
*/
func (a *Client) UpdateCostReport(params *UpdateCostReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCostReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCostReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCostReport",
		Method:             "PUT",
		PathPattern:        "/cost_reports/{cost_report_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCostReportReader{formats: a.formats},
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
	success, ok := result.(*UpdateCostReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCostReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
