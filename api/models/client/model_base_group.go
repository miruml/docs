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

// checks if the BaseGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseGroup{}

// BaseGroup struct for BaseGroup
type BaseGroup struct {
	Object string `json:"object"`
	Id string `json:"id"`
	Name string `json:"name"`
}

type _BaseGroup BaseGroup

// NewBaseGroup instantiates a new BaseGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseGroup(object string, id string, name string) *BaseGroup {
	this := BaseGroup{}
	this.Object = object
	this.Id = id
	this.Name = name
	return &this
}

// NewBaseGroupWithDefaults instantiates a new BaseGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseGroupWithDefaults() *BaseGroup {
	this := BaseGroup{}
	return &this
}

// GetObject returns the Object field value
func (o *BaseGroup) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *BaseGroup) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *BaseGroup) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *BaseGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BaseGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BaseGroup) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BaseGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseGroup) SetName(v string) {
	o.Name = v
}

func (o BaseGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *BaseGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"name",
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

	varBaseGroup := _BaseGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseGroup)

	if err != nil {
		return err
	}

	*o = BaseGroup(varBaseGroup)

	return err
}

type NullableBaseGroup struct {
	value *BaseGroup
	isSet bool
}

func (v NullableBaseGroup) Get() *BaseGroup {
	return v.value
}

func (v *NullableBaseGroup) Set(val *BaseGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseGroup(val *BaseGroup) *NullableBaseGroup {
	return &NullableBaseGroup{value: val, isSet: true}
}

func (v NullableBaseGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


