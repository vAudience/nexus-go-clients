# \FilesMetadataAPI

All URIs are relative to *https://file-manager.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFileMetadata**](FilesMetadataAPI.md#GetFileMetadata) | **Get** /v1/metadata/{storage_path} | Get file metadata by path
[**GetFileMetadataBatch**](FilesMetadataAPI.md#GetFileMetadataBatch) | **Post** /v1/metadata | Get metadata for multiple files



## GetFileMetadata

> MetadataFileMetadataResponse GetFileMetadata(ctx, storagePath).Execute()

Get file metadata by path



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	storagePath := "storagePath_example" // string | Storage path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.GetFileMetadata(context.Background(), storagePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.GetFileMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileMetadata`: MetadataFileMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.GetFileMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storagePath** | **string** | Storage path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataFileMetadataResponse**](MetadataFileMetadataResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileMetadataBatch

> MetadataFileMetadataBatchResponse GetFileMetadataBatch(ctx).Batchrequest(batchrequest).Execute()

Get metadata for multiple files



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	batchrequest := *openapiclient.NewMetadataGetFileMetadataBatchRequest([]string{"StoragePaths_example"}) // MetadataGetFileMetadataBatchRequest | Batch request with file paths

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.GetFileMetadataBatch(context.Background()).Batchrequest(batchrequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.GetFileMetadataBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileMetadataBatch`: MetadataFileMetadataBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.GetFileMetadataBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileMetadataBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchrequest** | [**MetadataGetFileMetadataBatchRequest**](MetadataGetFileMetadataBatchRequest.md) | Batch request with file paths | 

### Return type

[**MetadataFileMetadataBatchResponse**](MetadataFileMetadataBatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

