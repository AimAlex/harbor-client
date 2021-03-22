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

// NewPostRetentionsParams creates a new PostRetentionsParams object
// with the default values initialized.
func NewPostRetentionsParams() *PostRetentionsParams {
	var ()
	return &PostRetentionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRetentionsParamsWithTimeout creates a new PostRetentionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRetentionsParamsWithTimeout(timeout time.Duration) *PostRetentionsParams {
	var ()
	return &PostRetentionsParams{

		timeout: timeout,
	}
}

// NewPostRetentionsParamsWithContext creates a new PostRetentionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRetentionsParamsWithContext(ctx context.Context) *PostRetentionsParams {
	var ()
	return &PostRetentionsParams{

		Context: ctx,
	}
}

// NewPostRetentionsParamsWithHTTPClient creates a new PostRetentionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRetentionsParamsWithHTTPClient(client *http.Client) *PostRetentionsParams {
	var ()
	return &PostRetentionsParams{
		HTTPClient: client,
	}
}

/*PostRetentionsParams contains all the parameters to send to the API endpoint
for the post retentions operation typically these are written to a http.Request
*/
type PostRetentionsParams struct {

	/*Policy
	  Create Retention Policy successfully.

	*/
	Policy *models.RetentionPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post retentions params
func (o *PostRetentionsParams) WithTimeout(timeout time.Duration) *PostRetentionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post retentions params
func (o *PostRetentionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post retentions params
func (o *PostRetentionsParams) WithContext(ctx context.Context) *PostRetentionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post retentions params
func (o *PostRetentionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post retentions params
func (o *PostRetentionsParams) WithHTTPClient(client *http.Client) *PostRetentionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post retentions params
func (o *PostRetentionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the post retentions params
func (o *PostRetentionsParams) WithPolicy(policy *models.RetentionPolicy) *PostRetentionsParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the post retentions params
func (o *PostRetentionsParams) SetPolicy(policy *models.RetentionPolicy) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *PostRetentionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}