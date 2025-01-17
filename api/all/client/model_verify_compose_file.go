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

// checks if the VerifyComposeFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyComposeFile{}

// VerifyComposeFile struct for VerifyComposeFile
type VerifyComposeFile struct {
	ComposeFile string `json:"compose_file"`
	Architecture string `json:"architecture"`
}

type _VerifyComposeFile VerifyComposeFile

// NewVerifyComposeFile instantiates a new VerifyComposeFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyComposeFile(composeFile string, architecture string) *VerifyComposeFile {
	this := VerifyComposeFile{}
	this.ComposeFile = composeFile
	this.Architecture = architecture
	return &this
}

// NewVerifyComposeFileWithDefaults instantiates a new VerifyComposeFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyComposeFileWithDefaults() *VerifyComposeFile {
	this := VerifyComposeFile{}
	return &this
}

// GetComposeFile returns the ComposeFile field value
func (o *VerifyComposeFile) GetComposeFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComposeFile
}

// GetComposeFileOk returns a tuple with the ComposeFile field value
// and a boolean to check if the value has been set.
func (o *VerifyComposeFile) GetComposeFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComposeFile, true
}

// SetComposeFile sets field value
func (o *VerifyComposeFile) SetComposeFile(v string) {
	o.ComposeFile = v
}

// GetArchitecture returns the Architecture field value
func (o *VerifyComposeFile) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *VerifyComposeFile) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *VerifyComposeFile) SetArchitecture(v string) {
	o.Architecture = v
}

func (o VerifyComposeFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyComposeFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["compose_file"] = o.ComposeFile
	toSerialize["architecture"] = o.Architecture
	return toSerialize, nil
}

func (o *VerifyComposeFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"compose_file",
		"architecture",
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

	varVerifyComposeFile := _VerifyComposeFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifyComposeFile)

	if err != nil {
		return err
	}

	*o = VerifyComposeFile(varVerifyComposeFile)

	return err
}

type NullableVerifyComposeFile struct {
	value *VerifyComposeFile
	isSet bool
}

func (v NullableVerifyComposeFile) Get() *VerifyComposeFile {
	return v.value
}

func (v *NullableVerifyComposeFile) Set(val *VerifyComposeFile) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyComposeFile) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyComposeFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyComposeFile(val *VerifyComposeFile) *NullableVerifyComposeFile {
	return &NullableVerifyComposeFile{value: val, isSet: true}
}

func (v NullableVerifyComposeFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyComposeFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


