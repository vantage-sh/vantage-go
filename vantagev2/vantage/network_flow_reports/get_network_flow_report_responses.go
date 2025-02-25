// Code generated by go-swagger; DO NOT EDIT.

package network_flow_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// GetNetworkFlowReportReader is a Reader for the GetNetworkFlowReport structure.
type GetNetworkFlowReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkFlowReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkFlowReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNetworkFlowReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /network_flow_reports/{network_flow_report_token}] getNetworkFlowReport", response, response.Code())
	}
}

// NewGetNetworkFlowReportOK creates a GetNetworkFlowReportOK with default headers values
func NewGetNetworkFlowReportOK() *GetNetworkFlowReportOK {
	return &GetNetworkFlowReportOK{}
}

/*
GetNetworkFlowReportOK describes a response with status code 200, with default header values.

GetNetworkFlowReportOK get network flow report o k
*/
type GetNetworkFlowReportOK struct {
	Payload *models.NetworkFlowReport
}

// IsSuccess returns true when this get network flow report o k response has a 2xx status code
func (o *GetNetworkFlowReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network flow report o k response has a 3xx status code
func (o *GetNetworkFlowReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network flow report o k response has a 4xx status code
func (o *GetNetworkFlowReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network flow report o k response has a 5xx status code
func (o *GetNetworkFlowReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network flow report o k response a status code equal to that given
func (o *GetNetworkFlowReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network flow report o k response
func (o *GetNetworkFlowReportOK) Code() int {
	return 200
}

func (o *GetNetworkFlowReportOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network_flow_reports/{network_flow_report_token}][%d] getNetworkFlowReportOK %s", 200, payload)
}

func (o *GetNetworkFlowReportOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network_flow_reports/{network_flow_report_token}][%d] getNetworkFlowReportOK %s", 200, payload)
}

func (o *GetNetworkFlowReportOK) GetPayload() *models.NetworkFlowReport {
	return o.Payload
}

func (o *GetNetworkFlowReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkFlowReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkFlowReportNotFound creates a GetNetworkFlowReportNotFound with default headers values
func NewGetNetworkFlowReportNotFound() *GetNetworkFlowReportNotFound {
	return &GetNetworkFlowReportNotFound{}
}

/*
GetNetworkFlowReportNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetNetworkFlowReportNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this get network flow report not found response has a 2xx status code
func (o *GetNetworkFlowReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network flow report not found response has a 3xx status code
func (o *GetNetworkFlowReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network flow report not found response has a 4xx status code
func (o *GetNetworkFlowReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network flow report not found response has a 5xx status code
func (o *GetNetworkFlowReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get network flow report not found response a status code equal to that given
func (o *GetNetworkFlowReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get network flow report not found response
func (o *GetNetworkFlowReportNotFound) Code() int {
	return 404
}

func (o *GetNetworkFlowReportNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network_flow_reports/{network_flow_report_token}][%d] getNetworkFlowReportNotFound %s", 404, payload)
}

func (o *GetNetworkFlowReportNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network_flow_reports/{network_flow_report_token}][%d] getNetworkFlowReportNotFound %s", 404, payload)
}

func (o *GetNetworkFlowReportNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetNetworkFlowReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
