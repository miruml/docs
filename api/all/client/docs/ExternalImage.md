# ExternalImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Digest** | **string** |  | 
**Tags** | **[]string** |  | 
**Uri** | **string** |  | 
**Bytes** | **int32** |  | 
**UploadedAt** | **string** |  | 

## Methods

### NewExternalImage

`func NewExternalImage(digest string, tags []string, uri string, bytes int32, uploadedAt string, ) *ExternalImage`

NewExternalImage instantiates a new ExternalImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalImageWithDefaults

`func NewExternalImageWithDefaults() *ExternalImage`

NewExternalImageWithDefaults instantiates a new ExternalImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ExternalImage) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ExternalImage) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ExternalImage) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ExternalImage) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDigest

`func (o *ExternalImage) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ExternalImage) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ExternalImage) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetTags

`func (o *ExternalImage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExternalImage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExternalImage) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUri

`func (o *ExternalImage) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ExternalImage) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ExternalImage) SetUri(v string)`

SetUri sets Uri field to given value.


### GetBytes

`func (o *ExternalImage) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ExternalImage) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ExternalImage) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetUploadedAt

`func (o *ExternalImage) GetUploadedAt() string`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *ExternalImage) GetUploadedAtOk() (*string, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *ExternalImage) SetUploadedAt(v string)`

SetUploadedAt sets UploadedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


