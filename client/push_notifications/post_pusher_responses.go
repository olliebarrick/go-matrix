// Code generated by go-swagger; DO NOT EDIT.

package push_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/models"
)

// PostPusherReader is a Reader for the PostPusher structure.
type PostPusherReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPusherReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPusherOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPusherBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostPusherTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPusherOK creates a PostPusherOK with default headers values
func NewPostPusherOK() *PostPusherOK {
	return &PostPusherOK{}
}

/*PostPusherOK handles this case with default header values.

The pusher was set.
*/
type PostPusherOK struct {
	Payload PostPusherOKBody
}

func (o *PostPusherOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/pushers/set][%d] postPusherOK  %+v", 200, o.Payload)
}

func (o *PostPusherOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPusherBadRequest creates a PostPusherBadRequest with default headers values
func NewPostPusherBadRequest() *PostPusherBadRequest {
	return &PostPusherBadRequest{}
}

/*PostPusherBadRequest handles this case with default header values.

One or more of the pusher values were invalid.
*/
type PostPusherBadRequest struct {
	Payload PostPusherBadRequestBody
}

func (o *PostPusherBadRequest) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/pushers/set][%d] postPusherBadRequest  %+v", 400, o.Payload)
}

func (o *PostPusherBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPusherTooManyRequests creates a PostPusherTooManyRequests with default headers values
func NewPostPusherTooManyRequests() *PostPusherTooManyRequests {
	return &PostPusherTooManyRequests{}
}

/*PostPusherTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type PostPusherTooManyRequests struct {
	Payload *models.PostPusherTooManyRequestsBody
}

func (o *PostPusherTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/pushers/set][%d] postPusherTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPusherTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostPusherTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostPusherBadRequestBody post pusher bad request body
swagger:model PostPusherBadRequestBody
*/
type PostPusherBadRequestBody interface{}

/*PostPusherOKBody post pusher o k body
swagger:model PostPusherOKBody
*/
type PostPusherOKBody interface{}
