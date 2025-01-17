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

// checks if the RepositoryNotFound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryNotFound{}

// RepositoryNotFound struct for RepositoryNotFound
type RepositoryNotFound struct {
	Code string `json:"code"`
	Params RepositoryNotFoundParams `json:"params"`
	Message string `json:"message"`
}

type _RepositoryNotFound RepositoryNotFound

// NewRepositoryNotFound instantiates a new RepositoryNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryNotFound(code string, params RepositoryNotFoundParams, message string) *RepositoryNotFound {
	this := RepositoryNotFound{}
	this.Code = code
	this.Params = params
	this.Message = message
	return &this
}

// NewRepositoryNotFoundWithDefaults instantiates a new RepositoryNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryNotFoundWithDefaults() *RepositoryNotFound {
	this := RepositoryNotFound{}
	return &this
}

// GetCode returns the Code field value
func (o *RepositoryNotFound) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *RepositoryNotFound) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *RepositoryNotFound) SetCode(v string) {
	o.Code = v
}

// GetParams returns the Params field value
func (o *RepositoryNotFound) GetParams() RepositoryNotFoundParams {
	if o == nil {
		var ret RepositoryNotFoundParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *RepositoryNotFound) GetParamsOk() (*RepositoryNotFoundParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *RepositoryNotFound) SetParams(v RepositoryNotFoundParams) {
	o.Params = v
}

// GetMessage returns the Message field value
func (o *RepositoryNotFound) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RepositoryNotFound) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RepositoryNotFound) SetMessage(v string) {
	o.Message = v
}

func (o RepositoryNotFound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryNotFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["params"] = o.Params
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *RepositoryNotFound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"params",
		"message",
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

	varRepositoryNotFound := _RepositoryNotFound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRepositoryNotFound)

	if err != nil {
		return err
	}

	*o = RepositoryNotFound(varRepositoryNotFound)

	return err
}

type NullableRepositoryNotFound struct {
	value *RepositoryNotFound
	isSet bool
}

func (v NullableRepositoryNotFound) Get() *RepositoryNotFound {
	return v.value
}

func (v *NullableRepositoryNotFound) Set(val *RepositoryNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryNotFound(val *RepositoryNotFound) *NullableRepositoryNotFound {
	return &NullableRepositoryNotFound{value: val, isSet: true}
}

func (v NullableRepositoryNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

