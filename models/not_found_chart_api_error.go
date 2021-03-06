// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotFoundChartAPIError Not found
//
// swagger:model NotFoundChartAPIError
type NotFoundChartAPIError struct {
}

// Validate validates this not found chart API error
func (m *NotFoundChartAPIError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotFoundChartAPIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotFoundChartAPIError) UnmarshalBinary(b []byte) error {
	var res NotFoundChartAPIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
