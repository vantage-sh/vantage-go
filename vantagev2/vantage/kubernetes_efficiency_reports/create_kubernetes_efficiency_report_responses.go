// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_efficiency_reports

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

// CreateKubernetesEfficiencyReportReader is a Reader for the CreateKubernetesEfficiencyReport structure.
type CreateKubernetesEfficiencyReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateKubernetesEfficiencyReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateKubernetesEfficiencyReportCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateKubernetesEfficiencyReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateKubernetesEfficiencyReportUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /kubernetes_efficiency_reports] createKubernetesEfficiencyReport", response, response.Code())
	}
}

// NewCreateKubernetesEfficiencyReportCreated creates a CreateKubernetesEfficiencyReportCreated with default headers values
func NewCreateKubernetesEfficiencyReportCreated() *CreateKubernetesEfficiencyReportCreated {
	return &CreateKubernetesEfficiencyReportCreated{}
}

/*
CreateKubernetesEfficiencyReportCreated describes a response with status code 201, with default header values.

CreateKubernetesEfficiencyReportCreated create kubernetes efficiency report created
*/
type CreateKubernetesEfficiencyReportCreated struct {
	Payload *models.KubernetesEfficiencyReport
}

// IsSuccess returns true when this create kubernetes efficiency report created response has a 2xx status code
func (o *CreateKubernetesEfficiencyReportCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create kubernetes efficiency report created response has a 3xx status code
func (o *CreateKubernetesEfficiencyReportCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kubernetes efficiency report created response has a 4xx status code
func (o *CreateKubernetesEfficiencyReportCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create kubernetes efficiency report created response has a 5xx status code
func (o *CreateKubernetesEfficiencyReportCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create kubernetes efficiency report created response a status code equal to that given
func (o *CreateKubernetesEfficiencyReportCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create kubernetes efficiency report created response
func (o *CreateKubernetesEfficiencyReportCreated) Code() int {
	return 201
}

func (o *CreateKubernetesEfficiencyReportCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kubernetes_efficiency_reports][%d] createKubernetesEfficiencyReportCreated %s", 201, payload)
}

func (o *CreateKubernetesEfficiencyReportCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kubernetes_efficiency_reports][%d] createKubernetesEfficiencyReportCreated %s", 201, payload)
}

func (o *CreateKubernetesEfficiencyReportCreated) GetPayload() *models.KubernetesEfficiencyReport {
	return o.Payload
}

func (o *CreateKubernetesEfficiencyReportCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesEfficiencyReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateKubernetesEfficiencyReportBadRequest creates a CreateKubernetesEfficiencyReportBadRequest with default headers values
func NewCreateKubernetesEfficiencyReportBadRequest() *CreateKubernetesEfficiencyReportBadRequest {
	return &CreateKubernetesEfficiencyReportBadRequest{}
}

/*
CreateKubernetesEfficiencyReportBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateKubernetesEfficiencyReportBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create kubernetes efficiency report bad request response has a 2xx status code
func (o *CreateKubernetesEfficiencyReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create kubernetes efficiency report bad request response has a 3xx status code
func (o *CreateKubernetesEfficiencyReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kubernetes efficiency report bad request response has a 4xx status code
func (o *CreateKubernetesEfficiencyReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create kubernetes efficiency report bad request response has a 5xx status code
func (o *CreateKubernetesEfficiencyReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create kubernetes efficiency report bad request response a status code equal to that given
func (o *CreateKubernetesEfficiencyReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create kubernetes efficiency report bad request response
func (o *CreateKubernetesEfficiencyReportBadRequest) Code() int {
	return 400
}

func (o *CreateKubernetesEfficiencyReportBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kubernetes_efficiency_reports][%d] createKubernetesEfficiencyReportBadRequest %s", 400, payload)
}

func (o *CreateKubernetesEfficiencyReportBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kubernetes_efficiency_reports][%d] createKubernetesEfficiencyReportBadRequest %s", 400, payload)
}

func (o *CreateKubernetesEfficiencyReportBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateKubernetesEfficiencyReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateKubernetesEfficiencyReportUnprocessableEntity creates a CreateKubernetesEfficiencyReportUnprocessableEntity with default headers values
func NewCreateKubernetesEfficiencyReportUnprocessableEntity() *CreateKubernetesEfficiencyReportUnprocessableEntity {
	return &CreateKubernetesEfficiencyReportUnprocessableEntity{}
}

/*
CreateKubernetesEfficiencyReportUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type CreateKubernetesEfficiencyReportUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create kubernetes efficiency report unprocessable entity response has a 2xx status code
func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create kubernetes efficiency report unprocessable entity response has a 3xx status code
func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kubernetes efficiency report unprocessable entity response has a 4xx status code
func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create kubernetes efficiency report unprocessable entity response has a 5xx status code
func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create kubernetes efficiency report unprocessable entity response a status code equal to that given
func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create kubernetes efficiency report unprocessable entity response
func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kubernetes_efficiency_reports][%d] createKubernetesEfficiencyReportUnprocessableEntity %s", 422, payload)
}

func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kubernetes_efficiency_reports][%d] createKubernetesEfficiencyReportUnprocessableEntity %s", 422, payload)
}

func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateKubernetesEfficiencyReportUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
