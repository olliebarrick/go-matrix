// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new presence API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for presence API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPresence gets this user s presence state

Get the given user's presence state.
*/
func (a *Client) GetPresence(params *GetPresenceParams) (*GetPresenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPresenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPresence",
		Method:             "GET",
		PathPattern:        "/_matrix/client/r0/presence/{userId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPresenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPresenceOK), nil

}

/*
GetPresenceForList gets presence events for this presence list

Retrieve a list of presence events for every user on this list.
*/
func (a *Client) GetPresenceForList(params *GetPresenceForListParams) (*GetPresenceForListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPresenceForListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPresenceForList",
		Method:             "GET",
		PathPattern:        "/_matrix/client/r0/presence/list/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPresenceForListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPresenceForListOK), nil

}

/*
ModifyPresenceList adds or remove users from this presence list

Adds or removes users from this presence list.
*/
func (a *Client) ModifyPresenceList(params *ModifyPresenceListParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyPresenceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyPresenceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyPresenceList",
		Method:             "POST",
		PathPattern:        "/_matrix/client/r0/presence/list/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyPresenceListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyPresenceListOK), nil

}

/*
SetPresence updates this user s presence state

This API sets the given user's presence state. When setting the status,
the activity time is updated to reflect that activity; the client does
not need to specify the ``last_active_ago`` field. You cannot set the
presence state of another user.
*/
func (a *Client) SetPresence(params *SetPresenceParams, authInfo runtime.ClientAuthInfoWriter) (*SetPresenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetPresenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setPresence",
		Method:             "PUT",
		PathPattern:        "/_matrix/client/r0/presence/{userId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetPresenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetPresenceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
