# VolumeKerberosProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** |  | 
**ProxyUsers** | Pointer to **string** |  | [optional] [default to "*"]
**ProxyGroups** | Pointer to **string** |  | [optional] [default to "*"]
**ProxyHosts** | Pointer to **string** |  | [optional] [default to "*"]

## Methods

### NewVolumeKerberosProxy

`func NewVolumeKerberosProxy(user string, ) *VolumeKerberosProxy`

NewVolumeKerberosProxy instantiates a new VolumeKerberosProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeKerberosProxyWithDefaults

`func NewVolumeKerberosProxyWithDefaults() *VolumeKerberosProxy`

NewVolumeKerberosProxyWithDefaults instantiates a new VolumeKerberosProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *VolumeKerberosProxy) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VolumeKerberosProxy) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VolumeKerberosProxy) SetUser(v string)`

SetUser sets User field to given value.


### GetProxyUsers

`func (o *VolumeKerberosProxy) GetProxyUsers() string`

GetProxyUsers returns the ProxyUsers field if non-nil, zero value otherwise.

### GetProxyUsersOk

`func (o *VolumeKerberosProxy) GetProxyUsersOk() (*string, bool)`

GetProxyUsersOk returns a tuple with the ProxyUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUsers

`func (o *VolumeKerberosProxy) SetProxyUsers(v string)`

SetProxyUsers sets ProxyUsers field to given value.

### HasProxyUsers

`func (o *VolumeKerberosProxy) HasProxyUsers() bool`

HasProxyUsers returns a boolean if a field has been set.

### GetProxyGroups

`func (o *VolumeKerberosProxy) GetProxyGroups() string`

GetProxyGroups returns the ProxyGroups field if non-nil, zero value otherwise.

### GetProxyGroupsOk

`func (o *VolumeKerberosProxy) GetProxyGroupsOk() (*string, bool)`

GetProxyGroupsOk returns a tuple with the ProxyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyGroups

`func (o *VolumeKerberosProxy) SetProxyGroups(v string)`

SetProxyGroups sets ProxyGroups field to given value.

### HasProxyGroups

`func (o *VolumeKerberosProxy) HasProxyGroups() bool`

HasProxyGroups returns a boolean if a field has been set.

### GetProxyHosts

`func (o *VolumeKerberosProxy) GetProxyHosts() string`

GetProxyHosts returns the ProxyHosts field if non-nil, zero value otherwise.

### GetProxyHostsOk

`func (o *VolumeKerberosProxy) GetProxyHostsOk() (*string, bool)`

GetProxyHostsOk returns a tuple with the ProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHosts

`func (o *VolumeKerberosProxy) SetProxyHosts(v string)`

SetProxyHosts sets ProxyHosts field to given value.

### HasProxyHosts

`func (o *VolumeKerberosProxy) HasProxyHosts() bool`

HasProxyHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


