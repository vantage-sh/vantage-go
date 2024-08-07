// Code generated by go-swagger; DO NOT EDIT.

package anomaly_notifications

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

// DeleteAnomalyNotificationReader is a Reader for the DeleteAnomalyNotification structure.
type DeleteAnomalyNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAnomalyNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAnomalyNotificationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAnomalyNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /anomaly_notifications/{anomaly_notification_token}] deleteAnomalyNotification", response, response.Code())
	}
}

// NewDeleteAnomalyNotificationNoContent creates a DeleteAnomalyNotificationNoContent with default headers values
func NewDeleteAnomalyNotificationNoContent() *DeleteAnomalyNotificationNoContent {
	return &DeleteAnomalyNotificationNoContent{}
}

/*
DeleteAnomalyNotificationNoContent describes a response with status code 204, with default header values.

DeleteAnomalyNotificationNoContent delete anomaly notification no content
*/
type DeleteAnomalyNotificationNoContent struct {
	Payload *models.AnomalyNotification
}

// IsSuccess returns true when this delete anomaly notification no content response has a 2xx status code
func (o *DeleteAnomalyNotificationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete anomaly notification no content response has a 3xx status code
func (o *DeleteAnomalyNotificationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete anomaly notification no content response has a 4xx status code
func (o *DeleteAnomalyNotificationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete anomaly notification no content response has a 5xx status code
func (o *DeleteAnomalyNotificationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete anomaly notification no content response a status code equal to that given
func (o *DeleteAnomalyNotificationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete anomaly notification no content response
func (o *DeleteAnomalyNotificationNoContent) Code() int {
	return 204
}

func (o *DeleteAnomalyNotificationNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /anomaly_notifications/{anomaly_notification_token}][%d] deleteAnomalyNotificationNoContent %s", 204, payload)
}

func (o *DeleteAnomalyNotificationNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /anomaly_notifications/{anomaly_notification_token}][%d] deleteAnomalyNotificationNoContent %s", 204, payload)
}

func (o *DeleteAnomalyNotificationNoContent) GetPayload() *models.AnomalyNotification {
	return o.Payload
}

func (o *DeleteAnomalyNotificationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnomalyNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnomalyNotificationNotFound creates a DeleteAnomalyNotificationNotFound with default headers values
func NewDeleteAnomalyNotificationNotFound() *DeleteAnomalyNotificationNotFound {
	return &DeleteAnomalyNotificationNotFound{}
}

/*
DeleteAnomalyNotificationNotFound describes a response with status code 404, with default header values.

NotFound
*/
type DeleteAnomalyNotificationNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this delete anomaly notification not found response has a 2xx status code
func (o *DeleteAnomalyNotificationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete anomaly notification not found response has a 3xx status code
func (o *DeleteAnomalyNotificationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete anomaly notification not found response has a 4xx status code
func (o *DeleteAnomalyNotificationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete anomaly notification not found response has a 5xx status code
func (o *DeleteAnomalyNotificationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete anomaly notification not found response a status code equal to that given
func (o *DeleteAnomalyNotificationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete anomaly notification not found response
func (o *DeleteAnomalyNotificationNotFound) Code() int {
	return 404
}

func (o *DeleteAnomalyNotificationNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /anomaly_notifications/{anomaly_notification_token}][%d] deleteAnomalyNotificationNotFound %s", 404, payload)
}

func (o *DeleteAnomalyNotificationNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /anomaly_notifications/{anomaly_notification_token}][%d] deleteAnomalyNotificationNotFound %s", 404, payload)
}

func (o *DeleteAnomalyNotificationNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteAnomalyNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
