// Code generated by go-swagger; DO NOT EDIT.

package saved_filters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// DeleteSavedFilterReader is a Reader for the DeleteSavedFilter structure.
type DeleteSavedFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSavedFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSavedFilterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSavedFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /saved_filters/{saved_filter_token}] deleteSavedFilter", response, response.Code())
	}
}

// NewDeleteSavedFilterNoContent creates a DeleteSavedFilterNoContent with default headers values
func NewDeleteSavedFilterNoContent() *DeleteSavedFilterNoContent {
	return &DeleteSavedFilterNoContent{}
}

/*
DeleteSavedFilterNoContent describes a response with status code 204, with default header values.

DeleteSavedFilterNoContent delete saved filter no content
*/
type DeleteSavedFilterNoContent struct {
	Payload *models.SavedFilter
}

// IsSuccess returns true when this delete saved filter no content response has a 2xx status code
func (o *DeleteSavedFilterNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete saved filter no content response has a 3xx status code
func (o *DeleteSavedFilterNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete saved filter no content response has a 4xx status code
func (o *DeleteSavedFilterNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete saved filter no content response has a 5xx status code
func (o *DeleteSavedFilterNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete saved filter no content response a status code equal to that given
func (o *DeleteSavedFilterNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete saved filter no content response
func (o *DeleteSavedFilterNoContent) Code() int {
	return 204
}

func (o *DeleteSavedFilterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /saved_filters/{saved_filter_token}][%d] deleteSavedFilterNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSavedFilterNoContent) String() string {
	return fmt.Sprintf("[DELETE /saved_filters/{saved_filter_token}][%d] deleteSavedFilterNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSavedFilterNoContent) GetPayload() *models.SavedFilter {
	return o.Payload
}

func (o *DeleteSavedFilterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SavedFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedFilterNotFound creates a DeleteSavedFilterNotFound with default headers values
func NewDeleteSavedFilterNotFound() *DeleteSavedFilterNotFound {
	return &DeleteSavedFilterNotFound{}
}

/*
DeleteSavedFilterNotFound describes a response with status code 404, with default header values.

NotFound
*/
type DeleteSavedFilterNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this delete saved filter not found response has a 2xx status code
func (o *DeleteSavedFilterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete saved filter not found response has a 3xx status code
func (o *DeleteSavedFilterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete saved filter not found response has a 4xx status code
func (o *DeleteSavedFilterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete saved filter not found response has a 5xx status code
func (o *DeleteSavedFilterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete saved filter not found response a status code equal to that given
func (o *DeleteSavedFilterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete saved filter not found response
func (o *DeleteSavedFilterNotFound) Code() int {
	return 404
}

func (o *DeleteSavedFilterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /saved_filters/{saved_filter_token}][%d] deleteSavedFilterNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSavedFilterNotFound) String() string {
	return fmt.Sprintf("[DELETE /saved_filters/{saved_filter_token}][%d] deleteSavedFilterNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSavedFilterNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteSavedFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
