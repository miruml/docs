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

// ContainerRepositoriesAPIService is a service that implements the logic for the ContainerRepositoriesAPIServicer
// This service should implement the business logic for every endpoint for the ContainerRepositoriesAPI API.
// Include any external packages or services that will be required by this service.
type ContainerRepositoriesAPIService struct {
}

// NewContainerRepositoriesAPIService creates a default api service
func NewContainerRepositoriesAPIService() *ContainerRepositoriesAPIService {
	return &ContainerRepositoriesAPIService{}
}

// RepositoriesContainerContainerRepositoryIdGet - 
func (s *ContainerRepositoriesAPIService) RepositoriesContainerContainerRepositoryIdGet(ctx context.Context, containerRepositoryId string) (ImplResponse, error) {
	// TODO - update RepositoriesContainerContainerRepositoryIdGet with the required logic for this service method.
	// Add api_container_repositories_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ContainerRepository{}) or use other options such as http.Ok ...
	// return Response(200, ContainerRepository{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RepositoriesContainerContainerRepositoryIdGet method not implemented")
}

// IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet - 
func (s *ContainerRepositoriesAPIService) IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet(ctx context.Context, dockerhubIntegrationId string, pageSize int32) (ImplResponse, error) {
	// TODO - update IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet with the required logic for this service method.
	// Add api_container_repositories_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExternalContainerRepositoryList{}) or use other options such as http.Ok ...
	// return Response(200, ExternalContainerRepositoryList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet method not implemented")
}

// IntegrationsGarGarIntegrationIdRepositoriesExternalGet - 
func (s *ContainerRepositoriesAPIService) IntegrationsGarGarIntegrationIdRepositoriesExternalGet(ctx context.Context, garIntegrationId string, pageSize int32) (ImplResponse, error) {
	// TODO - update IntegrationsGarGarIntegrationIdRepositoriesExternalGet with the required logic for this service method.
	// Add api_container_repositories_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExternalContainerRepositoryList{}) or use other options such as http.Ok ...
	// return Response(200, ExternalContainerRepositoryList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("IntegrationsGarGarIntegrationIdRepositoriesExternalGet method not implemented")
}

// WorkspacesWorkspaceIdRepositoriesContainerExternalGet - 
func (s *ContainerRepositoriesAPIService) WorkspacesWorkspaceIdRepositoriesContainerExternalGet(ctx context.Context, workspaceId string, pageSize int32) (ImplResponse, error) {
	// TODO - update WorkspacesWorkspaceIdRepositoriesContainerExternalGet with the required logic for this service method.
	// Add api_container_repositories_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExternalContainerRepositoryList{}) or use other options such as http.Ok ...
	// return Response(200, ExternalContainerRepositoryList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("WorkspacesWorkspaceIdRepositoriesContainerExternalGet method not implemented")
}