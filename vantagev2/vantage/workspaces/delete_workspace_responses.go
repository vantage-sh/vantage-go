// Code generated by go-swagger; DO NOT EDIT.

package workspaces

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

// DeleteWorkspaceReader is a Reader for the DeleteWorkspace structure.
type DeleteWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkspaceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteWorkspaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /workspaces/{workspace_token}] deleteWorkspace", response, response.Code())
	}
}

// NewDeleteWorkspaceNoContent creates a DeleteWorkspaceNoContent with default headers values
func NewDeleteWorkspaceNoContent() *DeleteWorkspaceNoContent {
	return &DeleteWorkspaceNoContent{}
}

/*
DeleteWorkspaceNoContent describes a response with status code 204, with default header values.

DeleteWorkspaceNoContent delete workspace no content
*/
type DeleteWorkspaceNoContent struct {
	Payload *models.Workspace
}

// IsSuccess returns true when this delete workspace no content response has a 2xx status code
func (o *DeleteWorkspaceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete workspace no content response has a 3xx status code
func (o *DeleteWorkspaceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workspace no content response has a 4xx status code
func (o *DeleteWorkspaceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workspace no content response has a 5xx status code
func (o *DeleteWorkspaceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workspace no content response a status code equal to that given
func (o *DeleteWorkspaceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete workspace no content response
func (o *DeleteWorkspaceNoContent) Code() int {
	return 204
}

func (o *DeleteWorkspaceNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /workspaces/{workspace_token}][%d] deleteWorkspaceNoContent %s", 204, payload)
}

func (o *DeleteWorkspaceNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /workspaces/{workspace_token}][%d] deleteWorkspaceNoContent %s", 204, payload)
}

func (o *DeleteWorkspaceNoContent) GetPayload() *models.Workspace {
	return o.Payload
}

func (o *DeleteWorkspaceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkspaceNotFound creates a DeleteWorkspaceNotFound with default headers values
func NewDeleteWorkspaceNotFound() *DeleteWorkspaceNotFound {
	return &DeleteWorkspaceNotFound{}
}

/*
DeleteWorkspaceNotFound describes a response with status code 404, with default header values.

NotFound
*/
type DeleteWorkspaceNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this delete workspace not found response has a 2xx status code
func (o *DeleteWorkspaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workspace not found response has a 3xx status code
func (o *DeleteWorkspaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workspace not found response has a 4xx status code
func (o *DeleteWorkspaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workspace not found response has a 5xx status code
func (o *DeleteWorkspaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workspace not found response a status code equal to that given
func (o *DeleteWorkspaceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete workspace not found response
func (o *DeleteWorkspaceNotFound) Code() int {
	return 404
}

func (o *DeleteWorkspaceNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /workspaces/{workspace_token}][%d] deleteWorkspaceNotFound %s", 404, payload)
}

func (o *DeleteWorkspaceNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /workspaces/{workspace_token}][%d] deleteWorkspaceNotFound %s", 404, payload)
}

func (o *DeleteWorkspaceNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteWorkspaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
