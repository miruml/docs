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

// checks if the GroupDeviceDeploymentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupDeviceDeploymentList{}

// GroupDeviceDeploymentList struct for GroupDeviceDeploymentList
type GroupDeviceDeploymentList struct {
	Object string `json:"object"`
	Data []GroupDeviceDeployment `json:"data"`
}

type _GroupDeviceDeploymentList GroupDeviceDeploymentList

// NewGroupDeviceDeploymentList instantiates a new GroupDeviceDeploymentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDeviceDeploymentList(object string, data []GroupDeviceDeployment) *GroupDeviceDeploymentList {
	this := GroupDeviceDeploymentList{}
	this.Object = object
	this.Data = data
	return &this
}

// NewGroupDeviceDeploymentListWithDefaults instantiates a new GroupDeviceDeploymentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDeviceDeploymentListWithDefaults() *GroupDeviceDeploymentList {
	this := GroupDeviceDeploymentList{}
	return &this
}

// GetObject returns the Object field value
func (o *GroupDeviceDeploymentList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GroupDeviceDeploymentList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GroupDeviceDeploymentList) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *GroupDeviceDeploymentList) GetData() []GroupDeviceDeployment {
	if o == nil {
		var ret []GroupDeviceDeployment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GroupDeviceDeploymentList) GetDataOk() ([]GroupDeviceDeployment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GroupDeviceDeploymentList) SetData(v []GroupDeviceDeployment) {
	o.Data = v
}

func (o GroupDeviceDeploymentList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupDeviceDeploymentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GroupDeviceDeploymentList) UnmarshalJSON(data []byte) (err error) {
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

	varGroupDeviceDeploymentList := _GroupDeviceDeploymentList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupDeviceDeploymentList)

	if err != nil {
		return err
	}

	*o = GroupDeviceDeploymentList(varGroupDeviceDeploymentList)

	return err
}

type NullableGroupDeviceDeploymentList struct {
	value *GroupDeviceDeploymentList
	isSet bool
}

func (v NullableGroupDeviceDeploymentList) Get() *GroupDeviceDeploymentList {
	return v.value
}

func (v *NullableGroupDeviceDeploymentList) Set(val *GroupDeviceDeploymentList) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDeviceDeploymentList) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDeviceDeploymentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDeviceDeploymentList(val *GroupDeviceDeploymentList) *NullableGroupDeviceDeploymentList {
	return &NullableGroupDeviceDeploymentList{value: val, isSet: true}
}

func (v NullableGroupDeviceDeploymentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDeviceDeploymentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


