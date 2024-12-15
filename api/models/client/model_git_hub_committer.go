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

// checks if the GitHubCommitter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitHubCommitter{}

// GitHubCommitter struct for GitHubCommitter
type GitHubCommitter struct {
	Id int64 `json:"id"`
	Object string `json:"object"`
	Login string `json:"login"`
	Type string `json:"type"`
	HtmlUrl string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
}

type _GitHubCommitter GitHubCommitter

// NewGitHubCommitter instantiates a new GitHubCommitter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitHubCommitter(id int64, object string, login string, type_ string, htmlUrl string, avatarUrl string) *GitHubCommitter {
	this := GitHubCommitter{}
	this.Id = id
	this.Object = object
	this.Login = login
	this.Type = type_
	this.HtmlUrl = htmlUrl
	this.AvatarUrl = avatarUrl
	return &this
}

// NewGitHubCommitterWithDefaults instantiates a new GitHubCommitter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubCommitterWithDefaults() *GitHubCommitter {
	this := GitHubCommitter{}
	return &this
}

// GetId returns the Id field value
func (o *GitHubCommitter) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitter) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GitHubCommitter) SetId(v int64) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *GitHubCommitter) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitter) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GitHubCommitter) SetObject(v string) {
	o.Object = v
}

// GetLogin returns the Login field value
func (o *GitHubCommitter) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitter) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *GitHubCommitter) SetLogin(v string) {
	o.Login = v
}

// GetType returns the Type field value
func (o *GitHubCommitter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GitHubCommitter) SetType(v string) {
	o.Type = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *GitHubCommitter) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitter) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *GitHubCommitter) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *GitHubCommitter) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *GitHubCommitter) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *GitHubCommitter) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

func (o GitHubCommitter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitHubCommitter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["login"] = o.Login
	toSerialize["type"] = o.Type
	toSerialize["html_url"] = o.HtmlUrl
	toSerialize["avatar_url"] = o.AvatarUrl
	return toSerialize, nil
}

func (o *GitHubCommitter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"login",
		"type",
		"html_url",
		"avatar_url",
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

	varGitHubCommitter := _GitHubCommitter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitHubCommitter)

	if err != nil {
		return err
	}

	*o = GitHubCommitter(varGitHubCommitter)

	return err
}

type NullableGitHubCommitter struct {
	value *GitHubCommitter
	isSet bool
}

func (v NullableGitHubCommitter) Get() *GitHubCommitter {
	return v.value
}

func (v *NullableGitHubCommitter) Set(val *GitHubCommitter) {
	v.value = val
	v.isSet = true
}

func (v NullableGitHubCommitter) IsSet() bool {
	return v.isSet
}

func (v *NullableGitHubCommitter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitHubCommitter(val *GitHubCommitter) *NullableGitHubCommitter {
	return &NullableGitHubCommitter{value: val, isSet: true}
}

func (v NullableGitHubCommitter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitHubCommitter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

