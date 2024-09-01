# CreateVolumeQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The directory to be set quota on | 
**Inodes** | Pointer to **int32** | Total inodes, &#x60;0&#x60; or blank means no limit | [optional] 
**Size** | Pointer to **int32** | Total size in bytes, &#x60;0&#x60; or blank means no limit | [optional] 

## Methods

### NewCreateVolumeQuota

`func NewCreateVolumeQuota(path string, ) *CreateVolumeQuota`

NewCreateVolumeQuota instantiates a new CreateVolumeQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeQuotaWithDefaults

`func NewCreateVolumeQuotaWithDefaults() *CreateVolumeQuota`

NewCreateVolumeQuotaWithDefaults instantiates a new CreateVolumeQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CreateVolumeQuota) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateVolumeQuota) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateVolumeQuota) SetPath(v string)`

SetPath sets Path field to given value.


### GetInodes

`func (o *CreateVolumeQuota) GetInodes() int32`

GetInodes returns the Inodes field if non-nil, zero value otherwise.

### GetInodesOk

`func (o *CreateVolumeQuota) GetInodesOk() (*int32, bool)`

GetInodesOk returns a tuple with the Inodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInodes

`func (o *CreateVolumeQuota) SetInodes(v int32)`

SetInodes sets Inodes field to given value.

### HasInodes

`func (o *CreateVolumeQuota) HasInodes() bool`

HasInodes returns a boolean if a field has been set.

### GetSize

`func (o *CreateVolumeQuota) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVolumeQuota) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVolumeQuota) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateVolumeQuota) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


