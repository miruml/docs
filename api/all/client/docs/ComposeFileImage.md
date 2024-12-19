# ComposeFileImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Digest** | **string** |  | 
**Tags** | **[]string** |  | 
**Uri** | **string** |  | 
**Bytes** | **NullableInt64** |  | 
**UploadedAt** | **time.Time** |  | 
**IsValid** | **bool** |  | 

## Methods

### NewComposeFileImage

`func NewComposeFileImage(object string, digest string, tags []string, uri string, bytes NullableInt64, uploadedAt time.Time, isValid bool, ) *ComposeFileImage`

NewComposeFileImage instantiates a new ComposeFileImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeFileImageWithDefaults

`func NewComposeFileImageWithDefaults() *ComposeFileImage`

NewComposeFileImageWithDefaults instantiates a new ComposeFileImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ComposeFileImage) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ComposeFileImage) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ComposeFileImage) SetObject(v string)`

SetObject sets Object field to given value.


### GetDigest

`func (o *ComposeFileImage) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ComposeFileImage) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ComposeFileImage) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetTags

`func (o *ComposeFileImage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ComposeFileImage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ComposeFileImage) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUri

`func (o *ComposeFileImage) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ComposeFileImage) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ComposeFileImage) SetUri(v string)`

SetUri sets Uri field to given value.


### GetBytes

`func (o *ComposeFileImage) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ComposeFileImage) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ComposeFileImage) SetBytes(v int64)`

SetBytes sets Bytes field to given value.


### SetBytesNil

`func (o *ComposeFileImage) SetBytesNil(b bool)`

 SetBytesNil sets the value for Bytes to be an explicit nil

### UnsetBytes
`func (o *ComposeFileImage) UnsetBytes()`

UnsetBytes ensures that no value is present for Bytes, not even an explicit nil
### GetUploadedAt

`func (o *ComposeFileImage) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *ComposeFileImage) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *ComposeFileImage) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.


### GetIsValid

`func (o *ComposeFileImage) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ComposeFileImage) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ComposeFileImage) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


