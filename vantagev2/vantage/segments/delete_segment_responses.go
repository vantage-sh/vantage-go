// Code generated by go-swagger; DO NOT EDIT.

package segments

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

// DeleteSegmentReader is a Reader for the DeleteSegment structure.
type DeleteSegmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSegmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSegmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSegmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /segments/{segment_token}] deleteSegment", response, response.Code())
	}
}

// NewDeleteSegmentNoContent creates a DeleteSegmentNoContent with default headers values
func NewDeleteSegmentNoContent() *DeleteSegmentNoContent {
	return &DeleteSegmentNoContent{}
}

/*
DeleteSegmentNoContent describes a response with status code 204, with default header values.

DeleteSegmentNoContent delete segment no content
*/
type DeleteSegmentNoContent struct {
	Payload *models.Segment
}

// IsSuccess returns true when this delete segment no content response has a 2xx status code
func (o *DeleteSegmentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete segment no content response has a 3xx status code
func (o *DeleteSegmentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete segment no content response has a 4xx status code
func (o *DeleteSegmentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete segment no content response has a 5xx status code
func (o *DeleteSegmentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete segment no content response a status code equal to that given
func (o *DeleteSegmentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete segment no content response
func (o *DeleteSegmentNoContent) Code() int {
	return 204
}

func (o *DeleteSegmentNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /segments/{segment_token}][%d] deleteSegmentNoContent %s", 204, payload)
}

func (o *DeleteSegmentNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /segments/{segment_token}][%d] deleteSegmentNoContent %s", 204, payload)
}

func (o *DeleteSegmentNoContent) GetPayload() *models.Segment {
	return o.Payload
}

func (o *DeleteSegmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Segment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSegmentNotFound creates a DeleteSegmentNotFound with default headers values
func NewDeleteSegmentNotFound() *DeleteSegmentNotFound {
	return &DeleteSegmentNotFound{}
}

/*
DeleteSegmentNotFound describes a response with status code 404, with default header values.

NotFound
*/
type DeleteSegmentNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this delete segment not found response has a 2xx status code
func (o *DeleteSegmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete segment not found response has a 3xx status code
func (o *DeleteSegmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete segment not found response has a 4xx status code
func (o *DeleteSegmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete segment not found response has a 5xx status code
func (o *DeleteSegmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete segment not found response a status code equal to that given
func (o *DeleteSegmentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete segment not found response
func (o *DeleteSegmentNotFound) Code() int {
	return 404
}

func (o *DeleteSegmentNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /segments/{segment_token}][%d] deleteSegmentNotFound %s", 404, payload)
}

func (o *DeleteSegmentNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /segments/{segment_token}][%d] deleteSegmentNotFound %s", 404, payload)
}

func (o *DeleteSegmentNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteSegmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
