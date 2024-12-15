# Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**Status** | [**ArtifactStatus**](ArtifactStatus.md) |  | 
**SourceId** | **string** |  | 
**SourceType** | **string** |  | 
**GithubSourceData** | [**GitHubSourceData**](GitHubSourceData.md) |  | 
**Aarch64** | **bool** |  | 
**X8664** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**ReadyAt** | **NullableTime** |  | 
**FailedAt** | **NullableTime** |  | 

## Methods

### NewArtifact

`func NewArtifact(object string, id string, status ArtifactStatus, sourceId string, sourceType string, githubSourceData GitHubSourceData, aarch64 bool, x8664 bool, createdAt time.Time, readyAt NullableTime, failedAt NullableTime, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Artifact) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Artifact) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Artifact) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Artifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Artifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Artifact) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Artifact) GetStatus() ArtifactStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Artifact) GetStatusOk() (*ArtifactStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Artifact) SetStatus(v ArtifactStatus)`

SetStatus sets Status field to given value.


### GetSourceId

`func (o *Artifact) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Artifact) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Artifact) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceType

`func (o *Artifact) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Artifact) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Artifact) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetGithubSourceData

`func (o *Artifact) GetGithubSourceData() GitHubSourceData`

GetGithubSourceData returns the GithubSourceData field if non-nil, zero value otherwise.

### GetGithubSourceDataOk

`func (o *Artifact) GetGithubSourceDataOk() (*GitHubSourceData, bool)`

GetGithubSourceDataOk returns a tuple with the GithubSourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSourceData

`func (o *Artifact) SetGithubSourceData(v GitHubSourceData)`

SetGithubSourceData sets GithubSourceData field to given value.


### GetAarch64

`func (o *Artifact) GetAarch64() bool`

GetAarch64 returns the Aarch64 field if non-nil, zero value otherwise.

### GetAarch64Ok

`func (o *Artifact) GetAarch64Ok() (*bool, bool)`

GetAarch64Ok returns a tuple with the Aarch64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAarch64

`func (o *Artifact) SetAarch64(v bool)`

SetAarch64 sets Aarch64 field to given value.


### GetX8664

`func (o *Artifact) GetX8664() bool`

GetX8664 returns the X8664 field if non-nil, zero value otherwise.

### GetX8664Ok

`func (o *Artifact) GetX8664Ok() (*bool, bool)`

GetX8664Ok returns a tuple with the X8664 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX8664

`func (o *Artifact) SetX8664(v bool)`

SetX8664 sets X8664 field to given value.


### GetCreatedAt

`func (o *Artifact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Artifact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Artifact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReadyAt

`func (o *Artifact) GetReadyAt() time.Time`

GetReadyAt returns the ReadyAt field if non-nil, zero value otherwise.

### GetReadyAtOk

`func (o *Artifact) GetReadyAtOk() (*time.Time, bool)`

GetReadyAtOk returns a tuple with the ReadyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyAt

`func (o *Artifact) SetReadyAt(v time.Time)`

SetReadyAt sets ReadyAt field to given value.


### SetReadyAtNil

`func (o *Artifact) SetReadyAtNil(b bool)`

 SetReadyAtNil sets the value for ReadyAt to be an explicit nil

### UnsetReadyAt
`func (o *Artifact) UnsetReadyAt()`

UnsetReadyAt ensures that no value is present for ReadyAt, not even an explicit nil
### GetFailedAt

`func (o *Artifact) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *Artifact) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *Artifact) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.


### SetFailedAtNil

`func (o *Artifact) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *Artifact) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


