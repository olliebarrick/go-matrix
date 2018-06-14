// Code generated by go-swagger; DO NOT EDIT.

package room_membership

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/models"
)

// JoinRoomReader is a Reader for the JoinRoom structure.
type JoinRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JoinRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJoinRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewJoinRoomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewJoinRoomTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJoinRoomOK creates a JoinRoomOK with default headers values
func NewJoinRoomOK() *JoinRoomOK {
	return &JoinRoomOK{}
}

/*JoinRoomOK handles this case with default header values.

The room has been joined.

The joined room ID must be returned in the ``room_id`` field.
*/
type JoinRoomOK struct {
	Payload JoinRoomOKBody
}

func (o *JoinRoomOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/join/{roomIdOrAlias}][%d] joinRoomOK  %+v", 200, o.Payload)
}

func (o *JoinRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJoinRoomForbidden creates a JoinRoomForbidden with default headers values
func NewJoinRoomForbidden() *JoinRoomForbidden {
	return &JoinRoomForbidden{}
}

/*JoinRoomForbidden handles this case with default header values.

You do not have permission to join the room. A meaningful ``errcode`` and description error text will be returned. Example reasons for rejection are:

- The room is invite-only and the user was not invited.
- The user has been banned from the room.
*/
type JoinRoomForbidden struct {
}

func (o *JoinRoomForbidden) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/join/{roomIdOrAlias}][%d] joinRoomForbidden ", 403)
}

func (o *JoinRoomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJoinRoomTooManyRequests creates a JoinRoomTooManyRequests with default headers values
func NewJoinRoomTooManyRequests() *JoinRoomTooManyRequests {
	return &JoinRoomTooManyRequests{}
}

/*JoinRoomTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type JoinRoomTooManyRequests struct {
	Payload *models.JoinRoomTooManyRequestsBody
}

func (o *JoinRoomTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/join/{roomIdOrAlias}][%d] joinRoomTooManyRequests  %+v", 429, o.Payload)
}

func (o *JoinRoomTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JoinRoomTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JoinRoomOKBody join room o k body
swagger:model JoinRoomOKBody
*/
type JoinRoomOKBody interface{}
