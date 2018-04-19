// Code generated by go-swagger; DO NOT EDIT.

package blog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/bbcyyb/bunkerhill/models"
)

// GetBlogsOKCode is the HTTP code returned for type GetBlogsOK
const GetBlogsOKCode int = 200

/*GetBlogsOK Successful response

swagger:response getBlogsOK
*/
type GetBlogsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Blog `json:"body,omitempty"`
}

// NewGetBlogsOK creates GetBlogsOK with default headers values
func NewGetBlogsOK() *GetBlogsOK {

	return &GetBlogsOK{}
}

// WithPayload adds the payload to the get blogs o k response
func (o *GetBlogsOK) WithPayload(payload []*models.Blog) *GetBlogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blogs o k response
func (o *GetBlogsOK) SetPayload(payload []*models.Blog) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Blog, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetBlogsBadRequestCode is the HTTP code returned for type GetBlogsBadRequest
const GetBlogsBadRequestCode int = 400

/*GetBlogsBadRequest Bad Request

swagger:response getBlogsBadRequest
*/
type GetBlogsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GenericError `json:"body,omitempty"`
}

// NewGetBlogsBadRequest creates GetBlogsBadRequest with default headers values
func NewGetBlogsBadRequest() *GetBlogsBadRequest {

	return &GetBlogsBadRequest{}
}

// WithPayload adds the payload to the get blogs bad request response
func (o *GetBlogsBadRequest) WithPayload(payload *models.GenericError) *GetBlogsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blogs bad request response
func (o *GetBlogsBadRequest) SetPayload(payload *models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlogsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBlogsInternalServerErrorCode is the HTTP code returned for type GetBlogsInternalServerError
const GetBlogsInternalServerErrorCode int = 500

/*GetBlogsInternalServerError Internal Server Error

swagger:response getBlogsInternalServerError
*/
type GetBlogsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GenericError `json:"body,omitempty"`
}

// NewGetBlogsInternalServerError creates GetBlogsInternalServerError with default headers values
func NewGetBlogsInternalServerError() *GetBlogsInternalServerError {

	return &GetBlogsInternalServerError{}
}

// WithPayload adds the payload to the get blogs internal server error response
func (o *GetBlogsInternalServerError) WithPayload(payload *models.GenericError) *GetBlogsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blogs internal server error response
func (o *GetBlogsInternalServerError) SetPayload(payload *models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlogsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
