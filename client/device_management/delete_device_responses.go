// Code generated by go-swagger; DO NOT EDIT.

package device_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/models"
)

// DeleteDeviceReader is a Reader for the DeleteDevice structure.
type DeleteDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDeviceOK creates a DeleteDeviceOK with default headers values
func NewDeleteDeviceOK() *DeleteDeviceOK {
	return &DeleteDeviceOK{}
}

/*DeleteDeviceOK handles this case with default header values.

The device was successfully removed, or had been removed
previously.
*/
type DeleteDeviceOK struct {
	Payload DeleteDeviceOKBody
}

func (o *DeleteDeviceOK) Error() string {
	return fmt.Sprintf("[DELETE /_matrix/client/r0/devices/{deviceId}][%d] deleteDeviceOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceUnauthorized creates a DeleteDeviceUnauthorized with default headers values
func NewDeleteDeviceUnauthorized() *DeleteDeviceUnauthorized {
	return &DeleteDeviceUnauthorized{}
}

/*DeleteDeviceUnauthorized handles this case with default header values.

The homeserver requires additional authentication information.
*/
type DeleteDeviceUnauthorized struct {
	Payload *models.DeleteDeviceUnauthorizedBody
}

func (o *DeleteDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /_matrix/client/r0/devices/{deviceId}][%d] deleteDeviceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDeviceUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteDeviceOKBody delete device o k body
swagger:model DeleteDeviceOKBody
*/
type DeleteDeviceOKBody interface{}
