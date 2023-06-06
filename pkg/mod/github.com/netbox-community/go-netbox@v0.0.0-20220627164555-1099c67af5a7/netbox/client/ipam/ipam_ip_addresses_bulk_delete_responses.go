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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamIPAddressesBulkDeleteReader is a Reader for the IpamIPAddressesBulkDelete structure.
type IpamIPAddressesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamIPAddressesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPAddressesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPAddressesBulkDeleteNoContent creates a IpamIPAddressesBulkDeleteNoContent with default headers values
func NewIpamIPAddressesBulkDeleteNoContent() *IpamIPAddressesBulkDeleteNoContent {
	return &IpamIPAddressesBulkDeleteNoContent{}
}

/* IpamIPAddressesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamIPAddressesBulkDeleteNoContent ipam Ip addresses bulk delete no content
*/
type IpamIPAddressesBulkDeleteNoContent struct {
}

func (o *IpamIPAddressesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-addresses/][%d] ipamIpAddressesBulkDeleteNoContent ", 204)
}

func (o *IpamIPAddressesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamIPAddressesBulkDeleteDefault creates a IpamIPAddressesBulkDeleteDefault with default headers values
func NewIpamIPAddressesBulkDeleteDefault(code int) *IpamIPAddressesBulkDeleteDefault {
	return &IpamIPAddressesBulkDeleteDefault{
		_statusCode: code,
	}
}

/* IpamIPAddressesBulkDeleteDefault describes a response with status code -1, with default header values.

IpamIPAddressesBulkDeleteDefault ipam ip addresses bulk delete default
*/
type IpamIPAddressesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip addresses bulk delete default response
func (o *IpamIPAddressesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPAddressesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-addresses/][%d] ipam_ip-addresses_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamIPAddressesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPAddressesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}