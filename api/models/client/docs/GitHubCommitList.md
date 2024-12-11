# GitHubCommitList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GitHubCommit**](GitHubCommit.md) |  | [optional] 

## Methods

### NewGitHubCommitList

`func NewGitHubCommitList() *GitHubCommitList`

NewGitHubCommitList instantiates a new GitHubCommitList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubCommitListWithDefaults

`func NewGitHubCommitListWithDefaults() *GitHubCommitList`

NewGitHubCommitListWithDefaults instantiates a new GitHubCommitList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *GitHubCommitList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GitHubCommitList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GitHubCommitList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GitHubCommitList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *GitHubCommitList) GetData() []GitHubCommit`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GitHubCommitList) GetDataOk() (*[]GitHubCommit, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GitHubCommitList) SetData(v []GitHubCommit)`

SetData sets Data field to given value.

### HasData

`func (o *GitHubCommitList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


