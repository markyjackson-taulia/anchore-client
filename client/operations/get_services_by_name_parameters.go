// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServicesByNameParams creates a new GetServicesByNameParams object
// with the default values initialized.
func NewGetServicesByNameParams() *GetServicesByNameParams {
	var ()
	return &GetServicesByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicesByNameParamsWithTimeout creates a new GetServicesByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServicesByNameParamsWithTimeout(timeout time.Duration) *GetServicesByNameParams {
	var ()
	return &GetServicesByNameParams{

		timeout: timeout,
	}
}

// NewGetServicesByNameParamsWithContext creates a new GetServicesByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServicesByNameParamsWithContext(ctx context.Context) *GetServicesByNameParams {
	var ()
	return &GetServicesByNameParams{

		Context: ctx,
	}
}

// NewGetServicesByNameParamsWithHTTPClient creates a new GetServicesByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServicesByNameParamsWithHTTPClient(client *http.Client) *GetServicesByNameParams {
	var ()
	return &GetServicesByNameParams{
		HTTPClient: client,
	}
}

/*GetServicesByNameParams contains all the parameters to send to the API endpoint
for the get services by name operation typically these are written to a http.Request
*/
type GetServicesByNameParams struct {

	/*Servicename*/
	Servicename string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get services by name params
func (o *GetServicesByNameParams) WithTimeout(timeout time.Duration) *GetServicesByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get services by name params
func (o *GetServicesByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get services by name params
func (o *GetServicesByNameParams) WithContext(ctx context.Context) *GetServicesByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get services by name params
func (o *GetServicesByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get services by name params
func (o *GetServicesByNameParams) WithHTTPClient(client *http.Client) *GetServicesByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get services by name params
func (o *GetServicesByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServicename adds the servicename to the get services by name params
func (o *GetServicesByNameParams) WithServicename(servicename string) *GetServicesByNameParams {
	o.SetServicename(servicename)
	return o
}

// SetServicename adds the servicename to the get services by name params
func (o *GetServicesByNameParams) SetServicename(servicename string) {
	o.Servicename = servicename
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicesByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param servicename
	if err := r.SetPathParam("servicename", o.Servicename); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}