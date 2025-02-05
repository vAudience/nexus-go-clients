# \AudioAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAudioTranscription**](AudioAPI.md#CreateAudioTranscription) | **Post** /v1/organizations/{org_id}/audio/transcribe | Transcribe speech to text
[**CreateTextToSpeech**](AudioAPI.md#CreateTextToSpeech) | **Post** /v1/organizations/{org_id}/audio/tts | Text-To-Speech



## CreateAudioTranscription

> AIgencyMessage CreateAudioTranscription(ctx, orgId).Request(request).Execute()

Transcribe speech to text



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
	request := *openapiclient.NewAudioTranscriptionRequest("Language_example", "TranscriptionFormat_example") // AudioTranscriptionRequest | Audio Transcription Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.CreateAudioTranscription(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.CreateAudioTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAudioTranscription`: AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.CreateAudioTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudioTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**AudioTranscriptionRequest**](AudioTranscriptionRequest.md) | Audio Transcription Request | 

### Return type

[**AIgencyMessage**](AIgencyMessage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTextToSpeech

> AIgencyMessage CreateTextToSpeech(ctx, orgId).Request(request).Execute()

Text-To-Speech



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
	request := *openapiclient.NewAudioGenerationRequest() // AudioGenerationRequest | Audio-Generation Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.CreateTextToSpeech(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.CreateTextToSpeech``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTextToSpeech`: AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.CreateTextToSpeech`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTextToSpeechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**AudioGenerationRequest**](AudioGenerationRequest.md) | Audio-Generation Request | 

### Return type

[**AIgencyMessage**](AIgencyMessage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

