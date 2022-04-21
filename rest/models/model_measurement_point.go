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

// ModelMeasurementPoint model measurement point
//
// swagger:model model.MeasurementPoint
type ModelMeasurementPoint struct {

	// all parents
	AllParents string `json:"allParents,omitempty"`

	// asset Id
	AssetID uuid.UUID `json:"assetId,omitempty"`

	// collected at
	// Format: date-time
	CollectedAt strfmt.DateTime `json:"collectedAt,omitempty"`

	// company Id
	CompanyID uuid.UUID `json:"companyId,omitempty"`

	// frequency
	// Enum: [very-high high medium low very-low unknown]
	Frequency string `json:"frequency,omitempty"`

	// id
	ID float64 `json:"id,omitempty"`

	// last data point
	LastDataPoint float32 `json:"lastDataPoint,omitempty"`

	// measurement count
	MeasurementCount float64 `json:"measurementCount,omitempty"`

	// node Id
	NodeID uuid.UUID `json:"nodeId,omitempty"`

	// origin Id
	OriginID string `json:"originId,omitempty"`

	// plant Id
	PlantID uuid.UUID `json:"plantId,omitempty"`

	// site Id
	SiteID uuid.UUID `json:"siteId,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this model measurement point
func (m *ModelMeasurementPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelMeasurementPoint) validateCollectedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CollectedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("collectedAt", "body", "date-time", m.CollectedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var modelMeasurementPointTypeFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["very-high","high","medium","low","very-low","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelMeasurementPointTypeFrequencyPropEnum = append(modelMeasurementPointTypeFrequencyPropEnum, v)
	}
}

const (

	// ModelMeasurementPointFrequencyVeryHigh captures enum value "very-high"
	ModelMeasurementPointFrequencyVeryHigh string = "very-high"

	// ModelMeasurementPointFrequencyHigh captures enum value "high"
	ModelMeasurementPointFrequencyHigh string = "high"

	// ModelMeasurementPointFrequencyMedium captures enum value "medium"
	ModelMeasurementPointFrequencyMedium string = "medium"

	// ModelMeasurementPointFrequencyLow captures enum value "low"
	ModelMeasurementPointFrequencyLow string = "low"

	// ModelMeasurementPointFrequencyVeryLow captures enum value "very-low"
	ModelMeasurementPointFrequencyVeryLow string = "very-low"

	// ModelMeasurementPointFrequencyUnknown captures enum value "unknown"
	ModelMeasurementPointFrequencyUnknown string = "unknown"
)

// prop value enum
func (m *ModelMeasurementPoint) validateFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modelMeasurementPointTypeFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModelMeasurementPoint) validateFrequency(formats strfmt.Registry) error {

	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	// value enum
	if err := m.validateFrequencyEnum("frequency", "body", m.Frequency); err != nil {
		return err
	}

	return nil
}

func (m *ModelMeasurementPoint) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelMeasurementPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelMeasurementPoint) UnmarshalBinary(b []byte) error {
	var res ModelMeasurementPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}