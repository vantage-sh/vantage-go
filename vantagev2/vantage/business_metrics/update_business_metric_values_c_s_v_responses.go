// Code generated by go-swagger; DO NOT EDIT.

package business_metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// UpdateBusinessMetricValuesCSVReader is a Reader for the UpdateBusinessMetricValuesCSV structure.
type UpdateBusinessMetricValuesCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBusinessMetricValuesCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateBusinessMetricValuesCSVCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateBusinessMetricValuesCSVBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateBusinessMetricValuesCSVForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateBusinessMetricValuesCSVNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateBusinessMetricValuesCSVUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /business_metrics/{business_metric_token}/values.csv] updateBusinessMetricValuesCSV", response, response.Code())
	}
}

// NewUpdateBusinessMetricValuesCSVCreated creates a UpdateBusinessMetricValuesCSVCreated with default headers values
func NewUpdateBusinessMetricValuesCSVCreated() *UpdateBusinessMetricValuesCSVCreated {
	return &UpdateBusinessMetricValuesCSVCreated{}
}

/*
UpdateBusinessMetricValuesCSVCreated describes a response with status code 201, with default header values.

UpdateBusinessMetricValuesCSVCreated update business metric values c s v created
*/
type UpdateBusinessMetricValuesCSVCreated struct {
	Payload *models.BusinessMetric
}

// IsSuccess returns true when this update business metric values c s v created response has a 2xx status code
func (o *UpdateBusinessMetricValuesCSVCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update business metric values c s v created response has a 3xx status code
func (o *UpdateBusinessMetricValuesCSVCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric values c s v created response has a 4xx status code
func (o *UpdateBusinessMetricValuesCSVCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update business metric values c s v created response has a 5xx status code
func (o *UpdateBusinessMetricValuesCSVCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric values c s v created response a status code equal to that given
func (o *UpdateBusinessMetricValuesCSVCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update business metric values c s v created response
func (o *UpdateBusinessMetricValuesCSVCreated) Code() int {
	return 201
}

func (o *UpdateBusinessMetricValuesCSVCreated) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVCreated  %+v", 201, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVCreated) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVCreated  %+v", 201, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVCreated) GetPayload() *models.BusinessMetric {
	return o.Payload
}

func (o *UpdateBusinessMetricValuesCSVCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BusinessMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBusinessMetricValuesCSVBadRequest creates a UpdateBusinessMetricValuesCSVBadRequest with default headers values
func NewUpdateBusinessMetricValuesCSVBadRequest() *UpdateBusinessMetricValuesCSVBadRequest {
	return &UpdateBusinessMetricValuesCSVBadRequest{}
}

/*
UpdateBusinessMetricValuesCSVBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateBusinessMetricValuesCSVBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update business metric values c s v bad request response has a 2xx status code
func (o *UpdateBusinessMetricValuesCSVBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update business metric values c s v bad request response has a 3xx status code
func (o *UpdateBusinessMetricValuesCSVBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric values c s v bad request response has a 4xx status code
func (o *UpdateBusinessMetricValuesCSVBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update business metric values c s v bad request response has a 5xx status code
func (o *UpdateBusinessMetricValuesCSVBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric values c s v bad request response a status code equal to that given
func (o *UpdateBusinessMetricValuesCSVBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update business metric values c s v bad request response
func (o *UpdateBusinessMetricValuesCSVBadRequest) Code() int {
	return 400
}

func (o *UpdateBusinessMetricValuesCSVBadRequest) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVBadRequest) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBusinessMetricValuesCSVBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBusinessMetricValuesCSVForbidden creates a UpdateBusinessMetricValuesCSVForbidden with default headers values
func NewUpdateBusinessMetricValuesCSVForbidden() *UpdateBusinessMetricValuesCSVForbidden {
	return &UpdateBusinessMetricValuesCSVForbidden{}
}

/*
UpdateBusinessMetricValuesCSVForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateBusinessMetricValuesCSVForbidden struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update business metric values c s v forbidden response has a 2xx status code
func (o *UpdateBusinessMetricValuesCSVForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update business metric values c s v forbidden response has a 3xx status code
func (o *UpdateBusinessMetricValuesCSVForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric values c s v forbidden response has a 4xx status code
func (o *UpdateBusinessMetricValuesCSVForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update business metric values c s v forbidden response has a 5xx status code
func (o *UpdateBusinessMetricValuesCSVForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric values c s v forbidden response a status code equal to that given
func (o *UpdateBusinessMetricValuesCSVForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update business metric values c s v forbidden response
func (o *UpdateBusinessMetricValuesCSVForbidden) Code() int {
	return 403
}

func (o *UpdateBusinessMetricValuesCSVForbidden) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVForbidden  %+v", 403, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVForbidden) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVForbidden  %+v", 403, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBusinessMetricValuesCSVForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBusinessMetricValuesCSVNotFound creates a UpdateBusinessMetricValuesCSVNotFound with default headers values
func NewUpdateBusinessMetricValuesCSVNotFound() *UpdateBusinessMetricValuesCSVNotFound {
	return &UpdateBusinessMetricValuesCSVNotFound{}
}

/*
UpdateBusinessMetricValuesCSVNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateBusinessMetricValuesCSVNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update business metric values c s v not found response has a 2xx status code
func (o *UpdateBusinessMetricValuesCSVNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update business metric values c s v not found response has a 3xx status code
func (o *UpdateBusinessMetricValuesCSVNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric values c s v not found response has a 4xx status code
func (o *UpdateBusinessMetricValuesCSVNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update business metric values c s v not found response has a 5xx status code
func (o *UpdateBusinessMetricValuesCSVNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric values c s v not found response a status code equal to that given
func (o *UpdateBusinessMetricValuesCSVNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update business metric values c s v not found response
func (o *UpdateBusinessMetricValuesCSVNotFound) Code() int {
	return 404
}

func (o *UpdateBusinessMetricValuesCSVNotFound) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVNotFound) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBusinessMetricValuesCSVNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBusinessMetricValuesCSVUnprocessableEntity creates a UpdateBusinessMetricValuesCSVUnprocessableEntity with default headers values
func NewUpdateBusinessMetricValuesCSVUnprocessableEntity() *UpdateBusinessMetricValuesCSVUnprocessableEntity {
	return &UpdateBusinessMetricValuesCSVUnprocessableEntity{}
}

/*
UpdateBusinessMetricValuesCSVUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type UpdateBusinessMetricValuesCSVUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update business metric values c s v unprocessable entity response has a 2xx status code
func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update business metric values c s v unprocessable entity response has a 3xx status code
func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric values c s v unprocessable entity response has a 4xx status code
func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update business metric values c s v unprocessable entity response has a 5xx status code
func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric values c s v unprocessable entity response a status code equal to that given
func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update business metric values c s v unprocessable entity response
func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}/values.csv][%d] updateBusinessMetricValuesCSVUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBusinessMetricValuesCSVUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
