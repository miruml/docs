# GitHubCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Sha** | **string** |  | 
**Message** | **string** |  | 
**HtmlUrl** | **string** |  | 
**PushedAt** | **time.Time** |  | 
**Committer** | [**GitHubCommitter**](GitHubCommitter.md) |  | 
**IsBuilt** | **bool** |  | 
**Artifacts** | Pointer to [**GitHubCommitArtifactList**](GitHubCommitArtifactList.md) |  | [optional] 

## Methods

### NewGitHubCommit

`func NewGitHubCommit(object string, sha string, message string, htmlUrl string, pushedAt time.Time, committer GitHubCommitter, isBuilt bool, ) *GitHubCommit`

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


### GetPushedAt

`func (o *GitHubCommit) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *GitHubCommit) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *GitHubCommit) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.


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


### GetIsBuilt

`func (o *GitHubCommit) GetIsBuilt() bool`

GetIsBuilt returns the IsBuilt field if non-nil, zero value otherwise.

### GetIsBuiltOk

`func (o *GitHubCommit) GetIsBuiltOk() (*bool, bool)`

GetIsBuiltOk returns a tuple with the IsBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuilt

`func (o *GitHubCommit) SetIsBuilt(v bool)`

SetIsBuilt sets IsBuilt field to given value.


### GetArtifacts

`func (o *GitHubCommit) GetArtifacts() GitHubCommitArtifactList`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *GitHubCommit) GetArtifactsOk() (*GitHubCommitArtifactList, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *GitHubCommit) SetArtifacts(v GitHubCommitArtifactList)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *GitHubCommit) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


