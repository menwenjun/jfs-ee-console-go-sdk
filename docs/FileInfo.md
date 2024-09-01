# FileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**IsDir** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | - f: file - l: symlink - d: directory - q: fifo - b: block device - c: character device - s: socket - t: trash file  | [optional] 
**Length** | Pointer to **int64** |  | [optional] 
**Nlink** | Pointer to **int32** |  | [optional] 
**Mtime** | Pointer to **int32** |  | [optional] 

## Methods

### NewFileInfo

`func NewFileInfo() *FileInfo`

NewFileInfo instantiates a new FileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoWithDefaults

`func NewFileInfoWithDefaults() *FileInfo`

NewFileInfoWithDefaults instantiates a new FileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDir

`func (o *FileInfo) GetIsDir() bool`

GetIsDir returns the IsDir field if non-nil, zero value otherwise.

### GetIsDirOk

`func (o *FileInfo) GetIsDirOk() (*bool, bool)`

GetIsDirOk returns a tuple with the IsDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDir

`func (o *FileInfo) SetIsDir(v bool)`

SetIsDir sets IsDir field to given value.

### HasIsDir

`func (o *FileInfo) HasIsDir() bool`

HasIsDir returns a boolean if a field has been set.

### GetType

`func (o *FileInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLength

`func (o *FileInfo) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *FileInfo) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *FileInfo) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *FileInfo) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetNlink

`func (o *FileInfo) GetNlink() int32`

GetNlink returns the Nlink field if non-nil, zero value otherwise.

### GetNlinkOk

`func (o *FileInfo) GetNlinkOk() (*int32, bool)`

GetNlinkOk returns a tuple with the Nlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNlink

`func (o *FileInfo) SetNlink(v int32)`

SetNlink sets Nlink field to given value.

### HasNlink

`func (o *FileInfo) HasNlink() bool`

HasNlink returns a boolean if a field has been set.

### GetMtime

`func (o *FileInfo) GetMtime() int32`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *FileInfo) GetMtimeOk() (*int32, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *FileInfo) SetMtime(v int32)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *FileInfo) HasMtime() bool`

HasMtime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


