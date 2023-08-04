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

// ModelNodeDataRequest model node data request
//
// swagger:model model.NodeDataRequest
type ModelNodeDataRequest struct {

	// content type
	ContentType string `json:"contentType,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// data point
	DataPoint *ModelDataPoint `json:"dataPoint,omitempty"`

	// media
	Media *ModelMedia `json:"media,omitempty"`

	// node Id
	NodeID uuid.UUID `json:"nodeId,omitempty"`

	// note
	Note string `json:"note,omitempty"`

	// origin measurement Id
	OriginMeasurementID string `json:"originMeasurementId,omitempty"`

	// question answers
	QuestionAnswers []string `json:"questionAnswers"`

	// spectrum
	Spectrum *ModelSpectrum `json:"spectrum,omitempty"`

	// A `reason` tag is required for measurements with content type `MISSING_VALUE`. Please refer to the
	// [service documentation](/docs/service/node-data#missing-value) for examples and recommended values for this tag.
	Tags map[string]interface{} `json:"tags,omitempty"`

	// time series
	TimeSeries *ModelTimeSeries `json:"timeSeries,omitempty"`

	// user Id
	UserID uuid.UUID `json:"userId,omitempty"`
}

// Validate validates this model node data request
func (m *ModelNodeDataRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpectrum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeSeries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelNodeDataRequest) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelNodeDataRequest) validateDataPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.DataPoint) { // not required
		return nil
	}

	if m.DataPoint != nil {
		if err := m.DataPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataPoint")
			}
			return err
		}
	}

	return nil
}

func (m *ModelNodeDataRequest) validateMedia(formats strfmt.Registry) error {

	if swag.IsZero(m.Media) { // not required
		return nil
	}

	if m.Media != nil {
		if err := m.Media.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("media")
			}
			return err
		}
	}

	return nil
}

func (m *ModelNodeDataRequest) validateSpectrum(formats strfmt.Registry) error {

	if swag.IsZero(m.Spectrum) { // not required
		return nil
	}

	if m.Spectrum != nil {
		if err := m.Spectrum.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spectrum")
			}
			return err
		}
	}

	return nil
}

func (m *ModelNodeDataRequest) validateTimeSeries(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeSeries) { // not required
		return nil
	}

	if m.TimeSeries != nil {
		if err := m.TimeSeries.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeSeries")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelNodeDataRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelNodeDataRequest) UnmarshalBinary(b []byte) error {
	var res ModelNodeDataRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
