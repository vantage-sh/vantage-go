// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// GetTeamsReader is a Reader for the GetTeams structure.
type GetTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /teams] getTeams", response, response.Code())
	}
}

// NewGetTeamsOK creates a GetTeamsOK with default headers values
func NewGetTeamsOK() *GetTeamsOK {
	return &GetTeamsOK{}
}

/*
GetTeamsOK describes a response with status code 200, with default header values.

GetTeamsOK get teams o k
*/
type GetTeamsOK struct {
	Payload *models.Teams
}

// IsSuccess returns true when this get teams o k response has a 2xx status code
func (o *GetTeamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get teams o k response has a 3xx status code
func (o *GetTeamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get teams o k response has a 4xx status code
func (o *GetTeamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get teams o k response has a 5xx status code
func (o *GetTeamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get teams o k response a status code equal to that given
func (o *GetTeamsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get teams o k response
func (o *GetTeamsOK) Code() int {
	return 200
}

func (o *GetTeamsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams][%d] getTeamsOK %s", 200, payload)
}

func (o *GetTeamsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams][%d] getTeamsOK %s", 200, payload)
}

func (o *GetTeamsOK) GetPayload() *models.Teams {
	return o.Payload
}

func (o *GetTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Teams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
