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

// CreateDashboardReader is a Reader for the CreateDashboard structure.
type CreateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDashboardCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /dashboards] createDashboard", response, response.Code())
	}
}

// NewCreateDashboardCreated creates a CreateDashboardCreated with default headers values
func NewCreateDashboardCreated() *CreateDashboardCreated {
	return &CreateDashboardCreated{}
}

/*
CreateDashboardCreated describes a response with status code 201, with default header values.

CreateDashboardCreated create dashboard created
*/
type CreateDashboardCreated struct {
	Payload *models.Dashboard
}

// IsSuccess returns true when this create dashboard created response has a 2xx status code
func (o *CreateDashboardCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create dashboard created response has a 3xx status code
func (o *CreateDashboardCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dashboard created response has a 4xx status code
func (o *CreateDashboardCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create dashboard created response has a 5xx status code
func (o *CreateDashboardCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create dashboard created response a status code equal to that given
func (o *CreateDashboardCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create dashboard created response
func (o *CreateDashboardCreated) Code() int {
	return 201
}

func (o *CreateDashboardCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardCreated %s", 201, payload)
}

func (o *CreateDashboardCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardCreated %s", 201, payload)
}

func (o *CreateDashboardCreated) GetPayload() *models.Dashboard {
	return o.Payload
}

func (o *CreateDashboardCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardBadRequest creates a CreateDashboardBadRequest with default headers values
func NewCreateDashboardBadRequest() *CreateDashboardBadRequest {
	return &CreateDashboardBadRequest{}
}

/*
CreateDashboardBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateDashboardBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create dashboard bad request response has a 2xx status code
func (o *CreateDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create dashboard bad request response has a 3xx status code
func (o *CreateDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dashboard bad request response has a 4xx status code
func (o *CreateDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create dashboard bad request response has a 5xx status code
func (o *CreateDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create dashboard bad request response a status code equal to that given
func (o *CreateDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create dashboard bad request response
func (o *CreateDashboardBadRequest) Code() int {
	return 400
}

func (o *CreateDashboardBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardBadRequest %s", 400, payload)
}

func (o *CreateDashboardBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardBadRequest %s", 400, payload)
}

func (o *CreateDashboardBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
