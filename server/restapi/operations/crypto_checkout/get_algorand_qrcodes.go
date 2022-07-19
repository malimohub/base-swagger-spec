// Code generated by go-swagger; DO NOT EDIT.

package crypto_checkout

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAlgorandQrcodesHandlerFunc turns a function with the right signature into a get algorand qrcodes handler
type GetAlgorandQrcodesHandlerFunc func(GetAlgorandQrcodesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAlgorandQrcodesHandlerFunc) Handle(params GetAlgorandQrcodesParams) middleware.Responder {
	return fn(params)
}

// GetAlgorandQrcodesHandler interface for that can handle valid get algorand qrcodes params
type GetAlgorandQrcodesHandler interface {
	Handle(GetAlgorandQrcodesParams) middleware.Responder
}

// NewGetAlgorandQrcodes creates a new http.Handler for the get algorand qrcodes operation
func NewGetAlgorandQrcodes(ctx *middleware.Context, handler GetAlgorandQrcodesHandler) *GetAlgorandQrcodes {
	return &GetAlgorandQrcodes{Context: ctx, Handler: handler}
}

/* GetAlgorandQrcodes swagger:route GET /algorand/qrcodes crypto-checkout getAlgorandQrcodes

GET Algorand redeem rewards QR code

GET Algorand redeem rewards QR code


*/
type GetAlgorandQrcodes struct {
	Context *middleware.Context
	Handler GetAlgorandQrcodesHandler
}

func (o *GetAlgorandQrcodes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAlgorandQrcodesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}