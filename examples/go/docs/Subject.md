# Subject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **map[string]string** | Represents a set of digests, mapping algorithms to their respective digest strings. | 
**Name** | **string** |  | 

## Methods

### NewSubject

`func NewSubject(digest map[string]string, name string, ) *Subject`

NewSubject instantiates a new Subject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectWithDefaults

`func NewSubjectWithDefaults() *Subject`

NewSubjectWithDefaults instantiates a new Subject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *Subject) GetDigest() map[string]string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *Subject) GetDigestOk() (*map[string]string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *Subject) SetDigest(v map[string]string)`

SetDigest sets Digest field to given value.


### GetName

`func (o *Subject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subject) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


