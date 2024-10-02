// Code generated by go-swagger; DO NOT EDIT.

package resource_reports

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

// CreateResourceReportReader is a Reader for the CreateResourceReport structure.
type CreateResourceReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResourceReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateResourceReportCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateResourceReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateResourceReportUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /resource_reports] createResourceReport", response, response.Code())
	}
}

// NewCreateResourceReportCreated creates a CreateResourceReportCreated with default headers values
func NewCreateResourceReportCreated() *CreateResourceReportCreated {
	return &CreateResourceReportCreated{}
}

/*
CreateResourceReportCreated describes a response with status code 201, with default header values.

CreateResourceReportCreated create resource report created
*/
type CreateResourceReportCreated struct {
	Payload *models.ResourceReport
}

// IsSuccess returns true when this create resource report created response has a 2xx status code
func (o *CreateResourceReportCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create resource report created response has a 3xx status code
func (o *CreateResourceReportCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource report created response has a 4xx status code
func (o *CreateResourceReportCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create resource report created response has a 5xx status code
func (o *CreateResourceReportCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource report created response a status code equal to that given
func (o *CreateResourceReportCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create resource report created response
func (o *CreateResourceReportCreated) Code() int {
	return 201
}

func (o *CreateResourceReportCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource_reports][%d] createResourceReportCreated %s", 201, payload)
}

func (o *CreateResourceReportCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource_reports][%d] createResourceReportCreated %s", 201, payload)
}

func (o *CreateResourceReportCreated) GetPayload() *models.ResourceReport {
	return o.Payload
}

func (o *CreateResourceReportCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceReportBadRequest creates a CreateResourceReportBadRequest with default headers values
func NewCreateResourceReportBadRequest() *CreateResourceReportBadRequest {
	return &CreateResourceReportBadRequest{}
}

/*
CreateResourceReportBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateResourceReportBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create resource report bad request response has a 2xx status code
func (o *CreateResourceReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create resource report bad request response has a 3xx status code
func (o *CreateResourceReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource report bad request response has a 4xx status code
func (o *CreateResourceReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create resource report bad request response has a 5xx status code
func (o *CreateResourceReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource report bad request response a status code equal to that given
func (o *CreateResourceReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create resource report bad request response
func (o *CreateResourceReportBadRequest) Code() int {
	return 400
}

func (o *CreateResourceReportBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource_reports][%d] createResourceReportBadRequest %s", 400, payload)
}

func (o *CreateResourceReportBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource_reports][%d] createResourceReportBadRequest %s", 400, payload)
}

func (o *CreateResourceReportBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateResourceReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceReportUnprocessableEntity creates a CreateResourceReportUnprocessableEntity with default headers values
func NewCreateResourceReportUnprocessableEntity() *CreateResourceReportUnprocessableEntity {
	return &CreateResourceReportUnprocessableEntity{}
}

/*
CreateResourceReportUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type CreateResourceReportUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create resource report unprocessable entity response has a 2xx status code
func (o *CreateResourceReportUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create resource report unprocessable entity response has a 3xx status code
func (o *CreateResourceReportUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource report unprocessable entity response has a 4xx status code
func (o *CreateResourceReportUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create resource report unprocessable entity response has a 5xx status code
func (o *CreateResourceReportUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource report unprocessable entity response a status code equal to that given
func (o *CreateResourceReportUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create resource report unprocessable entity response
func (o *CreateResourceReportUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateResourceReportUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource_reports][%d] createResourceReportUnprocessableEntity %s", 422, payload)
}

func (o *CreateResourceReportUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource_reports][%d] createResourceReportUnprocessableEntity %s", 422, payload)
}

func (o *CreateResourceReportUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateResourceReportUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
