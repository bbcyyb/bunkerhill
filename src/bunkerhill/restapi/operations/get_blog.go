// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBlogHandlerFunc turns a function with the right signature into a get blog handler
type GetBlogHandlerFunc func(GetBlogParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBlogHandlerFunc) Handle(params GetBlogParams) middleware.Responder {
	return fn(params)
}

// GetBlogHandler interface for that can handle valid get blog params
type GetBlogHandler interface {
	Handle(GetBlogParams) middleware.Responder
}

// NewGetBlog creates a new http.Handler for the get blog operation
func NewGetBlog(ctx *middleware.Context, handler GetBlogHandler) *GetBlog {
	return &GetBlog{Context: ctx, Handler: handler}
}

/*GetBlog swagger:route GET /blog getBlog

GetBlog get blog API

*/
type GetBlog struct {
	Context *middleware.Context
	Handler GetBlogHandler
}

func (o *GetBlog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBlogParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
