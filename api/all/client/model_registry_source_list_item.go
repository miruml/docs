/*
Miru API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the RegistrySourceListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySourceListItem{}

// RegistrySourceListItem struct for RegistrySourceListItem
type RegistrySourceListItem struct {
	Id string `json:"id"`
	Object string `json:"object"`
	Name string `json:"name"`
	Repositories RepositoryList `json:"repositories"`
	Aarch64 bool `json:"aarch64"`
	X8664 bool `json:"x86_64"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _RegistrySourceListItem RegistrySourceListItem

// NewRegistrySourceListItem instantiates a new RegistrySourceListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySourceListItem(id string, object string, name string, repositories RepositoryList, aarch64 bool, x8664 bool, createdAt time.Time, updatedAt time.Time) *RegistrySourceListItem {
	this := RegistrySourceListItem{}
	this.Id = id
	this.Object = object
	this.Name = name
	this.Repositories = repositories
	this.Aarch64 = aarch64
	this.X8664 = x8664
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewRegistrySourceListItemWithDefaults instantiates a new RegistrySourceListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySourceListItemWithDefaults() *RegistrySourceListItem {
	this := RegistrySourceListItem{}
	return &this
}

// GetId returns the Id field value
func (o *RegistrySourceListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegistrySourceListItem) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *RegistrySourceListItem) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceListItem) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *RegistrySourceListItem) SetObject(v string) {
	o.Object = v
}

// GetName returns the Name field value
func (o *RegistrySourceListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RegistrySourceListItem) SetName(v string) {
	o.Name = v
}

// GetRepositories returns the Repositories field value
func (o *RegistrySourceListItem) GetRepositories() RepositoryList {
	if o == nil {
		var ret RepositoryList
		return ret
	}

	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceListItem) GetRepositoriesOk() (*RepositoryList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repositories, true
}

// SetRepositories sets field value
func (o *RegistrySourceListItem) SetRepositories(v RepositoryList) {
	o.Repositories = v
}

// GetAarch64 returns the Aarch64 field value
func (o *RegistrySourceListItem) GetAarch64() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Aarch64
}

// GetAarch64Ok returns a tuple with the Aarch64 field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceListItem) GetAarch64Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aarch64, true
}

// SetAarch64 sets field value
func (o *RegistrySourceListItem) SetAarch64(v bool) {
	o.Aarch64 = v
}

// GetX8664 returns the X8664 field value
func (o *RegistrySourceListItem) GetX8664() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.X8664
}

// GetX8664Ok returns a tuple with the X8664 field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceListItem) GetX8664Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X8664, true
}

// SetX8664 sets field value
func (o *RegistrySourceListItem) SetX8664(v bool) {
	o.X8664 = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RegistrySourceListItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RegistrySourceListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RegistrySourceListItem) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RegistrySourceListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RegistrySourceListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o RegistrySourceListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySourceListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["name"] = o.Name
	toSerialize["repositories"] = o.Repositories
	toSerialize["aarch64"] = o.Aarch64
	toSerialize["x86_64"] = o.X8664
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *RegistrySourceListItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"name",
		"repositories",
		"aarch64",
		"x86_64",
		"created_at",
		"updated_at",
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

	varRegistrySourceListItem := _RegistrySourceListItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySourceListItem)

	if err != nil {
		return err
	}

	*o = RegistrySourceListItem(varRegistrySourceListItem)

	return err
}

type NullableRegistrySourceListItem struct {
	value *RegistrySourceListItem
	isSet bool
}

func (v NullableRegistrySourceListItem) Get() *RegistrySourceListItem {
	return v.value
}

func (v *NullableRegistrySourceListItem) Set(val *RegistrySourceListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySourceListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySourceListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySourceListItem(val *RegistrySourceListItem) *NullableRegistrySourceListItem {
	return &NullableRegistrySourceListItem{value: val, isSet: true}
}

func (v NullableRegistrySourceListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySourceListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


