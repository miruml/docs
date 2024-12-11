# GitHubRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Object** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**FullName** | **string** |  | 
**HtmlUrl** | **string** |  | 

## Methods

### NewGitHubRepository

`func NewGitHubRepository(id int64, name string, fullName string, htmlUrl string, ) *GitHubRepository`

NewGitHubRepository instantiates a new GitHubRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubRepositoryWithDefaults

`func NewGitHubRepositoryWithDefaults() *GitHubRepository`

NewGitHubRepositoryWithDefaults instantiates a new GitHubRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitHubRepository) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitHubRepository) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitHubRepository) SetId(v int64)`

SetId sets Id field to given value.


### GetObject

`func (o *GitHubRepository) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GitHubRepository) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GitHubRepository) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GitHubRepository) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *GitHubRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitHubRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitHubRepository) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *GitHubRepository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GitHubRepository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GitHubRepository) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetHtmlUrl

`func (o *GitHubRepository) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GitHubRepository) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GitHubRepository) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


