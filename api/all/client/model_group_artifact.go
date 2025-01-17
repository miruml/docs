/*
Miru API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the GroupArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupArtifact{}

// GroupArtifact struct for GroupArtifact
type GroupArtifact struct {
	Object string `json:"object"`
	Id string `json:"id"`
	Status ArtifactStatus `json:"status"`
	Digest string `json:"digest"`
	Aarch64 bool `json:"aarch64"`
	X8664 bool `json:"x86_64"`
	CreatedAt time.Time `json:"created_at"`
	ReadyAt NullableTime `json:"ready_at"`
	FailedAt NullableTime `json:"failed_at"`
	SourceId string `json:"source_id"`
	SourceType string `json:"source_type"`
	CreatedBy NullableUser `json:"created_by"`
	RegistrySource NullableRegistrySource `json:"registry_source"`
	GithubSource NullableGitHubSource `json:"github_source"`
	GithubSourceData NullableGitHubSourceData `json:"github_source_data"`
	Images ImageList `json:"images"`
	Staged bool `json:"staged"`
	Deployments GroupArtifactDeploymentList `json:"deployments"`
}

type _GroupArtifact GroupArtifact

// NewGroupArtifact instantiates a new GroupArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupArtifact(object string, id string, status ArtifactStatus, digest string, aarch64 bool, x8664 bool, createdAt time.Time, readyAt NullableTime, failedAt NullableTime, sourceId string, sourceType string, createdBy NullableUser, registrySource NullableRegistrySource, githubSource NullableGitHubSource, githubSourceData NullableGitHubSourceData, images ImageList, staged bool, deployments GroupArtifactDeploymentList) *GroupArtifact {
	this := GroupArtifact{}
	this.Object = object
	this.Id = id
	this.Status = status
	this.Digest = digest
	this.Aarch64 = aarch64
	this.X8664 = x8664
	this.CreatedAt = createdAt
	this.ReadyAt = readyAt
	this.FailedAt = failedAt
	this.SourceId = sourceId
	this.SourceType = sourceType
	this.CreatedBy = createdBy
	this.RegistrySource = registrySource
	this.GithubSource = githubSource
	this.GithubSourceData = githubSourceData
	this.Images = images
	this.Staged = staged
	this.Deployments = deployments
	return &this
}

// NewGroupArtifactWithDefaults instantiates a new GroupArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupArtifactWithDefaults() *GroupArtifact {
	this := GroupArtifact{}
	return &this
}

// GetObject returns the Object field value
func (o *GroupArtifact) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GroupArtifact) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *GroupArtifact) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupArtifact) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *GroupArtifact) GetStatus() ArtifactStatus {
	if o == nil {
		var ret ArtifactStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetStatusOk() (*ArtifactStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GroupArtifact) SetStatus(v ArtifactStatus) {
	o.Status = v
}

// GetDigest returns the Digest field value
func (o *GroupArtifact) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *GroupArtifact) SetDigest(v string) {
	o.Digest = v
}

// GetAarch64 returns the Aarch64 field value
func (o *GroupArtifact) GetAarch64() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Aarch64
}

// GetAarch64Ok returns a tuple with the Aarch64 field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetAarch64Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aarch64, true
}

// SetAarch64 sets field value
func (o *GroupArtifact) SetAarch64(v bool) {
	o.Aarch64 = v
}

// GetX8664 returns the X8664 field value
func (o *GroupArtifact) GetX8664() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.X8664
}

// GetX8664Ok returns a tuple with the X8664 field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetX8664Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X8664, true
}

// SetX8664 sets field value
func (o *GroupArtifact) SetX8664(v bool) {
	o.X8664 = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GroupArtifact) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GroupArtifact) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetReadyAt returns the ReadyAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *GroupArtifact) GetReadyAt() time.Time {
	if o == nil || o.ReadyAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ReadyAt.Get()
}

// GetReadyAtOk returns a tuple with the ReadyAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupArtifact) GetReadyAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadyAt.Get(), o.ReadyAt.IsSet()
}

// SetReadyAt sets field value
func (o *GroupArtifact) SetReadyAt(v time.Time) {
	o.ReadyAt.Set(&v)
}

// GetFailedAt returns the FailedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *GroupArtifact) GetFailedAt() time.Time {
	if o == nil || o.FailedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.FailedAt.Get()
}

// GetFailedAtOk returns a tuple with the FailedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupArtifact) GetFailedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedAt.Get(), o.FailedAt.IsSet()
}

// SetFailedAt sets field value
func (o *GroupArtifact) SetFailedAt(v time.Time) {
	o.FailedAt.Set(&v)
}

// GetSourceId returns the SourceId field value
func (o *GroupArtifact) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *GroupArtifact) SetSourceId(v string) {
	o.SourceId = v
}

// GetSourceType returns the SourceType field value
func (o *GroupArtifact) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *GroupArtifact) SetSourceType(v string) {
	o.SourceType = v
}

// GetCreatedBy returns the CreatedBy field value
// If the value is explicit nil, the zero value for User will be returned
func (o *GroupArtifact) GetCreatedBy() User {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret User
		return ret
	}

	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupArtifact) GetCreatedByOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// SetCreatedBy sets field value
func (o *GroupArtifact) SetCreatedBy(v User) {
	o.CreatedBy.Set(&v)
}

// GetRegistrySource returns the RegistrySource field value
// If the value is explicit nil, the zero value for RegistrySource will be returned
func (o *GroupArtifact) GetRegistrySource() RegistrySource {
	if o == nil || o.RegistrySource.Get() == nil {
		var ret RegistrySource
		return ret
	}

	return *o.RegistrySource.Get()
}

// GetRegistrySourceOk returns a tuple with the RegistrySource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupArtifact) GetRegistrySourceOk() (*RegistrySource, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistrySource.Get(), o.RegistrySource.IsSet()
}

// SetRegistrySource sets field value
func (o *GroupArtifact) SetRegistrySource(v RegistrySource) {
	o.RegistrySource.Set(&v)
}

// GetGithubSource returns the GithubSource field value
// If the value is explicit nil, the zero value for GitHubSource will be returned
func (o *GroupArtifact) GetGithubSource() GitHubSource {
	if o == nil || o.GithubSource.Get() == nil {
		var ret GitHubSource
		return ret
	}

	return *o.GithubSource.Get()
}

// GetGithubSourceOk returns a tuple with the GithubSource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupArtifact) GetGithubSourceOk() (*GitHubSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.GithubSource.Get(), o.GithubSource.IsSet()
}

// SetGithubSource sets field value
func (o *GroupArtifact) SetGithubSource(v GitHubSource) {
	o.GithubSource.Set(&v)
}

// GetGithubSourceData returns the GithubSourceData field value
// If the value is explicit nil, the zero value for GitHubSourceData will be returned
func (o *GroupArtifact) GetGithubSourceData() GitHubSourceData {
	if o == nil || o.GithubSourceData.Get() == nil {
		var ret GitHubSourceData
		return ret
	}

	return *o.GithubSourceData.Get()
}

// GetGithubSourceDataOk returns a tuple with the GithubSourceData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupArtifact) GetGithubSourceDataOk() (*GitHubSourceData, bool) {
	if o == nil {
		return nil, false
	}
	return o.GithubSourceData.Get(), o.GithubSourceData.IsSet()
}

// SetGithubSourceData sets field value
func (o *GroupArtifact) SetGithubSourceData(v GitHubSourceData) {
	o.GithubSourceData.Set(&v)
}

// GetImages returns the Images field value
func (o *GroupArtifact) GetImages() ImageList {
	if o == nil {
		var ret ImageList
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetImagesOk() (*ImageList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Images, true
}

// SetImages sets field value
func (o *GroupArtifact) SetImages(v ImageList) {
	o.Images = v
}

// GetStaged returns the Staged field value
func (o *GroupArtifact) GetStaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Staged
}

// GetStagedOk returns a tuple with the Staged field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetStagedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Staged, true
}

// SetStaged sets field value
func (o *GroupArtifact) SetStaged(v bool) {
	o.Staged = v
}

// GetDeployments returns the Deployments field value
func (o *GroupArtifact) GetDeployments() GroupArtifactDeploymentList {
	if o == nil {
		var ret GroupArtifactDeploymentList
		return ret
	}

	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetDeploymentsOk() (*GroupArtifactDeploymentList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deployments, true
}

// SetDeployments sets field value
func (o *GroupArtifact) SetDeployments(v GroupArtifactDeploymentList) {
	o.Deployments = v
}

func (o GroupArtifact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["digest"] = o.Digest
	toSerialize["aarch64"] = o.Aarch64
	toSerialize["x86_64"] = o.X8664
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["ready_at"] = o.ReadyAt.Get()
	toSerialize["failed_at"] = o.FailedAt.Get()
	toSerialize["source_id"] = o.SourceId
	toSerialize["source_type"] = o.SourceType
	toSerialize["created_by"] = o.CreatedBy.Get()
	toSerialize["registry_source"] = o.RegistrySource.Get()
	toSerialize["github_source"] = o.GithubSource.Get()
	toSerialize["github_source_data"] = o.GithubSourceData.Get()
	toSerialize["images"] = o.Images
	toSerialize["staged"] = o.Staged
	toSerialize["deployments"] = o.Deployments
	return toSerialize, nil
}

func (o *GroupArtifact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"status",
		"digest",
		"aarch64",
		"x86_64",
		"created_at",
		"ready_at",
		"failed_at",
		"source_id",
		"source_type",
		"created_by",
		"registry_source",
		"github_source",
		"github_source_data",
		"images",
		"staged",
		"deployments",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGroupArtifact := _GroupArtifact{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupArtifact)

	if err != nil {
		return err
	}

	*o = GroupArtifact(varGroupArtifact)

	return err
}

type NullableGroupArtifact struct {
	value *GroupArtifact
	isSet bool
}

func (v NullableGroupArtifact) Get() *GroupArtifact {
	return v.value
}

func (v *NullableGroupArtifact) Set(val *GroupArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupArtifact(val *GroupArtifact) *NullableGroupArtifact {
	return &NullableGroupArtifact{value: val, isSet: true}
}

func (v NullableGroupArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

