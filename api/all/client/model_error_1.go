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

// checks if the Error1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error1{}

// Error1 struct for Error1
type Error1 struct {
	Error BadRequestError `json:"error"`
}

type _Error1 Error1

// NewError1 instantiates a new Error1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError1(error_ BadRequestError) *Error1 {
	this := Error1{}
	this.Error = error_
	return &this
}

// NewError1WithDefaults instantiates a new Error1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewError1WithDefaults() *Error1 {
	this := Error1{}
	return &this
}

// GetError returns the Error field value
func (o *Error1) GetError() BadRequestError {
	if o == nil {
		var ret BadRequestError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *Error1) GetErrorOk() (*BadRequestError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *Error1) SetError(v BadRequestError) {
	o.Error = v
}

func (o Error1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *Error1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varError1 := _Error1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varError1)

	if err != nil {
		return err
	}

	*o = Error1(varError1)

	return err
}

type NullableError1 struct {
	value *Error1
	isSet bool
}

func (v NullableError1) Get() *Error1 {
	return v.value
}

func (v *NullableError1) Set(val *Error1) {
	v.value = val
	v.isSet = true
}

func (v NullableError1) IsSet() bool {
	return v.isSet
}

func (v *NullableError1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError1(val *Error1) *NullableError1 {
	return &NullableError1{value: val, isSet: true}
}

func (v NullableError1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


