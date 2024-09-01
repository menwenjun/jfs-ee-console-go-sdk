# \VolumeAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumesGet**](VolumeAPI.md#VolumesGet) | **Get** /volumes | Get volume list
[**VolumesPost**](VolumeAPI.md#VolumesPost) | **Post** /volumes | 
[**VolumesVolumeIDClientsGet**](VolumeAPI.md#VolumesVolumeIDClientsGet) | **Get** /volumes/{volumeID}/clients | Get mounted clients
[**VolumesVolumeIDDelete**](VolumeAPI.md#VolumesVolumeIDDelete) | **Delete** /volumes/{volumeID} | 
[**VolumesVolumeIDGet**](VolumeAPI.md#VolumesVolumeIDGet) | **Get** /volumes/{volumeID} | 
[**VolumesVolumeIDIsReadyGet**](VolumeAPI.md#VolumesVolumeIDIsReadyGet) | **Get** /volumes/{volumeID}/is_ready | Volume has been ready or not
[**VolumesVolumeIDKerberosGet**](VolumeAPI.md#VolumesVolumeIDKerberosGet) | **Get** /volumes/{volumeID}/kerberos | 
[**VolumesVolumeIDKerberosPatch**](VolumeAPI.md#VolumesVolumeIDKerberosPatch) | **Patch** /volumes/{volumeID}/kerberos | 
[**VolumesVolumeIDListDirGet**](VolumeAPI.md#VolumesVolumeIDListDirGet) | **Get** /volumes/{volumeID}/listDir | List directory
[**VolumesVolumeIDMetricsGet**](VolumeAPI.md#VolumesVolumeIDMetricsGet) | **Get** /volumes/{volumeID}/metrics | Get volume metrics
[**VolumesVolumeIDMirrorsGet**](VolumeAPI.md#VolumesVolumeIDMirrorsGet) | **Get** /volumes/{volumeID}/mirrors | 
[**VolumesVolumeIDMirrorsPost**](VolumeAPI.md#VolumesVolumeIDMirrorsPost) | **Post** /volumes/{volumeID}/mirrors | 
[**VolumesVolumeIDRangerGet**](VolumeAPI.md#VolumesVolumeIDRangerGet) | **Get** /volumes/{volumeID}/ranger | 
[**VolumesVolumeIDRangerPatch**](VolumeAPI.md#VolumesVolumeIDRangerPatch) | **Patch** /volumes/{volumeID}/ranger | 
[**VolumesVolumeIDSettingGet**](VolumeAPI.md#VolumesVolumeIDSettingGet) | **Get** /volumes/{volumeID}/setting | 
[**VolumesVolumeIDSettingPost**](VolumeAPI.md#VolumesVolumeIDSettingPost) | **Post** /volumes/{volumeID}/setting | 



## VolumesGet

> []GettableVolume VolumesGet(ctx).Token(token).Execute()

Get volume list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesGet(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesGet`: []GettableVolume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVolumesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**[]GettableVolume**](GettableVolume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesPost

> GettableVolume VolumesPost(ctx).Token(token).PostableVolume(postableVolume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	postableVolume := *openapiclient.NewPostableVolume("demo-vol1", int32(1)) // PostableVolume |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesPost(context.Background()).Token(token).PostableVolume(postableVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesPost`: GettableVolume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVolumesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **postableVolume** | [**PostableVolume**](PostableVolume.md) |  | 

### Return type

[**GettableVolume**](GettableVolume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDClientsGet

> []VolumeClient VolumesVolumeIDClientsGet(ctx, volumeID).Token(token).Execute()

Get mounted clients



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDClientsGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDClientsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDClientsGet`: []VolumeClient
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDClientsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDClientsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**[]VolumeClient**](VolumeClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDDelete

> VolumesVolumeIDDelete(ctx, volumeID).Token(token).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeAPI.VolumesVolumeIDDelete(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVolumesVolumeIDDeleteRequest struct via the builder pattern


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


## VolumesVolumeIDGet

> GettableVolume VolumesVolumeIDGet(ctx, volumeID).Token(token).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDGet`: GettableVolume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**GettableVolume**](GettableVolume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDIsReadyGet

> VolumesVolumeIDIsReadyGet200Response VolumesVolumeIDIsReadyGet(ctx, volumeID).Token(token).Execute()

Volume has been ready or not

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDIsReadyGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDIsReadyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDIsReadyGet`: VolumesVolumeIDIsReadyGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDIsReadyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDIsReadyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**VolumesVolumeIDIsReadyGet200Response**](VolumesVolumeIDIsReadyGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDKerberosGet

> VolumeKerberos VolumesVolumeIDKerberosGet(ctx, volumeID).Token(token).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDKerberosGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDKerberosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDKerberosGet`: VolumeKerberos
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDKerberosGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDKerberosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**VolumeKerberos**](VolumeKerberos.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDKerberosPatch

> VolumeKerberos VolumesVolumeIDKerberosPatch(ctx, volumeID).Token(token).VolumeKerberos(volumeKerberos).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	volumeKerberos := *openapiclient.NewVolumeKerberos(false) // VolumeKerberos |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDKerberosPatch(context.Background(), volumeID).Token(token).VolumeKerberos(volumeKerberos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDKerberosPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDKerberosPatch`: VolumeKerberos
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDKerberosPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDKerberosPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **volumeKerberos** | [**VolumeKerberos**](VolumeKerberos.md) |  | 

### Return type

[**VolumeKerberos**](VolumeKerberos.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDListDirGet

> []FileInfo VolumesVolumeIDListDirGet(ctx, volumeID).Token(token).Path(path).Query(query).Page(page).N(n).ChildrenLimit(childrenLimit).Execute()

List directory



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	path := "path_example" // string | The directory to be listed (optional)
	query := "query_example" // string | A keyword used to search files(case insensitive) (optional)
	page := int32(56) // int32 | Page number, starting with `0` (optional)
	n := int32(56) // int32 | Page size (optional)
	childrenLimit := int32(56) // int32 | Limit the number of children to be listed. The maximum value is 1000000 for On-premises, 100000 for SaaS.  You should be careful when setting this value to a large number, listDir may become slow or even timeout.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDListDirGet(context.Background(), volumeID).Token(token).Path(path).Query(query).Page(page).N(n).ChildrenLimit(childrenLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDListDirGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDListDirGet`: []FileInfo
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDListDirGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDListDirGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **path** | **string** | The directory to be listed | 
 **query** | **string** | A keyword used to search files(case insensitive) | 
 **page** | **int32** | Page number, starting with &#x60;0&#x60; | 
 **n** | **int32** | Page size | 
 **childrenLimit** | **int32** | Limit the number of children to be listed. The maximum value is 1000000 for On-premises, 100000 for SaaS.  You should be careful when setting this value to a large number, listDir may become slow or even timeout.  | 

### Return type

[**[]FileInfo**](FileInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDMetricsGet

> string VolumesVolumeIDMetricsGet(ctx, volumeID).Token(token).Execute()

Get volume metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDMetricsGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDMetricsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDMetricsGet`: string
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDMetricsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDMetricsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDMirrorsGet

> MirrorVolume VolumesVolumeIDMirrorsGet(ctx, volumeID).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDMirrorsGet(context.Background(), volumeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDMirrorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDMirrorsGet`: MirrorVolume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDMirrorsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDMirrorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MirrorVolume**](MirrorVolume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDMirrorsPost

> CreateMirrorVolume VolumesVolumeIDMirrorsPost(ctx, volumeID).CreateMirrorVolume(createMirrorVolume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	createMirrorVolume := *openapiclient.NewCreateMirrorVolume("demo-vol1", int32(1)) // CreateMirrorVolume |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDMirrorsPost(context.Background(), volumeID).CreateMirrorVolume(createMirrorVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDMirrorsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDMirrorsPost`: CreateMirrorVolume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDMirrorsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDMirrorsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMirrorVolume** | [**CreateMirrorVolume**](CreateMirrorVolume.md) |  | 

### Return type

[**CreateMirrorVolume**](CreateMirrorVolume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDRangerGet

> VolumeRanger VolumesVolumeIDRangerGet(ctx, volumeID).Token(token).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDRangerGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDRangerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDRangerGet`: VolumeRanger
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDRangerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDRangerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**VolumeRanger**](VolumeRanger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDRangerPatch

> VolumeRanger VolumesVolumeIDRangerPatch(ctx, volumeID).Token(token).VolumeRanger(volumeRanger).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	volumeRanger := *openapiclient.NewVolumeRanger(false, "AdminUrl_example", "ServiceName_example") // VolumeRanger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDRangerPatch(context.Background(), volumeID).Token(token).VolumeRanger(volumeRanger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDRangerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDRangerPatch`: VolumeRanger
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDRangerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDRangerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **volumeRanger** | [**VolumeRanger**](VolumeRanger.md) |  | 

### Return type

[**VolumeRanger**](VolumeRanger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDSettingGet

> TrashTime VolumesVolumeIDSettingGet(ctx, volumeID).Token(token).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDSettingGet(context.Background(), volumeID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDSettingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDSettingGet`: TrashTime
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDSettingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDSettingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**TrashTime**](TrashTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesVolumeIDSettingPost

> TrashTime VolumesVolumeIDSettingPost(ctx, volumeID).Token(token).TrashTime(trashTime).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
)

func main() {
	volumeID := int32(56) // int32 | ID of the JuiceFS volume
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	trashTime := *openapiclient.NewTrashTime() // TrashTime |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.VolumesVolumeIDSettingPost(context.Background(), volumeID).Token(token).TrashTime(trashTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.VolumesVolumeIDSettingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesVolumeIDSettingPost`: TrashTime
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.VolumesVolumeIDSettingPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeID** | **int32** | ID of the JuiceFS volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesVolumeIDSettingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **trashTime** | [**TrashTime**](TrashTime.md) |  | 

### Return type

[**TrashTime**](TrashTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

