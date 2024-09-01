# \VolumeTrashAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumesVolumeIDDeleteTrashPost**](VolumeTrashAPI.md#VolumesVolumeIDDeleteTrashPost) | **Post** /volumes/{volumeID}/deleteTrash | Delete from trash
[**VolumesVolumeIDEmptyTrashPost**](VolumeTrashAPI.md#VolumesVolumeIDEmptyTrashPost) | **Post** /volumes/{volumeID}/emptyTrash | Empty trash asynchronously
[**VolumesVolumeIDListTrashGet**](VolumeTrashAPI.md#VolumesVolumeIDListTrashGet) | **Get** /volumes/{volumeID}/listTrash | List trash files
[**VolumesVolumeIDRestoreFilePost**](VolumeTrashAPI.md#VolumesVolumeIDRestoreFilePost) | **Post** /volumes/{volumeID}/restoreFile | Restore from trash



## VolumesVolumeIDDeleteTrashPost

> VolumesVolumeIDDeleteTrashPost(ctx, volumeID).Token(token).Inodes(inodes).Execute()

Delete from trash



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	inodes := *openapiclient.NewInodes() // Inodes | List of file inodes to be deleted in trash (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeTrashAPI.VolumesVolumeIDDeleteTrashPost(context.Background(), volumeID).Token(token).Inodes(inodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeTrashAPI.VolumesVolumeIDDeleteTrashPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDDeleteTrashPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **inodes** | [**Inodes**](Inodes.md) | List of file inodes to be deleted in trash | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDEmptyTrashPost

> VolumesVolumeIDEmptyTrashPost(ctx, volumeID).Token(token).Execute()

Empty trash asynchronously



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeTrashAPI.VolumesVolumeIDEmptyTrashPost(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeTrashAPI.VolumesVolumeIDEmptyTrashPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDEmptyTrashPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDListTrashGet

> VolumesVolumeIDListTrashGet200Response VolumesVolumeIDListTrashGet(ctx, volumeID).Token(token).Query(query).Page(page).N(n).Execute()

List trash files



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	query := "query_example" // string | keyword (optional)
	page := int32(56) // int32 | Page number, starting with `0` (optional)
	n := int32(56) // int32 | Page size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeTrashAPI.VolumesVolumeIDListTrashGet(context.Background(), volumeID).Token(token).Query(query).Page(page).N(n).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeTrashAPI.VolumesVolumeIDListTrashGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDListTrashGet`: VolumesVolumeIDListTrashGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VolumeTrashAPI.VolumesVolumeIDListTrashGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDListTrashGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **query** | **string** | keyword | 
 **page** | **int32** | Page number, starting with &#x60;0&#x60; | 
 **n** | **int32** | Page size | 

### Return type

[**VolumesVolumeIDListTrashGet200Response**](VolumesVolumeIDListTrashGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDRestoreFilePost

> VolumesVolumeIDRestoreFilePost(ctx, volumeID).Token(token).Inodes(inodes).Execute()

Restore from trash



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	inodes := *openapiclient.NewInodes() // Inodes | List of file inodes to be deleted in trash (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeTrashAPI.VolumesVolumeIDRestoreFilePost(context.Background(), volumeID).Token(token).Inodes(inodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeTrashAPI.VolumesVolumeIDRestoreFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDRestoreFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **inodes** | [**Inodes**](Inodes.md) | List of file inodes to be deleted in trash | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

