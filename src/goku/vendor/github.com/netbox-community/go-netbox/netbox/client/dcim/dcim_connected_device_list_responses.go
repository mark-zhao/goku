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

// DcimConnectedDeviceListReader is a Reader for the DcimConnectedDeviceList structure.
type DcimConnectedDeviceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConnectedDeviceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConnectedDeviceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConnectedDeviceListOK creates a DcimConnectedDeviceListOK with default headers values
func NewDcimConnectedDeviceListOK() *DcimConnectedDeviceListOK {
	return &DcimConnectedDeviceListOK{}
}

/* DcimConnectedDeviceListOK describes a response with status code 200, with default header values.

DcimConnectedDeviceListOK dcim connected device list o k
*/
type DcimConnectedDeviceListOK struct {
	Payload *models.Device
}

func (o *DcimConnectedDeviceListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/connected-device/][%d] dcimConnectedDeviceListOK  %+v", 200, o.Payload)
}
func (o *DcimConnectedDeviceListOK) GetPayload() *models.Device {
	return o.Payload
}

func (o *DcimConnectedDeviceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Device)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}