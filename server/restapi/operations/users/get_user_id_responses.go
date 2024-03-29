// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/base-swagger-spec/server/models"
)

// GetUserIDOKCode is the HTTP code returned for type GetUserIDOK
const GetUserIDOKCode int = 200

/*GetUserIDOK Successfully got user

swagger:response getUserIdOK
*/
type GetUserIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserIDOK creates GetUserIDOK with default headers values
func NewGetUserIDOK() *GetUserIDOK {

	return &GetUserIDOK{}
}

// WithPayload adds the payload to the get user Id o k response
func (o *GetUserIDOK) WithPayload(payload *models.User) *GetUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user Id o k response
func (o *GetUserIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserIDInternalServerErrorCode is the HTTP code returned for type GetUserIDInternalServerError
const GetUserIDInternalServerErrorCode int = 500

/*GetUserIDInternalServerError Internal server error

swagger:response getUserIdInternalServerError
*/
type GetUserIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserIDInternalServerError creates GetUserIDInternalServerError with default headers values
func NewGetUserIDInternalServerError() *GetUserIDInternalServerError {

	return &GetUserIDInternalServerError{}
}

// WithPayload adds the payload to the get user Id internal server error response
func (o *GetUserIDInternalServerError) WithPayload(payload *models.Error) *GetUserIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user Id internal server error response
func (o *GetUserIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
