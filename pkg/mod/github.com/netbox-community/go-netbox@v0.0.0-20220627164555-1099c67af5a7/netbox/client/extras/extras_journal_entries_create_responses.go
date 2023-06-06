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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// ExtrasJournalEntriesCreateReader is a Reader for the ExtrasJournalEntriesCreate structure.
type ExtrasJournalEntriesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasJournalEntriesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasJournalEntriesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasJournalEntriesCreateCreated creates a ExtrasJournalEntriesCreateCreated with default headers values
func NewExtrasJournalEntriesCreateCreated() *ExtrasJournalEntriesCreateCreated {
	return &ExtrasJournalEntriesCreateCreated{}
}

/* ExtrasJournalEntriesCreateCreated describes a response with status code 201, with default header values.

ExtrasJournalEntriesCreateCreated extras journal entries create created
*/
type ExtrasJournalEntriesCreateCreated struct {
	Payload *models.JournalEntry
}

func (o *ExtrasJournalEntriesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/journal-entries/][%d] extrasJournalEntriesCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasJournalEntriesCreateCreated) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasJournalEntriesCreateDefault creates a ExtrasJournalEntriesCreateDefault with default headers values
func NewExtrasJournalEntriesCreateDefault(code int) *ExtrasJournalEntriesCreateDefault {
	return &ExtrasJournalEntriesCreateDefault{
		_statusCode: code,
	}
}

/* ExtrasJournalEntriesCreateDefault describes a response with status code -1, with default header values.

ExtrasJournalEntriesCreateDefault extras journal entries create default
*/
type ExtrasJournalEntriesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras journal entries create default response
func (o *ExtrasJournalEntriesCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasJournalEntriesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/journal-entries/][%d] extras_journal-entries_create default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasJournalEntriesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJournalEntriesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
