// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RobotAccount The object of robot account
//
// swagger:model RobotAccount
type RobotAccount struct {

	// The creation time of the robot account
	CreationTime string `json:"creation_time,omitempty"`

	// The description of robot account
	Description string `json:"description,omitempty"`

	// The robot account is disable or enable
	Disabled bool `json:"disabled,omitempty"`

	// The expiration of robot account (in seconds)
	ExpiresAt int64 `json:"expires_at,omitempty"`

	// The id of robot account
	ID int64 `json:"id,omitempty"`

	// The name of robot account
	Name string `json:"name,omitempty"`

	// The project id of robot account
	ProjectID int64 `json:"project_id,omitempty"`

	// The update time of the robot account
	UpdateTime string `json:"update_time,omitempty"`
}

// Validate validates this robot account
func (m *RobotAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RobotAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RobotAccount) UnmarshalBinary(b []byte) error {
	var res RobotAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
