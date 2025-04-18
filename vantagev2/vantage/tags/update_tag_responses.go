// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// UpdateTagReader is a Reader for the UpdateTag structure.
type UpdateTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateTagUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /tags] updateTag", response, response.Code())
	}
}

// NewUpdateTagOK creates a UpdateTagOK with default headers values
func NewUpdateTagOK() *UpdateTagOK {
	return &UpdateTagOK{}
}

/*
UpdateTagOK describes a response with status code 200, with default header values.

UpdateTagOK update tag o k
*/
type UpdateTagOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this update tag o k response has a 2xx status code
func (o *UpdateTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update tag o k response has a 3xx status code
func (o *UpdateTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update tag o k response has a 4xx status code
func (o *UpdateTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update tag o k response has a 5xx status code
func (o *UpdateTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update tag o k response a status code equal to that given
func (o *UpdateTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update tag o k response
func (o *UpdateTagOK) Code() int {
	return 200
}

func (o *UpdateTagOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagOK %s", 200, payload)
}

func (o *UpdateTagOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagOK %s", 200, payload)
}

func (o *UpdateTagOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *UpdateTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagBadRequest creates a UpdateTagBadRequest with default headers values
func NewUpdateTagBadRequest() *UpdateTagBadRequest {
	return &UpdateTagBadRequest{}
}

/*
UpdateTagBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateTagBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update tag bad request response has a 2xx status code
func (o *UpdateTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update tag bad request response has a 3xx status code
func (o *UpdateTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update tag bad request response has a 4xx status code
func (o *UpdateTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update tag bad request response has a 5xx status code
func (o *UpdateTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update tag bad request response a status code equal to that given
func (o *UpdateTagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update tag bad request response
func (o *UpdateTagBadRequest) Code() int {
	return 400
}

func (o *UpdateTagBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagBadRequest %s", 400, payload)
}

func (o *UpdateTagBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagBadRequest %s", 400, payload)
}

func (o *UpdateTagBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagForbidden creates a UpdateTagForbidden with default headers values
func NewUpdateTagForbidden() *UpdateTagForbidden {
	return &UpdateTagForbidden{}
}

/*
UpdateTagForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateTagForbidden struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update tag forbidden response has a 2xx status code
func (o *UpdateTagForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update tag forbidden response has a 3xx status code
func (o *UpdateTagForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update tag forbidden response has a 4xx status code
func (o *UpdateTagForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update tag forbidden response has a 5xx status code
func (o *UpdateTagForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update tag forbidden response a status code equal to that given
func (o *UpdateTagForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update tag forbidden response
func (o *UpdateTagForbidden) Code() int {
	return 403
}

func (o *UpdateTagForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagForbidden %s", 403, payload)
}

func (o *UpdateTagForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagForbidden %s", 403, payload)
}

func (o *UpdateTagForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagNotFound creates a UpdateTagNotFound with default headers values
func NewUpdateTagNotFound() *UpdateTagNotFound {
	return &UpdateTagNotFound{}
}

/*
UpdateTagNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateTagNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update tag not found response has a 2xx status code
func (o *UpdateTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update tag not found response has a 3xx status code
func (o *UpdateTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update tag not found response has a 4xx status code
func (o *UpdateTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update tag not found response has a 5xx status code
func (o *UpdateTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update tag not found response a status code equal to that given
func (o *UpdateTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update tag not found response
func (o *UpdateTagNotFound) Code() int {
	return 404
}

func (o *UpdateTagNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagNotFound %s", 404, payload)
}

func (o *UpdateTagNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagNotFound %s", 404, payload)
}

func (o *UpdateTagNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagUnprocessableEntity creates a UpdateTagUnprocessableEntity with default headers values
func NewUpdateTagUnprocessableEntity() *UpdateTagUnprocessableEntity {
	return &UpdateTagUnprocessableEntity{}
}

/*
UpdateTagUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type UpdateTagUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update tag unprocessable entity response has a 2xx status code
func (o *UpdateTagUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update tag unprocessable entity response has a 3xx status code
func (o *UpdateTagUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update tag unprocessable entity response has a 4xx status code
func (o *UpdateTagUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update tag unprocessable entity response has a 5xx status code
func (o *UpdateTagUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update tag unprocessable entity response a status code equal to that given
func (o *UpdateTagUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update tag unprocessable entity response
func (o *UpdateTagUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateTagUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagUnprocessableEntity %s", 422, payload)
}

func (o *UpdateTagUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] updateTagUnprocessableEntity %s", 422, payload)
}

func (o *UpdateTagUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateTagUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
