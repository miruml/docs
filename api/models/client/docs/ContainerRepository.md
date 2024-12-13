# ContainerRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**RegistryUrl** | **string** |  | 
**Name** | **string** |  | 
**Uri** | **string** |  | 
**Type** | [**ContainerRepositoryType**](ContainerRepositoryType.md) |  | 

## Methods

### NewContainerRepository

`func NewContainerRepository(object string, id string, registryUrl string, name string, uri string, type_ ContainerRepositoryType, ) *ContainerRepository`

NewContainerRepository instantiates a new ContainerRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRepositoryWithDefaults

`func NewContainerRepositoryWithDefaults() *ContainerRepository`

NewContainerRepositoryWithDefaults instantiates a new ContainerRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ContainerRepository) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ContainerRepository) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ContainerRepository) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *ContainerRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerRepository) SetId(v string)`

SetId sets Id field to given value.


### GetRegistryUrl

`func (o *ContainerRepository) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *ContainerRepository) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *ContainerRepository) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.


### GetName

`func (o *ContainerRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerRepository) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *ContainerRepository) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ContainerRepository) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ContainerRepository) SetUri(v string)`

SetUri sets Uri field to given value.


### GetType

`func (o *ContainerRepository) GetType() ContainerRepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerRepository) GetTypeOk() (*ContainerRepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerRepository) SetType(v ContainerRepositoryType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


