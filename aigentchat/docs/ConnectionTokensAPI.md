# \ConnectionTokensAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionTokenHandler**](ConnectionTokensAPI.md#CreateConnectionTokenHandler) | **Post** /v1/organizations/{org_id}/connection-tokens/{channel_id} | Create a connection token for SSE



## CreateConnectionTokenHandler

> ConnectionTokenResponse CreateConnectionTokenHandler(ctx, orgId, channelId).Execute()

Create a connection token for SSE



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
	channelId := "channelId_example" // string | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionTokensAPI.CreateConnectionTokenHandler(context.Background(), orgId, channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionTokensAPI.CreateConnectionTokenHandler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectionTokenHandler`: ConnectionTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectionTokensAPI.CreateConnectionTokenHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**channelId** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionTokenHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectionTokenResponse**](ConnectionTokenResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

