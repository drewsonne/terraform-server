// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/getlunaform/lunaform/models"
)

// ListDeploymentsReader is a Reader for the ListDeployments structure.
type ListDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDeploymentsOK creates a ListDeploymentsOK with default headers values
func NewListDeploymentsOK() *ListDeploymentsOK {
	return &ListDeploymentsOK{}
}

/*ListDeploymentsOK handles this case with default header values.

OK
*/
type ListDeploymentsOK struct {
	Payload *models.ResponseListTfDeployments
}

func (o *ListDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /tf/stack/{id}/deployments][%d] listDeploymentsOK  %+v", 200, o.Payload)
}

func (o *ListDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseListTfDeployments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
