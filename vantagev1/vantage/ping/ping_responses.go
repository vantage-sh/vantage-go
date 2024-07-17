// Code generated by go-swagger; DO NOT EDIT.

package ping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PingReader is a Reader for the Ping structure.
type PingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ping] ping", response, response.Code())
	}
}

// NewPingOK creates a PingOK with default headers values
func NewPingOK() *PingOK {
	return &PingOK{}
}

/*
PingOK describes a response with status code 200, with default header values.

This is a health check endpoint that can be used to determine Vantage API healthiness. It will return 200 if everything is running smoothly.
*/
type PingOK struct {
}

// IsSuccess returns true when this ping o k response has a 2xx status code
func (o *PingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ping o k response has a 3xx status code
func (o *PingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping o k response has a 4xx status code
func (o *PingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping o k response has a 5xx status code
func (o *PingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ping o k response a status code equal to that given
func (o *PingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ping o k response
func (o *PingOK) Code() int {
	return 200
}

func (o *PingOK) Error() string {
	return fmt.Sprintf("[GET /ping][%d] pingOK", 200)
}

func (o *PingOK) String() string {
	return fmt.Sprintf("[GET /ping][%d] pingOK", 200)
}

func (o *PingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
