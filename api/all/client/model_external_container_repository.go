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

// checks if the ExternalContainerRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalContainerRepository{}

// ExternalContainerRepository struct for ExternalContainerRepository
type ExternalContainerRepository struct {
	Object string `json:"object"`
	RegistryUrl string `json:"registry_url"`
	AccountLogin string `json:"account_login"`
	Name string `json:"name"`
	Uri string `json:"uri"`
	Type ContainerRepositoryType `json:"type"`
	Description NullableString `json:"description"`
	// The size of the repository in bytes
	Bytes NullableInt64 `json:"bytes"`
}

type _ExternalContainerRepository ExternalContainerRepository

// NewExternalContainerRepository instantiates a new ExternalContainerRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalContainerRepository(object string, registryUrl string, accountLogin string, name string, uri string, type_ ContainerRepositoryType, description NullableString, bytes NullableInt64) *ExternalContainerRepository {
	this := ExternalContainerRepository{}
	this.Object = object
	this.RegistryUrl = registryUrl
	this.AccountLogin = accountLogin
	this.Name = name
	this.Uri = uri
	this.Type = type_
	this.Description = description
	this.Bytes = bytes
	return &this
}

// NewExternalContainerRepositoryWithDefaults instantiates a new ExternalContainerRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalContainerRepositoryWithDefaults() *ExternalContainerRepository {
	this := ExternalContainerRepository{}
	return &this
}

// GetObject returns the Object field value
func (o *ExternalContainerRepository) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ExternalContainerRepository) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ExternalContainerRepository) SetObject(v string) {
	o.Object = v
}

// GetRegistryUrl returns the RegistryUrl field value
func (o *ExternalContainerRepository) GetRegistryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value
// and a boolean to check if the value has been set.
func (o *ExternalContainerRepository) GetRegistryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryUrl, true
}

// SetRegistryUrl sets field value
func (o *ExternalContainerRepository) SetRegistryUrl(v string) {
	o.RegistryUrl = v
}

// GetAccountLogin returns the AccountLogin field value
func (o *ExternalContainerRepository) GetAccountLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountLogin
}

// GetAccountLoginOk returns a tuple with the AccountLogin field value
// and a boolean to check if the value has been set.
func (o *ExternalContainerRepository) GetAccountLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountLogin, true
}

// SetAccountLogin sets field value
func (o *ExternalContainerRepository) SetAccountLogin(v string) {
	o.AccountLogin = v
}

// GetName returns the Name field value
func (o *ExternalContainerRepository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalContainerRepository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalContainerRepository) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *ExternalContainerRepository) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ExternalContainerRepository) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ExternalContainerRepository) SetUri(v string) {
	o.Uri = v
}

// GetType returns the Type field value
func (o *ExternalContainerRepository) GetType() ContainerRepositoryType {
	if o == nil {
		var ret ContainerRepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalContainerRepository) GetTypeOk() (*ContainerRepositoryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalContainerRepository) SetType(v ContainerRepositoryType) {
	o.Type = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalContainerRepository) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalContainerRepository) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ExternalContainerRepository) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetBytes returns the Bytes field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ExternalContainerRepository) GetBytes() int64 {
	if o == nil || o.Bytes.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Bytes.Get()
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalContainerRepository) GetBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bytes.Get(), o.Bytes.IsSet()
}

// SetBytes sets field value
func (o *ExternalContainerRepository) SetBytes(v int64) {
	o.Bytes.Set(&v)
}

func (o ExternalContainerRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalContainerRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["registry_url"] = o.RegistryUrl
	toSerialize["account_login"] = o.AccountLogin
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri
	toSerialize["type"] = o.Type
	toSerialize["description"] = o.Description.Get()
	toSerialize["bytes"] = o.Bytes.Get()
	return toSerialize, nil
}

func (o *ExternalContainerRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"registry_url",
		"account_login",
		"name",
		"uri",
		"type",
		"description",
		"bytes",
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

	varExternalContainerRepository := _ExternalContainerRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalContainerRepository)

	if err != nil {
		return err
	}

	*o = ExternalContainerRepository(varExternalContainerRepository)

	return err
}

type NullableExternalContainerRepository struct {
	value *ExternalContainerRepository
	isSet bool
}

func (v NullableExternalContainerRepository) Get() *ExternalContainerRepository {
	return v.value
}

func (v *NullableExternalContainerRepository) Set(val *ExternalContainerRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalContainerRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalContainerRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalContainerRepository(val *ExternalContainerRepository) *NullableExternalContainerRepository {
	return &NullableExternalContainerRepository{value: val, isSet: true}
}

func (v NullableExternalContainerRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalContainerRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

