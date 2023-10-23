# GitRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** |  | [optional] 
**CommitId** | Pointer to **string** |  | [optional] 

## Methods

### NewGitRef

`func NewGitRef() *GitRef`

NewGitRef instantiates a new GitRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRefWithDefaults

`func NewGitRefWithDefaults() *GitRef`

NewGitRefWithDefaults instantiates a new GitRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *GitRef) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *GitRef) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *GitRef) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *GitRef) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetCommitId

`func (o *GitRef) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *GitRef) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *GitRef) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *GitRef) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


