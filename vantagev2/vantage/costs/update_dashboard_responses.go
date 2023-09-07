// Code generated by go-swagger; DO NOT EDIT.

package costs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// UpdateDashboardReader is a Reader for the UpdateDashboard structure.
type UpdateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /dashboards/{dashboard_token}] updateDashboard", response, response.Code())
	}
}

// NewUpdateDashboardOK creates a UpdateDashboardOK with default headers values
func NewUpdateDashboardOK() *UpdateDashboardOK {
	return &UpdateDashboardOK{}
}

/*
UpdateDashboardOK describes a response with status code 200, with default header values.

UpdateDashboardOK update dashboard o k
*/
type UpdateDashboardOK struct {
	Payload *models.Dashboard
}

// IsSuccess returns true when this update dashboard o k response has a 2xx status code
func (o *UpdateDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update dashboard o k response has a 3xx status code
func (o *UpdateDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard o k response has a 4xx status code
func (o *UpdateDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update dashboard o k response has a 5xx status code
func (o *UpdateDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard o k response a status code equal to that given
func (o *UpdateDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update dashboard o k response
func (o *UpdateDashboardOK) Code() int {
	return 200
}

func (o *UpdateDashboardOK) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboard_token}][%d] updateDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateDashboardOK) String() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboard_token}][%d] updateDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateDashboardOK) GetPayload() *models.Dashboard {
	return o.Payload
}

func (o *UpdateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardBadRequest creates a UpdateDashboardBadRequest with default headers values
func NewUpdateDashboardBadRequest() *UpdateDashboardBadRequest {
	return &UpdateDashboardBadRequest{}
}

/*
UpdateDashboardBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateDashboardBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update dashboard bad request response has a 2xx status code
func (o *UpdateDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard bad request response has a 3xx status code
func (o *UpdateDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard bad request response has a 4xx status code
func (o *UpdateDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard bad request response has a 5xx status code
func (o *UpdateDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard bad request response a status code equal to that given
func (o *UpdateDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update dashboard bad request response
func (o *UpdateDashboardBadRequest) Code() int {
	return 400
}

func (o *UpdateDashboardBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboard_token}][%d] updateDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDashboardBadRequest) String() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboard_token}][%d] updateDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDashboardBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardNotFound creates a UpdateDashboardNotFound with default headers values
func NewUpdateDashboardNotFound() *UpdateDashboardNotFound {
	return &UpdateDashboardNotFound{}
}

/*
UpdateDashboardNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateDashboardNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update dashboard not found response has a 2xx status code
func (o *UpdateDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard not found response has a 3xx status code
func (o *UpdateDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard not found response has a 4xx status code
func (o *UpdateDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard not found response has a 5xx status code
func (o *UpdateDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard not found response a status code equal to that given
func (o *UpdateDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update dashboard not found response
func (o *UpdateDashboardNotFound) Code() int {
	return 404
}

func (o *UpdateDashboardNotFound) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboard_token}][%d] updateDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDashboardNotFound) String() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboard_token}][%d] updateDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDashboardNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}