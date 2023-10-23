# SlsaProvenanceV1PredicateRunDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**builder** | [**crate::models::SlsaProvenanceV1PredicateRunDetailsBuilder**](SLSAProvenanceV1Predicate_runDetails_builder.md) |  | 
**byproducts** | Option<[**Vec<crate::models::SlsaProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner>**](SLSAProvenanceV1Predicate_buildDefinition_resolvedDependencies_inner.md)> | Additional artifacts generated during the build that are not considered the “output” of the build but that might be needed during debugging or incident response. For example, this might reference logs generated during the build and/or a digest of the fully evaluated build configuration.\\nIn most cases, this SHOULD NOT contain all intermediate files generated during the build. Instead, this SHOULD only contain files that are likely to be useful later and that cannot be easily reproduced. | [optional]
**metadata** | Option<[**crate::models::SlsaProvenanceV1PredicateRunDetailsMetadata**](SLSAProvenanceV1Predicate_runDetails_metadata.md)> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


