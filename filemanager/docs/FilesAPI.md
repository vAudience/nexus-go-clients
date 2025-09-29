# \FilesAPI

All URIs are relative to *https://file-manager.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFile**](FilesAPI.md#DeleteFile) | **Delete** /v1/files/{storage_path} | Delete a file
[**GetFileUploadSettings**](FilesAPI.md#GetFileUploadSettings) | **Get** /v1/organizations/{org_id}/files/{category}/settings | Get file upload settings for a category
[**ServeFile**](FilesAPI.md#ServeFile) | **Get** /v1/files/{storage_path} | Serve a file
[**UploadFile**](FilesAPI.md#UploadFile) | **Post** /v1/organizations/{org_id}/files/{category} | Create a file for a channel



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


## GetFileUploadSettings

> FileUploadSettings GetFileUploadSettings(ctx, orgId, category).Execute()

Get file upload settings for a category



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileUploadSettings(context.Background(), orgId, category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileUploadSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileUploadSettings`: FileUploadSettings
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileUploadSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**category** | **string** | category ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileUploadSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileUploadSettings**](FileUploadSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServeFile

> *os.File ServeFile(ctx, storagePath).Download(download).Filename(filename).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.ServeFile(context.Background(), storagePath).Download(download).Filename(filename).Execute()
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

