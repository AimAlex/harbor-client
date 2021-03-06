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

// NewGetChartrepoRepoChartsParams creates a new GetChartrepoRepoChartsParams object
// with the default values initialized.
func NewGetChartrepoRepoChartsParams() *GetChartrepoRepoChartsParams {
	var ()
	return &GetChartrepoRepoChartsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChartrepoRepoChartsParamsWithTimeout creates a new GetChartrepoRepoChartsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChartrepoRepoChartsParamsWithTimeout(timeout time.Duration) *GetChartrepoRepoChartsParams {
	var ()
	return &GetChartrepoRepoChartsParams{

		timeout: timeout,
	}
}

// NewGetChartrepoRepoChartsParamsWithContext creates a new GetChartrepoRepoChartsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChartrepoRepoChartsParamsWithContext(ctx context.Context) *GetChartrepoRepoChartsParams {
	var ()
	return &GetChartrepoRepoChartsParams{

		Context: ctx,
	}
}

// NewGetChartrepoRepoChartsParamsWithHTTPClient creates a new GetChartrepoRepoChartsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChartrepoRepoChartsParamsWithHTTPClient(client *http.Client) *GetChartrepoRepoChartsParams {
	var ()
	return &GetChartrepoRepoChartsParams{
		HTTPClient: client,
	}
}

/*GetChartrepoRepoChartsParams contains all the parameters to send to the API endpoint
for the get chartrepo repo charts operation typically these are written to a http.Request
*/
type GetChartrepoRepoChartsParams struct {

	/*Repo
	  The project name

	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get chartrepo repo charts params
func (o *GetChartrepoRepoChartsParams) WithTimeout(timeout time.Duration) *GetChartrepoRepoChartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chartrepo repo charts params
func (o *GetChartrepoRepoChartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chartrepo repo charts params
func (o *GetChartrepoRepoChartsParams) WithContext(ctx context.Context) *GetChartrepoRepoChartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chartrepo repo charts params
func (o *GetChartrepoRepoChartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chartrepo repo charts params
func (o *GetChartrepoRepoChartsParams) WithHTTPClient(client *http.Client) *GetChartrepoRepoChartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chartrepo repo charts params
func (o *GetChartrepoRepoChartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepo adds the repo to the get chartrepo repo charts params
func (o *GetChartrepoRepoChartsParams) WithRepo(repo string) *GetChartrepoRepoChartsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the get chartrepo repo charts params
func (o *GetChartrepoRepoChartsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *GetChartrepoRepoChartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
