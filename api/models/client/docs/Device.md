# Device

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
**Group** | [**BaseGroup**](BaseGroup.md) |  | 

## Methods

### NewDevice

`func NewDevice(object string, id string, name string, hardware string, operatingSystem NullableString, architecture NullableString, status string, lastReportedStatus string, miruVersion NullableString, createdAt time.Time, syncedAt time.Time, group BaseGroup, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Device) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Device) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Device) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *Device) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Device) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Device) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *Device) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.


### GetHardware

`func (o *Device) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *Device) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *Device) SetHardware(v string)`

SetHardware sets Hardware field to given value.


### GetOperatingSystem

`func (o *Device) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Device) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Device) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### SetOperatingSystemNil

`func (o *Device) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *Device) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetArchitecture

`func (o *Device) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Device) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Device) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### SetArchitectureNil

`func (o *Device) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *Device) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetStatus

`func (o *Device) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Device) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Device) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastReportedStatus

`func (o *Device) GetLastReportedStatus() string`

GetLastReportedStatus returns the LastReportedStatus field if non-nil, zero value otherwise.

### GetLastReportedStatusOk

`func (o *Device) GetLastReportedStatusOk() (*string, bool)`

GetLastReportedStatusOk returns a tuple with the LastReportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedStatus

`func (o *Device) SetLastReportedStatus(v string)`

SetLastReportedStatus sets LastReportedStatus field to given value.


### GetMiruVersion

`func (o *Device) GetMiruVersion() string`

GetMiruVersion returns the MiruVersion field if non-nil, zero value otherwise.

### GetMiruVersionOk

`func (o *Device) GetMiruVersionOk() (*string, bool)`

GetMiruVersionOk returns a tuple with the MiruVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiruVersion

`func (o *Device) SetMiruVersion(v string)`

SetMiruVersion sets MiruVersion field to given value.


### SetMiruVersionNil

`func (o *Device) SetMiruVersionNil(b bool)`

 SetMiruVersionNil sets the value for MiruVersion to be an explicit nil

### UnsetMiruVersion
`func (o *Device) UnsetMiruVersion()`

UnsetMiruVersion ensures that no value is present for MiruVersion, not even an explicit nil
### GetCreatedAt

`func (o *Device) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Device) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Device) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSyncedAt

`func (o *Device) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *Device) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *Device) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.


### GetGroup

`func (o *Device) GetGroup() BaseGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Device) GetGroupOk() (*BaseGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Device) SetGroup(v BaseGroup)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


