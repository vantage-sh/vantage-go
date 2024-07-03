// Code generated by go-swagger; DO NOT EDIT.

package billing_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// UpdateBillingRuleReader is a Reader for the UpdateBillingRule structure.
type UpdateBillingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBillingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBillingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateBillingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateBillingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /billing_rules/{billing_rule_token}] updateBillingRule", response, response.Code())
	}
}

// NewUpdateBillingRuleOK creates a UpdateBillingRuleOK with default headers values
func NewUpdateBillingRuleOK() *UpdateBillingRuleOK {
	return &UpdateBillingRuleOK{}
}

/*
UpdateBillingRuleOK describes a response with status code 200, with default header values.

UpdateBillingRuleOK update billing rule o k
*/
type UpdateBillingRuleOK struct {
	Payload *models.BillingRule
}

// IsSuccess returns true when this update billing rule o k response has a 2xx status code
func (o *UpdateBillingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update billing rule o k response has a 3xx status code
func (o *UpdateBillingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update billing rule o k response has a 4xx status code
func (o *UpdateBillingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update billing rule o k response has a 5xx status code
func (o *UpdateBillingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update billing rule o k response a status code equal to that given
func (o *UpdateBillingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update billing rule o k response
func (o *UpdateBillingRuleOK) Code() int {
	return 200
}

func (o *UpdateBillingRuleOK) Error() string {
	return fmt.Sprintf("[PUT /billing_rules/{billing_rule_token}][%d] updateBillingRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateBillingRuleOK) String() string {
	return fmt.Sprintf("[PUT /billing_rules/{billing_rule_token}][%d] updateBillingRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateBillingRuleOK) GetPayload() *models.BillingRule {
	return o.Payload
}

func (o *UpdateBillingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBillingRuleBadRequest creates a UpdateBillingRuleBadRequest with default headers values
func NewUpdateBillingRuleBadRequest() *UpdateBillingRuleBadRequest {
	return &UpdateBillingRuleBadRequest{}
}

/*
UpdateBillingRuleBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateBillingRuleBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update billing rule bad request response has a 2xx status code
func (o *UpdateBillingRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update billing rule bad request response has a 3xx status code
func (o *UpdateBillingRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update billing rule bad request response has a 4xx status code
func (o *UpdateBillingRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update billing rule bad request response has a 5xx status code
func (o *UpdateBillingRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update billing rule bad request response a status code equal to that given
func (o *UpdateBillingRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update billing rule bad request response
func (o *UpdateBillingRuleBadRequest) Code() int {
	return 400
}

func (o *UpdateBillingRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /billing_rules/{billing_rule_token}][%d] updateBillingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBillingRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /billing_rules/{billing_rule_token}][%d] updateBillingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBillingRuleBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBillingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBillingRuleNotFound creates a UpdateBillingRuleNotFound with default headers values
func NewUpdateBillingRuleNotFound() *UpdateBillingRuleNotFound {
	return &UpdateBillingRuleNotFound{}
}

/*
UpdateBillingRuleNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateBillingRuleNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update billing rule not found response has a 2xx status code
func (o *UpdateBillingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update billing rule not found response has a 3xx status code
func (o *UpdateBillingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update billing rule not found response has a 4xx status code
func (o *UpdateBillingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update billing rule not found response has a 5xx status code
func (o *UpdateBillingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update billing rule not found response a status code equal to that given
func (o *UpdateBillingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update billing rule not found response
func (o *UpdateBillingRuleNotFound) Code() int {
	return 404
}

func (o *UpdateBillingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /billing_rules/{billing_rule_token}][%d] updateBillingRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBillingRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /billing_rules/{billing_rule_token}][%d] updateBillingRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBillingRuleNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBillingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
