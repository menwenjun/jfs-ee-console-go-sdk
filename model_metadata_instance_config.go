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

// checks if the MetadataInstanceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataInstanceConfig{}

// MetadataInstanceConfig struct for MetadataInstanceConfig
type MetadataInstanceConfig struct {
	ANY_CONFIG_KEY *string `json:"ANY_CONFIG_KEY,omitempty"`
}

// NewMetadataInstanceConfig instantiates a new MetadataInstanceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataInstanceConfig() *MetadataInstanceConfig {
	this := MetadataInstanceConfig{}
	return &this
}

// NewMetadataInstanceConfigWithDefaults instantiates a new MetadataInstanceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataInstanceConfigWithDefaults() *MetadataInstanceConfig {
	this := MetadataInstanceConfig{}
	return &this
}

// GetANY_CONFIG_KEY returns the ANY_CONFIG_KEY field value if set, zero value otherwise.
func (o *MetadataInstanceConfig) GetANY_CONFIG_KEY() string {
	if o == nil || IsNil(o.ANY_CONFIG_KEY) {
		var ret string
		return ret
	}
	return *o.ANY_CONFIG_KEY
}

// GetANY_CONFIG_KEYOk returns a tuple with the ANY_CONFIG_KEY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataInstanceConfig) GetANY_CONFIG_KEYOk() (*string, bool) {
	if o == nil || IsNil(o.ANY_CONFIG_KEY) {
		return nil, false
	}
	return o.ANY_CONFIG_KEY, true
}

// HasANY_CONFIG_KEY returns a boolean if a field has been set.
func (o *MetadataInstanceConfig) HasANY_CONFIG_KEY() bool {
	if o != nil && !IsNil(o.ANY_CONFIG_KEY) {
		return true
	}

	return false
}

// SetANY_CONFIG_KEY gets a reference to the given string and assigns it to the ANY_CONFIG_KEY field.
func (o *MetadataInstanceConfig) SetANY_CONFIG_KEY(v string) {
	o.ANY_CONFIG_KEY = &v
}

func (o MetadataInstanceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataInstanceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ANY_CONFIG_KEY) {
		toSerialize["ANY_CONFIG_KEY"] = o.ANY_CONFIG_KEY
	}
	return toSerialize, nil
}

type NullableMetadataInstanceConfig struct {
	value *MetadataInstanceConfig
	isSet bool
}

func (v NullableMetadataInstanceConfig) Get() *MetadataInstanceConfig {
	return v.value
}

func (v *NullableMetadataInstanceConfig) Set(val *MetadataInstanceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataInstanceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataInstanceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataInstanceConfig(val *MetadataInstanceConfig) *NullableMetadataInstanceConfig {
	return &NullableMetadataInstanceConfig{value: val, isSet: true}
}

func (v NullableMetadataInstanceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataInstanceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


