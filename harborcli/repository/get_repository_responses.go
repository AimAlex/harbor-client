// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"gitlab.4pd.io/liwenhao/pineapple/pineapple/apigen/harborcli/models"
)

// GetRepositoryReader is a Reader for the GetRepository structure.
type GetRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoryOK creates a GetRepositoryOK with default headers values
func NewGetRepositoryOK() *GetRepositoryOK {
	return &GetRepositoryOK{}
}

/*GetRepositoryOK handles this case with default header values.

Success
*/
type GetRepositoryOK struct {
	Payload *models.Repository
}

func (o *GetRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetRepositoryOK) GetPayload() *models.Repository {
	return o.Payload
}

func (o *GetRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryBadRequest creates a GetRepositoryBadRequest with default headers values
func NewGetRepositoryBadRequest() *GetRepositoryBadRequest {
	return &GetRepositoryBadRequest{}
}

/*GetRepositoryBadRequest handles this case with default header values.

Bad request
*/
type GetRepositoryBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepositoryBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryUnauthorized creates a GetRepositoryUnauthorized with default headers values
func NewGetRepositoryUnauthorized() *GetRepositoryUnauthorized {
	return &GetRepositoryUnauthorized{}
}

/*GetRepositoryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRepositoryUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepositoryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryForbidden creates a GetRepositoryForbidden with default headers values
func NewGetRepositoryForbidden() *GetRepositoryForbidden {
	return &GetRepositoryForbidden{}
}

/*GetRepositoryForbidden handles this case with default header values.

Forbidden
*/
type GetRepositoryForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRepositoryForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryForbidden  %+v", 403, o.Payload)
}

func (o *GetRepositoryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryNotFound creates a GetRepositoryNotFound with default headers values
func NewGetRepositoryNotFound() *GetRepositoryNotFound {
	return &GetRepositoryNotFound{}
}

/*GetRepositoryNotFound handles this case with default header values.

Not found
*/
type GetRepositoryNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRepositoryNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoryNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryInternalServerError creates a GetRepositoryInternalServerError with default headers values
func NewGetRepositoryInternalServerError() *GetRepositoryInternalServerError {
	return &GetRepositoryInternalServerError{}
}

/*GetRepositoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetRepositoryInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRepositoryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
