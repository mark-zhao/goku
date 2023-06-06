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
	"github.com/go-openapi/swag"
)

// NewDcimInventoryItemTemplatesReadParams creates a new DcimInventoryItemTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemTemplatesReadParams() *DcimInventoryItemTemplatesReadParams {
	return &DcimInventoryItemTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemTemplatesReadParamsWithTimeout creates a new DcimInventoryItemTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimInventoryItemTemplatesReadParams {
	return &DcimInventoryItemTemplatesReadParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemTemplatesReadParamsWithContext creates a new DcimInventoryItemTemplatesReadParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemTemplatesReadParamsWithContext(ctx context.Context) *DcimInventoryItemTemplatesReadParams {
	return &DcimInventoryItemTemplatesReadParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemTemplatesReadParamsWithHTTPClient creates a new DcimInventoryItemTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimInventoryItemTemplatesReadParams {
	return &DcimInventoryItemTemplatesReadParams{
		HTTPClient: client,
	}
}

/* DcimInventoryItemTemplatesReadParams contains all the parameters to send to the API endpoint
   for the dcim inventory item templates read operation.

   Typically these are written to a http.Request.
*/
type DcimInventoryItemTemplatesReadParams struct {

	/* ID.

	   A unique integer value identifying this inventory item template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory item templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemTemplatesReadParams) WithDefaults() *DcimInventoryItemTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory item templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory item templates read params
func (o *DcimInventoryItemTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimInventoryItemTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory item templates read params
func (o *DcimInventoryItemTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory item templates read params
func (o *DcimInventoryItemTemplatesReadParams) WithContext(ctx context.Context) *DcimInventoryItemTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory item templates read params
func (o *DcimInventoryItemTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory item templates read params
func (o *DcimInventoryItemTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimInventoryItemTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory item templates read params
func (o *DcimInventoryItemTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim inventory item templates read params
func (o *DcimInventoryItemTemplatesReadParams) WithID(id int64) *DcimInventoryItemTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory item templates read params
func (o *DcimInventoryItemTemplatesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
