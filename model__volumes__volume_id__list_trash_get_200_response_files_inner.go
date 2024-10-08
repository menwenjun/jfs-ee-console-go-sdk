/*
JuiceFS console API

API of the JuiceFS console (https://juicefs.com/api/v1)

API version: 0.0.1
Contact: team@juicedata.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the VolumesVolumeIDListTrashGet200ResponseFilesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumesVolumeIDListTrashGet200ResponseFilesInner{}

// VolumesVolumeIDListTrashGet200ResponseFilesInner struct for VolumesVolumeIDListTrashGet200ResponseFilesInner
type VolumesVolumeIDListTrashGet200ResponseFilesInner struct {
	Inode *int32 `json:"inode,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewVolumesVolumeIDListTrashGet200ResponseFilesInner instantiates a new VolumesVolumeIDListTrashGet200ResponseFilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumesVolumeIDListTrashGet200ResponseFilesInner() *VolumesVolumeIDListTrashGet200ResponseFilesInner {
	this := VolumesVolumeIDListTrashGet200ResponseFilesInner{}
	return &this
}

// NewVolumesVolumeIDListTrashGet200ResponseFilesInnerWithDefaults instantiates a new VolumesVolumeIDListTrashGet200ResponseFilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumesVolumeIDListTrashGet200ResponseFilesInnerWithDefaults() *VolumesVolumeIDListTrashGet200ResponseFilesInner {
	this := VolumesVolumeIDListTrashGet200ResponseFilesInner{}
	return &this
}

// GetInode returns the Inode field value if set, zero value otherwise.
func (o *VolumesVolumeIDListTrashGet200ResponseFilesInner) GetInode() int32 {
	if o == nil || IsNil(o.Inode) {
		var ret int32
		return ret
	}
	return *o.Inode
}

// GetInodeOk returns a tuple with the Inode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDListTrashGet200ResponseFilesInner) GetInodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Inode) {
		return nil, false
	}
	return o.Inode, true
}

// HasInode returns a boolean if a field has been set.
func (o *VolumesVolumeIDListTrashGet200ResponseFilesInner) HasInode() bool {
	if o != nil && !IsNil(o.Inode) {
		return true
	}

	return false
}

// SetInode gets a reference to the given int32 and assigns it to the Inode field.
func (o *VolumesVolumeIDListTrashGet200ResponseFilesInner) SetInode(v int32) {
	o.Inode = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VolumesVolumeIDListTrashGet200ResponseFilesInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDListTrashGet200ResponseFilesInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VolumesVolumeIDListTrashGet200ResponseFilesInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VolumesVolumeIDListTrashGet200ResponseFilesInner) SetPath(v string) {
	o.Path = &v
}

func (o VolumesVolumeIDListTrashGet200ResponseFilesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumesVolumeIDListTrashGet200ResponseFilesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inode) {
		toSerialize["inode"] = o.Inode
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableVolumesVolumeIDListTrashGet200ResponseFilesInner struct {
	value *VolumesVolumeIDListTrashGet200ResponseFilesInner
	isSet bool
}

func (v NullableVolumesVolumeIDListTrashGet200ResponseFilesInner) Get() *VolumesVolumeIDListTrashGet200ResponseFilesInner {
	return v.value
}

func (v *NullableVolumesVolumeIDListTrashGet200ResponseFilesInner) Set(val *VolumesVolumeIDListTrashGet200ResponseFilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumesVolumeIDListTrashGet200ResponseFilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumesVolumeIDListTrashGet200ResponseFilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumesVolumeIDListTrashGet200ResponseFilesInner(val *VolumesVolumeIDListTrashGet200ResponseFilesInner) *NullableVolumesVolumeIDListTrashGet200ResponseFilesInner {
	return &NullableVolumesVolumeIDListTrashGet200ResponseFilesInner{value: val, isSet: true}
}

func (v NullableVolumesVolumeIDListTrashGet200ResponseFilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumesVolumeIDListTrashGet200ResponseFilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


