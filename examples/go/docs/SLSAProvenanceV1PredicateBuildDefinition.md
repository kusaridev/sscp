# SLSAProvenanceV1PredicateBuildDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | **string** |  | 
**ExternalParameters** | **map[string]interface{}** | The parameters that are under external control, such as those set by a user or tenant of the build platform. They MUST be complete at SLSA Build L3, meaning that there is no additional mechanism for an external party to influence the build. (At lower SLSA Build levels, the completeness MAY be best effort.)\\nThe build platform SHOULD be designed to minimize the size and complexity of externalParameters, in order to reduce fragility and ease verification. Consumers SHOULD have an expectation of what “good” looks like; the more information that they need to check, the harder that task becomes.\\nVerifiers SHOULD reject unrecognized or unexpected fields within externalParameters. | 
**InternalParameters** | Pointer to **map[string]interface{}** | Unordered collection of artifacts needed at build time. Completeness is best effort, at least through SLSA Build L3. For example, if the build script fetches and executes “example.com/foo.sh”, which in turn fetches “example.com/bar.tar.gz”, then both “foo.sh” and “bar.tar.gz” SHOULD be listed here. | [optional] 
**ResolvedDependencies** | Pointer to [**[]SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner**](SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner.md) | Unordered collection of artifacts needed at build time. Completeness is best effort, at least through SLSA Build L3. For example, if the build script fetches and executes “example.com/foo.sh”, which in turn fetches “example.com/bar.tar.gz”, then both “foo.sh” and “bar.tar.gz” SHOULD be listed here. | [optional] 

## Methods

### NewSLSAProvenanceV1PredicateBuildDefinition

`func NewSLSAProvenanceV1PredicateBuildDefinition(buildType string, externalParameters map[string]interface{}, ) *SLSAProvenanceV1PredicateBuildDefinition`

NewSLSAProvenanceV1PredicateBuildDefinition instantiates a new SLSAProvenanceV1PredicateBuildDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLSAProvenanceV1PredicateBuildDefinitionWithDefaults

`func NewSLSAProvenanceV1PredicateBuildDefinitionWithDefaults() *SLSAProvenanceV1PredicateBuildDefinition`

NewSLSAProvenanceV1PredicateBuildDefinitionWithDefaults instantiates a new SLSAProvenanceV1PredicateBuildDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *SLSAProvenanceV1PredicateBuildDefinition) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *SLSAProvenanceV1PredicateBuildDefinition) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *SLSAProvenanceV1PredicateBuildDefinition) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.


### GetExternalParameters

`func (o *SLSAProvenanceV1PredicateBuildDefinition) GetExternalParameters() map[string]interface{}`

GetExternalParameters returns the ExternalParameters field if non-nil, zero value otherwise.

### GetExternalParametersOk

`func (o *SLSAProvenanceV1PredicateBuildDefinition) GetExternalParametersOk() (*map[string]interface{}, bool)`

GetExternalParametersOk returns a tuple with the ExternalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParameters

`func (o *SLSAProvenanceV1PredicateBuildDefinition) SetExternalParameters(v map[string]interface{})`

SetExternalParameters sets ExternalParameters field to given value.


### GetInternalParameters

`func (o *SLSAProvenanceV1PredicateBuildDefinition) GetInternalParameters() map[string]interface{}`

GetInternalParameters returns the InternalParameters field if non-nil, zero value otherwise.

### GetInternalParametersOk

`func (o *SLSAProvenanceV1PredicateBuildDefinition) GetInternalParametersOk() (*map[string]interface{}, bool)`

GetInternalParametersOk returns a tuple with the InternalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalParameters

`func (o *SLSAProvenanceV1PredicateBuildDefinition) SetInternalParameters(v map[string]interface{})`

SetInternalParameters sets InternalParameters field to given value.

### HasInternalParameters

`func (o *SLSAProvenanceV1PredicateBuildDefinition) HasInternalParameters() bool`

HasInternalParameters returns a boolean if a field has been set.

### SetInternalParametersNil

`func (o *SLSAProvenanceV1PredicateBuildDefinition) SetInternalParametersNil(b bool)`

 SetInternalParametersNil sets the value for InternalParameters to be an explicit nil

### UnsetInternalParameters
`func (o *SLSAProvenanceV1PredicateBuildDefinition) UnsetInternalParameters()`

UnsetInternalParameters ensures that no value is present for InternalParameters, not even an explicit nil
### GetResolvedDependencies

`func (o *SLSAProvenanceV1PredicateBuildDefinition) GetResolvedDependencies() []SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner`

GetResolvedDependencies returns the ResolvedDependencies field if non-nil, zero value otherwise.

### GetResolvedDependenciesOk

`func (o *SLSAProvenanceV1PredicateBuildDefinition) GetResolvedDependenciesOk() (*[]SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner, bool)`

GetResolvedDependenciesOk returns a tuple with the ResolvedDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDependencies

`func (o *SLSAProvenanceV1PredicateBuildDefinition) SetResolvedDependencies(v []SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner)`

SetResolvedDependencies sets ResolvedDependencies field to given value.

### HasResolvedDependencies

`func (o *SLSAProvenanceV1PredicateBuildDefinition) HasResolvedDependencies() bool`

HasResolvedDependencies returns a boolean if a field has been set.

### SetResolvedDependenciesNil

`func (o *SLSAProvenanceV1PredicateBuildDefinition) SetResolvedDependenciesNil(b bool)`

 SetResolvedDependenciesNil sets the value for ResolvedDependencies to be an explicit nil

### UnsetResolvedDependencies
`func (o *SLSAProvenanceV1PredicateBuildDefinition) UnsetResolvedDependencies()`

UnsetResolvedDependencies ensures that no value is present for ResolvedDependencies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


