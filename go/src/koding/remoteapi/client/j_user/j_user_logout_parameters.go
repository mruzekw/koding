package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewJUserLogoutParams creates a new JUserLogoutParams object
// with the default values initialized.
func NewJUserLogoutParams() *JUserLogoutParams {

	return &JUserLogoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJUserLogoutParamsWithTimeout creates a new JUserLogoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJUserLogoutParamsWithTimeout(timeout time.Duration) *JUserLogoutParams {

	return &JUserLogoutParams{

		timeout: timeout,
	}
}

// NewJUserLogoutParamsWithContext creates a new JUserLogoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewJUserLogoutParamsWithContext(ctx context.Context) *JUserLogoutParams {

	return &JUserLogoutParams{

		Context: ctx,
	}
}

/*JUserLogoutParams contains all the parameters to send to the API endpoint
for the j user logout operation typically these are written to a http.Request
*/
type JUserLogoutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j user logout params
func (o *JUserLogoutParams) WithTimeout(timeout time.Duration) *JUserLogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j user logout params
func (o *JUserLogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j user logout params
func (o *JUserLogoutParams) WithContext(ctx context.Context) *JUserLogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j user logout params
func (o *JUserLogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *JUserLogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
