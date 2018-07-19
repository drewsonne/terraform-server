// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceListTfDeployment resource list tf deployment
// swagger:model resource-list-tf-deployment
type ResourceListTfDeployment struct {

	// deployments
	// Required: true
	Deployments []*ResourceTfDeployment `json:"deployments"`

	// stack
	// Required: true
	Stack *ResourceTfStack `json:"stack"`
}

// Validate validates this resource list tf deployment
func (m *ResourceListTfDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceListTfDeployment) validateDeployments(formats strfmt.Registry) error {

	if err := validate.Required("deployments", "body", m.Deployments); err != nil {
		return err
	}

	for i := 0; i < len(m.Deployments); i++ {
		if swag.IsZero(m.Deployments[i]) { // not required
			continue
		}

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourceListTfDeployment) validateStack(formats strfmt.Registry) error {

	if err := validate.Required("stack", "body", m.Stack); err != nil {
		return err
	}

	if m.Stack != nil {
		if err := m.Stack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceListTfDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceListTfDeployment) UnmarshalBinary(b []byte) error {
	var res ResourceListTfDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}