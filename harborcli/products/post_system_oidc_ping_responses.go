// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostSystemOidcPingReader is a Reader for the PostSystemOidcPing structure.
type PostSystemOidcPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSystemOidcPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSystemOidcPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSystemOidcPingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSystemOidcPingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSystemOidcPingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSystemOidcPingOK creates a PostSystemOidcPingOK with default headers values
func NewPostSystemOidcPingOK() *PostSystemOidcPingOK {
	return &PostSystemOidcPingOK{}
}

/*PostSystemOidcPingOK handles this case with default header values.

Ping succeeded.  The OIDC endpoint is valid.
*/
type PostSystemOidcPingOK struct {
}

func (o *PostSystemOidcPingOK) Error() string {
	return fmt.Sprintf("[POST /system/oidc/ping][%d] postSystemOidcPingOK ", 200)
}

func (o *PostSystemOidcPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemOidcPingBadRequest creates a PostSystemOidcPingBadRequest with default headers values
func NewPostSystemOidcPingBadRequest() *PostSystemOidcPingBadRequest {
	return &PostSystemOidcPingBadRequest{}
}

/*PostSystemOidcPingBadRequest handles this case with default header values.

The ping failed
*/
type PostSystemOidcPingBadRequest struct {
}

func (o *PostSystemOidcPingBadRequest) Error() string {
	return fmt.Sprintf("[POST /system/oidc/ping][%d] postSystemOidcPingBadRequest ", 400)
}

func (o *PostSystemOidcPingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemOidcPingUnauthorized creates a PostSystemOidcPingUnauthorized with default headers values
func NewPostSystemOidcPingUnauthorized() *PostSystemOidcPingUnauthorized {
	return &PostSystemOidcPingUnauthorized{}
}

/*PostSystemOidcPingUnauthorized handles this case with default header values.

User need to log in first.
*/
type PostSystemOidcPingUnauthorized struct {
}

func (o *PostSystemOidcPingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /system/oidc/ping][%d] postSystemOidcPingUnauthorized ", 401)
}

func (o *PostSystemOidcPingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemOidcPingForbidden creates a PostSystemOidcPingForbidden with default headers values
func NewPostSystemOidcPingForbidden() *PostSystemOidcPingForbidden {
	return &PostSystemOidcPingForbidden{}
}

/*PostSystemOidcPingForbidden handles this case with default header values.

User does not have permission to call this API
*/
type PostSystemOidcPingForbidden struct {
}

func (o *PostSystemOidcPingForbidden) Error() string {
	return fmt.Sprintf("[POST /system/oidc/ping][%d] postSystemOidcPingForbidden ", 403)
}

func (o *PostSystemOidcPingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostSystemOidcPingBody post system oidc ping body
swagger:model PostSystemOidcPingBody
*/
type PostSystemOidcPingBody struct {

	// The URL of OIDC endpoint to be tested.
	URL string `json:"url,omitempty"`

	// Whether the certificate should be verified
	VerifyCert bool `json:"verify_cert,omitempty"`
}

// Validate validates this post system oidc ping body
func (o *PostSystemOidcPingBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostSystemOidcPingBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSystemOidcPingBody) UnmarshalBinary(b []byte) error {
	var res PostSystemOidcPingBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
