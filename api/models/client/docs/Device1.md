# Device1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Hardware** | **string** |  | 
**OperatingSystem** | **NullableString** |  | 
**Architecture** | **NullableString** |  | 
**Status** | **string** |  | 
**LastReportedStatus** | **string** |  | 
**MiruVersion** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**SyncedAt** | **time.Time** |  | 
**Containers** | Pointer to [**Device1AllOfContainers**](Device1AllOfContainers.md) |  | [optional] 
**Deployments** | Pointer to [**Device1AllOfDeployments**](Device1AllOfDeployments.md) |  | [optional] 

## Methods

### NewDevice1

`func NewDevice1(object string, id string, name string, hardware string, operatingSystem NullableString, architecture NullableString, status string, lastReportedStatus string, miruVersion NullableString, createdAt time.Time, syncedAt time.Time, ) *Device1`

NewDevice1 instantiates a new Device1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevice1WithDefaults

`func NewDevice1WithDefaults() *Device1`

NewDevice1WithDefaults instantiates a new Device1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Device1) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Device1) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Device1) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Device1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Device1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device1) SetName(v string)`

SetName sets Name field to given value.


### GetHardware

`func (o *Device1) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *Device1) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *Device1) SetHardware(v string)`

SetHardware sets Hardware field to given value.


### GetOperatingSystem

`func (o *Device1) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Device1) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Device1) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### SetOperatingSystemNil

`func (o *Device1) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *Device1) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetArchitecture

`func (o *Device1) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Device1) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Device1) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### SetArchitectureNil

`func (o *Device1) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *Device1) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetStatus

`func (o *Device1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Device1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Device1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastReportedStatus

`func (o *Device1) GetLastReportedStatus() string`

GetLastReportedStatus returns the LastReportedStatus field if non-nil, zero value otherwise.

### GetLastReportedStatusOk

`func (o *Device1) GetLastReportedStatusOk() (*string, bool)`

GetLastReportedStatusOk returns a tuple with the LastReportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedStatus

`func (o *Device1) SetLastReportedStatus(v string)`

SetLastReportedStatus sets LastReportedStatus field to given value.


### GetMiruVersion

`func (o *Device1) GetMiruVersion() string`

GetMiruVersion returns the MiruVersion field if non-nil, zero value otherwise.

### GetMiruVersionOk

`func (o *Device1) GetMiruVersionOk() (*string, bool)`

GetMiruVersionOk returns a tuple with the MiruVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiruVersion

`func (o *Device1) SetMiruVersion(v string)`

SetMiruVersion sets MiruVersion field to given value.


### SetMiruVersionNil

`func (o *Device1) SetMiruVersionNil(b bool)`

 SetMiruVersionNil sets the value for MiruVersion to be an explicit nil

### UnsetMiruVersion
`func (o *Device1) UnsetMiruVersion()`

UnsetMiruVersion ensures that no value is present for MiruVersion, not even an explicit nil
### GetCreatedAt

`func (o *Device1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Device1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Device1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSyncedAt

`func (o *Device1) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *Device1) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *Device1) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.


### GetContainers

`func (o *Device1) GetContainers() Device1AllOfContainers`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Device1) GetContainersOk() (*Device1AllOfContainers, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Device1) SetContainers(v Device1AllOfContainers)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *Device1) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetDeployments

`func (o *Device1) GetDeployments() Device1AllOfDeployments`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *Device1) GetDeploymentsOk() (*Device1AllOfDeployments, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *Device1) SetDeployments(v Device1AllOfDeployments)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *Device1) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


