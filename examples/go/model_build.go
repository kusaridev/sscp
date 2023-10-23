/*
Secure Build API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Build type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Build{}

// Build struct for Build
type Build struct {
	Id *int32 `json:"id,omitempty"`
	GitRef *GitRef `json:"gitRef,omitempty"`
	BuildCommand []string `json:"buildCommand,omitempty"`
	BuildType *string `json:"buildType,omitempty"`
}

// NewBuild instantiates a new Build object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuild() *Build {
	this := Build{}
	return &this
}

// NewBuildWithDefaults instantiates a new Build object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildWithDefaults() *Build {
	this := Build{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Build) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Build) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Build) SetId(v int32) {
	o.Id = &v
}

// GetGitRef returns the GitRef field value if set, zero value otherwise.
func (o *Build) GetGitRef() GitRef {
	if o == nil || IsNil(o.GitRef) {
		var ret GitRef
		return ret
	}
	return *o.GitRef
}

// GetGitRefOk returns a tuple with the GitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetGitRefOk() (*GitRef, bool) {
	if o == nil || IsNil(o.GitRef) {
		return nil, false
	}
	return o.GitRef, true
}

// HasGitRef returns a boolean if a field has been set.
func (o *Build) HasGitRef() bool {
	if o != nil && !IsNil(o.GitRef) {
		return true
	}

	return false
}

// SetGitRef gets a reference to the given GitRef and assigns it to the GitRef field.
func (o *Build) SetGitRef(v GitRef) {
	o.GitRef = &v
}

// GetBuildCommand returns the BuildCommand field value if set, zero value otherwise.
func (o *Build) GetBuildCommand() []string {
	if o == nil || IsNil(o.BuildCommand) {
		var ret []string
		return ret
	}
	return o.BuildCommand
}

// GetBuildCommandOk returns a tuple with the BuildCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetBuildCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.BuildCommand) {
		return nil, false
	}
	return o.BuildCommand, true
}

// HasBuildCommand returns a boolean if a field has been set.
func (o *Build) HasBuildCommand() bool {
	if o != nil && !IsNil(o.BuildCommand) {
		return true
	}

	return false
}

// SetBuildCommand gets a reference to the given []string and assigns it to the BuildCommand field.
func (o *Build) SetBuildCommand(v []string) {
	o.BuildCommand = v
}

// GetBuildType returns the BuildType field value if set, zero value otherwise.
func (o *Build) GetBuildType() string {
	if o == nil || IsNil(o.BuildType) {
		var ret string
		return ret
	}
	return *o.BuildType
}

// GetBuildTypeOk returns a tuple with the BuildType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetBuildTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BuildType) {
		return nil, false
	}
	return o.BuildType, true
}

// HasBuildType returns a boolean if a field has been set.
func (o *Build) HasBuildType() bool {
	if o != nil && !IsNil(o.BuildType) {
		return true
	}

	return false
}

// SetBuildType gets a reference to the given string and assigns it to the BuildType field.
func (o *Build) SetBuildType(v string) {
	o.BuildType = &v
}

func (o Build) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Build) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GitRef) {
		toSerialize["gitRef"] = o.GitRef
	}
	if !IsNil(o.BuildCommand) {
		toSerialize["buildCommand"] = o.BuildCommand
	}
	if !IsNil(o.BuildType) {
		toSerialize["buildType"] = o.BuildType
	}
	return toSerialize, nil
}

type NullableBuild struct {
	value *Build
	isSet bool
}

func (v NullableBuild) Get() *Build {
	return v.value
}

func (v *NullableBuild) Set(val *Build) {
	v.value = val
	v.isSet = true
}

func (v NullableBuild) IsSet() bool {
	return v.isSet
}

func (v *NullableBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuild(val *Build) *NullableBuild {
	return &NullableBuild{value: val, isSet: true}
}

func (v NullableBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

