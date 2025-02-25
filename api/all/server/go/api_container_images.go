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

	"github.com/gorilla/mux"
)

// ContainerImagesAPIController binds http requests to an api service and writes the service results to the http response
type ContainerImagesAPIController struct {
	service ContainerImagesAPIServicer
	errorHandler ErrorHandler
}

// ContainerImagesAPIOption for how the controller is set up.
type ContainerImagesAPIOption func(*ContainerImagesAPIController)

// WithContainerImagesAPIErrorHandler inject ErrorHandler into controller
func WithContainerImagesAPIErrorHandler(h ErrorHandler) ContainerImagesAPIOption {
	return func(c *ContainerImagesAPIController) {
		c.errorHandler = h
	}
}

// NewContainerImagesAPIController creates a default api controller
func NewContainerImagesAPIController(s ContainerImagesAPIServicer, opts ...ContainerImagesAPIOption) *ContainerImagesAPIController {
	controller := &ContainerImagesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ContainerImagesAPIController
func (c *ContainerImagesAPIController) Routes() Routes {
	return Routes{
		"RepositoriesContainerContainerRepositoryIdImagesExternalGet": Route{
			strings.ToUpper("Get"),
			"/internal/v1/repositories/container/{container_repository_id}/images/external",
			c.RepositoriesContainerContainerRepositoryIdImagesExternalGet,
		},
	}
}

// RepositoriesContainerContainerRepositoryIdImagesExternalGet - 
func (c *ContainerImagesAPIController) RepositoriesContainerContainerRepositoryIdImagesExternalGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	containerRepositoryIdParam := params["container_repository_id"]
	if containerRepositoryIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"container_repository_id"}, nil)
		return
	}
	var pageSizeParam int32
	if query.Has("page_size") {
		param, err := parseNumericParameter[int32](
			query.Get("page_size"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "page_size", Err: err}, nil)
			return
		}

		pageSizeParam = param
	} else {
		var param int32 = 25
		pageSizeParam = param
	}
	result, err := c.service.RepositoriesContainerContainerRepositoryIdImagesExternalGet(r.Context(), containerRepositoryIdParam, pageSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
