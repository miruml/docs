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
	"time"
)



type Device1 struct {

	Object string `json:"object"`

	Id string `json:"id"`

	Name string `json:"name"`

	Hardware string `json:"hardware"`

	OperatingSystem *string `json:"operating_system"`

	Architecture *string `json:"architecture"`

	Status string `json:"status"`

	LastReportedStatus string `json:"last_reported_status"`

	MiruVersion *string `json:"miru_version"`

	CreatedAt time.Time `json:"created_at"`

	SyncedAt time.Time `json:"synced_at"`

	Containers Device1AllOfContainers `json:"containers,omitempty"`

	Deployments Device1AllOfDeployments `json:"deployments,omitempty"`
}

// AssertDevice1Required checks if the required fields are not zero-ed
func AssertDevice1Required(obj Device1) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"name": obj.Name,
		"hardware": obj.Hardware,
		"operating_system": obj.OperatingSystem,
		"architecture": obj.Architecture,
		"status": obj.Status,
		"last_reported_status": obj.LastReportedStatus,
		"miru_version": obj.MiruVersion,
		"created_at": obj.CreatedAt,
		"synced_at": obj.SyncedAt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDevice1AllOfContainersRequired(obj.Containers); err != nil {
		return err
	}
	if err := AssertDevice1AllOfDeploymentsRequired(obj.Deployments); err != nil {
		return err
	}
	return nil
}

// AssertDevice1Constraints checks if the values respects the defined constraints
func AssertDevice1Constraints(obj Device1) error {
	if err := AssertDevice1AllOfContainersConstraints(obj.Containers); err != nil {
		return err
	}
	if err := AssertDevice1AllOfDeploymentsConstraints(obj.Deployments); err != nil {
		return err
	}
	return nil
}
