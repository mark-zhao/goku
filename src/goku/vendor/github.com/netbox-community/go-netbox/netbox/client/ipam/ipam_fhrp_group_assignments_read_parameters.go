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

package ipam

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

// NewIpamFhrpGroupAssignmentsReadParams creates a new IpamFhrpGroupAssignmentsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamFhrpGroupAssignmentsReadParams() *IpamFhrpGroupAssignmentsReadParams {
	return &IpamFhrpGroupAssignmentsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamFhrpGroupAssignmentsReadParamsWithTimeout creates a new IpamFhrpGroupAssignmentsReadParams object
// with the ability to set a timeout on a request.
func NewIpamFhrpGroupAssignmentsReadParamsWithTimeout(timeout time.Duration) *IpamFhrpGroupAssignmentsReadParams {
	return &IpamFhrpGroupAssignmentsReadParams{
		timeout: timeout,
	}
}

// NewIpamFhrpGroupAssignmentsReadParamsWithContext creates a new IpamFhrpGroupAssignmentsReadParams object
// with the ability to set a context for a request.
func NewIpamFhrpGroupAssignmentsReadParamsWithContext(ctx context.Context) *IpamFhrpGroupAssignmentsReadParams {
	return &IpamFhrpGroupAssignmentsReadParams{
		Context: ctx,
	}
}

// NewIpamFhrpGroupAssignmentsReadParamsWithHTTPClient creates a new IpamFhrpGroupAssignmentsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamFhrpGroupAssignmentsReadParamsWithHTTPClient(client *http.Client) *IpamFhrpGroupAssignmentsReadParams {
	return &IpamFhrpGroupAssignmentsReadParams{
		HTTPClient: client,
	}
}

/* IpamFhrpGroupAssignmentsReadParams contains all the parameters to send to the API endpoint
   for the ipam fhrp group assignments read operation.

   Typically these are written to a http.Request.
*/
type IpamFhrpGroupAssignmentsReadParams struct {

	/* ID.

	   A unique integer value identifying this FHRP group assignment.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam fhrp group assignments read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupAssignmentsReadParams) WithDefaults() *IpamFhrpGroupAssignmentsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam fhrp group assignments read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupAssignmentsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam fhrp group assignments read params
func (o *IpamFhrpGroupAssignmentsReadParams) WithTimeout(timeout time.Duration) *IpamFhrpGroupAssignmentsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam fhrp group assignments read params
func (o *IpamFhrpGroupAssignmentsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam fhrp group assignments read params
func (o *IpamFhrpGroupAssignmentsReadParams) WithContext(ctx context.Context) *IpamFhrpGroupAssignmentsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam fhrp group assignments read params
func (o *IpamFhrpGroupAssignmentsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam fhrp group assignments read params
func (o *IpamFhrpGroupAssignmentsReadParams) WithHTTPClient(client *http.Client) *IpamFhrpGroupAssignmentsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam fhrp group assignments read params
func (o *IpamFhrpGroupAssignmentsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam fhrp group assignments read params
func (o *IpamFhrpGroupAssignmentsReadParams) WithID(id int64) *IpamFhrpGroupAssignmentsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam fhrp group assignments read params
func (o *IpamFhrpGroupAssignmentsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamFhrpGroupAssignmentsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
