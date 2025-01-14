# \EmbeddingsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmbedText**](EmbeddingsAPI.md#EmbedText) | **Post** /organizations/{org_id}/embeddings/text | Get the embeddings for the texts



## EmbedText

> TextEmbedding EmbedText(ctx, orgId).Request(request).Execute()

Get the embeddings for the texts



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
	orgId := "orgId_example" // string | Organization ID
	request := *openapiclient.NewTextEmbeddingRequestDto("AgentId_example", []openapiclient.TextEmbeddingItemDto{*openapiclient.NewTextEmbeddingItemDto("Text_example")}) // TextEmbeddingRequestDto | Text Embedding Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.EmbedText(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.EmbedText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmbedText`: TextEmbedding
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.EmbedText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmbedTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**TextEmbeddingRequestDto**](TextEmbeddingRequestDto.md) | Text Embedding Request | 

### Return type

[**TextEmbedding**](TextEmbedding.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

