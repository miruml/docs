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



type ArtifactDeployment struct {

	Object string `json:"object"`

	Id string `json:"id"`

	Device Device `json:"device,omitempty"`

	Status string `json:"status"`

	ActivityStatus string `json:"activity_status"`

	ErrorStatus string `json:"error_status"`

	CreatedBy *User `json:"created_by"`

	CreatedAt time.Time `json:"created_at"`

	StartedAt *time.Time `json:"started_at"`

	FinishedAt *time.Time `json:"finished_at"`

	RemovedAt *time.Time `json:"removed_at"`

	OnDevice bool `json:"on_device"`
}

// AssertArtifactDeploymentRequired checks if the required fields are not zero-ed
func AssertArtifactDeploymentRequired(obj ArtifactDeployment) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"status": obj.Status,
		"activity_status": obj.ActivityStatus,
		"error_status": obj.ErrorStatus,
		"created_by": obj.CreatedBy,
		"created_at": obj.CreatedAt,
		"started_at": obj.StartedAt,
		"finished_at": obj.FinishedAt,
		"removed_at": obj.RemovedAt,
		"on_device": obj.OnDevice,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDeviceRequired(obj.Device); err != nil {
		return err
	}
	if obj.CreatedBy != nil {
		if err := AssertUserRequired(*obj.CreatedBy); err != nil {
			return err
		}
	}
	return nil
}

// AssertArtifactDeploymentConstraints checks if the values respects the defined constraints
func AssertArtifactDeploymentConstraints(obj ArtifactDeployment) error {
	if err := AssertDeviceConstraints(obj.Device); err != nil {
		return err
	}
    if obj.CreatedBy != nil {
     	if err := AssertUserConstraints(*obj.CreatedBy); err != nil {
     		return err
     	}
    }
	return nil
}
