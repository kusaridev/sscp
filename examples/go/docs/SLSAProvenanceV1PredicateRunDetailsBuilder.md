# SLSAProvenanceV1PredicateRunDetailsBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuilderDependencies** | Pointer to [**[]SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner**](SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner.md) |  | [optional] 
**Id** | **string** |  | 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSLSAProvenanceV1PredicateRunDetailsBuilder

`func NewSLSAProvenanceV1PredicateRunDetailsBuilder(id string, ) *SLSAProvenanceV1PredicateRunDetailsBuilder`

NewSLSAProvenanceV1PredicateRunDetailsBuilder instantiates a new SLSAProvenanceV1PredicateRunDetailsBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLSAProvenanceV1PredicateRunDetailsBuilderWithDefaults

`func NewSLSAProvenanceV1PredicateRunDetailsBuilderWithDefaults() *SLSAProvenanceV1PredicateRunDetailsBuilder`

NewSLSAProvenanceV1PredicateRunDetailsBuilderWithDefaults instantiates a new SLSAProvenanceV1PredicateRunDetailsBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilderDependencies

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) GetBuilderDependencies() []SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner`

GetBuilderDependencies returns the BuilderDependencies field if non-nil, zero value otherwise.

### GetBuilderDependenciesOk

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) GetBuilderDependenciesOk() (*[]SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner, bool)`

GetBuilderDependenciesOk returns a tuple with the BuilderDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderDependencies

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) SetBuilderDependencies(v []SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner)`

SetBuilderDependencies sets BuilderDependencies field to given value.

### HasBuilderDependencies

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) HasBuilderDependencies() bool`

HasBuilderDependencies returns a boolean if a field has been set.

### SetBuilderDependenciesNil

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) SetBuilderDependenciesNil(b bool)`

 SetBuilderDependenciesNil sets the value for BuilderDependencies to be an explicit nil

### UnsetBuilderDependencies
`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) UnsetBuilderDependencies()`

UnsetBuilderDependencies ensures that no value is present for BuilderDependencies, not even an explicit nil
### GetId

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SLSAProvenanceV1PredicateRunDetailsBuilder) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


