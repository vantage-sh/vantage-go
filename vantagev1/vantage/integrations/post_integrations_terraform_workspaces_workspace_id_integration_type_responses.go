// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeReader is a Reader for the PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationType structure.
type PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /integrations/terraform/workspaces/{workspace_id}/{integration_type}] postIntegrationsTerraformWorkspacesWorkspaceIdIntegrationType", response, response.Code())
	}
}

// NewPostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated creates a PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated with default headers values
func NewPostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated() *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated {
	return &PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated{}
}

/*
PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated describes a response with status code 201, with default header values.

created Workspace
*/
type PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated struct {
}

// IsSuccess returns true when this post integrations terraform workspaces workspace Id integration type created response has a 2xx status code
func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post integrations terraform workspaces workspace Id integration type created response has a 3xx status code
func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations terraform workspaces workspace Id integration type created response has a 4xx status code
func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations terraform workspaces workspace Id integration type created response has a 5xx status code
func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations terraform workspaces workspace Id integration type created response a status code equal to that given
func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post integrations terraform workspaces workspace Id integration type created response
func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) Code() int {
	return 201
}

func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) Error() string {
	return fmt.Sprintf("[POST /integrations/terraform/workspaces/{workspace_id}/{integration_type}][%d] postIntegrationsTerraformWorkspacesWorkspaceIdIntegrationTypeCreated", 201)
}

func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) String() string {
	return fmt.Sprintf("[POST /integrations/terraform/workspaces/{workspace_id}/{integration_type}][%d] postIntegrationsTerraformWorkspacesWorkspaceIdIntegrationTypeCreated", 201)
}

func (o *PostIntegrationsTerraformWorkspacesWorkspaceIDIntegrationTypeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
