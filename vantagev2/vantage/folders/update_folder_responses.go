// Code generated by go-swagger; DO NOT EDIT.

package folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// UpdateFolderReader is a Reader for the UpdateFolder structure.
type UpdateFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /folders/{folder_token}] updateFolder", response, response.Code())
	}
}

// NewUpdateFolderOK creates a UpdateFolderOK with default headers values
func NewUpdateFolderOK() *UpdateFolderOK {
	return &UpdateFolderOK{}
}

/*
UpdateFolderOK describes a response with status code 200, with default header values.

UpdateFolderOK update folder o k
*/
type UpdateFolderOK struct {
	Payload *models.Folder
}

// IsSuccess returns true when this update folder o k response has a 2xx status code
func (o *UpdateFolderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update folder o k response has a 3xx status code
func (o *UpdateFolderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update folder o k response has a 4xx status code
func (o *UpdateFolderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update folder o k response has a 5xx status code
func (o *UpdateFolderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update folder o k response a status code equal to that given
func (o *UpdateFolderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update folder o k response
func (o *UpdateFolderOK) Code() int {
	return 200
}

func (o *UpdateFolderOK) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_token}][%d] updateFolderOK  %+v", 200, o.Payload)
}

func (o *UpdateFolderOK) String() string {
	return fmt.Sprintf("[PUT /folders/{folder_token}][%d] updateFolderOK  %+v", 200, o.Payload)
}

func (o *UpdateFolderOK) GetPayload() *models.Folder {
	return o.Payload
}

func (o *UpdateFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderBadRequest creates a UpdateFolderBadRequest with default headers values
func NewUpdateFolderBadRequest() *UpdateFolderBadRequest {
	return &UpdateFolderBadRequest{}
}

/*
UpdateFolderBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateFolderBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update folder bad request response has a 2xx status code
func (o *UpdateFolderBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update folder bad request response has a 3xx status code
func (o *UpdateFolderBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update folder bad request response has a 4xx status code
func (o *UpdateFolderBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update folder bad request response has a 5xx status code
func (o *UpdateFolderBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update folder bad request response a status code equal to that given
func (o *UpdateFolderBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update folder bad request response
func (o *UpdateFolderBadRequest) Code() int {
	return 400
}

func (o *UpdateFolderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_token}][%d] updateFolderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateFolderBadRequest) String() string {
	return fmt.Sprintf("[PUT /folders/{folder_token}][%d] updateFolderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateFolderBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderNotFound creates a UpdateFolderNotFound with default headers values
func NewUpdateFolderNotFound() *UpdateFolderNotFound {
	return &UpdateFolderNotFound{}
}

/*
UpdateFolderNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateFolderNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update folder not found response has a 2xx status code
func (o *UpdateFolderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update folder not found response has a 3xx status code
func (o *UpdateFolderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update folder not found response has a 4xx status code
func (o *UpdateFolderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update folder not found response has a 5xx status code
func (o *UpdateFolderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update folder not found response a status code equal to that given
func (o *UpdateFolderNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update folder not found response
func (o *UpdateFolderNotFound) Code() int {
	return 404
}

func (o *UpdateFolderNotFound) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_token}][%d] updateFolderNotFound  %+v", 404, o.Payload)
}

func (o *UpdateFolderNotFound) String() string {
	return fmt.Sprintf("[PUT /folders/{folder_token}][%d] updateFolderNotFound  %+v", 404, o.Payload)
}

func (o *UpdateFolderNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
