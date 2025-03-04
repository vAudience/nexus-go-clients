# \MessagesAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessage**](MessagesAPI.md#CreateMessage) | **Post** /v1/organizations/{org_id}/messages | Create a new message
[**DeleteMessage**](MessagesAPI.md#DeleteMessage) | **Delete** /v1/organizations/{org_id}/messages/{id} | Delete a message
[**GetChannelMessages**](MessagesAPI.md#GetChannelMessages) | **Get** /v1/organizations/{org_id}/messages/channel/{channel_id} | Get messages for a channel
[**GetMessage**](MessagesAPI.md#GetMessage) | **Get** /v1/organizations/{org_id}/messages/{id} | Get a message by ID
[**SearchMessages**](MessagesAPI.md#SearchMessages) | **Get** /v1/organizations/{org_id}/messages/search | Search messages
[**UpdateMessage**](MessagesAPI.md#UpdateMessage) | **Put** /v1/organizations/{org_id}/messages/{id} | Update a message



## CreateMessage

> AIgencyMessage CreateMessage(ctx, orgId).Message(message).Execute()

Create a new message



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
	message := *openapiclient.NewAIgencyMessageWriteDto("AiModelId_example", "AiServiceId_example", "ChannelId_example", "ChannelName_example", "MissionId_example", openapiclient.ConversationRole("unknown"), "SenderName_example", openapiclient.AIgencyMessageType("message")) // AIgencyMessageWriteDto | Message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateMessage(context.Background(), orgId).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessage`: AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | [**AIgencyMessageWriteDto**](AIgencyMessageWriteDto.md) | Message | 

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


## DeleteMessage

> []AIgencyMessage DeleteMessage(ctx, orgId, id).Cascade(cascade).Execute()

Delete a message



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
	id := "id_example" // string | Message ID
	cascade := true // bool | Delete related message (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.DeleteMessage(context.Background(), orgId, id).Cascade(cascade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.DeleteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMessage`: []AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.DeleteMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cascade** | **bool** | Delete related message | 

### Return type

[**[]AIgencyMessage**](AIgencyMessage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelMessages

> []AIgencyMessage GetChannelMessages(ctx, orgId, channelId).Execute()

Get messages for a channel



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
	resp, r, err := apiClient.MessagesAPI.GetChannelMessages(context.Background(), orgId, channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetChannelMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelMessages`: []AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetChannelMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**channelId** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AIgencyMessage**](AIgencyMessage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessage

> AIgencyMessage GetMessage(ctx, orgId, id).Execute()

Get a message by ID



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
	id := "id_example" // string | Message ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.GetMessage(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessage`: AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AIgencyMessage**](AIgencyMessage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMessages

> AIgencyMessageResults SearchMessages(ctx, orgId).Content(content).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

Search messages



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
	content := "content_example" // string | Search by content (optional)
	startDate := "startDate_example" // string | Start date in Unix milliseconds (optional)
	endDate := "endDate_example" // string | End date in Unix milliseconds (optional)
	offset := int32(56) // int32 | Offset (optional)
	limit := int32(56) // int32 | Limit results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.SearchMessages(context.Background(), orgId).Content(content).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.SearchMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMessages`: AIgencyMessageResults
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.SearchMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** | Search by content | 
 **startDate** | **string** | Start date in Unix milliseconds | 
 **endDate** | **string** | End date in Unix milliseconds | 
 **offset** | **int32** | Offset | 
 **limit** | **int32** | Limit results | 

### Return type

[**AIgencyMessageResults**](AIgencyMessageResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> AIgencyMessage UpdateMessage(ctx, orgId, id).Message(message).Execute()

Update a message



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
	id := "id_example" // string | Message ID
	message := *openapiclient.NewAIgencyMessageWriteDto("AiModelId_example", "AiServiceId_example", "ChannelId_example", "ChannelName_example", "MissionId_example", openapiclient.ConversationRole("unknown"), "SenderName_example", openapiclient.AIgencyMessageType("message")) // AIgencyMessageWriteDto | Message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.UpdateMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UpdateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessage`: AIgencyMessage
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.UpdateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | [**AIgencyMessageWriteDto**](AIgencyMessageWriteDto.md) | Message | 

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

