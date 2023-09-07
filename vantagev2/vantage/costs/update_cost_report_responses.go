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

// UpdateCostReportReader is a Reader for the UpdateCostReport structure.
type UpdateCostReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCostReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCostReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCostReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCostReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /cost_reports/{cost_report_token}] updateCostReport", response, response.Code())
	}
}

// NewUpdateCostReportOK creates a UpdateCostReportOK with default headers values
func NewUpdateCostReportOK() *UpdateCostReportOK {
	return &UpdateCostReportOK{}
}

/*
UpdateCostReportOK describes a response with status code 200, with default header values.

UpdateCostReportOK update cost report o k
*/
type UpdateCostReportOK struct {
	Payload *models.CostReport
}

// IsSuccess returns true when this update cost report o k response has a 2xx status code
func (o *UpdateCostReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update cost report o k response has a 3xx status code
func (o *UpdateCostReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cost report o k response has a 4xx status code
func (o *UpdateCostReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cost report o k response has a 5xx status code
func (o *UpdateCostReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update cost report o k response a status code equal to that given
func (o *UpdateCostReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update cost report o k response
func (o *UpdateCostReportOK) Code() int {
	return 200
}

func (o *UpdateCostReportOK) Error() string {
	return fmt.Sprintf("[PUT /cost_reports/{cost_report_token}][%d] updateCostReportOK  %+v", 200, o.Payload)
}

func (o *UpdateCostReportOK) String() string {
	return fmt.Sprintf("[PUT /cost_reports/{cost_report_token}][%d] updateCostReportOK  %+v", 200, o.Payload)
}

func (o *UpdateCostReportOK) GetPayload() *models.CostReport {
	return o.Payload
}

func (o *UpdateCostReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CostReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCostReportBadRequest creates a UpdateCostReportBadRequest with default headers values
func NewUpdateCostReportBadRequest() *UpdateCostReportBadRequest {
	return &UpdateCostReportBadRequest{}
}

/*
UpdateCostReportBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateCostReportBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update cost report bad request response has a 2xx status code
func (o *UpdateCostReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cost report bad request response has a 3xx status code
func (o *UpdateCostReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cost report bad request response has a 4xx status code
func (o *UpdateCostReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cost report bad request response has a 5xx status code
func (o *UpdateCostReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update cost report bad request response a status code equal to that given
func (o *UpdateCostReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update cost report bad request response
func (o *UpdateCostReportBadRequest) Code() int {
	return 400
}

func (o *UpdateCostReportBadRequest) Error() string {
	return fmt.Sprintf("[PUT /cost_reports/{cost_report_token}][%d] updateCostReportBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCostReportBadRequest) String() string {
	return fmt.Sprintf("[PUT /cost_reports/{cost_report_token}][%d] updateCostReportBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCostReportBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateCostReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCostReportNotFound creates a UpdateCostReportNotFound with default headers values
func NewUpdateCostReportNotFound() *UpdateCostReportNotFound {
	return &UpdateCostReportNotFound{}
}

/*
UpdateCostReportNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateCostReportNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update cost report not found response has a 2xx status code
func (o *UpdateCostReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cost report not found response has a 3xx status code
func (o *UpdateCostReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cost report not found response has a 4xx status code
func (o *UpdateCostReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cost report not found response has a 5xx status code
func (o *UpdateCostReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update cost report not found response a status code equal to that given
func (o *UpdateCostReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update cost report not found response
func (o *UpdateCostReportNotFound) Code() int {
	return 404
}

func (o *UpdateCostReportNotFound) Error() string {
	return fmt.Sprintf("[PUT /cost_reports/{cost_report_token}][%d] updateCostReportNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCostReportNotFound) String() string {
	return fmt.Sprintf("[PUT /cost_reports/{cost_report_token}][%d] updateCostReportNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCostReportNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateCostReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}