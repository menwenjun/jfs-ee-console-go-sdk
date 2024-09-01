# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Dns** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **int32** |  | [optional] 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Zoneid** | Pointer to **int32** |  | [optional] 
**SizeCap** | Pointer to **int32** |  | [optional] 
**ExpireAt** | Pointer to **NullableTime** |  | [optional] 
**InstanceSet** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Metadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Metadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Metadata) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Metadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDns

`func (o *Metadata) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *Metadata) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *Metadata) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *Metadata) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetPort

`func (o *Metadata) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Metadata) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Metadata) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Metadata) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRegion

`func (o *Metadata) GetRegion() int32`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Metadata) GetRegionOk() (*int32, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Metadata) SetRegion(v int32)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Metadata) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetParent

`func (o *Metadata) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Metadata) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Metadata) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Metadata) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *Metadata) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Metadata) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetZoneid

`func (o *Metadata) GetZoneid() int32`

GetZoneid returns the Zoneid field if non-nil, zero value otherwise.

### GetZoneidOk

`func (o *Metadata) GetZoneidOk() (*int32, bool)`

GetZoneidOk returns a tuple with the Zoneid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneid

`func (o *Metadata) SetZoneid(v int32)`

SetZoneid sets Zoneid field to given value.

### HasZoneid

`func (o *Metadata) HasZoneid() bool`

HasZoneid returns a boolean if a field has been set.

### GetSizeCap

`func (o *Metadata) GetSizeCap() int32`

GetSizeCap returns the SizeCap field if non-nil, zero value otherwise.

### GetSizeCapOk

`func (o *Metadata) GetSizeCapOk() (*int32, bool)`

GetSizeCapOk returns a tuple with the SizeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeCap

`func (o *Metadata) SetSizeCap(v int32)`

SetSizeCap sets SizeCap field to given value.

### HasSizeCap

`func (o *Metadata) HasSizeCap() bool`

HasSizeCap returns a boolean if a field has been set.

### GetExpireAt

`func (o *Metadata) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *Metadata) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *Metadata) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *Metadata) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### SetExpireAtNil

`func (o *Metadata) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *Metadata) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetInstanceSet

`func (o *Metadata) GetInstanceSet() []int32`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *Metadata) GetInstanceSetOk() (*[]int32, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *Metadata) SetInstanceSet(v []int32)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *Metadata) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


