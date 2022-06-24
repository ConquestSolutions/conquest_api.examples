// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIStyle conquest api style
//
// swagger:model conquest_apiStyle
type ConquestAPIStyle struct {

	// font style
	FontStyle *ConquestAPIFontStyle `json:"FontStyle,omitempty"`

	// line style
	LineStyle *ConquestAPILineStyle `json:"LineStyle,omitempty"`

	// shape style
	ShapeStyle *ConquestAPIShapeStyle `json:"ShapeStyle,omitempty"`

	// style type
	StyleType string `json:"StyleType,omitempty"`
}

// Validate validates this conquest api style
func (m *ConquestAPIStyle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFontStyle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineStyle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShapeStyle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIStyle) validateFontStyle(formats strfmt.Registry) error {
	if swag.IsZero(m.FontStyle) { // not required
		return nil
	}

	if m.FontStyle != nil {
		if err := m.FontStyle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FontStyle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FontStyle")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIStyle) validateLineStyle(formats strfmt.Registry) error {
	if swag.IsZero(m.LineStyle) { // not required
		return nil
	}

	if m.LineStyle != nil {
		if err := m.LineStyle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LineStyle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LineStyle")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIStyle) validateShapeStyle(formats strfmt.Registry) error {
	if swag.IsZero(m.ShapeStyle) { // not required
		return nil
	}

	if m.ShapeStyle != nil {
		if err := m.ShapeStyle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShapeStyle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShapeStyle")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api style based on the context it is used
func (m *ConquestAPIStyle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFontStyle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineStyle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShapeStyle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIStyle) contextValidateFontStyle(ctx context.Context, formats strfmt.Registry) error {

	if m.FontStyle != nil {
		if err := m.FontStyle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FontStyle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FontStyle")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIStyle) contextValidateLineStyle(ctx context.Context, formats strfmt.Registry) error {

	if m.LineStyle != nil {
		if err := m.LineStyle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LineStyle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LineStyle")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIStyle) contextValidateShapeStyle(ctx context.Context, formats strfmt.Registry) error {

	if m.ShapeStyle != nil {
		if err := m.ShapeStyle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShapeStyle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShapeStyle")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIStyle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIStyle) UnmarshalBinary(b []byte) error {
	var res ConquestAPIStyle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
