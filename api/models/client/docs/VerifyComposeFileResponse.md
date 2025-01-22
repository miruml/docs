# VerifyComposeFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Content** | **string** |  | 
**IsValid** | **bool** |  | 
**IsSchemaValid** | **bool** |  | 
**SchemaError** | Pointer to **string** |  | [optional] 
**Images** | [**ComposeFileImageList**](ComposeFileImageList.md) |  | 

## Methods

### NewVerifyComposeFileResponse

`func NewVerifyComposeFileResponse(object string, content string, isValid bool, isSchemaValid bool, images ComposeFileImageList, ) *VerifyComposeFileResponse`

NewVerifyComposeFileResponse instantiates a new VerifyComposeFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyComposeFileResponseWithDefaults

`func NewVerifyComposeFileResponseWithDefaults() *VerifyComposeFileResponse`

NewVerifyComposeFileResponseWithDefaults instantiates a new VerifyComposeFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *VerifyComposeFileResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VerifyComposeFileResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VerifyComposeFileResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetContent

`func (o *VerifyComposeFileResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VerifyComposeFileResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VerifyComposeFileResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetIsValid

`func (o *VerifyComposeFileResponse) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *VerifyComposeFileResponse) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *VerifyComposeFileResponse) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


### GetIsSchemaValid

`func (o *VerifyComposeFileResponse) GetIsSchemaValid() bool`

GetIsSchemaValid returns the IsSchemaValid field if non-nil, zero value otherwise.

### GetIsSchemaValidOk

`func (o *VerifyComposeFileResponse) GetIsSchemaValidOk() (*bool, bool)`

GetIsSchemaValidOk returns a tuple with the IsSchemaValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSchemaValid

`func (o *VerifyComposeFileResponse) SetIsSchemaValid(v bool)`

SetIsSchemaValid sets IsSchemaValid field to given value.


### GetSchemaError

`func (o *VerifyComposeFileResponse) GetSchemaError() string`

GetSchemaError returns the SchemaError field if non-nil, zero value otherwise.

### GetSchemaErrorOk

`func (o *VerifyComposeFileResponse) GetSchemaErrorOk() (*string, bool)`

GetSchemaErrorOk returns a tuple with the SchemaError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaError

`func (o *VerifyComposeFileResponse) SetSchemaError(v string)`

SetSchemaError sets SchemaError field to given value.

### HasSchemaError

`func (o *VerifyComposeFileResponse) HasSchemaError() bool`

HasSchemaError returns a boolean if a field has been set.

### GetImages

`func (o *VerifyComposeFileResponse) GetImages() ComposeFileImageList`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *VerifyComposeFileResponse) GetImagesOk() (*ComposeFileImageList, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *VerifyComposeFileResponse) SetImages(v ComposeFileImageList)`

SetImages sets Images field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


