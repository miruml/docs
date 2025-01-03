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

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image struct for Image
type Image struct {
	Object string `json:"object"`
	Id string `json:"id"`
	RegistryUrl string `json:"registry_url"`
	AccountLogin string `json:"account_login"`
	Name string `json:"name"`
	Uri string `json:"uri"`
	Type string `json:"type"`
	Digest string `json:"digest"`
	Tag string `json:"tag"`
}

type _Image Image

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(object string, id string, registryUrl string, accountLogin string, name string, uri string, type_ string, digest string, tag string) *Image {
	this := Image{}
	this.Object = object
	this.Id = id
	this.RegistryUrl = registryUrl
	this.AccountLogin = accountLogin
	this.Name = name
	this.Uri = uri
	this.Type = type_
	this.Digest = digest
	this.Tag = tag
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetObject returns the Object field value
func (o *Image) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Image) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Image) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Image) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Image) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Image) SetId(v string) {
	o.Id = v
}

// GetRegistryUrl returns the RegistryUrl field value
func (o *Image) GetRegistryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value
// and a boolean to check if the value has been set.
func (o *Image) GetRegistryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryUrl, true
}

// SetRegistryUrl sets field value
func (o *Image) SetRegistryUrl(v string) {
	o.RegistryUrl = v
}

// GetAccountLogin returns the AccountLogin field value
func (o *Image) GetAccountLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountLogin
}

// GetAccountLoginOk returns a tuple with the AccountLogin field value
// and a boolean to check if the value has been set.
func (o *Image) GetAccountLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountLogin, true
}

// SetAccountLogin sets field value
func (o *Image) SetAccountLogin(v string) {
	o.AccountLogin = v
}

// GetName returns the Name field value
func (o *Image) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Image) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Image) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *Image) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Image) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Image) SetUri(v string) {
	o.Uri = v
}

// GetType returns the Type field value
func (o *Image) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Image) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Image) SetType(v string) {
	o.Type = v
}

// GetDigest returns the Digest field value
func (o *Image) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *Image) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *Image) SetDigest(v string) {
	o.Digest = v
}

// GetTag returns the Tag field value
func (o *Image) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *Image) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *Image) SetTag(v string) {
	o.Tag = v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["registry_url"] = o.RegistryUrl
	toSerialize["account_login"] = o.AccountLogin
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri
	toSerialize["type"] = o.Type
	toSerialize["digest"] = o.Digest
	toSerialize["tag"] = o.Tag
	return toSerialize, nil
}

func (o *Image) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"registry_url",
		"account_login",
		"name",
		"uri",
		"type",
		"digest",
		"tag",
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

	varImage := _Image{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImage)

	if err != nil {
		return err
	}

	*o = Image(varImage)

	return err
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


