# VerifiedComposeFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Content** | **string** |  | 
**IsValid** | **bool** |  | 
**IsSchemaValid** | **bool** |  | 
**Images** | [**[]ComposeFileImageList**](ComposeFileImageList.md) |  | 

## Methods

### NewVerifiedComposeFileResponse

`func NewVerifiedComposeFileResponse(object string, content string, isValid bool, isSchemaValid bool, images []ComposeFileImageList, ) *VerifiedComposeFileResponse`

NewVerifiedComposeFileResponse instantiates a new VerifiedComposeFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifiedComposeFileResponseWithDefaults

`func NewVerifiedComposeFileResponseWithDefaults() *VerifiedComposeFileResponse`

NewVerifiedComposeFileResponseWithDefaults instantiates a new VerifiedComposeFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *VerifiedComposeFileResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VerifiedComposeFileResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VerifiedComposeFileResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetContent

`func (o *VerifiedComposeFileResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VerifiedComposeFileResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VerifiedComposeFileResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetIsValid

`func (o *VerifiedComposeFileResponse) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *VerifiedComposeFileResponse) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *VerifiedComposeFileResponse) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


### GetIsSchemaValid

`func (o *VerifiedComposeFileResponse) GetIsSchemaValid() bool`

GetIsSchemaValid returns the IsSchemaValid field if non-nil, zero value otherwise.

### GetIsSchemaValidOk

`func (o *VerifiedComposeFileResponse) GetIsSchemaValidOk() (*bool, bool)`

GetIsSchemaValidOk returns a tuple with the IsSchemaValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSchemaValid

`func (o *VerifiedComposeFileResponse) SetIsSchemaValid(v bool)`

SetIsSchemaValid sets IsSchemaValid field to given value.


### GetImages

`func (o *VerifiedComposeFileResponse) GetImages() []ComposeFileImageList`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *VerifiedComposeFileResponse) GetImagesOk() (*[]ComposeFileImageList, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *VerifiedComposeFileResponse) SetImages(v []ComposeFileImageList)`

SetImages sets Images field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


