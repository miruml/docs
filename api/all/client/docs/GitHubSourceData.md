# GitHubSourceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | [**GitHubSourceDataSender**](GitHubSourceDataSender.md) |  | 
**Branch** | **string** |  | 
**BuildPath** | **string** |  | 
**Repository** | [**GitHubSourceDataRepository**](GitHubSourceDataRepository.md) |  | 
**HeadCommit** | [**GitHubSourceDataHeadCommit**](GitHubSourceDataHeadCommit.md) |  | 

## Methods

### NewGitHubSourceData

`func NewGitHubSourceData(sender GitHubSourceDataSender, branch string, buildPath string, repository GitHubSourceDataRepository, headCommit GitHubSourceDataHeadCommit, ) *GitHubSourceData`

NewGitHubSourceData instantiates a new GitHubSourceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubSourceDataWithDefaults

`func NewGitHubSourceDataWithDefaults() *GitHubSourceData`

NewGitHubSourceDataWithDefaults instantiates a new GitHubSourceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *GitHubSourceData) GetSender() GitHubSourceDataSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GitHubSourceData) GetSenderOk() (*GitHubSourceDataSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GitHubSourceData) SetSender(v GitHubSourceDataSender)`

SetSender sets Sender field to given value.


### GetBranch

`func (o *GitHubSourceData) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitHubSourceData) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitHubSourceData) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetBuildPath

`func (o *GitHubSourceData) GetBuildPath() string`

GetBuildPath returns the BuildPath field if non-nil, zero value otherwise.

### GetBuildPathOk

`func (o *GitHubSourceData) GetBuildPathOk() (*string, bool)`

GetBuildPathOk returns a tuple with the BuildPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildPath

`func (o *GitHubSourceData) SetBuildPath(v string)`

SetBuildPath sets BuildPath field to given value.


### GetRepository

`func (o *GitHubSourceData) GetRepository() GitHubSourceDataRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GitHubSourceData) GetRepositoryOk() (*GitHubSourceDataRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GitHubSourceData) SetRepository(v GitHubSourceDataRepository)`

SetRepository sets Repository field to given value.


### GetHeadCommit

`func (o *GitHubSourceData) GetHeadCommit() GitHubSourceDataHeadCommit`

GetHeadCommit returns the HeadCommit field if non-nil, zero value otherwise.

### GetHeadCommitOk

`func (o *GitHubSourceData) GetHeadCommitOk() (*GitHubSourceDataHeadCommit, bool)`

GetHeadCommitOk returns a tuple with the HeadCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommit

`func (o *GitHubSourceData) SetHeadCommit(v GitHubSourceDataHeadCommit)`

SetHeadCommit sets HeadCommit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


