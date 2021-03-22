// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AimAlex/harbor-client/models"
)

// StopExecutionReader is a Reader for the StopExecution structure.
type StopExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStopExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopExecutionOK creates a StopExecutionOK with default headers values
func NewStopExecutionOK() *StopExecutionOK {
	return &StopExecutionOK{}
}

/*StopExecutionOK handles this case with default header values.

Success
*/
type StopExecutionOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *StopExecutionOK) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionOK ", 200)
}

func (o *StopExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewStopExecutionBadRequest creates a StopExecutionBadRequest with default headers values
func NewStopExecutionBadRequest() *StopExecutionBadRequest {
	return &StopExecutionBadRequest{}
}

/*StopExecutionBadRequest handles this case with default header values.

Bad request
*/
type StopExecutionBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *StopExecutionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionBadRequest  %+v", 400, o.Payload)
}

func (o *StopExecutionBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopExecutionUnauthorized creates a StopExecutionUnauthorized with default headers values
func NewStopExecutionUnauthorized() *StopExecutionUnauthorized {
	return &StopExecutionUnauthorized{}
}

/*StopExecutionUnauthorized handles this case with default header values.

Unauthorized
*/
type StopExecutionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *StopExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *StopExecutionUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopExecutionForbidden creates a StopExecutionForbidden with default headers values
func NewStopExecutionForbidden() *StopExecutionForbidden {
	return &StopExecutionForbidden{}
}

/*StopExecutionForbidden handles this case with default header values.

Forbidden
*/
type StopExecutionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *StopExecutionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionForbidden  %+v", 403, o.Payload)
}

func (o *StopExecutionForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopExecutionNotFound creates a StopExecutionNotFound with default headers values
func NewStopExecutionNotFound() *StopExecutionNotFound {
	return &StopExecutionNotFound{}
}

/*StopExecutionNotFound handles this case with default header values.

Not found
*/
type StopExecutionNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *StopExecutionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionNotFound  %+v", 404, o.Payload)
}

func (o *StopExecutionNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopExecutionInternalServerError creates a StopExecutionInternalServerError with default headers values
func NewStopExecutionInternalServerError() *StopExecutionInternalServerError {
	return &StopExecutionInternalServerError{}
}

/*StopExecutionInternalServerError handles this case with default header values.

Internal server error
*/
type StopExecutionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *StopExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *StopExecutionInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
