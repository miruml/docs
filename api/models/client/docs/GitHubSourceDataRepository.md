# GitHubSourceDataRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**HtmlUrl** | **string** |  | 
**CloneUrl** | **string** |  | 
**FullName** | **string** |  | 
**PushedAt** | **time.Time** |  | 

## Methods

### NewGitHubSourceDataRepository

`func NewGitHubSourceDataRepository(id int64, htmlUrl string, cloneUrl string, fullName string, pushedAt time.Time, ) *GitHubSourceDataRepository`

NewGitHubSourceDataRepository instantiates a new GitHubSourceDataRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubSourceDataRepositoryWithDefaults

`func NewGitHubSourceDataRepositoryWithDefaults() *GitHubSourceDataRepository`

NewGitHubSourceDataRepositoryWithDefaults instantiates a new GitHubSourceDataRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitHubSourceDataRepository) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitHubSourceDataRepository) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitHubSourceDataRepository) SetId(v int64)`

SetId sets Id field to given value.


### GetHtmlUrl

`func (o *GitHubSourceDataRepository) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GitHubSourceDataRepository) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GitHubSourceDataRepository) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCloneUrl

`func (o *GitHubSourceDataRepository) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *GitHubSourceDataRepository) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *GitHubSourceDataRepository) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.


### GetFullName

`func (o *GitHubSourceDataRepository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GitHubSourceDataRepository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GitHubSourceDataRepository) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetPushedAt

`func (o *GitHubSourceDataRepository) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *GitHubSourceDataRepository) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *GitHubSourceDataRepository) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


