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

// checks if the RegistrySourceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySourceList{}

// RegistrySourceList struct for RegistrySourceList
type RegistrySourceList struct {
	Object string `json:"object"`
	Data []RegistrySourceListItem `json:"data"`
}

type _RegistrySourceList RegistrySourceList

// NewRegistrySourceList instantiates a new RegistrySourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySourceList(object string, data []RegistrySourceListItem) *RegistrySourceList {
	this := RegistrySourceList{}
	this.Object = object
	this.Data = data
	return &this
}

// NewRegistrySourceListWithDefaults instantiates a new RegistrySourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySourceListWithDefaults() *RegistrySourceList {
	this := RegistrySourceList{}
	return &this
}

// GetObject returns the Object field value
func (o *RegistrySourceList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *RegistrySourceList) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *RegistrySourceList) GetData() []RegistrySourceListItem {
	if o == nil {
		var ret []RegistrySourceListItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceList) GetDataOk() ([]RegistrySourceListItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *RegistrySourceList) SetData(v []RegistrySourceListItem) {
	o.Data = v
}

func (o RegistrySourceList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySourceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RegistrySourceList) UnmarshalJSON(data []byte) (err error) {
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

	varRegistrySourceList := _RegistrySourceList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySourceList)

	if err != nil {
		return err
	}

	*o = RegistrySourceList(varRegistrySourceList)

	return err
}

type NullableRegistrySourceList struct {
	value *RegistrySourceList
	isSet bool
}

func (v NullableRegistrySourceList) Get() *RegistrySourceList {
	return v.value
}

func (v *NullableRegistrySourceList) Set(val *RegistrySourceList) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySourceList) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySourceList(val *RegistrySourceList) *NullableRegistrySourceList {
	return &NullableRegistrySourceList{value: val, isSet: true}
}

func (v NullableRegistrySourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


