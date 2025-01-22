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

// checks if the ContainerList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerList{}

// ContainerList struct for ContainerList
type ContainerList struct {
	Object string `json:"object"`
	Data []Container `json:"data"`
}

type _ContainerList ContainerList

// NewContainerList instantiates a new ContainerList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerList(object string, data []Container) *ContainerList {
	this := ContainerList{}
	this.Object = object
	this.Data = data
	return &this
}

// NewContainerListWithDefaults instantiates a new ContainerList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerListWithDefaults() *ContainerList {
	this := ContainerList{}
	return &this
}

// GetObject returns the Object field value
func (o *ContainerList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ContainerList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ContainerList) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *ContainerList) GetData() []Container {
	if o == nil {
		var ret []Container
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ContainerList) GetDataOk() ([]Container, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ContainerList) SetData(v []Container) {
	o.Data = v
}

func (o ContainerList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ContainerList) UnmarshalJSON(data []byte) (err error) {
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

	varContainerList := _ContainerList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerList)

	if err != nil {
		return err
	}

	*o = ContainerList(varContainerList)

	return err
}

type NullableContainerList struct {
	value *ContainerList
	isSet bool
}

func (v NullableContainerList) Get() *ContainerList {
	return v.value
}

func (v *NullableContainerList) Set(val *ContainerList) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerList) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerList(val *ContainerList) *NullableContainerList {
	return &NullableContainerList{value: val, isSet: true}
}

func (v NullableContainerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


