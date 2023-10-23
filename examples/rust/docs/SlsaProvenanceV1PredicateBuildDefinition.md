# SlsaProvenanceV1PredicateBuildDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**build_type** | **String** |  | 
**external_parameters** | [**::std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md) | The parameters that are under external control, such as those set by a user or tenant of the build platform. They MUST be complete at SLSA Build L3, meaning that there is no additional mechanism for an external party to influence the build. (At lower SLSA Build levels, the completeness MAY be best effort.)\\nThe build platform SHOULD be designed to minimize the size and complexity of externalParameters, in order to reduce fragility and ease verification. Consumers SHOULD have an expectation of what “good” looks like; the more information that they need to check, the harder that task becomes.\\nVerifiers SHOULD reject unrecognized or unexpected fields within externalParameters. | 
**internal_parameters** | Option<[**::std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> | Unordered collection of artifacts needed at build time. Completeness is best effort, at least through SLSA Build L3. For example, if the build script fetches and executes “example.com/foo.sh”, which in turn fetches “example.com/bar.tar.gz”, then both “foo.sh” and “bar.tar.gz” SHOULD be listed here. | [optional]
**resolved_dependencies** | Option<[**Vec<crate::models::SlsaProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner>**](SLSAProvenanceV1Predicate_buildDefinition_resolvedDependencies_inner.md)> | Unordered collection of artifacts needed at build time. Completeness is best effort, at least through SLSA Build L3. For example, if the build script fetches and executes “example.com/foo.sh”, which in turn fetches “example.com/bar.tar.gz”, then both “foo.sh” and “bar.tar.gz” SHOULD be listed here. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


