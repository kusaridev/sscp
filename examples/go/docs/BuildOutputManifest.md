# BuildOutputManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outputs** | Pointer to [**map[string]BuildOutput**](BuildOutput.md) |  | [optional] 

## Methods

### NewBuildOutputManifest

`func NewBuildOutputManifest() *BuildOutputManifest`

NewBuildOutputManifest instantiates a new BuildOutputManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildOutputManifestWithDefaults

`func NewBuildOutputManifestWithDefaults() *BuildOutputManifest`

NewBuildOutputManifestWithDefaults instantiates a new BuildOutputManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputs

`func (o *BuildOutputManifest) GetOutputs() map[string]BuildOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *BuildOutputManifest) GetOutputsOk() (*map[string]BuildOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *BuildOutputManifest) SetOutputs(v map[string]BuildOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *BuildOutputManifest) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


