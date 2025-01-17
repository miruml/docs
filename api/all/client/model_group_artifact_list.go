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

// checks if the GroupArtifactList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupArtifactList{}

// GroupArtifactList struct for GroupArtifactList
type GroupArtifactList struct {
	Object string `json:"object"`
	Data []GroupArtifact `json:"data"`
}

type _GroupArtifactList GroupArtifactList

// NewGroupArtifactList instantiates a new GroupArtifactList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupArtifactList(object string, data []GroupArtifact) *GroupArtifactList {
	this := GroupArtifactList{}
	this.Object = object
	this.Data = data
	return &this
}

// NewGroupArtifactListWithDefaults instantiates a new GroupArtifactList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupArtifactListWithDefaults() *GroupArtifactList {
	this := GroupArtifactList{}
	return &this
}

// GetObject returns the Object field value
func (o *GroupArtifactList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GroupArtifactList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GroupArtifactList) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *GroupArtifactList) GetData() []GroupArtifact {
	if o == nil {
		var ret []GroupArtifact
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GroupArtifactList) GetDataOk() ([]GroupArtifact, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GroupArtifactList) SetData(v []GroupArtifact) {
	o.Data = v
}

func (o GroupArtifactList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupArtifactList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GroupArtifactList) UnmarshalJSON(data []byte) (err error) {
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

	varGroupArtifactList := _GroupArtifactList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupArtifactList)

	if err != nil {
		return err
	}

	*o = GroupArtifactList(varGroupArtifactList)

	return err
}

type NullableGroupArtifactList struct {
	value *GroupArtifactList
	isSet bool
}

func (v NullableGroupArtifactList) Get() *GroupArtifactList {
	return v.value
}

func (v *NullableGroupArtifactList) Set(val *GroupArtifactList) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupArtifactList) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupArtifactList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupArtifactList(val *GroupArtifactList) *NullableGroupArtifactList {
	return &NullableGroupArtifactList{value: val, isSet: true}
}

func (v NullableGroupArtifactList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupArtifactList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


