# \RegionAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegionsGet**](RegionAPI.md#RegionsGet) | **Get** /regions | Get region list
[**RegionsRegionIDGet**](RegionAPI.md#RegionsRegionIDGet) | **Get** /regions/{regionID} | Get region by ID
[**RegionsRegionIDMetricsGet**](RegionAPI.md#RegionsRegionIDMetricsGet) | **Get** /regions/{regionID}/metrics | Get region metrics
[**RegionsRegionIDTrashtimeGet**](RegionAPI.md#RegionsRegionIDTrashtimeGet) | **Get** /regions/{regionID}/trashtime | 
[**RegionsRegionIDTrashtimePut**](RegionAPI.md#RegionsRegionIDTrashtimePut) | **Put** /regions/{regionID}/trashtime | Set region trashtime(days)



## RegionsGet

> []Region RegionsGet(ctx).Execute()

Get region list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.RegionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.RegionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsGet`: []Region
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.RegionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsGetRequest struct via the builder pattern


### Return type

[**[]Region**](Region.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionsRegionIDGet

> Region RegionsRegionIDGet(ctx, regionID).Token(token).Execute()

Get region by ID



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
	regionID := int32(56) // int32 | ID of the region
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.RegionsRegionIDGet(context.Background(), regionID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.RegionsRegionIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsRegionIDGet`: Region
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.RegionsRegionIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionID** | **int32** | ID of the region | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsRegionIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**Region**](Region.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionsRegionIDMetricsGet

> string RegionsRegionIDMetricsGet(ctx, regionID).Token(token).Execute()

Get region metrics



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
	regionID := int32(56) // int32 | ID of the region
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.RegionsRegionIDMetricsGet(context.Background(), regionID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.RegionsRegionIDMetricsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsRegionIDMetricsGet`: string
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.RegionsRegionIDMetricsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionID** | **int32** | ID of the region | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsRegionIDMetricsGetRequest struct via the builder pattern


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


## RegionsRegionIDTrashtimeGet

> TrashTime RegionsRegionIDTrashtimeGet(ctx, regionID).Token(token).Execute()





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
	regionID := int32(56) // int32 | ID of the region
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.RegionsRegionIDTrashtimeGet(context.Background(), regionID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.RegionsRegionIDTrashtimeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsRegionIDTrashtimeGet`: TrashTime
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.RegionsRegionIDTrashtimeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionID** | **int32** | ID of the region | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsRegionIDTrashtimeGetRequest struct via the builder pattern


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


## RegionsRegionIDTrashtimePut

> TrashTime RegionsRegionIDTrashtimePut(ctx, regionID).Token(token).TrashTime(trashTime).Execute()

Set region trashtime(days)



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
	regionID := int32(56) // int32 | ID of the region
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	trashTime := *openapiclient.NewTrashTime() // TrashTime |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.RegionsRegionIDTrashtimePut(context.Background(), regionID).Token(token).TrashTime(trashTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.RegionsRegionIDTrashtimePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsRegionIDTrashtimePut`: TrashTime
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.RegionsRegionIDTrashtimePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionID** | **int32** | ID of the region | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsRegionIDTrashtimePutRequest struct via the builder pattern


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

