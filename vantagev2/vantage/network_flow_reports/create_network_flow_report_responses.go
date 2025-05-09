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

// CreateNetworkFlowReportReader is a Reader for the CreateNetworkFlowReport structure.
type CreateNetworkFlowReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkFlowReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkFlowReportCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNetworkFlowReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateNetworkFlowReportUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /network_flow_reports] createNetworkFlowReport", response, response.Code())
	}
}

// NewCreateNetworkFlowReportCreated creates a CreateNetworkFlowReportCreated with default headers values
func NewCreateNetworkFlowReportCreated() *CreateNetworkFlowReportCreated {
	return &CreateNetworkFlowReportCreated{}
}

/*
CreateNetworkFlowReportCreated describes a response with status code 201, with default header values.

CreateNetworkFlowReportCreated create network flow report created
*/
type CreateNetworkFlowReportCreated struct {
	Payload *models.NetworkFlowReport
}

// IsSuccess returns true when this create network flow report created response has a 2xx status code
func (o *CreateNetworkFlowReportCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network flow report created response has a 3xx status code
func (o *CreateNetworkFlowReportCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network flow report created response has a 4xx status code
func (o *CreateNetworkFlowReportCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network flow report created response has a 5xx status code
func (o *CreateNetworkFlowReportCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network flow report created response a status code equal to that given
func (o *CreateNetworkFlowReportCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network flow report created response
func (o *CreateNetworkFlowReportCreated) Code() int {
	return 201
}

func (o *CreateNetworkFlowReportCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network_flow_reports][%d] createNetworkFlowReportCreated %s", 201, payload)
}

func (o *CreateNetworkFlowReportCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network_flow_reports][%d] createNetworkFlowReportCreated %s", 201, payload)
}

func (o *CreateNetworkFlowReportCreated) GetPayload() *models.NetworkFlowReport {
	return o.Payload
}

func (o *CreateNetworkFlowReportCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkFlowReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNetworkFlowReportBadRequest creates a CreateNetworkFlowReportBadRequest with default headers values
func NewCreateNetworkFlowReportBadRequest() *CreateNetworkFlowReportBadRequest {
	return &CreateNetworkFlowReportBadRequest{}
}

/*
CreateNetworkFlowReportBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateNetworkFlowReportBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create network flow report bad request response has a 2xx status code
func (o *CreateNetworkFlowReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create network flow report bad request response has a 3xx status code
func (o *CreateNetworkFlowReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network flow report bad request response has a 4xx status code
func (o *CreateNetworkFlowReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create network flow report bad request response has a 5xx status code
func (o *CreateNetworkFlowReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create network flow report bad request response a status code equal to that given
func (o *CreateNetworkFlowReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create network flow report bad request response
func (o *CreateNetworkFlowReportBadRequest) Code() int {
	return 400
}

func (o *CreateNetworkFlowReportBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network_flow_reports][%d] createNetworkFlowReportBadRequest %s", 400, payload)
}

func (o *CreateNetworkFlowReportBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network_flow_reports][%d] createNetworkFlowReportBadRequest %s", 400, payload)
}

func (o *CreateNetworkFlowReportBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateNetworkFlowReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNetworkFlowReportUnprocessableEntity creates a CreateNetworkFlowReportUnprocessableEntity with default headers values
func NewCreateNetworkFlowReportUnprocessableEntity() *CreateNetworkFlowReportUnprocessableEntity {
	return &CreateNetworkFlowReportUnprocessableEntity{}
}

/*
CreateNetworkFlowReportUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type CreateNetworkFlowReportUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create network flow report unprocessable entity response has a 2xx status code
func (o *CreateNetworkFlowReportUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create network flow report unprocessable entity response has a 3xx status code
func (o *CreateNetworkFlowReportUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network flow report unprocessable entity response has a 4xx status code
func (o *CreateNetworkFlowReportUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create network flow report unprocessable entity response has a 5xx status code
func (o *CreateNetworkFlowReportUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create network flow report unprocessable entity response a status code equal to that given
func (o *CreateNetworkFlowReportUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create network flow report unprocessable entity response
func (o *CreateNetworkFlowReportUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateNetworkFlowReportUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network_flow_reports][%d] createNetworkFlowReportUnprocessableEntity %s", 422, payload)
}

func (o *CreateNetworkFlowReportUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network_flow_reports][%d] createNetworkFlowReportUnprocessableEntity %s", 422, payload)
}

func (o *CreateNetworkFlowReportUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateNetworkFlowReportUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
