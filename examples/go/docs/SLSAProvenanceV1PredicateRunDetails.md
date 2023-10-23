# SLSAProvenanceV1PredicateRunDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Builder** | [**SLSAProvenanceV1PredicateRunDetailsBuilder**](SLSAProvenanceV1PredicateRunDetailsBuilder.md) |  | 
**Byproducts** | Pointer to [**[]SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner**](SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner.md) | Additional artifacts generated during the build that are not considered the “output” of the build but that might be needed during debugging or incident response. For example, this might reference logs generated during the build and/or a digest of the fully evaluated build configuration.\\nIn most cases, this SHOULD NOT contain all intermediate files generated during the build. Instead, this SHOULD only contain files that are likely to be useful later and that cannot be easily reproduced. | [optional] 
**Metadata** | Pointer to [**NullableSLSAProvenanceV1PredicateRunDetailsMetadata**](SLSAProvenanceV1PredicateRunDetailsMetadata.md) |  | [optional] 

## Methods

### NewSLSAProvenanceV1PredicateRunDetails

`func NewSLSAProvenanceV1PredicateRunDetails(builder SLSAProvenanceV1PredicateRunDetailsBuilder, ) *SLSAProvenanceV1PredicateRunDetails`

NewSLSAProvenanceV1PredicateRunDetails instantiates a new SLSAProvenanceV1PredicateRunDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLSAProvenanceV1PredicateRunDetailsWithDefaults

`func NewSLSAProvenanceV1PredicateRunDetailsWithDefaults() *SLSAProvenanceV1PredicateRunDetails`

NewSLSAProvenanceV1PredicateRunDetailsWithDefaults instantiates a new SLSAProvenanceV1PredicateRunDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilder

`func (o *SLSAProvenanceV1PredicateRunDetails) GetBuilder() SLSAProvenanceV1PredicateRunDetailsBuilder`

GetBuilder returns the Builder field if non-nil, zero value otherwise.

### GetBuilderOk

`func (o *SLSAProvenanceV1PredicateRunDetails) GetBuilderOk() (*SLSAProvenanceV1PredicateRunDetailsBuilder, bool)`

GetBuilderOk returns a tuple with the Builder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilder

`func (o *SLSAProvenanceV1PredicateRunDetails) SetBuilder(v SLSAProvenanceV1PredicateRunDetailsBuilder)`

SetBuilder sets Builder field to given value.


### GetByproducts

`func (o *SLSAProvenanceV1PredicateRunDetails) GetByproducts() []SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner`

GetByproducts returns the Byproducts field if non-nil, zero value otherwise.

### GetByproductsOk

`func (o *SLSAProvenanceV1PredicateRunDetails) GetByproductsOk() (*[]SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner, bool)`

GetByproductsOk returns a tuple with the Byproducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByproducts

`func (o *SLSAProvenanceV1PredicateRunDetails) SetByproducts(v []SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner)`

SetByproducts sets Byproducts field to given value.

### HasByproducts

`func (o *SLSAProvenanceV1PredicateRunDetails) HasByproducts() bool`

HasByproducts returns a boolean if a field has been set.

### SetByproductsNil

`func (o *SLSAProvenanceV1PredicateRunDetails) SetByproductsNil(b bool)`

 SetByproductsNil sets the value for Byproducts to be an explicit nil

### UnsetByproducts
`func (o *SLSAProvenanceV1PredicateRunDetails) UnsetByproducts()`

UnsetByproducts ensures that no value is present for Byproducts, not even an explicit nil
### GetMetadata

`func (o *SLSAProvenanceV1PredicateRunDetails) GetMetadata() SLSAProvenanceV1PredicateRunDetailsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SLSAProvenanceV1PredicateRunDetails) GetMetadataOk() (*SLSAProvenanceV1PredicateRunDetailsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SLSAProvenanceV1PredicateRunDetails) SetMetadata(v SLSAProvenanceV1PredicateRunDetailsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SLSAProvenanceV1PredicateRunDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SLSAProvenanceV1PredicateRunDetails) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SLSAProvenanceV1PredicateRunDetails) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


