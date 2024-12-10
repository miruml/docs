# RecordNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Params** | [**RecordNotFoundErrorParams**](RecordNotFoundErrorParams.md) |  | 
**Message** | **string** |  | 

## Methods

### NewRecordNotFoundError

`func NewRecordNotFoundError(code string, params RecordNotFoundErrorParams, message string, ) *RecordNotFoundError`

NewRecordNotFoundError instantiates a new RecordNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordNotFoundErrorWithDefaults

`func NewRecordNotFoundErrorWithDefaults() *RecordNotFoundError`

NewRecordNotFoundErrorWithDefaults instantiates a new RecordNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RecordNotFoundError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RecordNotFoundError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RecordNotFoundError) SetCode(v string)`

SetCode sets Code field to given value.


### GetParams

`func (o *RecordNotFoundError) GetParams() RecordNotFoundErrorParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *RecordNotFoundError) GetParamsOk() (*RecordNotFoundErrorParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *RecordNotFoundError) SetParams(v RecordNotFoundErrorParams)`

SetParams sets Params field to given value.


### GetMessage

`func (o *RecordNotFoundError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RecordNotFoundError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RecordNotFoundError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


