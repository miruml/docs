# GitHubCommitter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Object** | Pointer to **string** |  | [optional] 
**Login** | **string** |  | 
**Type** | **string** |  | 
**HtmlUrl** | **string** |  | 
**AvatarUrl** | **string** |  | 

## Methods

### NewGitHubCommitter

`func NewGitHubCommitter(id int64, login string, type_ string, htmlUrl string, avatarUrl string, ) *GitHubCommitter`

NewGitHubCommitter instantiates a new GitHubCommitter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubCommitterWithDefaults

`func NewGitHubCommitterWithDefaults() *GitHubCommitter`

NewGitHubCommitterWithDefaults instantiates a new GitHubCommitter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitHubCommitter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitHubCommitter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitHubCommitter) SetId(v int64)`

SetId sets Id field to given value.


### GetObject

`func (o *GitHubCommitter) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GitHubCommitter) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GitHubCommitter) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GitHubCommitter) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetLogin

`func (o *GitHubCommitter) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GitHubCommitter) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GitHubCommitter) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetType

`func (o *GitHubCommitter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitHubCommitter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitHubCommitter) SetType(v string)`

SetType sets Type field to given value.


### GetHtmlUrl

`func (o *GitHubCommitter) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GitHubCommitter) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GitHubCommitter) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetAvatarUrl

`func (o *GitHubCommitter) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GitHubCommitter) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GitHubCommitter) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


