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
	"context"
	"net/http"
	"errors"
)

// RegistrySourcesAPIService is a service that implements the logic for the RegistrySourcesAPIServicer
// This service should implement the business logic for every endpoint for the RegistrySourcesAPI API.
// Include any external packages or services that will be required by this service.
type RegistrySourcesAPIService struct {
}

// NewRegistrySourcesAPIService creates a default api service
func NewRegistrySourcesAPIService() *RegistrySourcesAPIService {
	return &RegistrySourcesAPIService{}
}

// SourcesRegistryRegistrySourceIdGet - 
func (s *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdGet(ctx context.Context, registrySourceId string) (ImplResponse, error) {
	// TODO - update SourcesRegistryRegistrySourceIdGet with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySource{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySource{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SourcesRegistryRegistrySourceIdGet method not implemented")
}

// SourcesRegistryRegistrySourceIdDelete - 
func (s *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdDelete(ctx context.Context, registrySourceId string) (ImplResponse, error) {
	// TODO - update SourcesRegistryRegistrySourceIdDelete with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySource{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySource{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SourcesRegistryRegistrySourceIdDelete method not implemented")
}

// SourcesRegistryRegistrySourceIdPatch - 
func (s *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdPatch(ctx context.Context, registrySourceId string, updateRegistrySource UpdateRegistrySource) (ImplResponse, error) {
	// TODO - update SourcesRegistryRegistrySourceIdPatch with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySource{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySource{}), nil

	// TODO: Uncomment the next line to return response Response(400, InvalidComposeFile{}) or use other options such as http.Ok ...
	// return Response(400, InvalidComposeFile{}), nil

	// TODO: Uncomment the next line to return response Response(404, RepositoryNotFound{}) or use other options such as http.Ok ...
	// return Response(404, RepositoryNotFound{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SourcesRegistryRegistrySourceIdPatch method not implemented")
}

// SourcesRegistryRegistrySourceIdComposeFileGet - 
func (s *RegistrySourcesAPIService) SourcesRegistryRegistrySourceIdComposeFileGet(ctx context.Context, registrySourceId string) (ImplResponse, error) {
	// TODO - update SourcesRegistryRegistrySourceIdComposeFileGet with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ComposeFile{}) or use other options such as http.Ok ...
	// return Response(200, ComposeFile{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SourcesRegistryRegistrySourceIdComposeFileGet method not implemented")
}

// WorkspacesWorkspaceIdSourcesRegistryGet - 
func (s *RegistrySourcesAPIService) WorkspacesWorkspaceIdSourcesRegistryGet(ctx context.Context, workspaceId string) (ImplResponse, error) {
	// TODO - update WorkspacesWorkspaceIdSourcesRegistryGet with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySourceList{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySourceList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("WorkspacesWorkspaceIdSourcesRegistryGet method not implemented")
}

// WorkspacesWorkspaceIdSourcesRegistryPost - 
func (s *RegistrySourcesAPIService) WorkspacesWorkspaceIdSourcesRegistryPost(ctx context.Context, workspaceId string, createRegistrySource CreateRegistrySource) (ImplResponse, error) {
	// TODO - update WorkspacesWorkspaceIdSourcesRegistryPost with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySource{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySource{}), nil

	// TODO: Uncomment the next line to return response Response(400, InvalidComposeFile{}) or use other options such as http.Ok ...
	// return Response(400, InvalidComposeFile{}), nil

	// TODO: Uncomment the next line to return response Response(404, RepositoryNotFound{}) or use other options such as http.Ok ...
	// return Response(404, RepositoryNotFound{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("WorkspacesWorkspaceIdSourcesRegistryPost method not implemented")
}
