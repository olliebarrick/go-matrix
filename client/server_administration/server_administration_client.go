// Code generated by go-swagger; DO NOT EDIT.

package server_administration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new server administration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server administration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetVersions gets the versions of the specification supported by the server

Gets the versions of the specification supported by the server.

Values will take the form ``rX.Y.Z``.

Only the latest ``Z`` value will be reported for each supported ``X.Y`` value.
i.e. if the server implements ``r0.0.0``, ``r0.0.1``, and ``r1.2.0``, it will report ``r0.0.1`` and ``r1.2.0``.
*/
func (a *Client) GetVersions(params *GetVersionsParams) (*GetVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVersions",
		Method:             "GET",
		PathPattern:        "/_matrix/client/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVersionsOK), nil

}

/*
GetWhoIs gets information about a particular user

Gets information about a particular user.

This API may be restricted to only be called by the user being looked
up, or by a server admin. Server-local administrator privileges are not
specified in this document.
*/
func (a *Client) GetWhoIs(params *GetWhoIsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWhoIsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWhoIsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWhoIs",
		Method:             "GET",
		PathPattern:        "/_matrix/client/r0/admin/whois/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWhoIsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWhoIsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
