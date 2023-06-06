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

package users

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

// NewUsersConfigListParams creates a new UsersConfigListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersConfigListParams() *UsersConfigListParams {
	return &UsersConfigListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersConfigListParamsWithTimeout creates a new UsersConfigListParams object
// with the ability to set a timeout on a request.
func NewUsersConfigListParamsWithTimeout(timeout time.Duration) *UsersConfigListParams {
	return &UsersConfigListParams{
		timeout: timeout,
	}
}

// NewUsersConfigListParamsWithContext creates a new UsersConfigListParams object
// with the ability to set a context for a request.
func NewUsersConfigListParamsWithContext(ctx context.Context) *UsersConfigListParams {
	return &UsersConfigListParams{
		Context: ctx,
	}
}

// NewUsersConfigListParamsWithHTTPClient creates a new UsersConfigListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersConfigListParamsWithHTTPClient(client *http.Client) *UsersConfigListParams {
	return &UsersConfigListParams{
		HTTPClient: client,
	}
}

/* UsersConfigListParams contains all the parameters to send to the API endpoint
   for the users config list operation.

   Typically these are written to a http.Request.
*/
type UsersConfigListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users config list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersConfigListParams) WithDefaults() *UsersConfigListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users config list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersConfigListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users config list params
func (o *UsersConfigListParams) WithTimeout(timeout time.Duration) *UsersConfigListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users config list params
func (o *UsersConfigListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users config list params
func (o *UsersConfigListParams) WithContext(ctx context.Context) *UsersConfigListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users config list params
func (o *UsersConfigListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users config list params
func (o *UsersConfigListParams) WithHTTPClient(client *http.Client) *UsersConfigListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users config list params
func (o *UsersConfigListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UsersConfigListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
