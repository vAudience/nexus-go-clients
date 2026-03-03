# \PromptsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrompt**](PromptsAPI.md#CreatePrompt) | **Post** /v1/organizations/{org_id}/prompts | Create a new prompt
[**DeletePrompt**](PromptsAPI.md#DeletePrompt) | **Delete** /v1/organizations/{org_id}/prompts/{prompt_id} | Delete a prompt
[**DuplicatePrompt**](PromptsAPI.md#DuplicatePrompt) | **Post** /v1/organizations/{org_id}/prompts/{prompt_id}/duplicate | Duplicate Prompt
[**GetPrompt**](PromptsAPI.md#GetPrompt) | **Get** /v1/organizations/{org_id}/prompts/{prompt_id} | Get a specific prompt
[**ListPrompts**](PromptsAPI.md#ListPrompts) | **Get** /v1/organizations/{org_id}/prompts | List prompts
[**RenderPrompt**](PromptsAPI.md#RenderPrompt) | **Post** /v1/organizations/{org_id}/prompts/render | Render Prompt
[**UpdatePrompt**](PromptsAPI.md#UpdatePrompt) | **Put** /v1/organizations/{org_id}/prompts/{prompt_id} | Update a prompt



## CreatePrompt

> Prompt CreatePrompt(ctx, orgId).Prompt(prompt).Execute()

Create a new prompt



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
	prompt := *openapiclient.NewPromptWriteDto() // PromptWriteDto | Prompt object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.CreatePrompt(context.Background(), orgId).Prompt(prompt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CreatePrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrompt`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.CreatePrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prompt** | [**PromptWriteDto**](PromptWriteDto.md) | Prompt object | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrompt

> DeletePrompt(ctx, orgId, promptId).Execute()

Delete a prompt



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
	promptId := "promptId_example" // string | Prompt ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PromptsAPI.DeletePrompt(context.Background(), orgId, promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.DeletePrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**promptId** | **string** | Prompt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuplicatePrompt

> Prompt DuplicatePrompt(ctx, orgId, promptId).Execute()

Duplicate Prompt



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
	promptId := "promptId_example" // string | Prompt ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.DuplicatePrompt(context.Background(), orgId, promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.DuplicatePrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DuplicatePrompt`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.DuplicatePrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**promptId** | **string** | Prompt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicatePromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Prompt**](Prompt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrompt

> Prompt GetPrompt(ctx, orgId, promptId).Execute()

Get a specific prompt



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
	promptId := "promptId_example" // string | Prompt ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.GetPrompt(context.Background(), orgId, promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrompt`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**promptId** | **string** | Prompt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Prompt**](Prompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrompts

> PromptResults ListPrompts(ctx, orgId).OwnerId(ownerId).Tags(tags).SystemTags(systemTags).Q(q).Visibility(visibility).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()

List prompts



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
	orgId := "orgId_example" // string | Organization ID to filter by
	ownerId := "ownerId_example" // string | Owner ID to filter by (optional)
	tags := []string{"Inner_example"} // []string | Tags to filter by (comma separated) (optional)
	systemTags := []string{"Inner_example"} // []string | System tags to filter by (comma separated) (optional)
	q := "q_example" // string | Search term for title, description or tags (optional)
	visibility := "visibility_example" // string | Filter prompts by access visibility (public, private, organization) (optional)
	limit := int32(56) // int32 | Limit the number of results (optional) (default to 1000)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)
	sortBy := "sortBy_example" // string | Field to sort by (title, createdat, updatedat) (optional) (default to "\"title\"")
	sortOrder := "sortOrder_example" // string | Sort order (asc or desc) (optional) (default to "\"asc\"")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.ListPrompts(context.Background(), orgId).OwnerId(ownerId).Tags(tags).SystemTags(systemTags).Q(q).Visibility(visibility).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.ListPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrompts`: PromptResults
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.ListPrompts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID to filter by | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPromptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ownerId** | **string** | Owner ID to filter by | 
 **tags** | **[]string** | Tags to filter by (comma separated) | 
 **systemTags** | **[]string** | System tags to filter by (comma separated) | 
 **q** | **string** | Search term for title, description or tags | 
 **visibility** | **string** | Filter prompts by access visibility (public, private, organization) | 
 **limit** | **int32** | Limit the number of results | [default to 1000]
 **offset** | **int32** | Offset for pagination | [default to 0]
 **sortBy** | **string** | Field to sort by (title, createdat, updatedat) | [default to &quot;\&quot;title\&quot;&quot;]
 **sortOrder** | **string** | Sort order (asc or desc) | [default to &quot;\&quot;asc\&quot;&quot;]

### Return type

[**PromptResults**](PromptResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderPrompt

> RenderedPrompt RenderPrompt(ctx, orgId).RenderDto(renderDto).Execute()

Render Prompt



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
	renderDto := *openapiclient.NewPromptRenderDto() // PromptRenderDto | Prompt Render DTO

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.RenderPrompt(context.Background(), orgId).RenderDto(renderDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.RenderPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderPrompt`: RenderedPrompt
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.RenderPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenderPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renderDto** | [**PromptRenderDto**](PromptRenderDto.md) | Prompt Render DTO | 

### Return type

[**RenderedPrompt**](RenderedPrompt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrompt

> Prompt UpdatePrompt(ctx, orgId, promptId).Prompt(prompt).Execute()

Update a prompt



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
	promptId := "promptId_example" // string | Prompt ID
	prompt := *openapiclient.NewPromptWriteDto() // PromptWriteDto | Updated Prompt object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.UpdatePrompt(context.Background(), orgId, promptId).Prompt(prompt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.UpdatePrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrompt`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.UpdatePrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**promptId** | **string** | Prompt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **prompt** | [**PromptWriteDto**](PromptWriteDto.md) | Updated Prompt object | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

