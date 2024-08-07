// Code generated by go-swagger; DO NOT EDIT.

package budget_alerts

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

// DeleteBudgetAlertReader is a Reader for the DeleteBudgetAlert structure.
type DeleteBudgetAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBudgetAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBudgetAlertNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteBudgetAlertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /budget_alerts/{budget_alert_token}] deleteBudgetAlert", response, response.Code())
	}
}

// NewDeleteBudgetAlertNoContent creates a DeleteBudgetAlertNoContent with default headers values
func NewDeleteBudgetAlertNoContent() *DeleteBudgetAlertNoContent {
	return &DeleteBudgetAlertNoContent{}
}

/*
DeleteBudgetAlertNoContent describes a response with status code 204, with default header values.

DeleteBudgetAlertNoContent delete budget alert no content
*/
type DeleteBudgetAlertNoContent struct {
	Payload *models.BudgetAlert
}

// IsSuccess returns true when this delete budget alert no content response has a 2xx status code
func (o *DeleteBudgetAlertNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete budget alert no content response has a 3xx status code
func (o *DeleteBudgetAlertNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete budget alert no content response has a 4xx status code
func (o *DeleteBudgetAlertNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete budget alert no content response has a 5xx status code
func (o *DeleteBudgetAlertNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete budget alert no content response a status code equal to that given
func (o *DeleteBudgetAlertNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete budget alert no content response
func (o *DeleteBudgetAlertNoContent) Code() int {
	return 204
}

func (o *DeleteBudgetAlertNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /budget_alerts/{budget_alert_token}][%d] deleteBudgetAlertNoContent %s", 204, payload)
}

func (o *DeleteBudgetAlertNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /budget_alerts/{budget_alert_token}][%d] deleteBudgetAlertNoContent %s", 204, payload)
}

func (o *DeleteBudgetAlertNoContent) GetPayload() *models.BudgetAlert {
	return o.Payload
}

func (o *DeleteBudgetAlertNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BudgetAlert)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBudgetAlertNotFound creates a DeleteBudgetAlertNotFound with default headers values
func NewDeleteBudgetAlertNotFound() *DeleteBudgetAlertNotFound {
	return &DeleteBudgetAlertNotFound{}
}

/*
DeleteBudgetAlertNotFound describes a response with status code 404, with default header values.

NotFound
*/
type DeleteBudgetAlertNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this delete budget alert not found response has a 2xx status code
func (o *DeleteBudgetAlertNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete budget alert not found response has a 3xx status code
func (o *DeleteBudgetAlertNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete budget alert not found response has a 4xx status code
func (o *DeleteBudgetAlertNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete budget alert not found response has a 5xx status code
func (o *DeleteBudgetAlertNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete budget alert not found response a status code equal to that given
func (o *DeleteBudgetAlertNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete budget alert not found response
func (o *DeleteBudgetAlertNotFound) Code() int {
	return 404
}

func (o *DeleteBudgetAlertNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /budget_alerts/{budget_alert_token}][%d] deleteBudgetAlertNotFound %s", 404, payload)
}

func (o *DeleteBudgetAlertNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /budget_alerts/{budget_alert_token}][%d] deleteBudgetAlertNotFound %s", 404, payload)
}

func (o *DeleteBudgetAlertNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteBudgetAlertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
