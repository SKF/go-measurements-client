// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/SKF/go-utility/v2/uuid"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelNodeData model node data
//
// swagger:model model.NodeData
type ModelNodeData struct {

	// asset Id
	AssetID uuid.UUID `json:"assetId,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// data point
	DataPoint *ModelDataPoint `json:"dataPoint,omitempty"`

	// frequency
	// Enum: [very-high high medium low very-low unknown]
	Frequency string `json:"frequency,omitempty"`

	// measurement Id
	MeasurementID uuid.UUID `json:"measurementId,omitempty"`

	// node Id
	NodeID uuid.UUID `json:"nodeId,omitempty"`

	// note
	Note string `json:"note,omitempty"`

	// question answers
	QuestionAnswers []string `json:"questionAnswers"`

	// rate of change
	RateOfChange float32 `json:"rateOfChange,omitempty"`

	// site Id
	SiteID uuid.UUID `json:"siteId,omitempty"`

	// spectrum
	Spectrum *ModelSpectrum `json:"spectrum,omitempty"`

	// Tags is a JSON object encoded as a string
	Tags string `json:"tags,omitempty"`

	// time series
	TimeSeries *ModelTimeSeries `json:"timeSeries,omitempty"`
}

// Validate validates this model node data
func (m *ModelNodeData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
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

func (m *ModelNodeData) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelNodeData) validateDataPoint(formats strfmt.Registry) error {

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

var modelNodeDataTypeFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["very-high","high","medium","low","very-low","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelNodeDataTypeFrequencyPropEnum = append(modelNodeDataTypeFrequencyPropEnum, v)
	}
}

const (

	// ModelNodeDataFrequencyVeryHigh captures enum value "very-high"
	ModelNodeDataFrequencyVeryHigh string = "very-high"

	// ModelNodeDataFrequencyHigh captures enum value "high"
	ModelNodeDataFrequencyHigh string = "high"

	// ModelNodeDataFrequencyMedium captures enum value "medium"
	ModelNodeDataFrequencyMedium string = "medium"

	// ModelNodeDataFrequencyLow captures enum value "low"
	ModelNodeDataFrequencyLow string = "low"

	// ModelNodeDataFrequencyVeryLow captures enum value "very-low"
	ModelNodeDataFrequencyVeryLow string = "very-low"

	// ModelNodeDataFrequencyUnknown captures enum value "unknown"
	ModelNodeDataFrequencyUnknown string = "unknown"
)

// prop value enum
func (m *ModelNodeData) validateFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modelNodeDataTypeFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModelNodeData) validateFrequency(formats strfmt.Registry) error {

	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	// value enum
	if err := m.validateFrequencyEnum("frequency", "body", m.Frequency); err != nil {
		return err
	}

	return nil
}

func (m *ModelNodeData) validateSpectrum(formats strfmt.Registry) error {

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

func (m *ModelNodeData) validateTimeSeries(formats strfmt.Registry) error {

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
func (m *ModelNodeData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelNodeData) UnmarshalBinary(b []byte) error {
	var res ModelNodeData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
