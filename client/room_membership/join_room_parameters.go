// Code generated by go-swagger; DO NOT EDIT.

package room_membership

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

// NewJoinRoomParams creates a new JoinRoomParams object
// with the default values initialized.
func NewJoinRoomParams() *JoinRoomParams {
	var ()
	return &JoinRoomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJoinRoomParamsWithTimeout creates a new JoinRoomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJoinRoomParamsWithTimeout(timeout time.Duration) *JoinRoomParams {
	var ()
	return &JoinRoomParams{

		timeout: timeout,
	}
}

// NewJoinRoomParamsWithContext creates a new JoinRoomParams object
// with the default values initialized, and the ability to set a context for a request
func NewJoinRoomParamsWithContext(ctx context.Context) *JoinRoomParams {
	var ()
	return &JoinRoomParams{

		Context: ctx,
	}
}

// NewJoinRoomParamsWithHTTPClient creates a new JoinRoomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJoinRoomParamsWithHTTPClient(client *http.Client) *JoinRoomParams {
	var ()
	return &JoinRoomParams{
		HTTPClient: client,
	}
}

/*JoinRoomParams contains all the parameters to send to the API endpoint
for the join room operation typically these are written to a http.Request
*/
type JoinRoomParams struct {

	/*RoomIDOrAlias
	  The room identifier or alias to join.

	*/
	RoomIDOrAlias string
	/*ThirdPartySigned*/
	ThirdPartySigned *models.JoinRoomParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the join room params
func (o *JoinRoomParams) WithTimeout(timeout time.Duration) *JoinRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the join room params
func (o *JoinRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the join room params
func (o *JoinRoomParams) WithContext(ctx context.Context) *JoinRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the join room params
func (o *JoinRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the join room params
func (o *JoinRoomParams) WithHTTPClient(client *http.Client) *JoinRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the join room params
func (o *JoinRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoomIDOrAlias adds the roomIDOrAlias to the join room params
func (o *JoinRoomParams) WithRoomIDOrAlias(roomIDOrAlias string) *JoinRoomParams {
	o.SetRoomIDOrAlias(roomIDOrAlias)
	return o
}

// SetRoomIDOrAlias adds the roomIdOrAlias to the join room params
func (o *JoinRoomParams) SetRoomIDOrAlias(roomIDOrAlias string) {
	o.RoomIDOrAlias = roomIDOrAlias
}

// WithThirdPartySigned adds the thirdPartySigned to the join room params
func (o *JoinRoomParams) WithThirdPartySigned(thirdPartySigned *models.JoinRoomParamsBody) *JoinRoomParams {
	o.SetThirdPartySigned(thirdPartySigned)
	return o
}

// SetThirdPartySigned adds the thirdPartySigned to the join room params
func (o *JoinRoomParams) SetThirdPartySigned(thirdPartySigned *models.JoinRoomParamsBody) {
	o.ThirdPartySigned = thirdPartySigned
}

// WriteToRequest writes these params to a swagger request
func (o *JoinRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roomIdOrAlias
	if err := r.SetPathParam("roomIdOrAlias", o.RoomIDOrAlias); err != nil {
		return err
	}

	if o.ThirdPartySigned != nil {
		if err := r.SetBodyParam(o.ThirdPartySigned); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
