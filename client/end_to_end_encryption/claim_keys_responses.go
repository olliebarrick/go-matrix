// Code generated by go-swagger; DO NOT EDIT.

package end_to_end_encryption

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/models"
)

// ClaimKeysReader is a Reader for the ClaimKeys structure.
type ClaimKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClaimKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewClaimKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClaimKeysOK creates a ClaimKeysOK with default headers values
func NewClaimKeysOK() *ClaimKeysOK {
	return &ClaimKeysOK{}
}

/*ClaimKeysOK handles this case with default header values.

The claimed keys
*/
type ClaimKeysOK struct {
	Payload *models.ClaimKeysOKBody
}

func (o *ClaimKeysOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/r0/keys/claim][%d] claimKeysOK  %+v", 200, o.Payload)
}

func (o *ClaimKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClaimKeysOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
