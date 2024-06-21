// Code generated by go-swagger; DO NOT EDIT.

package virtual_tags

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

// GetVirtualTagConfigsReader is a Reader for the GetVirtualTagConfigs structure.
type GetVirtualTagConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVirtualTagConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVirtualTagConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /virtual_tag_configs] getVirtualTagConfigs", response, response.Code())
	}
}

// NewGetVirtualTagConfigsOK creates a GetVirtualTagConfigsOK with default headers values
func NewGetVirtualTagConfigsOK() *GetVirtualTagConfigsOK {
	return &GetVirtualTagConfigsOK{}
}

/*
GetVirtualTagConfigsOK describes a response with status code 200, with default header values.

GetVirtualTagConfigsOK get virtual tag configs o k
*/
type GetVirtualTagConfigsOK struct {
	Payload *models.VirtualTagConfigs
}

// IsSuccess returns true when this get virtual tag configs o k response has a 2xx status code
func (o *GetVirtualTagConfigsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get virtual tag configs o k response has a 3xx status code
func (o *GetVirtualTagConfigsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get virtual tag configs o k response has a 4xx status code
func (o *GetVirtualTagConfigsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get virtual tag configs o k response has a 5xx status code
func (o *GetVirtualTagConfigsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get virtual tag configs o k response a status code equal to that given
func (o *GetVirtualTagConfigsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get virtual tag configs o k response
func (o *GetVirtualTagConfigsOK) Code() int {
	return 200
}

func (o *GetVirtualTagConfigsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtual_tag_configs][%d] getVirtualTagConfigsOK %s", 200, payload)
}

func (o *GetVirtualTagConfigsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtual_tag_configs][%d] getVirtualTagConfigsOK %s", 200, payload)
}

func (o *GetVirtualTagConfigsOK) GetPayload() *models.VirtualTagConfigs {
	return o.Payload
}

func (o *GetVirtualTagConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualTagConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
