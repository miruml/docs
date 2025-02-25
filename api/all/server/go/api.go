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
)



// ArtifactsAPIRouter defines the required methods for binding the api requests to a responses for the ArtifactsAPI
// The ArtifactsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ArtifactsAPIServicer to perform the required actions, then write the service results to the http response.
type ArtifactsAPIRouter interface { 
	ArtifactsArtifactIdGet(http.ResponseWriter, *http.Request)
	ArtifactsArtifactIdDeploymentsGet(http.ResponseWriter, *http.Request)
	ArtifactsArtifactIdDeploymentFilesGet(http.ResponseWriter, *http.Request)
	GroupsGroupIdArtifactsGet(http.ResponseWriter, *http.Request)
	SourcesRegistryRegistrySourceIdArtifactsPost(http.ResponseWriter, *http.Request)
	SourcesGithubGithubSourceIdArtifactsPost(http.ResponseWriter, *http.Request)
}
// ComposeFileAPIRouter defines the required methods for binding the api requests to a responses for the ComposeFileAPI
// The ComposeFileAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ComposeFileAPIServicer to perform the required actions, then write the service results to the http response.
type ComposeFileAPIRouter interface { 
	WorkspaceWorkspaceIdComposeVerifyPost(http.ResponseWriter, *http.Request)
}
// ContainerImagesAPIRouter defines the required methods for binding the api requests to a responses for the ContainerImagesAPI
// The ContainerImagesAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ContainerImagesAPIServicer to perform the required actions, then write the service results to the http response.
type ContainerImagesAPIRouter interface { 
	RepositoriesContainerContainerRepositoryIdImagesExternalGet(http.ResponseWriter, *http.Request)
}
// ContainerRepositoriesAPIRouter defines the required methods for binding the api requests to a responses for the ContainerRepositoriesAPI
// The ContainerRepositoriesAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ContainerRepositoriesAPIServicer to perform the required actions, then write the service results to the http response.
type ContainerRepositoriesAPIRouter interface { 
	RepositoriesContainerContainerRepositoryIdGet(http.ResponseWriter, *http.Request)
	IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet(http.ResponseWriter, *http.Request)
	IntegrationsGarGarIntegrationIdRepositoriesExternalGet(http.ResponseWriter, *http.Request)
	WorkspacesWorkspaceIdRepositoriesContainerExternalGet(http.ResponseWriter, *http.Request)
}
// ContainersAPIRouter defines the required methods for binding the api requests to a responses for the ContainersAPI
// The ContainersAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ContainersAPIServicer to perform the required actions, then write the service results to the http response.
type ContainersAPIRouter interface { 
	DevicesDeviceIdContainersGet(http.ResponseWriter, *http.Request)
}
// DeploymentsAPIRouter defines the required methods for binding the api requests to a responses for the DeploymentsAPI
// The DeploymentsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DeploymentsAPIServicer to perform the required actions, then write the service results to the http response.
type DeploymentsAPIRouter interface { 
	DeploymentsDeploymentIdGet(http.ResponseWriter, *http.Request)
	DevicesDeviceIdDeploymentsGet(http.ResponseWriter, *http.Request)
	WorkspacesWorkspaceIdDeployPost(http.ResponseWriter, *http.Request)
	DeploymentsDeploymentIdLogsGet(http.ResponseWriter, *http.Request)
}
// DevicesAPIRouter defines the required methods for binding the api requests to a responses for the DevicesAPI
// The DevicesAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DevicesAPIServicer to perform the required actions, then write the service results to the http response.
type DevicesAPIRouter interface { 
	DevicesDeviceIdGet(http.ResponseWriter, *http.Request)
}
// GenericErrorsAPIRouter defines the required methods for binding the api requests to a responses for the GenericErrorsAPI
// The GenericErrorsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a GenericErrorsAPIServicer to perform the required actions, then write the service results to the http response.
type GenericErrorsAPIRouter interface { 
	ExampleErrors400Get(http.ResponseWriter, *http.Request)
	ExampleErrors401Get(http.ResponseWriter, *http.Request)
	ExampleErrors403Get(http.ResponseWriter, *http.Request)
	ExampleErrors404Get(http.ResponseWriter, *http.Request)
}
// GitHubSourcesAPIRouter defines the required methods for binding the api requests to a responses for the GitHubSourcesAPI
// The GitHubSourcesAPIRouter implementation should parse necessary information from the http request,
// pass the data to a GitHubSourcesAPIServicer to perform the required actions, then write the service results to the http response.
type GitHubSourcesAPIRouter interface { 
	SourcesGithubGithubSourceIdCommitsGet(http.ResponseWriter, *http.Request)
}
// GroupsAPIRouter defines the required methods for binding the api requests to a responses for the GroupsAPI
// The GroupsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a GroupsAPIServicer to perform the required actions, then write the service results to the http response.
type GroupsAPIRouter interface { 
	GroupsGroupIdGet(http.ResponseWriter, *http.Request)
}
// RegistrySourcesAPIRouter defines the required methods for binding the api requests to a responses for the RegistrySourcesAPI
// The RegistrySourcesAPIRouter implementation should parse necessary information from the http request,
// pass the data to a RegistrySourcesAPIServicer to perform the required actions, then write the service results to the http response.
type RegistrySourcesAPIRouter interface { 
	SourcesRegistryRegistrySourceIdGet(http.ResponseWriter, *http.Request)
	SourcesRegistryRegistrySourceIdDelete(http.ResponseWriter, *http.Request)
	SourcesRegistryRegistrySourceIdPatch(http.ResponseWriter, *http.Request)
	SourcesRegistryRegistrySourceIdComposeFileGet(http.ResponseWriter, *http.Request)
	WorkspacesWorkspaceIdSourcesRegistryGet(http.ResponseWriter, *http.Request)
	WorkspacesWorkspaceIdSourcesRegistryPost(http.ResponseWriter, *http.Request)
}


// ArtifactsAPIServicer defines the api actions for the ArtifactsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ArtifactsAPIServicer interface { 
	ArtifactsArtifactIdGet(context.Context, string) (ImplResponse, error)
	ArtifactsArtifactIdDeploymentsGet(context.Context, string) (ImplResponse, error)
	ArtifactsArtifactIdDeploymentFilesGet(context.Context, string) (ImplResponse, error)
	GroupsGroupIdArtifactsGet(context.Context, string, bool) (ImplResponse, error)
	SourcesRegistryRegistrySourceIdArtifactsPost(context.Context, string, CreateRegistrySourceArtifact) (ImplResponse, error)
	SourcesGithubGithubSourceIdArtifactsPost(context.Context, string, CreateGitHubSourceArtifact) (ImplResponse, error)
}


// ComposeFileAPIServicer defines the api actions for the ComposeFileAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ComposeFileAPIServicer interface { 
	WorkspaceWorkspaceIdComposeVerifyPost(context.Context, VerifyComposeFile) (ImplResponse, error)
}


// ContainerImagesAPIServicer defines the api actions for the ContainerImagesAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ContainerImagesAPIServicer interface { 
	RepositoriesContainerContainerRepositoryIdImagesExternalGet(context.Context, string, int32) (ImplResponse, error)
}


// ContainerRepositoriesAPIServicer defines the api actions for the ContainerRepositoriesAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ContainerRepositoriesAPIServicer interface { 
	RepositoriesContainerContainerRepositoryIdGet(context.Context, string) (ImplResponse, error)
	IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet(context.Context, string, int32) (ImplResponse, error)
	IntegrationsGarGarIntegrationIdRepositoriesExternalGet(context.Context, string, int32) (ImplResponse, error)
	WorkspacesWorkspaceIdRepositoriesContainerExternalGet(context.Context, string, int32) (ImplResponse, error)
}


// ContainersAPIServicer defines the api actions for the ContainersAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ContainersAPIServicer interface { 
	DevicesDeviceIdContainersGet(context.Context, string) (ImplResponse, error)
}


// DeploymentsAPIServicer defines the api actions for the DeploymentsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DeploymentsAPIServicer interface { 
	DeploymentsDeploymentIdGet(context.Context, string) (ImplResponse, error)
	DevicesDeviceIdDeploymentsGet(context.Context, string, bool) (ImplResponse, error)
	WorkspacesWorkspaceIdDeployPost(context.Context, string, CreateDeployments) (ImplResponse, error)
	DeploymentsDeploymentIdLogsGet(context.Context, string) (ImplResponse, error)
}


// DevicesAPIServicer defines the api actions for the DevicesAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DevicesAPIServicer interface { 
	DevicesDeviceIdGet(context.Context, string) (ImplResponse, error)
}


// GenericErrorsAPIServicer defines the api actions for the GenericErrorsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type GenericErrorsAPIServicer interface { 
	ExampleErrors400Get(context.Context) (ImplResponse, error)
	ExampleErrors401Get(context.Context) (ImplResponse, error)
	ExampleErrors403Get(context.Context) (ImplResponse, error)
	ExampleErrors404Get(context.Context) (ImplResponse, error)
}


// GitHubSourcesAPIServicer defines the api actions for the GitHubSourcesAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type GitHubSourcesAPIServicer interface { 
	SourcesGithubGithubSourceIdCommitsGet(context.Context, string, string) (ImplResponse, error)
}


// GroupsAPIServicer defines the api actions for the GroupsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type GroupsAPIServicer interface { 
	GroupsGroupIdGet(context.Context, string) (ImplResponse, error)
}


// RegistrySourcesAPIServicer defines the api actions for the RegistrySourcesAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type RegistrySourcesAPIServicer interface { 
	SourcesRegistryRegistrySourceIdGet(context.Context, string) (ImplResponse, error)
	SourcesRegistryRegistrySourceIdDelete(context.Context, string) (ImplResponse, error)
	SourcesRegistryRegistrySourceIdPatch(context.Context, string, UpdateRegistrySource) (ImplResponse, error)
	SourcesRegistryRegistrySourceIdComposeFileGet(context.Context, string) (ImplResponse, error)
	WorkspacesWorkspaceIdSourcesRegistryGet(context.Context, string) (ImplResponse, error)
	WorkspacesWorkspaceIdSourcesRegistryPost(context.Context, string, CreateRegistrySource) (ImplResponse, error)
}
