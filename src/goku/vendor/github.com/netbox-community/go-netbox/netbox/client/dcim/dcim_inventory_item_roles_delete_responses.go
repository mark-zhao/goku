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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInventoryItemRolesDeleteReader is a Reader for the DcimInventoryItemRolesDelete structure.
type DcimInventoryItemRolesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInventoryItemRolesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemRolesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemRolesDeleteNoContent creates a DcimInventoryItemRolesDeleteNoContent with default headers values
func NewDcimInventoryItemRolesDeleteNoContent() *DcimInventoryItemRolesDeleteNoContent {
	return &DcimInventoryItemRolesDeleteNoContent{}
}

/* DcimInventoryItemRolesDeleteNoContent describes a response with status code 204, with default header values.

DcimInventoryItemRolesDeleteNoContent dcim inventory item roles delete no content
*/
type DcimInventoryItemRolesDeleteNoContent struct {
}

func (o *DcimInventoryItemRolesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-item-roles/{id}/][%d] dcimInventoryItemRolesDeleteNoContent ", 204)
}

func (o *DcimInventoryItemRolesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimInventoryItemRolesDeleteDefault creates a DcimInventoryItemRolesDeleteDefault with default headers values
func NewDcimInventoryItemRolesDeleteDefault(code int) *DcimInventoryItemRolesDeleteDefault {
	return &DcimInventoryItemRolesDeleteDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemRolesDeleteDefault describes a response with status code -1, with default header values.

DcimInventoryItemRolesDeleteDefault dcim inventory item roles delete default
*/
type DcimInventoryItemRolesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item roles delete default response
func (o *DcimInventoryItemRolesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemRolesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-item-roles/{id}/][%d] dcim_inventory-item-roles_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemRolesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
