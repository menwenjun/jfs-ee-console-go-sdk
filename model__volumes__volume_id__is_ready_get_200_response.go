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

// checks if the VolumesVolumeIDIsReadyGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumesVolumeIDIsReadyGet200Response{}

// VolumesVolumeIDIsReadyGet200Response struct for VolumesVolumeIDIsReadyGet200Response
type VolumesVolumeIDIsReadyGet200Response struct {
	IsReady *bool `json:"is_ready,omitempty"`
}

// NewVolumesVolumeIDIsReadyGet200Response instantiates a new VolumesVolumeIDIsReadyGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumesVolumeIDIsReadyGet200Response() *VolumesVolumeIDIsReadyGet200Response {
	this := VolumesVolumeIDIsReadyGet200Response{}
	return &this
}

// NewVolumesVolumeIDIsReadyGet200ResponseWithDefaults instantiates a new VolumesVolumeIDIsReadyGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumesVolumeIDIsReadyGet200ResponseWithDefaults() *VolumesVolumeIDIsReadyGet200Response {
	this := VolumesVolumeIDIsReadyGet200Response{}
	return &this
}

// GetIsReady returns the IsReady field value if set, zero value otherwise.
func (o *VolumesVolumeIDIsReadyGet200Response) GetIsReady() bool {
	if o == nil || IsNil(o.IsReady) {
		var ret bool
		return ret
	}
	return *o.IsReady
}

// GetIsReadyOk returns a tuple with the IsReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDIsReadyGet200Response) GetIsReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReady) {
		return nil, false
	}
	return o.IsReady, true
}

// HasIsReady returns a boolean if a field has been set.
func (o *VolumesVolumeIDIsReadyGet200Response) HasIsReady() bool {
	if o != nil && !IsNil(o.IsReady) {
		return true
	}

	return false
}

// SetIsReady gets a reference to the given bool and assigns it to the IsReady field.
func (o *VolumesVolumeIDIsReadyGet200Response) SetIsReady(v bool) {
	o.IsReady = &v
}

func (o VolumesVolumeIDIsReadyGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumesVolumeIDIsReadyGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsReady) {
		toSerialize["is_ready"] = o.IsReady
	}
	return toSerialize, nil
}

type NullableVolumesVolumeIDIsReadyGet200Response struct {
	value *VolumesVolumeIDIsReadyGet200Response
	isSet bool
}

func (v NullableVolumesVolumeIDIsReadyGet200Response) Get() *VolumesVolumeIDIsReadyGet200Response {
	return v.value
}

func (v *NullableVolumesVolumeIDIsReadyGet200Response) Set(val *VolumesVolumeIDIsReadyGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumesVolumeIDIsReadyGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumesVolumeIDIsReadyGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumesVolumeIDIsReadyGet200Response(val *VolumesVolumeIDIsReadyGet200Response) *NullableVolumesVolumeIDIsReadyGet200Response {
	return &NullableVolumesVolumeIDIsReadyGet200Response{value: val, isSet: true}
}

func (v NullableVolumesVolumeIDIsReadyGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumesVolumeIDIsReadyGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

