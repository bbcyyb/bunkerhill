// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Paging paging
// swagger:model paging
type Paging struct {

	// next
	Next strfmt.URI `json:"next,omitempty"`

	// previous
	Previous strfmt.URI `json:"previous,omitempty"`
}

// Validate validates this paging
func (m *Paging) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrevious(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Paging) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("next", "body", "uri", m.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Paging) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(m.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("previous", "body", "uri", m.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Paging) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Paging) UnmarshalBinary(b []byte) error {
	var res Paging
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
