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

// DcimInterfaceTemplatesDeleteReader is a Reader for the DcimInterfaceTemplatesDelete structure.
type DcimInterfaceTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInterfaceTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfaceTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfaceTemplatesDeleteNoContent creates a DcimInterfaceTemplatesDeleteNoContent with default headers values
func NewDcimInterfaceTemplatesDeleteNoContent() *DcimInterfaceTemplatesDeleteNoContent {
	return &DcimInterfaceTemplatesDeleteNoContent{}
}

/* DcimInterfaceTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimInterfaceTemplatesDeleteNoContent dcim interface templates delete no content
*/
type DcimInterfaceTemplatesDeleteNoContent struct {
}

func (o *DcimInterfaceTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesDeleteNoContent ", 204)
}

func (o *DcimInterfaceTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimInterfaceTemplatesDeleteDefault creates a DcimInterfaceTemplatesDeleteDefault with default headers values
func NewDcimInterfaceTemplatesDeleteDefault(code int) *DcimInterfaceTemplatesDeleteDefault {
	return &DcimInterfaceTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/* DcimInterfaceTemplatesDeleteDefault describes a response with status code -1, with default header values.

DcimInterfaceTemplatesDeleteDefault dcim interface templates delete default
*/
type DcimInterfaceTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interface templates delete default response
func (o *DcimInterfaceTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimInterfaceTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interface-templates/{id}/][%d] dcim_interface-templates_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInterfaceTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfaceTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}