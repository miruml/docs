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



type Container struct {

	Object string `json:"object"`

	Id string `json:"id"`

	DeviceId string `json:"device_id"`

	ImageId *string `json:"image_id"`

	ImageName *string `json:"image_name,omitempty"`

	ImageTag *string `json:"image_tag"`

	ImageDigest *string `json:"image_digest"`

	DockerContainerId string `json:"docker_container_id"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

	StartedAt time.Time `json:"started_at"`

	FinishedAt *time.Time `json:"finished_at"`

	Status string `json:"status"`

	TargetStatus string `json:"target_status"`

	ExitCode int32 `json:"exit_code"`

	Error *string `json:"error"`
}

// AssertContainerRequired checks if the required fields are not zero-ed
func AssertContainerRequired(obj Container) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"device_id": obj.DeviceId,
		"image_id": obj.ImageId,
		"image_tag": obj.ImageTag,
		"image_digest": obj.ImageDigest,
		"docker_container_id": obj.DockerContainerId,
		"created_at": obj.CreatedAt,
		"updated_at": obj.UpdatedAt,
		"started_at": obj.StartedAt,
		"finished_at": obj.FinishedAt,
		"status": obj.Status,
		"target_status": obj.TargetStatus,
		"exit_code": obj.ExitCode,
		"error": obj.Error,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertContainerConstraints checks if the values respects the defined constraints
func AssertContainerConstraints(obj Container) error {
	return nil
}