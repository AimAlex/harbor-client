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
	"github.com/go-openapi/swag"

	"gitlab.4pd.io/liwenhao/pineapple/pineapple/apigen/harborcli/models"
)

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDParams creates a new PutProjectsProjectIDWebhookPoliciesPolicyIDParams object
// with the default values initialized.
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDParams() *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	var ()
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDParamsWithTimeout creates a new PutProjectsProjectIDWebhookPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDParamsWithTimeout(timeout time.Duration) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	var ()
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDParams{

		timeout: timeout,
	}
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDParamsWithContext creates a new PutProjectsProjectIDWebhookPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDParamsWithContext(ctx context.Context) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	var ()
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDParams{

		Context: ctx,
	}
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDParamsWithHTTPClient creates a new PutProjectsProjectIDWebhookPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDParamsWithHTTPClient(client *http.Client) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	var ()
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDParams{
		HTTPClient: client,
	}
}

/*PutProjectsProjectIDWebhookPoliciesPolicyIDParams contains all the parameters to send to the API endpoint
for the put projects project ID webhook policies policy ID operation typically these are written to a http.Request
*/
type PutProjectsProjectIDWebhookPoliciesPolicyIDParams struct {

	/*Policy
	  All properties needed except "id", "project_id", "creation_time", "update_time".

	*/
	Policy *models.WebhookPolicy
	/*PolicyID
	  The id of webhook policy.

	*/
	PolicyID int64
	/*ProjectID
	  Relevant project ID.

	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) WithTimeout(timeout time.Duration) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) WithContext(ctx context.Context) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) WithHTTPClient(client *http.Client) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) WithPolicy(policy *models.WebhookPolicy) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) SetPolicy(policy *models.WebhookPolicy) {
	o.Policy = policy
}

// WithPolicyID adds the policyID to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) WithPolicyID(policyID int64) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WithProjectID adds the projectID to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) WithProjectID(projectID int64) *PutProjectsProjectIDWebhookPoliciesPolicyIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put projects project ID webhook policies policy ID params
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	// path param policy_id
	if err := r.SetPathParam("policy_id", swag.FormatInt64(o.PolicyID)); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
