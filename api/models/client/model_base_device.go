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

// checks if the BaseDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseDevice{}

// BaseDevice struct for BaseDevice
type BaseDevice struct {
	Object string `json:"object"`
	Id string `json:"id"`
	Name string `json:"name"`
	Hardware string `json:"hardware"`
	OperatingSystem NullableString `json:"operating_system"`
	Architecture NullableString `json:"architecture"`
	Status string `json:"status"`
	LastReportedStatus string `json:"last_reported_status"`
	MiruVersion NullableString `json:"miru_version"`
	CreatedAt time.Time `json:"created_at"`
	SyncedAt time.Time `json:"synced_at"`
}

type _BaseDevice BaseDevice

// NewBaseDevice instantiates a new BaseDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseDevice(object string, id string, name string, hardware string, operatingSystem NullableString, architecture NullableString, status string, lastReportedStatus string, miruVersion NullableString, createdAt time.Time, syncedAt time.Time) *BaseDevice {
	this := BaseDevice{}
	this.Object = object
	this.Id = id
	this.Name = name
	this.Hardware = hardware
	this.OperatingSystem = operatingSystem
	this.Architecture = architecture
	this.Status = status
	this.LastReportedStatus = lastReportedStatus
	this.MiruVersion = miruVersion
	this.CreatedAt = createdAt
	this.SyncedAt = syncedAt
	return &this
}

// NewBaseDeviceWithDefaults instantiates a new BaseDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseDeviceWithDefaults() *BaseDevice {
	this := BaseDevice{}
	return &this
}

// GetObject returns the Object field value
func (o *BaseDevice) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *BaseDevice) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *BaseDevice) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *BaseDevice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BaseDevice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BaseDevice) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BaseDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseDevice) SetName(v string) {
	o.Name = v
}

// GetHardware returns the Hardware field value
func (o *BaseDevice) GetHardware() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hardware
}

// GetHardwareOk returns a tuple with the Hardware field value
// and a boolean to check if the value has been set.
func (o *BaseDevice) GetHardwareOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hardware, true
}

// SetHardware sets field value
func (o *BaseDevice) SetHardware(v string) {
	o.Hardware = v
}

// GetOperatingSystem returns the OperatingSystem field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BaseDevice) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem.Get() == nil {
		var ret string
		return ret
	}

	return *o.OperatingSystem.Get()
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDevice) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperatingSystem.Get(), o.OperatingSystem.IsSet()
}

// SetOperatingSystem sets field value
func (o *BaseDevice) SetOperatingSystem(v string) {
	o.OperatingSystem.Set(&v)
}

// GetArchitecture returns the Architecture field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BaseDevice) GetArchitecture() string {
	if o == nil || o.Architecture.Get() == nil {
		var ret string
		return ret
	}

	return *o.Architecture.Get()
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDevice) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Architecture.Get(), o.Architecture.IsSet()
}

// SetArchitecture sets field value
func (o *BaseDevice) SetArchitecture(v string) {
	o.Architecture.Set(&v)
}

// GetStatus returns the Status field value
func (o *BaseDevice) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BaseDevice) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BaseDevice) SetStatus(v string) {
	o.Status = v
}

// GetLastReportedStatus returns the LastReportedStatus field value
func (o *BaseDevice) GetLastReportedStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastReportedStatus
}

// GetLastReportedStatusOk returns a tuple with the LastReportedStatus field value
// and a boolean to check if the value has been set.
func (o *BaseDevice) GetLastReportedStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastReportedStatus, true
}

// SetLastReportedStatus sets field value
func (o *BaseDevice) SetLastReportedStatus(v string) {
	o.LastReportedStatus = v
}

// GetMiruVersion returns the MiruVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BaseDevice) GetMiruVersion() string {
	if o == nil || o.MiruVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.MiruVersion.Get()
}

// GetMiruVersionOk returns a tuple with the MiruVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDevice) GetMiruVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiruVersion.Get(), o.MiruVersion.IsSet()
}

// SetMiruVersion sets field value
func (o *BaseDevice) SetMiruVersion(v string) {
	o.MiruVersion.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *BaseDevice) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BaseDevice) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BaseDevice) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetSyncedAt returns the SyncedAt field value
func (o *BaseDevice) GetSyncedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SyncedAt
}

// GetSyncedAtOk returns a tuple with the SyncedAt field value
// and a boolean to check if the value has been set.
func (o *BaseDevice) GetSyncedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncedAt, true
}

// SetSyncedAt sets field value
func (o *BaseDevice) SetSyncedAt(v time.Time) {
	o.SyncedAt = v
}

func (o BaseDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["hardware"] = o.Hardware
	toSerialize["operating_system"] = o.OperatingSystem.Get()
	toSerialize["architecture"] = o.Architecture.Get()
	toSerialize["status"] = o.Status
	toSerialize["last_reported_status"] = o.LastReportedStatus
	toSerialize["miru_version"] = o.MiruVersion.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["synced_at"] = o.SyncedAt
	return toSerialize, nil
}

func (o *BaseDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"name",
		"hardware",
		"operating_system",
		"architecture",
		"status",
		"last_reported_status",
		"miru_version",
		"created_at",
		"synced_at",
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

	varBaseDevice := _BaseDevice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseDevice)

	if err != nil {
		return err
	}

	*o = BaseDevice(varBaseDevice)

	return err
}

type NullableBaseDevice struct {
	value *BaseDevice
	isSet bool
}

func (v NullableBaseDevice) Get() *BaseDevice {
	return v.value
}

func (v *NullableBaseDevice) Set(val *BaseDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseDevice(val *BaseDevice) *NullableBaseDevice {
	return &NullableBaseDevice{value: val, isSet: true}
}

func (v NullableBaseDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


