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
)

// NewDcimFrontPortsBulkDeleteParams creates a new DcimFrontPortsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortsBulkDeleteParams() *DcimFrontPortsBulkDeleteParams {
	return &DcimFrontPortsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsBulkDeleteParamsWithTimeout creates a new DcimFrontPortsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimFrontPortsBulkDeleteParams {
	return &DcimFrontPortsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortsBulkDeleteParamsWithContext creates a new DcimFrontPortsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimFrontPortsBulkDeleteParamsWithContext(ctx context.Context) *DcimFrontPortsBulkDeleteParams {
	return &DcimFrontPortsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimFrontPortsBulkDeleteParamsWithHTTPClient creates a new DcimFrontPortsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimFrontPortsBulkDeleteParams {
	return &DcimFrontPortsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim front ports bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front ports bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsBulkDeleteParams) WithDefaults() *DcimFrontPortsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front ports bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front ports bulk delete params
func (o *DcimFrontPortsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimFrontPortsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports bulk delete params
func (o *DcimFrontPortsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports bulk delete params
func (o *DcimFrontPortsBulkDeleteParams) WithContext(ctx context.Context) *DcimFrontPortsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports bulk delete params
func (o *DcimFrontPortsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports bulk delete params
func (o *DcimFrontPortsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimFrontPortsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports bulk delete params
func (o *DcimFrontPortsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
