# PostableVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Region** | **int32** |  | 
**Bucket** | Pointer to **string** |  | [optional] [default to "juicefs-<Volume name>"]
**Trashtime** | Pointer to **int64** | Days to keep deleted files, set to zero to disable. Files in trash remains billable, learn more at &lt;a href&#x3D;\&quot;https://juicefs.com/docs/cloud/trash\&quot; target&#x3D;\&quot;_blank\&quot;&gt;our docs&lt;/a&gt;. | [optional] 
**BlockSize** | Pointer to **int32** |  | [optional] [default to 4096]
**Compress** | Pointer to **string** |  | [optional] [default to "lz4"]
**Compatible** | Pointer to **bool** |  | [optional] [default to false]
**Extend** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostableVolume

`func NewPostableVolume(name string, region int32, ) *PostableVolume`

NewPostableVolume instantiates a new PostableVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostableVolumeWithDefaults

`func NewPostableVolumeWithDefaults() *PostableVolume`

NewPostableVolumeWithDefaults instantiates a new PostableVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostableVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostableVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostableVolume) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *PostableVolume) GetRegion() int32`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostableVolume) GetRegionOk() (*int32, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostableVolume) SetRegion(v int32)`

SetRegion sets Region field to given value.


### GetBucket

`func (o *PostableVolume) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *PostableVolume) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *PostableVolume) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *PostableVolume) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetTrashtime

`func (o *PostableVolume) GetTrashtime() int64`

GetTrashtime returns the Trashtime field if non-nil, zero value otherwise.

### GetTrashtimeOk

`func (o *PostableVolume) GetTrashtimeOk() (*int64, bool)`

GetTrashtimeOk returns a tuple with the Trashtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashtime

`func (o *PostableVolume) SetTrashtime(v int64)`

SetTrashtime sets Trashtime field to given value.

### HasTrashtime

`func (o *PostableVolume) HasTrashtime() bool`

HasTrashtime returns a boolean if a field has been set.

### GetBlockSize

`func (o *PostableVolume) GetBlockSize() int32`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *PostableVolume) GetBlockSizeOk() (*int32, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *PostableVolume) SetBlockSize(v int32)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *PostableVolume) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetCompress

`func (o *PostableVolume) GetCompress() string`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *PostableVolume) GetCompressOk() (*string, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *PostableVolume) SetCompress(v string)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *PostableVolume) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetCompatible

`func (o *PostableVolume) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *PostableVolume) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *PostableVolume) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *PostableVolume) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetExtend

`func (o *PostableVolume) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *PostableVolume) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *PostableVolume) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *PostableVolume) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### GetStorage

`func (o *PostableVolume) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PostableVolume) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PostableVolume) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *PostableVolume) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *PostableVolume) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *PostableVolume) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


