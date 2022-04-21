// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelNodeBandOverallResponse model node band overall response
//
// swagger:model model.NodeBandOverallResponse
type ModelNodeBandOverallResponse struct {

	// trends
	Trends []*ModelMeasurementBandOverall `json:"trends"`
}

// Validate validates this model node band overall response
func (m *ModelNodeBandOverallResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrends(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelNodeBandOverallResponse) validateTrends(formats strfmt.Registry) error {

	if swag.IsZero(m.Trends) { // not required
		return nil
	}

	for i := 0; i < len(m.Trends); i++ {
		if swag.IsZero(m.Trends[i]) { // not required
			continue
		}

		if m.Trends[i] != nil {
			if err := m.Trends[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelNodeBandOverallResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelNodeBandOverallResponse) UnmarshalBinary(b []byte) error {
	var res ModelNodeBandOverallResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}