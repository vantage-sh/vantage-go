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

// GetFoldersReader is a Reader for the GetFolders structure.
type GetFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFoldersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /folders] getFolders", response, response.Code())
	}
}

// NewGetFoldersOK creates a GetFoldersOK with default headers values
func NewGetFoldersOK() *GetFoldersOK {
	return &GetFoldersOK{}
}

/*
GetFoldersOK describes a response with status code 200, with default header values.

GetFoldersOK get folders o k
*/
type GetFoldersOK struct {
	Payload *models.Folders
}

// IsSuccess returns true when this get folders o k response has a 2xx status code
func (o *GetFoldersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get folders o k response has a 3xx status code
func (o *GetFoldersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folders o k response has a 4xx status code
func (o *GetFoldersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folders o k response has a 5xx status code
func (o *GetFoldersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get folders o k response a status code equal to that given
func (o *GetFoldersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get folders o k response
func (o *GetFoldersOK) Code() int {
	return 200
}

func (o *GetFoldersOK) Error() string {
	return fmt.Sprintf("[GET /folders][%d] getFoldersOK  %+v", 200, o.Payload)
}

func (o *GetFoldersOK) String() string {
	return fmt.Sprintf("[GET /folders][%d] getFoldersOK  %+v", 200, o.Payload)
}

func (o *GetFoldersOK) GetPayload() *models.Folders {
	return o.Payload
}

func (o *GetFoldersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folders)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
