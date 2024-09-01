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

// checks if the UpdateVolumeQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVolumeQuota{}

// UpdateVolumeQuota struct for UpdateVolumeQuota
type UpdateVolumeQuota struct {
	// The directory to be set quota on, it must be equal to the original path if it's provided
	Path *string `json:"path,omitempty"`
	// Total inodes, `0` means no limit
	Inodes *int32 `json:"inodes,omitempty"`
	// Total size in bytes, `0` means no limit
	Size *int32 `json:"size,omitempty"`
}

// NewUpdateVolumeQuota instantiates a new UpdateVolumeQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVolumeQuota() *UpdateVolumeQuota {
	this := UpdateVolumeQuota{}
	return &this
}

// NewUpdateVolumeQuotaWithDefaults instantiates a new UpdateVolumeQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVolumeQuotaWithDefaults() *UpdateVolumeQuota {
	this := UpdateVolumeQuota{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *UpdateVolumeQuota) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeQuota) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UpdateVolumeQuota) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *UpdateVolumeQuota) SetPath(v string) {
	o.Path = &v
}

// GetInodes returns the Inodes field value if set, zero value otherwise.
func (o *UpdateVolumeQuota) GetInodes() int32 {
	if o == nil || IsNil(o.Inodes) {
		var ret int32
		return ret
	}
	return *o.Inodes
}

// GetInodesOk returns a tuple with the Inodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeQuota) GetInodesOk() (*int32, bool) {
	if o == nil || IsNil(o.Inodes) {
		return nil, false
	}
	return o.Inodes, true
}

// HasInodes returns a boolean if a field has been set.
func (o *UpdateVolumeQuota) HasInodes() bool {
	if o != nil && !IsNil(o.Inodes) {
		return true
	}

	return false
}

// SetInodes gets a reference to the given int32 and assigns it to the Inodes field.
func (o *UpdateVolumeQuota) SetInodes(v int32) {
	o.Inodes = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UpdateVolumeQuota) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeQuota) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UpdateVolumeQuota) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *UpdateVolumeQuota) SetSize(v int32) {
	o.Size = &v
}

func (o UpdateVolumeQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVolumeQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Inodes) {
		toSerialize["inodes"] = o.Inodes
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableUpdateVolumeQuota struct {
	value *UpdateVolumeQuota
	isSet bool
}

func (v NullableUpdateVolumeQuota) Get() *UpdateVolumeQuota {
	return v.value
}

func (v *NullableUpdateVolumeQuota) Set(val *UpdateVolumeQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVolumeQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVolumeQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVolumeQuota(val *UpdateVolumeQuota) *NullableUpdateVolumeQuota {
	return &NullableUpdateVolumeQuota{value: val, isSet: true}
}

func (v NullableUpdateVolumeQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVolumeQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


