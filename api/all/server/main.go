// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	ArtifactsAPIService := openapi.NewArtifactsAPIService()
	ArtifactsAPIController := openapi.NewArtifactsAPIController(ArtifactsAPIService)

	ComposeFileAPIService := openapi.NewComposeFileAPIService()
	ComposeFileAPIController := openapi.NewComposeFileAPIController(ComposeFileAPIService)

	ContainerImagesAPIService := openapi.NewContainerImagesAPIService()
	ContainerImagesAPIController := openapi.NewContainerImagesAPIController(ContainerImagesAPIService)

	ContainerRepositoriesAPIService := openapi.NewContainerRepositoriesAPIService()
	ContainerRepositoriesAPIController := openapi.NewContainerRepositoriesAPIController(ContainerRepositoriesAPIService)

	ContainersAPIService := openapi.NewContainersAPIService()
	ContainersAPIController := openapi.NewContainersAPIController(ContainersAPIService)

	DeploymentsAPIService := openapi.NewDeploymentsAPIService()
	DeploymentsAPIController := openapi.NewDeploymentsAPIController(DeploymentsAPIService)

	DevicesAPIService := openapi.NewDevicesAPIService()
	DevicesAPIController := openapi.NewDevicesAPIController(DevicesAPIService)

	GenericErrorsAPIService := openapi.NewGenericErrorsAPIService()
	GenericErrorsAPIController := openapi.NewGenericErrorsAPIController(GenericErrorsAPIService)

	GitHubSourcesAPIService := openapi.NewGitHubSourcesAPIService()
	GitHubSourcesAPIController := openapi.NewGitHubSourcesAPIController(GitHubSourcesAPIService)

	GroupsAPIService := openapi.NewGroupsAPIService()
	GroupsAPIController := openapi.NewGroupsAPIController(GroupsAPIService)

	RegistrySourcesAPIService := openapi.NewRegistrySourcesAPIService()
	RegistrySourcesAPIController := openapi.NewRegistrySourcesAPIController(RegistrySourcesAPIService)

	router := openapi.NewRouter(ArtifactsAPIController, ComposeFileAPIController, ContainerImagesAPIController, ContainerRepositoriesAPIController, ContainersAPIController, DeploymentsAPIController, DevicesAPIController, GenericErrorsAPIController, GitHubSourcesAPIController, GroupsAPIController, RegistrySourcesAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}