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

// checks if the Region type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Region{}

// Region struct for Region
type Region struct {
	Id *int32 `json:"id,omitempty"`
	Cloud *int32 `json:"cloud,omitempty"`
	Name *string `json:"name,omitempty"`
	Desp *string `json:"desp,omitempty"`
	Owner *int32 `json:"owner,omitempty"`
	Token *string `json:"token,omitempty"`
	Trashtime *int64 `json:"trashtime,omitempty"`
}

// NewRegion instantiates a new Region object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegion() *Region {
	this := Region{}
	return &this
}

// NewRegionWithDefaults instantiates a new Region object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionWithDefaults() *Region {
	this := Region{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Region) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Region) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Region) SetId(v int32) {
	o.Id = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *Region) GetCloud() int32 {
	if o == nil || IsNil(o.Cloud) {
		var ret int32
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetCloudOk() (*int32, bool) {
	if o == nil || IsNil(o.Cloud) {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *Region) HasCloud() bool {
	if o != nil && !IsNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given int32 and assigns it to the Cloud field.
func (o *Region) SetCloud(v int32) {
	o.Cloud = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Region) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Region) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Region) SetName(v string) {
	o.Name = &v
}

// GetDesp returns the Desp field value if set, zero value otherwise.
func (o *Region) GetDesp() string {
	if o == nil || IsNil(o.Desp) {
		var ret string
		return ret
	}
	return *o.Desp
}

// GetDespOk returns a tuple with the Desp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetDespOk() (*string, bool) {
	if o == nil || IsNil(o.Desp) {
		return nil, false
	}
	return o.Desp, true
}

// HasDesp returns a boolean if a field has been set.
func (o *Region) HasDesp() bool {
	if o != nil && !IsNil(o.Desp) {
		return true
	}

	return false
}

// SetDesp gets a reference to the given string and assigns it to the Desp field.
func (o *Region) SetDesp(v string) {
	o.Desp = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Region) GetOwner() int32 {
	if o == nil || IsNil(o.Owner) {
		var ret int32
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetOwnerOk() (*int32, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Region) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given int32 and assigns it to the Owner field.
func (o *Region) SetOwner(v int32) {
	o.Owner = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Region) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Region) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Region) SetToken(v string) {
	o.Token = &v
}

// GetTrashtime returns the Trashtime field value if set, zero value otherwise.
func (o *Region) GetTrashtime() int64 {
	if o == nil || IsNil(o.Trashtime) {
		var ret int64
		return ret
	}
	return *o.Trashtime
}

// GetTrashtimeOk returns a tuple with the Trashtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetTrashtimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Trashtime) {
		return nil, false
	}
	return o.Trashtime, true
}

// HasTrashtime returns a boolean if a field has been set.
func (o *Region) HasTrashtime() bool {
	if o != nil && !IsNil(o.Trashtime) {
		return true
	}

	return false
}

// SetTrashtime gets a reference to the given int64 and assigns it to the Trashtime field.
func (o *Region) SetTrashtime(v int64) {
	o.Trashtime = &v
}

func (o Region) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Region) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Cloud) {
		toSerialize["cloud"] = o.Cloud
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Desp) {
		toSerialize["desp"] = o.Desp
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Trashtime) {
		toSerialize["trashtime"] = o.Trashtime
	}
	return toSerialize, nil
}

type NullableRegion struct {
	value *Region
	isSet bool
}

func (v NullableRegion) Get() *Region {
	return v.value
}

func (v *NullableRegion) Set(val *Region) {
	v.value = val
	v.isSet = true
}

func (v NullableRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegion(val *Region) *NullableRegion {
	return &NullableRegion{value: val, isSet: true}
}

func (v NullableRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


