// Code generated by go-swagger; DO NOT EDIT.

package account_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetAccountUsingGETParams creates a new GetAccountUsingGETParams object
//
// There are no default values defined in the spec.
func NewGetAccountUsingGETParams() GetAccountUsingGETParams {

	return GetAccountUsingGETParams{}
}

// GetAccountUsingGETParams contains all the bound params for the get account using g e t operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAccountUsingGET
type GetAccountUsingGETParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The address hash identifying the account.
	  In: path
	*/
	AddressHash *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAccountUsingGETParams() beforehand.
func (o *GetAccountUsingGETParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAddressHash, rhkAddressHash, _ := route.Params.GetOK("addressHash")
	if err := o.bindAddressHash(rAddressHash, rhkAddressHash, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAddressHash binds and validates parameter AddressHash from path.
func (o *GetAccountUsingGETParams) bindAddressHash(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// Parameter is provided by construction from the route
	o.AddressHash = &raw

	return nil
}