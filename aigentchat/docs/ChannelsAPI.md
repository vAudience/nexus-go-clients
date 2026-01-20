# \ChannelsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](ChannelsAPI.md#CreateChannel) | **Post** /v1/organizations/{org_id}/channels | Create a new channel
[**DeleteChannel**](ChannelsAPI.md#DeleteChannel) | **Delete** /v1/organizations/{org_id}/channels/{id} | Delete a channel
[**DeleteChannelsByOwnerId**](ChannelsAPI.md#DeleteChannelsByOwnerId) | **Delete** /v1/organizations/{org_id}/channels | Delete channels by their owner ID
[**GetChannel**](ChannelsAPI.md#GetChannel) | **Get** /v1/organizations/{org_id}/channels/{id} | Get a channel by ID
[**ListChannelsByOrgId**](ChannelsAPI.md#ListChannelsByOrgId) | **Get** /v1/organizations/{org_id}/channels | List channels by organization ID
[**SearchChannels**](ChannelsAPI.md#SearchChannels) | **Get** /v1/organizations/{org_id}/channels/search | Search channels by query
[**UpdateChannel**](ChannelsAPI.md#UpdateChannel) | **Put** /v1/organizations/{org_id}/channels/{id} | Update a channel



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
	resp, r, err := apiClient.ChannelsAPI.CreateChannel(context.Background(), orgId).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CreateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CreateChannel`: %v\n", resp)
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
	resp, r, err := apiClient.ChannelsAPI.DeleteChannel(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.DeleteChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.DeleteChannel`: %v\n", resp)
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


## DeleteChannelsByOwnerId

> []Channel DeleteChannelsByOwnerId(ctx, orgId).Execute()

Delete channels by their owner ID



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
	resp, r, err := apiClient.ChannelsAPI.DeleteChannelsByOwnerId(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.DeleteChannelsByOwnerId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChannelsByOwnerId`: []Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.DeleteChannelsByOwnerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChannelsByOwnerIdRequest struct via the builder pattern


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
	resp, r, err := apiClient.ChannelsAPI.GetChannel(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannel`: %v\n", resp)
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
	resp, r, err := apiClient.ChannelsAPI.ListChannelsByOrgId(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.ListChannelsByOrgId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChannelsByOrgId`: []Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.ListChannelsByOrgId`: %v\n", resp)
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


## SearchChannels

> ChannelResults SearchChannels(ctx, orgId).UserId(userId).Q(q).IncludeServices(includeServices).Limit(limit).Offset(offset).OffsetChannelId(offsetChannelId).Execute()

Search channels by query



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
	userId := "userId_example" // string | user ID or me for current user (optional)
	q := "q_example" // string | Search term for name, description or summary (optional)
	includeServices := true // bool | Whether to include service channels in the results (optional) (default to false)
	limit := int32(56) // int32 | Limit the number of results (optional) (default to 1000)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)
	offsetChannelId := "offsetChannelId_example" // string | Offset the results to center around the specified channel ID, disables offset parameter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.SearchChannels(context.Background(), orgId).UserId(userId).Q(q).IncludeServices(includeServices).Limit(limit).Offset(offset).OffsetChannelId(offsetChannelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.SearchChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchChannels`: ChannelResults
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.SearchChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | user ID or me for current user | 
 **q** | **string** | Search term for name, description or summary | 
 **includeServices** | **bool** | Whether to include service channels in the results | [default to false]
 **limit** | **int32** | Limit the number of results | [default to 1000]
 **offset** | **int32** | Offset for pagination | [default to 0]
 **offsetChannelId** | **string** | Offset the results to center around the specified channel ID, disables offset parameter | 

### Return type

[**ChannelResults**](ChannelResults.md)

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
	resp, r, err := apiClient.ChannelsAPI.UpdateChannel(context.Background(), orgId, id).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.UpdateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.UpdateChannel`: %v\n", resp)
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

