// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"gitlab.4pd.io/liwenhao/pineapple/pineapple/apigen/harborcli/models"
)

// ListPoliciesReader is a Reader for the ListPolicies structure.
type ListPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPoliciesOK creates a ListPoliciesOK with default headers values
func NewListPoliciesOK() *ListPoliciesOK {
	return &ListPoliciesOK{}
}

/*ListPoliciesOK handles this case with default header values.

List preheat policies success
*/
type ListPoliciesOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of policies
	 */
	XTotalCount int64

	Payload []*models.PreheatPolicy
}

func (o *ListPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListPoliciesOK) GetPayload() []*models.PreheatPolicy {
	return o.Payload
}

func (o *ListPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPoliciesBadRequest creates a ListPoliciesBadRequest with default headers values
func NewListPoliciesBadRequest() *ListPoliciesBadRequest {
	return &ListPoliciesBadRequest{}
}

/*ListPoliciesBadRequest handles this case with default header values.

Bad request
*/
type ListPoliciesBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *ListPoliciesBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPoliciesUnauthorized creates a ListPoliciesUnauthorized with default headers values
func NewListPoliciesUnauthorized() *ListPoliciesUnauthorized {
	return &ListPoliciesUnauthorized{}
}

/*ListPoliciesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListPoliciesUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListPoliciesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPoliciesForbidden creates a ListPoliciesForbidden with default headers values
func NewListPoliciesForbidden() *ListPoliciesForbidden {
	return &ListPoliciesForbidden{}
}

/*ListPoliciesForbidden handles this case with default header values.

Forbidden
*/
type ListPoliciesForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *ListPoliciesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPoliciesInternalServerError creates a ListPoliciesInternalServerError with default headers values
func NewListPoliciesInternalServerError() *ListPoliciesInternalServerError {
	return &ListPoliciesInternalServerError{}
}

/*ListPoliciesInternalServerError handles this case with default header values.

Internal server error
*/
type ListPoliciesInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPoliciesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
