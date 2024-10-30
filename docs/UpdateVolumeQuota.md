# UpdateVolumeQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The directory to be set quota on, it must be equal to the original path if it&#39;s provided | [optional] 
**Inodes** | Pointer to **int32** | Total inodes, &#x60;0&#x60; means no limit | [optional] 
**Size** | Pointer to **int64** | Total size in bytes, &#x60;0&#x60; means no limit | [optional] 

## Methods

### NewUpdateVolumeQuota

`func NewUpdateVolumeQuota() *UpdateVolumeQuota`

NewUpdateVolumeQuota instantiates a new UpdateVolumeQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeQuotaWithDefaults

`func NewUpdateVolumeQuotaWithDefaults() *UpdateVolumeQuota`

NewUpdateVolumeQuotaWithDefaults instantiates a new UpdateVolumeQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *UpdateVolumeQuota) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateVolumeQuota) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateVolumeQuota) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateVolumeQuota) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetInodes

`func (o *UpdateVolumeQuota) GetInodes() int32`

GetInodes returns the Inodes field if non-nil, zero value otherwise.

### GetInodesOk

`func (o *UpdateVolumeQuota) GetInodesOk() (*int32, bool)`

GetInodesOk returns a tuple with the Inodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInodes

`func (o *UpdateVolumeQuota) SetInodes(v int32)`

SetInodes sets Inodes field to given value.

### HasInodes

`func (o *UpdateVolumeQuota) HasInodes() bool`

HasInodes returns a boolean if a field has been set.

### GetSize

`func (o *UpdateVolumeQuota) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateVolumeQuota) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateVolumeQuota) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateVolumeQuota) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


