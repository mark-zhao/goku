// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomField custom field
//
// swagger:model CustomField
type CustomField struct {

	// Comma-separated list of available choices (for selection fields)
	Choices []string `json:"choices"`

	// content types
	// Required: true
	// Unique: true
	ContentTypes []string `json:"content_types"`

	// Default
	//
	// Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. "Foo").
	Default *string `json:"default,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// filter logic
	FilterLogic *CustomFieldFilterLogic `json:"filter_logic,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Name of the field as displayed to users (if not provided, the field's name will be used)
	// Max Length: 50
	Label string `json:"label,omitempty"`

	// Name
	//
	// Internal field name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Required
	//
	// If true, this field is required when creating new objects or editing an existing object.
	Required bool `json:"required,omitempty"`

	// type
	// Required: true
	Type *CustomFieldType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Maximum value
	//
	// Maximum allowed value (for numeric fields)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	ValidationMaximum *int64 `json:"validation_maximum,omitempty"`

	// Minimum value
	//
	// Minimum allowed value (for numeric fields)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	ValidationMinimum *int64 `json:"validation_minimum,omitempty"`

	// Validation regex
	//
	// Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, <code>^[A-Z]{3}$</code> will limit values to exactly three uppercase letters.
	// Max Length: 500
	ValidationRegex string `json:"validation_regex,omitempty"`

	// Weight
	//
	// Fields with higher weights appear lower in a form.
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this custom field
func (m *CustomField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChoices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterLogic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationMaximum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationMinimum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomField) validateChoices(formats strfmt.Registry) error {
	if swag.IsZero(m.Choices) { // not required
		return nil
	}

	for i := 0; i < len(m.Choices); i++ {

		if err := validate.MinLength("choices"+"."+strconv.Itoa(i), "body", m.Choices[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength("choices"+"."+strconv.Itoa(i), "body", m.Choices[i], 100); err != nil {
			return err
		}

	}

	return nil
}

func (m *CustomField) validateContentTypes(formats strfmt.Registry) error {

	if err := validate.Required("content_types", "body", m.ContentTypes); err != nil {
		return err
	}

	if err := validate.UniqueItems("content_types", "body", m.ContentTypes); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateFilterLogic(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterLogic) { // not required
		return nil
	}

	if m.FilterLogic != nil {
		if err := m.FilterLogic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_logic")
			}
			return err
		}
	}

	return nil
}

func (m *CustomField) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 50); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 50); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *CustomField) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateValidationMaximum(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationMaximum) { // not required
		return nil
	}

	if err := validate.MinimumInt("validation_maximum", "body", *m.ValidationMaximum, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("validation_maximum", "body", *m.ValidationMaximum, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateValidationMinimum(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationMinimum) { // not required
		return nil
	}

	if err := validate.MinimumInt("validation_minimum", "body", *m.ValidationMinimum, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("validation_minimum", "body", *m.ValidationMinimum, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateValidationRegex(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationRegex) { // not required
		return nil
	}

	if err := validate.MaxLength("validation_regex", "body", m.ValidationRegex, 500); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", *m.Weight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", *m.Weight, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this custom field based on the context it is used
func (m *CustomField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilterLogic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomField) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) contextValidateFilterLogic(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterLogic != nil {
		if err := m.FilterLogic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_logic")
			}
			return err
		}
	}

	return nil
}

func (m *CustomField) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *CustomField) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomField) UnmarshalBinary(b []byte) error {
	var res CustomField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CustomFieldFilterLogic Filter logic
//
// swagger:model CustomFieldFilterLogic
type CustomFieldFilterLogic struct {

	// label
	// Required: true
	// Enum: [Disabled Loose Exact]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [disabled loose exact]
	Value *string `json:"value"`
}

// Validate validates this custom field filter logic
func (m *CustomFieldFilterLogic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customFieldFilterLogicTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Disabled","Loose","Exact"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customFieldFilterLogicTypeLabelPropEnum = append(customFieldFilterLogicTypeLabelPropEnum, v)
	}
}

const (

	// CustomFieldFilterLogicLabelDisabled captures enum value "Disabled"
	CustomFieldFilterLogicLabelDisabled string = "Disabled"

	// CustomFieldFilterLogicLabelLoose captures enum value "Loose"
	CustomFieldFilterLogicLabelLoose string = "Loose"

	// CustomFieldFilterLogicLabelExact captures enum value "Exact"
	CustomFieldFilterLogicLabelExact string = "Exact"
)

// prop value enum
func (m *CustomFieldFilterLogic) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customFieldFilterLogicTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomFieldFilterLogic) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("filter_logic"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("filter_logic"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var customFieldFilterLogicTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disabled","loose","exact"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customFieldFilterLogicTypeValuePropEnum = append(customFieldFilterLogicTypeValuePropEnum, v)
	}
}

const (

	// CustomFieldFilterLogicValueDisabled captures enum value "disabled"
	CustomFieldFilterLogicValueDisabled string = "disabled"

	// CustomFieldFilterLogicValueLoose captures enum value "loose"
	CustomFieldFilterLogicValueLoose string = "loose"

	// CustomFieldFilterLogicValueExact captures enum value "exact"
	CustomFieldFilterLogicValueExact string = "exact"
)

// prop value enum
func (m *CustomFieldFilterLogic) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customFieldFilterLogicTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomFieldFilterLogic) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("filter_logic"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("filter_logic"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom field filter logic based on context it is used
func (m *CustomFieldFilterLogic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomFieldFilterLogic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomFieldFilterLogic) UnmarshalBinary(b []byte) error {
	var res CustomFieldFilterLogic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CustomFieldType Type
//
// swagger:model CustomFieldType
type CustomFieldType struct {

	// label
	// Required: true
	// Enum: [Text Integer Boolean (true/false) Date URL Selection Multiple selection]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [text integer boolean date url select multiselect]
	Value *string `json:"value"`
}

// Validate validates this custom field type
func (m *CustomFieldType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customFieldTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Text","Integer","Boolean (true/false)","Date","URL","Selection","Multiple selection"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customFieldTypeTypeLabelPropEnum = append(customFieldTypeTypeLabelPropEnum, v)
	}
}

const (

	// CustomFieldTypeLabelText captures enum value "Text"
	CustomFieldTypeLabelText string = "Text"

	// CustomFieldTypeLabelInteger captures enum value "Integer"
	CustomFieldTypeLabelInteger string = "Integer"

	// CustomFieldTypeLabelBooleanTrueFalse captures enum value "Boolean (true/false)"
	CustomFieldTypeLabelBooleanTrueFalse string = "Boolean (true/false)"

	// CustomFieldTypeLabelDate captures enum value "Date"
	CustomFieldTypeLabelDate string = "Date"

	// CustomFieldTypeLabelURL captures enum value "URL"
	CustomFieldTypeLabelURL string = "URL"

	// CustomFieldTypeLabelSelection captures enum value "Selection"
	CustomFieldTypeLabelSelection string = "Selection"

	// CustomFieldTypeLabelMultipleSelection captures enum value "Multiple selection"
	CustomFieldTypeLabelMultipleSelection string = "Multiple selection"
)

// prop value enum
func (m *CustomFieldType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customFieldTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomFieldType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var customFieldTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text","integer","boolean","date","url","select","multiselect"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customFieldTypeTypeValuePropEnum = append(customFieldTypeTypeValuePropEnum, v)
	}
}

const (

	// CustomFieldTypeValueText captures enum value "text"
	CustomFieldTypeValueText string = "text"

	// CustomFieldTypeValueInteger captures enum value "integer"
	CustomFieldTypeValueInteger string = "integer"

	// CustomFieldTypeValueBoolean captures enum value "boolean"
	CustomFieldTypeValueBoolean string = "boolean"

	// CustomFieldTypeValueDate captures enum value "date"
	CustomFieldTypeValueDate string = "date"

	// CustomFieldTypeValueURL captures enum value "url"
	CustomFieldTypeValueURL string = "url"

	// CustomFieldTypeValueSelect captures enum value "select"
	CustomFieldTypeValueSelect string = "select"

	// CustomFieldTypeValueMultiselect captures enum value "multiselect"
	CustomFieldTypeValueMultiselect string = "multiselect"
)

// prop value enum
func (m *CustomFieldType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customFieldTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomFieldType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom field type based on context it is used
func (m *CustomFieldType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomFieldType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomFieldType) UnmarshalBinary(b []byte) error {
	var res CustomFieldType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
