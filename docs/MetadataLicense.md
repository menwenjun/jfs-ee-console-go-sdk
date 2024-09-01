# MetadataLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **string** |  | [optional] 
**ExpiredAt** | Pointer to **time.Time** |  | [optional] 
**SizeCap** | Pointer to **int32** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadataLicense

`func NewMetadataLicense() *MetadataLicense`

NewMetadataLicense instantiates a new MetadataLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataLicenseWithDefaults

`func NewMetadataLicenseWithDefaults() *MetadataLicense`

NewMetadataLicenseWithDefaults instantiates a new MetadataLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *MetadataLicense) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *MetadataLicense) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *MetadataLicense) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *MetadataLicense) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetExpiredAt

`func (o *MetadataLicense) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *MetadataLicense) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *MetadataLicense) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *MetadataLicense) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetSizeCap

`func (o *MetadataLicense) GetSizeCap() int32`

GetSizeCap returns the SizeCap field if non-nil, zero value otherwise.

### GetSizeCapOk

`func (o *MetadataLicense) GetSizeCapOk() (*int32, bool)`

GetSizeCapOk returns a tuple with the SizeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeCap

`func (o *MetadataLicense) SetSizeCap(v int32)`

SetSizeCap sets SizeCap field to given value.

### HasSizeCap

`func (o *MetadataLicense) HasSizeCap() bool`

HasSizeCap returns a boolean if a field has been set.

### GetLicense

`func (o *MetadataLicense) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *MetadataLicense) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *MetadataLicense) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *MetadataLicense) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


