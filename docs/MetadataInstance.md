# MetadataInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Meta** | Pointer to **int32** |  | [optional] [readonly] 
**Node** | Pointer to **int32** |  | [optional] 
**LogOnly** | Pointer to **bool** |  | [optional] [default to false]
**Config** | Pointer to [**MetadataInstanceConfig**](MetadataInstanceConfig.md) |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewMetadataInstance

`func NewMetadataInstance() *MetadataInstance`

NewMetadataInstance instantiates a new MetadataInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataInstanceWithDefaults

`func NewMetadataInstanceWithDefaults() *MetadataInstance`

NewMetadataInstanceWithDefaults instantiates a new MetadataInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *MetadataInstance) GetMeta() int32`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MetadataInstance) GetMetaOk() (*int32, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MetadataInstance) SetMeta(v int32)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MetadataInstance) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNode

`func (o *MetadataInstance) GetNode() int32`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *MetadataInstance) GetNodeOk() (*int32, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *MetadataInstance) SetNode(v int32)`

SetNode sets Node field to given value.

### HasNode

`func (o *MetadataInstance) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetLogOnly

`func (o *MetadataInstance) GetLogOnly() bool`

GetLogOnly returns the LogOnly field if non-nil, zero value otherwise.

### GetLogOnlyOk

`func (o *MetadataInstance) GetLogOnlyOk() (*bool, bool)`

GetLogOnlyOk returns a tuple with the LogOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOnly

`func (o *MetadataInstance) SetLogOnly(v bool)`

SetLogOnly sets LogOnly field to given value.

### HasLogOnly

`func (o *MetadataInstance) HasLogOnly() bool`

HasLogOnly returns a boolean if a field has been set.

### GetConfig

`func (o *MetadataInstance) GetConfig() MetadataInstanceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MetadataInstance) GetConfigOk() (*MetadataInstanceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MetadataInstance) SetConfig(v MetadataInstanceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MetadataInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetUpdated

`func (o *MetadataInstance) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *MetadataInstance) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *MetadataInstance) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *MetadataInstance) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


