# \FilesAPI

All URIs are relative to *https://file-manager.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFileAccessToken**](FilesAPI.md#CreateFileAccessToken) | **Post** /v1/file-access-tokens | Create a file access token
[**DeleteFile**](FilesAPI.md#DeleteFile) | **Delete** /v1/files/{storage_path} | Delete a file
[**GetFileUploadCategories**](FilesAPI.md#GetFileUploadCategories) | **Get** /v1/organizations/{org_id}/files/categories | Get file upload categories
[**ServeFile**](FilesAPI.md#ServeFile) | **Get** /v1/files/{storage_path} | Serve a file
[**UploadFile**](FilesAPI.md#UploadFile) | **Post** /v1/organizations/{org_id}/files/{category} | Create a file for a channel



## CreateFileAccessToken

> FileAccessTokenResponse CreateFileAccessToken(ctx).FileAccessTokenRequest(fileAccessTokenRequest).Execute()

Create a file access token



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
	fileAccessTokenRequest := *openapiclient.NewFileAccessTokenRequest("StoragePath_example") // FileAccessTokenRequest | File access token request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.CreateFileAccessToken(context.Background()).FileAccessTokenRequest(fileAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.CreateFileAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileAccessToken`: FileAccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.CreateFileAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileAccessTokenRequest** | [**FileAccessTokenRequest**](FileAccessTokenRequest.md) | File access token request | 

### Return type

[**FileAccessTokenResponse**](FileAccessTokenResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFile

> FileMetadataResponse DeleteFile(ctx, storagePath).Execute()

Delete a file



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
	resp, r, err := apiClient.FilesAPI.DeleteFile(context.Background(), storagePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFile`: FileMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storagePath** | **string** | Storage path | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


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


## GetFileUploadCategories

> []FileUploadCategoryResponse GetFileUploadCategories(ctx, orgId).ModelCapabilities(modelCapabilities).Execute()

Get file upload categories



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
	orgId := "orgId_example" // string | organization ID
	modelCapabilities := "modelCapabilities_example" // string | Comma separated list of model capabilities to filter by, e.g. text-to-text,image-to-text (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileUploadCategories(context.Background(), orgId).ModelCapabilities(modelCapabilities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileUploadCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileUploadCategories`: []FileUploadCategoryResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileUploadCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileUploadCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelCapabilities** | **string** | Comma separated list of model capabilities to filter by, e.g. text-to-text,image-to-text | 

### Return type

[**[]FileUploadCategoryResponse**](FileUploadCategoryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServeFile

> *os.File ServeFile(ctx, storagePath).Download(download).Filename(filename).Fat(fat).Execute()

Serve a file



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
	download := true // bool | Force download as attachment (optional)
	filename := "filename_example" // string | Custom filename for download (optional)
	fat := "fat_example" // string | File access token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.ServeFile(context.Background(), storagePath).Download(download).Filename(filename).Fat(fat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.ServeFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServeFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.ServeFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storagePath** | **string** | Storage path | 

### Other Parameters

Other parameters are passed through a pointer to a apiServeFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **download** | **bool** | Force download as attachment | 
 **filename** | **string** | Custom filename for download | 
 **fat** | **string** | File access token | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, image/*, text/*, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> FileUploadResponse UploadFile(ctx, orgId, category).File(file).Metadata(metadata).Execute()

Create a file for a channel



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
	orgId := "orgId_example" // string | organization ID
	category := "category_example" // string | category ID
	file := os.NewFile(1234, "some_file") // *os.File | File to upload
	metadata := map[string]interface{}{ ... } // map[string]interface{} | Metadata for the uploaded file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadFile(context.Background(), orgId, category).File(file).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFile`: FileUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**category** | **string** | category ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | File to upload | 
 **metadata** | [**map[string]interface{}**](map[string]interface{}.md) | Metadata for the uploaded file | 

### Return type

[**FileUploadResponse**](FileUploadResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

