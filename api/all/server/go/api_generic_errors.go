// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi

import (
	"net/http"
	"strings"
)

// GenericErrorsAPIController binds http requests to an api service and writes the service results to the http response
type GenericErrorsAPIController struct {
	service GenericErrorsAPIServicer
	errorHandler ErrorHandler
}

// GenericErrorsAPIOption for how the controller is set up.
type GenericErrorsAPIOption func(*GenericErrorsAPIController)

// WithGenericErrorsAPIErrorHandler inject ErrorHandler into controller
func WithGenericErrorsAPIErrorHandler(h ErrorHandler) GenericErrorsAPIOption {
	return func(c *GenericErrorsAPIController) {
		c.errorHandler = h
	}
}

// NewGenericErrorsAPIController creates a default api controller
func NewGenericErrorsAPIController(s GenericErrorsAPIServicer, opts ...GenericErrorsAPIOption) *GenericErrorsAPIController {
	controller := &GenericErrorsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the GenericErrorsAPIController
func (c *GenericErrorsAPIController) Routes() Routes {
	return Routes{
		"ExampleErrors400Get": Route{
			strings.ToUpper("Get"),
			"/internal/v1/example/errors/400",
			c.ExampleErrors400Get,
		},
		"ExampleErrors401Get": Route{
			strings.ToUpper("Get"),
			"/internal/v1/example/errors/401",
			c.ExampleErrors401Get,
		},
		"ExampleErrors403Get": Route{
			strings.ToUpper("Get"),
			"/internal/v1/example/errors/403",
			c.ExampleErrors403Get,
		},
		"ExampleErrors404Get": Route{
			strings.ToUpper("Get"),
			"/internal/v1/example/errors/404",
			c.ExampleErrors404Get,
		},
	}
}

// ExampleErrors400Get - 
func (c *GenericErrorsAPIController) ExampleErrors400Get(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ExampleErrors400Get(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ExampleErrors401Get - 
func (c *GenericErrorsAPIController) ExampleErrors401Get(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ExampleErrors401Get(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ExampleErrors403Get - 
func (c *GenericErrorsAPIController) ExampleErrors403Get(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ExampleErrors403Get(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ExampleErrors404Get - 
func (c *GenericErrorsAPIController) ExampleErrors404Get(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ExampleErrors404Get(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
