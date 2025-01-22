# DeploymentLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**Level** | **string** |  | 
**Message** | **string** |  | 
**Code** | **string** |  | 

## Methods

### NewDeploymentLog

`func NewDeploymentLog(timestamp time.Time, level string, message string, code string, ) *DeploymentLog`

NewDeploymentLog instantiates a new DeploymentLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentLogWithDefaults

`func NewDeploymentLogWithDefaults() *DeploymentLog`

NewDeploymentLogWithDefaults instantiates a new DeploymentLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *DeploymentLog) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeploymentLog) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeploymentLog) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetLevel

`func (o *DeploymentLog) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *DeploymentLog) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *DeploymentLog) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *DeploymentLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeploymentLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeploymentLog) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *DeploymentLog) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeploymentLog) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeploymentLog) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


