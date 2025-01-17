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

// GroupsAPIService is a service that implements the logic for the GroupsAPIServicer
// This service should implement the business logic for every endpoint for the GroupsAPI API.
// Include any external packages or services that will be required by this service.
type GroupsAPIService struct {
}

// NewGroupsAPIService creates a default api service
func NewGroupsAPIService() *GroupsAPIService {
	return &GroupsAPIService{}
}

// GroupsGroupIdGet - 
func (s *GroupsAPIService) GroupsGroupIdGet(ctx context.Context, groupId string) (ImplResponse, error) {
	// TODO - update GroupsGroupIdGet with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Group{}) or use other options such as http.Ok ...
	// return Response(200, Group{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GroupsGroupIdGet method not implemented")
}