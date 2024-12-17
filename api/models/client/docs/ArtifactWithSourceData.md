# ArtifactWithSourceData

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
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**SourceId** | **string** |  | 
**SourceType** | **string** |  | 
**RegistrySource** | [**RegistrySource**](RegistrySource.md) |  | 
**GithubSource** | [**GitHubSource**](GitHubSource.md) |  | 
**GithubSourceData** | [**GitHubSourceData**](GitHubSourceData.md) |  | 

## Methods

### NewArtifactWithSourceData

`func NewArtifactWithSourceData(object string, id string, status ArtifactStatus, digest string, aarch64 bool, x8664 bool, createdAt time.Time, readyAt NullableTime, failedAt NullableTime, sourceId string, sourceType string, registrySource RegistrySource, githubSource GitHubSource, githubSourceData GitHubSourceData, ) *ArtifactWithSourceData`

NewArtifactWithSourceData instantiates a new ArtifactWithSourceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithSourceDataWithDefaults

`func NewArtifactWithSourceDataWithDefaults() *ArtifactWithSourceData`

NewArtifactWithSourceDataWithDefaults instantiates a new ArtifactWithSourceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ArtifactWithSourceData) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ArtifactWithSourceData) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ArtifactWithSourceData) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *ArtifactWithSourceData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactWithSourceData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactWithSourceData) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ArtifactWithSourceData) GetStatus() ArtifactStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArtifactWithSourceData) GetStatusOk() (*ArtifactStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArtifactWithSourceData) SetStatus(v ArtifactStatus)`

SetStatus sets Status field to given value.


### GetDigest

`func (o *ArtifactWithSourceData) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ArtifactWithSourceData) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ArtifactWithSourceData) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetAarch64

`func (o *ArtifactWithSourceData) GetAarch64() bool`

GetAarch64 returns the Aarch64 field if non-nil, zero value otherwise.

### GetAarch64Ok

`func (o *ArtifactWithSourceData) GetAarch64Ok() (*bool, bool)`

GetAarch64Ok returns a tuple with the Aarch64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAarch64

`func (o *ArtifactWithSourceData) SetAarch64(v bool)`

SetAarch64 sets Aarch64 field to given value.


### GetX8664

`func (o *ArtifactWithSourceData) GetX8664() bool`

GetX8664 returns the X8664 field if non-nil, zero value otherwise.

### GetX8664Ok

`func (o *ArtifactWithSourceData) GetX8664Ok() (*bool, bool)`

GetX8664Ok returns a tuple with the X8664 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX8664

`func (o *ArtifactWithSourceData) SetX8664(v bool)`

SetX8664 sets X8664 field to given value.


### GetCreatedAt

`func (o *ArtifactWithSourceData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactWithSourceData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactWithSourceData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReadyAt

`func (o *ArtifactWithSourceData) GetReadyAt() time.Time`

GetReadyAt returns the ReadyAt field if non-nil, zero value otherwise.

### GetReadyAtOk

`func (o *ArtifactWithSourceData) GetReadyAtOk() (*time.Time, bool)`

GetReadyAtOk returns a tuple with the ReadyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyAt

`func (o *ArtifactWithSourceData) SetReadyAt(v time.Time)`

SetReadyAt sets ReadyAt field to given value.


### SetReadyAtNil

`func (o *ArtifactWithSourceData) SetReadyAtNil(b bool)`

 SetReadyAtNil sets the value for ReadyAt to be an explicit nil

### UnsetReadyAt
`func (o *ArtifactWithSourceData) UnsetReadyAt()`

UnsetReadyAt ensures that no value is present for ReadyAt, not even an explicit nil
### GetFailedAt

`func (o *ArtifactWithSourceData) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *ArtifactWithSourceData) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *ArtifactWithSourceData) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.


### SetFailedAtNil

`func (o *ArtifactWithSourceData) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *ArtifactWithSourceData) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetCreatedBy

`func (o *ArtifactWithSourceData) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArtifactWithSourceData) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArtifactWithSourceData) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArtifactWithSourceData) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetSourceId

`func (o *ArtifactWithSourceData) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ArtifactWithSourceData) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ArtifactWithSourceData) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceType

`func (o *ArtifactWithSourceData) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ArtifactWithSourceData) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ArtifactWithSourceData) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetRegistrySource

`func (o *ArtifactWithSourceData) GetRegistrySource() RegistrySource`

GetRegistrySource returns the RegistrySource field if non-nil, zero value otherwise.

### GetRegistrySourceOk

`func (o *ArtifactWithSourceData) GetRegistrySourceOk() (*RegistrySource, bool)`

GetRegistrySourceOk returns a tuple with the RegistrySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrySource

`func (o *ArtifactWithSourceData) SetRegistrySource(v RegistrySource)`

SetRegistrySource sets RegistrySource field to given value.


### GetGithubSource

`func (o *ArtifactWithSourceData) GetGithubSource() GitHubSource`

GetGithubSource returns the GithubSource field if non-nil, zero value otherwise.

### GetGithubSourceOk

`func (o *ArtifactWithSourceData) GetGithubSourceOk() (*GitHubSource, bool)`

GetGithubSourceOk returns a tuple with the GithubSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSource

`func (o *ArtifactWithSourceData) SetGithubSource(v GitHubSource)`

SetGithubSource sets GithubSource field to given value.


### GetGithubSourceData

`func (o *ArtifactWithSourceData) GetGithubSourceData() GitHubSourceData`

GetGithubSourceData returns the GithubSourceData field if non-nil, zero value otherwise.

### GetGithubSourceDataOk

`func (o *ArtifactWithSourceData) GetGithubSourceDataOk() (*GitHubSourceData, bool)`

GetGithubSourceDataOk returns a tuple with the GithubSourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSourceData

`func (o *ArtifactWithSourceData) SetGithubSourceData(v GitHubSourceData)`

SetGithubSourceData sets GithubSourceData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


