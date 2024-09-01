# PostableVolumeExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** |  | [optional] 
**Iprange** | Pointer to **string** | allowed IP format:   - any IP address: *   - Single IP (dot-decimal): A.B.C.D   - IP range (dot-decimal): A.B.C.D-E.F.G.H   - IP range (CIDR): A.B.C.D/[0-32]  | [optional] [default to "*"]
**Readonly** | Pointer to **bool** |  | [optional] [default to false]
**Appendonly** | Pointer to **bool** |  | [optional] [default to false]
**Internalip** | Pointer to **bool** |  | [optional] [default to false]
**Qos** | Pointer to **string** | qos format: &#x60;&lt;put&gt;[:&lt;get&gt;[:&lt;compact&gt;]]&#x60;, put, get and compact has the same format &#x60;&lt;number&gt;[K|M|G|P|T|Z|E]&#x60;, 0 or empty value means unlimited.  the unit for qos is &#x60;B/s&#x60;.  examples: - &#x60;10M&#x60;, set put to 10MB/s, get and compact to unlimited - &#x60;10M:10M&#x60;, set put to 10MB/s, get to 10MB/s and compact to unlimited - &#x60;1k:1m:10m&#x60;, set put to 1KB/s, get to 1MB/s and compact to 10MB/s. the suffix(k, m) is case-insensitive. - &#x60;10M:0:10M&#x60;, set put to 10MB/s, get to unlimited and compact to 10MB/s  | [optional] [default to ""]
**Extend** | Pointer to **string** |  | [optional] [default to ""]
**Subdir** | Pointer to **string** |  | [optional] [default to "/"]

## Methods

### NewPostableVolumeExport

`func NewPostableVolumeExport() *PostableVolumeExport`

NewPostableVolumeExport instantiates a new PostableVolumeExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostableVolumeExportWithDefaults

`func NewPostableVolumeExportWithDefaults() *PostableVolumeExport`

NewPostableVolumeExportWithDefaults instantiates a new PostableVolumeExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *PostableVolumeExport) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *PostableVolumeExport) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *PostableVolumeExport) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *PostableVolumeExport) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetIprange

`func (o *PostableVolumeExport) GetIprange() string`

GetIprange returns the Iprange field if non-nil, zero value otherwise.

### GetIprangeOk

`func (o *PostableVolumeExport) GetIprangeOk() (*string, bool)`

GetIprangeOk returns a tuple with the Iprange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprange

`func (o *PostableVolumeExport) SetIprange(v string)`

SetIprange sets Iprange field to given value.

### HasIprange

`func (o *PostableVolumeExport) HasIprange() bool`

HasIprange returns a boolean if a field has been set.

### GetReadonly

`func (o *PostableVolumeExport) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *PostableVolumeExport) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *PostableVolumeExport) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *PostableVolumeExport) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetAppendonly

`func (o *PostableVolumeExport) GetAppendonly() bool`

GetAppendonly returns the Appendonly field if non-nil, zero value otherwise.

### GetAppendonlyOk

`func (o *PostableVolumeExport) GetAppendonlyOk() (*bool, bool)`

GetAppendonlyOk returns a tuple with the Appendonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendonly

`func (o *PostableVolumeExport) SetAppendonly(v bool)`

SetAppendonly sets Appendonly field to given value.

### HasAppendonly

`func (o *PostableVolumeExport) HasAppendonly() bool`

HasAppendonly returns a boolean if a field has been set.

### GetInternalip

`func (o *PostableVolumeExport) GetInternalip() bool`

GetInternalip returns the Internalip field if non-nil, zero value otherwise.

### GetInternalipOk

`func (o *PostableVolumeExport) GetInternalipOk() (*bool, bool)`

GetInternalipOk returns a tuple with the Internalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalip

`func (o *PostableVolumeExport) SetInternalip(v bool)`

SetInternalip sets Internalip field to given value.

### HasInternalip

`func (o *PostableVolumeExport) HasInternalip() bool`

HasInternalip returns a boolean if a field has been set.

### GetQos

`func (o *PostableVolumeExport) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *PostableVolumeExport) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *PostableVolumeExport) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *PostableVolumeExport) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetExtend

`func (o *PostableVolumeExport) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *PostableVolumeExport) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *PostableVolumeExport) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *PostableVolumeExport) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### GetSubdir

`func (o *PostableVolumeExport) GetSubdir() string`

GetSubdir returns the Subdir field if non-nil, zero value otherwise.

### GetSubdirOk

`func (o *PostableVolumeExport) GetSubdirOk() (*string, bool)`

GetSubdirOk returns a tuple with the Subdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdir

`func (o *PostableVolumeExport) SetSubdir(v string)`

SetSubdir sets Subdir field to given value.

### HasSubdir

`func (o *PostableVolumeExport) HasSubdir() bool`

HasSubdir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


