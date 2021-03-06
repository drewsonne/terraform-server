// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"

	hal "github.com/getlunaform/lunaform/models/hal"
)

// ResourceTfStateBackend resource tf state backend
// swagger:model resource-tf-state-backend
type ResourceTfStateBackend struct {

	// links
	Links *hal.HalRscLinks `json:"_links,omitempty"`

	// configuration
	Configuration interface{} `json:"configuration,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this resource tf state backend
func (m *ResourceTfStateBackend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceTfStateBackend) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceTfStateBackend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceTfStateBackend) UnmarshalBinary(b []byte) error {
	var res ResourceTfStateBackend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
