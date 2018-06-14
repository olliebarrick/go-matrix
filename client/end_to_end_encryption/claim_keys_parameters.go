// Code generated by go-swagger; DO NOT EDIT.

package end_to_end_encryption

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

// NewClaimKeysParams creates a new ClaimKeysParams object
// with the default values initialized.
func NewClaimKeysParams() *ClaimKeysParams {
	var ()
	return &ClaimKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewClaimKeysParamsWithTimeout creates a new ClaimKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewClaimKeysParamsWithTimeout(timeout time.Duration) *ClaimKeysParams {
	var ()
	return &ClaimKeysParams{

		timeout: timeout,
	}
}

// NewClaimKeysParamsWithContext creates a new ClaimKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewClaimKeysParamsWithContext(ctx context.Context) *ClaimKeysParams {
	var ()
	return &ClaimKeysParams{

		Context: ctx,
	}
}

// NewClaimKeysParamsWithHTTPClient creates a new ClaimKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewClaimKeysParamsWithHTTPClient(client *http.Client) *ClaimKeysParams {
	var ()
	return &ClaimKeysParams{
		HTTPClient: client,
	}
}

/*ClaimKeysParams contains all the parameters to send to the API endpoint
for the claim keys operation typically these are written to a http.Request
*/
type ClaimKeysParams struct {

	/*Query
	  Query defining the keys to be claimed

	*/
	Query *models.ClaimKeysParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the claim keys params
func (o *ClaimKeysParams) WithTimeout(timeout time.Duration) *ClaimKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the claim keys params
func (o *ClaimKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the claim keys params
func (o *ClaimKeysParams) WithContext(ctx context.Context) *ClaimKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the claim keys params
func (o *ClaimKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the claim keys params
func (o *ClaimKeysParams) WithHTTPClient(client *http.Client) *ClaimKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the claim keys params
func (o *ClaimKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the claim keys params
func (o *ClaimKeysParams) WithQuery(query *models.ClaimKeysParamsBody) *ClaimKeysParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the claim keys params
func (o *ClaimKeysParams) SetQuery(query *models.ClaimKeysParamsBody) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *ClaimKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Query != nil {
		if err := r.SetBodyParam(o.Query); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
