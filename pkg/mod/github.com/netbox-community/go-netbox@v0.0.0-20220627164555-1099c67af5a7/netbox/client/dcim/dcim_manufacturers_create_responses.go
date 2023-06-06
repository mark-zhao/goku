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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimManufacturersCreateReader is a Reader for the DcimManufacturersCreate structure.
type DcimManufacturersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimManufacturersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimManufacturersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersCreateCreated creates a DcimManufacturersCreateCreated with default headers values
func NewDcimManufacturersCreateCreated() *DcimManufacturersCreateCreated {
	return &DcimManufacturersCreateCreated{}
}

/* DcimManufacturersCreateCreated describes a response with status code 201, with default header values.

DcimManufacturersCreateCreated dcim manufacturers create created
*/
type DcimManufacturersCreateCreated struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/manufacturers/][%d] dcimManufacturersCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimManufacturersCreateCreated) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersCreateDefault creates a DcimManufacturersCreateDefault with default headers values
func NewDcimManufacturersCreateDefault(code int) *DcimManufacturersCreateDefault {
	return &DcimManufacturersCreateDefault{
		_statusCode: code,
	}
}

/* DcimManufacturersCreateDefault describes a response with status code -1, with default header values.

DcimManufacturersCreateDefault dcim manufacturers create default
*/
type DcimManufacturersCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim manufacturers create default response
func (o *DcimManufacturersCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimManufacturersCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/manufacturers/][%d] dcim_manufacturers_create default  %+v", o._statusCode, o.Payload)
}
func (o *DcimManufacturersCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
