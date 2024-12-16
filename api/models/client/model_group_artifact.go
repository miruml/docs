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
	Digest *string `json:"digest,omitempty"`
	Aarch64 bool `json:"aarch64"`
	X8664 bool `json:"x86_64"`
	CreatedAt time.Time `json:"created_at"`
	ReadyAt NullableTime `json:"ready_at"`
	FailedAt NullableTime `json:"failed_at"`
	CreatedBy *User `json:"created_by,omitempty"`
	Deployments *ArtifactDeploymentList `json:"deployments,omitempty"`
	Images ImageList `json:"images"`
	SourceId string `json:"source_id"`
	SourceType string `json:"source_type"`
	RegistrySource *RegistrySource `json:"registry_source,omitempty"`
	GithubSource *GitHubSource `json:"github_source,omitempty"`
	GithubSourceData *GitHubSourceData `json:"github_source_data,omitempty"`
	Staged bool `json:"staged"`
}

type _GroupArtifact GroupArtifact

// NewGroupArtifact instantiates a new GroupArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupArtifact(object string, id string, status ArtifactStatus, aarch64 bool, x8664 bool, createdAt time.Time, readyAt NullableTime, failedAt NullableTime, images ImageList, sourceId string, sourceType string, staged bool) *GroupArtifact {
	this := GroupArtifact{}
	this.Object = object
	this.Id = id
	this.Status = status
	this.Aarch64 = aarch64
	this.X8664 = x8664
	this.CreatedAt = createdAt
	this.ReadyAt = readyAt
	this.FailedAt = failedAt
	this.Images = images
	this.SourceId = sourceId
	this.SourceType = sourceType
	this.Staged = staged
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

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *GroupArtifact) GetDigest() string {
	if o == nil || IsNil(o.Digest) {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetDigestOk() (*string, bool) {
	if o == nil || IsNil(o.Digest) {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *GroupArtifact) HasDigest() bool {
	if o != nil && !IsNil(o.Digest) {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *GroupArtifact) SetDigest(v string) {
	o.Digest = &v
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

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *GroupArtifact) GetCreatedBy() User {
	if o == nil || IsNil(o.CreatedBy) {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetCreatedByOk() (*User, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GroupArtifact) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *GroupArtifact) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *GroupArtifact) GetDeployments() ArtifactDeploymentList {
	if o == nil || IsNil(o.Deployments) {
		var ret ArtifactDeploymentList
		return ret
	}
	return *o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetDeploymentsOk() (*ArtifactDeploymentList, bool) {
	if o == nil || IsNil(o.Deployments) {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *GroupArtifact) HasDeployments() bool {
	if o != nil && !IsNil(o.Deployments) {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given ArtifactDeploymentList and assigns it to the Deployments field.
func (o *GroupArtifact) SetDeployments(v ArtifactDeploymentList) {
	o.Deployments = &v
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

// GetRegistrySource returns the RegistrySource field value if set, zero value otherwise.
func (o *GroupArtifact) GetRegistrySource() RegistrySource {
	if o == nil || IsNil(o.RegistrySource) {
		var ret RegistrySource
		return ret
	}
	return *o.RegistrySource
}

// GetRegistrySourceOk returns a tuple with the RegistrySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetRegistrySourceOk() (*RegistrySource, bool) {
	if o == nil || IsNil(o.RegistrySource) {
		return nil, false
	}
	return o.RegistrySource, true
}

// HasRegistrySource returns a boolean if a field has been set.
func (o *GroupArtifact) HasRegistrySource() bool {
	if o != nil && !IsNil(o.RegistrySource) {
		return true
	}

	return false
}

// SetRegistrySource gets a reference to the given RegistrySource and assigns it to the RegistrySource field.
func (o *GroupArtifact) SetRegistrySource(v RegistrySource) {
	o.RegistrySource = &v
}

// GetGithubSource returns the GithubSource field value if set, zero value otherwise.
func (o *GroupArtifact) GetGithubSource() GitHubSource {
	if o == nil || IsNil(o.GithubSource) {
		var ret GitHubSource
		return ret
	}
	return *o.GithubSource
}

// GetGithubSourceOk returns a tuple with the GithubSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetGithubSourceOk() (*GitHubSource, bool) {
	if o == nil || IsNil(o.GithubSource) {
		return nil, false
	}
	return o.GithubSource, true
}

// HasGithubSource returns a boolean if a field has been set.
func (o *GroupArtifact) HasGithubSource() bool {
	if o != nil && !IsNil(o.GithubSource) {
		return true
	}

	return false
}

// SetGithubSource gets a reference to the given GitHubSource and assigns it to the GithubSource field.
func (o *GroupArtifact) SetGithubSource(v GitHubSource) {
	o.GithubSource = &v
}

// GetGithubSourceData returns the GithubSourceData field value if set, zero value otherwise.
func (o *GroupArtifact) GetGithubSourceData() GitHubSourceData {
	if o == nil || IsNil(o.GithubSourceData) {
		var ret GitHubSourceData
		return ret
	}
	return *o.GithubSourceData
}

// GetGithubSourceDataOk returns a tuple with the GithubSourceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupArtifact) GetGithubSourceDataOk() (*GitHubSourceData, bool) {
	if o == nil || IsNil(o.GithubSourceData) {
		return nil, false
	}
	return o.GithubSourceData, true
}

// HasGithubSourceData returns a boolean if a field has been set.
func (o *GroupArtifact) HasGithubSourceData() bool {
	if o != nil && !IsNil(o.GithubSourceData) {
		return true
	}

	return false
}

// SetGithubSourceData gets a reference to the given GitHubSourceData and assigns it to the GithubSourceData field.
func (o *GroupArtifact) SetGithubSourceData(v GitHubSourceData) {
	o.GithubSourceData = &v
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
	if !IsNil(o.Digest) {
		toSerialize["digest"] = o.Digest
	}
	toSerialize["aarch64"] = o.Aarch64
	toSerialize["x86_64"] = o.X8664
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["ready_at"] = o.ReadyAt.Get()
	toSerialize["failed_at"] = o.FailedAt.Get()
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Deployments) {
		toSerialize["deployments"] = o.Deployments
	}
	toSerialize["images"] = o.Images
	toSerialize["source_id"] = o.SourceId
	toSerialize["source_type"] = o.SourceType
	if !IsNil(o.RegistrySource) {
		toSerialize["registry_source"] = o.RegistrySource
	}
	if !IsNil(o.GithubSource) {
		toSerialize["github_source"] = o.GithubSource
	}
	if !IsNil(o.GithubSourceData) {
		toSerialize["github_source_data"] = o.GithubSourceData
	}
	toSerialize["staged"] = o.Staged
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
		"aarch64",
		"x86_64",
		"created_at",
		"ready_at",
		"failed_at",
		"images",
		"source_id",
		"source_type",
		"staged",
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

