// Code generated by go-swagger; DO NOT EDIT.

package anomaly_alerts

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

// GetAnomalyAlertsReader is a Reader for the GetAnomalyAlerts structure.
type GetAnomalyAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnomalyAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnomalyAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /anomaly_alerts] getAnomalyAlerts", response, response.Code())
	}
}

// NewGetAnomalyAlertsOK creates a GetAnomalyAlertsOK with default headers values
func NewGetAnomalyAlertsOK() *GetAnomalyAlertsOK {
	return &GetAnomalyAlertsOK{}
}

/*
GetAnomalyAlertsOK describes a response with status code 200, with default header values.

GetAnomalyAlertsOK get anomaly alerts o k
*/
type GetAnomalyAlertsOK struct {
	Payload *models.AnomalyAlerts
}

// IsSuccess returns true when this get anomaly alerts o k response has a 2xx status code
func (o *GetAnomalyAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get anomaly alerts o k response has a 3xx status code
func (o *GetAnomalyAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get anomaly alerts o k response has a 4xx status code
func (o *GetAnomalyAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get anomaly alerts o k response has a 5xx status code
func (o *GetAnomalyAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get anomaly alerts o k response a status code equal to that given
func (o *GetAnomalyAlertsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get anomaly alerts o k response
func (o *GetAnomalyAlertsOK) Code() int {
	return 200
}

func (o *GetAnomalyAlertsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /anomaly_alerts][%d] getAnomalyAlertsOK %s", 200, payload)
}

func (o *GetAnomalyAlertsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /anomaly_alerts][%d] getAnomalyAlertsOK %s", 200, payload)
}

func (o *GetAnomalyAlertsOK) GetPayload() *models.AnomalyAlerts {
	return o.Payload
}

func (o *GetAnomalyAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnomalyAlerts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
