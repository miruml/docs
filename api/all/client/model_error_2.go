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

// checks if the Error2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error2{}

// Error2 struct for Error2
type Error2 struct {
	Error UnauthorizedError `json:"error"`
}

type _Error2 Error2

// NewError2 instantiates a new Error2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError2(error_ UnauthorizedError) *Error2 {
	this := Error2{}
	this.Error = error_
	return &this
}

// NewError2WithDefaults instantiates a new Error2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewError2WithDefaults() *Error2 {
	this := Error2{}
	return &this
}

// GetError returns the Error field value
func (o *Error2) GetError() UnauthorizedError {
	if o == nil {
		var ret UnauthorizedError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *Error2) GetErrorOk() (*UnauthorizedError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *Error2) SetError(v UnauthorizedError) {
	o.Error = v
}

func (o Error2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *Error2) UnmarshalJSON(data []byte) (err error) {
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

	varError2 := _Error2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varError2)

	if err != nil {
		return err
	}

	*o = Error2(varError2)

	return err
}

type NullableError2 struct {
	value *Error2
	isSet bool
}

func (v NullableError2) Get() *Error2 {
	return v.value
}

func (v *NullableError2) Set(val *Error2) {
	v.value = val
	v.isSet = true
}

func (v NullableError2) IsSet() bool {
	return v.isSet
}

func (v *NullableError2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError2(val *Error2) *NullableError2 {
	return &NullableError2{value: val, isSet: true}
}

func (v NullableError2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

