# CdBuildstartedSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "build"]
**Content** | **map[string]interface{}** |  | 

## Methods

### NewCdBuildstartedSubject

`func NewCdBuildstartedSubject(id string, type_ string, content map[string]interface{}, ) *CdBuildstartedSubject`

NewCdBuildstartedSubject instantiates a new CdBuildstartedSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdBuildstartedSubjectWithDefaults

`func NewCdBuildstartedSubjectWithDefaults() *CdBuildstartedSubject`

NewCdBuildstartedSubjectWithDefaults instantiates a new CdBuildstartedSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdBuildstartedSubject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdBuildstartedSubject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdBuildstartedSubject) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *CdBuildstartedSubject) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CdBuildstartedSubject) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CdBuildstartedSubject) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CdBuildstartedSubject) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *CdBuildstartedSubject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CdBuildstartedSubject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CdBuildstartedSubject) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *CdBuildstartedSubject) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CdBuildstartedSubject) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CdBuildstartedSubject) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


