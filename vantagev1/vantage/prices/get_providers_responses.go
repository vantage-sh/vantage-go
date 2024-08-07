// Code generated by go-swagger; DO NOT EDIT.

package prices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev1/models"
)

// GetProvidersReader is a Reader for the GetProviders structure.
type GetProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /providers] getProviders", response, response.Code())
	}
}

// NewGetProvidersOK creates a GetProvidersOK with default headers values
func NewGetProvidersOK() *GetProvidersOK {
	return &GetProvidersOK{}
}

/*
GetProvidersOK describes a response with status code 200, with default header values.

GetProvidersOK get providers o k
*/
type GetProvidersOK struct {
	Payload *models.Providers
}

// IsSuccess returns true when this get providers o k response has a 2xx status code
func (o *GetProvidersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get providers o k response has a 3xx status code
func (o *GetProvidersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get providers o k response has a 4xx status code
func (o *GetProvidersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get providers o k response has a 5xx status code
func (o *GetProvidersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get providers o k response a status code equal to that given
func (o *GetProvidersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get providers o k response
func (o *GetProvidersOK) Code() int {
	return 200
}

func (o *GetProvidersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /providers][%d] getProvidersOK %s", 200, payload)
}

func (o *GetProvidersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /providers][%d] getProvidersOK %s", 200, payload)
}

func (o *GetProvidersOK) GetPayload() *models.Providers {
	return o.Payload
}

func (o *GetProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Providers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
