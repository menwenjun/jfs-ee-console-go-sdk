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

// checks if the VolumesVolumeIDListTrashGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumesVolumeIDListTrashGet200Response{}

// VolumesVolumeIDListTrashGet200Response struct for VolumesVolumeIDListTrashGet200Response
type VolumesVolumeIDListTrashGet200Response struct {
	Total *int32 `json:"total,omitempty"`
	TotalBytes *int32 `json:"totalBytes,omitempty"`
	Query *string `json:"query,omitempty"`
	CurPage *int32 `json:"curPage,omitempty"`
	Files []VolumesVolumeIDListTrashGet200ResponseFilesInner `json:"files,omitempty"`
}

// NewVolumesVolumeIDListTrashGet200Response instantiates a new VolumesVolumeIDListTrashGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumesVolumeIDListTrashGet200Response() *VolumesVolumeIDListTrashGet200Response {
	this := VolumesVolumeIDListTrashGet200Response{}
	return &this
}

// NewVolumesVolumeIDListTrashGet200ResponseWithDefaults instantiates a new VolumesVolumeIDListTrashGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumesVolumeIDListTrashGet200ResponseWithDefaults() *VolumesVolumeIDListTrashGet200Response {
	this := VolumesVolumeIDListTrashGet200Response{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *VolumesVolumeIDListTrashGet200Response) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDListTrashGet200Response) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *VolumesVolumeIDListTrashGet200Response) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *VolumesVolumeIDListTrashGet200Response) SetTotal(v int32) {
	o.Total = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *VolumesVolumeIDListTrashGet200Response) GetTotalBytes() int32 {
	if o == nil || IsNil(o.TotalBytes) {
		var ret int32
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDListTrashGet200Response) GetTotalBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalBytes) {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *VolumesVolumeIDListTrashGet200Response) HasTotalBytes() bool {
	if o != nil && !IsNil(o.TotalBytes) {
		return true
	}

	return false
}

// SetTotalBytes gets a reference to the given int32 and assigns it to the TotalBytes field.
func (o *VolumesVolumeIDListTrashGet200Response) SetTotalBytes(v int32) {
	o.TotalBytes = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *VolumesVolumeIDListTrashGet200Response) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDListTrashGet200Response) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *VolumesVolumeIDListTrashGet200Response) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *VolumesVolumeIDListTrashGet200Response) SetQuery(v string) {
	o.Query = &v
}

// GetCurPage returns the CurPage field value if set, zero value otherwise.
func (o *VolumesVolumeIDListTrashGet200Response) GetCurPage() int32 {
	if o == nil || IsNil(o.CurPage) {
		var ret int32
		return ret
	}
	return *o.CurPage
}

// GetCurPageOk returns a tuple with the CurPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDListTrashGet200Response) GetCurPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurPage) {
		return nil, false
	}
	return o.CurPage, true
}

// HasCurPage returns a boolean if a field has been set.
func (o *VolumesVolumeIDListTrashGet200Response) HasCurPage() bool {
	if o != nil && !IsNil(o.CurPage) {
		return true
	}

	return false
}

// SetCurPage gets a reference to the given int32 and assigns it to the CurPage field.
func (o *VolumesVolumeIDListTrashGet200Response) SetCurPage(v int32) {
	o.CurPage = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *VolumesVolumeIDListTrashGet200Response) GetFiles() []VolumesVolumeIDListTrashGet200ResponseFilesInner {
	if o == nil || IsNil(o.Files) {
		var ret []VolumesVolumeIDListTrashGet200ResponseFilesInner
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumesVolumeIDListTrashGet200Response) GetFilesOk() ([]VolumesVolumeIDListTrashGet200ResponseFilesInner, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *VolumesVolumeIDListTrashGet200Response) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []VolumesVolumeIDListTrashGet200ResponseFilesInner and assigns it to the Files field.
func (o *VolumesVolumeIDListTrashGet200Response) SetFiles(v []VolumesVolumeIDListTrashGet200ResponseFilesInner) {
	o.Files = v
}

func (o VolumesVolumeIDListTrashGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumesVolumeIDListTrashGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.TotalBytes) {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.CurPage) {
		toSerialize["curPage"] = o.CurPage
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

type NullableVolumesVolumeIDListTrashGet200Response struct {
	value *VolumesVolumeIDListTrashGet200Response
	isSet bool
}

func (v NullableVolumesVolumeIDListTrashGet200Response) Get() *VolumesVolumeIDListTrashGet200Response {
	return v.value
}

func (v *NullableVolumesVolumeIDListTrashGet200Response) Set(val *VolumesVolumeIDListTrashGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumesVolumeIDListTrashGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumesVolumeIDListTrashGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumesVolumeIDListTrashGet200Response(val *VolumesVolumeIDListTrashGet200Response) *NullableVolumesVolumeIDListTrashGet200Response {
	return &NullableVolumesVolumeIDListTrashGet200Response{value: val, isSet: true}
}

func (v NullableVolumesVolumeIDListTrashGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumesVolumeIDListTrashGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


