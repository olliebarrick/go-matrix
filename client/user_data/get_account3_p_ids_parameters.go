// Code generated by go-swagger; DO NOT EDIT.

package user_data

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

// NewGetAccount3PIdsParams creates a new GetAccount3PIdsParams object
// with the default values initialized.
func NewGetAccount3PIdsParams() *GetAccount3PIdsParams {

	return &GetAccount3PIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccount3PIdsParamsWithTimeout creates a new GetAccount3PIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccount3PIdsParamsWithTimeout(timeout time.Duration) *GetAccount3PIdsParams {

	return &GetAccount3PIdsParams{

		timeout: timeout,
	}
}

// NewGetAccount3PIdsParamsWithContext creates a new GetAccount3PIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccount3PIdsParamsWithContext(ctx context.Context) *GetAccount3PIdsParams {

	return &GetAccount3PIdsParams{

		Context: ctx,
	}
}

// NewGetAccount3PIdsParamsWithHTTPClient creates a new GetAccount3PIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccount3PIdsParamsWithHTTPClient(client *http.Client) *GetAccount3PIdsParams {

	return &GetAccount3PIdsParams{
		HTTPClient: client,
	}
}

/*GetAccount3PIdsParams contains all the parameters to send to the API endpoint
for the get account3 p ids operation typically these are written to a http.Request
*/
type GetAccount3PIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account3 p ids params
func (o *GetAccount3PIdsParams) WithTimeout(timeout time.Duration) *GetAccount3PIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account3 p ids params
func (o *GetAccount3PIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account3 p ids params
func (o *GetAccount3PIdsParams) WithContext(ctx context.Context) *GetAccount3PIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account3 p ids params
func (o *GetAccount3PIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account3 p ids params
func (o *GetAccount3PIdsParams) WithHTTPClient(client *http.Client) *GetAccount3PIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account3 p ids params
func (o *GetAccount3PIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccount3PIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
