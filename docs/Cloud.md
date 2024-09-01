# Cloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] 

## Methods

### NewCloud

`func NewCloud() *Cloud`

NewCloud instantiates a new Cloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWithDefaults

`func NewCloudWithDefaults() *Cloud`

NewCloudWithDefaults instantiates a new Cloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cloud) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cloud) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorage

`func (o *Cloud) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Cloud) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Cloud) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Cloud) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


