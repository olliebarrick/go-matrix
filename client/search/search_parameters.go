// Code generated by go-swagger; DO NOT EDIT.

package search

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

	models "github.com/justinbarrick/go-matrix/models"
)

// NewSearchParams creates a new SearchParams object
// with the default values initialized.
func NewSearchParams() *SearchParams {
	var ()
	return &SearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchParamsWithTimeout creates a new SearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchParamsWithTimeout(timeout time.Duration) *SearchParams {
	var ()
	return &SearchParams{

		timeout: timeout,
	}
}

// NewSearchParamsWithContext creates a new SearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchParamsWithContext(ctx context.Context) *SearchParams {
	var ()
	return &SearchParams{

		Context: ctx,
	}
}

// NewSearchParamsWithHTTPClient creates a new SearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchParamsWithHTTPClient(client *http.Client) *SearchParams {
	var ()
	return &SearchParams{
		HTTPClient: client,
	}
}

/*SearchParams contains all the parameters to send to the API endpoint
for the search operation typically these are written to a http.Request
*/
type SearchParams struct {

	/*Body*/
	Body *models.SearchParamsBody
	/*NextBatch
	  The point to return events from. If given, this should be a
	`next_batch` result from a previous call to this endpoint.

	*/
	NextBatch *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search params
func (o *SearchParams) WithTimeout(timeout time.Duration) *SearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search params
func (o *SearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search params
func (o *SearchParams) WithContext(ctx context.Context) *SearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search params
func (o *SearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) WithHTTPClient(client *http.Client) *SearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search params
func (o *SearchParams) WithBody(body *models.SearchParamsBody) *SearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search params
func (o *SearchParams) SetBody(body *models.SearchParamsBody) {
	o.Body = body
}

// WithNextBatch adds the nextBatch to the search params
func (o *SearchParams) WithNextBatch(nextBatch *string) *SearchParams {
	o.SetNextBatch(nextBatch)
	return o
}

// SetNextBatch adds the nextBatch to the search params
func (o *SearchParams) SetNextBatch(nextBatch *string) {
	o.NextBatch = nextBatch
}

// WriteToRequest writes these params to a swagger request
func (o *SearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.NextBatch != nil {

		// query param next_batch
		var qrNextBatch string
		if o.NextBatch != nil {
			qrNextBatch = *o.NextBatch
		}
		qNextBatch := qrNextBatch
		if qNextBatch != "" {
			if err := r.SetQueryParam("next_batch", qNextBatch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
