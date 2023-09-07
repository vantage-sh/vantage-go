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

// UpdateSavedFilterReader is a Reader for the UpdateSavedFilter structure.
type UpdateSavedFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSavedFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateSavedFilterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSavedFilterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSavedFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /saved_filters/{saved_filter_token}] updateSavedFilter", response, response.Code())
	}
}

// NewUpdateSavedFilterCreated creates a UpdateSavedFilterCreated with default headers values
func NewUpdateSavedFilterCreated() *UpdateSavedFilterCreated {
	return &UpdateSavedFilterCreated{}
}

/*
UpdateSavedFilterCreated describes a response with status code 201, with default header values.

UpdateSavedFilterCreated update saved filter created
*/
type UpdateSavedFilterCreated struct {
	Payload *models.SavedFilter
}

// IsSuccess returns true when this update saved filter created response has a 2xx status code
func (o *UpdateSavedFilterCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update saved filter created response has a 3xx status code
func (o *UpdateSavedFilterCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update saved filter created response has a 4xx status code
func (o *UpdateSavedFilterCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update saved filter created response has a 5xx status code
func (o *UpdateSavedFilterCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update saved filter created response a status code equal to that given
func (o *UpdateSavedFilterCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update saved filter created response
func (o *UpdateSavedFilterCreated) Code() int {
	return 201
}

func (o *UpdateSavedFilterCreated) Error() string {
	return fmt.Sprintf("[PUT /saved_filters/{saved_filter_token}][%d] updateSavedFilterCreated  %+v", 201, o.Payload)
}

func (o *UpdateSavedFilterCreated) String() string {
	return fmt.Sprintf("[PUT /saved_filters/{saved_filter_token}][%d] updateSavedFilterCreated  %+v", 201, o.Payload)
}

func (o *UpdateSavedFilterCreated) GetPayload() *models.SavedFilter {
	return o.Payload
}

func (o *UpdateSavedFilterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SavedFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSavedFilterBadRequest creates a UpdateSavedFilterBadRequest with default headers values
func NewUpdateSavedFilterBadRequest() *UpdateSavedFilterBadRequest {
	return &UpdateSavedFilterBadRequest{}
}

/*
UpdateSavedFilterBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateSavedFilterBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update saved filter bad request response has a 2xx status code
func (o *UpdateSavedFilterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update saved filter bad request response has a 3xx status code
func (o *UpdateSavedFilterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update saved filter bad request response has a 4xx status code
func (o *UpdateSavedFilterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update saved filter bad request response has a 5xx status code
func (o *UpdateSavedFilterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update saved filter bad request response a status code equal to that given
func (o *UpdateSavedFilterBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update saved filter bad request response
func (o *UpdateSavedFilterBadRequest) Code() int {
	return 400
}

func (o *UpdateSavedFilterBadRequest) Error() string {
	return fmt.Sprintf("[PUT /saved_filters/{saved_filter_token}][%d] updateSavedFilterBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSavedFilterBadRequest) String() string {
	return fmt.Sprintf("[PUT /saved_filters/{saved_filter_token}][%d] updateSavedFilterBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSavedFilterBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateSavedFilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSavedFilterNotFound creates a UpdateSavedFilterNotFound with default headers values
func NewUpdateSavedFilterNotFound() *UpdateSavedFilterNotFound {
	return &UpdateSavedFilterNotFound{}
}

/*
UpdateSavedFilterNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateSavedFilterNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update saved filter not found response has a 2xx status code
func (o *UpdateSavedFilterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update saved filter not found response has a 3xx status code
func (o *UpdateSavedFilterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update saved filter not found response has a 4xx status code
func (o *UpdateSavedFilterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update saved filter not found response has a 5xx status code
func (o *UpdateSavedFilterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update saved filter not found response a status code equal to that given
func (o *UpdateSavedFilterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update saved filter not found response
func (o *UpdateSavedFilterNotFound) Code() int {
	return 404
}

func (o *UpdateSavedFilterNotFound) Error() string {
	return fmt.Sprintf("[PUT /saved_filters/{saved_filter_token}][%d] updateSavedFilterNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSavedFilterNotFound) String() string {
	return fmt.Sprintf("[PUT /saved_filters/{saved_filter_token}][%d] updateSavedFilterNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSavedFilterNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateSavedFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}