// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelEvent model event
//
// swagger:model model.Event
type ModelEvent struct {

	// content type
	ContentType string `json:"contentType,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// node Id
	// Format: uuid
	NodeID strfmt.UUID `json:"nodeId,omitempty"`

	// seq no
	SeqNo string `json:"seqNo,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this model event
func (m *ModelEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelEvent) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelEvent) validateNodeID(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeID) { // not required
		return nil
	}

	if err := validate.FormatOf("nodeId", "body", "uuid", m.NodeID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelEvent) UnmarshalBinary(b []byte) error {
	var res ModelEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
