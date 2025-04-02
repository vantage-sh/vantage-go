// Code generated by go-swagger; DO NOT EDIT.

package costs

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

// CreateCostExportReader is a Reader for the CreateCostExport structure.
type CreateCostExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCostExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateCostExportAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCostExportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewCreateCostExportPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCostExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /costs/data_exports] createCostExport", response, response.Code())
	}
}

// NewCreateCostExportAccepted creates a CreateCostExportAccepted with default headers values
func NewCreateCostExportAccepted() *CreateCostExportAccepted {
	return &CreateCostExportAccepted{}
}

/*
CreateCostExportAccepted describes a response with status code 202, with default header values.

The data export has been queued and will be available at the location specified in the Location header.
*/
type CreateCostExportAccepted struct {
}

// IsSuccess returns true when this create cost export accepted response has a 2xx status code
func (o *CreateCostExportAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cost export accepted response has a 3xx status code
func (o *CreateCostExportAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cost export accepted response has a 4xx status code
func (o *CreateCostExportAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cost export accepted response has a 5xx status code
func (o *CreateCostExportAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create cost export accepted response a status code equal to that given
func (o *CreateCostExportAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the create cost export accepted response
func (o *CreateCostExportAccepted) Code() int {
	return 202
}

func (o *CreateCostExportAccepted) Error() string {
	return fmt.Sprintf("[POST /costs/data_exports][%d] createCostExportAccepted", 202)
}

func (o *CreateCostExportAccepted) String() string {
	return fmt.Sprintf("[POST /costs/data_exports][%d] createCostExportAccepted", 202)
}

func (o *CreateCostExportAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCostExportBadRequest creates a CreateCostExportBadRequest with default headers values
func NewCreateCostExportBadRequest() *CreateCostExportBadRequest {
	return &CreateCostExportBadRequest{}
}

/*
CreateCostExportBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateCostExportBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create cost export bad request response has a 2xx status code
func (o *CreateCostExportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cost export bad request response has a 3xx status code
func (o *CreateCostExportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cost export bad request response has a 4xx status code
func (o *CreateCostExportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cost export bad request response has a 5xx status code
func (o *CreateCostExportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create cost export bad request response a status code equal to that given
func (o *CreateCostExportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create cost export bad request response
func (o *CreateCostExportBadRequest) Code() int {
	return 400
}

func (o *CreateCostExportBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /costs/data_exports][%d] createCostExportBadRequest %s", 400, payload)
}

func (o *CreateCostExportBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /costs/data_exports][%d] createCostExportBadRequest %s", 400, payload)
}

func (o *CreateCostExportBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateCostExportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCostExportPaymentRequired creates a CreateCostExportPaymentRequired with default headers values
func NewCreateCostExportPaymentRequired() *CreateCostExportPaymentRequired {
	return &CreateCostExportPaymentRequired{}
}

/*
CreateCostExportPaymentRequired describes a response with status code 402, with default header values.

PaymentRequired
*/
type CreateCostExportPaymentRequired struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create cost export payment required response has a 2xx status code
func (o *CreateCostExportPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cost export payment required response has a 3xx status code
func (o *CreateCostExportPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cost export payment required response has a 4xx status code
func (o *CreateCostExportPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cost export payment required response has a 5xx status code
func (o *CreateCostExportPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this create cost export payment required response a status code equal to that given
func (o *CreateCostExportPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the create cost export payment required response
func (o *CreateCostExportPaymentRequired) Code() int {
	return 402
}

func (o *CreateCostExportPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /costs/data_exports][%d] createCostExportPaymentRequired %s", 402, payload)
}

func (o *CreateCostExportPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /costs/data_exports][%d] createCostExportPaymentRequired %s", 402, payload)
}

func (o *CreateCostExportPaymentRequired) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateCostExportPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCostExportNotFound creates a CreateCostExportNotFound with default headers values
func NewCreateCostExportNotFound() *CreateCostExportNotFound {
	return &CreateCostExportNotFound{}
}

/*
CreateCostExportNotFound describes a response with status code 404, with default header values.

NotFound
*/
type CreateCostExportNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create cost export not found response has a 2xx status code
func (o *CreateCostExportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cost export not found response has a 3xx status code
func (o *CreateCostExportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cost export not found response has a 4xx status code
func (o *CreateCostExportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cost export not found response has a 5xx status code
func (o *CreateCostExportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create cost export not found response a status code equal to that given
func (o *CreateCostExportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create cost export not found response
func (o *CreateCostExportNotFound) Code() int {
	return 404
}

func (o *CreateCostExportNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /costs/data_exports][%d] createCostExportNotFound %s", 404, payload)
}

func (o *CreateCostExportNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /costs/data_exports][%d] createCostExportNotFound %s", 404, payload)
}

func (o *CreateCostExportNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateCostExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
