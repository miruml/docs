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
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewArtifactDeployment

`func NewArtifactDeployment(object string, id string, status string, activityStatus string, errorStatus string, ) *ArtifactDeployment`

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

### HasCreatedAt

`func (o *ArtifactDeployment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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

### HasFinishedAt

`func (o *ArtifactDeployment) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *ArtifactDeployment) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ArtifactDeployment) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


