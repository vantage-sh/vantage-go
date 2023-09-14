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

// GetFolderReader is a Reader for the GetFolder structure.
type GetFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /folders/{folder_token}] getFolder", response, response.Code())
	}
}

// NewGetFolderOK creates a GetFolderOK with default headers values
func NewGetFolderOK() *GetFolderOK {
	return &GetFolderOK{}
}

/*
GetFolderOK describes a response with status code 200, with default header values.

GetFolderOK get folder o k
*/
type GetFolderOK struct {
	Payload *models.Folder
}

// IsSuccess returns true when this get folder o k response has a 2xx status code
func (o *GetFolderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get folder o k response has a 3xx status code
func (o *GetFolderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder o k response has a 4xx status code
func (o *GetFolderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folder o k response has a 5xx status code
func (o *GetFolderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder o k response a status code equal to that given
func (o *GetFolderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get folder o k response
func (o *GetFolderOK) Code() int {
	return 200
}

func (o *GetFolderOK) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_token}][%d] getFolderOK  %+v", 200, o.Payload)
}

func (o *GetFolderOK) String() string {
	return fmt.Sprintf("[GET /folders/{folder_token}][%d] getFolderOK  %+v", 200, o.Payload)
}

func (o *GetFolderOK) GetPayload() *models.Folder {
	return o.Payload
}

func (o *GetFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderNotFound creates a GetFolderNotFound with default headers values
func NewGetFolderNotFound() *GetFolderNotFound {
	return &GetFolderNotFound{}
}

/*
GetFolderNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetFolderNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this get folder not found response has a 2xx status code
func (o *GetFolderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder not found response has a 3xx status code
func (o *GetFolderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder not found response has a 4xx status code
func (o *GetFolderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder not found response has a 5xx status code
func (o *GetFolderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder not found response a status code equal to that given
func (o *GetFolderNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get folder not found response
func (o *GetFolderNotFound) Code() int {
	return 404
}

func (o *GetFolderNotFound) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_token}][%d] getFolderNotFound  %+v", 404, o.Payload)
}

func (o *GetFolderNotFound) String() string {
	return fmt.Sprintf("[GET /folders/{folder_token}][%d] getFolderNotFound  %+v", 404, o.Payload)
}

func (o *GetFolderNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
