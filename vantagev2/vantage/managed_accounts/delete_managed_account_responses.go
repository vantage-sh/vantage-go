// Code generated by go-swagger; DO NOT EDIT.

package managed_accounts

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

// DeleteManagedAccountReader is a Reader for the DeleteManagedAccount structure.
type DeleteManagedAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteManagedAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteManagedAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteManagedAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /managed_accounts/{managed_account_token}] deleteManagedAccount", response, response.Code())
	}
}

// NewDeleteManagedAccountNoContent creates a DeleteManagedAccountNoContent with default headers values
func NewDeleteManagedAccountNoContent() *DeleteManagedAccountNoContent {
	return &DeleteManagedAccountNoContent{}
}

/*
DeleteManagedAccountNoContent describes a response with status code 204, with default header values.

DeleteManagedAccountNoContent delete managed account no content
*/
type DeleteManagedAccountNoContent struct {
	Payload *models.ManagedAccount
}

// IsSuccess returns true when this delete managed account no content response has a 2xx status code
func (o *DeleteManagedAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete managed account no content response has a 3xx status code
func (o *DeleteManagedAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete managed account no content response has a 4xx status code
func (o *DeleteManagedAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete managed account no content response has a 5xx status code
func (o *DeleteManagedAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete managed account no content response a status code equal to that given
func (o *DeleteManagedAccountNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete managed account no content response
func (o *DeleteManagedAccountNoContent) Code() int {
	return 204
}

func (o *DeleteManagedAccountNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /managed_accounts/{managed_account_token}][%d] deleteManagedAccountNoContent %s", 204, payload)
}

func (o *DeleteManagedAccountNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /managed_accounts/{managed_account_token}][%d] deleteManagedAccountNoContent %s", 204, payload)
}

func (o *DeleteManagedAccountNoContent) GetPayload() *models.ManagedAccount {
	return o.Payload
}

func (o *DeleteManagedAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagedAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteManagedAccountNotFound creates a DeleteManagedAccountNotFound with default headers values
func NewDeleteManagedAccountNotFound() *DeleteManagedAccountNotFound {
	return &DeleteManagedAccountNotFound{}
}

/*
DeleteManagedAccountNotFound describes a response with status code 404, with default header values.

NotFound
*/
type DeleteManagedAccountNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this delete managed account not found response has a 2xx status code
func (o *DeleteManagedAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete managed account not found response has a 3xx status code
func (o *DeleteManagedAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete managed account not found response has a 4xx status code
func (o *DeleteManagedAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete managed account not found response has a 5xx status code
func (o *DeleteManagedAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete managed account not found response a status code equal to that given
func (o *DeleteManagedAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete managed account not found response
func (o *DeleteManagedAccountNotFound) Code() int {
	return 404
}

func (o *DeleteManagedAccountNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /managed_accounts/{managed_account_token}][%d] deleteManagedAccountNotFound %s", 404, payload)
}

func (o *DeleteManagedAccountNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /managed_accounts/{managed_account_token}][%d] deleteManagedAccountNotFound %s", 404, payload)
}

func (o *DeleteManagedAccountNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteManagedAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
