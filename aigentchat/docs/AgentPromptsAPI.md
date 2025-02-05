# \AgentPromptsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgentPrompt**](AgentPromptsAPI.md#CreateAgentPrompt) | **Post** /v1/organizations/{org_id}/agent-prompts | Create a new agent prompt
[**DeleteAgentPrompt**](AgentPromptsAPI.md#DeleteAgentPrompt) | **Delete** /v1/organizations/{org_id}/agent-prompts/{prompt_id} | Delete an agent prompt
[**GetAgentPrompt**](AgentPromptsAPI.md#GetAgentPrompt) | **Get** /v1/organizations/{org_id}/agent-prompts/{prompt_id} | Get a specific agent prompt
[**ListAgentPrompts**](AgentPromptsAPI.md#ListAgentPrompts) | **Get** /v1/organizations/{org_id}/agent-prompts | List agent prompts
[**RenderAgentPrompt**](AgentPromptsAPI.md#RenderAgentPrompt) | **Post** /v1/organizations/{org_id}/agent-prompts/render | Render Agent Prompt
[**UpdateAgentPrompt**](AgentPromptsAPI.md#UpdateAgentPrompt) | **Put** /v1/organizations/{org_id}/agent-prompts/{prompt_id} | Update an agent prompt



## CreateAgentPrompt

> AgentPrompt CreateAgentPrompt(ctx, orgId).Prompt(prompt).Execute()

Create a new agent prompt



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
	prompt := *openapiclient.NewAgentPrompt("Id_example", "OwnerId_example", "OwnerOrganizationId_example", "Title_example") // AgentPrompt | Agent Prompt object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentPromptsAPI.CreateAgentPrompt(context.Background(), orgId).Prompt(prompt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentPromptsAPI.CreateAgentPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgentPrompt`: AgentPrompt
	fmt.Fprintf(os.Stdout, "Response from `AgentPromptsAPI.CreateAgentPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prompt** | [**AgentPrompt**](AgentPrompt.md) | Agent Prompt object | 

### Return type

[**AgentPrompt**](AgentPrompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentPrompt

> DeleteAgentPrompt(ctx, orgId, promptId).Execute()

Delete an agent prompt



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
	r, err := apiClient.AgentPromptsAPI.DeleteAgentPrompt(context.Background(), orgId, promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentPromptsAPI.DeleteAgentPrompt``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAgentPromptRequest struct via the builder pattern


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


## GetAgentPrompt

> AgentPrompt GetAgentPrompt(ctx, orgId, promptId).Execute()

Get a specific agent prompt



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
	resp, r, err := apiClient.AgentPromptsAPI.GetAgentPrompt(context.Background(), orgId, promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentPromptsAPI.GetAgentPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentPrompt`: AgentPrompt
	fmt.Fprintf(os.Stdout, "Response from `AgentPromptsAPI.GetAgentPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**promptId** | **string** | Prompt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentPrompt**](AgentPrompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentPrompts

> []AgentPrompt ListAgentPrompts(ctx, orgId).Offset(offset).Limit(limit).Execute()

List agent prompts



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
	offset := int32(56) // int32 | Offset (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentPromptsAPI.ListAgentPrompts(context.Background(), orgId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentPromptsAPI.ListAgentPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentPrompts`: []AgentPrompt
	fmt.Fprintf(os.Stdout, "Response from `AgentPromptsAPI.ListAgentPrompts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentPromptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Offset | 
 **limit** | **int32** | Limit | 

### Return type

[**[]AgentPrompt**](AgentPrompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderAgentPrompt

> AgentPromptRenderedDto RenderAgentPrompt(ctx, orgId).RenderDto(renderDto).Execute()

Render Agent Prompt



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
	renderDto := *openapiclient.NewAgentPromptRenderDto() // AgentPromptRenderDto | Agent Prompt Render DTO

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentPromptsAPI.RenderAgentPrompt(context.Background(), orgId).RenderDto(renderDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentPromptsAPI.RenderAgentPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderAgentPrompt`: AgentPromptRenderedDto
	fmt.Fprintf(os.Stdout, "Response from `AgentPromptsAPI.RenderAgentPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenderAgentPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renderDto** | [**AgentPromptRenderDto**](AgentPromptRenderDto.md) | Agent Prompt Render DTO | 

### Return type

[**AgentPromptRenderedDto**](AgentPromptRenderedDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentPrompt

> AgentPrompt UpdateAgentPrompt(ctx, orgId, promptId).Prompt(prompt).Execute()

Update an agent prompt



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
	prompt := *openapiclient.NewAgentPrompt("Id_example", "OwnerId_example", "OwnerOrganizationId_example", "Title_example") // AgentPrompt | Updated Agent Prompt object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentPromptsAPI.UpdateAgentPrompt(context.Background(), orgId, promptId).Prompt(prompt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentPromptsAPI.UpdateAgentPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgentPrompt`: AgentPrompt
	fmt.Fprintf(os.Stdout, "Response from `AgentPromptsAPI.UpdateAgentPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**promptId** | **string** | Prompt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **prompt** | [**AgentPrompt**](AgentPrompt.md) | Updated Agent Prompt object | 

### Return type

[**AgentPrompt**](AgentPrompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

