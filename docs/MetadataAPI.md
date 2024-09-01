# \MetadataAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadataDnsGet**](MetadataAPI.md#MetadataDnsGet) | **Get** /metadata-dns | Resolve domain name
[**MetadatasGet**](MetadataAPI.md#MetadatasGet) | **Get** /metadatas | Get metadata list
[**MetadatasMetadataIDGet**](MetadataAPI.md#MetadatasMetadataIDGet) | **Get** /metadatas/{metadataID} | Get metadata by ID
[**MetadatasMetadataIDLicenseGet**](MetadataAPI.md#MetadatasMetadataIDLicenseGet) | **Get** /metadatas/{metadataID}/license | Get current license
[**MetadatasMetadataIDLicensePut**](MetadataAPI.md#MetadatasMetadataIDLicensePut) | **Put** /metadatas/{metadataID}/license | Set license



## MetadataDnsGet

> MetadataDnsGet200Response MetadataDnsGet(ctx).Name(name).Cache(cache).Execute()

Resolve domain name



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
	name := "abc.meta.juicefs.com" // string | 
	cache := TODO // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataDnsGet(context.Background()).Name(name).Cache(cache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataDnsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataDnsGet`: MetadataDnsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataDnsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataDnsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **cache** | [**bool**](bool.md) |  | 

### Return type

[**MetadataDnsGet200Response**](MetadataDnsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadatasGet

> []Metadata MetadatasGet(ctx).Token(token).Execute()

Get metadata list



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
	token := "token_example" // string | Token of the region, it's once used for authentication, now deprecated. please use access key/secret key instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadatasGet(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadatasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadatasGet`: []Metadata
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadatasGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadatasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**[]Metadata**](Metadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadatasMetadataIDGet

> Metadata MetadatasMetadataIDGet(ctx, metadataID).Token(token).Execute()

Get metadata by ID



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
	resp, r, err := apiClient.MetadataAPI.MetadatasMetadataIDGet(context.Background(), metadataID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadatasMetadataIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadatasMetadataIDGet`: Metadata
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadatasMetadataIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataID** | **int32** | ID of the metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadatasMetadataIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**Metadata**](Metadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadatasMetadataIDLicenseGet

> MetadataLicense MetadatasMetadataIDLicenseGet(ctx, metadataID).Token(token).Execute()

Get current license



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
	resp, r, err := apiClient.MetadataAPI.MetadatasMetadataIDLicenseGet(context.Background(), metadataID).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadatasMetadataIDLicenseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadatasMetadataIDLicenseGet`: MetadataLicense
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadatasMetadataIDLicenseGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataID** | **int32** | ID of the metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadatasMetadataIDLicenseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 

### Return type

[**MetadataLicense**](MetadataLicense.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadatasMetadataIDLicensePut

> MetadataLicense MetadatasMetadataIDLicensePut(ctx, metadataID).Token(token).MetadataLicense(metadataLicense).Execute()

Set license



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
	metadataLicense := *openapiclient.NewMetadataLicense() // MetadataLicense |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadatasMetadataIDLicensePut(context.Background(), metadataID).Token(token).MetadataLicense(metadataLicense).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadatasMetadataIDLicensePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadatasMetadataIDLicensePut`: MetadataLicense
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadatasMetadataIDLicensePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataID** | **int32** | ID of the metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadatasMetadataIDLicensePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead. | 
 **metadataLicense** | [**MetadataLicense**](MetadataLicense.md) |  | 

### Return type

[**MetadataLicense**](MetadataLicense.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

