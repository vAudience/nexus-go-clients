# \FilesMetadataAPI

All URIs are relative to *https://file-manager.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFileMetadata**](FilesMetadataAPI.md#GetFileMetadata) | **Get** /v1/metadata/{storage_path} | Get file metadata by path
[**GetFileMetadataBatch**](FilesMetadataAPI.md#GetFileMetadataBatch) | **Post** /v1/metadata | Get metadata for multiple files
[**ListFileMetadata**](FilesMetadataAPI.md#ListFileMetadata) | **Get** /v1/metadata | List file metadata



## GetFileMetadata

> FileMetadataResponse GetFileMetadata(ctx, storagePath).Execute()

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
	// response from `GetFileMetadata`: FileMetadataResponse
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

[**FileMetadataResponse**](FileMetadataResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileMetadataBatch

> FileMetadataBatchResponse GetFileMetadataBatch(ctx).Batchrequest(batchrequest).Execute()

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
	batchrequest := *openapiclient.NewFileMetadataBatchGetRequest([]string{"StoragePaths_example"}) // FileMetadataBatchGetRequest | Batch request with file paths

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.GetFileMetadataBatch(context.Background()).Batchrequest(batchrequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.GetFileMetadataBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileMetadataBatch`: FileMetadataBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.GetFileMetadataBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileMetadataBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchrequest** | [**FileMetadataBatchGetRequest**](FileMetadataBatchGetRequest.md) | Batch request with file paths | 

### Return type

[**FileMetadataBatchResponse**](FileMetadataBatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFileMetadata

> FileMetadataListResponse ListFileMetadata(ctx).Scope(scope).UserId(userId).OrgId(orgId).TeamId(teamId).UploadCategory(uploadCategory).Offset(offset).Limit(limit).Sort(sort).Order(order).Execute()

List file metadata



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
	scope := "scope_example" // string | Scope: temp, public, private-org, private-org-user, private-team
	userId := "userId_example" // string | User ID filter (optional)
	orgId := "orgId_example" // string | Organization ID filter (optional)
	teamId := "teamId_example" // string | Team ID filter (optional)
	uploadCategory := "uploadCategory_example" // string | Upload category filter (optional)
	offset := int32(56) // int32 | Offset (default: 0) (optional)
	limit := int32(56) // int32 | Limit (default: 50, max: 100) (optional)
	sort := "sort_example" // string | Sort field: created_at, file_name, file_size (default: created_at) (optional)
	order := "order_example" // string | Sort order: asc, desc (default: desc) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.ListFileMetadata(context.Background()).Scope(scope).UserId(userId).OrgId(orgId).TeamId(teamId).UploadCategory(uploadCategory).Offset(offset).Limit(limit).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.ListFileMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFileMetadata`: FileMetadataListResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.ListFileMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFileMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope: temp, public, private-org, private-org-user, private-team | 
 **userId** | **string** | User ID filter | 
 **orgId** | **string** | Organization ID filter | 
 **teamId** | **string** | Team ID filter | 
 **uploadCategory** | **string** | Upload category filter | 
 **offset** | **int32** | Offset (default: 0) | 
 **limit** | **int32** | Limit (default: 50, max: 100) | 
 **sort** | **string** | Sort field: created_at, file_name, file_size (default: created_at) | 
 **order** | **string** | Sort order: asc, desc (default: desc) | 

### Return type

[**FileMetadataListResponse**](FileMetadataListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

