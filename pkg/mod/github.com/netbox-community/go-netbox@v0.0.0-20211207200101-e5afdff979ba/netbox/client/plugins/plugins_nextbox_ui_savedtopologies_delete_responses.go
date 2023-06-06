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

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsNextboxUISavedtopologiesDeleteReader is a Reader for the PluginsNextboxUISavedtopologiesDelete structure.
type PluginsNextboxUISavedtopologiesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNextboxUISavedtopologiesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsNextboxUISavedtopologiesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsNextboxUISavedtopologiesDeleteNoContent creates a PluginsNextboxUISavedtopologiesDeleteNoContent with default headers values
func NewPluginsNextboxUISavedtopologiesDeleteNoContent() *PluginsNextboxUISavedtopologiesDeleteNoContent {
	return &PluginsNextboxUISavedtopologiesDeleteNoContent{}
}

/* PluginsNextboxUISavedtopologiesDeleteNoContent describes a response with status code 204, with default header values.

PluginsNextboxUISavedtopologiesDeleteNoContent plugins nextbox Ui savedtopologies delete no content
*/
type PluginsNextboxUISavedtopologiesDeleteNoContent struct {
}

func (o *PluginsNextboxUISavedtopologiesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/nextbox-ui/savedtopologies/{id}/][%d] pluginsNextboxUiSavedtopologiesDeleteNoContent ", 204)
}

func (o *PluginsNextboxUISavedtopologiesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
