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

// GetFinancialCommitmentReportReader is a Reader for the GetFinancialCommitmentReport structure.
type GetFinancialCommitmentReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFinancialCommitmentReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFinancialCommitmentReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetFinancialCommitmentReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /financial_commitment_reports/{financial_commitment_report_token}] getFinancialCommitmentReport", response, response.Code())
	}
}

// NewGetFinancialCommitmentReportOK creates a GetFinancialCommitmentReportOK with default headers values
func NewGetFinancialCommitmentReportOK() *GetFinancialCommitmentReportOK {
	return &GetFinancialCommitmentReportOK{}
}

/*
GetFinancialCommitmentReportOK describes a response with status code 200, with default header values.

GetFinancialCommitmentReportOK get financial commitment report o k
*/
type GetFinancialCommitmentReportOK struct {
	Payload *models.FinancialCommitmentReport
}

// IsSuccess returns true when this get financial commitment report o k response has a 2xx status code
func (o *GetFinancialCommitmentReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get financial commitment report o k response has a 3xx status code
func (o *GetFinancialCommitmentReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get financial commitment report o k response has a 4xx status code
func (o *GetFinancialCommitmentReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get financial commitment report o k response has a 5xx status code
func (o *GetFinancialCommitmentReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get financial commitment report o k response a status code equal to that given
func (o *GetFinancialCommitmentReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get financial commitment report o k response
func (o *GetFinancialCommitmentReportOK) Code() int {
	return 200
}

func (o *GetFinancialCommitmentReportOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /financial_commitment_reports/{financial_commitment_report_token}][%d] getFinancialCommitmentReportOK %s", 200, payload)
}

func (o *GetFinancialCommitmentReportOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /financial_commitment_reports/{financial_commitment_report_token}][%d] getFinancialCommitmentReportOK %s", 200, payload)
}

func (o *GetFinancialCommitmentReportOK) GetPayload() *models.FinancialCommitmentReport {
	return o.Payload
}

func (o *GetFinancialCommitmentReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FinancialCommitmentReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFinancialCommitmentReportNotFound creates a GetFinancialCommitmentReportNotFound with default headers values
func NewGetFinancialCommitmentReportNotFound() *GetFinancialCommitmentReportNotFound {
	return &GetFinancialCommitmentReportNotFound{}
}

/*
GetFinancialCommitmentReportNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetFinancialCommitmentReportNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this get financial commitment report not found response has a 2xx status code
func (o *GetFinancialCommitmentReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get financial commitment report not found response has a 3xx status code
func (o *GetFinancialCommitmentReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get financial commitment report not found response has a 4xx status code
func (o *GetFinancialCommitmentReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get financial commitment report not found response has a 5xx status code
func (o *GetFinancialCommitmentReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get financial commitment report not found response a status code equal to that given
func (o *GetFinancialCommitmentReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get financial commitment report not found response
func (o *GetFinancialCommitmentReportNotFound) Code() int {
	return 404
}

func (o *GetFinancialCommitmentReportNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /financial_commitment_reports/{financial_commitment_report_token}][%d] getFinancialCommitmentReportNotFound %s", 404, payload)
}

func (o *GetFinancialCommitmentReportNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /financial_commitment_reports/{financial_commitment_report_token}][%d] getFinancialCommitmentReportNotFound %s", 404, payload)
}

func (o *GetFinancialCommitmentReportNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetFinancialCommitmentReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
