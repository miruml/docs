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

// checks if the GroupArtifactDeploymentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupArtifactDeploymentList{}

// GroupArtifactDeploymentList struct for GroupArtifactDeploymentList
type GroupArtifactDeploymentList struct {
	Object string `json:"object"`
	Data []GroupArtifactDeployment `json:"data"`
}

type _GroupArtifactDeploymentList GroupArtifactDeploymentList

// NewGroupArtifactDeploymentList instantiates a new GroupArtifactDeploymentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupArtifactDeploymentList(object string, data []GroupArtifactDeployment) *GroupArtifactDeploymentList {
	this := GroupArtifactDeploymentList{}
	this.Object = object
	this.Data = data
	return &this
}

// NewGroupArtifactDeploymentListWithDefaults instantiates a new GroupArtifactDeploymentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupArtifactDeploymentListWithDefaults() *GroupArtifactDeploymentList {
	this := GroupArtifactDeploymentList{}
	return &this
}

// GetObject returns the Object field value
func (o *GroupArtifactDeploymentList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GroupArtifactDeploymentList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GroupArtifactDeploymentList) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *GroupArtifactDeploymentList) GetData() []GroupArtifactDeployment {
	if o == nil {
		var ret []GroupArtifactDeployment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GroupArtifactDeploymentList) GetDataOk() ([]GroupArtifactDeployment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GroupArtifactDeploymentList) SetData(v []GroupArtifactDeployment) {
	o.Data = v
}

func (o GroupArtifactDeploymentList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupArtifactDeploymentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GroupArtifactDeploymentList) UnmarshalJSON(data []byte) (err error) {
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

	varGroupArtifactDeploymentList := _GroupArtifactDeploymentList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupArtifactDeploymentList)

	if err != nil {
		return err
	}

	*o = GroupArtifactDeploymentList(varGroupArtifactDeploymentList)

	return err
}

type NullableGroupArtifactDeploymentList struct {
	value *GroupArtifactDeploymentList
	isSet bool
}

func (v NullableGroupArtifactDeploymentList) Get() *GroupArtifactDeploymentList {
	return v.value
}

func (v *NullableGroupArtifactDeploymentList) Set(val *GroupArtifactDeploymentList) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupArtifactDeploymentList) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupArtifactDeploymentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupArtifactDeploymentList(val *GroupArtifactDeploymentList) *NullableGroupArtifactDeploymentList {
	return &NullableGroupArtifactDeploymentList{value: val, isSet: true}
}

func (v NullableGroupArtifactDeploymentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupArtifactDeploymentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


