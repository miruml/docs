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

// checks if the ContainerRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRepository{}

// ContainerRepository struct for ContainerRepository
type ContainerRepository struct {
	Object string `json:"object"`
	Id string `json:"id"`
	RegistryUrl string `json:"registry_url"`
	AccountLogin string `json:"account_login"`
	Name string `json:"name"`
	Uri string `json:"uri"`
	Type ContainerRepositoryType `json:"type"`
}

type _ContainerRepository ContainerRepository

// NewContainerRepository instantiates a new ContainerRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRepository(object string, id string, registryUrl string, accountLogin string, name string, uri string, type_ ContainerRepositoryType) *ContainerRepository {
	this := ContainerRepository{}
	this.Object = object
	this.Id = id
	this.RegistryUrl = registryUrl
	this.AccountLogin = accountLogin
	this.Name = name
	this.Uri = uri
	this.Type = type_
	return &this
}

// NewContainerRepositoryWithDefaults instantiates a new ContainerRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRepositoryWithDefaults() *ContainerRepository {
	this := ContainerRepository{}
	return &this
}

// GetObject returns the Object field value
func (o *ContainerRepository) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ContainerRepository) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *ContainerRepository) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContainerRepository) SetId(v string) {
	o.Id = v
}

// GetRegistryUrl returns the RegistryUrl field value
func (o *ContainerRepository) GetRegistryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetRegistryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryUrl, true
}

// SetRegistryUrl sets field value
func (o *ContainerRepository) SetRegistryUrl(v string) {
	o.RegistryUrl = v
}

// GetAccountLogin returns the AccountLogin field value
func (o *ContainerRepository) GetAccountLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountLogin
}

// GetAccountLoginOk returns a tuple with the AccountLogin field value
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetAccountLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountLogin, true
}

// SetAccountLogin sets field value
func (o *ContainerRepository) SetAccountLogin(v string) {
	o.AccountLogin = v
}

// GetName returns the Name field value
func (o *ContainerRepository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerRepository) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *ContainerRepository) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ContainerRepository) SetUri(v string) {
	o.Uri = v
}

// GetType returns the Type field value
func (o *ContainerRepository) GetType() ContainerRepositoryType {
	if o == nil {
		var ret ContainerRepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetTypeOk() (*ContainerRepositoryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContainerRepository) SetType(v ContainerRepositoryType) {
	o.Type = v
}

func (o ContainerRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["registry_url"] = o.RegistryUrl
	toSerialize["account_login"] = o.AccountLogin
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ContainerRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"registry_url",
		"account_login",
		"name",
		"uri",
		"type",
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

	varContainerRepository := _ContainerRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerRepository)

	if err != nil {
		return err
	}

	*o = ContainerRepository(varContainerRepository)

	return err
}

type NullableContainerRepository struct {
	value *ContainerRepository
	isSet bool
}

func (v NullableContainerRepository) Get() *ContainerRepository {
	return v.value
}

func (v *NullableContainerRepository) Set(val *ContainerRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRepository(val *ContainerRepository) *NullableContainerRepository {
	return &NullableContainerRepository{value: val, isSet: true}
}

func (v NullableContainerRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


