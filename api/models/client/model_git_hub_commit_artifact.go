/*
Miru API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GitHubCommitArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitHubCommitArtifact{}

// GitHubCommitArtifact struct for GitHubCommitArtifact
type GitHubCommitArtifact struct {
	Object string `json:"object"`
	Id string `json:"id"`
	Status string `json:"status"`
	Ready bool `json:"ready"`
	Failed bool `json:"failed"`
	InProgress bool `json:"in_progress"`
}

type _GitHubCommitArtifact GitHubCommitArtifact

// NewGitHubCommitArtifact instantiates a new GitHubCommitArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitHubCommitArtifact(object string, id string, status string, ready bool, failed bool, inProgress bool) *GitHubCommitArtifact {
	this := GitHubCommitArtifact{}
	this.Object = object
	this.Id = id
	this.Status = status
	this.Ready = ready
	this.Failed = failed
	this.InProgress = inProgress
	return &this
}

// NewGitHubCommitArtifactWithDefaults instantiates a new GitHubCommitArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubCommitArtifactWithDefaults() *GitHubCommitArtifact {
	this := GitHubCommitArtifact{}
	return &this
}

// GetObject returns the Object field value
func (o *GitHubCommitArtifact) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitArtifact) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GitHubCommitArtifact) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *GitHubCommitArtifact) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitArtifact) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GitHubCommitArtifact) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *GitHubCommitArtifact) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitArtifact) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GitHubCommitArtifact) SetStatus(v string) {
	o.Status = v
}

// GetReady returns the Ready field value
func (o *GitHubCommitArtifact) GetReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitArtifact) GetReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *GitHubCommitArtifact) SetReady(v bool) {
	o.Ready = v
}

// GetFailed returns the Failed field value
func (o *GitHubCommitArtifact) GetFailed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitArtifact) GetFailedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failed, true
}

// SetFailed sets field value
func (o *GitHubCommitArtifact) SetFailed(v bool) {
	o.Failed = v
}

// GetInProgress returns the InProgress field value
func (o *GitHubCommitArtifact) GetInProgress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InProgress
}

// GetInProgressOk returns a tuple with the InProgress field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitArtifact) GetInProgressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InProgress, true
}

// SetInProgress sets field value
func (o *GitHubCommitArtifact) SetInProgress(v bool) {
	o.InProgress = v
}

func (o GitHubCommitArtifact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitHubCommitArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["ready"] = o.Ready
	toSerialize["failed"] = o.Failed
	toSerialize["in_progress"] = o.InProgress
	return toSerialize, nil
}

func (o *GitHubCommitArtifact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"status",
		"ready",
		"failed",
		"in_progress",
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

	varGitHubCommitArtifact := _GitHubCommitArtifact{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitHubCommitArtifact)

	if err != nil {
		return err
	}

	*o = GitHubCommitArtifact(varGitHubCommitArtifact)

	return err
}

type NullableGitHubCommitArtifact struct {
	value *GitHubCommitArtifact
	isSet bool
}

func (v NullableGitHubCommitArtifact) Get() *GitHubCommitArtifact {
	return v.value
}

func (v *NullableGitHubCommitArtifact) Set(val *GitHubCommitArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableGitHubCommitArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableGitHubCommitArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitHubCommitArtifact(val *GitHubCommitArtifact) *NullableGitHubCommitArtifact {
	return &NullableGitHubCommitArtifact{value: val, isSet: true}
}

func (v NullableGitHubCommitArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitHubCommitArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

