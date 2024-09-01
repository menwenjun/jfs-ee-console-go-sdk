# VolumeKerberos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** |  | 
**Keytab** | Pointer to **string** |  | [optional] [default to ""]
**Superuser** | Pointer to **string** |  | [optional] [default to ""]
**Supergroup** | Pointer to **string** |  | [optional] [default to ""]
**Proxies** | Pointer to [**[]VolumeKerberosProxy**](VolumeKerberosProxy.md) |  | [optional] 

## Methods

### NewVolumeKerberos

`func NewVolumeKerberos(enable bool, ) *VolumeKerberos`

NewVolumeKerberos instantiates a new VolumeKerberos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeKerberosWithDefaults

`func NewVolumeKerberosWithDefaults() *VolumeKerberos`

NewVolumeKerberosWithDefaults instantiates a new VolumeKerberos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *VolumeKerberos) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *VolumeKerberos) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *VolumeKerberos) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetKeytab

`func (o *VolumeKerberos) GetKeytab() string`

GetKeytab returns the Keytab field if non-nil, zero value otherwise.

### GetKeytabOk

`func (o *VolumeKerberos) GetKeytabOk() (*string, bool)`

GetKeytabOk returns a tuple with the Keytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytab

`func (o *VolumeKerberos) SetKeytab(v string)`

SetKeytab sets Keytab field to given value.

### HasKeytab

`func (o *VolumeKerberos) HasKeytab() bool`

HasKeytab returns a boolean if a field has been set.

### GetSuperuser

`func (o *VolumeKerberos) GetSuperuser() string`

GetSuperuser returns the Superuser field if non-nil, zero value otherwise.

### GetSuperuserOk

`func (o *VolumeKerberos) GetSuperuserOk() (*string, bool)`

GetSuperuserOk returns a tuple with the Superuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperuser

`func (o *VolumeKerberos) SetSuperuser(v string)`

SetSuperuser sets Superuser field to given value.

### HasSuperuser

`func (o *VolumeKerberos) HasSuperuser() bool`

HasSuperuser returns a boolean if a field has been set.

### GetSupergroup

`func (o *VolumeKerberos) GetSupergroup() string`

GetSupergroup returns the Supergroup field if non-nil, zero value otherwise.

### GetSupergroupOk

`func (o *VolumeKerberos) GetSupergroupOk() (*string, bool)`

GetSupergroupOk returns a tuple with the Supergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupergroup

`func (o *VolumeKerberos) SetSupergroup(v string)`

SetSupergroup sets Supergroup field to given value.

### HasSupergroup

`func (o *VolumeKerberos) HasSupergroup() bool`

HasSupergroup returns a boolean if a field has been set.

### GetProxies

`func (o *VolumeKerberos) GetProxies() []VolumeKerberosProxy`

GetProxies returns the Proxies field if non-nil, zero value otherwise.

### GetProxiesOk

`func (o *VolumeKerberos) GetProxiesOk() (*[]VolumeKerberosProxy, bool)`

GetProxiesOk returns a tuple with the Proxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxies

`func (o *VolumeKerberos) SetProxies(v []VolumeKerberosProxy)`

SetProxies sets Proxies field to given value.

### HasProxies

`func (o *VolumeKerberos) HasProxies() bool`

HasProxies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


