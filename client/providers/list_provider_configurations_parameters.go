// Code generated by go-swagger; DO NOT EDIT.

package providers

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
)

// NewListProviderConfigurationsParams creates a new ListProviderConfigurationsParams object
// with the default values initialized.
func NewListProviderConfigurationsParams() *ListProviderConfigurationsParams {
	var ()
	return &ListProviderConfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListProviderConfigurationsParamsWithTimeout creates a new ListProviderConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListProviderConfigurationsParamsWithTimeout(timeout time.Duration) *ListProviderConfigurationsParams {
	var ()
	return &ListProviderConfigurationsParams{

		timeout: timeout,
	}
}

// NewListProviderConfigurationsParamsWithContext creates a new ListProviderConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListProviderConfigurationsParamsWithContext(ctx context.Context) *ListProviderConfigurationsParams {
	var ()
	return &ListProviderConfigurationsParams{

		Context: ctx,
	}
}

// NewListProviderConfigurationsParamsWithHTTPClient creates a new ListProviderConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListProviderConfigurationsParamsWithHTTPClient(client *http.Client) *ListProviderConfigurationsParams {
	var ()
	return &ListProviderConfigurationsParams{
		HTTPClient: client,
	}
}

/*ListProviderConfigurationsParams contains all the parameters to send to the API endpoint
for the list provider configurations operation typically these are written to a http.Request
*/
type ListProviderConfigurationsParams struct {

	/*ProviderName
	  Terraform Provider Name

	*/
	ProviderName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list provider configurations params
func (o *ListProviderConfigurationsParams) WithTimeout(timeout time.Duration) *ListProviderConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list provider configurations params
func (o *ListProviderConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list provider configurations params
func (o *ListProviderConfigurationsParams) WithContext(ctx context.Context) *ListProviderConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list provider configurations params
func (o *ListProviderConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list provider configurations params
func (o *ListProviderConfigurationsParams) WithHTTPClient(client *http.Client) *ListProviderConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list provider configurations params
func (o *ListProviderConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProviderName adds the providerName to the list provider configurations params
func (o *ListProviderConfigurationsParams) WithProviderName(providerName string) *ListProviderConfigurationsParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the list provider configurations params
func (o *ListProviderConfigurationsParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WriteToRequest writes these params to a swagger request
func (o *ListProviderConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param provider-name
	if err := r.SetPathParam("provider-name", o.ProviderName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
