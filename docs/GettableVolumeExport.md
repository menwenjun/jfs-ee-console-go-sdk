# GettableVolumeExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Desc** | Pointer to **string** |  | [optional] 
**Iprange** | Pointer to **string** |  | [optional] 
**Apionly** | Pointer to **bool** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**Appendonly** | Pointer to **bool** |  | [optional] 
**Internalip** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Passwd** | Pointer to **string** |  | [optional] 
**Qos** | Pointer to **string** |  | [optional] 
**Extend** | Pointer to **string** |  | [optional] 
**Subdir** | Pointer to **string** |  | [optional] 

## Methods

### NewGettableVolumeExport

`func NewGettableVolumeExport() *GettableVolumeExport`

NewGettableVolumeExport instantiates a new GettableVolumeExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGettableVolumeExportWithDefaults

`func NewGettableVolumeExportWithDefaults() *GettableVolumeExport`

NewGettableVolumeExportWithDefaults instantiates a new GettableVolumeExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GettableVolumeExport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GettableVolumeExport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GettableVolumeExport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GettableVolumeExport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDesc

`func (o *GettableVolumeExport) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *GettableVolumeExport) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *GettableVolumeExport) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *GettableVolumeExport) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetIprange

`func (o *GettableVolumeExport) GetIprange() string`

GetIprange returns the Iprange field if non-nil, zero value otherwise.

### GetIprangeOk

`func (o *GettableVolumeExport) GetIprangeOk() (*string, bool)`

GetIprangeOk returns a tuple with the Iprange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprange

`func (o *GettableVolumeExport) SetIprange(v string)`

SetIprange sets Iprange field to given value.

### HasIprange

`func (o *GettableVolumeExport) HasIprange() bool`

HasIprange returns a boolean if a field has been set.

### GetApionly

`func (o *GettableVolumeExport) GetApionly() bool`

GetApionly returns the Apionly field if non-nil, zero value otherwise.

### GetApionlyOk

`func (o *GettableVolumeExport) GetApionlyOk() (*bool, bool)`

GetApionlyOk returns a tuple with the Apionly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApionly

`func (o *GettableVolumeExport) SetApionly(v bool)`

SetApionly sets Apionly field to given value.

### HasApionly

`func (o *GettableVolumeExport) HasApionly() bool`

HasApionly returns a boolean if a field has been set.

### GetReadonly

`func (o *GettableVolumeExport) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *GettableVolumeExport) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *GettableVolumeExport) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *GettableVolumeExport) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetAppendonly

`func (o *GettableVolumeExport) GetAppendonly() bool`

GetAppendonly returns the Appendonly field if non-nil, zero value otherwise.

### GetAppendonlyOk

`func (o *GettableVolumeExport) GetAppendonlyOk() (*bool, bool)`

GetAppendonlyOk returns a tuple with the Appendonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendonly

`func (o *GettableVolumeExport) SetAppendonly(v bool)`

SetAppendonly sets Appendonly field to given value.

### HasAppendonly

`func (o *GettableVolumeExport) HasAppendonly() bool`

HasAppendonly returns a boolean if a field has been set.

### GetInternalip

`func (o *GettableVolumeExport) GetInternalip() bool`

GetInternalip returns the Internalip field if non-nil, zero value otherwise.

### GetInternalipOk

`func (o *GettableVolumeExport) GetInternalipOk() (*bool, bool)`

GetInternalipOk returns a tuple with the Internalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalip

`func (o *GettableVolumeExport) SetInternalip(v bool)`

SetInternalip sets Internalip field to given value.

### HasInternalip

`func (o *GettableVolumeExport) HasInternalip() bool`

HasInternalip returns a boolean if a field has been set.

### GetToken

`func (o *GettableVolumeExport) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GettableVolumeExport) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GettableVolumeExport) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GettableVolumeExport) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetPasswd

`func (o *GettableVolumeExport) GetPasswd() string`

GetPasswd returns the Passwd field if non-nil, zero value otherwise.

### GetPasswdOk

`func (o *GettableVolumeExport) GetPasswdOk() (*string, bool)`

GetPasswdOk returns a tuple with the Passwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswd

`func (o *GettableVolumeExport) SetPasswd(v string)`

SetPasswd sets Passwd field to given value.

### HasPasswd

`func (o *GettableVolumeExport) HasPasswd() bool`

HasPasswd returns a boolean if a field has been set.

### GetQos

`func (o *GettableVolumeExport) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *GettableVolumeExport) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *GettableVolumeExport) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *GettableVolumeExport) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetExtend

`func (o *GettableVolumeExport) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *GettableVolumeExport) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *GettableVolumeExport) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *GettableVolumeExport) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### GetSubdir

`func (o *GettableVolumeExport) GetSubdir() string`

GetSubdir returns the Subdir field if non-nil, zero value otherwise.

### GetSubdirOk

`func (o *GettableVolumeExport) GetSubdirOk() (*string, bool)`

GetSubdirOk returns a tuple with the Subdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdir

`func (o *GettableVolumeExport) SetSubdir(v string)`

SetSubdir sets Subdir field to given value.

### HasSubdir

`func (o *GettableVolumeExport) HasSubdir() bool`

HasSubdir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


