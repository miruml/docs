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

	"github.com/gorilla/mux"
)

// RegistrySourcesAPIController binds http requests to an api service and writes the service results to the http response
type RegistrySourcesAPIController struct {
	service RegistrySourcesAPIServicer
	errorHandler ErrorHandler
}

// RegistrySourcesAPIOption for how the controller is set up.
type RegistrySourcesAPIOption func(*RegistrySourcesAPIController)

// WithRegistrySourcesAPIErrorHandler inject ErrorHandler into controller
func WithRegistrySourcesAPIErrorHandler(h ErrorHandler) RegistrySourcesAPIOption {
	return func(c *RegistrySourcesAPIController) {
		c.errorHandler = h
	}
}

// NewRegistrySourcesAPIController creates a default api controller
func NewRegistrySourcesAPIController(s RegistrySourcesAPIServicer, opts ...RegistrySourcesAPIOption) *RegistrySourcesAPIController {
	controller := &RegistrySourcesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the RegistrySourcesAPIController
func (c *RegistrySourcesAPIController) Routes() Routes {
	return Routes{
		"SourcesRegistryRegistrySourceIdGet": Route{
			strings.ToUpper("Get"),
			"/internal/v1/sources/registry/{registry_source_id}",
			c.SourcesRegistryRegistrySourceIdGet,
		},
		"SourcesRegistryRegistrySourceIdDelete": Route{
			strings.ToUpper("Delete"),
			"/internal/v1/sources/registry/{registry_source_id}",
			c.SourcesRegistryRegistrySourceIdDelete,
		},
		"SourcesRegistryRegistrySourceIdPatch": Route{
			strings.ToUpper("Patch"),
			"/internal/v1/sources/registry/{registry_source_id}",
			c.SourcesRegistryRegistrySourceIdPatch,
		},
		"SourcesRegistryRegistrySourceIdComposeFileGet": Route{
			strings.ToUpper("Get"),
			"/internal/v1/sources/registry/{registry_source_id}/compose_file",
			c.SourcesRegistryRegistrySourceIdComposeFileGet,
		},
		"WorkspacesWorkspaceIdSourcesRegistryGet": Route{
			strings.ToUpper("Get"),
			"/internal/v1/workspaces/{workspace_id}/sources/registry",
			c.WorkspacesWorkspaceIdSourcesRegistryGet,
		},
		"WorkspacesWorkspaceIdSourcesRegistryPost": Route{
			strings.ToUpper("Post"),
			"/internal/v1/workspaces/{workspace_id}/sources/registry",
			c.WorkspacesWorkspaceIdSourcesRegistryPost,
		},
	}
}

// SourcesRegistryRegistrySourceIdGet - 
func (c *RegistrySourcesAPIController) SourcesRegistryRegistrySourceIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrySourceIdParam := params["registry_source_id"]
	if registrySourceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"registry_source_id"}, nil)
		return
	}
	result, err := c.service.SourcesRegistryRegistrySourceIdGet(r.Context(), registrySourceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// SourcesRegistryRegistrySourceIdDelete - 
func (c *RegistrySourcesAPIController) SourcesRegistryRegistrySourceIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrySourceIdParam := params["registry_source_id"]
	if registrySourceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"registry_source_id"}, nil)
		return
	}
	result, err := c.service.SourcesRegistryRegistrySourceIdDelete(r.Context(), registrySourceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// SourcesRegistryRegistrySourceIdPatch - 
func (c *RegistrySourcesAPIController) SourcesRegistryRegistrySourceIdPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrySourceIdParam := params["registry_source_id"]
	if registrySourceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"registry_source_id"}, nil)
		return
	}
	updateRegistrySourceParam := UpdateRegistrySource{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&updateRegistrySourceParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUpdateRegistrySourceRequired(updateRegistrySourceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUpdateRegistrySourceConstraints(updateRegistrySourceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.SourcesRegistryRegistrySourceIdPatch(r.Context(), registrySourceIdParam, updateRegistrySourceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// SourcesRegistryRegistrySourceIdComposeFileGet - 
func (c *RegistrySourcesAPIController) SourcesRegistryRegistrySourceIdComposeFileGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrySourceIdParam := params["registry_source_id"]
	if registrySourceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"registry_source_id"}, nil)
		return
	}
	result, err := c.service.SourcesRegistryRegistrySourceIdComposeFileGet(r.Context(), registrySourceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// WorkspacesWorkspaceIdSourcesRegistryGet - 
func (c *RegistrySourcesAPIController) WorkspacesWorkspaceIdSourcesRegistryGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workspaceIdParam := params["workspace_id"]
	if workspaceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"workspace_id"}, nil)
		return
	}
	result, err := c.service.WorkspacesWorkspaceIdSourcesRegistryGet(r.Context(), workspaceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// WorkspacesWorkspaceIdSourcesRegistryPost - 
func (c *RegistrySourcesAPIController) WorkspacesWorkspaceIdSourcesRegistryPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workspaceIdParam := params["workspace_id"]
	if workspaceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"workspace_id"}, nil)
		return
	}
	createRegistrySourceParam := CreateRegistrySource{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createRegistrySourceParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateRegistrySourceRequired(createRegistrySourceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCreateRegistrySourceConstraints(createRegistrySourceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.WorkspacesWorkspaceIdSourcesRegistryPost(r.Context(), workspaceIdParam, createRegistrySourceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
