// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelAssetsSearchRequest model assets search request
//
// swagger:model model.AssetsSearchRequest
type ModelAssetsSearchRequest struct {

	// The list of asset IDs with a max per request of 1000
	AssetIDs []string `json:"assetIDs"`

	// tags
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// Validate validates this model assets search request
func (m *ModelAssetsSearchRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelAssetsSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelAssetsSearchRequest) UnmarshalBinary(b []byte) error {
	var res ModelAssetsSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
