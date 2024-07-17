// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev1/models"
)

// GetIntegrationsAWSReader is a Reader for the GetIntegrationsAWS structure.
type GetIntegrationsAWSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsAWSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsAWSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIntegrationsAWSNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /integrations/aws/{access_credential_id}] getIntegrationsAWS", response, response.Code())
	}
}

// NewGetIntegrationsAWSOK creates a GetIntegrationsAWSOK with default headers values
func NewGetIntegrationsAWSOK() *GetIntegrationsAWSOK {
	return &GetIntegrationsAWSOK{}
}

/*
GetIntegrationsAWSOK describes a response with status code 200, with default header values.

GetIntegrationsAWSOK get integrations a w s o k
*/
type GetIntegrationsAWSOK struct {
	Payload *models.AwsAccessCredential
}

// IsSuccess returns true when this get integrations a w s o k response has a 2xx status code
func (o *GetIntegrationsAWSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations a w s o k response has a 3xx status code
func (o *GetIntegrationsAWSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations a w s o k response has a 4xx status code
func (o *GetIntegrationsAWSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations a w s o k response has a 5xx status code
func (o *GetIntegrationsAWSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations a w s o k response a status code equal to that given
func (o *GetIntegrationsAWSOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get integrations a w s o k response
func (o *GetIntegrationsAWSOK) Code() int {
	return 200
}

func (o *GetIntegrationsAWSOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /integrations/aws/{access_credential_id}][%d] getIntegrationsAWSOK %s", 200, payload)
}

func (o *GetIntegrationsAWSOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /integrations/aws/{access_credential_id}][%d] getIntegrationsAWSOK %s", 200, payload)
}

func (o *GetIntegrationsAWSOK) GetPayload() *models.AwsAccessCredential {
	return o.Payload
}

func (o *GetIntegrationsAWSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsAccessCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsAWSNotFound creates a GetIntegrationsAWSNotFound with default headers values
func NewGetIntegrationsAWSNotFound() *GetIntegrationsAWSNotFound {
	return &GetIntegrationsAWSNotFound{}
}

/*
GetIntegrationsAWSNotFound describes a response with status code 404, with default header values.

GetIntegrationsAWSNotFound get integrations a w s not found
*/
type GetIntegrationsAWSNotFound struct {
}

// IsSuccess returns true when this get integrations a w s not found response has a 2xx status code
func (o *GetIntegrationsAWSNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations a w s not found response has a 3xx status code
func (o *GetIntegrationsAWSNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations a w s not found response has a 4xx status code
func (o *GetIntegrationsAWSNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations a w s not found response has a 5xx status code
func (o *GetIntegrationsAWSNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations a w s not found response a status code equal to that given
func (o *GetIntegrationsAWSNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get integrations a w s not found response
func (o *GetIntegrationsAWSNotFound) Code() int {
	return 404
}

func (o *GetIntegrationsAWSNotFound) Error() string {
	return fmt.Sprintf("[GET /integrations/aws/{access_credential_id}][%d] getIntegrationsAWSNotFound", 404)
}

func (o *GetIntegrationsAWSNotFound) String() string {
	return fmt.Sprintf("[GET /integrations/aws/{access_credential_id}][%d] getIntegrationsAWSNotFound", 404)
}

func (o *GetIntegrationsAWSNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
