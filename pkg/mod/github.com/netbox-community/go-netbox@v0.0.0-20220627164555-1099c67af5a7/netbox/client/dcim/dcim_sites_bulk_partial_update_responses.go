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

// DcimSitesBulkPartialUpdateReader is a Reader for the DcimSitesBulkPartialUpdate structure.
type DcimSitesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesBulkPartialUpdateOK creates a DcimSitesBulkPartialUpdateOK with default headers values
func NewDcimSitesBulkPartialUpdateOK() *DcimSitesBulkPartialUpdateOK {
	return &DcimSitesBulkPartialUpdateOK{}
}

/* DcimSitesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimSitesBulkPartialUpdateOK dcim sites bulk partial update o k
*/
type DcimSitesBulkPartialUpdateOK struct {
	Payload *models.Site
}

func (o *DcimSitesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcimSitesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSitesBulkPartialUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesBulkPartialUpdateDefault creates a DcimSitesBulkPartialUpdateDefault with default headers values
func NewDcimSitesBulkPartialUpdateDefault(code int) *DcimSitesBulkPartialUpdateDefault {
	return &DcimSitesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimSitesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimSitesBulkPartialUpdateDefault dcim sites bulk partial update default
*/
type DcimSitesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites bulk partial update default response
func (o *DcimSitesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSitesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcim_sites_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimSitesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
