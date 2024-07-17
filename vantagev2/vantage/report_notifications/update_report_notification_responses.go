// Code generated by go-swagger; DO NOT EDIT.

package report_notifications

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

// UpdateReportNotificationReader is a Reader for the UpdateReportNotification structure.
type UpdateReportNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReportNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateReportNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateReportNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateReportNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /report_notifications/{report_notification_token}] updateReportNotification", response, response.Code())
	}
}

// NewUpdateReportNotificationOK creates a UpdateReportNotificationOK with default headers values
func NewUpdateReportNotificationOK() *UpdateReportNotificationOK {
	return &UpdateReportNotificationOK{}
}

/*
UpdateReportNotificationOK describes a response with status code 200, with default header values.

UpdateReportNotificationOK update report notification o k
*/
type UpdateReportNotificationOK struct {
	Payload *models.ReportNotification
}

// IsSuccess returns true when this update report notification o k response has a 2xx status code
func (o *UpdateReportNotificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update report notification o k response has a 3xx status code
func (o *UpdateReportNotificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update report notification o k response has a 4xx status code
func (o *UpdateReportNotificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update report notification o k response has a 5xx status code
func (o *UpdateReportNotificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update report notification o k response a status code equal to that given
func (o *UpdateReportNotificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update report notification o k response
func (o *UpdateReportNotificationOK) Code() int {
	return 200
}

func (o *UpdateReportNotificationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /report_notifications/{report_notification_token}][%d] updateReportNotificationOK %s", 200, payload)
}

func (o *UpdateReportNotificationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /report_notifications/{report_notification_token}][%d] updateReportNotificationOK %s", 200, payload)
}

func (o *UpdateReportNotificationOK) GetPayload() *models.ReportNotification {
	return o.Payload
}

func (o *UpdateReportNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReportNotificationBadRequest creates a UpdateReportNotificationBadRequest with default headers values
func NewUpdateReportNotificationBadRequest() *UpdateReportNotificationBadRequest {
	return &UpdateReportNotificationBadRequest{}
}

/*
UpdateReportNotificationBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateReportNotificationBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update report notification bad request response has a 2xx status code
func (o *UpdateReportNotificationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update report notification bad request response has a 3xx status code
func (o *UpdateReportNotificationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update report notification bad request response has a 4xx status code
func (o *UpdateReportNotificationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update report notification bad request response has a 5xx status code
func (o *UpdateReportNotificationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update report notification bad request response a status code equal to that given
func (o *UpdateReportNotificationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update report notification bad request response
func (o *UpdateReportNotificationBadRequest) Code() int {
	return 400
}

func (o *UpdateReportNotificationBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /report_notifications/{report_notification_token}][%d] updateReportNotificationBadRequest %s", 400, payload)
}

func (o *UpdateReportNotificationBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /report_notifications/{report_notification_token}][%d] updateReportNotificationBadRequest %s", 400, payload)
}

func (o *UpdateReportNotificationBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateReportNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReportNotificationNotFound creates a UpdateReportNotificationNotFound with default headers values
func NewUpdateReportNotificationNotFound() *UpdateReportNotificationNotFound {
	return &UpdateReportNotificationNotFound{}
}

/*
UpdateReportNotificationNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateReportNotificationNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update report notification not found response has a 2xx status code
func (o *UpdateReportNotificationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update report notification not found response has a 3xx status code
func (o *UpdateReportNotificationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update report notification not found response has a 4xx status code
func (o *UpdateReportNotificationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update report notification not found response has a 5xx status code
func (o *UpdateReportNotificationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update report notification not found response a status code equal to that given
func (o *UpdateReportNotificationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update report notification not found response
func (o *UpdateReportNotificationNotFound) Code() int {
	return 404
}

func (o *UpdateReportNotificationNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /report_notifications/{report_notification_token}][%d] updateReportNotificationNotFound %s", 404, payload)
}

func (o *UpdateReportNotificationNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /report_notifications/{report_notification_token}][%d] updateReportNotificationNotFound %s", 404, payload)
}

func (o *UpdateReportNotificationNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateReportNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
