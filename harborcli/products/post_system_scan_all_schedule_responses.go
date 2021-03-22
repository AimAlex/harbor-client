// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostSystemScanAllScheduleReader is a Reader for the PostSystemScanAllSchedule structure.
type PostSystemScanAllScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSystemScanAllScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSystemScanAllScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSystemScanAllScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSystemScanAllScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSystemScanAllScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSystemScanAllScheduleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSystemScanAllScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostSystemScanAllScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSystemScanAllScheduleOK creates a PostSystemScanAllScheduleOK with default headers values
func NewPostSystemScanAllScheduleOK() *PostSystemScanAllScheduleOK {
	return &PostSystemScanAllScheduleOK{}
}

/*PostSystemScanAllScheduleOK handles this case with default header values.

Updated scan_all's schedule successfully.
*/
type PostSystemScanAllScheduleOK struct {
}

func (o *PostSystemScanAllScheduleOK) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] postSystemScanAllScheduleOK ", 200)
}

func (o *PostSystemScanAllScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemScanAllScheduleBadRequest creates a PostSystemScanAllScheduleBadRequest with default headers values
func NewPostSystemScanAllScheduleBadRequest() *PostSystemScanAllScheduleBadRequest {
	return &PostSystemScanAllScheduleBadRequest{}
}

/*PostSystemScanAllScheduleBadRequest handles this case with default header values.

Invalid schedule type.
*/
type PostSystemScanAllScheduleBadRequest struct {
}

func (o *PostSystemScanAllScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] postSystemScanAllScheduleBadRequest ", 400)
}

func (o *PostSystemScanAllScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemScanAllScheduleUnauthorized creates a PostSystemScanAllScheduleUnauthorized with default headers values
func NewPostSystemScanAllScheduleUnauthorized() *PostSystemScanAllScheduleUnauthorized {
	return &PostSystemScanAllScheduleUnauthorized{}
}

/*PostSystemScanAllScheduleUnauthorized handles this case with default header values.

User need to log in first.
*/
type PostSystemScanAllScheduleUnauthorized struct {
}

func (o *PostSystemScanAllScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] postSystemScanAllScheduleUnauthorized ", 401)
}

func (o *PostSystemScanAllScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemScanAllScheduleForbidden creates a PostSystemScanAllScheduleForbidden with default headers values
func NewPostSystemScanAllScheduleForbidden() *PostSystemScanAllScheduleForbidden {
	return &PostSystemScanAllScheduleForbidden{}
}

/*PostSystemScanAllScheduleForbidden handles this case with default header values.

User does not have permission of admin role.
*/
type PostSystemScanAllScheduleForbidden struct {
}

func (o *PostSystemScanAllScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] postSystemScanAllScheduleForbidden ", 403)
}

func (o *PostSystemScanAllScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemScanAllScheduleConflict creates a PostSystemScanAllScheduleConflict with default headers values
func NewPostSystemScanAllScheduleConflict() *PostSystemScanAllScheduleConflict {
	return &PostSystemScanAllScheduleConflict{}
}

/*PostSystemScanAllScheduleConflict handles this case with default header values.

There is a "scanall" job in progress, so the request cannot be served.
*/
type PostSystemScanAllScheduleConflict struct {
}

func (o *PostSystemScanAllScheduleConflict) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] postSystemScanAllScheduleConflict ", 409)
}

func (o *PostSystemScanAllScheduleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemScanAllScheduleInternalServerError creates a PostSystemScanAllScheduleInternalServerError with default headers values
func NewPostSystemScanAllScheduleInternalServerError() *PostSystemScanAllScheduleInternalServerError {
	return &PostSystemScanAllScheduleInternalServerError{}
}

/*PostSystemScanAllScheduleInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PostSystemScanAllScheduleInternalServerError struct {
}

func (o *PostSystemScanAllScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] postSystemScanAllScheduleInternalServerError ", 500)
}

func (o *PostSystemScanAllScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemScanAllScheduleServiceUnavailable creates a PostSystemScanAllScheduleServiceUnavailable with default headers values
func NewPostSystemScanAllScheduleServiceUnavailable() *PostSystemScanAllScheduleServiceUnavailable {
	return &PostSystemScanAllScheduleServiceUnavailable{}
}

/*PostSystemScanAllScheduleServiceUnavailable handles this case with default header values.

Harbor is not deployed with scanners.
*/
type PostSystemScanAllScheduleServiceUnavailable struct {
}

func (o *PostSystemScanAllScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] postSystemScanAllScheduleServiceUnavailable ", 503)
}

func (o *PostSystemScanAllScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
