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

// checks if the GroupDeviceContainerList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupDeviceContainerList{}

// GroupDeviceContainerList struct for GroupDeviceContainerList
type GroupDeviceContainerList struct {
	Object string `json:"object"`
	Data []GroupDeviceContainer `json:"data"`
}

type _GroupDeviceContainerList GroupDeviceContainerList

// NewGroupDeviceContainerList instantiates a new GroupDeviceContainerList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDeviceContainerList(object string, data []GroupDeviceContainer) *GroupDeviceContainerList {
	this := GroupDeviceContainerList{}
	this.Object = object
	this.Data = data
	return &this
}

// NewGroupDeviceContainerListWithDefaults instantiates a new GroupDeviceContainerList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDeviceContainerListWithDefaults() *GroupDeviceContainerList {
	this := GroupDeviceContainerList{}
	return &this
}

// GetObject returns the Object field value
func (o *GroupDeviceContainerList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GroupDeviceContainerList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GroupDeviceContainerList) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *GroupDeviceContainerList) GetData() []GroupDeviceContainer {
	if o == nil {
		var ret []GroupDeviceContainer
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GroupDeviceContainerList) GetDataOk() ([]GroupDeviceContainer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GroupDeviceContainerList) SetData(v []GroupDeviceContainer) {
	o.Data = v
}

func (o GroupDeviceContainerList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupDeviceContainerList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GroupDeviceContainerList) UnmarshalJSON(data []byte) (err error) {
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

	varGroupDeviceContainerList := _GroupDeviceContainerList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupDeviceContainerList)

	if err != nil {
		return err
	}

	*o = GroupDeviceContainerList(varGroupDeviceContainerList)

	return err
}

type NullableGroupDeviceContainerList struct {
	value *GroupDeviceContainerList
	isSet bool
}

func (v NullableGroupDeviceContainerList) Get() *GroupDeviceContainerList {
	return v.value
}

func (v *NullableGroupDeviceContainerList) Set(val *GroupDeviceContainerList) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDeviceContainerList) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDeviceContainerList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDeviceContainerList(val *GroupDeviceContainerList) *NullableGroupDeviceContainerList {
	return &NullableGroupDeviceContainerList{value: val, isSet: true}
}

func (v NullableGroupDeviceContainerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDeviceContainerList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


