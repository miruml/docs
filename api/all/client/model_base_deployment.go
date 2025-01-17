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

// checks if the BaseDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseDeployment{}

// BaseDeployment struct for BaseDeployment
type BaseDeployment struct {
	Object string `json:"object"`
	Id string `json:"id"`
	DeviceId string `json:"device_id"`
	Status string `json:"status"`
	ActivityStatus string `json:"activity_status"`
	ErrorStatus string `json:"error_status"`
	TargetStatus string `json:"target_status"`
	CreatedAt time.Time `json:"created_at"`
	DownloadingAt NullableTime `json:"downloading_at"`
	DownloadedAt NullableTime `json:"downloaded_at"`
	BootingAt NullableTime `json:"booting_at"`
	ActiveAt NullableTime `json:"active_at"`
	StoppingAt NullableTime `json:"stopping_at"`
	StoppedAt NullableTime `json:"stopped_at"`
	RemovingAt NullableTime `json:"removing_at"`
	ArchivedAt NullableTime `json:"archived_at"`
	CooldownEndsAt time.Time `json:"cooldown_ends_at"`
}

type _BaseDeployment BaseDeployment

// NewBaseDeployment instantiates a new BaseDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseDeployment(object string, id string, deviceId string, status string, activityStatus string, errorStatus string, targetStatus string, createdAt time.Time, downloadingAt NullableTime, downloadedAt NullableTime, bootingAt NullableTime, activeAt NullableTime, stoppingAt NullableTime, stoppedAt NullableTime, removingAt NullableTime, archivedAt NullableTime, cooldownEndsAt time.Time) *BaseDeployment {
	this := BaseDeployment{}
	this.Object = object
	this.Id = id
	this.DeviceId = deviceId
	this.Status = status
	this.ActivityStatus = activityStatus
	this.ErrorStatus = errorStatus
	this.TargetStatus = targetStatus
	this.CreatedAt = createdAt
	this.DownloadingAt = downloadingAt
	this.DownloadedAt = downloadedAt
	this.BootingAt = bootingAt
	this.ActiveAt = activeAt
	this.StoppingAt = stoppingAt
	this.StoppedAt = stoppedAt
	this.RemovingAt = removingAt
	this.ArchivedAt = archivedAt
	this.CooldownEndsAt = cooldownEndsAt
	return &this
}

// NewBaseDeploymentWithDefaults instantiates a new BaseDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseDeploymentWithDefaults() *BaseDeployment {
	this := BaseDeployment{}
	return &this
}

// GetObject returns the Object field value
func (o *BaseDeployment) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *BaseDeployment) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *BaseDeployment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BaseDeployment) SetId(v string) {
	o.Id = v
}

// GetDeviceId returns the DeviceId field value
func (o *BaseDeployment) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *BaseDeployment) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetStatus returns the Status field value
func (o *BaseDeployment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BaseDeployment) SetStatus(v string) {
	o.Status = v
}

// GetActivityStatus returns the ActivityStatus field value
func (o *BaseDeployment) GetActivityStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivityStatus
}

// GetActivityStatusOk returns a tuple with the ActivityStatus field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetActivityStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivityStatus, true
}

// SetActivityStatus sets field value
func (o *BaseDeployment) SetActivityStatus(v string) {
	o.ActivityStatus = v
}

// GetErrorStatus returns the ErrorStatus field value
func (o *BaseDeployment) GetErrorStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorStatus
}

// GetErrorStatusOk returns a tuple with the ErrorStatus field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetErrorStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorStatus, true
}

// SetErrorStatus sets field value
func (o *BaseDeployment) SetErrorStatus(v string) {
	o.ErrorStatus = v
}

// GetTargetStatus returns the TargetStatus field value
func (o *BaseDeployment) GetTargetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetStatus
}

// GetTargetStatusOk returns a tuple with the TargetStatus field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetTargetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetStatus, true
}

// SetTargetStatus sets field value
func (o *BaseDeployment) SetTargetStatus(v string) {
	o.TargetStatus = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BaseDeployment) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BaseDeployment) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDownloadingAt returns the DownloadingAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BaseDeployment) GetDownloadingAt() time.Time {
	if o == nil || o.DownloadingAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DownloadingAt.Get()
}

// GetDownloadingAtOk returns a tuple with the DownloadingAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDeployment) GetDownloadingAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadingAt.Get(), o.DownloadingAt.IsSet()
}

// SetDownloadingAt sets field value
func (o *BaseDeployment) SetDownloadingAt(v time.Time) {
	o.DownloadingAt.Set(&v)
}

// GetDownloadedAt returns the DownloadedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BaseDeployment) GetDownloadedAt() time.Time {
	if o == nil || o.DownloadedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DownloadedAt.Get()
}

// GetDownloadedAtOk returns a tuple with the DownloadedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDeployment) GetDownloadedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadedAt.Get(), o.DownloadedAt.IsSet()
}

// SetDownloadedAt sets field value
func (o *BaseDeployment) SetDownloadedAt(v time.Time) {
	o.DownloadedAt.Set(&v)
}

// GetBootingAt returns the BootingAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BaseDeployment) GetBootingAt() time.Time {
	if o == nil || o.BootingAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.BootingAt.Get()
}

// GetBootingAtOk returns a tuple with the BootingAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDeployment) GetBootingAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.BootingAt.Get(), o.BootingAt.IsSet()
}

// SetBootingAt sets field value
func (o *BaseDeployment) SetBootingAt(v time.Time) {
	o.BootingAt.Set(&v)
}

// GetActiveAt returns the ActiveAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BaseDeployment) GetActiveAt() time.Time {
	if o == nil || o.ActiveAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveAt.Get()
}

// GetActiveAtOk returns a tuple with the ActiveAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDeployment) GetActiveAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveAt.Get(), o.ActiveAt.IsSet()
}

// SetActiveAt sets field value
func (o *BaseDeployment) SetActiveAt(v time.Time) {
	o.ActiveAt.Set(&v)
}

// GetStoppingAt returns the StoppingAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BaseDeployment) GetStoppingAt() time.Time {
	if o == nil || o.StoppingAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StoppingAt.Get()
}

// GetStoppingAtOk returns a tuple with the StoppingAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDeployment) GetStoppingAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoppingAt.Get(), o.StoppingAt.IsSet()
}

// SetStoppingAt sets field value
func (o *BaseDeployment) SetStoppingAt(v time.Time) {
	o.StoppingAt.Set(&v)
}

// GetStoppedAt returns the StoppedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BaseDeployment) GetStoppedAt() time.Time {
	if o == nil || o.StoppedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StoppedAt.Get()
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDeployment) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoppedAt.Get(), o.StoppedAt.IsSet()
}

// SetStoppedAt sets field value
func (o *BaseDeployment) SetStoppedAt(v time.Time) {
	o.StoppedAt.Set(&v)
}

// GetRemovingAt returns the RemovingAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BaseDeployment) GetRemovingAt() time.Time {
	if o == nil || o.RemovingAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.RemovingAt.Get()
}

// GetRemovingAtOk returns a tuple with the RemovingAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDeployment) GetRemovingAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovingAt.Get(), o.RemovingAt.IsSet()
}

// SetRemovingAt sets field value
func (o *BaseDeployment) SetRemovingAt(v time.Time) {
	o.RemovingAt.Set(&v)
}

// GetArchivedAt returns the ArchivedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BaseDeployment) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ArchivedAt.Get()
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseDeployment) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchivedAt.Get(), o.ArchivedAt.IsSet()
}

// SetArchivedAt sets field value
func (o *BaseDeployment) SetArchivedAt(v time.Time) {
	o.ArchivedAt.Set(&v)
}

// GetCooldownEndsAt returns the CooldownEndsAt field value
func (o *BaseDeployment) GetCooldownEndsAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CooldownEndsAt
}

// GetCooldownEndsAtOk returns a tuple with the CooldownEndsAt field value
// and a boolean to check if the value has been set.
func (o *BaseDeployment) GetCooldownEndsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CooldownEndsAt, true
}

// SetCooldownEndsAt sets field value
func (o *BaseDeployment) SetCooldownEndsAt(v time.Time) {
	o.CooldownEndsAt = v
}

func (o BaseDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["device_id"] = o.DeviceId
	toSerialize["status"] = o.Status
	toSerialize["activity_status"] = o.ActivityStatus
	toSerialize["error_status"] = o.ErrorStatus
	toSerialize["target_status"] = o.TargetStatus
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["downloading_at"] = o.DownloadingAt.Get()
	toSerialize["downloaded_at"] = o.DownloadedAt.Get()
	toSerialize["booting_at"] = o.BootingAt.Get()
	toSerialize["active_at"] = o.ActiveAt.Get()
	toSerialize["stopping_at"] = o.StoppingAt.Get()
	toSerialize["stopped_at"] = o.StoppedAt.Get()
	toSerialize["removing_at"] = o.RemovingAt.Get()
	toSerialize["archived_at"] = o.ArchivedAt.Get()
	toSerialize["cooldown_ends_at"] = o.CooldownEndsAt
	return toSerialize, nil
}

func (o *BaseDeployment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"device_id",
		"status",
		"activity_status",
		"error_status",
		"target_status",
		"created_at",
		"downloading_at",
		"downloaded_at",
		"booting_at",
		"active_at",
		"stopping_at",
		"stopped_at",
		"removing_at",
		"archived_at",
		"cooldown_ends_at",
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

	varBaseDeployment := _BaseDeployment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseDeployment)

	if err != nil {
		return err
	}

	*o = BaseDeployment(varBaseDeployment)

	return err
}

type NullableBaseDeployment struct {
	value *BaseDeployment
	isSet bool
}

func (v NullableBaseDeployment) Get() *BaseDeployment {
	return v.value
}

func (v *NullableBaseDeployment) Set(val *BaseDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseDeployment(val *BaseDeployment) *NullableBaseDeployment {
	return &NullableBaseDeployment{value: val, isSet: true}
}

func (v NullableBaseDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

