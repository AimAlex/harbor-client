// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutSystemScanAllScheduleReader is a Reader for the PutSystemScanAllSchedule structure.
type PutSystemScanAllScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSystemScanAllScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSystemScanAllScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutSystemScanAllScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutSystemScanAllScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSystemScanAllScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutSystemScanAllScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutSystemScanAllScheduleOK creates a PutSystemScanAllScheduleOK with default headers values
func NewPutSystemScanAllScheduleOK() *PutSystemScanAllScheduleOK {
	return &PutSystemScanAllScheduleOK{}
}

/*PutSystemScanAllScheduleOK handles this case with default header values.

Updated scan_all's schedule successfully.
*/
type PutSystemScanAllScheduleOK struct {
}

func (o *PutSystemScanAllScheduleOK) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] putSystemScanAllScheduleOK ", 200)
}

func (o *PutSystemScanAllScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemScanAllScheduleBadRequest creates a PutSystemScanAllScheduleBadRequest with default headers values
func NewPutSystemScanAllScheduleBadRequest() *PutSystemScanAllScheduleBadRequest {
	return &PutSystemScanAllScheduleBadRequest{}
}

/*PutSystemScanAllScheduleBadRequest handles this case with default header values.

Invalid schedule type.
*/
type PutSystemScanAllScheduleBadRequest struct {
}

func (o *PutSystemScanAllScheduleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] putSystemScanAllScheduleBadRequest ", 400)
}

func (o *PutSystemScanAllScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemScanAllScheduleUnauthorized creates a PutSystemScanAllScheduleUnauthorized with default headers values
func NewPutSystemScanAllScheduleUnauthorized() *PutSystemScanAllScheduleUnauthorized {
	return &PutSystemScanAllScheduleUnauthorized{}
}

/*PutSystemScanAllScheduleUnauthorized handles this case with default header values.

User need to log in first.
*/
type PutSystemScanAllScheduleUnauthorized struct {
}

func (o *PutSystemScanAllScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] putSystemScanAllScheduleUnauthorized ", 401)
}

func (o *PutSystemScanAllScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemScanAllScheduleForbidden creates a PutSystemScanAllScheduleForbidden with default headers values
func NewPutSystemScanAllScheduleForbidden() *PutSystemScanAllScheduleForbidden {
	return &PutSystemScanAllScheduleForbidden{}
}

/*PutSystemScanAllScheduleForbidden handles this case with default header values.

User does not have permission of admin role.
*/
type PutSystemScanAllScheduleForbidden struct {
}

func (o *PutSystemScanAllScheduleForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] putSystemScanAllScheduleForbidden ", 403)
}

func (o *PutSystemScanAllScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemScanAllScheduleInternalServerError creates a PutSystemScanAllScheduleInternalServerError with default headers values
func NewPutSystemScanAllScheduleInternalServerError() *PutSystemScanAllScheduleInternalServerError {
	return &PutSystemScanAllScheduleInternalServerError{}
}

/*PutSystemScanAllScheduleInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PutSystemScanAllScheduleInternalServerError struct {
}

func (o *PutSystemScanAllScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] putSystemScanAllScheduleInternalServerError ", 500)
}

func (o *PutSystemScanAllScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
