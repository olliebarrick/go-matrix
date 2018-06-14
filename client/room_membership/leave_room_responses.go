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

// LeaveRoomReader is a Reader for the LeaveRoom structure.
type LeaveRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LeaveRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLeaveRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 429:
		result := NewLeaveRoomTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLeaveRoomOK creates a LeaveRoomOK with default headers values
func NewLeaveRoomOK() *LeaveRoomOK {
	return &LeaveRoomOK{}
}

/*LeaveRoomOK handles this case with default header values.

The room has been left.
*/
type LeaveRoomOK struct {
	Payload LeaveRoomOKBody
}

func (o *LeaveRoomOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/rooms/{roomId}/leave][%d] leaveRoomOK  %+v", 200, o.Payload)
}

func (o *LeaveRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaveRoomTooManyRequests creates a LeaveRoomTooManyRequests with default headers values
func NewLeaveRoomTooManyRequests() *LeaveRoomTooManyRequests {
	return &LeaveRoomTooManyRequests{}
}

/*LeaveRoomTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type LeaveRoomTooManyRequests struct {
	Payload *models.LeaveRoomTooManyRequestsBody
}

func (o *LeaveRoomTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/rooms/{roomId}/leave][%d] leaveRoomTooManyRequests  %+v", 429, o.Payload)
}

func (o *LeaveRoomTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LeaveRoomTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LeaveRoomOKBody leave room o k body
swagger:model LeaveRoomOKBody
*/
type LeaveRoomOKBody interface{}
