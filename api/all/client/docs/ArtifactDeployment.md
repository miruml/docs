# ArtifactDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**Status** | **string** |  | 
**ActivityStatus** | **string** |  | 
**ErrorStatus** | **string** |  | 
**CreatedBy** | [**NullableUser**](User.md) |  | 
**CreatedAt** | **time.Time** |  | 
**StartedAt** | **NullableTime** |  | 
**FinishedAt** | **NullableTime** |  | 
**RemovedAt** | **NullableTime** |  | 
**OnDevice** | **bool** |  | 

## Methods

### NewArtifactDeployment

`func NewArtifactDeployment(object string, id string, status string, activityStatus string, errorStatus string, createdBy NullableUser, createdAt time.Time, startedAt NullableTime, finishedAt NullableTime, removedAt NullableTime, onDevice bool, ) *ArtifactDeployment`

NewArtifactDeployment instantiates a new ArtifactDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactDeploymentWithDefaults

`func NewArtifactDeploymentWithDefaults() *ArtifactDeployment`

NewArtifactDeploymentWithDefaults instantiates a new ArtifactDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ArtifactDeployment) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ArtifactDeployment) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ArtifactDeployment) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *ArtifactDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactDeployment) SetId(v string)`

SetId sets Id field to given value.


### GetDevice

`func (o *ArtifactDeployment) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ArtifactDeployment) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ArtifactDeployment) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ArtifactDeployment) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetStatus

`func (o *ArtifactDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArtifactDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArtifactDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetActivityStatus

`func (o *ArtifactDeployment) GetActivityStatus() string`

GetActivityStatus returns the ActivityStatus field if non-nil, zero value otherwise.

### GetActivityStatusOk

`func (o *ArtifactDeployment) GetActivityStatusOk() (*string, bool)`

GetActivityStatusOk returns a tuple with the ActivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatus

`func (o *ArtifactDeployment) SetActivityStatus(v string)`

SetActivityStatus sets ActivityStatus field to given value.


### GetErrorStatus

`func (o *ArtifactDeployment) GetErrorStatus() string`

GetErrorStatus returns the ErrorStatus field if non-nil, zero value otherwise.

### GetErrorStatusOk

`func (o *ArtifactDeployment) GetErrorStatusOk() (*string, bool)`

GetErrorStatusOk returns a tuple with the ErrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStatus

`func (o *ArtifactDeployment) SetErrorStatus(v string)`

SetErrorStatus sets ErrorStatus field to given value.


### GetCreatedBy

`func (o *ArtifactDeployment) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArtifactDeployment) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArtifactDeployment) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *ArtifactDeployment) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ArtifactDeployment) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ArtifactDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *ArtifactDeployment) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ArtifactDeployment) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ArtifactDeployment) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *ArtifactDeployment) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ArtifactDeployment) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetFinishedAt

`func (o *ArtifactDeployment) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ArtifactDeployment) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ArtifactDeployment) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *ArtifactDeployment) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ArtifactDeployment) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetRemovedAt

`func (o *ArtifactDeployment) GetRemovedAt() time.Time`

GetRemovedAt returns the RemovedAt field if non-nil, zero value otherwise.

### GetRemovedAtOk

`func (o *ArtifactDeployment) GetRemovedAtOk() (*time.Time, bool)`

GetRemovedAtOk returns a tuple with the RemovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAt

`func (o *ArtifactDeployment) SetRemovedAt(v time.Time)`

SetRemovedAt sets RemovedAt field to given value.


### SetRemovedAtNil

`func (o *ArtifactDeployment) SetRemovedAtNil(b bool)`

 SetRemovedAtNil sets the value for RemovedAt to be an explicit nil

### UnsetRemovedAt
`func (o *ArtifactDeployment) UnsetRemovedAt()`

UnsetRemovedAt ensures that no value is present for RemovedAt, not even an explicit nil
### GetOnDevice

`func (o *ArtifactDeployment) GetOnDevice() bool`

GetOnDevice returns the OnDevice field if non-nil, zero value otherwise.

### GetOnDeviceOk

`func (o *ArtifactDeployment) GetOnDeviceOk() (*bool, bool)`

GetOnDeviceOk returns a tuple with the OnDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDevice

`func (o *ArtifactDeployment) SetOnDevice(v bool)`

SetOnDevice sets OnDevice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


