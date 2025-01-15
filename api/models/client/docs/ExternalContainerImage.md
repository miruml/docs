# ExternalContainerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Digest** | **string** |  | 
**Tags** | **[]string** |  | 
**Uri** | **string** |  | 
**Bytes** | **int64** |  | 
**UploadedAt** | **time.Time** |  | 

## Methods

### NewExternalContainerImage

`func NewExternalContainerImage(object string, digest string, tags []string, uri string, bytes int64, uploadedAt time.Time, ) *ExternalContainerImage`

NewExternalContainerImage instantiates a new ExternalContainerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalContainerImageWithDefaults

`func NewExternalContainerImageWithDefaults() *ExternalContainerImage`

NewExternalContainerImageWithDefaults instantiates a new ExternalContainerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ExternalContainerImage) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ExternalContainerImage) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ExternalContainerImage) SetObject(v string)`

SetObject sets Object field to given value.


### GetDigest

`func (o *ExternalContainerImage) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ExternalContainerImage) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ExternalContainerImage) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetTags

`func (o *ExternalContainerImage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExternalContainerImage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExternalContainerImage) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUri

`func (o *ExternalContainerImage) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ExternalContainerImage) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ExternalContainerImage) SetUri(v string)`

SetUri sets Uri field to given value.


### GetBytes

`func (o *ExternalContainerImage) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ExternalContainerImage) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ExternalContainerImage) SetBytes(v int64)`

SetBytes sets Bytes field to given value.


### GetUploadedAt

`func (o *ExternalContainerImage) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *ExternalContainerImage) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *ExternalContainerImage) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


