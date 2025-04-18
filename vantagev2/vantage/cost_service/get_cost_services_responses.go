// Code generated by go-swagger; DO NOT EDIT.

package cost_service

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

// GetCostServicesReader is a Reader for the GetCostServices structure.
type GetCostServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCostServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCostServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /cost_services] getCostServices", response, response.Code())
	}
}

// NewGetCostServicesOK creates a GetCostServicesOK with default headers values
func NewGetCostServicesOK() *GetCostServicesOK {
	return &GetCostServicesOK{}
}

/*
GetCostServicesOK describes a response with status code 200, with default header values.

List of CostServices, used to query costs using VQL.
*/
type GetCostServicesOK struct {
	Payload *models.CostServices
}

// IsSuccess returns true when this get cost services o k response has a 2xx status code
func (o *GetCostServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cost services o k response has a 3xx status code
func (o *GetCostServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cost services o k response has a 4xx status code
func (o *GetCostServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cost services o k response has a 5xx status code
func (o *GetCostServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cost services o k response a status code equal to that given
func (o *GetCostServicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cost services o k response
func (o *GetCostServicesOK) Code() int {
	return 200
}

func (o *GetCostServicesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cost_services][%d] getCostServicesOK %s", 200, payload)
}

func (o *GetCostServicesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cost_services][%d] getCostServicesOK %s", 200, payload)
}

func (o *GetCostServicesOK) GetPayload() *models.CostServices {
	return o.Payload
}

func (o *GetCostServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CostServices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
