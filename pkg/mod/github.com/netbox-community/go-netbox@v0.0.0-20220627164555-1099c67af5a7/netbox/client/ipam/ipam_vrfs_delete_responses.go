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

// IpamVrfsDeleteReader is a Reader for the IpamVrfsDelete structure.
type IpamVrfsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVrfsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsDeleteNoContent creates a IpamVrfsDeleteNoContent with default headers values
func NewIpamVrfsDeleteNoContent() *IpamVrfsDeleteNoContent {
	return &IpamVrfsDeleteNoContent{}
}

/* IpamVrfsDeleteNoContent describes a response with status code 204, with default header values.

IpamVrfsDeleteNoContent ipam vrfs delete no content
*/
type IpamVrfsDeleteNoContent struct {
}

func (o *IpamVrfsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/{id}/][%d] ipamVrfsDeleteNoContent ", 204)
}

func (o *IpamVrfsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamVrfsDeleteDefault creates a IpamVrfsDeleteDefault with default headers values
func NewIpamVrfsDeleteDefault(code int) *IpamVrfsDeleteDefault {
	return &IpamVrfsDeleteDefault{
		_statusCode: code,
	}
}

/* IpamVrfsDeleteDefault describes a response with status code -1, with default header values.

IpamVrfsDeleteDefault ipam vrfs delete default
*/
type IpamVrfsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vrfs delete default response
func (o *IpamVrfsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamVrfsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/{id}/][%d] ipam_vrfs_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVrfsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
