# BaseDevice

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

## Methods

### NewBaseDevice

`func NewBaseDevice(object string, id string, name string, hardware string, operatingSystem NullableString, architecture NullableString, status string, lastReportedStatus string, miruVersion NullableString, createdAt time.Time, syncedAt time.Time, ) *BaseDevice`

NewBaseDevice instantiates a new BaseDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDeviceWithDefaults

`func NewBaseDeviceWithDefaults() *BaseDevice`

NewBaseDeviceWithDefaults instantiates a new BaseDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BaseDevice) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BaseDevice) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BaseDevice) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *BaseDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseDevice) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *BaseDevice) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BaseDevice) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BaseDevice) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BaseDevice) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetName

`func (o *BaseDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseDevice) SetName(v string)`

SetName sets Name field to given value.


### GetHardware

`func (o *BaseDevice) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *BaseDevice) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *BaseDevice) SetHardware(v string)`

SetHardware sets Hardware field to given value.


### GetOperatingSystem

`func (o *BaseDevice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *BaseDevice) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *BaseDevice) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### SetOperatingSystemNil

`func (o *BaseDevice) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *BaseDevice) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetArchitecture

`func (o *BaseDevice) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *BaseDevice) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *BaseDevice) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### SetArchitectureNil

`func (o *BaseDevice) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *BaseDevice) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetStatus

`func (o *BaseDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseDevice) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastReportedStatus

`func (o *BaseDevice) GetLastReportedStatus() string`

GetLastReportedStatus returns the LastReportedStatus field if non-nil, zero value otherwise.

### GetLastReportedStatusOk

`func (o *BaseDevice) GetLastReportedStatusOk() (*string, bool)`

GetLastReportedStatusOk returns a tuple with the LastReportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedStatus

`func (o *BaseDevice) SetLastReportedStatus(v string)`

SetLastReportedStatus sets LastReportedStatus field to given value.


### GetMiruVersion

`func (o *BaseDevice) GetMiruVersion() string`

GetMiruVersion returns the MiruVersion field if non-nil, zero value otherwise.

### GetMiruVersionOk

`func (o *BaseDevice) GetMiruVersionOk() (*string, bool)`

GetMiruVersionOk returns a tuple with the MiruVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiruVersion

`func (o *BaseDevice) SetMiruVersion(v string)`

SetMiruVersion sets MiruVersion field to given value.


### SetMiruVersionNil

`func (o *BaseDevice) SetMiruVersionNil(b bool)`

 SetMiruVersionNil sets the value for MiruVersion to be an explicit nil

### UnsetMiruVersion
`func (o *BaseDevice) UnsetMiruVersion()`

UnsetMiruVersion ensures that no value is present for MiruVersion, not even an explicit nil
### GetCreatedAt

`func (o *BaseDevice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseDevice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseDevice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSyncedAt

`func (o *BaseDevice) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *BaseDevice) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *BaseDevice) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


