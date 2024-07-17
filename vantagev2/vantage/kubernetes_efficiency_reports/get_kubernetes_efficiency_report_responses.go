// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_efficiency_reports

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

// GetKubernetesEfficiencyReportReader is a Reader for the GetKubernetesEfficiencyReport structure.
type GetKubernetesEfficiencyReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesEfficiencyReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesEfficiencyReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetKubernetesEfficiencyReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes_efficiency_reports/{kubernetes_efficiency_report_token}] getKubernetesEfficiencyReport", response, response.Code())
	}
}

// NewGetKubernetesEfficiencyReportOK creates a GetKubernetesEfficiencyReportOK with default headers values
func NewGetKubernetesEfficiencyReportOK() *GetKubernetesEfficiencyReportOK {
	return &GetKubernetesEfficiencyReportOK{}
}

/*
GetKubernetesEfficiencyReportOK describes a response with status code 200, with default header values.

GetKubernetesEfficiencyReportOK get kubernetes efficiency report o k
*/
type GetKubernetesEfficiencyReportOK struct {
	Payload *models.KubernetesEfficiencyReport
}

// IsSuccess returns true when this get kubernetes efficiency report o k response has a 2xx status code
func (o *GetKubernetesEfficiencyReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes efficiency report o k response has a 3xx status code
func (o *GetKubernetesEfficiencyReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes efficiency report o k response has a 4xx status code
func (o *GetKubernetesEfficiencyReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes efficiency report o k response has a 5xx status code
func (o *GetKubernetesEfficiencyReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes efficiency report o k response a status code equal to that given
func (o *GetKubernetesEfficiencyReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kubernetes efficiency report o k response
func (o *GetKubernetesEfficiencyReportOK) Code() int {
	return 200
}

func (o *GetKubernetesEfficiencyReportOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes_efficiency_reports/{kubernetes_efficiency_report_token}][%d] getKubernetesEfficiencyReportOK %s", 200, payload)
}

func (o *GetKubernetesEfficiencyReportOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes_efficiency_reports/{kubernetes_efficiency_report_token}][%d] getKubernetesEfficiencyReportOK %s", 200, payload)
}

func (o *GetKubernetesEfficiencyReportOK) GetPayload() *models.KubernetesEfficiencyReport {
	return o.Payload
}

func (o *GetKubernetesEfficiencyReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesEfficiencyReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesEfficiencyReportNotFound creates a GetKubernetesEfficiencyReportNotFound with default headers values
func NewGetKubernetesEfficiencyReportNotFound() *GetKubernetesEfficiencyReportNotFound {
	return &GetKubernetesEfficiencyReportNotFound{}
}

/*
GetKubernetesEfficiencyReportNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetKubernetesEfficiencyReportNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this get kubernetes efficiency report not found response has a 2xx status code
func (o *GetKubernetesEfficiencyReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes efficiency report not found response has a 3xx status code
func (o *GetKubernetesEfficiencyReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes efficiency report not found response has a 4xx status code
func (o *GetKubernetesEfficiencyReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes efficiency report not found response has a 5xx status code
func (o *GetKubernetesEfficiencyReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes efficiency report not found response a status code equal to that given
func (o *GetKubernetesEfficiencyReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get kubernetes efficiency report not found response
func (o *GetKubernetesEfficiencyReportNotFound) Code() int {
	return 404
}

func (o *GetKubernetesEfficiencyReportNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes_efficiency_reports/{kubernetes_efficiency_report_token}][%d] getKubernetesEfficiencyReportNotFound %s", 404, payload)
}

func (o *GetKubernetesEfficiencyReportNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes_efficiency_reports/{kubernetes_efficiency_report_token}][%d] getKubernetesEfficiencyReportNotFound %s", 404, payload)
}

func (o *GetKubernetesEfficiencyReportNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetKubernetesEfficiencyReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
