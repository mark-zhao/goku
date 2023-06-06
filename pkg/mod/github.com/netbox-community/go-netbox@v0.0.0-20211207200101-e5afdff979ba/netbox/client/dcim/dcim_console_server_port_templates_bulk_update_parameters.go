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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewDcimConsoleServerPortTemplatesBulkUpdateParams creates a new DcimConsoleServerPortTemplatesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortTemplatesBulkUpdateParams() *DcimConsoleServerPortTemplatesBulkUpdateParams {
	return &DcimConsoleServerPortTemplatesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortTemplatesBulkUpdateParamsWithTimeout creates a new DcimConsoleServerPortTemplatesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortTemplatesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesBulkUpdateParams {
	return &DcimConsoleServerPortTemplatesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortTemplatesBulkUpdateParamsWithContext creates a new DcimConsoleServerPortTemplatesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortTemplatesBulkUpdateParamsWithContext(ctx context.Context) *DcimConsoleServerPortTemplatesBulkUpdateParams {
	return &DcimConsoleServerPortTemplatesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortTemplatesBulkUpdateParamsWithHTTPClient creates a new DcimConsoleServerPortTemplatesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortTemplatesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesBulkUpdateParams {
	return &DcimConsoleServerPortTemplatesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsoleServerPortTemplatesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console server port templates bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleServerPortTemplatesBulkUpdateParams struct {

	// Data.
	Data *models.WritableConsoleServerPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server port templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) WithDefaults() *DcimConsoleServerPortTemplatesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server port templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server port templates bulk update params
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server port templates bulk update params
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server port templates bulk update params
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) WithContext(ctx context.Context) *DcimConsoleServerPortTemplatesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server port templates bulk update params
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server port templates bulk update params
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server port templates bulk update params
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server port templates bulk update params
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) WithData(data *models.WritableConsoleServerPortTemplate) *DcimConsoleServerPortTemplatesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server port templates bulk update params
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) SetData(data *models.WritableConsoleServerPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortTemplatesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
