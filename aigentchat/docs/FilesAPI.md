# \FilesAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleFileUpload**](FilesAPI.md#HandleFileUpload) | **Post** /organizations/{org_id}/files/{recipe_name}/{channel_id} | Upload and process a file



## HandleFileUpload

> FileUploadResponse HandleFileUpload(ctx, orgId, channelId, recipeName).File(file).Execute()

Upload and process a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/aigentchat"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	channelId := "channelId_example" // string | channelID to receive updates
	recipeName := "recipeName_example" // string | Recipe name
	file := os.NewFile(1234, "some_file") // *os.File | File to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.HandleFileUpload(context.Background(), orgId, channelId, recipeName).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.HandleFileUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleFileUpload`: FileUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.HandleFileUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**channelId** | **string** | channelID to receive updates | 
**recipeName** | **string** | Recipe name | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **file** | ***os.File** | File to upload | 

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

