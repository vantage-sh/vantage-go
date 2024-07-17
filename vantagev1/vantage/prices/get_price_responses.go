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

	"github.com/vantage-sh/vantage-go/vantagev1/models"
)

// GetPriceReader is a Reader for the GetPrice structure.
type GetPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /products/{product_id}/prices/{id}] getPrice", response, response.Code())
	}
}

// NewGetPriceOK creates a GetPriceOK with default headers values
func NewGetPriceOK() *GetPriceOK {
	return &GetPriceOK{}
}

/*
GetPriceOK describes a response with status code 200, with default header values.

GetPriceOK get price o k
*/
type GetPriceOK struct {
	Payload *models.Price
}

// IsSuccess returns true when this get price o k response has a 2xx status code
func (o *GetPriceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get price o k response has a 3xx status code
func (o *GetPriceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get price o k response has a 4xx status code
func (o *GetPriceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get price o k response has a 5xx status code
func (o *GetPriceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get price o k response a status code equal to that given
func (o *GetPriceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get price o k response
func (o *GetPriceOK) Code() int {
	return 200
}

func (o *GetPriceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /products/{product_id}/prices/{id}][%d] getPriceOK %s", 200, payload)
}

func (o *GetPriceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /products/{product_id}/prices/{id}][%d] getPriceOK %s", 200, payload)
}

func (o *GetPriceOK) GetPayload() *models.Price {
	return o.Payload
}

func (o *GetPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Price)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
