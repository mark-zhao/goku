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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// TenancyContactsBulkPartialUpdateReader is a Reader for the TenancyContactsBulkPartialUpdate structure.
type TenancyContactsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactsBulkPartialUpdateOK creates a TenancyContactsBulkPartialUpdateOK with default headers values
func NewTenancyContactsBulkPartialUpdateOK() *TenancyContactsBulkPartialUpdateOK {
	return &TenancyContactsBulkPartialUpdateOK{}
}

/* TenancyContactsBulkPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactsBulkPartialUpdateOK tenancy contacts bulk partial update o k
*/
type TenancyContactsBulkPartialUpdateOK struct {
	Payload *models.Contact
}

func (o *TenancyContactsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contacts/][%d] tenancyContactsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyContactsBulkPartialUpdateOK) GetPayload() *models.Contact {
	return o.Payload
}

func (o *TenancyContactsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactsBulkPartialUpdateDefault creates a TenancyContactsBulkPartialUpdateDefault with default headers values
func NewTenancyContactsBulkPartialUpdateDefault(code int) *TenancyContactsBulkPartialUpdateDefault {
	return &TenancyContactsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* TenancyContactsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyContactsBulkPartialUpdateDefault tenancy contacts bulk partial update default
*/
type TenancyContactsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contacts bulk partial update default response
func (o *TenancyContactsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contacts/][%d] tenancy_contacts_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyContactsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
