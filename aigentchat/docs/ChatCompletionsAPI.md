# \ChatCompletionsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelChatCompletion**](ChatCompletionsAPI.md#CancelChatCompletion) | **Post** /v1/organizations/{org_id}/completions/cancel/{channel_id} | Cancel a chat-completion
[**CreateChatCompletion**](ChatCompletionsAPI.md#CreateChatCompletion) | **Post** /v1/organizations/{org_id}/completions | Create a chat-completion
[**CreateChatCompletionStreaming**](ChatCompletionsAPI.md#CreateChatCompletionStreaming) | **Post** /v1/organizations/{org_id}/completions/stream | Create a streaming chat-completion
[**GetChatCompletionFileSettings**](ChatCompletionsAPI.md#GetChatCompletionFileSettings) | **Get** /v1/organizations/{org_id}/completions/file-settings | Get file settings for chat completions



## CancelChatCompletion

> CancelChatCompletion(ctx, orgId, channelId).Execute()

Cancel a chat-completion



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
	r, err := apiClient.ChatCompletionsAPI.CancelChatCompletion(context.Background(), orgId, channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatCompletionsAPI.CancelChatCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**channelId** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelChatCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChatCompletion

> []AIgencyMessage CreateChatCompletion(ctx, orgId).Request(request).Execute()

Create a chat-completion



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
	request := *openapiclient.NewChatCompletionRequestDto("AgentId_example", "Message_example") // ChatCompletionRequestDto | Chat Completion Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatCompletionsAPI.CreateChatCompletion(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatCompletionsAPI.CreateChatCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChatCompletion`: []AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `ChatCompletionsAPI.CreateChatCompletion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ChatCompletionRequestDto**](ChatCompletionRequestDto.md) | Chat Completion Request | 

### Return type

[**[]AIgencyMessage**](AIgencyMessage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChatCompletionStreaming

> []AIgencyMessage CreateChatCompletionStreaming(ctx, orgId).Request(request).Execute()

Create a streaming chat-completion



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
	request := *openapiclient.NewChatCompletionRequestDto("AgentId_example", "Message_example") // ChatCompletionRequestDto | Chat Completion Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatCompletionsAPI.CreateChatCompletionStreaming(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatCompletionsAPI.CreateChatCompletionStreaming``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChatCompletionStreaming`: []AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `ChatCompletionsAPI.CreateChatCompletionStreaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatCompletionStreamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ChatCompletionRequestDto**](ChatCompletionRequestDto.md) | Chat Completion Request | 

### Return type

[**[]AIgencyMessage**](AIgencyMessage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatCompletionFileSettings

> ChatCompletionFileSettings GetChatCompletionFileSettings(ctx, orgId).Execute()

Get file settings for chat completions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatCompletionsAPI.GetChatCompletionFileSettings(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatCompletionsAPI.GetChatCompletionFileSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatCompletionFileSettings`: ChatCompletionFileSettings
	fmt.Fprintf(os.Stdout, "Response from `ChatCompletionsAPI.GetChatCompletionFileSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatCompletionFileSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatCompletionFileSettings**](ChatCompletionFileSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

