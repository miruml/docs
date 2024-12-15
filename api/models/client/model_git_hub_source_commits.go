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

// checks if the GitHubSourceCommits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitHubSourceCommits{}

// GitHubSourceCommits struct for GitHubSourceCommits
type GitHubSourceCommits struct {
	Object string `json:"object"`
	Repository GitHubRepository `json:"repository"`
	Branch string `json:"branch"`
	Commits GitHubCommitList `json:"commits"`
}

type _GitHubSourceCommits GitHubSourceCommits

// NewGitHubSourceCommits instantiates a new GitHubSourceCommits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitHubSourceCommits(object string, repository GitHubRepository, branch string, commits GitHubCommitList) *GitHubSourceCommits {
	this := GitHubSourceCommits{}
	this.Object = object
	this.Repository = repository
	this.Branch = branch
	this.Commits = commits
	return &this
}

// NewGitHubSourceCommitsWithDefaults instantiates a new GitHubSourceCommits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubSourceCommitsWithDefaults() *GitHubSourceCommits {
	this := GitHubSourceCommits{}
	return &this
}

// GetObject returns the Object field value
func (o *GitHubSourceCommits) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GitHubSourceCommits) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GitHubSourceCommits) SetObject(v string) {
	o.Object = v
}

// GetRepository returns the Repository field value
func (o *GitHubSourceCommits) GetRepository() GitHubRepository {
	if o == nil {
		var ret GitHubRepository
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *GitHubSourceCommits) GetRepositoryOk() (*GitHubRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *GitHubSourceCommits) SetRepository(v GitHubRepository) {
	o.Repository = v
}

// GetBranch returns the Branch field value
func (o *GitHubSourceCommits) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *GitHubSourceCommits) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *GitHubSourceCommits) SetBranch(v string) {
	o.Branch = v
}

// GetCommits returns the Commits field value
func (o *GitHubSourceCommits) GetCommits() GitHubCommitList {
	if o == nil {
		var ret GitHubCommitList
		return ret
	}

	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value
// and a boolean to check if the value has been set.
func (o *GitHubSourceCommits) GetCommitsOk() (*GitHubCommitList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commits, true
}

// SetCommits sets field value
func (o *GitHubSourceCommits) SetCommits(v GitHubCommitList) {
	o.Commits = v
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
	toSerialize["object"] = o.Object
	toSerialize["repository"] = o.Repository
	toSerialize["branch"] = o.Branch
	toSerialize["commits"] = o.Commits
	return toSerialize, nil
}

func (o *GitHubSourceCommits) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"repository",
		"branch",
		"commits",
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

	varGitHubSourceCommits := _GitHubSourceCommits{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitHubSourceCommits)

	if err != nil {
		return err
	}

	*o = GitHubSourceCommits(varGitHubSourceCommits)

	return err
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

