# CreateRegistrySourceArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDuplicate** | **bool** |  | 
**Images** | [**[]CreateRegistrySourceArtifactImagesInner**](CreateRegistrySourceArtifactImagesInner.md) |  | 
**Groups** | **[]string** |  | 

## Methods

### NewCreateRegistrySourceArtifact

`func NewCreateRegistrySourceArtifact(allowDuplicate bool, images []CreateRegistrySourceArtifactImagesInner, groups []string, ) *CreateRegistrySourceArtifact`

NewCreateRegistrySourceArtifact instantiates a new CreateRegistrySourceArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRegistrySourceArtifactWithDefaults

`func NewCreateRegistrySourceArtifactWithDefaults() *CreateRegistrySourceArtifact`

NewCreateRegistrySourceArtifactWithDefaults instantiates a new CreateRegistrySourceArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDuplicate

`func (o *CreateRegistrySourceArtifact) GetAllowDuplicate() bool`

GetAllowDuplicate returns the AllowDuplicate field if non-nil, zero value otherwise.

### GetAllowDuplicateOk

`func (o *CreateRegistrySourceArtifact) GetAllowDuplicateOk() (*bool, bool)`

GetAllowDuplicateOk returns a tuple with the AllowDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicate

`func (o *CreateRegistrySourceArtifact) SetAllowDuplicate(v bool)`

SetAllowDuplicate sets AllowDuplicate field to given value.


### GetImages

`func (o *CreateRegistrySourceArtifact) GetImages() []CreateRegistrySourceArtifactImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CreateRegistrySourceArtifact) GetImagesOk() (*[]CreateRegistrySourceArtifactImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CreateRegistrySourceArtifact) SetImages(v []CreateRegistrySourceArtifactImagesInner)`

SetImages sets Images field to given value.


### GetGroups

`func (o *CreateRegistrySourceArtifact) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateRegistrySourceArtifact) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateRegistrySourceArtifact) SetGroups(v []string)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


