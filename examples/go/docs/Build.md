# Build

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**GitRef** | Pointer to [**GitRef**](GitRef.md) |  | [optional] 
**BuildCommand** | Pointer to **[]string** |  | [optional] 
**BuildType** | Pointer to **string** |  | [optional] 

## Methods

### NewBuild

`func NewBuild() *Build`

NewBuild instantiates a new Build object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildWithDefaults

`func NewBuildWithDefaults() *Build`

NewBuildWithDefaults instantiates a new Build object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Build) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Build) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Build) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Build) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGitRef

`func (o *Build) GetGitRef() GitRef`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *Build) GetGitRefOk() (*GitRef, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *Build) SetGitRef(v GitRef)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *Build) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### GetBuildCommand

`func (o *Build) GetBuildCommand() []string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *Build) GetBuildCommandOk() (*[]string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *Build) SetBuildCommand(v []string)`

SetBuildCommand sets BuildCommand field to given value.

### HasBuildCommand

`func (o *Build) HasBuildCommand() bool`

HasBuildCommand returns a boolean if a field has been set.

### GetBuildType

`func (o *Build) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *Build) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *Build) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *Build) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


