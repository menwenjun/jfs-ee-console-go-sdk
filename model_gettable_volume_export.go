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

// checks if the GettableVolumeExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GettableVolumeExport{}

// GettableVolumeExport struct for GettableVolumeExport
type GettableVolumeExport struct {
	Id *int32 `json:"id,omitempty"`
	Desc *string `json:"desc,omitempty"`
	Iprange *string `json:"iprange,omitempty"`
	Apionly *bool `json:"apionly,omitempty"`
	Readonly *bool `json:"readonly,omitempty"`
	Appendonly *bool `json:"appendonly,omitempty"`
	Internalip *bool `json:"internalip,omitempty"`
	Token *string `json:"token,omitempty"`
	Passwd *string `json:"passwd,omitempty"`
	Qos *string `json:"qos,omitempty"`
	Extend *string `json:"extend,omitempty"`
	Subdir *string `json:"subdir,omitempty"`
}

// NewGettableVolumeExport instantiates a new GettableVolumeExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGettableVolumeExport() *GettableVolumeExport {
	this := GettableVolumeExport{}
	return &this
}

// NewGettableVolumeExportWithDefaults instantiates a new GettableVolumeExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGettableVolumeExportWithDefaults() *GettableVolumeExport {
	this := GettableVolumeExport{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GettableVolumeExport) SetId(v int32) {
	o.Id = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *GettableVolumeExport) SetDesc(v string) {
	o.Desc = &v
}

// GetIprange returns the Iprange field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetIprange() string {
	if o == nil || IsNil(o.Iprange) {
		var ret string
		return ret
	}
	return *o.Iprange
}

// GetIprangeOk returns a tuple with the Iprange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetIprangeOk() (*string, bool) {
	if o == nil || IsNil(o.Iprange) {
		return nil, false
	}
	return o.Iprange, true
}

// HasIprange returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasIprange() bool {
	if o != nil && !IsNil(o.Iprange) {
		return true
	}

	return false
}

// SetIprange gets a reference to the given string and assigns it to the Iprange field.
func (o *GettableVolumeExport) SetIprange(v string) {
	o.Iprange = &v
}

// GetApionly returns the Apionly field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetApionly() bool {
	if o == nil || IsNil(o.Apionly) {
		var ret bool
		return ret
	}
	return *o.Apionly
}

// GetApionlyOk returns a tuple with the Apionly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetApionlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Apionly) {
		return nil, false
	}
	return o.Apionly, true
}

// HasApionly returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasApionly() bool {
	if o != nil && !IsNil(o.Apionly) {
		return true
	}

	return false
}

// SetApionly gets a reference to the given bool and assigns it to the Apionly field.
func (o *GettableVolumeExport) SetApionly(v bool) {
	o.Apionly = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *GettableVolumeExport) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetAppendonly returns the Appendonly field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetAppendonly() bool {
	if o == nil || IsNil(o.Appendonly) {
		var ret bool
		return ret
	}
	return *o.Appendonly
}

// GetAppendonlyOk returns a tuple with the Appendonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetAppendonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Appendonly) {
		return nil, false
	}
	return o.Appendonly, true
}

// HasAppendonly returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasAppendonly() bool {
	if o != nil && !IsNil(o.Appendonly) {
		return true
	}

	return false
}

// SetAppendonly gets a reference to the given bool and assigns it to the Appendonly field.
func (o *GettableVolumeExport) SetAppendonly(v bool) {
	o.Appendonly = &v
}

// GetInternalip returns the Internalip field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetInternalip() bool {
	if o == nil || IsNil(o.Internalip) {
		var ret bool
		return ret
	}
	return *o.Internalip
}

// GetInternalipOk returns a tuple with the Internalip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetInternalipOk() (*bool, bool) {
	if o == nil || IsNil(o.Internalip) {
		return nil, false
	}
	return o.Internalip, true
}

// HasInternalip returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasInternalip() bool {
	if o != nil && !IsNil(o.Internalip) {
		return true
	}

	return false
}

// SetInternalip gets a reference to the given bool and assigns it to the Internalip field.
func (o *GettableVolumeExport) SetInternalip(v bool) {
	o.Internalip = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GettableVolumeExport) SetToken(v string) {
	o.Token = &v
}

// GetPasswd returns the Passwd field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetPasswd() string {
	if o == nil || IsNil(o.Passwd) {
		var ret string
		return ret
	}
	return *o.Passwd
}

// GetPasswdOk returns a tuple with the Passwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetPasswdOk() (*string, bool) {
	if o == nil || IsNil(o.Passwd) {
		return nil, false
	}
	return o.Passwd, true
}

// HasPasswd returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasPasswd() bool {
	if o != nil && !IsNil(o.Passwd) {
		return true
	}

	return false
}

// SetPasswd gets a reference to the given string and assigns it to the Passwd field.
func (o *GettableVolumeExport) SetPasswd(v string) {
	o.Passwd = &v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetQos() string {
	if o == nil || IsNil(o.Qos) {
		var ret string
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetQosOk() (*string, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given string and assigns it to the Qos field.
func (o *GettableVolumeExport) SetQos(v string) {
	o.Qos = &v
}

// GetExtend returns the Extend field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetExtend() string {
	if o == nil || IsNil(o.Extend) {
		var ret string
		return ret
	}
	return *o.Extend
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetExtendOk() (*string, bool) {
	if o == nil || IsNil(o.Extend) {
		return nil, false
	}
	return o.Extend, true
}

// HasExtend returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasExtend() bool {
	if o != nil && !IsNil(o.Extend) {
		return true
	}

	return false
}

// SetExtend gets a reference to the given string and assigns it to the Extend field.
func (o *GettableVolumeExport) SetExtend(v string) {
	o.Extend = &v
}

// GetSubdir returns the Subdir field value if set, zero value otherwise.
func (o *GettableVolumeExport) GetSubdir() string {
	if o == nil || IsNil(o.Subdir) {
		var ret string
		return ret
	}
	return *o.Subdir
}

// GetSubdirOk returns a tuple with the Subdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolumeExport) GetSubdirOk() (*string, bool) {
	if o == nil || IsNil(o.Subdir) {
		return nil, false
	}
	return o.Subdir, true
}

// HasSubdir returns a boolean if a field has been set.
func (o *GettableVolumeExport) HasSubdir() bool {
	if o != nil && !IsNil(o.Subdir) {
		return true
	}

	return false
}

// SetSubdir gets a reference to the given string and assigns it to the Subdir field.
func (o *GettableVolumeExport) SetSubdir(v string) {
	o.Subdir = &v
}

func (o GettableVolumeExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GettableVolumeExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Iprange) {
		toSerialize["iprange"] = o.Iprange
	}
	if !IsNil(o.Apionly) {
		toSerialize["apionly"] = o.Apionly
	}
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	if !IsNil(o.Appendonly) {
		toSerialize["appendonly"] = o.Appendonly
	}
	if !IsNil(o.Internalip) {
		toSerialize["internalip"] = o.Internalip
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Passwd) {
		toSerialize["passwd"] = o.Passwd
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Extend) {
		toSerialize["extend"] = o.Extend
	}
	if !IsNil(o.Subdir) {
		toSerialize["subdir"] = o.Subdir
	}
	return toSerialize, nil
}

type NullableGettableVolumeExport struct {
	value *GettableVolumeExport
	isSet bool
}

func (v NullableGettableVolumeExport) Get() *GettableVolumeExport {
	return v.value
}

func (v *NullableGettableVolumeExport) Set(val *GettableVolumeExport) {
	v.value = val
	v.isSet = true
}

func (v NullableGettableVolumeExport) IsSet() bool {
	return v.isSet
}

func (v *NullableGettableVolumeExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGettableVolumeExport(val *GettableVolumeExport) *NullableGettableVolumeExport {
	return &NullableGettableVolumeExport{value: val, isSet: true}
}

func (v NullableGettableVolumeExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGettableVolumeExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

