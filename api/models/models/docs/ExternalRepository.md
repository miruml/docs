# ExternalRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**RegistryUrl** | **string** |  | 
**Name** | **string** |  | 
**Uri** | **string** |  | 

## Methods

### NewExternalRepository

`func NewExternalRepository(registryUrl string, name string, uri string, ) *ExternalRepository`

NewExternalRepository instantiates a new ExternalRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalRepositoryWithDefaults

`func NewExternalRepositoryWithDefaults() *ExternalRepository`

NewExternalRepositoryWithDefaults instantiates a new ExternalRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ExternalRepository) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ExternalRepository) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ExternalRepository) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ExternalRepository) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRegistryUrl

`func (o *ExternalRepository) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *ExternalRepository) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *ExternalRepository) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.


### GetName

`func (o *ExternalRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalRepository) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *ExternalRepository) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ExternalRepository) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ExternalRepository) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


