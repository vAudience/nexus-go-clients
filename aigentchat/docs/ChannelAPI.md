# \ChannelAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](ChannelAPI.md#CreateChannel) | **Post** /organizations/{org_id}/channels | Create a new channel
[**DeleteChannel**](ChannelAPI.md#DeleteChannel) | **Delete** /organizations/{org_id}/channels/{id} | Delete a channel
[**GetActiveChannels**](ChannelAPI.md#GetActiveChannels) | **Get** /organizations/{org_id}/channels/active | Get active channels
[**GetChannel**](ChannelAPI.md#GetChannel) | **Get** /organizations/{org_id}/channels/{id} | Get a channel by ID
[**GetChannelPresence**](ChannelAPI.md#GetChannelPresence) | **Get** /organizations/{org_id}/channels/{channel_id}/presence | Get channel presence
[**GetUserSubscribedChannels**](ChannelAPI.md#GetUserSubscribedChannels) | **Get** /organizations/{org_id}/channels/subscribed/{user_id} | Get user&#39;s subscribed channels
[**ListChannelsByOrgId**](ChannelAPI.md#ListChannelsByOrgId) | **Get** /organizations/{org_id}/channels | List channels by organization ID
[**ListChannelsByOwnerId**](ChannelAPI.md#ListChannelsByOwnerId) | **Get** /organizations/{org_id}/channels/owner | List channels by owner ID
[**UpdateChannel**](ChannelAPI.md#UpdateChannel) | **Put** /organizations/{org_id}/channels/{id} | Update a channel



## CreateChannel

> Channel CreateChannel(ctx, orgId).Channel(channel).Execute()

Create a new channel



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
	channel := *openapiclient.NewChannelWriteDto() // ChannelWriteDto | Channel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.CreateChannel(context.Background(), orgId).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.CreateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.CreateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | [**ChannelWriteDto**](ChannelWriteDto.md) | Channel | 

### Return type

[**Channel**](Channel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> Channel DeleteChannel(ctx, orgId, id).Execute()

Delete a channel



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
	id := "id_example" // string | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.DeleteChannel(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.DeleteChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.DeleteChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Channel**](Channel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveChannels

> []Channel GetActiveChannels(ctx, orgId).Execute()

Get active channels



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
	resp, r, err := apiClient.ChannelAPI.GetActiveChannels(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetActiveChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveChannels`: []Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetActiveChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Channel**](Channel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannel

> Channel GetChannel(ctx, orgId, id).Execute()

Get a channel by ID



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
	id := "id_example" // string | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.GetChannel(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Channel**](Channel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelPresence

> []string GetChannelPresence(ctx, orgId, channelId).Execute()

Get channel presence



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
	resp, r, err := apiClient.ChannelAPI.GetChannelPresence(context.Background(), orgId, channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetChannelPresence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelPresence`: []string
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannelPresence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**channelId** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelPresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSubscribedChannels

> []Channel GetUserSubscribedChannels(ctx, orgId, userId).Execute()

Get user's subscribed channels



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
	userId := "userId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.GetUserSubscribedChannels(context.Background(), orgId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetUserSubscribedChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSubscribedChannels`: []Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetUserSubscribedChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSubscribedChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Channel**](Channel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelsByOrgId

> []Channel ListChannelsByOrgId(ctx, orgId).Execute()

List channels by organization ID



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
	resp, r, err := apiClient.ChannelAPI.ListChannelsByOrgId(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ListChannelsByOrgId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChannelsByOrgId`: []Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ListChannelsByOrgId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListChannelsByOrgIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Channel**](Channel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelsByOwnerId

> []Channel ListChannelsByOwnerId(ctx, orgId).Execute()

List channels by owner ID



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
	resp, r, err := apiClient.ChannelAPI.ListChannelsByOwnerId(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ListChannelsByOwnerId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChannelsByOwnerId`: []Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ListChannelsByOwnerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListChannelsByOwnerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Channel**](Channel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> Channel UpdateChannel(ctx, orgId, id).Channel(channel).Execute()

Update a channel



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
	id := "id_example" // string | Channel ID
	channel := *openapiclient.NewChannelWriteDto() // ChannelWriteDto | Channel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.UpdateChannel(context.Background(), orgId, id).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.UpdateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.UpdateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **channel** | [**ChannelWriteDto**](ChannelWriteDto.md) | Channel | 

### Return type

[**Channel**](Channel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

