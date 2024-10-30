# CreateMirrorVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Region** | **int32** |  | 
**Bucket** | Pointer to **string** | The bucket of the mirror volume. - If it is not provided or equal to the source volume&#39;s bucket, then the mirror volume and the source volume will share the same bucket. - Otherwise, the mirror volume&#39;s bucket will replicate the source volume&#39;s bucket.  | [optional] [default to ""]
**Trashtime** | Pointer to **int64** | Days to keep deleted files, set to zero to disable. Files in trash remains billable, learn more at &lt;a href&#x3D;\&quot;https://juicefs.com/docs/cloud/trash\&quot; target&#x3D;\&quot;_blank\&quot;&gt;our docs&lt;/a&gt;. | [optional] 
**BlockSize** | Pointer to **int32** |  | [optional] [default to 4096]
**Compress** | Pointer to **string** |  | [optional] [default to "lz4"]
**Compatible** | Pointer to **bool** |  | [optional] [default to false]
**Extend** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateMirrorVolume

`func NewCreateMirrorVolume(name string, region int32, ) *CreateMirrorVolume`

NewCreateMirrorVolume instantiates a new CreateMirrorVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMirrorVolumeWithDefaults

`func NewCreateMirrorVolumeWithDefaults() *CreateMirrorVolume`

NewCreateMirrorVolumeWithDefaults instantiates a new CreateMirrorVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMirrorVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMirrorVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMirrorVolume) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *CreateMirrorVolume) GetRegion() int32`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateMirrorVolume) GetRegionOk() (*int32, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateMirrorVolume) SetRegion(v int32)`

SetRegion sets Region field to given value.


### GetBucket

`func (o *CreateMirrorVolume) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CreateMirrorVolume) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CreateMirrorVolume) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CreateMirrorVolume) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetTrashtime

`func (o *CreateMirrorVolume) GetTrashtime() int64`

GetTrashtime returns the Trashtime field if non-nil, zero value otherwise.

### GetTrashtimeOk

`func (o *CreateMirrorVolume) GetTrashtimeOk() (*int64, bool)`

GetTrashtimeOk returns a tuple with the Trashtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashtime

`func (o *CreateMirrorVolume) SetTrashtime(v int64)`

SetTrashtime sets Trashtime field to given value.

### HasTrashtime

`func (o *CreateMirrorVolume) HasTrashtime() bool`

HasTrashtime returns a boolean if a field has been set.

### GetBlockSize

`func (o *CreateMirrorVolume) GetBlockSize() int32`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *CreateMirrorVolume) GetBlockSizeOk() (*int32, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *CreateMirrorVolume) SetBlockSize(v int32)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *CreateMirrorVolume) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetCompress

`func (o *CreateMirrorVolume) GetCompress() string`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *CreateMirrorVolume) GetCompressOk() (*string, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *CreateMirrorVolume) SetCompress(v string)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *CreateMirrorVolume) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetCompatible

`func (o *CreateMirrorVolume) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *CreateMirrorVolume) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *CreateMirrorVolume) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *CreateMirrorVolume) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetExtend

`func (o *CreateMirrorVolume) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *CreateMirrorVolume) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *CreateMirrorVolume) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *CreateMirrorVolume) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### GetStorage

`func (o *CreateMirrorVolume) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateMirrorVolume) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateMirrorVolume) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateMirrorVolume) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *CreateMirrorVolume) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *CreateMirrorVolume) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


