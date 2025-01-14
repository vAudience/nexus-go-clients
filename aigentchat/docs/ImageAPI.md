# \ImageAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](ImageAPI.md#CreateImage) | **Post** /organizations/{org_id}/image/generate | Generates an image



## CreateImage

> AIgencyImage CreateImage(ctx, orgId).Request(request).Execute()

Generates an image



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
	request := *openapiclient.NewImageGenerationRequest() // ImageGenerationRequest | Image Generation Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.CreateImage(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.CreateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImage`: AIgencyImage
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.CreateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ImageGenerationRequest**](ImageGenerationRequest.md) | Image Generation Request | 

### Return type

[**AIgencyImage**](AIgencyImage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

