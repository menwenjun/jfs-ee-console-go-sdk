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

// checks if the VolumesVolumeIDGet404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumesVolumeIDGet404Response{}

// VolumesVolumeIDGet404Response struct for VolumesVolumeIDGet404Response
type VolumesVolumeIDGet404Response struct {
	Detail *string `json:"detail,omitempty"`
}

// NewVolumesVolumeIDGet404Response instantiates a new VolumesVolumeIDGet404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumesVolumeIDGet404Response() *VolumesVolumeIDGet404Response {
	this := VolumesVolumeIDGet404Response{}
	return &this
}

// NewVolumesVolumeIDGet404ResponseWithDefaults instantiates a new VolumesVolumeIDGet404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumesVolumeIDGet404ResponseWithDefaults() *VolumesVolumeIDGet404Response {
	this := VolumesVolumeIDGet404Response{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *VolumesVolumeIDGet404Response) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDGet404Response) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *VolumesVolumeIDGet404Response) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *VolumesVolumeIDGet404Response) SetDetail(v string) {
	o.Detail = &v
}

func (o VolumesVolumeIDGet404Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumesVolumeIDGet404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableVolumesVolumeIDGet404Response struct {
	value *VolumesVolumeIDGet404Response
	isSet bool
}

func (v NullableVolumesVolumeIDGet404Response) Get() *VolumesVolumeIDGet404Response {
	return v.value
}

func (v *NullableVolumesVolumeIDGet404Response) Set(val *VolumesVolumeIDGet404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumesVolumeIDGet404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumesVolumeIDGet404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumesVolumeIDGet404Response(val *VolumesVolumeIDGet404Response) *NullableVolumesVolumeIDGet404Response {
	return &NullableVolumesVolumeIDGet404Response{value: val, isSet: true}
}

func (v NullableVolumesVolumeIDGet404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumesVolumeIDGet404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

