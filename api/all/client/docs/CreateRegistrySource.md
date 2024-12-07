# CreateRegistrySource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ComposeFile** | **string** |  | 
**ExtraRepositories** | **[]string** |  | 
**Aarch64** | **bool** |  | 
**X8664** | **bool** |  | 

## Methods

### NewCreateRegistrySource

`func NewCreateRegistrySource(name string, composeFile string, extraRepositories []string, aarch64 bool, x8664 bool, ) *CreateRegistrySource`

NewCreateRegistrySource instantiates a new CreateRegistrySource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRegistrySourceWithDefaults

`func NewCreateRegistrySourceWithDefaults() *CreateRegistrySource`

NewCreateRegistrySourceWithDefaults instantiates a new CreateRegistrySource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRegistrySource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRegistrySource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRegistrySource) SetName(v string)`

SetName sets Name field to given value.


### GetComposeFile

`func (o *CreateRegistrySource) GetComposeFile() string`

GetComposeFile returns the ComposeFile field if non-nil, zero value otherwise.

### GetComposeFileOk

`func (o *CreateRegistrySource) GetComposeFileOk() (*string, bool)`

GetComposeFileOk returns a tuple with the ComposeFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeFile

`func (o *CreateRegistrySource) SetComposeFile(v string)`

SetComposeFile sets ComposeFile field to given value.


### GetExtraRepositories

`func (o *CreateRegistrySource) GetExtraRepositories() []string`

GetExtraRepositories returns the ExtraRepositories field if non-nil, zero value otherwise.

### GetExtraRepositoriesOk

`func (o *CreateRegistrySource) GetExtraRepositoriesOk() (*[]string, bool)`

GetExtraRepositoriesOk returns a tuple with the ExtraRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRepositories

`func (o *CreateRegistrySource) SetExtraRepositories(v []string)`

SetExtraRepositories sets ExtraRepositories field to given value.


### GetAarch64

`func (o *CreateRegistrySource) GetAarch64() bool`

GetAarch64 returns the Aarch64 field if non-nil, zero value otherwise.

### GetAarch64Ok

`func (o *CreateRegistrySource) GetAarch64Ok() (*bool, bool)`

GetAarch64Ok returns a tuple with the Aarch64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAarch64

`func (o *CreateRegistrySource) SetAarch64(v bool)`

SetAarch64 sets Aarch64 field to given value.


### GetX8664

`func (o *CreateRegistrySource) GetX8664() bool`

GetX8664 returns the X8664 field if non-nil, zero value otherwise.

### GetX8664Ok

`func (o *CreateRegistrySource) GetX8664Ok() (*bool, bool)`

GetX8664Ok returns a tuple with the X8664 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX8664

`func (o *CreateRegistrySource) SetX8664(v bool)`

SetX8664 sets X8664 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


