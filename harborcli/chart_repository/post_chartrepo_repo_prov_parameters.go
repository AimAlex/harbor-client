// Code generated by go-swagger; DO NOT EDIT.

package chart_repository

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
)

// NewPostChartrepoRepoProvParams creates a new PostChartrepoRepoProvParams object
// with the default values initialized.
func NewPostChartrepoRepoProvParams() *PostChartrepoRepoProvParams {
	var ()
	return &PostChartrepoRepoProvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostChartrepoRepoProvParamsWithTimeout creates a new PostChartrepoRepoProvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostChartrepoRepoProvParamsWithTimeout(timeout time.Duration) *PostChartrepoRepoProvParams {
	var ()
	return &PostChartrepoRepoProvParams{

		timeout: timeout,
	}
}

// NewPostChartrepoRepoProvParamsWithContext creates a new PostChartrepoRepoProvParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostChartrepoRepoProvParamsWithContext(ctx context.Context) *PostChartrepoRepoProvParams {
	var ()
	return &PostChartrepoRepoProvParams{

		Context: ctx,
	}
}

// NewPostChartrepoRepoProvParamsWithHTTPClient creates a new PostChartrepoRepoProvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostChartrepoRepoProvParamsWithHTTPClient(client *http.Client) *PostChartrepoRepoProvParams {
	var ()
	return &PostChartrepoRepoProvParams{
		HTTPClient: client,
	}
}

/*PostChartrepoRepoProvParams contains all the parameters to send to the API endpoint
for the post chartrepo repo prov operation typically these are written to a http.Request
*/
type PostChartrepoRepoProvParams struct {

	/*Prov
	  The provance file

	*/
	Prov runtime.NamedReadCloser
	/*Repo
	  The project name

	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) WithTimeout(timeout time.Duration) *PostChartrepoRepoProvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) WithContext(ctx context.Context) *PostChartrepoRepoProvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) WithHTTPClient(client *http.Client) *PostChartrepoRepoProvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProv adds the prov to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) WithProv(prov runtime.NamedReadCloser) *PostChartrepoRepoProvParams {
	o.SetProv(prov)
	return o
}

// SetProv adds the prov to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) SetProv(prov runtime.NamedReadCloser) {
	o.Prov = prov
}

// WithRepo adds the repo to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) WithRepo(repo string) *PostChartrepoRepoProvParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the post chartrepo repo prov params
func (o *PostChartrepoRepoProvParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *PostChartrepoRepoProvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param prov
	if err := r.SetFileParam("prov", o.Prov); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
