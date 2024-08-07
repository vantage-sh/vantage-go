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

// CreateTeamReader is a Reader for the CreateTeam structure.
type CreateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTeamCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /teams] createTeam", response, response.Code())
	}
}

// NewCreateTeamCreated creates a CreateTeamCreated with default headers values
func NewCreateTeamCreated() *CreateTeamCreated {
	return &CreateTeamCreated{}
}

/*
CreateTeamCreated describes a response with status code 201, with default header values.

CreateTeamCreated create team created
*/
type CreateTeamCreated struct {
	Payload *models.Team
}

// IsSuccess returns true when this create team created response has a 2xx status code
func (o *CreateTeamCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create team created response has a 3xx status code
func (o *CreateTeamCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team created response has a 4xx status code
func (o *CreateTeamCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create team created response has a 5xx status code
func (o *CreateTeamCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create team created response a status code equal to that given
func (o *CreateTeamCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create team created response
func (o *CreateTeamCreated) Code() int {
	return 201
}

func (o *CreateTeamCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamCreated %s", 201, payload)
}

func (o *CreateTeamCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamCreated %s", 201, payload)
}

func (o *CreateTeamCreated) GetPayload() *models.Team {
	return o.Payload
}

func (o *CreateTeamCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamBadRequest creates a CreateTeamBadRequest with default headers values
func NewCreateTeamBadRequest() *CreateTeamBadRequest {
	return &CreateTeamBadRequest{}
}

/*
CreateTeamBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateTeamBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create team bad request response has a 2xx status code
func (o *CreateTeamBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create team bad request response has a 3xx status code
func (o *CreateTeamBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team bad request response has a 4xx status code
func (o *CreateTeamBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create team bad request response has a 5xx status code
func (o *CreateTeamBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create team bad request response a status code equal to that given
func (o *CreateTeamBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create team bad request response
func (o *CreateTeamBadRequest) Code() int {
	return 400
}

func (o *CreateTeamBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamBadRequest %s", 400, payload)
}

func (o *CreateTeamBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamBadRequest %s", 400, payload)
}

func (o *CreateTeamBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamForbidden creates a CreateTeamForbidden with default headers values
func NewCreateTeamForbidden() *CreateTeamForbidden {
	return &CreateTeamForbidden{}
}

/*
CreateTeamForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateTeamForbidden struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create team forbidden response has a 2xx status code
func (o *CreateTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create team forbidden response has a 3xx status code
func (o *CreateTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team forbidden response has a 4xx status code
func (o *CreateTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create team forbidden response has a 5xx status code
func (o *CreateTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create team forbidden response a status code equal to that given
func (o *CreateTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create team forbidden response
func (o *CreateTeamForbidden) Code() int {
	return 403
}

func (o *CreateTeamForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamForbidden %s", 403, payload)
}

func (o *CreateTeamForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamForbidden %s", 403, payload)
}

func (o *CreateTeamForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamNotFound creates a CreateTeamNotFound with default headers values
func NewCreateTeamNotFound() *CreateTeamNotFound {
	return &CreateTeamNotFound{}
}

/*
CreateTeamNotFound describes a response with status code 404, with default header values.

NotFound
*/
type CreateTeamNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create team not found response has a 2xx status code
func (o *CreateTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create team not found response has a 3xx status code
func (o *CreateTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team not found response has a 4xx status code
func (o *CreateTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create team not found response has a 5xx status code
func (o *CreateTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create team not found response a status code equal to that given
func (o *CreateTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create team not found response
func (o *CreateTeamNotFound) Code() int {
	return 404
}

func (o *CreateTeamNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamNotFound %s", 404, payload)
}

func (o *CreateTeamNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamNotFound %s", 404, payload)
}

func (o *CreateTeamNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamUnprocessableEntity creates a CreateTeamUnprocessableEntity with default headers values
func NewCreateTeamUnprocessableEntity() *CreateTeamUnprocessableEntity {
	return &CreateTeamUnprocessableEntity{}
}

/*
CreateTeamUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type CreateTeamUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create team unprocessable entity response has a 2xx status code
func (o *CreateTeamUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create team unprocessable entity response has a 3xx status code
func (o *CreateTeamUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team unprocessable entity response has a 4xx status code
func (o *CreateTeamUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create team unprocessable entity response has a 5xx status code
func (o *CreateTeamUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create team unprocessable entity response a status code equal to that given
func (o *CreateTeamUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create team unprocessable entity response
func (o *CreateTeamUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateTeamUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamUnprocessableEntity %s", 422, payload)
}

func (o *CreateTeamUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /teams][%d] createTeamUnprocessableEntity %s", 422, payload)
}

func (o *CreateTeamUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
