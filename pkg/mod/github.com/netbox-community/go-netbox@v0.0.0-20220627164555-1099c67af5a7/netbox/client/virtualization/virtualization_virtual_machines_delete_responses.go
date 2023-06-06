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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationVirtualMachinesDeleteReader is a Reader for the VirtualizationVirtualMachinesDelete structure.
type VirtualizationVirtualMachinesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationVirtualMachinesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualMachinesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesDeleteNoContent creates a VirtualizationVirtualMachinesDeleteNoContent with default headers values
func NewVirtualizationVirtualMachinesDeleteNoContent() *VirtualizationVirtualMachinesDeleteNoContent {
	return &VirtualizationVirtualMachinesDeleteNoContent{}
}

/* VirtualizationVirtualMachinesDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationVirtualMachinesDeleteNoContent virtualization virtual machines delete no content
*/
type VirtualizationVirtualMachinesDeleteNoContent struct {
}

func (o *VirtualizationVirtualMachinesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesDeleteNoContent ", 204)
}

func (o *VirtualizationVirtualMachinesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationVirtualMachinesDeleteDefault creates a VirtualizationVirtualMachinesDeleteDefault with default headers values
func NewVirtualizationVirtualMachinesDeleteDefault(code int) *VirtualizationVirtualMachinesDeleteDefault {
	return &VirtualizationVirtualMachinesDeleteDefault{
		_statusCode: code,
	}
}

/* VirtualizationVirtualMachinesDeleteDefault describes a response with status code -1, with default header values.

VirtualizationVirtualMachinesDeleteDefault virtualization virtual machines delete default
*/
type VirtualizationVirtualMachinesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization virtual machines delete default response
func (o *VirtualizationVirtualMachinesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationVirtualMachinesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_delete default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationVirtualMachinesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}