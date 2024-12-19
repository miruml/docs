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

// checks if the VerifiedComposeFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifiedComposeFileResponse{}

// VerifiedComposeFileResponse struct for VerifiedComposeFileResponse
type VerifiedComposeFileResponse struct {
	Object string `json:"object"`
	Content string `json:"content"`
	IsValid bool `json:"is_valid"`
	IsSchemaValid bool `json:"is_schema_valid"`
	Images []ComposeFileImageList `json:"images"`
}

type _VerifiedComposeFileResponse VerifiedComposeFileResponse

// NewVerifiedComposeFileResponse instantiates a new VerifiedComposeFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifiedComposeFileResponse(object string, content string, isValid bool, isSchemaValid bool, images []ComposeFileImageList) *VerifiedComposeFileResponse {
	this := VerifiedComposeFileResponse{}
	this.Object = object
	this.Content = content
	this.IsValid = isValid
	this.IsSchemaValid = isSchemaValid
	this.Images = images
	return &this
}

// NewVerifiedComposeFileResponseWithDefaults instantiates a new VerifiedComposeFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifiedComposeFileResponseWithDefaults() *VerifiedComposeFileResponse {
	this := VerifiedComposeFileResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *VerifiedComposeFileResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *VerifiedComposeFileResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *VerifiedComposeFileResponse) SetObject(v string) {
	o.Object = v
}

// GetContent returns the Content field value
func (o *VerifiedComposeFileResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *VerifiedComposeFileResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *VerifiedComposeFileResponse) SetContent(v string) {
	o.Content = v
}

// GetIsValid returns the IsValid field value
func (o *VerifiedComposeFileResponse) GetIsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *VerifiedComposeFileResponse) GetIsValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value
func (o *VerifiedComposeFileResponse) SetIsValid(v bool) {
	o.IsValid = v
}

// GetIsSchemaValid returns the IsSchemaValid field value
func (o *VerifiedComposeFileResponse) GetIsSchemaValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSchemaValid
}

// GetIsSchemaValidOk returns a tuple with the IsSchemaValid field value
// and a boolean to check if the value has been set.
func (o *VerifiedComposeFileResponse) GetIsSchemaValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSchemaValid, true
}

// SetIsSchemaValid sets field value
func (o *VerifiedComposeFileResponse) SetIsSchemaValid(v bool) {
	o.IsSchemaValid = v
}

// GetImages returns the Images field value
func (o *VerifiedComposeFileResponse) GetImages() []ComposeFileImageList {
	if o == nil {
		var ret []ComposeFileImageList
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *VerifiedComposeFileResponse) GetImagesOk() ([]ComposeFileImageList, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *VerifiedComposeFileResponse) SetImages(v []ComposeFileImageList) {
	o.Images = v
}

func (o VerifiedComposeFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifiedComposeFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["content"] = o.Content
	toSerialize["is_valid"] = o.IsValid
	toSerialize["is_schema_valid"] = o.IsSchemaValid
	toSerialize["images"] = o.Images
	return toSerialize, nil
}

func (o *VerifiedComposeFileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"content",
		"is_valid",
		"is_schema_valid",
		"images",
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

	varVerifiedComposeFileResponse := _VerifiedComposeFileResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifiedComposeFileResponse)

	if err != nil {
		return err
	}

	*o = VerifiedComposeFileResponse(varVerifiedComposeFileResponse)

	return err
}

type NullableVerifiedComposeFileResponse struct {
	value *VerifiedComposeFileResponse
	isSet bool
}

func (v NullableVerifiedComposeFileResponse) Get() *VerifiedComposeFileResponse {
	return v.value
}

func (v *NullableVerifiedComposeFileResponse) Set(val *VerifiedComposeFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiedComposeFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiedComposeFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiedComposeFileResponse(val *VerifiedComposeFileResponse) *NullableVerifiedComposeFileResponse {
	return &NullableVerifiedComposeFileResponse{value: val, isSet: true}
}

func (v NullableVerifiedComposeFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiedComposeFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

