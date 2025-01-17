# ComposeFileImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**ComposeReference** | **string** |  | 
**ImageUri** | **string** |  | 
**RepositoryUri** | **string** |  | 
**RegistryUrl** | **string** |  | 
**RegistryType** | **string** |  | 
**Name** | **string** |  | 
**Digest** | **string** |  | 
**Tag** | **string** |  | 
**IsImageValid** | **bool** |  | 
**IsRepositoryValid** | **bool** |  | 
**Error** | **string** |  | 

## Methods

### NewComposeFileImage

`func NewComposeFileImage(object string, composeReference string, imageUri string, repositoryUri string, registryUrl string, registryType string, name string, digest string, tag string, isImageValid bool, isRepositoryValid bool, error_ string, ) *ComposeFileImage`

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


### GetComposeReference

`func (o *ComposeFileImage) GetComposeReference() string`

GetComposeReference returns the ComposeReference field if non-nil, zero value otherwise.

### GetComposeReferenceOk

`func (o *ComposeFileImage) GetComposeReferenceOk() (*string, bool)`

GetComposeReferenceOk returns a tuple with the ComposeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeReference

`func (o *ComposeFileImage) SetComposeReference(v string)`

SetComposeReference sets ComposeReference field to given value.


### GetImageUri

`func (o *ComposeFileImage) GetImageUri() string`

GetImageUri returns the ImageUri field if non-nil, zero value otherwise.

### GetImageUriOk

`func (o *ComposeFileImage) GetImageUriOk() (*string, bool)`

GetImageUriOk returns a tuple with the ImageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUri

`func (o *ComposeFileImage) SetImageUri(v string)`

SetImageUri sets ImageUri field to given value.


### GetRepositoryUri

`func (o *ComposeFileImage) GetRepositoryUri() string`

GetRepositoryUri returns the RepositoryUri field if non-nil, zero value otherwise.

### GetRepositoryUriOk

`func (o *ComposeFileImage) GetRepositoryUriOk() (*string, bool)`

GetRepositoryUriOk returns a tuple with the RepositoryUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUri

`func (o *ComposeFileImage) SetRepositoryUri(v string)`

SetRepositoryUri sets RepositoryUri field to given value.


### GetRegistryUrl

`func (o *ComposeFileImage) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *ComposeFileImage) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *ComposeFileImage) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.


### GetRegistryType

`func (o *ComposeFileImage) GetRegistryType() string`

GetRegistryType returns the RegistryType field if non-nil, zero value otherwise.

### GetRegistryTypeOk

`func (o *ComposeFileImage) GetRegistryTypeOk() (*string, bool)`

GetRegistryTypeOk returns a tuple with the RegistryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryType

`func (o *ComposeFileImage) SetRegistryType(v string)`

SetRegistryType sets RegistryType field to given value.


### GetName

`func (o *ComposeFileImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComposeFileImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComposeFileImage) SetName(v string)`

SetName sets Name field to given value.


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


### GetTag

`func (o *ComposeFileImage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ComposeFileImage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ComposeFileImage) SetTag(v string)`

SetTag sets Tag field to given value.


### GetIsImageValid

`func (o *ComposeFileImage) GetIsImageValid() bool`

GetIsImageValid returns the IsImageValid field if non-nil, zero value otherwise.

### GetIsImageValidOk

`func (o *ComposeFileImage) GetIsImageValidOk() (*bool, bool)`

GetIsImageValidOk returns a tuple with the IsImageValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImageValid

`func (o *ComposeFileImage) SetIsImageValid(v bool)`

SetIsImageValid sets IsImageValid field to given value.


### GetIsRepositoryValid

`func (o *ComposeFileImage) GetIsRepositoryValid() bool`

GetIsRepositoryValid returns the IsRepositoryValid field if non-nil, zero value otherwise.

### GetIsRepositoryValidOk

`func (o *ComposeFileImage) GetIsRepositoryValidOk() (*bool, bool)`

GetIsRepositoryValidOk returns a tuple with the IsRepositoryValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepositoryValid

`func (o *ComposeFileImage) SetIsRepositoryValid(v bool)`

SetIsRepositoryValid sets IsRepositoryValid field to given value.


### GetError

`func (o *ComposeFileImage) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ComposeFileImage) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ComposeFileImage) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


