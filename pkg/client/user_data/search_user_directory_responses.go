// Code generated by go-swagger; DO NOT EDIT.

package user_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/pkg/models"
)

// SearchUserDirectoryReader is a Reader for the SearchUserDirectory structure.
type SearchUserDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUserDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchUserDirectoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 429:
		result := NewSearchUserDirectoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchUserDirectoryOK creates a SearchUserDirectoryOK with default headers values
func NewSearchUserDirectoryOK() *SearchUserDirectoryOK {
	return &SearchUserDirectoryOK{}
}

/*SearchUserDirectoryOK handles this case with default header values.

The results of the search.
*/
type SearchUserDirectoryOK struct {
	Payload *models.SearchUserDirectoryOKBody
}

func (o *SearchUserDirectoryOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/user_directory/search][%d] searchUserDirectoryOK  %+v", 200, o.Payload)
}

func (o *SearchUserDirectoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchUserDirectoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserDirectoryTooManyRequests creates a SearchUserDirectoryTooManyRequests with default headers values
func NewSearchUserDirectoryTooManyRequests() *SearchUserDirectoryTooManyRequests {
	return &SearchUserDirectoryTooManyRequests{}
}

/*SearchUserDirectoryTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type SearchUserDirectoryTooManyRequests struct {
	Payload *models.SearchUserDirectoryTooManyRequestsBody
}

func (o *SearchUserDirectoryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/user_directory/search][%d] searchUserDirectoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *SearchUserDirectoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchUserDirectoryTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
