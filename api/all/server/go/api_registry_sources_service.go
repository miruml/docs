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

// WorkspacesWorkspaceIdSourcesRegistriesGet - 
func (s *RegistrySourcesAPIService) WorkspacesWorkspaceIdSourcesRegistriesGet(ctx context.Context, workspaceId string) (ImplResponse, error) {
	// TODO - update WorkspacesWorkspaceIdSourcesRegistriesGet with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySourceList{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySourceList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("WorkspacesWorkspaceIdSourcesRegistriesGet method not implemented")
}

// WorkspacesWorkspaceIdSourcesRegistriesPost - 
func (s *RegistrySourcesAPIService) WorkspacesWorkspaceIdSourcesRegistriesPost(ctx context.Context, workspaceId string, createRegistrySource CreateRegistrySource) (ImplResponse, error) {
	// TODO - update WorkspacesWorkspaceIdSourcesRegistriesPost with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySource{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySource{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("WorkspacesWorkspaceIdSourcesRegistriesPost method not implemented")
}

// SourcesRegistriesRegistrySourceIdGet - 
func (s *RegistrySourcesAPIService) SourcesRegistriesRegistrySourceIdGet(ctx context.Context, registrySourceId string) (ImplResponse, error) {
	// TODO - update SourcesRegistriesRegistrySourceIdGet with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySource{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySource{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SourcesRegistriesRegistrySourceIdGet method not implemented")
}

// SourcesRegistriesRegistrySourceIdPatch - 
func (s *RegistrySourcesAPIService) SourcesRegistriesRegistrySourceIdPatch(ctx context.Context, registrySourceId string, updateRegistrySource UpdateRegistrySource) (ImplResponse, error) {
	// TODO - update SourcesRegistriesRegistrySourceIdPatch with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegistrySource{}) or use other options such as http.Ok ...
	// return Response(200, RegistrySource{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SourcesRegistriesRegistrySourceIdPatch method not implemented")
}

// SourcesRegistriesRegistrySourceIdComposeFileGet - 
func (s *RegistrySourcesAPIService) SourcesRegistriesRegistrySourceIdComposeFileGet(ctx context.Context, registrySourceId string) (ImplResponse, error) {
	// TODO - update SourcesRegistriesRegistrySourceIdComposeFileGet with the required logic for this service method.
	// Add api_registry_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ComposeFile{}) or use other options such as http.Ok ...
	// return Response(200, ComposeFile{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SourcesRegistriesRegistrySourceIdComposeFileGet method not implemented")
}
