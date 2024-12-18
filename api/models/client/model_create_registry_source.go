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

// checks if the CreateRegistrySource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRegistrySource{}

// CreateRegistrySource struct for CreateRegistrySource
type CreateRegistrySource struct {
	Name string `json:"name"`
	ComposeFile string `json:"compose_file"`
	ExtraRepositories []string `json:"extra_repositories"`
	Aarch64 bool `json:"aarch64"`
	X8664 bool `json:"x86_64"`
}

type _CreateRegistrySource CreateRegistrySource

// NewCreateRegistrySource instantiates a new CreateRegistrySource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRegistrySource(name string, composeFile string, extraRepositories []string, aarch64 bool, x8664 bool) *CreateRegistrySource {
	this := CreateRegistrySource{}
	this.Name = name
	this.ComposeFile = composeFile
	this.ExtraRepositories = extraRepositories
	this.Aarch64 = aarch64
	this.X8664 = x8664
	return &this
}

// NewCreateRegistrySourceWithDefaults instantiates a new CreateRegistrySource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRegistrySourceWithDefaults() *CreateRegistrySource {
	this := CreateRegistrySource{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRegistrySource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRegistrySource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRegistrySource) SetName(v string) {
	o.Name = v
}

// GetComposeFile returns the ComposeFile field value
func (o *CreateRegistrySource) GetComposeFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComposeFile
}

// GetComposeFileOk returns a tuple with the ComposeFile field value
// and a boolean to check if the value has been set.
func (o *CreateRegistrySource) GetComposeFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComposeFile, true
}

// SetComposeFile sets field value
func (o *CreateRegistrySource) SetComposeFile(v string) {
	o.ComposeFile = v
}

// GetExtraRepositories returns the ExtraRepositories field value
func (o *CreateRegistrySource) GetExtraRepositories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExtraRepositories
}

// GetExtraRepositoriesOk returns a tuple with the ExtraRepositories field value
// and a boolean to check if the value has been set.
func (o *CreateRegistrySource) GetExtraRepositoriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraRepositories, true
}

// SetExtraRepositories sets field value
func (o *CreateRegistrySource) SetExtraRepositories(v []string) {
	o.ExtraRepositories = v
}

// GetAarch64 returns the Aarch64 field value
func (o *CreateRegistrySource) GetAarch64() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Aarch64
}

// GetAarch64Ok returns a tuple with the Aarch64 field value
// and a boolean to check if the value has been set.
func (o *CreateRegistrySource) GetAarch64Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aarch64, true
}

// SetAarch64 sets field value
func (o *CreateRegistrySource) SetAarch64(v bool) {
	o.Aarch64 = v
}

// GetX8664 returns the X8664 field value
func (o *CreateRegistrySource) GetX8664() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.X8664
}

// GetX8664Ok returns a tuple with the X8664 field value
// and a boolean to check if the value has been set.
func (o *CreateRegistrySource) GetX8664Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X8664, true
}

// SetX8664 sets field value
func (o *CreateRegistrySource) SetX8664(v bool) {
	o.X8664 = v
}

func (o CreateRegistrySource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRegistrySource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["compose_file"] = o.ComposeFile
	toSerialize["extra_repositories"] = o.ExtraRepositories
	toSerialize["aarch64"] = o.Aarch64
	toSerialize["x86_64"] = o.X8664
	return toSerialize, nil
}

func (o *CreateRegistrySource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"compose_file",
		"extra_repositories",
		"aarch64",
		"x86_64",
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

	varCreateRegistrySource := _CreateRegistrySource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRegistrySource)

	if err != nil {
		return err
	}

	*o = CreateRegistrySource(varCreateRegistrySource)

	return err
}

type NullableCreateRegistrySource struct {
	value *CreateRegistrySource
	isSet bool
}

func (v NullableCreateRegistrySource) Get() *CreateRegistrySource {
	return v.value
}

func (v *NullableCreateRegistrySource) Set(val *CreateRegistrySource) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRegistrySource) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRegistrySource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRegistrySource(val *CreateRegistrySource) *NullableCreateRegistrySource {
	return &NullableCreateRegistrySource{value: val, isSet: true}
}

func (v NullableCreateRegistrySource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRegistrySource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


