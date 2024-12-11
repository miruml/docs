# GitHubSourceCommits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**GitHubRepository**](GitHubRepository.md) |  | [optional] 
**Commits** | Pointer to [**GitHubCommitList**](GitHubCommitList.md) |  | [optional] 

## Methods

### NewGitHubSourceCommits

`func NewGitHubSourceCommits() *GitHubSourceCommits`

NewGitHubSourceCommits instantiates a new GitHubSourceCommits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubSourceCommitsWithDefaults

`func NewGitHubSourceCommitsWithDefaults() *GitHubSourceCommits`

NewGitHubSourceCommitsWithDefaults instantiates a new GitHubSourceCommits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *GitHubSourceCommits) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GitHubSourceCommits) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GitHubSourceCommits) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GitHubSourceCommits) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRepository

`func (o *GitHubSourceCommits) GetRepository() GitHubRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GitHubSourceCommits) GetRepositoryOk() (*GitHubRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GitHubSourceCommits) SetRepository(v GitHubRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GitHubSourceCommits) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetCommits

`func (o *GitHubSourceCommits) GetCommits() GitHubCommitList`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GitHubSourceCommits) GetCommitsOk() (*GitHubCommitList, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GitHubSourceCommits) SetCommits(v GitHubCommitList)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *GitHubSourceCommits) HasCommits() bool`

HasCommits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


