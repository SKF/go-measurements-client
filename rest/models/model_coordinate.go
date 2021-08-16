// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelCoordinate model coordinate
//
// swagger:model model.Coordinate
type ModelCoordinate struct {

	// x
	X float32 `json:"x,omitempty"`

	// y
	Y float32 `json:"y,omitempty"`
}

// Validate validates this model coordinate
func (m *ModelCoordinate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelCoordinate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelCoordinate) UnmarshalBinary(b []byte) error {
	var res ModelCoordinate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
