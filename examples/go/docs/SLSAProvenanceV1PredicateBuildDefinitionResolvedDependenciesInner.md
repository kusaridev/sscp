# SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]interface{}** | This field MAY be used to provide additional information or metadata about the resource or artifact that may be useful to the consumer when evaluating the attestation against a policy. | [optional] 
**Content** | Pointer to **string** | The contents of the resource or artifact. This field is REQUIRED unless either uri or digest is set. | [optional] 
**Digest** | Pointer to **map[string]string** | A set of cryptographic digests of the contents of the resource or artifact. This field is REQUIRED unless either uri or content is set. | [optional] 
**DownloadLocation** | Pointer to **string** | The location of the described resource or artifact, if different from the uri. | [optional] 
**MediaType** | Pointer to **NullableString** | The MIME Type (i.e., media type) of the described resource or artifact. | [optional] 
**Name** | Pointer to **NullableString** | Machine-readable identifier for distinguishing between descriptors. | [optional] 
**Uri** | **string** | A URI used to identify the resource or artifact globally. This field is REQUIRED unless either digest or content is set. | 

## Methods

### NewSLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner

`func NewSLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner(uri string, ) *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner`

NewSLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner instantiates a new SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInnerWithDefaults

`func NewSLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInnerWithDefaults() *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner`

NewSLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInnerWithDefaults instantiates a new SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetContent

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDigest

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetDigest() map[string]string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetDigestOk() (*map[string]string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetDigest(v map[string]string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### SetDigestNil

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetDigestNil(b bool)`

 SetDigestNil sets the value for Digest to be an explicit nil

### UnsetDigest
`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) UnsetDigest()`

UnsetDigest ensures that no value is present for Digest, not even an explicit nil
### GetDownloadLocation

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetDownloadLocation() string`

GetDownloadLocation returns the DownloadLocation field if non-nil, zero value otherwise.

### GetDownloadLocationOk

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetDownloadLocationOk() (*string, bool)`

GetDownloadLocationOk returns a tuple with the DownloadLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLocation

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetDownloadLocation(v string)`

SetDownloadLocation sets DownloadLocation field to given value.

### HasDownloadLocation

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) HasDownloadLocation() bool`

HasDownloadLocation returns a boolean if a field has been set.

### GetMediaType

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetName

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUri

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SLSAProvenanceV1PredicateBuildDefinitionResolvedDependenciesInner) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


