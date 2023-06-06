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

// DcimVirtualChassisPartialUpdateReader is a Reader for the DcimVirtualChassisPartialUpdate structure.
type DcimVirtualChassisPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualChassisPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisPartialUpdateOK creates a DcimVirtualChassisPartialUpdateOK with default headers values
func NewDcimVirtualChassisPartialUpdateOK() *DcimVirtualChassisPartialUpdateOK {
	return &DcimVirtualChassisPartialUpdateOK{}
}

/* DcimVirtualChassisPartialUpdateOK describes a response with status code 200, with default header values.

DcimVirtualChassisPartialUpdateOK dcim virtual chassis partial update o k
*/
type DcimVirtualChassisPartialUpdateOK struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimVirtualChassisPartialUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisPartialUpdateDefault creates a DcimVirtualChassisPartialUpdateDefault with default headers values
func NewDcimVirtualChassisPartialUpdateDefault(code int) *DcimVirtualChassisPartialUpdateDefault {
	return &DcimVirtualChassisPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimVirtualChassisPartialUpdateDefault describes a response with status code -1, with default header values.

DcimVirtualChassisPartialUpdateDefault dcim virtual chassis partial update default
*/
type DcimVirtualChassisPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim virtual chassis partial update default response
func (o *DcimVirtualChassisPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimVirtualChassisPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcim_virtual-chassis_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimVirtualChassisPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}