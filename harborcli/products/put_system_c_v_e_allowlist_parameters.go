// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"gitlab.4pd.io/liwenhao/pineapple/pineapple/apigen/harborcli/models"
)

// NewPutSystemCVEAllowlistParams creates a new PutSystemCVEAllowlistParams object
// with the default values initialized.
func NewPutSystemCVEAllowlistParams() *PutSystemCVEAllowlistParams {
	var ()
	return &PutSystemCVEAllowlistParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSystemCVEAllowlistParamsWithTimeout creates a new PutSystemCVEAllowlistParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSystemCVEAllowlistParamsWithTimeout(timeout time.Duration) *PutSystemCVEAllowlistParams {
	var ()
	return &PutSystemCVEAllowlistParams{

		timeout: timeout,
	}
}

// NewPutSystemCVEAllowlistParamsWithContext creates a new PutSystemCVEAllowlistParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSystemCVEAllowlistParamsWithContext(ctx context.Context) *PutSystemCVEAllowlistParams {
	var ()
	return &PutSystemCVEAllowlistParams{

		Context: ctx,
	}
}

// NewPutSystemCVEAllowlistParamsWithHTTPClient creates a new PutSystemCVEAllowlistParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSystemCVEAllowlistParamsWithHTTPClient(client *http.Client) *PutSystemCVEAllowlistParams {
	var ()
	return &PutSystemCVEAllowlistParams{
		HTTPClient: client,
	}
}

/*PutSystemCVEAllowlistParams contains all the parameters to send to the API endpoint
for the put system c v e allowlist operation typically these are written to a http.Request
*/
type PutSystemCVEAllowlistParams struct {

	/*Allowlist
	  The allowlist with new content

	*/
	Allowlist *models.CVEAllowlist

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put system c v e allowlist params
func (o *PutSystemCVEAllowlistParams) WithTimeout(timeout time.Duration) *PutSystemCVEAllowlistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put system c v e allowlist params
func (o *PutSystemCVEAllowlistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put system c v e allowlist params
func (o *PutSystemCVEAllowlistParams) WithContext(ctx context.Context) *PutSystemCVEAllowlistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put system c v e allowlist params
func (o *PutSystemCVEAllowlistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put system c v e allowlist params
func (o *PutSystemCVEAllowlistParams) WithHTTPClient(client *http.Client) *PutSystemCVEAllowlistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put system c v e allowlist params
func (o *PutSystemCVEAllowlistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowlist adds the allowlist to the put system c v e allowlist params
func (o *PutSystemCVEAllowlistParams) WithAllowlist(allowlist *models.CVEAllowlist) *PutSystemCVEAllowlistParams {
	o.SetAllowlist(allowlist)
	return o
}

// SetAllowlist adds the allowlist to the put system c v e allowlist params
func (o *PutSystemCVEAllowlistParams) SetAllowlist(allowlist *models.CVEAllowlist) {
	o.Allowlist = allowlist
}

// WriteToRequest writes these params to a swagger request
func (o *PutSystemCVEAllowlistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Allowlist != nil {
		if err := r.SetBodyParam(o.Allowlist); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
