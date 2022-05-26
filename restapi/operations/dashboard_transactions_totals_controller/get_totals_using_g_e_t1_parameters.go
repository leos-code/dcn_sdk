// Code generated by go-swagger; DO NOT EDIT.

package dashboard_transactions_totals_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetTotalsUsingGET1Params creates a new GetTotalsUsingGET1Params object
//
// There are no default values defined in the spec.
func NewGetTotalsUsingGET1Params() GetTotalsUsingGET1Params {

	return GetTotalsUsingGET1Params{}
}

// GetTotalsUsingGET1Params contains all the bound params for the get totals using g e t 1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTotalsUsingGET_1
type GetTotalsUsingGET1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTotalsUsingGET1Params() beforehand.
func (o *GetTotalsUsingGET1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}