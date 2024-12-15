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

// checks if the GitHubRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitHubRepository{}

// GitHubRepository struct for GitHubRepository
type GitHubRepository struct {
	Id int64 `json:"id"`
	Object string `json:"object"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	HtmlUrl string `json:"html_url"`
}

type _GitHubRepository GitHubRepository

// NewGitHubRepository instantiates a new GitHubRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitHubRepository(id int64, object string, name string, fullName string, htmlUrl string) *GitHubRepository {
	this := GitHubRepository{}
	this.Id = id
	this.Object = object
	this.Name = name
	this.FullName = fullName
	this.HtmlUrl = htmlUrl
	return &this
}

// NewGitHubRepositoryWithDefaults instantiates a new GitHubRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubRepositoryWithDefaults() *GitHubRepository {
	this := GitHubRepository{}
	return &this
}

// GetId returns the Id field value
func (o *GitHubRepository) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GitHubRepository) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GitHubRepository) SetId(v int64) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *GitHubRepository) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GitHubRepository) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GitHubRepository) SetObject(v string) {
	o.Object = v
}

// GetName returns the Name field value
func (o *GitHubRepository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitHubRepository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitHubRepository) SetName(v string) {
	o.Name = v
}

// GetFullName returns the FullName field value
func (o *GitHubRepository) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *GitHubRepository) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *GitHubRepository) SetFullName(v string) {
	o.FullName = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *GitHubRepository) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *GitHubRepository) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *GitHubRepository) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

func (o GitHubRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitHubRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["name"] = o.Name
	toSerialize["full_name"] = o.FullName
	toSerialize["html_url"] = o.HtmlUrl
	return toSerialize, nil
}

func (o *GitHubRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"name",
		"full_name",
		"html_url",
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

	varGitHubRepository := _GitHubRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitHubRepository)

	if err != nil {
		return err
	}

	*o = GitHubRepository(varGitHubRepository)

	return err
}

type NullableGitHubRepository struct {
	value *GitHubRepository
	isSet bool
}

func (v NullableGitHubRepository) Get() *GitHubRepository {
	return v.value
}

func (v *NullableGitHubRepository) Set(val *GitHubRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableGitHubRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableGitHubRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitHubRepository(val *GitHubRepository) *NullableGitHubRepository {
	return &NullableGitHubRepository{value: val, isSet: true}
}

func (v NullableGitHubRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitHubRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


