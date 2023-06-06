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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// VirtualizationClustersListReader is a Reader for the VirtualizationClustersList structure.
type VirtualizationClustersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClustersListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClustersListOK creates a VirtualizationClustersListOK with default headers values
func NewVirtualizationClustersListOK() *VirtualizationClustersListOK {
	return &VirtualizationClustersListOK{}
}

/* VirtualizationClustersListOK describes a response with status code 200, with default header values.

VirtualizationClustersListOK virtualization clusters list o k
*/
type VirtualizationClustersListOK struct {
	Payload *VirtualizationClustersListOKBody
}

func (o *VirtualizationClustersListOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/clusters/][%d] virtualizationClustersListOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClustersListOK) GetPayload() *VirtualizationClustersListOKBody {
	return o.Payload
}

func (o *VirtualizationClustersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VirtualizationClustersListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersListDefault creates a VirtualizationClustersListDefault with default headers values
func NewVirtualizationClustersListDefault(code int) *VirtualizationClustersListDefault {
	return &VirtualizationClustersListDefault{
		_statusCode: code,
	}
}

/* VirtualizationClustersListDefault describes a response with status code -1, with default header values.

VirtualizationClustersListDefault virtualization clusters list default
*/
type VirtualizationClustersListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization clusters list default response
func (o *VirtualizationClustersListDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClustersListDefault) Error() string {
	return fmt.Sprintf("[GET /virtualization/clusters/][%d] virtualization_clusters_list default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationClustersListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VirtualizationClustersListOKBody virtualization clusters list o k body
swagger:model VirtualizationClustersListOKBody
*/
type VirtualizationClustersListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.Cluster `json:"results"`
}

// Validate validates this virtualization clusters list o k body
func (o *VirtualizationClustersListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VirtualizationClustersListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("virtualizationClustersListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *VirtualizationClustersListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("virtualizationClustersListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *VirtualizationClustersListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("virtualizationClustersListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *VirtualizationClustersListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("virtualizationClustersListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualizationClustersListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualizationClustersListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this virtualization clusters list o k body based on the context it is used
func (o *VirtualizationClustersListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VirtualizationClustersListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualizationClustersListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualizationClustersListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *VirtualizationClustersListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VirtualizationClustersListOKBody) UnmarshalBinary(b []byte) error {
	var res VirtualizationClustersListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
