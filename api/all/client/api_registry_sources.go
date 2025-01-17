/*
Miru API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// RegistrySourcesAPIService RegistrySourcesAPI service
type RegistrySourcesAPIService service

type ApiSourcesRegistryRegistrySourceIdComposeFileGetRequest struct {
	ctx context.Context
	ApiService *RegistrySourcesAPIService
	registrySourceId string
}

func (r ApiSourcesRegistryRegistrySourceIdComposeFileGetRequest) Execute() (*ComposeFile, *http.Response, error) {
	return r.ApiService.SourcesRegistryRegistrySourceIdComposeFileGetExecute(r)
}

/*
SourcesRegistryRegistrySourceIdComposeFileGet Method for SourcesRegistryRegistrySourceIdComposeFileGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registrySourceId The unique identifier of the registry source
 @return ApiSourcesRegistryRegistrySourceIdComposeFileGetRequest
*/
func (a *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdComposeFileGet(ctx context.Context, registrySourceId string) ApiSourcesRegistryRegistrySourceIdComposeFileGetRequest {
	return ApiSourcesRegistryRegistrySourceIdComposeFileGetRequest{
		ApiService: a,
		ctx: ctx,
		registrySourceId: registrySourceId,
	}
}

// Execute executes the request
//  @return ComposeFile
func (a *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdComposeFileGetExecute(r ApiSourcesRegistryRegistrySourceIdComposeFileGetRequest) (*ComposeFile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComposeFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrySourcesAPIService.SourcesRegistryRegistrySourceIdComposeFileGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/registry/{registry_source_id}/compose_file"
	localVarPath = strings.Replace(localVarPath, "{"+"registry_source_id"+"}", url.PathEscape(parameterValueToString(r.registrySourceId, "registrySourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSourcesRegistryRegistrySourceIdDeleteRequest struct {
	ctx context.Context
	ApiService *RegistrySourcesAPIService
	registrySourceId string
}

func (r ApiSourcesRegistryRegistrySourceIdDeleteRequest) Execute() (*RegistrySource, *http.Response, error) {
	return r.ApiService.SourcesRegistryRegistrySourceIdDeleteExecute(r)
}

/*
SourcesRegistryRegistrySourceIdDelete Method for SourcesRegistryRegistrySourceIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registrySourceId The unique identifier of the registry source
 @return ApiSourcesRegistryRegistrySourceIdDeleteRequest
*/
func (a *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdDelete(ctx context.Context, registrySourceId string) ApiSourcesRegistryRegistrySourceIdDeleteRequest {
	return ApiSourcesRegistryRegistrySourceIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		registrySourceId: registrySourceId,
	}
}

// Execute executes the request
//  @return RegistrySource
func (a *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdDeleteExecute(r ApiSourcesRegistryRegistrySourceIdDeleteRequest) (*RegistrySource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegistrySource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrySourcesAPIService.SourcesRegistryRegistrySourceIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/registry/{registry_source_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"registry_source_id"+"}", url.PathEscape(parameterValueToString(r.registrySourceId, "registrySourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSourcesRegistryRegistrySourceIdGetRequest struct {
	ctx context.Context
	ApiService *RegistrySourcesAPIService
	registrySourceId string
}

func (r ApiSourcesRegistryRegistrySourceIdGetRequest) Execute() (*RegistrySource, *http.Response, error) {
	return r.ApiService.SourcesRegistryRegistrySourceIdGetExecute(r)
}

/*
SourcesRegistryRegistrySourceIdGet Method for SourcesRegistryRegistrySourceIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registrySourceId The unique identifier of the registry source
 @return ApiSourcesRegistryRegistrySourceIdGetRequest
*/
func (a *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdGet(ctx context.Context, registrySourceId string) ApiSourcesRegistryRegistrySourceIdGetRequest {
	return ApiSourcesRegistryRegistrySourceIdGetRequest{
		ApiService: a,
		ctx: ctx,
		registrySourceId: registrySourceId,
	}
}

// Execute executes the request
//  @return RegistrySource
func (a *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdGetExecute(r ApiSourcesRegistryRegistrySourceIdGetRequest) (*RegistrySource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegistrySource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrySourcesAPIService.SourcesRegistryRegistrySourceIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/registry/{registry_source_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"registry_source_id"+"}", url.PathEscape(parameterValueToString(r.registrySourceId, "registrySourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSourcesRegistryRegistrySourceIdPatchRequest struct {
	ctx context.Context
	ApiService *RegistrySourcesAPIService
	registrySourceId string
	updateRegistrySource *UpdateRegistrySource
}

func (r ApiSourcesRegistryRegistrySourceIdPatchRequest) UpdateRegistrySource(updateRegistrySource UpdateRegistrySource) ApiSourcesRegistryRegistrySourceIdPatchRequest {
	r.updateRegistrySource = &updateRegistrySource
	return r
}

func (r ApiSourcesRegistryRegistrySourceIdPatchRequest) Execute() (*RegistrySource, *http.Response, error) {
	return r.ApiService.SourcesRegistryRegistrySourceIdPatchExecute(r)
}

/*
SourcesRegistryRegistrySourceIdPatch Method for SourcesRegistryRegistrySourceIdPatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registrySourceId The unique identifier of the registry source
 @return ApiSourcesRegistryRegistrySourceIdPatchRequest
*/
func (a *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdPatch(ctx context.Context, registrySourceId string) ApiSourcesRegistryRegistrySourceIdPatchRequest {
	return ApiSourcesRegistryRegistrySourceIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		registrySourceId: registrySourceId,
	}
}

// Execute executes the request
//  @return RegistrySource
func (a *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdPatchExecute(r ApiSourcesRegistryRegistrySourceIdPatchRequest) (*RegistrySource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegistrySource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrySourcesAPIService.SourcesRegistryRegistrySourceIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/registry/{registry_source_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"registry_source_id"+"}", url.PathEscape(parameterValueToString(r.registrySourceId, "registrySourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRegistrySource == nil {
		return localVarReturnValue, nil, reportError("updateRegistrySource is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateRegistrySource
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidComposeFile
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v RepositoryNotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiWorkspacesWorkspaceIdSourcesRegistryGetRequest struct {
	ctx context.Context
	ApiService *RegistrySourcesAPIService
	workspaceId string
}

func (r ApiWorkspacesWorkspaceIdSourcesRegistryGetRequest) Execute() (*RegistrySourceList, *http.Response, error) {
	return r.ApiService.WorkspacesWorkspaceIdSourcesRegistryGetExecute(r)
}

/*
WorkspacesWorkspaceIdSourcesRegistryGet Method for WorkspacesWorkspaceIdSourcesRegistryGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The unique identifier of the workspace
 @return ApiWorkspacesWorkspaceIdSourcesRegistryGetRequest
*/
func (a *RegistrySourcesAPIService) WorkspacesWorkspaceIdSourcesRegistryGet(ctx context.Context, workspaceId string) ApiWorkspacesWorkspaceIdSourcesRegistryGetRequest {
	return ApiWorkspacesWorkspaceIdSourcesRegistryGetRequest{
		ApiService: a,
		ctx: ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return RegistrySourceList
func (a *RegistrySourcesAPIService) WorkspacesWorkspaceIdSourcesRegistryGetExecute(r ApiWorkspacesWorkspaceIdSourcesRegistryGetRequest) (*RegistrySourceList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegistrySourceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrySourcesAPIService.WorkspacesWorkspaceIdSourcesRegistryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces/{workspace_id}/sources/registry"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_id"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiWorkspacesWorkspaceIdSourcesRegistryPostRequest struct {
	ctx context.Context
	ApiService *RegistrySourcesAPIService
	workspaceId string
	createRegistrySource *CreateRegistrySource
}

func (r ApiWorkspacesWorkspaceIdSourcesRegistryPostRequest) CreateRegistrySource(createRegistrySource CreateRegistrySource) ApiWorkspacesWorkspaceIdSourcesRegistryPostRequest {
	r.createRegistrySource = &createRegistrySource
	return r
}

func (r ApiWorkspacesWorkspaceIdSourcesRegistryPostRequest) Execute() (*RegistrySource, *http.Response, error) {
	return r.ApiService.WorkspacesWorkspaceIdSourcesRegistryPostExecute(r)
}

/*
WorkspacesWorkspaceIdSourcesRegistryPost Method for WorkspacesWorkspaceIdSourcesRegistryPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The unique identifier of the workspace
 @return ApiWorkspacesWorkspaceIdSourcesRegistryPostRequest
*/
func (a *RegistrySourcesAPIService) WorkspacesWorkspaceIdSourcesRegistryPost(ctx context.Context, workspaceId string) ApiWorkspacesWorkspaceIdSourcesRegistryPostRequest {
	return ApiWorkspacesWorkspaceIdSourcesRegistryPostRequest{
		ApiService: a,
		ctx: ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return RegistrySource
func (a *RegistrySourcesAPIService) WorkspacesWorkspaceIdSourcesRegistryPostExecute(r ApiWorkspacesWorkspaceIdSourcesRegistryPostRequest) (*RegistrySource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegistrySource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrySourcesAPIService.WorkspacesWorkspaceIdSourcesRegistryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces/{workspace_id}/sources/registry"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_id"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createRegistrySource == nil {
		return localVarReturnValue, nil, reportError("createRegistrySource is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createRegistrySource
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidComposeFile
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v RepositoryNotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
