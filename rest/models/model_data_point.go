// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelDataPoint model data point
//
// swagger:model model.DataPoint
type ModelDataPoint struct {

	// coordinate
	Coordinate *ModelCoordinate `json:"coordinate,omitempty"`

	// XUnit describes the X unit, such as a ms (millisecond)
	XUnit string `json:"xUnit,omitempty"`

	// YUnit describes the Y unit, such as acceleration enveloping peak-to-peak (gE PtP), mm/s (millimeters per second), in/s (inch per second), Celsius (°C) or Fahrenheit (°F)
	YUnit string `json:"yUnit,omitempty"`
}

// Validate validates this model data point
func (m *ModelDataPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoordinate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelDataPoint) validateCoordinate(formats strfmt.Registry) error {

	if swag.IsZero(m.Coordinate) { // not required
		return nil
	}

	if m.Coordinate != nil {
		if err := m.Coordinate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coordinate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelDataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelDataPoint) UnmarshalBinary(b []byte) error {
	var res ModelDataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
