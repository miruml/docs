# ExternalContainerRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**RegistryUrl** | **string** |  | 
**AccountLogin** | **string** |  | 
**Name** | **string** |  | 
**Uri** | **string** |  | 
**Type** | [**ContainerRepositoryType**](ContainerRepositoryType.md) |  | 
**Description** | **NullableString** |  | 
**Bytes** | **NullableInt64** | The size of the repository in bytes | 

## Methods

### NewExternalContainerRepository

`func NewExternalContainerRepository(object string, registryUrl string, accountLogin string, name string, uri string, type_ ContainerRepositoryType, description NullableString, bytes NullableInt64, ) *ExternalContainerRepository`

NewExternalContainerRepository instantiates a new ExternalContainerRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalContainerRepositoryWithDefaults

`func NewExternalContainerRepositoryWithDefaults() *ExternalContainerRepository`

NewExternalContainerRepositoryWithDefaults instantiates a new ExternalContainerRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ExternalContainerRepository) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ExternalContainerRepository) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ExternalContainerRepository) SetObject(v string)`

SetObject sets Object field to given value.


### GetRegistryUrl

`func (o *ExternalContainerRepository) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *ExternalContainerRepository) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *ExternalContainerRepository) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.


### GetAccountLogin

`func (o *ExternalContainerRepository) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *ExternalContainerRepository) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *ExternalContainerRepository) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.


### GetName

`func (o *ExternalContainerRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalContainerRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalContainerRepository) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *ExternalContainerRepository) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ExternalContainerRepository) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ExternalContainerRepository) SetUri(v string)`

SetUri sets Uri field to given value.


### GetType

`func (o *ExternalContainerRepository) GetType() ContainerRepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalContainerRepository) GetTypeOk() (*ContainerRepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalContainerRepository) SetType(v ContainerRepositoryType)`

SetType sets Type field to given value.


### GetDescription

`func (o *ExternalContainerRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalContainerRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalContainerRepository) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ExternalContainerRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExternalContainerRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBytes

`func (o *ExternalContainerRepository) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ExternalContainerRepository) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ExternalContainerRepository) SetBytes(v int64)`

SetBytes sets Bytes field to given value.


### SetBytesNil

`func (o *ExternalContainerRepository) SetBytesNil(b bool)`

 SetBytesNil sets the value for Bytes to be an explicit nil

### UnsetBytes
`func (o *ExternalContainerRepository) UnsetBytes()`

UnsetBytes ensures that no value is present for Bytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


