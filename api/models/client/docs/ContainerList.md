# ContainerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]Container**](Container.md) |  | [optional] 

## Methods

### NewContainerList

`func NewContainerList() *ContainerList`

NewContainerList instantiates a new ContainerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerListWithDefaults

`func NewContainerListWithDefaults() *ContainerList`

NewContainerListWithDefaults instantiates a new ContainerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ContainerList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ContainerList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ContainerList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ContainerList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *ContainerList) GetData() []Container`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContainerList) GetDataOk() (*[]Container, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContainerList) SetData(v []Container)`

SetData sets Data field to given value.

### HasData

`func (o *ContainerList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

