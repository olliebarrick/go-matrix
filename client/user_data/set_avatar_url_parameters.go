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

// NewSetAvatarURLParams creates a new SetAvatarURLParams object
// with the default values initialized.
func NewSetAvatarURLParams() *SetAvatarURLParams {
	var ()
	return &SetAvatarURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetAvatarURLParamsWithTimeout creates a new SetAvatarURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetAvatarURLParamsWithTimeout(timeout time.Duration) *SetAvatarURLParams {
	var ()
	return &SetAvatarURLParams{

		timeout: timeout,
	}
}

// NewSetAvatarURLParamsWithContext creates a new SetAvatarURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetAvatarURLParamsWithContext(ctx context.Context) *SetAvatarURLParams {
	var ()
	return &SetAvatarURLParams{

		Context: ctx,
	}
}

// NewSetAvatarURLParamsWithHTTPClient creates a new SetAvatarURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetAvatarURLParamsWithHTTPClient(client *http.Client) *SetAvatarURLParams {
	var ()
	return &SetAvatarURLParams{
		HTTPClient: client,
	}
}

/*SetAvatarURLParams contains all the parameters to send to the API endpoint
for the set avatar Url operation typically these are written to a http.Request
*/
type SetAvatarURLParams struct {

	/*AvatarURL
	  The avatar url info.

	*/
	AvatarURL *models.SetAvatarURLParamsBody
	/*UserID
	  The user whose avatar URL to set.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set avatar Url params
func (o *SetAvatarURLParams) WithTimeout(timeout time.Duration) *SetAvatarURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set avatar Url params
func (o *SetAvatarURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set avatar Url params
func (o *SetAvatarURLParams) WithContext(ctx context.Context) *SetAvatarURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set avatar Url params
func (o *SetAvatarURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set avatar Url params
func (o *SetAvatarURLParams) WithHTTPClient(client *http.Client) *SetAvatarURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set avatar Url params
func (o *SetAvatarURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvatarURL adds the avatarURL to the set avatar Url params
func (o *SetAvatarURLParams) WithAvatarURL(avatarURL *models.SetAvatarURLParamsBody) *SetAvatarURLParams {
	o.SetAvatarURL(avatarURL)
	return o
}

// SetAvatarURL adds the avatarUrl to the set avatar Url params
func (o *SetAvatarURLParams) SetAvatarURL(avatarURL *models.SetAvatarURLParamsBody) {
	o.AvatarURL = avatarURL
}

// WithUserID adds the userID to the set avatar Url params
func (o *SetAvatarURLParams) WithUserID(userID string) *SetAvatarURLParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set avatar Url params
func (o *SetAvatarURLParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetAvatarURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvatarURL != nil {
		if err := r.SetBodyParam(o.AvatarURL); err != nil {
			return err
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
