// Code generated by go-swagger; DO NOT EDIT.

package state_backends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/getlunaform/lunaform/models"
)

// NewUpdateStateBackendParams creates a new UpdateStateBackendParams object
// with the default values initialized.
func NewUpdateStateBackendParams() *UpdateStateBackendParams {
	var ()
	return &UpdateStateBackendParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateStateBackendParamsWithTimeout creates a new UpdateStateBackendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateStateBackendParamsWithTimeout(timeout time.Duration) *UpdateStateBackendParams {
	var ()
	return &UpdateStateBackendParams{

		timeout: timeout,
	}
}

// NewUpdateStateBackendParamsWithContext creates a new UpdateStateBackendParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateStateBackendParamsWithContext(ctx context.Context) *UpdateStateBackendParams {
	var ()
	return &UpdateStateBackendParams{

		Context: ctx,
	}
}

// NewUpdateStateBackendParamsWithHTTPClient creates a new UpdateStateBackendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateStateBackendParamsWithHTTPClient(client *http.Client) *UpdateStateBackendParams {
	var ()
	return &UpdateStateBackendParams{
		HTTPClient: client,
	}
}

/*UpdateStateBackendParams contains all the parameters to send to the API endpoint
for the update state backend operation typically these are written to a http.Request
*/
type UpdateStateBackendParams struct {

	/*ID
	  ID of a terraform state backend

	*/
	ID string
	/*TerraformStateBackend
	  A terraform state backend

	*/
	TerraformStateBackend *models.ResourceTfStateBackend

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update state backend params
func (o *UpdateStateBackendParams) WithTimeout(timeout time.Duration) *UpdateStateBackendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update state backend params
func (o *UpdateStateBackendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update state backend params
func (o *UpdateStateBackendParams) WithContext(ctx context.Context) *UpdateStateBackendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update state backend params
func (o *UpdateStateBackendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update state backend params
func (o *UpdateStateBackendParams) WithHTTPClient(client *http.Client) *UpdateStateBackendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update state backend params
func (o *UpdateStateBackendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update state backend params
func (o *UpdateStateBackendParams) WithID(id string) *UpdateStateBackendParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update state backend params
func (o *UpdateStateBackendParams) SetID(id string) {
	o.ID = id
}

// WithTerraformStateBackend adds the terraformStateBackend to the update state backend params
func (o *UpdateStateBackendParams) WithTerraformStateBackend(terraformStateBackend *models.ResourceTfStateBackend) *UpdateStateBackendParams {
	o.SetTerraformStateBackend(terraformStateBackend)
	return o
}

// SetTerraformStateBackend adds the terraformStateBackend to the update state backend params
func (o *UpdateStateBackendParams) SetTerraformStateBackend(terraformStateBackend *models.ResourceTfStateBackend) {
	o.TerraformStateBackend = terraformStateBackend
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateStateBackendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.TerraformStateBackend != nil {
		if err := r.SetBodyParam(o.TerraformStateBackend); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}