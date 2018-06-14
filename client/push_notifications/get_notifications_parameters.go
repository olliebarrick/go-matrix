// Code generated by go-swagger; DO NOT EDIT.

package push_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNotificationsParams creates a new GetNotificationsParams object
// with the default values initialized.
func NewGetNotificationsParams() *GetNotificationsParams {
	var ()
	return &GetNotificationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNotificationsParamsWithTimeout creates a new GetNotificationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNotificationsParamsWithTimeout(timeout time.Duration) *GetNotificationsParams {
	var ()
	return &GetNotificationsParams{

		timeout: timeout,
	}
}

// NewGetNotificationsParamsWithContext creates a new GetNotificationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNotificationsParamsWithContext(ctx context.Context) *GetNotificationsParams {
	var ()
	return &GetNotificationsParams{

		Context: ctx,
	}
}

// NewGetNotificationsParamsWithHTTPClient creates a new GetNotificationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNotificationsParamsWithHTTPClient(client *http.Client) *GetNotificationsParams {
	var ()
	return &GetNotificationsParams{
		HTTPClient: client,
	}
}

/*GetNotificationsParams contains all the parameters to send to the API endpoint
for the get notifications operation typically these are written to a http.Request
*/
type GetNotificationsParams struct {

	/*From
	  Pagination token given to retrieve the next set of events.

	*/
	From *string
	/*Limit
	  Limit on the number of events to return in this request.

	*/
	Limit *float64
	/*Only
	  Allows basic filtering of events returned. Supply ``highlight``
	to return only events where the notification had the highlight
	tweak set.

	*/
	Only *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get notifications params
func (o *GetNotificationsParams) WithTimeout(timeout time.Duration) *GetNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get notifications params
func (o *GetNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get notifications params
func (o *GetNotificationsParams) WithContext(ctx context.Context) *GetNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get notifications params
func (o *GetNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get notifications params
func (o *GetNotificationsParams) WithHTTPClient(client *http.Client) *GetNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get notifications params
func (o *GetNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get notifications params
func (o *GetNotificationsParams) WithFrom(from *string) *GetNotificationsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get notifications params
func (o *GetNotificationsParams) SetFrom(from *string) {
	o.From = from
}

// WithLimit adds the limit to the get notifications params
func (o *GetNotificationsParams) WithLimit(limit *float64) *GetNotificationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get notifications params
func (o *GetNotificationsParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithOnly adds the only to the get notifications params
func (o *GetNotificationsParams) WithOnly(only *string) *GetNotificationsParams {
	o.SetOnly(only)
	return o
}

// SetOnly adds the only to the get notifications params
func (o *GetNotificationsParams) SetOnly(only *string) {
	o.Only = only
}

// WriteToRequest writes these params to a swagger request
func (o *GetNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit float64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Only != nil {

		// query param only
		var qrOnly string
		if o.Only != nil {
			qrOnly = *o.Only
		}
		qOnly := qrOnly
		if qOnly != "" {
			if err := r.SetQueryParam("only", qOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
