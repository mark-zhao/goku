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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// IpamPrefixesAvailablePrefixesListReader is a Reader for the IpamPrefixesAvailablePrefixesList structure.
type IpamPrefixesAvailablePrefixesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailablePrefixesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesAvailablePrefixesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesAvailablePrefixesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesAvailablePrefixesListOK creates a IpamPrefixesAvailablePrefixesListOK with default headers values
func NewIpamPrefixesAvailablePrefixesListOK() *IpamPrefixesAvailablePrefixesListOK {
	return &IpamPrefixesAvailablePrefixesListOK{}
}

/* IpamPrefixesAvailablePrefixesListOK describes a response with status code 200, with default header values.

IpamPrefixesAvailablePrefixesListOK ipam prefixes available prefixes list o k
*/
type IpamPrefixesAvailablePrefixesListOK struct {
	Payload []*models.AvailablePrefix
}

func (o *IpamPrefixesAvailablePrefixesListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-prefixes/][%d] ipamPrefixesAvailablePrefixesListOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesAvailablePrefixesListOK) GetPayload() []*models.AvailablePrefix {
	return o.Payload
}

func (o *IpamPrefixesAvailablePrefixesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesAvailablePrefixesListDefault creates a IpamPrefixesAvailablePrefixesListDefault with default headers values
func NewIpamPrefixesAvailablePrefixesListDefault(code int) *IpamPrefixesAvailablePrefixesListDefault {
	return &IpamPrefixesAvailablePrefixesListDefault{
		_statusCode: code,
	}
}

/* IpamPrefixesAvailablePrefixesListDefault describes a response with status code -1, with default header values.

IpamPrefixesAvailablePrefixesListDefault ipam prefixes available prefixes list default
*/
type IpamPrefixesAvailablePrefixesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam prefixes available prefixes list default response
func (o *IpamPrefixesAvailablePrefixesListDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesAvailablePrefixesListDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-prefixes/][%d] ipam_prefixes_available-prefixes_list default  %+v", o._statusCode, o.Payload)
}
func (o *IpamPrefixesAvailablePrefixesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesAvailablePrefixesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
