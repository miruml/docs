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

// checks if the ContainerRepositoryListWithIsExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRepositoryListWithIsExtra{}

// ContainerRepositoryListWithIsExtra struct for ContainerRepositoryListWithIsExtra
type ContainerRepositoryListWithIsExtra struct {
	Object string `json:"object"`
	Data []ContainerRepositoryWithIsExtra `json:"data"`
}

type _ContainerRepositoryListWithIsExtra ContainerRepositoryListWithIsExtra

// NewContainerRepositoryListWithIsExtra instantiates a new ContainerRepositoryListWithIsExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRepositoryListWithIsExtra(object string, data []ContainerRepositoryWithIsExtra) *ContainerRepositoryListWithIsExtra {
	this := ContainerRepositoryListWithIsExtra{}
	this.Object = object
	this.Data = data
	return &this
}

// NewContainerRepositoryListWithIsExtraWithDefaults instantiates a new ContainerRepositoryListWithIsExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRepositoryListWithIsExtraWithDefaults() *ContainerRepositoryListWithIsExtra {
	this := ContainerRepositoryListWithIsExtra{}
	return &this
}

// GetObject returns the Object field value
func (o *ContainerRepositoryListWithIsExtra) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ContainerRepositoryListWithIsExtra) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ContainerRepositoryListWithIsExtra) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *ContainerRepositoryListWithIsExtra) GetData() []ContainerRepositoryWithIsExtra {
	if o == nil {
		var ret []ContainerRepositoryWithIsExtra
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ContainerRepositoryListWithIsExtra) GetDataOk() ([]ContainerRepositoryWithIsExtra, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ContainerRepositoryListWithIsExtra) SetData(v []ContainerRepositoryWithIsExtra) {
	o.Data = v
}

func (o ContainerRepositoryListWithIsExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRepositoryListWithIsExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ContainerRepositoryListWithIsExtra) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"data",
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

	varContainerRepositoryListWithIsExtra := _ContainerRepositoryListWithIsExtra{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerRepositoryListWithIsExtra)

	if err != nil {
		return err
	}

	*o = ContainerRepositoryListWithIsExtra(varContainerRepositoryListWithIsExtra)

	return err
}

type NullableContainerRepositoryListWithIsExtra struct {
	value *ContainerRepositoryListWithIsExtra
	isSet bool
}

func (v NullableContainerRepositoryListWithIsExtra) Get() *ContainerRepositoryListWithIsExtra {
	return v.value
}

func (v *NullableContainerRepositoryListWithIsExtra) Set(val *ContainerRepositoryListWithIsExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRepositoryListWithIsExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRepositoryListWithIsExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRepositoryListWithIsExtra(val *ContainerRepositoryListWithIsExtra) *NullableContainerRepositoryListWithIsExtra {
	return &NullableContainerRepositoryListWithIsExtra{value: val, isSet: true}
}

func (v NullableContainerRepositoryListWithIsExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRepositoryListWithIsExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


