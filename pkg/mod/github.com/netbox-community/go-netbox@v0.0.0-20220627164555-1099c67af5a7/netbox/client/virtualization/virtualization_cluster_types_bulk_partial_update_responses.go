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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// VirtualizationClusterTypesBulkPartialUpdateReader is a Reader for the VirtualizationClusterTypesBulkPartialUpdate structure.
type VirtualizationClusterTypesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesBulkPartialUpdateOK creates a VirtualizationClusterTypesBulkPartialUpdateOK with default headers values
func NewVirtualizationClusterTypesBulkPartialUpdateOK() *VirtualizationClusterTypesBulkPartialUpdateOK {
	return &VirtualizationClusterTypesBulkPartialUpdateOK{}
}

/* VirtualizationClusterTypesBulkPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterTypesBulkPartialUpdateOK virtualization cluster types bulk partial update o k
*/
type VirtualizationClusterTypesBulkPartialUpdateOK struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-types/][%d] virtualizationClusterTypesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClusterTypesBulkPartialUpdateOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterTypesBulkPartialUpdateDefault creates a VirtualizationClusterTypesBulkPartialUpdateDefault with default headers values
func NewVirtualizationClusterTypesBulkPartialUpdateDefault(code int) *VirtualizationClusterTypesBulkPartialUpdateDefault {
	return &VirtualizationClusterTypesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* VirtualizationClusterTypesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClusterTypesBulkPartialUpdateDefault virtualization cluster types bulk partial update default
*/
type VirtualizationClusterTypesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster types bulk partial update default response
func (o *VirtualizationClusterTypesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterTypesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-types/][%d] virtualization_cluster-types_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationClusterTypesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
