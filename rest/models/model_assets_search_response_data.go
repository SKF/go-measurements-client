// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/SKF/go-utility/v2/uuid"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelAssetsSearchResponseData model assets search response data
//
// swagger:model model.AssetsSearchResponseData
type ModelAssetsSearchResponseData struct {

	// asset Id
	AssetID uuid.UUID `json:"assetId,omitempty"`

	// collected at
	// Format: date-time
	CollectedAt strfmt.DateTime `json:"collectedAt,omitempty"`

	// reason
	Reason *ModelAssetsSearchResponseDataReason `json:"reason,omitempty"`
}

// Validate validates this model assets search response data
func (m *ModelAssetsSearchResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelAssetsSearchResponseData) validateCollectedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CollectedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("collectedAt", "body", "date-time", m.CollectedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelAssetsSearchResponseData) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if m.Reason != nil {
		if err := m.Reason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelAssetsSearchResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelAssetsSearchResponseData) UnmarshalBinary(b []byte) error {
	var res ModelAssetsSearchResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}