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

// checks if the DeploymentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentList{}

// DeploymentList struct for DeploymentList
type DeploymentList struct {
	Object string `json:"object"`
	Data []Deployment `json:"data"`
}

type _DeploymentList DeploymentList

// NewDeploymentList instantiates a new DeploymentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentList(object string, data []Deployment) *DeploymentList {
	this := DeploymentList{}
	this.Object = object
	this.Data = data
	return &this
}

// NewDeploymentListWithDefaults instantiates a new DeploymentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentListWithDefaults() *DeploymentList {
	this := DeploymentList{}
	return &this
}

// GetObject returns the Object field value
func (o *DeploymentList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *DeploymentList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *DeploymentList) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *DeploymentList) GetData() []Deployment {
	if o == nil {
		var ret []Deployment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeploymentList) GetDataOk() ([]Deployment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DeploymentList) SetData(v []Deployment) {
	o.Data = v
}

func (o DeploymentList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DeploymentList) UnmarshalJSON(data []byte) (err error) {
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

	varDeploymentList := _DeploymentList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeploymentList)

	if err != nil {
		return err
	}

	*o = DeploymentList(varDeploymentList)

	return err
}

type NullableDeploymentList struct {
	value *DeploymentList
	isSet bool
}

func (v NullableDeploymentList) Get() *DeploymentList {
	return v.value
}

func (v *NullableDeploymentList) Set(val *DeploymentList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentList(val *DeploymentList) *NullableDeploymentList {
	return &NullableDeploymentList{value: val, isSet: true}
}

func (v NullableDeploymentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


