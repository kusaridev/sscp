# CdBuildstarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**CdBuildstartedContext**](CdBuildstartedContext.md) |  | 
**Subject** | [**CdBuildstartedSubject**](CdBuildstartedSubject.md) |  | 
**CustomData** | Pointer to [**CdBuildstartedCustomData**](CdBuildstartedCustomData.md) |  | [optional] 
**CustomDataContentType** | Pointer to **string** |  | [optional] 

## Methods

### NewCdBuildstarted

`func NewCdBuildstarted(context CdBuildstartedContext, subject CdBuildstartedSubject, ) *CdBuildstarted`

NewCdBuildstarted instantiates a new CdBuildstarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdBuildstartedWithDefaults

`func NewCdBuildstartedWithDefaults() *CdBuildstarted`

NewCdBuildstartedWithDefaults instantiates a new CdBuildstarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *CdBuildstarted) GetContext() CdBuildstartedContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CdBuildstarted) GetContextOk() (*CdBuildstartedContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CdBuildstarted) SetContext(v CdBuildstartedContext)`

SetContext sets Context field to given value.


### GetSubject

`func (o *CdBuildstarted) GetSubject() CdBuildstartedSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CdBuildstarted) GetSubjectOk() (*CdBuildstartedSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CdBuildstarted) SetSubject(v CdBuildstartedSubject)`

SetSubject sets Subject field to given value.


### GetCustomData

`func (o *CdBuildstarted) GetCustomData() CdBuildstartedCustomData`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *CdBuildstarted) GetCustomDataOk() (*CdBuildstartedCustomData, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *CdBuildstarted) SetCustomData(v CdBuildstartedCustomData)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *CdBuildstarted) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetCustomDataContentType

`func (o *CdBuildstarted) GetCustomDataContentType() string`

GetCustomDataContentType returns the CustomDataContentType field if non-nil, zero value otherwise.

### GetCustomDataContentTypeOk

`func (o *CdBuildstarted) GetCustomDataContentTypeOk() (*string, bool)`

GetCustomDataContentTypeOk returns a tuple with the CustomDataContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDataContentType

`func (o *CdBuildstarted) SetCustomDataContentType(v string)`

SetCustomDataContentType sets CustomDataContentType field to given value.

### HasCustomDataContentType

`func (o *CdBuildstarted) HasCustomDataContentType() bool`

HasCustomDataContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


