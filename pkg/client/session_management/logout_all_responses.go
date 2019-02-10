// Code generated by go-swagger; DO NOT EDIT.

package session_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// LogoutAllReader is a Reader for the LogoutAll structure.
type LogoutAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLogoutAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLogoutAllOK creates a LogoutAllOK with default headers values
func NewLogoutAllOK() *LogoutAllOK {
	return &LogoutAllOK{}
}

/*LogoutAllOK handles this case with default header values.

The user's access tokens were succesfully invalidated.
*/
type LogoutAllOK struct {
	Payload LogoutAllOKBody
}

func (o *LogoutAllOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/logout/all][%d] logoutAllOK  %+v", 200, o.Payload)
}

func (o *LogoutAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LogoutAllOKBody logout all o k body
swagger:model LogoutAllOKBody
*/
type LogoutAllOKBody interface{}
