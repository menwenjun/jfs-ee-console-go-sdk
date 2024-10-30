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
	"time"
)

// checks if the GettableVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GettableVolume{}

// GettableVolume struct for GettableVolume
type GettableVolume struct {
	Name *string `json:"name,omitempty" validate:"regexp=^[a-z][0-9a-z-]{1,38}[0-9a-z]$"`
	Region *int32 `json:"region,omitempty"`
	Bucket *string `json:"bucket,omitempty"`
	// Days to keep deleted files, set to zero to disable. Files in trash remains billable, learn more at <a href=\"https://juicefs.com/docs/cloud/trash\" target=\"_blank\">our docs</a>.
	Trashtime *int64 `json:"trashtime,omitempty"`
	BlockSize *int32 `json:"blockSize,omitempty"`
	Compress *string `json:"compress,omitempty"`
	Compatible *bool `json:"compatible,omitempty"`
	Extend *string `json:"extend,omitempty"`
	Storage NullableString `json:"storage,omitempty"`
	Id *int32 `json:"id,omitempty"`
	AccessRules []GettableVolumeAllOfAccessRules `json:"access_rules,omitempty"`
	Owner *int32 `json:"owner,omitempty"`
	Size NullableInt64 `json:"size,omitempty"`
	Inodes NullableInt32 `json:"inodes,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewGettableVolume instantiates a new GettableVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGettableVolume() *GettableVolume {
	this := GettableVolume{}
	var bucket string = "juicefs-<Volume name>"
	this.Bucket = &bucket
	var blockSize int32 = 4096
	this.BlockSize = &blockSize
	var compress string = "lz4"
	this.Compress = &compress
	var compatible bool = false
	this.Compatible = &compatible
	return &this
}

// NewGettableVolumeWithDefaults instantiates a new GettableVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGettableVolumeWithDefaults() *GettableVolume {
	this := GettableVolume{}
	var bucket string = "juicefs-<Volume name>"
	this.Bucket = &bucket
	var blockSize int32 = 4096
	this.BlockSize = &blockSize
	var compress string = "lz4"
	this.Compress = &compress
	var compatible bool = false
	this.Compatible = &compatible
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GettableVolume) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GettableVolume) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GettableVolume) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GettableVolume) GetRegion() int32 {
	if o == nil || IsNil(o.Region) {
		var ret int32
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetRegionOk() (*int32, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GettableVolume) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given int32 and assigns it to the Region field.
func (o *GettableVolume) SetRegion(v int32) {
	o.Region = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *GettableVolume) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *GettableVolume) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *GettableVolume) SetBucket(v string) {
	o.Bucket = &v
}

// GetTrashtime returns the Trashtime field value if set, zero value otherwise.
func (o *GettableVolume) GetTrashtime() int64 {
	if o == nil || IsNil(o.Trashtime) {
		var ret int64
		return ret
	}
	return *o.Trashtime
}

// GetTrashtimeOk returns a tuple with the Trashtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetTrashtimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Trashtime) {
		return nil, false
	}
	return o.Trashtime, true
}

// HasTrashtime returns a boolean if a field has been set.
func (o *GettableVolume) HasTrashtime() bool {
	if o != nil && !IsNil(o.Trashtime) {
		return true
	}

	return false
}

// SetTrashtime gets a reference to the given int64 and assigns it to the Trashtime field.
func (o *GettableVolume) SetTrashtime(v int64) {
	o.Trashtime = &v
}

// GetBlockSize returns the BlockSize field value if set, zero value otherwise.
func (o *GettableVolume) GetBlockSize() int32 {
	if o == nil || IsNil(o.BlockSize) {
		var ret int32
		return ret
	}
	return *o.BlockSize
}

// GetBlockSizeOk returns a tuple with the BlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetBlockSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockSize) {
		return nil, false
	}
	return o.BlockSize, true
}

// HasBlockSize returns a boolean if a field has been set.
func (o *GettableVolume) HasBlockSize() bool {
	if o != nil && !IsNil(o.BlockSize) {
		return true
	}

	return false
}

// SetBlockSize gets a reference to the given int32 and assigns it to the BlockSize field.
func (o *GettableVolume) SetBlockSize(v int32) {
	o.BlockSize = &v
}

// GetCompress returns the Compress field value if set, zero value otherwise.
func (o *GettableVolume) GetCompress() string {
	if o == nil || IsNil(o.Compress) {
		var ret string
		return ret
	}
	return *o.Compress
}

// GetCompressOk returns a tuple with the Compress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetCompressOk() (*string, bool) {
	if o == nil || IsNil(o.Compress) {
		return nil, false
	}
	return o.Compress, true
}

// HasCompress returns a boolean if a field has been set.
func (o *GettableVolume) HasCompress() bool {
	if o != nil && !IsNil(o.Compress) {
		return true
	}

	return false
}

// SetCompress gets a reference to the given string and assigns it to the Compress field.
func (o *GettableVolume) SetCompress(v string) {
	o.Compress = &v
}

// GetCompatible returns the Compatible field value if set, zero value otherwise.
func (o *GettableVolume) GetCompatible() bool {
	if o == nil || IsNil(o.Compatible) {
		var ret bool
		return ret
	}
	return *o.Compatible
}

// GetCompatibleOk returns a tuple with the Compatible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetCompatibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Compatible) {
		return nil, false
	}
	return o.Compatible, true
}

// HasCompatible returns a boolean if a field has been set.
func (o *GettableVolume) HasCompatible() bool {
	if o != nil && !IsNil(o.Compatible) {
		return true
	}

	return false
}

// SetCompatible gets a reference to the given bool and assigns it to the Compatible field.
func (o *GettableVolume) SetCompatible(v bool) {
	o.Compatible = &v
}

// GetExtend returns the Extend field value if set, zero value otherwise.
func (o *GettableVolume) GetExtend() string {
	if o == nil || IsNil(o.Extend) {
		var ret string
		return ret
	}
	return *o.Extend
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetExtendOk() (*string, bool) {
	if o == nil || IsNil(o.Extend) {
		return nil, false
	}
	return o.Extend, true
}

// HasExtend returns a boolean if a field has been set.
func (o *GettableVolume) HasExtend() bool {
	if o != nil && !IsNil(o.Extend) {
		return true
	}

	return false
}

// SetExtend gets a reference to the given string and assigns it to the Extend field.
func (o *GettableVolume) SetExtend(v string) {
	o.Extend = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GettableVolume) GetStorage() string {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret string
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GettableVolume) GetStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *GettableVolume) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableString and assigns it to the Storage field.
func (o *GettableVolume) SetStorage(v string) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *GettableVolume) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *GettableVolume) UnsetStorage() {
	o.Storage.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GettableVolume) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GettableVolume) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GettableVolume) SetId(v int32) {
	o.Id = &v
}

// GetAccessRules returns the AccessRules field value if set, zero value otherwise.
func (o *GettableVolume) GetAccessRules() []GettableVolumeAllOfAccessRules {
	if o == nil || IsNil(o.AccessRules) {
		var ret []GettableVolumeAllOfAccessRules
		return ret
	}
	return o.AccessRules
}

// GetAccessRulesOk returns a tuple with the AccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetAccessRulesOk() ([]GettableVolumeAllOfAccessRules, bool) {
	if o == nil || IsNil(o.AccessRules) {
		return nil, false
	}
	return o.AccessRules, true
}

// HasAccessRules returns a boolean if a field has been set.
func (o *GettableVolume) HasAccessRules() bool {
	if o != nil && !IsNil(o.AccessRules) {
		return true
	}

	return false
}

// SetAccessRules gets a reference to the given []GettableVolumeAllOfAccessRules and assigns it to the AccessRules field.
func (o *GettableVolume) SetAccessRules(v []GettableVolumeAllOfAccessRules) {
	o.AccessRules = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *GettableVolume) GetOwner() int32 {
	if o == nil || IsNil(o.Owner) {
		var ret int32
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetOwnerOk() (*int32, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *GettableVolume) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given int32 and assigns it to the Owner field.
func (o *GettableVolume) SetOwner(v int32) {
	o.Owner = &v
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GettableVolume) GetSize() int64 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int64
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GettableVolume) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *GettableVolume) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt64 and assigns it to the Size field.
func (o *GettableVolume) SetSize(v int64) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *GettableVolume) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *GettableVolume) UnsetSize() {
	o.Size.Unset()
}

// GetInodes returns the Inodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GettableVolume) GetInodes() int32 {
	if o == nil || IsNil(o.Inodes.Get()) {
		var ret int32
		return ret
	}
	return *o.Inodes.Get()
}

// GetInodesOk returns a tuple with the Inodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GettableVolume) GetInodesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inodes.Get(), o.Inodes.IsSet()
}

// HasInodes returns a boolean if a field has been set.
func (o *GettableVolume) HasInodes() bool {
	if o != nil && o.Inodes.IsSet() {
		return true
	}

	return false
}

// SetInodes gets a reference to the given NullableInt32 and assigns it to the Inodes field.
func (o *GettableVolume) SetInodes(v int32) {
	o.Inodes.Set(&v)
}
// SetInodesNil sets the value for Inodes to be an explicit nil
func (o *GettableVolume) SetInodesNil() {
	o.Inodes.Set(nil)
}

// UnsetInodes ensures that no value is present for Inodes, not even an explicit nil
func (o *GettableVolume) UnsetInodes() {
	o.Inodes.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GettableVolume) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GettableVolume) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GettableVolume) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GettableVolume) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GettableVolume) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GettableVolume) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GettableVolume) SetUuid(v string) {
	o.Uuid = &v
}

func (o GettableVolume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GettableVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Trashtime) {
		toSerialize["trashtime"] = o.Trashtime
	}
	if !IsNil(o.BlockSize) {
		toSerialize["blockSize"] = o.BlockSize
	}
	if !IsNil(o.Compress) {
		toSerialize["compress"] = o.Compress
	}
	if !IsNil(o.Compatible) {
		toSerialize["compatible"] = o.Compatible
	}
	if !IsNil(o.Extend) {
		toSerialize["extend"] = o.Extend
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccessRules) {
		toSerialize["access_rules"] = o.AccessRules
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.Inodes.IsSet() {
		toSerialize["inodes"] = o.Inodes.Get()
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableGettableVolume struct {
	value *GettableVolume
	isSet bool
}

func (v NullableGettableVolume) Get() *GettableVolume {
	return v.value
}

func (v *NullableGettableVolume) Set(val *GettableVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableGettableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableGettableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGettableVolume(val *GettableVolume) *NullableGettableVolume {
	return &NullableGettableVolume{value: val, isSet: true}
}

func (v NullableGettableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGettableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


