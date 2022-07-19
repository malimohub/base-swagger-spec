// Code generated by go-swagger; DO NOT EDIT.

package crypto_checkout

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostCheckoutHandlerFunc turns a function with the right signature into a post checkout handler
type PostCheckoutHandlerFunc func(PostCheckoutParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCheckoutHandlerFunc) Handle(params PostCheckoutParams) middleware.Responder {
	return fn(params)
}

// PostCheckoutHandler interface for that can handle valid post checkout params
type PostCheckoutHandler interface {
	Handle(PostCheckoutParams) middleware.Responder
}

// NewPostCheckout creates a new http.Handler for the post checkout operation
func NewPostCheckout(ctx *middleware.Context, handler PostCheckoutHandler) *PostCheckout {
	return &PostCheckout{Context: ctx, Handler: handler}
}

/* PostCheckout swagger:route POST /checkout crypto-checkout postCheckout

Post start checkout

Post start checkout


*/
type PostCheckout struct {
	Context *middleware.Context
	Handler PostCheckoutHandler
}

func (o *PostCheckout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostCheckoutParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
