# \VolumeACLAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumesVolumeIDExportsGet**](VolumeACLAPI.md#VolumesVolumeIDExportsGet) | **Get** /volumes/{volumeID}/exports | Get ACL list
[**VolumesVolumeIDExportsPost**](VolumeACLAPI.md#VolumesVolumeIDExportsPost) | **Post** /volumes/{volumeID}/exports | Create ACL
[**VolumesVolumeIDExportsVolumeExportIDDelete**](VolumeACLAPI.md#VolumesVolumeIDExportsVolumeExportIDDelete) | **Delete** /volumes/{volumeID}/exports/{volumeExportID} | Delete ACL
[**VolumesVolumeIDExportsVolumeExportIDGet**](VolumeACLAPI.md#VolumesVolumeIDExportsVolumeExportIDGet) | **Get** /volumes/{volumeID}/exports/{volumeExportID} | Get ACL by ID
[**VolumesVolumeIDExportsVolumeExportIDPut**](VolumeACLAPI.md#VolumesVolumeIDExportsVolumeExportIDPut) | **Put** /volumes/{volumeID}/exports/{volumeExportID} | Update ACL by ID



## VolumesVolumeIDExportsGet

> []GettableVolumeExport VolumesVolumeIDExportsGet(ctx, volumeID).Token(token).Execute()

Get ACL list



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
	resp, r, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeACLAPI.VolumesVolumeIDExportsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDExportsGet`: []GettableVolumeExport
	fmt.Fprintf(os.Stdout, "Response from `VolumeACLAPI.VolumesVolumeIDExportsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDExportsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**[]GettableVolumeExport**](GettableVolumeExport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDExportsPost

> GettableVolumeExport VolumesVolumeIDExportsPost(ctx, volumeID).Token(token).PostableVolumeExport(postableVolumeExport).Execute()

Create ACL



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
	postableVolumeExport := *openapiclient.NewPostableVolumeExport() // PostableVolumeExport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsPost(context.Background(), volumeID).Token(token).PostableVolumeExport(postableVolumeExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeACLAPI.VolumesVolumeIDExportsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDExportsPost`: GettableVolumeExport
	fmt.Fprintf(os.Stdout, "Response from `VolumeACLAPI.VolumesVolumeIDExportsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDExportsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **postableVolumeExport** | [**PostableVolumeExport**](PostableVolumeExport.md) |  | 

### Return type

[**GettableVolumeExport**](GettableVolumeExport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDExportsVolumeExportIDDelete

> VolumesVolumeIDExportsVolumeExportIDDelete(ctx, volumeID, volumeExportID).Token(token).Execute()

Delete ACL



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
	volumeExportID := int32(56) // int32 | ID of the JuiceFS volume export
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDDelete(context.Background(), volumeID, volumeExportID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 
**volumeExportID** | **int32** | ID of the JuiceFS volume export | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDExportsVolumeExportIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDExportsVolumeExportIDGet

> GettableVolumeExport VolumesVolumeIDExportsVolumeExportIDGet(ctx, volumeID, volumeExportID).Token(token).Execute()

Get ACL by ID



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
	volumeExportID := int32(56) // int32 | ID of the JuiceFS volume export
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDGet(context.Background(), volumeID, volumeExportID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDExportsVolumeExportIDGet`: GettableVolumeExport
	fmt.Fprintf(os.Stdout, "Response from `VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 
**volumeExportID** | **int32** | ID of the JuiceFS volume export | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDExportsVolumeExportIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**GettableVolumeExport**](GettableVolumeExport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDExportsVolumeExportIDPut

> GettableVolumeExport VolumesVolumeIDExportsVolumeExportIDPut(ctx, volumeID, volumeExportID).Token(token).PostableVolumeExport(postableVolumeExport).Execute()

Update ACL by ID



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
	volumeExportID := int32(56) // int32 | ID of the JuiceFS volume export
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	postableVolumeExport := *openapiclient.NewPostableVolumeExport() // PostableVolumeExport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDPut(context.Background(), volumeID, volumeExportID).Token(token).PostableVolumeExport(postableVolumeExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDExportsVolumeExportIDPut`: GettableVolumeExport
	fmt.Fprintf(os.Stdout, "Response from `VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 
**volumeExportID** | **int32** | ID of the JuiceFS volume export | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDExportsVolumeExportIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **postableVolumeExport** | [**PostableVolumeExport**](PostableVolumeExport.md) |  | 

### Return type

[**GettableVolumeExport**](GettableVolumeExport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

