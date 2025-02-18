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

// UpdateNetworkFlowReportReader is a Reader for the UpdateNetworkFlowReport structure.
type UpdateNetworkFlowReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkFlowReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkFlowReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNetworkFlowReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateNetworkFlowReportUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /network_flow_reports/{network_flow_report_token}] updateNetworkFlowReport", response, response.Code())
	}
}

// NewUpdateNetworkFlowReportOK creates a UpdateNetworkFlowReportOK with default headers values
func NewUpdateNetworkFlowReportOK() *UpdateNetworkFlowReportOK {
	return &UpdateNetworkFlowReportOK{}
}

/*
UpdateNetworkFlowReportOK describes a response with status code 200, with default header values.

UpdateNetworkFlowReportOK update network flow report o k
*/
type UpdateNetworkFlowReportOK struct {
	Payload *models.NetworkFlowReport
}

// IsSuccess returns true when this update network flow report o k response has a 2xx status code
func (o *UpdateNetworkFlowReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network flow report o k response has a 3xx status code
func (o *UpdateNetworkFlowReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network flow report o k response has a 4xx status code
func (o *UpdateNetworkFlowReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network flow report o k response has a 5xx status code
func (o *UpdateNetworkFlowReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network flow report o k response a status code equal to that given
func (o *UpdateNetworkFlowReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network flow report o k response
func (o *UpdateNetworkFlowReportOK) Code() int {
	return 200
}

func (o *UpdateNetworkFlowReportOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network_flow_reports/{network_flow_report_token}][%d] updateNetworkFlowReportOK %s", 200, payload)
}

func (o *UpdateNetworkFlowReportOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network_flow_reports/{network_flow_report_token}][%d] updateNetworkFlowReportOK %s", 200, payload)
}

func (o *UpdateNetworkFlowReportOK) GetPayload() *models.NetworkFlowReport {
	return o.Payload
}

func (o *UpdateNetworkFlowReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkFlowReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkFlowReportBadRequest creates a UpdateNetworkFlowReportBadRequest with default headers values
func NewUpdateNetworkFlowReportBadRequest() *UpdateNetworkFlowReportBadRequest {
	return &UpdateNetworkFlowReportBadRequest{}
}

/*
UpdateNetworkFlowReportBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateNetworkFlowReportBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update network flow report bad request response has a 2xx status code
func (o *UpdateNetworkFlowReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update network flow report bad request response has a 3xx status code
func (o *UpdateNetworkFlowReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network flow report bad request response has a 4xx status code
func (o *UpdateNetworkFlowReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update network flow report bad request response has a 5xx status code
func (o *UpdateNetworkFlowReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update network flow report bad request response a status code equal to that given
func (o *UpdateNetworkFlowReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update network flow report bad request response
func (o *UpdateNetworkFlowReportBadRequest) Code() int {
	return 400
}

func (o *UpdateNetworkFlowReportBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network_flow_reports/{network_flow_report_token}][%d] updateNetworkFlowReportBadRequest %s", 400, payload)
}

func (o *UpdateNetworkFlowReportBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network_flow_reports/{network_flow_report_token}][%d] updateNetworkFlowReportBadRequest %s", 400, payload)
}

func (o *UpdateNetworkFlowReportBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateNetworkFlowReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkFlowReportUnprocessableEntity creates a UpdateNetworkFlowReportUnprocessableEntity with default headers values
func NewUpdateNetworkFlowReportUnprocessableEntity() *UpdateNetworkFlowReportUnprocessableEntity {
	return &UpdateNetworkFlowReportUnprocessableEntity{}
}

/*
UpdateNetworkFlowReportUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type UpdateNetworkFlowReportUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update network flow report unprocessable entity response has a 2xx status code
func (o *UpdateNetworkFlowReportUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update network flow report unprocessable entity response has a 3xx status code
func (o *UpdateNetworkFlowReportUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network flow report unprocessable entity response has a 4xx status code
func (o *UpdateNetworkFlowReportUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update network flow report unprocessable entity response has a 5xx status code
func (o *UpdateNetworkFlowReportUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update network flow report unprocessable entity response a status code equal to that given
func (o *UpdateNetworkFlowReportUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update network flow report unprocessable entity response
func (o *UpdateNetworkFlowReportUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateNetworkFlowReportUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network_flow_reports/{network_flow_report_token}][%d] updateNetworkFlowReportUnprocessableEntity %s", 422, payload)
}

func (o *UpdateNetworkFlowReportUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network_flow_reports/{network_flow_report_token}][%d] updateNetworkFlowReportUnprocessableEntity %s", 422, payload)
}

func (o *UpdateNetworkFlowReportUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateNetworkFlowReportUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
