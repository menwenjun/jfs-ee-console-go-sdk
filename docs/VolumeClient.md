# VolumeClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Mountpoint** | Pointer to **string** |  | [optional] 
**Openfiles** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to **map[string]interface{}** |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVolumeClient

`func NewVolumeClient() *VolumeClient`

NewVolumeClient instantiates a new VolumeClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeClientWithDefaults

`func NewVolumeClientWithDefaults() *VolumeClient`

NewVolumeClientWithDefaults instantiates a new VolumeClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VolumeClient) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VolumeClient) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VolumeClient) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VolumeClient) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *VolumeClient) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VolumeClient) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VolumeClient) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *VolumeClient) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMountpoint

`func (o *VolumeClient) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *VolumeClient) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *VolumeClient) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *VolumeClient) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetOpenfiles

`func (o *VolumeClient) GetOpenfiles() int32`

GetOpenfiles returns the Openfiles field if non-nil, zero value otherwise.

### GetOpenfilesOk

`func (o *VolumeClient) GetOpenfilesOk() (*int32, bool)`

GetOpenfilesOk returns a tuple with the Openfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenfiles

`func (o *VolumeClient) SetOpenfiles(v int32)`

SetOpenfiles sets Openfiles field to given value.

### HasOpenfiles

`func (o *VolumeClient) HasOpenfiles() bool`

HasOpenfiles returns a boolean if a field has been set.

### GetVersion

`func (o *VolumeClient) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VolumeClient) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VolumeClient) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VolumeClient) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOperations

`func (o *VolumeClient) GetOperations() map[string]interface{}`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *VolumeClient) GetOperationsOk() (*map[string]interface{}, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *VolumeClient) SetOperations(v map[string]interface{})`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *VolumeClient) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetStats

`func (o *VolumeClient) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *VolumeClient) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *VolumeClient) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *VolumeClient) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


