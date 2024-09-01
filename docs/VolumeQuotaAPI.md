# \VolumeQuotaAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumesVolumeIDQuotasGet**](VolumeQuotaAPI.md#VolumesVolumeIDQuotasGet) | **Get** /volumes/{volumeID}/quotas | Get quota list
[**VolumesVolumeIDQuotasPost**](VolumeQuotaAPI.md#VolumesVolumeIDQuotasPost) | **Post** /volumes/{volumeID}/quotas | Set quota
[**VolumesVolumeIDQuotasVolumeQuotaIDDelete**](VolumeQuotaAPI.md#VolumesVolumeIDQuotasVolumeQuotaIDDelete) | **Delete** /volumes/{volumeID}/quotas/{volumeQuotaID} | 
[**VolumesVolumeIDQuotasVolumeQuotaIDGet**](VolumeQuotaAPI.md#VolumesVolumeIDQuotasVolumeQuotaIDGet) | **Get** /volumes/{volumeID}/quotas/{volumeQuotaID} | 
[**VolumesVolumeIDQuotasVolumeQuotaIDPut**](VolumeQuotaAPI.md#VolumesVolumeIDQuotasVolumeQuotaIDPut) | **Put** /volumes/{volumeID}/quotas/{volumeQuotaID} | Update quota by ID



## VolumesVolumeIDQuotasGet

> []GettableVolumeQuota VolumesVolumeIDQuotasGet(ctx, volumeID).Token(token).Execute()

Get quota list



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
	resp, r, err := apiClient.VolumeQuotaAPI.VolumesVolumeIDQuotasGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeQuotaAPI.VolumesVolumeIDQuotasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDQuotasGet`: []GettableVolumeQuota
	fmt.Fprintf(os.Stdout, "Response from `VolumeQuotaAPI.VolumesVolumeIDQuotasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDQuotasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**[]GettableVolumeQuota**](GettableVolumeQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDQuotasPost

> []GettableVolumeQuota VolumesVolumeIDQuotasPost(ctx, volumeID).Token(token).CreateVolumeQuota(createVolumeQuota).Execute()

Set quota



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
	createVolumeQuota := *openapiclient.NewCreateVolumeQuota("/dir1/dir2") // CreateVolumeQuota |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeQuotaAPI.VolumesVolumeIDQuotasPost(context.Background(), volumeID).Token(token).CreateVolumeQuota(createVolumeQuota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeQuotaAPI.VolumesVolumeIDQuotasPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDQuotasPost`: []GettableVolumeQuota
	fmt.Fprintf(os.Stdout, "Response from `VolumeQuotaAPI.VolumesVolumeIDQuotasPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDQuotasPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **createVolumeQuota** | [**CreateVolumeQuota**](CreateVolumeQuota.md) |  | 

### Return type

[**[]GettableVolumeQuota**](GettableVolumeQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDQuotasVolumeQuotaIDDelete

> VolumesVolumeIDQuotasVolumeQuotaIDDelete(ctx, volumeID, volumeQuotaID).Token(token).Execute()





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
	volumeQuotaID := int32(56) // int32 | ID of the JuiceFS volume quota
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeQuotaAPI.VolumesVolumeIDQuotasVolumeQuotaIDDelete(context.Background(), volumeID, volumeQuotaID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeQuotaAPI.VolumesVolumeIDQuotasVolumeQuotaIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 
**volumeQuotaID** | **int32** | ID of the JuiceFS volume quota | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDQuotasVolumeQuotaIDDeleteRequest struct via the builder pattern


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


## VolumesVolumeIDQuotasVolumeQuotaIDGet

> GettableVolumeQuota VolumesVolumeIDQuotasVolumeQuotaIDGet(ctx, volumeID, volumeQuotaID).Token(token).Execute()





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
	volumeQuotaID := int32(56) // int32 | ID of the JuiceFS volume quota
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeQuotaAPI.VolumesVolumeIDQuotasVolumeQuotaIDGet(context.Background(), volumeID, volumeQuotaID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeQuotaAPI.VolumesVolumeIDQuotasVolumeQuotaIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDQuotasVolumeQuotaIDGet`: GettableVolumeQuota
	fmt.Fprintf(os.Stdout, "Response from `VolumeQuotaAPI.VolumesVolumeIDQuotasVolumeQuotaIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 
**volumeQuotaID** | **int32** | ID of the JuiceFS volume quota | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDQuotasVolumeQuotaIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**GettableVolumeQuota**](GettableVolumeQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDQuotasVolumeQuotaIDPut

> GettableVolumeQuota VolumesVolumeIDQuotasVolumeQuotaIDPut(ctx, volumeID, volumeQuotaID).Token(token).UpdateVolumeQuota(updateVolumeQuota).Execute()

Update quota by ID



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
	volumeQuotaID := int32(56) // int32 | ID of the JuiceFS volume quota
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	updateVolumeQuota := *openapiclient.NewUpdateVolumeQuota() // UpdateVolumeQuota |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeQuotaAPI.VolumesVolumeIDQuotasVolumeQuotaIDPut(context.Background(), volumeID, volumeQuotaID).Token(token).UpdateVolumeQuota(updateVolumeQuota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeQuotaAPI.VolumesVolumeIDQuotasVolumeQuotaIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDQuotasVolumeQuotaIDPut`: GettableVolumeQuota
	fmt.Fprintf(os.Stdout, "Response from `VolumeQuotaAPI.VolumesVolumeIDQuotasVolumeQuotaIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 
**volumeQuotaID** | **int32** | ID of the JuiceFS volume quota | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDQuotasVolumeQuotaIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **updateVolumeQuota** | [**UpdateVolumeQuota**](UpdateVolumeQuota.md) |  | 

### Return type

[**GettableVolumeQuota**](GettableVolumeQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

