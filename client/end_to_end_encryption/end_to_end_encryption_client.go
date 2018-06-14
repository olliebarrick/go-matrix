// Code generated by go-swagger; DO NOT EDIT.

package end_to_end_encryption

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new end to end encryption API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for end to end encryption API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ClaimKeys claims one time encryption keys

Claims one-time keys for use in pre-key messages.
*/
func (a *Client) ClaimKeys(params *ClaimKeysParams, authInfo runtime.ClientAuthInfoWriter) (*ClaimKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClaimKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "claimKeys",
		Method:             "POST",
		PathPattern:        "/_matrix/client/r0/keys/claim",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ClaimKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ClaimKeysOK), nil

}

/*
GetKeysChanges queries users with recent device key updates

Gets a list of users who have updated their device identity keys since a
previous sync token.

The server should include in the results any users who:

* currently share a room with the calling user (ie, both users have
  membership state ``join``); *and*
* added new device identity keys or removed an existing device with
  identity keys, between ``from`` and ``to``.
*/
func (a *Client) GetKeysChanges(params *GetKeysChangesParams, authInfo runtime.ClientAuthInfoWriter) (*GetKeysChangesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeysChangesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getKeysChanges",
		Method:             "GET",
		PathPattern:        "/_matrix/client/r0/keys/changes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKeysChangesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKeysChangesOK), nil

}

/*
QueryKeys downloads device identity keys

Returns the current devices and identity keys for the given users.
*/
func (a *Client) QueryKeys(params *QueryKeysParams, authInfo runtime.ClientAuthInfoWriter) (*QueryKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryKeys",
		Method:             "POST",
		PathPattern:        "/_matrix/client/r0/keys/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryKeysOK), nil

}

/*
UploadKeys uploads end to end encryption keys

Publishes end-to-end encryption keys for the device.
*/
func (a *Client) UploadKeys(params *UploadKeysParams, authInfo runtime.ClientAuthInfoWriter) (*UploadKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadKeys",
		Method:             "POST",
		PathPattern:        "/_matrix/client/r0/keys/upload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadKeysOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
