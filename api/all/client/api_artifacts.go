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


// ArtifactsAPIService ArtifactsAPI service
type ArtifactsAPIService service

type ApiArtifactsArtifactIdComposeFileGetRequest struct {
	ctx context.Context
	ApiService *ArtifactsAPIService
	artifactId string
}

func (r ApiArtifactsArtifactIdComposeFileGetRequest) Execute() (*ComposeFile, *http.Response, error) {
	return r.ApiService.ArtifactsArtifactIdComposeFileGetExecute(r)
}

/*
ArtifactsArtifactIdComposeFileGet Method for ArtifactsArtifactIdComposeFileGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param artifactId The unique identifier of the artifact
 @return ApiArtifactsArtifactIdComposeFileGetRequest
*/
func (a *ArtifactsAPIService) ArtifactsArtifactIdComposeFileGet(ctx context.Context, artifactId string) ApiArtifactsArtifactIdComposeFileGetRequest {
	return ApiArtifactsArtifactIdComposeFileGetRequest{
		ApiService: a,
		ctx: ctx,
		artifactId: artifactId,
	}
}

// Execute executes the request
//  @return ComposeFile
func (a *ArtifactsAPIService) ArtifactsArtifactIdComposeFileGetExecute(r ApiArtifactsArtifactIdComposeFileGetRequest) (*ComposeFile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComposeFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsAPIService.ArtifactsArtifactIdComposeFileGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifacts/{artifact_id}/compose_file"
	localVarPath = strings.Replace(localVarPath, "{"+"artifact_id"+"}", url.PathEscape(parameterValueToString(r.artifactId, "artifactId")), -1)

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

type ApiArtifactsArtifactIdGetRequest struct {
	ctx context.Context
	ApiService *ArtifactsAPIService
	artifactId string
}

func (r ApiArtifactsArtifactIdGetRequest) Execute() (*Artifact, *http.Response, error) {
	return r.ApiService.ArtifactsArtifactIdGetExecute(r)
}

/*
ArtifactsArtifactIdGet Method for ArtifactsArtifactIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param artifactId The unique identifier of the artifact
 @return ApiArtifactsArtifactIdGetRequest
*/
func (a *ArtifactsAPIService) ArtifactsArtifactIdGet(ctx context.Context, artifactId string) ApiArtifactsArtifactIdGetRequest {
	return ApiArtifactsArtifactIdGetRequest{
		ApiService: a,
		ctx: ctx,
		artifactId: artifactId,
	}
}

// Execute executes the request
//  @return Artifact
func (a *ArtifactsAPIService) ArtifactsArtifactIdGetExecute(r ApiArtifactsArtifactIdGetRequest) (*Artifact, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Artifact
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsAPIService.ArtifactsArtifactIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifacts/{artifact_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"artifact_id"+"}", url.PathEscape(parameterValueToString(r.artifactId, "artifactId")), -1)

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

type ApiSourcesGithubGithubSourceIdArtifactsPostRequest struct {
	ctx context.Context
	ApiService *ArtifactsAPIService
	githubSourceId string
	createGitHubSourceArtifact *CreateGitHubSourceArtifact
}

func (r ApiSourcesGithubGithubSourceIdArtifactsPostRequest) CreateGitHubSourceArtifact(createGitHubSourceArtifact CreateGitHubSourceArtifact) ApiSourcesGithubGithubSourceIdArtifactsPostRequest {
	r.createGitHubSourceArtifact = &createGitHubSourceArtifact
	return r
}

func (r ApiSourcesGithubGithubSourceIdArtifactsPostRequest) Execute() (*Artifact, *http.Response, error) {
	return r.ApiService.SourcesGithubGithubSourceIdArtifactsPostExecute(r)
}

/*
SourcesGithubGithubSourceIdArtifactsPost Method for SourcesGithubGithubSourceIdArtifactsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param githubSourceId The unique identifier of the github source
 @return ApiSourcesGithubGithubSourceIdArtifactsPostRequest
*/
func (a *ArtifactsAPIService) SourcesGithubGithubSourceIdArtifactsPost(ctx context.Context, githubSourceId string) ApiSourcesGithubGithubSourceIdArtifactsPostRequest {
	return ApiSourcesGithubGithubSourceIdArtifactsPostRequest{
		ApiService: a,
		ctx: ctx,
		githubSourceId: githubSourceId,
	}
}

// Execute executes the request
//  @return Artifact
func (a *ArtifactsAPIService) SourcesGithubGithubSourceIdArtifactsPostExecute(r ApiSourcesGithubGithubSourceIdArtifactsPostRequest) (*Artifact, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Artifact
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsAPIService.SourcesGithubGithubSourceIdArtifactsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/github/{github_source_id}/artifacts"
	localVarPath = strings.Replace(localVarPath, "{"+"github_source_id"+"}", url.PathEscape(parameterValueToString(r.githubSourceId, "githubSourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createGitHubSourceArtifact == nil {
		return localVarReturnValue, nil, reportError("createGitHubSourceArtifact is required and must be specified")
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
	localVarPostBody = r.createGitHubSourceArtifact
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

type ApiSourcesRegistriesRegistrySourceIdArtifactsPostRequest struct {
	ctx context.Context
	ApiService *ArtifactsAPIService
	registrySourceId string
	createRegistrySourceArtifact *CreateRegistrySourceArtifact
}

func (r ApiSourcesRegistriesRegistrySourceIdArtifactsPostRequest) CreateRegistrySourceArtifact(createRegistrySourceArtifact CreateRegistrySourceArtifact) ApiSourcesRegistriesRegistrySourceIdArtifactsPostRequest {
	r.createRegistrySourceArtifact = &createRegistrySourceArtifact
	return r
}

func (r ApiSourcesRegistriesRegistrySourceIdArtifactsPostRequest) Execute() (*Artifact, *http.Response, error) {
	return r.ApiService.SourcesRegistriesRegistrySourceIdArtifactsPostExecute(r)
}

/*
SourcesRegistriesRegistrySourceIdArtifactsPost Method for SourcesRegistriesRegistrySourceIdArtifactsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registrySourceId The unique identifier of the registry source
 @return ApiSourcesRegistriesRegistrySourceIdArtifactsPostRequest
*/
func (a *ArtifactsAPIService) SourcesRegistriesRegistrySourceIdArtifactsPost(ctx context.Context, registrySourceId string) ApiSourcesRegistriesRegistrySourceIdArtifactsPostRequest {
	return ApiSourcesRegistriesRegistrySourceIdArtifactsPostRequest{
		ApiService: a,
		ctx: ctx,
		registrySourceId: registrySourceId,
	}
}

// Execute executes the request
//  @return Artifact
func (a *ArtifactsAPIService) SourcesRegistriesRegistrySourceIdArtifactsPostExecute(r ApiSourcesRegistriesRegistrySourceIdArtifactsPostRequest) (*Artifact, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Artifact
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsAPIService.SourcesRegistriesRegistrySourceIdArtifactsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/registries/{registry_source_id}/artifacts"
	localVarPath = strings.Replace(localVarPath, "{"+"registry_source_id"+"}", url.PathEscape(parameterValueToString(r.registrySourceId, "registrySourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createRegistrySourceArtifact == nil {
		return localVarReturnValue, nil, reportError("createRegistrySourceArtifact is required and must be specified")
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
	localVarPostBody = r.createRegistrySourceArtifact
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
