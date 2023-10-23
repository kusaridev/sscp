# SlsaProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**annotations** | Option<[**::std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> | This field MAY be used to provide additional information or metadata about the resource or artifact that may be useful to the consumer when evaluating the attestation against a policy. | [optional]
**content** | Option<**String**> | The contents of the resource or artifact. This field is REQUIRED unless either uri or digest is set. | [optional]
**digest** | Option<**::std::collections::HashMap<String, String>**> | A set of cryptographic digests of the contents of the resource or artifact. This field is REQUIRED unless either uri or content is set. | [optional]
**download_location** | Option<**String**> | The location of the described resource or artifact, if different from the uri. | [optional]
**media_type** | Option<**String**> | The MIME Type (i.e., media type) of the described resource or artifact. | [optional]
**name** | Option<**String**> | Machine-readable identifier for distinguishing between descriptors. | [optional]
**uri** | **String** | A URI used to identify the resource or artifact globally. This field is REQUIRED unless either digest or content is set. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


