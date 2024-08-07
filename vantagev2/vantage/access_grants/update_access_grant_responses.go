// Code generated by go-swagger; DO NOT EDIT.

package access_grants

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

// UpdateAccessGrantReader is a Reader for the UpdateAccessGrant structure.
type UpdateAccessGrantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccessGrantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAccessGrantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAccessGrantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAccessGrantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /access_grants/{access_grant_token}] updateAccessGrant", response, response.Code())
	}
}

// NewUpdateAccessGrantOK creates a UpdateAccessGrantOK with default headers values
func NewUpdateAccessGrantOK() *UpdateAccessGrantOK {
	return &UpdateAccessGrantOK{}
}

/*
UpdateAccessGrantOK describes a response with status code 200, with default header values.

UpdateAccessGrantOK update access grant o k
*/
type UpdateAccessGrantOK struct {
	Payload *models.AccessGrant
}

// IsSuccess returns true when this update access grant o k response has a 2xx status code
func (o *UpdateAccessGrantOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update access grant o k response has a 3xx status code
func (o *UpdateAccessGrantOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access grant o k response has a 4xx status code
func (o *UpdateAccessGrantOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update access grant o k response has a 5xx status code
func (o *UpdateAccessGrantOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update access grant o k response a status code equal to that given
func (o *UpdateAccessGrantOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update access grant o k response
func (o *UpdateAccessGrantOK) Code() int {
	return 200
}

func (o *UpdateAccessGrantOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access_grants/{access_grant_token}][%d] updateAccessGrantOK %s", 200, payload)
}

func (o *UpdateAccessGrantOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access_grants/{access_grant_token}][%d] updateAccessGrantOK %s", 200, payload)
}

func (o *UpdateAccessGrantOK) GetPayload() *models.AccessGrant {
	return o.Payload
}

func (o *UpdateAccessGrantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessGrant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAccessGrantBadRequest creates a UpdateAccessGrantBadRequest with default headers values
func NewUpdateAccessGrantBadRequest() *UpdateAccessGrantBadRequest {
	return &UpdateAccessGrantBadRequest{}
}

/*
UpdateAccessGrantBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateAccessGrantBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update access grant bad request response has a 2xx status code
func (o *UpdateAccessGrantBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update access grant bad request response has a 3xx status code
func (o *UpdateAccessGrantBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access grant bad request response has a 4xx status code
func (o *UpdateAccessGrantBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update access grant bad request response has a 5xx status code
func (o *UpdateAccessGrantBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update access grant bad request response a status code equal to that given
func (o *UpdateAccessGrantBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update access grant bad request response
func (o *UpdateAccessGrantBadRequest) Code() int {
	return 400
}

func (o *UpdateAccessGrantBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access_grants/{access_grant_token}][%d] updateAccessGrantBadRequest %s", 400, payload)
}

func (o *UpdateAccessGrantBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access_grants/{access_grant_token}][%d] updateAccessGrantBadRequest %s", 400, payload)
}

func (o *UpdateAccessGrantBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateAccessGrantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAccessGrantNotFound creates a UpdateAccessGrantNotFound with default headers values
func NewUpdateAccessGrantNotFound() *UpdateAccessGrantNotFound {
	return &UpdateAccessGrantNotFound{}
}

/*
UpdateAccessGrantNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateAccessGrantNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update access grant not found response has a 2xx status code
func (o *UpdateAccessGrantNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update access grant not found response has a 3xx status code
func (o *UpdateAccessGrantNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access grant not found response has a 4xx status code
func (o *UpdateAccessGrantNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update access grant not found response has a 5xx status code
func (o *UpdateAccessGrantNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update access grant not found response a status code equal to that given
func (o *UpdateAccessGrantNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update access grant not found response
func (o *UpdateAccessGrantNotFound) Code() int {
	return 404
}

func (o *UpdateAccessGrantNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access_grants/{access_grant_token}][%d] updateAccessGrantNotFound %s", 404, payload)
}

func (o *UpdateAccessGrantNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access_grants/{access_grant_token}][%d] updateAccessGrantNotFound %s", 404, payload)
}

func (o *UpdateAccessGrantNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateAccessGrantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
