/*
Miru API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GitHubSourceCommits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitHubSourceCommits{}

// GitHubSourceCommits struct for GitHubSourceCommits
type GitHubSourceCommits struct {
	Object *string `json:"object,omitempty"`
	Repository *GitHubRepository `json:"repository,omitempty"`
	Commits *GitHubCommitList `json:"commits,omitempty"`
}

// NewGitHubSourceCommits instantiates a new GitHubSourceCommits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitHubSourceCommits() *GitHubSourceCommits {
	this := GitHubSourceCommits{}
	return &this
}

// NewGitHubSourceCommitsWithDefaults instantiates a new GitHubSourceCommits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubSourceCommitsWithDefaults() *GitHubSourceCommits {
	this := GitHubSourceCommits{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *GitHubSourceCommits) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubSourceCommits) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *GitHubSourceCommits) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *GitHubSourceCommits) SetObject(v string) {
	o.Object = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *GitHubSourceCommits) GetRepository() GitHubRepository {
	if o == nil || IsNil(o.Repository) {
		var ret GitHubRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubSourceCommits) GetRepositoryOk() (*GitHubRepository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GitHubSourceCommits) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given GitHubRepository and assigns it to the Repository field.
func (o *GitHubSourceCommits) SetRepository(v GitHubRepository) {
	o.Repository = &v
}

// GetCommits returns the Commits field value if set, zero value otherwise.
func (o *GitHubSourceCommits) GetCommits() GitHubCommitList {
	if o == nil || IsNil(o.Commits) {
		var ret GitHubCommitList
		return ret
	}
	return *o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubSourceCommits) GetCommitsOk() (*GitHubCommitList, bool) {
	if o == nil || IsNil(o.Commits) {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *GitHubSourceCommits) HasCommits() bool {
	if o != nil && !IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given GitHubCommitList and assigns it to the Commits field.
func (o *GitHubSourceCommits) SetCommits(v GitHubCommitList) {
	o.Commits = &v
}

func (o GitHubSourceCommits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitHubSourceCommits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Commits) {
		toSerialize["commits"] = o.Commits
	}
	return toSerialize, nil
}

type NullableGitHubSourceCommits struct {
	value *GitHubSourceCommits
	isSet bool
}

func (v NullableGitHubSourceCommits) Get() *GitHubSourceCommits {
	return v.value
}

func (v *NullableGitHubSourceCommits) Set(val *GitHubSourceCommits) {
	v.value = val
	v.isSet = true
}

func (v NullableGitHubSourceCommits) IsSet() bool {
	return v.isSet
}

func (v *NullableGitHubSourceCommits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitHubSourceCommits(val *GitHubSourceCommits) *NullableGitHubSourceCommits {
	return &NullableGitHubSourceCommits{value: val, isSet: true}
}

func (v NullableGitHubSourceCommits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitHubSourceCommits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


