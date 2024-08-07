// Code generated by go-swagger; DO NOT EDIT.

package billing_rules

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

// GetBillingRuleReader is a Reader for the GetBillingRule structure.
type GetBillingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBillingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBillingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBillingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /billing_rules/{billing_rule_token}] getBillingRule", response, response.Code())
	}
}

// NewGetBillingRuleOK creates a GetBillingRuleOK with default headers values
func NewGetBillingRuleOK() *GetBillingRuleOK {
	return &GetBillingRuleOK{}
}

/*
GetBillingRuleOK describes a response with status code 200, with default header values.

GetBillingRuleOK get billing rule o k
*/
type GetBillingRuleOK struct {
	Payload *models.BillingRule
}

// IsSuccess returns true when this get billing rule o k response has a 2xx status code
func (o *GetBillingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get billing rule o k response has a 3xx status code
func (o *GetBillingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing rule o k response has a 4xx status code
func (o *GetBillingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get billing rule o k response has a 5xx status code
func (o *GetBillingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing rule o k response a status code equal to that given
func (o *GetBillingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get billing rule o k response
func (o *GetBillingRuleOK) Code() int {
	return 200
}

func (o *GetBillingRuleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing_rules/{billing_rule_token}][%d] getBillingRuleOK %s", 200, payload)
}

func (o *GetBillingRuleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing_rules/{billing_rule_token}][%d] getBillingRuleOK %s", 200, payload)
}

func (o *GetBillingRuleOK) GetPayload() *models.BillingRule {
	return o.Payload
}

func (o *GetBillingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingRuleNotFound creates a GetBillingRuleNotFound with default headers values
func NewGetBillingRuleNotFound() *GetBillingRuleNotFound {
	return &GetBillingRuleNotFound{}
}

/*
GetBillingRuleNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetBillingRuleNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this get billing rule not found response has a 2xx status code
func (o *GetBillingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing rule not found response has a 3xx status code
func (o *GetBillingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing rule not found response has a 4xx status code
func (o *GetBillingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing rule not found response has a 5xx status code
func (o *GetBillingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing rule not found response a status code equal to that given
func (o *GetBillingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get billing rule not found response
func (o *GetBillingRuleNotFound) Code() int {
	return 404
}

func (o *GetBillingRuleNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing_rules/{billing_rule_token}][%d] getBillingRuleNotFound %s", 404, payload)
}

func (o *GetBillingRuleNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /billing_rules/{billing_rule_token}][%d] getBillingRuleNotFound %s", 404, payload)
}

func (o *GetBillingRuleNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetBillingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
