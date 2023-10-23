# SLSAProvenanceV1PredicateRunDetailsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishedOn** | Pointer to **NullableTime** | The timestamp of when the build completed. | [optional] 
**InvocationId** | Pointer to **NullableString** | Identifies this particular build invocation, which can be useful for finding associated logs or other ad-hoc analysis. The exact meaning and format is defined by builder.id; by default it is treated as opaque and case-sensitive. The value SHOULD be globally unique. | [optional] 
**StartedOn** | Pointer to **NullableTime** | The timestamp of when the build started. | [optional] 

## Methods

### NewSLSAProvenanceV1PredicateRunDetailsMetadata

`func NewSLSAProvenanceV1PredicateRunDetailsMetadata() *SLSAProvenanceV1PredicateRunDetailsMetadata`

NewSLSAProvenanceV1PredicateRunDetailsMetadata instantiates a new SLSAProvenanceV1PredicateRunDetailsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLSAProvenanceV1PredicateRunDetailsMetadataWithDefaults

`func NewSLSAProvenanceV1PredicateRunDetailsMetadataWithDefaults() *SLSAProvenanceV1PredicateRunDetailsMetadata`

NewSLSAProvenanceV1PredicateRunDetailsMetadataWithDefaults instantiates a new SLSAProvenanceV1PredicateRunDetailsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishedOn

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) GetFinishedOn() time.Time`

GetFinishedOn returns the FinishedOn field if non-nil, zero value otherwise.

### GetFinishedOnOk

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) GetFinishedOnOk() (*time.Time, bool)`

GetFinishedOnOk returns a tuple with the FinishedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedOn

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) SetFinishedOn(v time.Time)`

SetFinishedOn sets FinishedOn field to given value.

### HasFinishedOn

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) HasFinishedOn() bool`

HasFinishedOn returns a boolean if a field has been set.

### SetFinishedOnNil

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) SetFinishedOnNil(b bool)`

 SetFinishedOnNil sets the value for FinishedOn to be an explicit nil

### UnsetFinishedOn
`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) UnsetFinishedOn()`

UnsetFinishedOn ensures that no value is present for FinishedOn, not even an explicit nil
### GetInvocationId

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) GetInvocationId() string`

GetInvocationId returns the InvocationId field if non-nil, zero value otherwise.

### GetInvocationIdOk

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) GetInvocationIdOk() (*string, bool)`

GetInvocationIdOk returns a tuple with the InvocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationId

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) SetInvocationId(v string)`

SetInvocationId sets InvocationId field to given value.

### HasInvocationId

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) HasInvocationId() bool`

HasInvocationId returns a boolean if a field has been set.

### SetInvocationIdNil

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) SetInvocationIdNil(b bool)`

 SetInvocationIdNil sets the value for InvocationId to be an explicit nil

### UnsetInvocationId
`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) UnsetInvocationId()`

UnsetInvocationId ensures that no value is present for InvocationId, not even an explicit nil
### GetStartedOn

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *SLSAProvenanceV1PredicateRunDetailsMetadata) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


