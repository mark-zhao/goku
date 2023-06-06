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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewIpamPrefixesAvailablePrefixesCreateParams creates a new IpamPrefixesAvailablePrefixesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesAvailablePrefixesCreateParams() *IpamPrefixesAvailablePrefixesCreateParams {
	return &IpamPrefixesAvailablePrefixesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesAvailablePrefixesCreateParamsWithTimeout creates a new IpamPrefixesAvailablePrefixesCreateParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesAvailablePrefixesCreateParamsWithTimeout(timeout time.Duration) *IpamPrefixesAvailablePrefixesCreateParams {
	return &IpamPrefixesAvailablePrefixesCreateParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesAvailablePrefixesCreateParamsWithContext creates a new IpamPrefixesAvailablePrefixesCreateParams object
// with the ability to set a context for a request.
func NewIpamPrefixesAvailablePrefixesCreateParamsWithContext(ctx context.Context) *IpamPrefixesAvailablePrefixesCreateParams {
	return &IpamPrefixesAvailablePrefixesCreateParams{
		Context: ctx,
	}
}

// NewIpamPrefixesAvailablePrefixesCreateParamsWithHTTPClient creates a new IpamPrefixesAvailablePrefixesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesAvailablePrefixesCreateParamsWithHTTPClient(client *http.Client) *IpamPrefixesAvailablePrefixesCreateParams {
	return &IpamPrefixesAvailablePrefixesCreateParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesAvailablePrefixesCreateParams contains all the parameters to send to the API endpoint
   for the ipam prefixes available prefixes create operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesAvailablePrefixesCreateParams struct {

	// Data.
	Data *models.PrefixLength

	/* ID.

	   A unique integer value identifying this prefix.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes available prefixes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithDefaults() *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes available prefixes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithTimeout(timeout time.Duration) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithContext(ctx context.Context) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithHTTPClient(client *http.Client) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithData(data *models.PrefixLength) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetData(data *models.PrefixLength) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) WithID(id int64) *IpamPrefixesAvailablePrefixesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes available prefixes create params
func (o *IpamPrefixesAvailablePrefixesCreateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesAvailablePrefixesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
