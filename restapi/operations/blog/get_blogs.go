// Code generated by go-swagger; DO NOT EDIT.

package blog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBlogsHandlerFunc turns a function with the right signature into a get blogs handler
type GetBlogsHandlerFunc func(GetBlogsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBlogsHandlerFunc) Handle(params GetBlogsParams) middleware.Responder {
	return fn(params)
}

// GetBlogsHandler interface for that can handle valid get blogs params
type GetBlogsHandler interface {
	Handle(GetBlogsParams) middleware.Responder
}

// NewGetBlogs creates a new http.Handler for the get blogs operation
func NewGetBlogs(ctx *middleware.Context, handler GetBlogsHandler) *GetBlogs {
	return &GetBlogs{Context: ctx, Handler: handler}
}

/*GetBlogs swagger:route GET /blogs blog getBlogs

GetBlogs get blogs API

*/
type GetBlogs struct {
	Context *middleware.Context
	Handler GetBlogsHandler
}

func (o *GetBlogs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBlogsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
