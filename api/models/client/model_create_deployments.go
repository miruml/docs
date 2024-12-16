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

// checks if the CreateDeployments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeployments{}

// CreateDeployments struct for CreateDeployments
type CreateDeployments struct {
	ArtifactIds []string `json:"artifact_ids"`
	DeviceIds []string `json:"device_ids"`
}

type _CreateDeployments CreateDeployments

// NewCreateDeployments instantiates a new CreateDeployments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeployments(artifactIds []string, deviceIds []string) *CreateDeployments {
	this := CreateDeployments{}
	this.ArtifactIds = artifactIds
	this.DeviceIds = deviceIds
	return &this
}

// NewCreateDeploymentsWithDefaults instantiates a new CreateDeployments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeploymentsWithDefaults() *CreateDeployments {
	this := CreateDeployments{}
	return &this
}

// GetArtifactIds returns the ArtifactIds field value
func (o *CreateDeployments) GetArtifactIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ArtifactIds
}

// GetArtifactIdsOk returns a tuple with the ArtifactIds field value
// and a boolean to check if the value has been set.
func (o *CreateDeployments) GetArtifactIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArtifactIds, true
}

// SetArtifactIds sets field value
func (o *CreateDeployments) SetArtifactIds(v []string) {
	o.ArtifactIds = v
}

// GetDeviceIds returns the DeviceIds field value
func (o *CreateDeployments) GetDeviceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value
// and a boolean to check if the value has been set.
func (o *CreateDeployments) GetDeviceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// SetDeviceIds sets field value
func (o *CreateDeployments) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

func (o CreateDeployments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeployments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artifact_ids"] = o.ArtifactIds
	toSerialize["device_ids"] = o.DeviceIds
	return toSerialize, nil
}

func (o *CreateDeployments) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"artifact_ids",
		"device_ids",
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

	varCreateDeployments := _CreateDeployments{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDeployments)

	if err != nil {
		return err
	}

	*o = CreateDeployments(varCreateDeployments)

	return err
}

type NullableCreateDeployments struct {
	value *CreateDeployments
	isSet bool
}

func (v NullableCreateDeployments) Get() *CreateDeployments {
	return v.value
}

func (v *NullableCreateDeployments) Set(val *CreateDeployments) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeployments) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeployments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeployments(val *CreateDeployments) *NullableCreateDeployments {
	return &NullableCreateDeployments{value: val, isSet: true}
}

func (v NullableCreateDeployments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeployments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


