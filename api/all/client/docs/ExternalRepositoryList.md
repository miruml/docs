# ExternalRepositoryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]ExternalRepository**](ExternalRepository.md) |  | [optional] 

## Methods

### NewExternalRepositoryList

`func NewExternalRepositoryList() *ExternalRepositoryList`

NewExternalRepositoryList instantiates a new ExternalRepositoryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalRepositoryListWithDefaults

`func NewExternalRepositoryListWithDefaults() *ExternalRepositoryList`

NewExternalRepositoryListWithDefaults instantiates a new ExternalRepositoryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ExternalRepositoryList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ExternalRepositoryList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ExternalRepositoryList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ExternalRepositoryList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *ExternalRepositoryList) GetData() []ExternalRepository`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalRepositoryList) GetDataOk() (*[]ExternalRepository, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalRepositoryList) SetData(v []ExternalRepository)`

SetData sets Data field to given value.

### HasData

`func (o *ExternalRepositoryList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


