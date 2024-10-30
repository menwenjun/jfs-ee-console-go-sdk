# GettableVolumeQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Path** | **string** |  | 
**Inodes** | **int32** |  | 
**Size** | **int64** |  | 
**UsedInodes** | Pointer to **int32** |  | [optional] 
**UsedSize** | Pointer to **int64** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGettableVolumeQuota

`func NewGettableVolumeQuota(id int32, path string, inodes int32, size int64, ) *GettableVolumeQuota`

NewGettableVolumeQuota instantiates a new GettableVolumeQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGettableVolumeQuotaWithDefaults

`func NewGettableVolumeQuotaWithDefaults() *GettableVolumeQuota`

NewGettableVolumeQuotaWithDefaults instantiates a new GettableVolumeQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GettableVolumeQuota) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GettableVolumeQuota) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GettableVolumeQuota) SetId(v int32)`

SetId sets Id field to given value.


### GetPath

`func (o *GettableVolumeQuota) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GettableVolumeQuota) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GettableVolumeQuota) SetPath(v string)`

SetPath sets Path field to given value.


### GetInodes

`func (o *GettableVolumeQuota) GetInodes() int32`

GetInodes returns the Inodes field if non-nil, zero value otherwise.

### GetInodesOk

`func (o *GettableVolumeQuota) GetInodesOk() (*int32, bool)`

GetInodesOk returns a tuple with the Inodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInodes

`func (o *GettableVolumeQuota) SetInodes(v int32)`

SetInodes sets Inodes field to given value.


### GetSize

`func (o *GettableVolumeQuota) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GettableVolumeQuota) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GettableVolumeQuota) SetSize(v int64)`

SetSize sets Size field to given value.


### GetUsedInodes

`func (o *GettableVolumeQuota) GetUsedInodes() int32`

GetUsedInodes returns the UsedInodes field if non-nil, zero value otherwise.

### GetUsedInodesOk

`func (o *GettableVolumeQuota) GetUsedInodesOk() (*int32, bool)`

GetUsedInodesOk returns a tuple with the UsedInodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInodes

`func (o *GettableVolumeQuota) SetUsedInodes(v int32)`

SetUsedInodes sets UsedInodes field to given value.

### HasUsedInodes

`func (o *GettableVolumeQuota) HasUsedInodes() bool`

HasUsedInodes returns a boolean if a field has been set.

### GetUsedSize

`func (o *GettableVolumeQuota) GetUsedSize() int64`

GetUsedSize returns the UsedSize field if non-nil, zero value otherwise.

### GetUsedSizeOk

`func (o *GettableVolumeQuota) GetUsedSizeOk() (*int64, bool)`

GetUsedSizeOk returns a tuple with the UsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSize

`func (o *GettableVolumeQuota) SetUsedSize(v int64)`

SetUsedSize sets UsedSize field to given value.

### HasUsedSize

`func (o *GettableVolumeQuota) HasUsedSize() bool`

HasUsedSize returns a boolean if a field has been set.

### GetCreated

`func (o *GettableVolumeQuota) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GettableVolumeQuota) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GettableVolumeQuota) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GettableVolumeQuota) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *GettableVolumeQuota) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GettableVolumeQuota) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GettableVolumeQuota) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GettableVolumeQuota) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


