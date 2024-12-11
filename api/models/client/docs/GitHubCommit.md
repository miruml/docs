# GitHubCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**Committer** | Pointer to [**GitHubCommitter**](GitHubCommitter.md) |  | [optional] 

## Methods

### NewGitHubCommit

`func NewGitHubCommit() *GitHubCommit`

NewGitHubCommit instantiates a new GitHubCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubCommitWithDefaults

`func NewGitHubCommitWithDefaults() *GitHubCommit`

NewGitHubCommitWithDefaults instantiates a new GitHubCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *GitHubCommit) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GitHubCommit) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GitHubCommit) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GitHubCommit) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSha

`func (o *GitHubCommit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitHubCommit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitHubCommit) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitHubCommit) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetMessage

`func (o *GitHubCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitHubCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitHubCommit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GitHubCommit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GitHubCommit) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GitHubCommit) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GitHubCommit) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GitHubCommit) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetCommitter

`func (o *GitHubCommit) GetCommitter() GitHubCommitter`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *GitHubCommit) GetCommitterOk() (*GitHubCommitter, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *GitHubCommit) SetCommitter(v GitHubCommitter)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *GitHubCommit) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


