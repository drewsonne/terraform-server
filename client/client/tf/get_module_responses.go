// Code generated by go-swagger; DO NOT EDIT.

package tf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetModuleReader is a Reader for the GetModule structure.
type GetModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetModuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetModuleOK creates a GetModuleOK with default headers values
func NewGetModuleOK() *GetModuleOK {
	return &GetModuleOK{}
}

/*GetModuleOK handles this case with default header values.

OK
*/
type GetModuleOK struct {
}

func (o *GetModuleOK) Error() string {
	return fmt.Sprintf("[GET /tf/module/{id}][%d] getModuleOK ", 200)
}

func (o *GetModuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
