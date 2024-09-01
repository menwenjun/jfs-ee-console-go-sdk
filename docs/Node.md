# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Ip** | Pointer to **string** |  | [optional] [readonly] 
**Disks** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**Memory** | Pointer to **float64** |  | [optional] [readonly] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Node) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Node) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *Node) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Node) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Node) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Node) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetDisks

`func (o *Node) GetDisks() string`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Node) GetDisksOk() (*string, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Node) SetDisks(v string)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Node) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetActive

`func (o *Node) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Node) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Node) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Node) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMemory

`func (o *Node) GetMemory() float64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Node) GetMemoryOk() (*float64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Node) SetMemory(v float64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Node) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


