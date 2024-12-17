# BaseDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**Status** | **string** |  | 
**ActivityStatus** | **string** |  | 
**ErrorStatus** | **string** |  | 
**TargetStatus** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**DownloadingAt** | **NullableTime** |  | 
**DownloadedAt** | **NullableTime** |  | 
**BootingAt** | **NullableTime** |  | 
**ActiveAt** | **NullableTime** |  | 
**StoppingAt** | **NullableTime** |  | 
**StoppedAt** | **NullableTime** |  | 
**RemovingAt** | **NullableTime** |  | 
**ArchivedAt** | **NullableTime** |  | 
**CooldownEndsAt** | **time.Time** |  | 

## Methods

### NewBaseDeployment

`func NewBaseDeployment(object string, id string, deviceId string, status string, activityStatus string, errorStatus string, targetStatus string, createdAt time.Time, downloadingAt NullableTime, downloadedAt NullableTime, bootingAt NullableTime, activeAt NullableTime, stoppingAt NullableTime, stoppedAt NullableTime, removingAt NullableTime, archivedAt NullableTime, cooldownEndsAt time.Time, ) *BaseDeployment`

NewBaseDeployment instantiates a new BaseDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDeploymentWithDefaults

`func NewBaseDeploymentWithDefaults() *BaseDeployment`

NewBaseDeploymentWithDefaults instantiates a new BaseDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BaseDeployment) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BaseDeployment) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BaseDeployment) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *BaseDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseDeployment) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *BaseDeployment) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BaseDeployment) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BaseDeployment) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetStatus

`func (o *BaseDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetActivityStatus

`func (o *BaseDeployment) GetActivityStatus() string`

GetActivityStatus returns the ActivityStatus field if non-nil, zero value otherwise.

### GetActivityStatusOk

`func (o *BaseDeployment) GetActivityStatusOk() (*string, bool)`

GetActivityStatusOk returns a tuple with the ActivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatus

`func (o *BaseDeployment) SetActivityStatus(v string)`

SetActivityStatus sets ActivityStatus field to given value.


### GetErrorStatus

`func (o *BaseDeployment) GetErrorStatus() string`

GetErrorStatus returns the ErrorStatus field if non-nil, zero value otherwise.

### GetErrorStatusOk

`func (o *BaseDeployment) GetErrorStatusOk() (*string, bool)`

GetErrorStatusOk returns a tuple with the ErrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStatus

`func (o *BaseDeployment) SetErrorStatus(v string)`

SetErrorStatus sets ErrorStatus field to given value.


### GetTargetStatus

`func (o *BaseDeployment) GetTargetStatus() string`

GetTargetStatus returns the TargetStatus field if non-nil, zero value otherwise.

### GetTargetStatusOk

`func (o *BaseDeployment) GetTargetStatusOk() (*string, bool)`

GetTargetStatusOk returns a tuple with the TargetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatus

`func (o *BaseDeployment) SetTargetStatus(v string)`

SetTargetStatus sets TargetStatus field to given value.


### GetCreatedAt

`func (o *BaseDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownloadingAt

`func (o *BaseDeployment) GetDownloadingAt() time.Time`

GetDownloadingAt returns the DownloadingAt field if non-nil, zero value otherwise.

### GetDownloadingAtOk

`func (o *BaseDeployment) GetDownloadingAtOk() (*time.Time, bool)`

GetDownloadingAtOk returns a tuple with the DownloadingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadingAt

`func (o *BaseDeployment) SetDownloadingAt(v time.Time)`

SetDownloadingAt sets DownloadingAt field to given value.


### SetDownloadingAtNil

`func (o *BaseDeployment) SetDownloadingAtNil(b bool)`

 SetDownloadingAtNil sets the value for DownloadingAt to be an explicit nil

### UnsetDownloadingAt
`func (o *BaseDeployment) UnsetDownloadingAt()`

UnsetDownloadingAt ensures that no value is present for DownloadingAt, not even an explicit nil
### GetDownloadedAt

`func (o *BaseDeployment) GetDownloadedAt() time.Time`

GetDownloadedAt returns the DownloadedAt field if non-nil, zero value otherwise.

### GetDownloadedAtOk

`func (o *BaseDeployment) GetDownloadedAtOk() (*time.Time, bool)`

GetDownloadedAtOk returns a tuple with the DownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedAt

`func (o *BaseDeployment) SetDownloadedAt(v time.Time)`

SetDownloadedAt sets DownloadedAt field to given value.


### SetDownloadedAtNil

`func (o *BaseDeployment) SetDownloadedAtNil(b bool)`

 SetDownloadedAtNil sets the value for DownloadedAt to be an explicit nil

### UnsetDownloadedAt
`func (o *BaseDeployment) UnsetDownloadedAt()`

UnsetDownloadedAt ensures that no value is present for DownloadedAt, not even an explicit nil
### GetBootingAt

`func (o *BaseDeployment) GetBootingAt() time.Time`

GetBootingAt returns the BootingAt field if non-nil, zero value otherwise.

### GetBootingAtOk

`func (o *BaseDeployment) GetBootingAtOk() (*time.Time, bool)`

GetBootingAtOk returns a tuple with the BootingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootingAt

`func (o *BaseDeployment) SetBootingAt(v time.Time)`

SetBootingAt sets BootingAt field to given value.


### SetBootingAtNil

`func (o *BaseDeployment) SetBootingAtNil(b bool)`

 SetBootingAtNil sets the value for BootingAt to be an explicit nil

### UnsetBootingAt
`func (o *BaseDeployment) UnsetBootingAt()`

UnsetBootingAt ensures that no value is present for BootingAt, not even an explicit nil
### GetActiveAt

`func (o *BaseDeployment) GetActiveAt() time.Time`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *BaseDeployment) GetActiveAtOk() (*time.Time, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *BaseDeployment) SetActiveAt(v time.Time)`

SetActiveAt sets ActiveAt field to given value.


### SetActiveAtNil

`func (o *BaseDeployment) SetActiveAtNil(b bool)`

 SetActiveAtNil sets the value for ActiveAt to be an explicit nil

### UnsetActiveAt
`func (o *BaseDeployment) UnsetActiveAt()`

UnsetActiveAt ensures that no value is present for ActiveAt, not even an explicit nil
### GetStoppingAt

`func (o *BaseDeployment) GetStoppingAt() time.Time`

GetStoppingAt returns the StoppingAt field if non-nil, zero value otherwise.

### GetStoppingAtOk

`func (o *BaseDeployment) GetStoppingAtOk() (*time.Time, bool)`

GetStoppingAtOk returns a tuple with the StoppingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppingAt

`func (o *BaseDeployment) SetStoppingAt(v time.Time)`

SetStoppingAt sets StoppingAt field to given value.


### SetStoppingAtNil

`func (o *BaseDeployment) SetStoppingAtNil(b bool)`

 SetStoppingAtNil sets the value for StoppingAt to be an explicit nil

### UnsetStoppingAt
`func (o *BaseDeployment) UnsetStoppingAt()`

UnsetStoppingAt ensures that no value is present for StoppingAt, not even an explicit nil
### GetStoppedAt

`func (o *BaseDeployment) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *BaseDeployment) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *BaseDeployment) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.


### SetStoppedAtNil

`func (o *BaseDeployment) SetStoppedAtNil(b bool)`

 SetStoppedAtNil sets the value for StoppedAt to be an explicit nil

### UnsetStoppedAt
`func (o *BaseDeployment) UnsetStoppedAt()`

UnsetStoppedAt ensures that no value is present for StoppedAt, not even an explicit nil
### GetRemovingAt

`func (o *BaseDeployment) GetRemovingAt() time.Time`

GetRemovingAt returns the RemovingAt field if non-nil, zero value otherwise.

### GetRemovingAtOk

`func (o *BaseDeployment) GetRemovingAtOk() (*time.Time, bool)`

GetRemovingAtOk returns a tuple with the RemovingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovingAt

`func (o *BaseDeployment) SetRemovingAt(v time.Time)`

SetRemovingAt sets RemovingAt field to given value.


### SetRemovingAtNil

`func (o *BaseDeployment) SetRemovingAtNil(b bool)`

 SetRemovingAtNil sets the value for RemovingAt to be an explicit nil

### UnsetRemovingAt
`func (o *BaseDeployment) UnsetRemovingAt()`

UnsetRemovingAt ensures that no value is present for RemovingAt, not even an explicit nil
### GetArchivedAt

`func (o *BaseDeployment) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *BaseDeployment) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *BaseDeployment) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.


### SetArchivedAtNil

`func (o *BaseDeployment) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *BaseDeployment) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetCooldownEndsAt

`func (o *BaseDeployment) GetCooldownEndsAt() time.Time`

GetCooldownEndsAt returns the CooldownEndsAt field if non-nil, zero value otherwise.

### GetCooldownEndsAtOk

`func (o *BaseDeployment) GetCooldownEndsAtOk() (*time.Time, bool)`

GetCooldownEndsAtOk returns a tuple with the CooldownEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownEndsAt

`func (o *BaseDeployment) SetCooldownEndsAt(v time.Time)`

SetCooldownEndsAt sets CooldownEndsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


