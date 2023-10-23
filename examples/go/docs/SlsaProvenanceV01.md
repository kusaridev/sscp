# SlsaProvenanceV01

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Predicate** | [**SLSAProvenanceV1Predicate**](SLSAProvenanceV1Predicate.md) |  | 
**PredicateType** | **string** |  | 
**Subject** | [**[]Subject**](Subject.md) |  | 

## Methods

### NewSlsaProvenanceV01

`func NewSlsaProvenanceV01(type_ string, predicate SLSAProvenanceV1Predicate, predicateType string, subject []Subject, ) *SlsaProvenanceV01`

NewSlsaProvenanceV01 instantiates a new SlsaProvenanceV01 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlsaProvenanceV01WithDefaults

`func NewSlsaProvenanceV01WithDefaults() *SlsaProvenanceV01`

NewSlsaProvenanceV01WithDefaults instantiates a new SlsaProvenanceV01 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SlsaProvenanceV01) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlsaProvenanceV01) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlsaProvenanceV01) SetType(v string)`

SetType sets Type field to given value.


### GetPredicate

`func (o *SlsaProvenanceV01) GetPredicate() SLSAProvenanceV1Predicate`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *SlsaProvenanceV01) GetPredicateOk() (*SLSAProvenanceV1Predicate, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *SlsaProvenanceV01) SetPredicate(v SLSAProvenanceV1Predicate)`

SetPredicate sets Predicate field to given value.


### GetPredicateType

`func (o *SlsaProvenanceV01) GetPredicateType() string`

GetPredicateType returns the PredicateType field if non-nil, zero value otherwise.

### GetPredicateTypeOk

`func (o *SlsaProvenanceV01) GetPredicateTypeOk() (*string, bool)`

GetPredicateTypeOk returns a tuple with the PredicateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateType

`func (o *SlsaProvenanceV01) SetPredicateType(v string)`

SetPredicateType sets PredicateType field to given value.


### GetSubject

`func (o *SlsaProvenanceV01) GetSubject() []Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SlsaProvenanceV01) GetSubjectOk() (*[]Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SlsaProvenanceV01) SetSubject(v []Subject)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


