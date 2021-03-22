// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UnauthorizedChartAPIError Unauthorized
//
// swagger:model UnauthorizedChartAPIError
type UnauthorizedChartAPIError struct {
}

// Validate validates this unauthorized chart API error
func (m *UnauthorizedChartAPIError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UnauthorizedChartAPIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnauthorizedChartAPIError) UnmarshalBinary(b []byte) error {
	var res UnauthorizedChartAPIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
