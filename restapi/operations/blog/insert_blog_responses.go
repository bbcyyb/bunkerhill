// Code generated by go-swagger; DO NOT EDIT.

package blog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/bbcyyb/bunkerhill/models"
)

// InsertBlogCreatedCode is the HTTP code returned for type InsertBlogCreated
const InsertBlogCreatedCode int = 201

/*InsertBlogCreated The request has been fulfilled, resulting in the creation of a new resource.

swagger:response insertBlogCreated
*/
type InsertBlogCreated struct {

	/*
	  In: Body
	*/
	Payload *models.InsertBlogCreatedBody `json:"body,omitempty"`
}

// NewInsertBlogCreated creates InsertBlogCreated with default headers values
func NewInsertBlogCreated() *InsertBlogCreated {

	return &InsertBlogCreated{}
}

// WithPayload adds the payload to the insert blog created response
func (o *InsertBlogCreated) WithPayload(payload *models.InsertBlogCreatedBody) *InsertBlogCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the insert blog created response
func (o *InsertBlogCreated) SetPayload(payload *models.InsertBlogCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InsertBlogCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// InsertBlogBadRequestCode is the HTTP code returned for type InsertBlogBadRequest
const InsertBlogBadRequestCode int = 400

/*InsertBlogBadRequest Bad Request

swagger:response insertBlogBadRequest
*/
type InsertBlogBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GenericError `json:"body,omitempty"`
}

// NewInsertBlogBadRequest creates InsertBlogBadRequest with default headers values
func NewInsertBlogBadRequest() *InsertBlogBadRequest {

	return &InsertBlogBadRequest{}
}

// WithPayload adds the payload to the insert blog bad request response
func (o *InsertBlogBadRequest) WithPayload(payload *models.GenericError) *InsertBlogBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the insert blog bad request response
func (o *InsertBlogBadRequest) SetPayload(payload *models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InsertBlogBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// InsertBlogInternalServerErrorCode is the HTTP code returned for type InsertBlogInternalServerError
const InsertBlogInternalServerErrorCode int = 500

/*InsertBlogInternalServerError Internal Server Error

swagger:response insertBlogInternalServerError
*/
type InsertBlogInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GenericError `json:"body,omitempty"`
}

// NewInsertBlogInternalServerError creates InsertBlogInternalServerError with default headers values
func NewInsertBlogInternalServerError() *InsertBlogInternalServerError {

	return &InsertBlogInternalServerError{}
}

// WithPayload adds the payload to the insert blog internal server error response
func (o *InsertBlogInternalServerError) WithPayload(payload *models.GenericError) *InsertBlogInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the insert blog internal server error response
func (o *InsertBlogInternalServerError) SetPayload(payload *models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InsertBlogInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
