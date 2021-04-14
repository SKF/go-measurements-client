// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelExportMetadata model export metadata
//
// swagger:model model.ExportMetadata
type ModelExportMetadata struct {

	// from
	From string `json:"from,omitempty"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`

	// to
	To string `json:"to,omitempty"`
}

// Validate validates this model export metadata
func (m *ModelExportMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelExportMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelExportMetadata) UnmarshalBinary(b []byte) error {
	var res ModelExportMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
