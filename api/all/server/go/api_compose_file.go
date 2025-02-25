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
	"encoding/json"
	"net/http"
	"strings"
)

// ComposeFileAPIController binds http requests to an api service and writes the service results to the http response
type ComposeFileAPIController struct {
	service ComposeFileAPIServicer
	errorHandler ErrorHandler
}

// ComposeFileAPIOption for how the controller is set up.
type ComposeFileAPIOption func(*ComposeFileAPIController)

// WithComposeFileAPIErrorHandler inject ErrorHandler into controller
func WithComposeFileAPIErrorHandler(h ErrorHandler) ComposeFileAPIOption {
	return func(c *ComposeFileAPIController) {
		c.errorHandler = h
	}
}

// NewComposeFileAPIController creates a default api controller
func NewComposeFileAPIController(s ComposeFileAPIServicer, opts ...ComposeFileAPIOption) *ComposeFileAPIController {
	controller := &ComposeFileAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ComposeFileAPIController
func (c *ComposeFileAPIController) Routes() Routes {
	return Routes{
		"WorkspaceWorkspaceIdComposeVerifyPost": Route{
			strings.ToUpper("Post"),
			"/internal/v1/workspace/{workspace_id}/compose/verify",
			c.WorkspaceWorkspaceIdComposeVerifyPost,
		},
	}
}

// WorkspaceWorkspaceIdComposeVerifyPost - 
func (c *ComposeFileAPIController) WorkspaceWorkspaceIdComposeVerifyPost(w http.ResponseWriter, r *http.Request) {
	verifyComposeFileParam := VerifyComposeFile{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&verifyComposeFileParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertVerifyComposeFileRequired(verifyComposeFileParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertVerifyComposeFileConstraints(verifyComposeFileParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.WorkspaceWorkspaceIdComposeVerifyPost(r.Context(), verifyComposeFileParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
