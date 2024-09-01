# \InstanceAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadatasMetadataIDInstancesGet**](InstanceAPI.md#MetadatasMetadataIDInstancesGet) | **Get** /metadatas/{metadataID}/instances | Get instance list
[**MetadatasMetadataIDInstancesInstanceIDGet**](InstanceAPI.md#MetadatasMetadataIDInstancesInstanceIDGet) | **Get** /metadatas/{metadataID}/instances/{instanceID} | Get instance by ID
[**MetadatasMetadataIDInstancesInstanceIDPatch**](InstanceAPI.md#MetadatasMetadataIDInstancesInstanceIDPatch) | **Patch** /metadatas/{metadataID}/instances/{instanceID} | Update instance properties



## MetadatasMetadataIDInstancesGet

> []MetadataInstance MetadatasMetadataIDInstancesGet(ctx, metadataID).Token(token).Execute()

Get instance list



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
	metadataID := int32(56) // int32 | ID of the metadata
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.MetadatasMetadataIDInstancesGet(context.Background(), metadataID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.MetadatasMetadataIDInstancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadatasMetadataIDInstancesGet`: []MetadataInstance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.MetadatasMetadataIDInstancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataID** | **int32** | ID of the metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadatasMetadataIDInstancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**[]MetadataInstance**](MetadataInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadatasMetadataIDInstancesInstanceIDGet

> MetadataInstance MetadatasMetadataIDInstancesInstanceIDGet(ctx, metadataID, instanceID).Token(token).Execute()

Get instance by ID



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
	metadataID := int32(56) // int32 | ID of the metadata
	instanceID := int32(56) // int32 | ID of the JuiceFS metadata instance
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.MetadatasMetadataIDInstancesInstanceIDGet(context.Background(), metadataID, instanceID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.MetadatasMetadataIDInstancesInstanceIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadatasMetadataIDInstancesInstanceIDGet`: MetadataInstance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.MetadatasMetadataIDInstancesInstanceIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataID** | **int32** | ID of the metadata | 
**instanceID** | **int32** | ID of the JuiceFS metadata instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadatasMetadataIDInstancesInstanceIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**MetadataInstance**](MetadataInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadatasMetadataIDInstancesInstanceIDPatch

> MetadataInstance MetadatasMetadataIDInstancesInstanceIDPatch(ctx, metadataID, instanceID).Token(token).MetadataInstance(metadataInstance).Execute()

Update instance properties



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
	metadataID := int32(56) // int32 | ID of the metadata
	instanceID := int32(56) // int32 | ID of the JuiceFS metadata instance
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)
	metadataInstance := *openapiclient.NewMetadataInstance() // MetadataInstance |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.MetadatasMetadataIDInstancesInstanceIDPatch(context.Background(), metadataID, instanceID).Token(token).MetadataInstance(metadataInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.MetadatasMetadataIDInstancesInstanceIDPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadatasMetadataIDInstancesInstanceIDPatch`: MetadataInstance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.MetadatasMetadataIDInstancesInstanceIDPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataID** | **int32** | ID of the metadata | 
**instanceID** | **int32** | ID of the JuiceFS metadata instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadatasMetadataIDInstancesInstanceIDPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **metadataInstance** | [**MetadataInstance**](MetadataInstance.md) |  | 

### Return type

[**MetadataInstance**](MetadataInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

