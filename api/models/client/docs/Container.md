# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**ImageId** | **NullableString** |  | 
**ImageTag** | **NullableString** |  | 
**ImageDigest** | **NullableString** |  | 
**DockerContainerId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartedAt** | **time.Time** |  | 
**FinishedAt** | **NullableTime** |  | 
**Status** | **string** |  | 
**TargetStatus** | **string** |  | 
**ExitCode** | **int32** |  | 
**Error** | **NullableString** |  | 

## Methods

### NewContainer

`func NewContainer(object string, id string, deviceId string, imageId NullableString, imageTag NullableString, imageDigest NullableString, dockerContainerId string, createdAt time.Time, updatedAt time.Time, startedAt time.Time, finishedAt NullableTime, status string, targetStatus string, exitCode int32, error_ NullableString, ) *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Container) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Container) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Container) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Container) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Container) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Container) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *Container) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Container) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Container) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetImageId

`func (o *Container) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Container) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *Container) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### SetImageIdNil

`func (o *Container) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *Container) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetImageTag

`func (o *Container) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *Container) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *Container) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.


### SetImageTagNil

`func (o *Container) SetImageTagNil(b bool)`

 SetImageTagNil sets the value for ImageTag to be an explicit nil

### UnsetImageTag
`func (o *Container) UnsetImageTag()`

UnsetImageTag ensures that no value is present for ImageTag, not even an explicit nil
### GetImageDigest

`func (o *Container) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *Container) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *Container) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.


### SetImageDigestNil

`func (o *Container) SetImageDigestNil(b bool)`

 SetImageDigestNil sets the value for ImageDigest to be an explicit nil

### UnsetImageDigest
`func (o *Container) UnsetImageDigest()`

UnsetImageDigest ensures that no value is present for ImageDigest, not even an explicit nil
### GetDockerContainerId

`func (o *Container) GetDockerContainerId() string`

GetDockerContainerId returns the DockerContainerId field if non-nil, zero value otherwise.

### GetDockerContainerIdOk

`func (o *Container) GetDockerContainerIdOk() (*string, bool)`

GetDockerContainerIdOk returns a tuple with the DockerContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerId

`func (o *Container) SetDockerContainerId(v string)`

SetDockerContainerId sets DockerContainerId field to given value.


### GetCreatedAt

`func (o *Container) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Container) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Container) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Container) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Container) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Container) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartedAt

`func (o *Container) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Container) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Container) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *Container) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Container) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Container) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *Container) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *Container) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetStatus

`func (o *Container) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Container) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Container) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTargetStatus

`func (o *Container) GetTargetStatus() string`

GetTargetStatus returns the TargetStatus field if non-nil, zero value otherwise.

### GetTargetStatusOk

`func (o *Container) GetTargetStatusOk() (*string, bool)`

GetTargetStatusOk returns a tuple with the TargetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatus

`func (o *Container) SetTargetStatus(v string)`

SetTargetStatus sets TargetStatus field to given value.


### GetExitCode

`func (o *Container) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Container) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Container) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetError

`func (o *Container) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Container) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Container) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *Container) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Container) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


