// Code generated by go-swagger; DO NOT EDIT.

package blog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/bbcyyb/bunkerhill/models"
)

// NewInsertBlogParams creates a new InsertBlogParams object
// no default values defined in spec.
func NewInsertBlogParams() InsertBlogParams {

	return InsertBlogParams{}
}

// InsertBlogParams contains all the bound params for the insert blog operation
// typically these are obtained from a http.Request
//
// swagger:parameters InsertBlog
type InsertBlogParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Content of blog to be saved
	  Required: true
	  In: body
	*/
	Blog *models.Blog
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewInsertBlogParams() beforehand.
func (o *InsertBlogParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Blog
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("blog", "body"))
			} else {
				res = append(res, errors.NewParseError("blog", "body", "", err))
			}
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Blog = &body
			}
		}
	} else {
		res = append(res, errors.Required("blog", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
