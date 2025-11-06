# \AudioAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoiceChatInput**](AudioAPI.md#CreateVoiceChatInput) | **Post** /v1/organizations/{org_id}/audio/voice-chat-input | Voice Chat Input



## CreateVoiceChatInput

> AudioTranscription CreateVoiceChatInput(ctx, orgId).Request(request).Execute()

Voice Chat Input



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
	request := *openapiclient.NewVoiceChatInputRequestDto("FileUrl_example") // VoiceChatInputRequestDto | Voice Chat Input Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.CreateVoiceChatInput(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.CreateVoiceChatInput``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoiceChatInput`: AudioTranscription
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.CreateVoiceChatInput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoiceChatInputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**VoiceChatInputRequestDto**](VoiceChatInputRequestDto.md) | Voice Chat Input Request | 

### Return type

[**AudioTranscription**](AudioTranscription.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

