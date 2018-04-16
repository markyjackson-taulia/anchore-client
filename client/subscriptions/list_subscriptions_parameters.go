// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewListSubscriptionsParams creates a new ListSubscriptionsParams object
// with the default values initialized.
func NewListSubscriptionsParams() *ListSubscriptionsParams {

	return &ListSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSubscriptionsParamsWithTimeout creates a new ListSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSubscriptionsParamsWithTimeout(timeout time.Duration) *ListSubscriptionsParams {

	return &ListSubscriptionsParams{

		timeout: timeout,
	}
}

// NewListSubscriptionsParamsWithContext creates a new ListSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSubscriptionsParamsWithContext(ctx context.Context) *ListSubscriptionsParams {

	return &ListSubscriptionsParams{

		Context: ctx,
	}
}

// NewListSubscriptionsParamsWithHTTPClient creates a new ListSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSubscriptionsParamsWithHTTPClient(client *http.Client) *ListSubscriptionsParams {

	return &ListSubscriptionsParams{
		HTTPClient: client,
	}
}

/*ListSubscriptionsParams contains all the parameters to send to the API endpoint
for the list subscriptions operation typically these are written to a http.Request
*/
type ListSubscriptionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list subscriptions params
func (o *ListSubscriptionsParams) WithTimeout(timeout time.Duration) *ListSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list subscriptions params
func (o *ListSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list subscriptions params
func (o *ListSubscriptionsParams) WithContext(ctx context.Context) *ListSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list subscriptions params
func (o *ListSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list subscriptions params
func (o *ListSubscriptionsParams) WithHTTPClient(client *http.Client) *ListSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list subscriptions params
func (o *ListSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}