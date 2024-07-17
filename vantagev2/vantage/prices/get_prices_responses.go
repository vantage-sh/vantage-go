// Code generated by go-swagger; DO NOT EDIT.

package prices

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

// GetPricesReader is a Reader for the GetPrices structure.
type GetPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /products/{product_id}/prices] getPrices", response, response.Code())
	}
}

// NewGetPricesOK creates a GetPricesOK with default headers values
func NewGetPricesOK() *GetPricesOK {
	return &GetPricesOK{}
}

/*
GetPricesOK describes a response with status code 200, with default header values.

GetPricesOK get prices o k
*/
type GetPricesOK struct {
	Payload *models.Prices
}

// IsSuccess returns true when this get prices o k response has a 2xx status code
func (o *GetPricesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get prices o k response has a 3xx status code
func (o *GetPricesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get prices o k response has a 4xx status code
func (o *GetPricesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get prices o k response has a 5xx status code
func (o *GetPricesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get prices o k response a status code equal to that given
func (o *GetPricesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get prices o k response
func (o *GetPricesOK) Code() int {
	return 200
}

func (o *GetPricesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /products/{product_id}/prices][%d] getPricesOK %s", 200, payload)
}

func (o *GetPricesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /products/{product_id}/prices][%d] getPricesOK %s", 200, payload)
}

func (o *GetPricesOK) GetPayload() *models.Prices {
	return o.Payload
}

func (o *GetPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
