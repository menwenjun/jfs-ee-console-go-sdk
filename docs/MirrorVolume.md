# MirrorVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AccessRules** | Pointer to [**[]GettableVolumeAllOfAccessRules**](GettableVolumeAllOfAccessRules.md) |  | [optional] 
**Owner** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Inodes** | Pointer to **NullableInt64** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **int32** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] [default to "juicefs-<Volume name>"]
**Trashtime** | Pointer to **int64** | Days to keep deleted files, set to zero to disable. Files in trash remains billable, learn more at &lt;a href&#x3D;\&quot;https://juicefs.com/docs/cloud/trash\&quot; target&#x3D;\&quot;_blank\&quot;&gt;our docs&lt;/a&gt;. | [optional] 
**BlockSize** | Pointer to **int32** |  | [optional] [default to 4096]
**Compress** | Pointer to **string** |  | [optional] [default to "lz4"]
**Compatible** | Pointer to **bool** |  | [optional] [default to false]
**Extend** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **int32** |  | [optional] 

## Methods

### NewMirrorVolume

`func NewMirrorVolume() *MirrorVolume`

NewMirrorVolume instantiates a new MirrorVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMirrorVolumeWithDefaults

`func NewMirrorVolumeWithDefaults() *MirrorVolume`

NewMirrorVolumeWithDefaults instantiates a new MirrorVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MirrorVolume) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MirrorVolume) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MirrorVolume) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MirrorVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessRules

`func (o *MirrorVolume) GetAccessRules() []GettableVolumeAllOfAccessRules`

GetAccessRules returns the AccessRules field if non-nil, zero value otherwise.

### GetAccessRulesOk

`func (o *MirrorVolume) GetAccessRulesOk() (*[]GettableVolumeAllOfAccessRules, bool)`

GetAccessRulesOk returns a tuple with the AccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRules

`func (o *MirrorVolume) SetAccessRules(v []GettableVolumeAllOfAccessRules)`

SetAccessRules sets AccessRules field to given value.

### HasAccessRules

`func (o *MirrorVolume) HasAccessRules() bool`

HasAccessRules returns a boolean if a field has been set.

### GetOwner

`func (o *MirrorVolume) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MirrorVolume) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MirrorVolume) SetOwner(v int32)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MirrorVolume) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSize

`func (o *MirrorVolume) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MirrorVolume) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MirrorVolume) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MirrorVolume) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MirrorVolume) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MirrorVolume) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetInodes

`func (o *MirrorVolume) GetInodes() int64`

GetInodes returns the Inodes field if non-nil, zero value otherwise.

### GetInodesOk

`func (o *MirrorVolume) GetInodesOk() (*int64, bool)`

GetInodesOk returns a tuple with the Inodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInodes

`func (o *MirrorVolume) SetInodes(v int64)`

SetInodes sets Inodes field to given value.

### HasInodes

`func (o *MirrorVolume) HasInodes() bool`

HasInodes returns a boolean if a field has been set.

### SetInodesNil

`func (o *MirrorVolume) SetInodesNil(b bool)`

 SetInodesNil sets the value for Inodes to be an explicit nil

### UnsetInodes
`func (o *MirrorVolume) UnsetInodes()`

UnsetInodes ensures that no value is present for Inodes, not even an explicit nil
### GetCreated

`func (o *MirrorVolume) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MirrorVolume) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MirrorVolume) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MirrorVolume) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUuid

`func (o *MirrorVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MirrorVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MirrorVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MirrorVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *MirrorVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MirrorVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MirrorVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MirrorVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *MirrorVolume) GetRegion() int32`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MirrorVolume) GetRegionOk() (*int32, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MirrorVolume) SetRegion(v int32)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MirrorVolume) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetBucket

`func (o *MirrorVolume) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *MirrorVolume) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *MirrorVolume) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *MirrorVolume) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetTrashtime

`func (o *MirrorVolume) GetTrashtime() int64`

GetTrashtime returns the Trashtime field if non-nil, zero value otherwise.

### GetTrashtimeOk

`func (o *MirrorVolume) GetTrashtimeOk() (*int64, bool)`

GetTrashtimeOk returns a tuple with the Trashtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashtime

`func (o *MirrorVolume) SetTrashtime(v int64)`

SetTrashtime sets Trashtime field to given value.

### HasTrashtime

`func (o *MirrorVolume) HasTrashtime() bool`

HasTrashtime returns a boolean if a field has been set.

### GetBlockSize

`func (o *MirrorVolume) GetBlockSize() int32`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *MirrorVolume) GetBlockSizeOk() (*int32, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *MirrorVolume) SetBlockSize(v int32)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *MirrorVolume) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetCompress

`func (o *MirrorVolume) GetCompress() string`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *MirrorVolume) GetCompressOk() (*string, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *MirrorVolume) SetCompress(v string)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *MirrorVolume) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetCompatible

`func (o *MirrorVolume) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *MirrorVolume) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *MirrorVolume) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *MirrorVolume) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetExtend

`func (o *MirrorVolume) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *MirrorVolume) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *MirrorVolume) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *MirrorVolume) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### GetStorage

`func (o *MirrorVolume) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *MirrorVolume) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *MirrorVolume) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *MirrorVolume) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *MirrorVolume) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *MirrorVolume) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetSource

`func (o *MirrorVolume) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MirrorVolume) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MirrorVolume) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *MirrorVolume) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


