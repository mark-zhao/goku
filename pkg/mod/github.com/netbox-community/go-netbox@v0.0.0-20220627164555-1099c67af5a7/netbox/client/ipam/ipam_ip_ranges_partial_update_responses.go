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

// IpamIPRangesPartialUpdateReader is a Reader for the IpamIPRangesPartialUpdate structure.
type IpamIPRangesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPRangesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPRangesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesPartialUpdateOK creates a IpamIPRangesPartialUpdateOK with default headers values
func NewIpamIPRangesPartialUpdateOK() *IpamIPRangesPartialUpdateOK {
	return &IpamIPRangesPartialUpdateOK{}
}

/* IpamIPRangesPartialUpdateOK describes a response with status code 200, with default header values.

IpamIPRangesPartialUpdateOK ipam Ip ranges partial update o k
*/
type IpamIPRangesPartialUpdateOK struct {
	Payload *models.IPRange
}

func (o *IpamIPRangesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-ranges/{id}/][%d] ipamIpRangesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamIPRangesPartialUpdateOK) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesPartialUpdateDefault creates a IpamIPRangesPartialUpdateDefault with default headers values
func NewIpamIPRangesPartialUpdateDefault(code int) *IpamIPRangesPartialUpdateDefault {
	return &IpamIPRangesPartialUpdateDefault{
		_statusCode: code,
	}
}

/* IpamIPRangesPartialUpdateDefault describes a response with status code -1, with default header values.

IpamIPRangesPartialUpdateDefault ipam ip ranges partial update default
*/
type IpamIPRangesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip ranges partial update default response
func (o *IpamIPRangesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPRangesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-ranges/{id}/][%d] ipam_ip-ranges_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamIPRangesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
