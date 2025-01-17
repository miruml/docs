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


// ContainerRepositoriesAPIService ContainerRepositoriesAPI service
type ContainerRepositoriesAPIService service

type ApiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest struct {
	ctx context.Context
	ApiService *ContainerRepositoriesAPIService
	dockerhubIntegrationId string
	pageSize *int32
}

// The number of items to return per page
func (r ApiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest) PageSize(pageSize int32) ApiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest) Execute() (*ExternalContainerRepositoryList, *http.Response, error) {
	return r.ApiService.IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetExecute(r)
}

/*
IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet Method for IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dockerhubIntegrationId The unique identifier of the dockerhub integration
 @return ApiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest
*/
func (a *ContainerRepositoriesAPIService) IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet(ctx context.Context, dockerhubIntegrationId string) ApiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest {
	return ApiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest{
		ApiService: a,
		ctx: ctx,
		dockerhubIntegrationId: dockerhubIntegrationId,
	}
}

// Execute executes the request
//  @return ExternalContainerRepositoryList
func (a *ContainerRepositoriesAPIService) IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetExecute(r ApiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest) (*ExternalContainerRepositoryList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalContainerRepositoryList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerRepositoriesAPIService.IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/integrations/dockerhub/{dockerhub_integration_id}/repositories/external"
	localVarPath = strings.Replace(localVarPath, "{"+"dockerhub_integration_id"+"}", url.PathEscape(parameterValueToString(r.dockerhubIntegrationId, "dockerhubIntegrationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "", "")
	} else {
		var defaultValue int32 = 25
		r.pageSize = &defaultValue
	}
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

type ApiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest struct {
	ctx context.Context
	ApiService *ContainerRepositoriesAPIService
	garIntegrationId string
	pageSize *int32
}

// The number of items to return per page
func (r ApiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest) PageSize(pageSize int32) ApiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest) Execute() (*ExternalContainerRepositoryList, *http.Response, error) {
	return r.ApiService.IntegrationsGarGarIntegrationIdRepositoriesExternalGetExecute(r)
}

/*
IntegrationsGarGarIntegrationIdRepositoriesExternalGet Method for IntegrationsGarGarIntegrationIdRepositoriesExternalGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param garIntegrationId The unique identifier of the google artifact registry integration
 @return ApiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest
*/
func (a *ContainerRepositoriesAPIService) IntegrationsGarGarIntegrationIdRepositoriesExternalGet(ctx context.Context, garIntegrationId string) ApiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest {
	return ApiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest{
		ApiService: a,
		ctx: ctx,
		garIntegrationId: garIntegrationId,
	}
}

// Execute executes the request
//  @return ExternalContainerRepositoryList
func (a *ContainerRepositoriesAPIService) IntegrationsGarGarIntegrationIdRepositoriesExternalGetExecute(r ApiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest) (*ExternalContainerRepositoryList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalContainerRepositoryList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerRepositoriesAPIService.IntegrationsGarGarIntegrationIdRepositoriesExternalGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/integrations/gar/{gar_integration_id}/repositories/external"
	localVarPath = strings.Replace(localVarPath, "{"+"gar_integration_id"+"}", url.PathEscape(parameterValueToString(r.garIntegrationId, "garIntegrationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "", "")
	} else {
		var defaultValue int32 = 25
		r.pageSize = &defaultValue
	}
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

type ApiRepositoriesContainerContainerRepositoryIdGetRequest struct {
	ctx context.Context
	ApiService *ContainerRepositoriesAPIService
	containerRepositoryId string
}

func (r ApiRepositoriesContainerContainerRepositoryIdGetRequest) Execute() (*ContainerRepository, *http.Response, error) {
	return r.ApiService.RepositoriesContainerContainerRepositoryIdGetExecute(r)
}

/*
RepositoriesContainerContainerRepositoryIdGet Method for RepositoriesContainerContainerRepositoryIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerRepositoryId The unique identifier of the container repository
 @return ApiRepositoriesContainerContainerRepositoryIdGetRequest
*/
func (a *ContainerRepositoriesAPIService) RepositoriesContainerContainerRepositoryIdGet(ctx context.Context, containerRepositoryId string) ApiRepositoriesContainerContainerRepositoryIdGetRequest {
	return ApiRepositoriesContainerContainerRepositoryIdGetRequest{
		ApiService: a,
		ctx: ctx,
		containerRepositoryId: containerRepositoryId,
	}
}

// Execute executes the request
//  @return ContainerRepository
func (a *ContainerRepositoriesAPIService) RepositoriesContainerContainerRepositoryIdGetExecute(r ApiRepositoriesContainerContainerRepositoryIdGetRequest) (*ContainerRepository, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContainerRepository
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerRepositoriesAPIService.RepositoriesContainerContainerRepositoryIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repositories/container/{container_repository_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"container_repository_id"+"}", url.PathEscape(parameterValueToString(r.containerRepositoryId, "containerRepositoryId")), -1)

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

type ApiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest struct {
	ctx context.Context
	ApiService *ContainerRepositoriesAPIService
	workspaceId string
	pageSize *int32
}

// The number of items to return per page
func (r ApiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest) PageSize(pageSize int32) ApiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest) Execute() (*ExternalContainerRepositoryList, *http.Response, error) {
	return r.ApiService.WorkspacesWorkspaceIdRepositoriesContainerExternalGetExecute(r)
}

/*
WorkspacesWorkspaceIdRepositoriesContainerExternalGet Method for WorkspacesWorkspaceIdRepositoriesContainerExternalGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId The unique identifier of the workspace
 @return ApiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest
*/
func (a *ContainerRepositoriesAPIService) WorkspacesWorkspaceIdRepositoriesContainerExternalGet(ctx context.Context, workspaceId string) ApiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest {
	return ApiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest{
		ApiService: a,
		ctx: ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return ExternalContainerRepositoryList
func (a *ContainerRepositoriesAPIService) WorkspacesWorkspaceIdRepositoriesContainerExternalGetExecute(r ApiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest) (*ExternalContainerRepositoryList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalContainerRepositoryList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerRepositoriesAPIService.WorkspacesWorkspaceIdRepositoriesContainerExternalGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces/{workspace_id}/repositories/container/external"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_id"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "", "")
	} else {
		var defaultValue int32 = 25
		r.pageSize = &defaultValue
	}
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