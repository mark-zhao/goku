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

// NewDcimModulesBulkDeleteParams creates a new DcimModulesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModulesBulkDeleteParams() *DcimModulesBulkDeleteParams {
	return &DcimModulesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModulesBulkDeleteParamsWithTimeout creates a new DcimModulesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimModulesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimModulesBulkDeleteParams {
	return &DcimModulesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimModulesBulkDeleteParamsWithContext creates a new DcimModulesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimModulesBulkDeleteParamsWithContext(ctx context.Context) *DcimModulesBulkDeleteParams {
	return &DcimModulesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimModulesBulkDeleteParamsWithHTTPClient creates a new DcimModulesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModulesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimModulesBulkDeleteParams {
	return &DcimModulesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimModulesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim modules bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimModulesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim modules bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesBulkDeleteParams) WithDefaults() *DcimModulesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim modules bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim modules bulk delete params
func (o *DcimModulesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimModulesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim modules bulk delete params
func (o *DcimModulesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim modules bulk delete params
func (o *DcimModulesBulkDeleteParams) WithContext(ctx context.Context) *DcimModulesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim modules bulk delete params
func (o *DcimModulesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim modules bulk delete params
func (o *DcimModulesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimModulesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim modules bulk delete params
func (o *DcimModulesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModulesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
