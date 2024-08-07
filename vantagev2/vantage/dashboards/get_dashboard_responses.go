// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// GetDashboardReader is a Reader for the GetDashboard structure.
type GetDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /dashboards/{dashboard_token}] getDashboard", response, response.Code())
	}
}

// NewGetDashboardOK creates a GetDashboardOK with default headers values
func NewGetDashboardOK() *GetDashboardOK {
	return &GetDashboardOK{}
}

/*
GetDashboardOK describes a response with status code 200, with default header values.

GetDashboardOK get dashboard o k
*/
type GetDashboardOK struct {
	Payload *models.Dashboard
}

// IsSuccess returns true when this get dashboard o k response has a 2xx status code
func (o *GetDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dashboard o k response has a 3xx status code
func (o *GetDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard o k response has a 4xx status code
func (o *GetDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard o k response has a 5xx status code
func (o *GetDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard o k response a status code equal to that given
func (o *GetDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get dashboard o k response
func (o *GetDashboardOK) Code() int {
	return 200
}

func (o *GetDashboardOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/{dashboard_token}][%d] getDashboardOK %s", 200, payload)
}

func (o *GetDashboardOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/{dashboard_token}][%d] getDashboardOK %s", 200, payload)
}

func (o *GetDashboardOK) GetPayload() *models.Dashboard {
	return o.Payload
}

func (o *GetDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardNotFound creates a GetDashboardNotFound with default headers values
func NewGetDashboardNotFound() *GetDashboardNotFound {
	return &GetDashboardNotFound{}
}

/*
GetDashboardNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetDashboardNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this get dashboard not found response has a 2xx status code
func (o *GetDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard not found response has a 3xx status code
func (o *GetDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard not found response has a 4xx status code
func (o *GetDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard not found response has a 5xx status code
func (o *GetDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard not found response a status code equal to that given
func (o *GetDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get dashboard not found response
func (o *GetDashboardNotFound) Code() int {
	return 404
}

func (o *GetDashboardNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/{dashboard_token}][%d] getDashboardNotFound %s", 404, payload)
}

func (o *GetDashboardNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/{dashboard_token}][%d] getDashboardNotFound %s", 404, payload)
}

func (o *GetDashboardNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
