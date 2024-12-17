# GroupDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Hardware** | **string** |  | 
**OperatingSystem** | **NullableString** |  | 
**Architecture** | **NullableString** |  | 
**Status** | **string** |  | 
**LastReportedStatus** | **string** |  | 
**MiruVersion** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**SyncedAt** | **time.Time** |  | 
**Containers** | [**GroupDeviceContainerList**](GroupDeviceContainerList.md) |  | 
**Deployments** | [**GroupDeviceDeploymentList**](GroupDeviceDeploymentList.md) |  | 

## Methods

### NewGroupDevice

`func NewGroupDevice(object string, id string, name string, hardware string, operatingSystem NullableString, architecture NullableString, status string, lastReportedStatus string, miruVersion NullableString, createdAt time.Time, syncedAt time.Time, containers GroupDeviceContainerList, deployments GroupDeviceDeploymentList, ) *GroupDevice`

NewGroupDevice instantiates a new GroupDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDeviceWithDefaults

`func NewGroupDeviceWithDefaults() *GroupDevice`

NewGroupDeviceWithDefaults instantiates a new GroupDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *GroupDevice) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GroupDevice) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GroupDevice) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *GroupDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupDevice) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *GroupDevice) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GroupDevice) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GroupDevice) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *GroupDevice) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetName

`func (o *GroupDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupDevice) SetName(v string)`

SetName sets Name field to given value.


### GetHardware

`func (o *GroupDevice) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *GroupDevice) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *GroupDevice) SetHardware(v string)`

SetHardware sets Hardware field to given value.


### GetOperatingSystem

`func (o *GroupDevice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *GroupDevice) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *GroupDevice) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### SetOperatingSystemNil

`func (o *GroupDevice) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *GroupDevice) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetArchitecture

`func (o *GroupDevice) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *GroupDevice) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *GroupDevice) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### SetArchitectureNil

`func (o *GroupDevice) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *GroupDevice) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetStatus

`func (o *GroupDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupDevice) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastReportedStatus

`func (o *GroupDevice) GetLastReportedStatus() string`

GetLastReportedStatus returns the LastReportedStatus field if non-nil, zero value otherwise.

### GetLastReportedStatusOk

`func (o *GroupDevice) GetLastReportedStatusOk() (*string, bool)`

GetLastReportedStatusOk returns a tuple with the LastReportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedStatus

`func (o *GroupDevice) SetLastReportedStatus(v string)`

SetLastReportedStatus sets LastReportedStatus field to given value.


### GetMiruVersion

`func (o *GroupDevice) GetMiruVersion() string`

GetMiruVersion returns the MiruVersion field if non-nil, zero value otherwise.

### GetMiruVersionOk

`func (o *GroupDevice) GetMiruVersionOk() (*string, bool)`

GetMiruVersionOk returns a tuple with the MiruVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiruVersion

`func (o *GroupDevice) SetMiruVersion(v string)`

SetMiruVersion sets MiruVersion field to given value.


### SetMiruVersionNil

`func (o *GroupDevice) SetMiruVersionNil(b bool)`

 SetMiruVersionNil sets the value for MiruVersion to be an explicit nil

### UnsetMiruVersion
`func (o *GroupDevice) UnsetMiruVersion()`

UnsetMiruVersion ensures that no value is present for MiruVersion, not even an explicit nil
### GetCreatedAt

`func (o *GroupDevice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupDevice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupDevice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSyncedAt

`func (o *GroupDevice) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *GroupDevice) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *GroupDevice) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.


### GetContainers

`func (o *GroupDevice) GetContainers() GroupDeviceContainerList`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GroupDevice) GetContainersOk() (*GroupDeviceContainerList, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GroupDevice) SetContainers(v GroupDeviceContainerList)`

SetContainers sets Containers field to given value.


### GetDeployments

`func (o *GroupDevice) GetDeployments() GroupDeviceDeploymentList`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *GroupDevice) GetDeploymentsOk() (*GroupDeviceDeploymentList, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *GroupDevice) SetDeployments(v GroupDeviceDeploymentList)`

SetDeployments sets Deployments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


