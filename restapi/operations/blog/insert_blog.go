// Code generated by go-swagger; DO NOT EDIT.

package blog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// InsertBlogHandlerFunc turns a function with the right signature into a insert blog handler
type InsertBlogHandlerFunc func(InsertBlogParams) middleware.Responder

// Handle executing the request and returning a response
func (fn InsertBlogHandlerFunc) Handle(params InsertBlogParams) middleware.Responder {
	return fn(params)
}

// InsertBlogHandler interface for that can handle valid insert blog params
type InsertBlogHandler interface {
	Handle(InsertBlogParams) middleware.Responder
}

// NewInsertBlog creates a new http.Handler for the insert blog operation
func NewInsertBlog(ctx *middleware.Context, handler InsertBlogHandler) *InsertBlog {
	return &InsertBlog{Context: ctx, Handler: handler}
}

/*InsertBlog swagger:route POST /blogs blog insertBlog

InsertBlog insert blog API

*/
type InsertBlog struct {
	Context *middleware.Context
	Handler InsertBlogHandler
}

func (o *InsertBlog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewInsertBlogParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
