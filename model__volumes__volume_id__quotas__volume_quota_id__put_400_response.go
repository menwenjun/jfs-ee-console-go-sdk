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

// checks if the VolumesVolumeIDQuotasVolumeQuotaIDPut400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumesVolumeIDQuotasVolumeQuotaIDPut400Response{}

// VolumesVolumeIDQuotasVolumeQuotaIDPut400Response struct for VolumesVolumeIDQuotasVolumeQuotaIDPut400Response
type VolumesVolumeIDQuotasVolumeQuotaIDPut400Response struct {
	NonFieldErrors []string `json:"non_field_errors,omitempty"`
}

// NewVolumesVolumeIDQuotasVolumeQuotaIDPut400Response instantiates a new VolumesVolumeIDQuotasVolumeQuotaIDPut400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumesVolumeIDQuotasVolumeQuotaIDPut400Response() *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response {
	this := VolumesVolumeIDQuotasVolumeQuotaIDPut400Response{}
	return &this
}

// NewVolumesVolumeIDQuotasVolumeQuotaIDPut400ResponseWithDefaults instantiates a new VolumesVolumeIDQuotasVolumeQuotaIDPut400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumesVolumeIDQuotasVolumeQuotaIDPut400ResponseWithDefaults() *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response {
	this := VolumesVolumeIDQuotasVolumeQuotaIDPut400Response{}
	return &this
}

// GetNonFieldErrors returns the NonFieldErrors field value if set, zero value otherwise.
func (o *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response) GetNonFieldErrors() []string {
	if o == nil || IsNil(o.NonFieldErrors) {
		var ret []string
		return ret
	}
	return o.NonFieldErrors
}

// GetNonFieldErrorsOk returns a tuple with the NonFieldErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response) GetNonFieldErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.NonFieldErrors) {
		return nil, false
	}
	return o.NonFieldErrors, true
}

// HasNonFieldErrors returns a boolean if a field has been set.
func (o *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response) HasNonFieldErrors() bool {
	if o != nil && !IsNil(o.NonFieldErrors) {
		return true
	}

	return false
}

// SetNonFieldErrors gets a reference to the given []string and assigns it to the NonFieldErrors field.
func (o *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response) SetNonFieldErrors(v []string) {
	o.NonFieldErrors = v
}

func (o VolumesVolumeIDQuotasVolumeQuotaIDPut400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumesVolumeIDQuotasVolumeQuotaIDPut400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NonFieldErrors) {
		toSerialize["non_field_errors"] = o.NonFieldErrors
	}
	return toSerialize, nil
}

type NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response struct {
	value *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response
	isSet bool
}

func (v NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response) Get() *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response {
	return v.value
}

func (v *NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response) Set(val *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response(val *VolumesVolumeIDQuotasVolumeQuotaIDPut400Response) *NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response {
	return &NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response{value: val, isSet: true}
}

func (v NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumesVolumeIDQuotasVolumeQuotaIDPut400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

