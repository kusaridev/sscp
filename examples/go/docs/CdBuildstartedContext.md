# CdBuildstartedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Id** | **string** |  | 
**Source** | **string** |  | 
**Type** | **string** |  | [default to "dev.cdevents.build.started.0.1.1"]
**Timestamp** | **time.Time** |  | 

## Methods

### NewCdBuildstartedContext

`func NewCdBuildstartedContext(version string, id string, source string, type_ string, timestamp time.Time, ) *CdBuildstartedContext`

NewCdBuildstartedContext instantiates a new CdBuildstartedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdBuildstartedContextWithDefaults

`func NewCdBuildstartedContextWithDefaults() *CdBuildstartedContext`

NewCdBuildstartedContextWithDefaults instantiates a new CdBuildstartedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CdBuildstartedContext) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CdBuildstartedContext) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CdBuildstartedContext) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetId

`func (o *CdBuildstartedContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdBuildstartedContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdBuildstartedContext) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *CdBuildstartedContext) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CdBuildstartedContext) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CdBuildstartedContext) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *CdBuildstartedContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CdBuildstartedContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CdBuildstartedContext) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *CdBuildstartedContext) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CdBuildstartedContext) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CdBuildstartedContext) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


