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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExtrasCustomFieldsReadParams creates a new ExtrasCustomFieldsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldsReadParams() *ExtrasCustomFieldsReadParams {
	return &ExtrasCustomFieldsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldsReadParamsWithTimeout creates a new ExtrasCustomFieldsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldsReadParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldsReadParams {
	return &ExtrasCustomFieldsReadParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldsReadParamsWithContext creates a new ExtrasCustomFieldsReadParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldsReadParamsWithContext(ctx context.Context) *ExtrasCustomFieldsReadParams {
	return &ExtrasCustomFieldsReadParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldsReadParamsWithHTTPClient creates a new ExtrasCustomFieldsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldsReadParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldsReadParams {
	return &ExtrasCustomFieldsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldsReadParams contains all the parameters to send to the API endpoint
   for the extras custom fields read operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldsReadParams struct {

	/* ID.

	   A unique integer value identifying this custom field.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom fields read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsReadParams) WithDefaults() *ExtrasCustomFieldsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom fields read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) WithContext(ctx context.Context) *ExtrasCustomFieldsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) WithID(id int64) *ExtrasCustomFieldsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
