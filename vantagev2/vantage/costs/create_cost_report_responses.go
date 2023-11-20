// Code generated by go-swagger; DO NOT EDIT.

package costs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// CreateCostReportReader is a Reader for the CreateCostReport structure.
type CreateCostReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCostReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCostReportCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCostReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateCostReportUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cost_reports] createCostReport", response, response.Code())
	}
}

// NewCreateCostReportCreated creates a CreateCostReportCreated with default headers values
func NewCreateCostReportCreated() *CreateCostReportCreated {
	return &CreateCostReportCreated{}
}

/*
CreateCostReportCreated describes a response with status code 201, with default header values.

CreateCostReportCreated create cost report created
*/
type CreateCostReportCreated struct {
	Payload *models.CostReport
}

// IsSuccess returns true when this create cost report created response has a 2xx status code
func (o *CreateCostReportCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cost report created response has a 3xx status code
func (o *CreateCostReportCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cost report created response has a 4xx status code
func (o *CreateCostReportCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cost report created response has a 5xx status code
func (o *CreateCostReportCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create cost report created response a status code equal to that given
func (o *CreateCostReportCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create cost report created response
func (o *CreateCostReportCreated) Code() int {
	return 201
}

func (o *CreateCostReportCreated) Error() string {
	return fmt.Sprintf("[POST /cost_reports][%d] createCostReportCreated  %+v", 201, o.Payload)
}

func (o *CreateCostReportCreated) String() string {
	return fmt.Sprintf("[POST /cost_reports][%d] createCostReportCreated  %+v", 201, o.Payload)
}

func (o *CreateCostReportCreated) GetPayload() *models.CostReport {
	return o.Payload
}

func (o *CreateCostReportCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CostReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCostReportBadRequest creates a CreateCostReportBadRequest with default headers values
func NewCreateCostReportBadRequest() *CreateCostReportBadRequest {
	return &CreateCostReportBadRequest{}
}

/*
CreateCostReportBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateCostReportBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create cost report bad request response has a 2xx status code
func (o *CreateCostReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cost report bad request response has a 3xx status code
func (o *CreateCostReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cost report bad request response has a 4xx status code
func (o *CreateCostReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cost report bad request response has a 5xx status code
func (o *CreateCostReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create cost report bad request response a status code equal to that given
func (o *CreateCostReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create cost report bad request response
func (o *CreateCostReportBadRequest) Code() int {
	return 400
}

func (o *CreateCostReportBadRequest) Error() string {
	return fmt.Sprintf("[POST /cost_reports][%d] createCostReportBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCostReportBadRequest) String() string {
	return fmt.Sprintf("[POST /cost_reports][%d] createCostReportBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCostReportBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateCostReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCostReportUnprocessableEntity creates a CreateCostReportUnprocessableEntity with default headers values
func NewCreateCostReportUnprocessableEntity() *CreateCostReportUnprocessableEntity {
	return &CreateCostReportUnprocessableEntity{}
}

/*
CreateCostReportUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type CreateCostReportUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create cost report unprocessable entity response has a 2xx status code
func (o *CreateCostReportUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cost report unprocessable entity response has a 3xx status code
func (o *CreateCostReportUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cost report unprocessable entity response has a 4xx status code
func (o *CreateCostReportUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cost report unprocessable entity response has a 5xx status code
func (o *CreateCostReportUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create cost report unprocessable entity response a status code equal to that given
func (o *CreateCostReportUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create cost report unprocessable entity response
func (o *CreateCostReportUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateCostReportUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /cost_reports][%d] createCostReportUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateCostReportUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /cost_reports][%d] createCostReportUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateCostReportUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateCostReportUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
