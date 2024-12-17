# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**Artifact** | [**ArtifactBase**](ArtifactBase.md) |  | 
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

### NewDeployment

`func NewDeployment(object string, id string, deviceId string, artifact ArtifactBase, status string, activityStatus string, errorStatus string, targetStatus string, createdAt time.Time, downloadingAt NullableTime, downloadedAt NullableTime, bootingAt NullableTime, activeAt NullableTime, stoppingAt NullableTime, stoppedAt NullableTime, removingAt NullableTime, archivedAt NullableTime, cooldownEndsAt time.Time, ) *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Deployment) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Deployment) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Deployment) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Deployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *Deployment) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Deployment) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Deployment) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetArtifact

`func (o *Deployment) GetArtifact() ArtifactBase`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *Deployment) GetArtifactOk() (*ArtifactBase, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *Deployment) SetArtifact(v ArtifactBase)`

SetArtifact sets Artifact field to given value.


### GetStatus

`func (o *Deployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetActivityStatus

`func (o *Deployment) GetActivityStatus() string`

GetActivityStatus returns the ActivityStatus field if non-nil, zero value otherwise.

### GetActivityStatusOk

`func (o *Deployment) GetActivityStatusOk() (*string, bool)`

GetActivityStatusOk returns a tuple with the ActivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatus

`func (o *Deployment) SetActivityStatus(v string)`

SetActivityStatus sets ActivityStatus field to given value.


### GetErrorStatus

`func (o *Deployment) GetErrorStatus() string`

GetErrorStatus returns the ErrorStatus field if non-nil, zero value otherwise.

### GetErrorStatusOk

`func (o *Deployment) GetErrorStatusOk() (*string, bool)`

GetErrorStatusOk returns a tuple with the ErrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStatus

`func (o *Deployment) SetErrorStatus(v string)`

SetErrorStatus sets ErrorStatus field to given value.


### GetTargetStatus

`func (o *Deployment) GetTargetStatus() string`

GetTargetStatus returns the TargetStatus field if non-nil, zero value otherwise.

### GetTargetStatusOk

`func (o *Deployment) GetTargetStatusOk() (*string, bool)`

GetTargetStatusOk returns a tuple with the TargetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatus

`func (o *Deployment) SetTargetStatus(v string)`

SetTargetStatus sets TargetStatus field to given value.


### GetCreatedAt

`func (o *Deployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Deployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Deployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownloadingAt

`func (o *Deployment) GetDownloadingAt() time.Time`

GetDownloadingAt returns the DownloadingAt field if non-nil, zero value otherwise.

### GetDownloadingAtOk

`func (o *Deployment) GetDownloadingAtOk() (*time.Time, bool)`

GetDownloadingAtOk returns a tuple with the DownloadingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadingAt

`func (o *Deployment) SetDownloadingAt(v time.Time)`

SetDownloadingAt sets DownloadingAt field to given value.


### SetDownloadingAtNil

`func (o *Deployment) SetDownloadingAtNil(b bool)`

 SetDownloadingAtNil sets the value for DownloadingAt to be an explicit nil

### UnsetDownloadingAt
`func (o *Deployment) UnsetDownloadingAt()`

UnsetDownloadingAt ensures that no value is present for DownloadingAt, not even an explicit nil
### GetDownloadedAt

`func (o *Deployment) GetDownloadedAt() time.Time`

GetDownloadedAt returns the DownloadedAt field if non-nil, zero value otherwise.

### GetDownloadedAtOk

`func (o *Deployment) GetDownloadedAtOk() (*time.Time, bool)`

GetDownloadedAtOk returns a tuple with the DownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedAt

`func (o *Deployment) SetDownloadedAt(v time.Time)`

SetDownloadedAt sets DownloadedAt field to given value.


### SetDownloadedAtNil

`func (o *Deployment) SetDownloadedAtNil(b bool)`

 SetDownloadedAtNil sets the value for DownloadedAt to be an explicit nil

### UnsetDownloadedAt
`func (o *Deployment) UnsetDownloadedAt()`

UnsetDownloadedAt ensures that no value is present for DownloadedAt, not even an explicit nil
### GetBootingAt

`func (o *Deployment) GetBootingAt() time.Time`

GetBootingAt returns the BootingAt field if non-nil, zero value otherwise.

### GetBootingAtOk

`func (o *Deployment) GetBootingAtOk() (*time.Time, bool)`

GetBootingAtOk returns a tuple with the BootingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootingAt

`func (o *Deployment) SetBootingAt(v time.Time)`

SetBootingAt sets BootingAt field to given value.


### SetBootingAtNil

`func (o *Deployment) SetBootingAtNil(b bool)`

 SetBootingAtNil sets the value for BootingAt to be an explicit nil

### UnsetBootingAt
`func (o *Deployment) UnsetBootingAt()`

UnsetBootingAt ensures that no value is present for BootingAt, not even an explicit nil
### GetActiveAt

`func (o *Deployment) GetActiveAt() time.Time`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *Deployment) GetActiveAtOk() (*time.Time, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *Deployment) SetActiveAt(v time.Time)`

SetActiveAt sets ActiveAt field to given value.


### SetActiveAtNil

`func (o *Deployment) SetActiveAtNil(b bool)`

 SetActiveAtNil sets the value for ActiveAt to be an explicit nil

### UnsetActiveAt
`func (o *Deployment) UnsetActiveAt()`

UnsetActiveAt ensures that no value is present for ActiveAt, not even an explicit nil
### GetStoppingAt

`func (o *Deployment) GetStoppingAt() time.Time`

GetStoppingAt returns the StoppingAt field if non-nil, zero value otherwise.

### GetStoppingAtOk

`func (o *Deployment) GetStoppingAtOk() (*time.Time, bool)`

GetStoppingAtOk returns a tuple with the StoppingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppingAt

`func (o *Deployment) SetStoppingAt(v time.Time)`

SetStoppingAt sets StoppingAt field to given value.


### SetStoppingAtNil

`func (o *Deployment) SetStoppingAtNil(b bool)`

 SetStoppingAtNil sets the value for StoppingAt to be an explicit nil

### UnsetStoppingAt
`func (o *Deployment) UnsetStoppingAt()`

UnsetStoppingAt ensures that no value is present for StoppingAt, not even an explicit nil
### GetStoppedAt

`func (o *Deployment) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *Deployment) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *Deployment) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.


### SetStoppedAtNil

`func (o *Deployment) SetStoppedAtNil(b bool)`

 SetStoppedAtNil sets the value for StoppedAt to be an explicit nil

### UnsetStoppedAt
`func (o *Deployment) UnsetStoppedAt()`

UnsetStoppedAt ensures that no value is present for StoppedAt, not even an explicit nil
### GetRemovingAt

`func (o *Deployment) GetRemovingAt() time.Time`

GetRemovingAt returns the RemovingAt field if non-nil, zero value otherwise.

### GetRemovingAtOk

`func (o *Deployment) GetRemovingAtOk() (*time.Time, bool)`

GetRemovingAtOk returns a tuple with the RemovingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovingAt

`func (o *Deployment) SetRemovingAt(v time.Time)`

SetRemovingAt sets RemovingAt field to given value.


### SetRemovingAtNil

`func (o *Deployment) SetRemovingAtNil(b bool)`

 SetRemovingAtNil sets the value for RemovingAt to be an explicit nil

### UnsetRemovingAt
`func (o *Deployment) UnsetRemovingAt()`

UnsetRemovingAt ensures that no value is present for RemovingAt, not even an explicit nil
### GetArchivedAt

`func (o *Deployment) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Deployment) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Deployment) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.


### SetArchivedAtNil

`func (o *Deployment) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *Deployment) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetCooldownEndsAt

`func (o *Deployment) GetCooldownEndsAt() time.Time`

GetCooldownEndsAt returns the CooldownEndsAt field if non-nil, zero value otherwise.

### GetCooldownEndsAtOk

`func (o *Deployment) GetCooldownEndsAtOk() (*time.Time, bool)`

GetCooldownEndsAtOk returns a tuple with the CooldownEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownEndsAt

`func (o *Deployment) SetCooldownEndsAt(v time.Time)`

SetCooldownEndsAt sets CooldownEndsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


