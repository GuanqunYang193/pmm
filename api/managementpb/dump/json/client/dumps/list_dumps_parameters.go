// Code generated by go-swagger; DO NOT EDIT.

package dumps

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

// NewListDumpsParams creates a new ListDumpsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDumpsParams() *ListDumpsParams {
	return &ListDumpsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDumpsParamsWithTimeout creates a new ListDumpsParams object
// with the ability to set a timeout on a request.
func NewListDumpsParamsWithTimeout(timeout time.Duration) *ListDumpsParams {
	return &ListDumpsParams{
		timeout: timeout,
	}
}

// NewListDumpsParamsWithContext creates a new ListDumpsParams object
// with the ability to set a context for a request.
func NewListDumpsParamsWithContext(ctx context.Context) *ListDumpsParams {
	return &ListDumpsParams{
		Context: ctx,
	}
}

// NewListDumpsParamsWithHTTPClient creates a new ListDumpsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDumpsParamsWithHTTPClient(client *http.Client) *ListDumpsParams {
	return &ListDumpsParams{
		HTTPClient: client,
	}
}

/*
ListDumpsParams contains all the parameters to send to the API endpoint

	for the list dumps operation.

	Typically these are written to a http.Request.
*/
type ListDumpsParams struct {
	// Body.
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list dumps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDumpsParams) WithDefaults() *ListDumpsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list dumps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDumpsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list dumps params
func (o *ListDumpsParams) WithTimeout(timeout time.Duration) *ListDumpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list dumps params
func (o *ListDumpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list dumps params
func (o *ListDumpsParams) WithContext(ctx context.Context) *ListDumpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list dumps params
func (o *ListDumpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list dumps params
func (o *ListDumpsParams) WithHTTPClient(client *http.Client) *ListDumpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list dumps params
func (o *ListDumpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list dumps params
func (o *ListDumpsParams) WithBody(body interface{}) *ListDumpsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list dumps params
func (o *ListDumpsParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListDumpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
