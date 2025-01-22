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

// checks if the BadRequestError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequestError{}

// BadRequestError struct for BadRequestError
type BadRequestError struct {
	Code string `json:"code"`
	Params map[string]interface{} `json:"params"`
	Message string `json:"message"`
}

type _BadRequestError BadRequestError

// NewBadRequestError instantiates a new BadRequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestError(code string, params map[string]interface{}, message string) *BadRequestError {
	this := BadRequestError{}
	this.Code = code
	this.Params = params
	this.Message = message
	return &this
}

// NewBadRequestErrorWithDefaults instantiates a new BadRequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestErrorWithDefaults() *BadRequestError {
	this := BadRequestError{}
	return &this
}

// GetCode returns the Code field value
func (o *BadRequestError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *BadRequestError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *BadRequestError) SetCode(v string) {
	o.Code = v
}

// GetParams returns the Params field value
func (o *BadRequestError) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *BadRequestError) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *BadRequestError) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetMessage returns the Message field value
func (o *BadRequestError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BadRequestError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BadRequestError) SetMessage(v string) {
	o.Message = v
}

func (o BadRequestError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequestError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["params"] = o.Params
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *BadRequestError) UnmarshalJSON(data []byte) (err error) {
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

	varBadRequestError := _BadRequestError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBadRequestError)

	if err != nil {
		return err
	}

	*o = BadRequestError(varBadRequestError)

	return err
}

type NullableBadRequestError struct {
	value *BadRequestError
	isSet bool
}

func (v NullableBadRequestError) Get() *BadRequestError {
	return v.value
}

func (v *NullableBadRequestError) Set(val *BadRequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestError(val *BadRequestError) *NullableBadRequestError {
	return &NullableBadRequestError{value: val, isSet: true}
}

func (v NullableBadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


