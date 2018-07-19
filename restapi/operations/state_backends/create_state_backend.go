// Code generated by go-swagger; DO NOT EDIT.

package state_backends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/getlunaform/lunaform/models"
)

// CreateStateBackendHandlerFunc turns a function with the right signature into a create state backend handler
type CreateStateBackendHandlerFunc func(CreateStateBackendParams, *models.ResourceAuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateStateBackendHandlerFunc) Handle(params CreateStateBackendParams, principal *models.ResourceAuthUser) middleware.Responder {
	return fn(params, principal)
}

// CreateStateBackendHandler interface for that can handle valid create state backend params
type CreateStateBackendHandler interface {
	Handle(CreateStateBackendParams, *models.ResourceAuthUser) middleware.Responder
}

// NewCreateStateBackend creates a new http.Handler for the create state backend operation
func NewCreateStateBackend(ctx *middleware.Context, handler CreateStateBackendHandler) *CreateStateBackend {
	return &CreateStateBackend{Context: ctx, Handler: handler}
}

/*CreateStateBackend swagger:route POST /tf/state-backends state-backends createStateBackend

Define a Terraform state backend

*/
type CreateStateBackend struct {
	Context *middleware.Context
	Handler CreateStateBackendHandler
}

func (o *CreateStateBackend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateStateBackendParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.ResourceAuthUser
	if uprinc != nil {
		principal = uprinc.(*models.ResourceAuthUser) // this is really a models.ResourceAuthUser, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}