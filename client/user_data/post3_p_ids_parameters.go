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

	models "github.com/justinbarrick/go-matrix/models"
)

// NewPost3PIdsParams creates a new Post3PIdsParams object
// with the default values initialized.
func NewPost3PIdsParams() *Post3PIdsParams {
	var ()
	return &Post3PIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPost3PIdsParamsWithTimeout creates a new Post3PIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPost3PIdsParamsWithTimeout(timeout time.Duration) *Post3PIdsParams {
	var ()
	return &Post3PIdsParams{

		timeout: timeout,
	}
}

// NewPost3PIdsParamsWithContext creates a new Post3PIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPost3PIdsParamsWithContext(ctx context.Context) *Post3PIdsParams {
	var ()
	return &Post3PIdsParams{

		Context: ctx,
	}
}

// NewPost3PIdsParamsWithHTTPClient creates a new Post3PIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPost3PIdsParamsWithHTTPClient(client *http.Client) *Post3PIdsParams {
	var ()
	return &Post3PIdsParams{
		HTTPClient: client,
	}
}

/*Post3PIdsParams contains all the parameters to send to the API endpoint
for the post3 p ids operation typically these are written to a http.Request
*/
type Post3PIdsParams struct {

	/*Body*/
	Body *models.Post3PIdsParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post3 p ids params
func (o *Post3PIdsParams) WithTimeout(timeout time.Duration) *Post3PIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post3 p ids params
func (o *Post3PIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post3 p ids params
func (o *Post3PIdsParams) WithContext(ctx context.Context) *Post3PIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post3 p ids params
func (o *Post3PIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post3 p ids params
func (o *Post3PIdsParams) WithHTTPClient(client *http.Client) *Post3PIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post3 p ids params
func (o *Post3PIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post3 p ids params
func (o *Post3PIdsParams) WithBody(body *models.Post3PIdsParamsBody) *Post3PIdsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post3 p ids params
func (o *Post3PIdsParams) SetBody(body *models.Post3PIdsParamsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *Post3PIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
