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

// checks if the Inodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Inodes{}

// Inodes struct for Inodes
type Inodes struct {
	Inodes []int32 `json:"inodes,omitempty"`
}

// NewInodes instantiates a new Inodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInodes() *Inodes {
	this := Inodes{}
	return &this
}

// NewInodesWithDefaults instantiates a new Inodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInodesWithDefaults() *Inodes {
	this := Inodes{}
	return &this
}

// GetInodes returns the Inodes field value if set, zero value otherwise.
func (o *Inodes) GetInodes() []int32 {
	if o == nil || IsNil(o.Inodes) {
		var ret []int32
		return ret
	}
	return o.Inodes
}

// GetInodesOk returns a tuple with the Inodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inodes) GetInodesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Inodes) {
		return nil, false
	}
	return o.Inodes, true
}

// HasInodes returns a boolean if a field has been set.
func (o *Inodes) HasInodes() bool {
	if o != nil && !IsNil(o.Inodes) {
		return true
	}

	return false
}

// SetInodes gets a reference to the given []int32 and assigns it to the Inodes field.
func (o *Inodes) SetInodes(v []int32) {
	o.Inodes = v
}

func (o Inodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Inodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inodes) {
		toSerialize["inodes"] = o.Inodes
	}
	return toSerialize, nil
}

type NullableInodes struct {
	value *Inodes
	isSet bool
}

func (v NullableInodes) Get() *Inodes {
	return v.value
}

func (v *NullableInodes) Set(val *Inodes) {
	v.value = val
	v.isSet = true
}

func (v NullableInodes) IsSet() bool {
	return v.isSet
}

func (v *NullableInodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInodes(val *Inodes) *NullableInodes {
	return &NullableInodes{value: val, isSet: true}
}

func (v NullableInodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


