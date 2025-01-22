# UnauthorizedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Params** | **map[string]interface{}** |  | 
**Message** | **string** |  | 

## Methods

### NewUnauthorizedError

`func NewUnauthorizedError(code string, params map[string]interface{}, message string, ) *UnauthorizedError`

NewUnauthorizedError instantiates a new UnauthorizedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorWithDefaults

`func NewUnauthorizedErrorWithDefaults() *UnauthorizedError`

NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnauthorizedError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnauthorizedError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnauthorizedError) SetCode(v string)`

SetCode sets Code field to given value.


### GetParams

`func (o *UnauthorizedError) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *UnauthorizedError) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *UnauthorizedError) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.


### GetMessage

`func (o *UnauthorizedError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthorizedError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthorizedError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


