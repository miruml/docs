# RegistrySourceContainerRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**RegistryUrl** | **string** |  | 
**AccountLogin** | **string** |  | 
**Name** | **string** |  | 
**Uri** | **string** |  | 
**Type** | [**ContainerRepositoryType**](ContainerRepositoryType.md) |  | 
**IsExtra** | **bool** |  | 
**ComposeService** | **NullableString** |  | 

## Methods

### NewRegistrySourceContainerRepository

`func NewRegistrySourceContainerRepository(object string, id string, registryUrl string, accountLogin string, name string, uri string, type_ ContainerRepositoryType, isExtra bool, composeService NullableString, ) *RegistrySourceContainerRepository`

NewRegistrySourceContainerRepository instantiates a new RegistrySourceContainerRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrySourceContainerRepositoryWithDefaults

`func NewRegistrySourceContainerRepositoryWithDefaults() *RegistrySourceContainerRepository`

NewRegistrySourceContainerRepositoryWithDefaults instantiates a new RegistrySourceContainerRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *RegistrySourceContainerRepository) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RegistrySourceContainerRepository) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RegistrySourceContainerRepository) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *RegistrySourceContainerRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistrySourceContainerRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistrySourceContainerRepository) SetId(v string)`

SetId sets Id field to given value.


### GetRegistryUrl

`func (o *RegistrySourceContainerRepository) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *RegistrySourceContainerRepository) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *RegistrySourceContainerRepository) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.


### GetAccountLogin

`func (o *RegistrySourceContainerRepository) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *RegistrySourceContainerRepository) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *RegistrySourceContainerRepository) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.


### GetName

`func (o *RegistrySourceContainerRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistrySourceContainerRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistrySourceContainerRepository) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *RegistrySourceContainerRepository) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RegistrySourceContainerRepository) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RegistrySourceContainerRepository) SetUri(v string)`

SetUri sets Uri field to given value.


### GetType

`func (o *RegistrySourceContainerRepository) GetType() ContainerRepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistrySourceContainerRepository) GetTypeOk() (*ContainerRepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistrySourceContainerRepository) SetType(v ContainerRepositoryType)`

SetType sets Type field to given value.


### GetIsExtra

`func (o *RegistrySourceContainerRepository) GetIsExtra() bool`

GetIsExtra returns the IsExtra field if non-nil, zero value otherwise.

### GetIsExtraOk

`func (o *RegistrySourceContainerRepository) GetIsExtraOk() (*bool, bool)`

GetIsExtraOk returns a tuple with the IsExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExtra

`func (o *RegistrySourceContainerRepository) SetIsExtra(v bool)`

SetIsExtra sets IsExtra field to given value.


### GetComposeService

`func (o *RegistrySourceContainerRepository) GetComposeService() string`

GetComposeService returns the ComposeService field if non-nil, zero value otherwise.

### GetComposeServiceOk

`func (o *RegistrySourceContainerRepository) GetComposeServiceOk() (*string, bool)`

GetComposeServiceOk returns a tuple with the ComposeService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeService

`func (o *RegistrySourceContainerRepository) SetComposeService(v string)`

SetComposeService sets ComposeService field to given value.


### SetComposeServiceNil

`func (o *RegistrySourceContainerRepository) SetComposeServiceNil(b bool)`

 SetComposeServiceNil sets the value for ComposeService to be an explicit nil

### UnsetComposeService
`func (o *RegistrySourceContainerRepository) UnsetComposeService()`

UnsetComposeService ensures that no value is present for ComposeService, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


