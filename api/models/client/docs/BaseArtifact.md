# BaseArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**Status** | [**ArtifactStatus**](ArtifactStatus.md) |  | 
**Digest** | **string** |  | 
**Aarch64** | **bool** |  | 
**X8664** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**ReadyAt** | **NullableTime** |  | 
**FailedAt** | **NullableTime** |  | 
**SourceId** | **string** |  | 
**SourceType** | **string** |  | 

## Methods

### NewBaseArtifact

`func NewBaseArtifact(object string, id string, status ArtifactStatus, digest string, aarch64 bool, x8664 bool, createdAt time.Time, readyAt NullableTime, failedAt NullableTime, sourceId string, sourceType string, ) *BaseArtifact`

NewBaseArtifact instantiates a new BaseArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseArtifactWithDefaults

`func NewBaseArtifactWithDefaults() *BaseArtifact`

NewBaseArtifactWithDefaults instantiates a new BaseArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BaseArtifact) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BaseArtifact) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BaseArtifact) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *BaseArtifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseArtifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseArtifact) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BaseArtifact) GetStatus() ArtifactStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseArtifact) GetStatusOk() (*ArtifactStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseArtifact) SetStatus(v ArtifactStatus)`

SetStatus sets Status field to given value.


### GetDigest

`func (o *BaseArtifact) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *BaseArtifact) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *BaseArtifact) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetAarch64

`func (o *BaseArtifact) GetAarch64() bool`

GetAarch64 returns the Aarch64 field if non-nil, zero value otherwise.

### GetAarch64Ok

`func (o *BaseArtifact) GetAarch64Ok() (*bool, bool)`

GetAarch64Ok returns a tuple with the Aarch64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAarch64

`func (o *BaseArtifact) SetAarch64(v bool)`

SetAarch64 sets Aarch64 field to given value.


### GetX8664

`func (o *BaseArtifact) GetX8664() bool`

GetX8664 returns the X8664 field if non-nil, zero value otherwise.

### GetX8664Ok

`func (o *BaseArtifact) GetX8664Ok() (*bool, bool)`

GetX8664Ok returns a tuple with the X8664 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX8664

`func (o *BaseArtifact) SetX8664(v bool)`

SetX8664 sets X8664 field to given value.


### GetCreatedAt

`func (o *BaseArtifact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseArtifact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseArtifact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReadyAt

`func (o *BaseArtifact) GetReadyAt() time.Time`

GetReadyAt returns the ReadyAt field if non-nil, zero value otherwise.

### GetReadyAtOk

`func (o *BaseArtifact) GetReadyAtOk() (*time.Time, bool)`

GetReadyAtOk returns a tuple with the ReadyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyAt

`func (o *BaseArtifact) SetReadyAt(v time.Time)`

SetReadyAt sets ReadyAt field to given value.


### SetReadyAtNil

`func (o *BaseArtifact) SetReadyAtNil(b bool)`

 SetReadyAtNil sets the value for ReadyAt to be an explicit nil

### UnsetReadyAt
`func (o *BaseArtifact) UnsetReadyAt()`

UnsetReadyAt ensures that no value is present for ReadyAt, not even an explicit nil
### GetFailedAt

`func (o *BaseArtifact) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *BaseArtifact) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *BaseArtifact) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.


### SetFailedAtNil

`func (o *BaseArtifact) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *BaseArtifact) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetSourceId

`func (o *BaseArtifact) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseArtifact) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseArtifact) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceType

`func (o *BaseArtifact) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BaseArtifact) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BaseArtifact) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


