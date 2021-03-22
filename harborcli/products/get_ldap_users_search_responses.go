// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"gitlab.4pd.io/liwenhao/pineapple/pineapple/apigen/harborcli/models"
)

// GetLdapUsersSearchReader is a Reader for the GetLdapUsersSearch structure.
type GetLdapUsersSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLdapUsersSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLdapUsersSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLdapUsersSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLdapUsersSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLdapUsersSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLdapUsersSearchOK creates a GetLdapUsersSearchOK with default headers values
func NewGetLdapUsersSearchOK() *GetLdapUsersSearchOK {
	return &GetLdapUsersSearchOK{}
}

/*GetLdapUsersSearchOK handles this case with default header values.

Search ldap users successfully.
*/
type GetLdapUsersSearchOK struct {
	Payload []*models.LdapUsers
}

func (o *GetLdapUsersSearchOK) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] getLdapUsersSearchOK  %+v", 200, o.Payload)
}

func (o *GetLdapUsersSearchOK) GetPayload() []*models.LdapUsers {
	return o.Payload
}

func (o *GetLdapUsersSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLdapUsersSearchUnauthorized creates a GetLdapUsersSearchUnauthorized with default headers values
func NewGetLdapUsersSearchUnauthorized() *GetLdapUsersSearchUnauthorized {
	return &GetLdapUsersSearchUnauthorized{}
}

/*GetLdapUsersSearchUnauthorized handles this case with default header values.

User need to login first.
*/
type GetLdapUsersSearchUnauthorized struct {
}

func (o *GetLdapUsersSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] getLdapUsersSearchUnauthorized ", 401)
}

func (o *GetLdapUsersSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLdapUsersSearchForbidden creates a GetLdapUsersSearchForbidden with default headers values
func NewGetLdapUsersSearchForbidden() *GetLdapUsersSearchForbidden {
	return &GetLdapUsersSearchForbidden{}
}

/*GetLdapUsersSearchForbidden handles this case with default header values.

Only admin has this authority.
*/
type GetLdapUsersSearchForbidden struct {
}

func (o *GetLdapUsersSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] getLdapUsersSearchForbidden ", 403)
}

func (o *GetLdapUsersSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLdapUsersSearchInternalServerError creates a GetLdapUsersSearchInternalServerError with default headers values
func NewGetLdapUsersSearchInternalServerError() *GetLdapUsersSearchInternalServerError {
	return &GetLdapUsersSearchInternalServerError{}
}

/*GetLdapUsersSearchInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetLdapUsersSearchInternalServerError struct {
}

func (o *GetLdapUsersSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] getLdapUsersSearchInternalServerError ", 500)
}

func (o *GetLdapUsersSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
