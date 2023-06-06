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

// DcimConsoleServerPortTemplatesUpdateReader is a Reader for the DcimConsoleServerPortTemplatesUpdate structure.
type DcimConsoleServerPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesUpdateOK creates a DcimConsoleServerPortTemplatesUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesUpdateOK() *DcimConsoleServerPortTemplatesUpdateOK {
	return &DcimConsoleServerPortTemplatesUpdateOK{}
}

/* DcimConsoleServerPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesUpdateOK dcim console server port templates update o k
*/
type DcimConsoleServerPortTemplatesUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesUpdateDefault creates a DcimConsoleServerPortTemplatesUpdateDefault with default headers values
func NewDcimConsoleServerPortTemplatesUpdateDefault(code int) *DcimConsoleServerPortTemplatesUpdateDefault {
	return &DcimConsoleServerPortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/* DcimConsoleServerPortTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortTemplatesUpdateDefault dcim console server port templates update default
*/
type DcimConsoleServerPortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates update default response
func (o *DcimConsoleServerPortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
