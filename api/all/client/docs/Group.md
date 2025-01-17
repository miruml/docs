# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Devices** | Pointer to [**GroupDeviceList**](GroupDeviceList.md) |  | [optional] 
**GithubSources** | Pointer to [**GitHubSourceList**](GitHubSourceList.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup(object string, id string, name string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Group) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Group) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Group) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *Group) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Group) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Group) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *Group) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetDevices

`func (o *Group) GetDevices() GroupDeviceList`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *Group) GetDevicesOk() (*GroupDeviceList, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *Group) SetDevices(v GroupDeviceList)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *Group) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetGithubSources

`func (o *Group) GetGithubSources() GitHubSourceList`

GetGithubSources returns the GithubSources field if non-nil, zero value otherwise.

### GetGithubSourcesOk

`func (o *Group) GetGithubSourcesOk() (*GitHubSourceList, bool)`

GetGithubSourcesOk returns a tuple with the GithubSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSources

`func (o *Group) SetGithubSources(v GitHubSourceList)`

SetGithubSources sets GithubSources field to given value.

### HasGithubSources

`func (o *Group) HasGithubSources() bool`

HasGithubSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


