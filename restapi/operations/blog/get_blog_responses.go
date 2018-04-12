// Code generated by go-swagger; DO NOT EDIT.

package blog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/bbcyyb/bunkerhill/models"
)

// GetBlogOKCode is the HTTP code returned for type GetBlogOK
const GetBlogOKCode int = 200

/*GetBlogOK Successful response

swagger:response getBlogOK
*/
type GetBlogOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetBlogOKBody `json:"body,omitempty"`
}

// NewGetBlogOK creates GetBlogOK with default headers values
func NewGetBlogOK() *GetBlogOK {

	return &GetBlogOK{}
}

// WithPayload adds the payload to the get blog o k response
func (o *GetBlogOK) WithPayload(payload *models.GetBlogOKBody) *GetBlogOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blog o k response
func (o *GetBlogOK) SetPayload(payload *models.GetBlogOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlogOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}