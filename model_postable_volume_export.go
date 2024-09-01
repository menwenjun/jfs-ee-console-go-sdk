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

// checks if the PostableVolumeExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostableVolumeExport{}

// PostableVolumeExport struct for PostableVolumeExport
type PostableVolumeExport struct {
	Desc *string `json:"desc,omitempty"`
	// allowed IP format:   - any IP address: *   - Single IP (dot-decimal): A.B.C.D   - IP range (dot-decimal): A.B.C.D-E.F.G.H   - IP range (CIDR): A.B.C.D/[0-32] 
	Iprange *string `json:"iprange,omitempty"`
	Readonly *bool `json:"readonly,omitempty"`
	Appendonly *bool `json:"appendonly,omitempty"`
	Internalip *bool `json:"internalip,omitempty"`
	// qos format: `<put>[:<get>[:<compact>]]`, put, get and compact has the same format `<number>[K|M|G|P|T|Z|E]`, 0 or empty value means unlimited.  the unit for qos is `B/s`.  examples: - `10M`, set put to 10MB/s, get and compact to unlimited - `10M:10M`, set put to 10MB/s, get to 10MB/s and compact to unlimited - `1k:1m:10m`, set put to 1KB/s, get to 1MB/s and compact to 10MB/s. the suffix(k, m) is case-insensitive. - `10M:0:10M`, set put to 10MB/s, get to unlimited and compact to 10MB/s 
	Qos *string `json:"qos,omitempty"`
	Extend *string `json:"extend,omitempty"`
	Subdir *string `json:"subdir,omitempty"`
}

// NewPostableVolumeExport instantiates a new PostableVolumeExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostableVolumeExport() *PostableVolumeExport {
	this := PostableVolumeExport{}
	var iprange string = "*"
	this.Iprange = &iprange
	var readonly bool = false
	this.Readonly = &readonly
	var appendonly bool = false
	this.Appendonly = &appendonly
	var internalip bool = false
	this.Internalip = &internalip
	var qos string = ""
	this.Qos = &qos
	var extend string = ""
	this.Extend = &extend
	var subdir string = "/"
	this.Subdir = &subdir
	return &this
}

// NewPostableVolumeExportWithDefaults instantiates a new PostableVolumeExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostableVolumeExportWithDefaults() *PostableVolumeExport {
	this := PostableVolumeExport{}
	var iprange string = "*"
	this.Iprange = &iprange
	var readonly bool = false
	this.Readonly = &readonly
	var appendonly bool = false
	this.Appendonly = &appendonly
	var internalip bool = false
	this.Internalip = &internalip
	var qos string = ""
	this.Qos = &qos
	var extend string = ""
	this.Extend = &extend
	var subdir string = "/"
	this.Subdir = &subdir
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *PostableVolumeExport) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostableVolumeExport) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *PostableVolumeExport) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *PostableVolumeExport) SetDesc(v string) {
	o.Desc = &v
}

// GetIprange returns the Iprange field value if set, zero value otherwise.
func (o *PostableVolumeExport) GetIprange() string {
	if o == nil || IsNil(o.Iprange) {
		var ret string
		return ret
	}
	return *o.Iprange
}

// GetIprangeOk returns a tuple with the Iprange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostableVolumeExport) GetIprangeOk() (*string, bool) {
	if o == nil || IsNil(o.Iprange) {
		return nil, false
	}
	return o.Iprange, true
}

// HasIprange returns a boolean if a field has been set.
func (o *PostableVolumeExport) HasIprange() bool {
	if o != nil && !IsNil(o.Iprange) {
		return true
	}

	return false
}

// SetIprange gets a reference to the given string and assigns it to the Iprange field.
func (o *PostableVolumeExport) SetIprange(v string) {
	o.Iprange = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *PostableVolumeExport) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostableVolumeExport) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *PostableVolumeExport) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *PostableVolumeExport) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetAppendonly returns the Appendonly field value if set, zero value otherwise.
func (o *PostableVolumeExport) GetAppendonly() bool {
	if o == nil || IsNil(o.Appendonly) {
		var ret bool
		return ret
	}
	return *o.Appendonly
}

// GetAppendonlyOk returns a tuple with the Appendonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostableVolumeExport) GetAppendonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Appendonly) {
		return nil, false
	}
	return o.Appendonly, true
}

// HasAppendonly returns a boolean if a field has been set.
func (o *PostableVolumeExport) HasAppendonly() bool {
	if o != nil && !IsNil(o.Appendonly) {
		return true
	}

	return false
}

// SetAppendonly gets a reference to the given bool and assigns it to the Appendonly field.
func (o *PostableVolumeExport) SetAppendonly(v bool) {
	o.Appendonly = &v
}

// GetInternalip returns the Internalip field value if set, zero value otherwise.
func (o *PostableVolumeExport) GetInternalip() bool {
	if o == nil || IsNil(o.Internalip) {
		var ret bool
		return ret
	}
	return *o.Internalip
}

// GetInternalipOk returns a tuple with the Internalip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostableVolumeExport) GetInternalipOk() (*bool, bool) {
	if o == nil || IsNil(o.Internalip) {
		return nil, false
	}
	return o.Internalip, true
}

// HasInternalip returns a boolean if a field has been set.
func (o *PostableVolumeExport) HasInternalip() bool {
	if o != nil && !IsNil(o.Internalip) {
		return true
	}

	return false
}

// SetInternalip gets a reference to the given bool and assigns it to the Internalip field.
func (o *PostableVolumeExport) SetInternalip(v bool) {
	o.Internalip = &v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *PostableVolumeExport) GetQos() string {
	if o == nil || IsNil(o.Qos) {
		var ret string
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostableVolumeExport) GetQosOk() (*string, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *PostableVolumeExport) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given string and assigns it to the Qos field.
func (o *PostableVolumeExport) SetQos(v string) {
	o.Qos = &v
}

// GetExtend returns the Extend field value if set, zero value otherwise.
func (o *PostableVolumeExport) GetExtend() string {
	if o == nil || IsNil(o.Extend) {
		var ret string
		return ret
	}
	return *o.Extend
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostableVolumeExport) GetExtendOk() (*string, bool) {
	if o == nil || IsNil(o.Extend) {
		return nil, false
	}
	return o.Extend, true
}

// HasExtend returns a boolean if a field has been set.
func (o *PostableVolumeExport) HasExtend() bool {
	if o != nil && !IsNil(o.Extend) {
		return true
	}

	return false
}

// SetExtend gets a reference to the given string and assigns it to the Extend field.
func (o *PostableVolumeExport) SetExtend(v string) {
	o.Extend = &v
}

// GetSubdir returns the Subdir field value if set, zero value otherwise.
func (o *PostableVolumeExport) GetSubdir() string {
	if o == nil || IsNil(o.Subdir) {
		var ret string
		return ret
	}
	return *o.Subdir
}

// GetSubdirOk returns a tuple with the Subdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostableVolumeExport) GetSubdirOk() (*string, bool) {
	if o == nil || IsNil(o.Subdir) {
		return nil, false
	}
	return o.Subdir, true
}

// HasSubdir returns a boolean if a field has been set.
func (o *PostableVolumeExport) HasSubdir() bool {
	if o != nil && !IsNil(o.Subdir) {
		return true
	}

	return false
}

// SetSubdir gets a reference to the given string and assigns it to the Subdir field.
func (o *PostableVolumeExport) SetSubdir(v string) {
	o.Subdir = &v
}

func (o PostableVolumeExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostableVolumeExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Iprange) {
		toSerialize["iprange"] = o.Iprange
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

type NullablePostableVolumeExport struct {
	value *PostableVolumeExport
	isSet bool
}

func (v NullablePostableVolumeExport) Get() *PostableVolumeExport {
	return v.value
}

func (v *NullablePostableVolumeExport) Set(val *PostableVolumeExport) {
	v.value = val
	v.isSet = true
}

func (v NullablePostableVolumeExport) IsSet() bool {
	return v.isSet
}

func (v *NullablePostableVolumeExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostableVolumeExport(val *PostableVolumeExport) *NullablePostableVolumeExport {
	return &NullablePostableVolumeExport{value: val, isSet: true}
}

func (v NullablePostableVolumeExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostableVolumeExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


