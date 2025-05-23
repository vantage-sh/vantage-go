// Code generated by go-swagger; DO NOT EDIT.

package financial_commitment_reports

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

// UpdateFinancialCommitmentReportReader is a Reader for the UpdateFinancialCommitmentReport structure.
type UpdateFinancialCommitmentReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFinancialCommitmentReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFinancialCommitmentReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFinancialCommitmentReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateFinancialCommitmentReportUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /financial_commitment_reports/{financial_commitment_report_token}] updateFinancialCommitmentReport", response, response.Code())
	}
}

// NewUpdateFinancialCommitmentReportOK creates a UpdateFinancialCommitmentReportOK with default headers values
func NewUpdateFinancialCommitmentReportOK() *UpdateFinancialCommitmentReportOK {
	return &UpdateFinancialCommitmentReportOK{}
}

/*
UpdateFinancialCommitmentReportOK describes a response with status code 200, with default header values.

UpdateFinancialCommitmentReportOK update financial commitment report o k
*/
type UpdateFinancialCommitmentReportOK struct {
	Payload *models.FinancialCommitmentReport
}

// IsSuccess returns true when this update financial commitment report o k response has a 2xx status code
func (o *UpdateFinancialCommitmentReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update financial commitment report o k response has a 3xx status code
func (o *UpdateFinancialCommitmentReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update financial commitment report o k response has a 4xx status code
func (o *UpdateFinancialCommitmentReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update financial commitment report o k response has a 5xx status code
func (o *UpdateFinancialCommitmentReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update financial commitment report o k response a status code equal to that given
func (o *UpdateFinancialCommitmentReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update financial commitment report o k response
func (o *UpdateFinancialCommitmentReportOK) Code() int {
	return 200
}

func (o *UpdateFinancialCommitmentReportOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /financial_commitment_reports/{financial_commitment_report_token}][%d] updateFinancialCommitmentReportOK %s", 200, payload)
}

func (o *UpdateFinancialCommitmentReportOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /financial_commitment_reports/{financial_commitment_report_token}][%d] updateFinancialCommitmentReportOK %s", 200, payload)
}

func (o *UpdateFinancialCommitmentReportOK) GetPayload() *models.FinancialCommitmentReport {
	return o.Payload
}

func (o *UpdateFinancialCommitmentReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FinancialCommitmentReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFinancialCommitmentReportBadRequest creates a UpdateFinancialCommitmentReportBadRequest with default headers values
func NewUpdateFinancialCommitmentReportBadRequest() *UpdateFinancialCommitmentReportBadRequest {
	return &UpdateFinancialCommitmentReportBadRequest{}
}

/*
UpdateFinancialCommitmentReportBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateFinancialCommitmentReportBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update financial commitment report bad request response has a 2xx status code
func (o *UpdateFinancialCommitmentReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update financial commitment report bad request response has a 3xx status code
func (o *UpdateFinancialCommitmentReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update financial commitment report bad request response has a 4xx status code
func (o *UpdateFinancialCommitmentReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update financial commitment report bad request response has a 5xx status code
func (o *UpdateFinancialCommitmentReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update financial commitment report bad request response a status code equal to that given
func (o *UpdateFinancialCommitmentReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update financial commitment report bad request response
func (o *UpdateFinancialCommitmentReportBadRequest) Code() int {
	return 400
}

func (o *UpdateFinancialCommitmentReportBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /financial_commitment_reports/{financial_commitment_report_token}][%d] updateFinancialCommitmentReportBadRequest %s", 400, payload)
}

func (o *UpdateFinancialCommitmentReportBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /financial_commitment_reports/{financial_commitment_report_token}][%d] updateFinancialCommitmentReportBadRequest %s", 400, payload)
}

func (o *UpdateFinancialCommitmentReportBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateFinancialCommitmentReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFinancialCommitmentReportUnprocessableEntity creates a UpdateFinancialCommitmentReportUnprocessableEntity with default headers values
func NewUpdateFinancialCommitmentReportUnprocessableEntity() *UpdateFinancialCommitmentReportUnprocessableEntity {
	return &UpdateFinancialCommitmentReportUnprocessableEntity{}
}

/*
UpdateFinancialCommitmentReportUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type UpdateFinancialCommitmentReportUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update financial commitment report unprocessable entity response has a 2xx status code
func (o *UpdateFinancialCommitmentReportUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update financial commitment report unprocessable entity response has a 3xx status code
func (o *UpdateFinancialCommitmentReportUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update financial commitment report unprocessable entity response has a 4xx status code
func (o *UpdateFinancialCommitmentReportUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update financial commitment report unprocessable entity response has a 5xx status code
func (o *UpdateFinancialCommitmentReportUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update financial commitment report unprocessable entity response a status code equal to that given
func (o *UpdateFinancialCommitmentReportUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update financial commitment report unprocessable entity response
func (o *UpdateFinancialCommitmentReportUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateFinancialCommitmentReportUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /financial_commitment_reports/{financial_commitment_report_token}][%d] updateFinancialCommitmentReportUnprocessableEntity %s", 422, payload)
}

func (o *UpdateFinancialCommitmentReportUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /financial_commitment_reports/{financial_commitment_report_token}][%d] updateFinancialCommitmentReportUnprocessableEntity %s", 422, payload)
}

func (o *UpdateFinancialCommitmentReportUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateFinancialCommitmentReportUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
