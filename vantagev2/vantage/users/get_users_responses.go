// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users] getUsers", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*
GetUsersOK describes a response with status code 200, with default header values.

GetUsersOK get users o k
*/
type GetUsersOK struct {
	Payload *models.Users
}

// IsSuccess returns true when this get users o k response has a 2xx status code
func (o *GetUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users o k response has a 3xx status code
func (o *GetUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users o k response has a 4xx status code
func (o *GetUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users o k response has a 5xx status code
func (o *GetUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users o k response a status code equal to that given
func (o *GetUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users o k response
func (o *GetUsersOK) Code() int {
	return 200
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) String() string {
	return fmt.Sprintf("[GET /users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) GetPayload() *models.Users {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Users)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForbidden creates a GetUsersForbidden with default headers values
func NewGetUsersForbidden() *GetUsersForbidden {
	return &GetUsersForbidden{}
}

/*
GetUsersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsersForbidden struct {
	Payload *models.Errors
}

// IsSuccess returns true when this get users forbidden response has a 2xx status code
func (o *GetUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users forbidden response has a 3xx status code
func (o *GetUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users forbidden response has a 4xx status code
func (o *GetUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users forbidden response has a 5xx status code
func (o *GetUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get users forbidden response a status code equal to that given
func (o *GetUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get users forbidden response
func (o *GetUsersForbidden) Code() int {
	return 403
}

func (o *GetUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersForbidden) String() string {
	return fmt.Sprintf("[GET /users][%d] getUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
