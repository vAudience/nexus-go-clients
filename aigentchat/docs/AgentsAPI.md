# \AgentsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAssignedTool**](AgentsAPI.md#AddAssignedTool) | **Patch** /v1/organizations/{org_id}/agents/{id}/add-tool/{tool_id} | Add an assigned tool to an agent
[**AddAttachedFileID**](AgentsAPI.md#AddAttachedFileID) | **Patch** /v1/organizations/{org_id}/agents/{id}/add-file/{file_id} | Add an attached file ID to an agent
[**AddInitialUserMessage**](AgentsAPI.md#AddInitialUserMessage) | **Patch** /v1/organizations/{org_id}/agents/{id}/messages/add-user-message | Add an initial user message to an agent
[**AddSystemMessage**](AgentsAPI.md#AddSystemMessage) | **Patch** /v1/organizations/{org_id}/agents/{id}/messages/add-system-message | Add a system message to an agent
[**CreateAgent**](AgentsAPI.md#CreateAgent) | **Post** /v1/organizations/{org_id}/agents | Create a new agent
[**DeleteAgent**](AgentsAPI.md#DeleteAgent) | **Delete** /v1/organizations/{org_id}/agents/{id} | Delete an agent
[**GetAgent**](AgentsAPI.md#GetAgent) | **Get** /v1/organizations/{org_id}/agents/{id} | Get an agent
[**ListAgents**](AgentsAPI.md#ListAgents) | **Get** /v1/organizations/{org_id}/agents | List agents
[**RemoveAssignedTool**](AgentsAPI.md#RemoveAssignedTool) | **Patch** /v1/organizations/{org_id}/agents/{id}/remove-tool/{tool_id} | Remove an assigned tool from an agent
[**RemoveAttachedFileID**](AgentsAPI.md#RemoveAttachedFileID) | **Patch** /v1/organizations/{org_id}/agents/{id}/remove-file/{file_id} | Remove an attached file ID from an agent
[**RemoveInitialUserMessage**](AgentsAPI.md#RemoveInitialUserMessage) | **Patch** /v1/organizations/{org_id}/agents/{id}/messages/remove-user-message | Remove an initial user message from an agent
[**RemoveSystemMessage**](AgentsAPI.md#RemoveSystemMessage) | **Patch** /v1/organizations/{org_id}/agents/{id}/messages/remove-system-message | Remove a system message from an agent
[**UpdateAgent**](AgentsAPI.md#UpdateAgent) | **Put** /v1/organizations/{org_id}/agents/{id} | Update an agent



## AddAssignedTool

> Agent AddAssignedTool(ctx, orgId, id, toolId).Execute()

Add an assigned tool to an agent



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
	id := "id_example" // string | Agent ID
	toolId := "toolId_example" // string | Tool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AddAssignedTool(context.Background(), orgId, id, toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AddAssignedTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAssignedTool`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AddAssignedTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 
**toolId** | **string** | Tool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAssignedToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAttachedFileID

> Agent AddAttachedFileID(ctx, orgId, id, fileId).Execute()

Add an attached file ID to an agent



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
	id := "id_example" // string | Agent ID
	fileId := "fileId_example" // string | File ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AddAttachedFileID(context.Background(), orgId, id, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AddAttachedFileID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAttachedFileID`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AddAttachedFileID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 
**fileId** | **string** | File ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAttachedFileIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddInitialUserMessage

> Agent AddInitialUserMessage(ctx, orgId, id).Message(message).Execute()

Add an initial user message to an agent



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
	id := "id_example" // string | Agent ID
	message := "message_example" // string | Initial User Message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AddInitialUserMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AddInitialUserMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddInitialUserMessage`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AddInitialUserMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInitialUserMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | **string** | Initial User Message | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSystemMessage

> Agent AddSystemMessage(ctx, orgId, id).Message(message).Execute()

Add a system message to an agent



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
	id := "id_example" // string | Agent ID
	message := "message_example" // string | System Message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AddSystemMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AddSystemMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSystemMessage`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AddSystemMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSystemMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | **string** | System Message | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgent

> Agent CreateAgent(ctx, orgId).Agent(agent).Execute()

Create a new agent



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
	agent := *openapiclient.NewAgentWriteDto() // AgentWriteDto | Agent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CreateAgent(context.Background(), orgId).Agent(agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CreateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CreateAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agent** | [**AgentWriteDto**](AgentWriteDto.md) | Agent | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgent

> Agent DeleteAgent(ctx, orgId, id).Execute()

Delete an agent



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
	id := "id_example" // string | Agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.DeleteAgent(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.DeleteAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.DeleteAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgent

> Agent GetAgent(ctx, orgId, id).Execute()

Get an agent



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
	id := "id_example" // string | Agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgent(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgents

> AgentResults ListAgents(ctx, orgId).ModelIds(modelIds).TagIds(tagIds).Q(q).Action(action).Types(types).AddPredefinedAgents(addPredefinedAgents).AdminMode(adminMode).Visibility(visibility).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()

List agents



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
	modelIds := []string{"Inner_example"} // []string | Model ID to filter by (comma separated) (optional)
	tagIds := []string{"Inner_example"} // []string | Tag IDs to filter by (comma separated) (optional)
	q := "q_example" // string | Search term for name or description (optional)
	action := "action_example" // string | Filter agents by model action (chat, image, etc.) (optional)
	types := []string{"Types_example"} // []string | Filter agents by types (basic, background, service - comma separated) (optional)
	addPredefinedAgents := true // bool | Include default agents to the list of org owned agents (optional)
	adminMode := true // bool | Admin mode to bypass certain permission checks (optional)
	visibility := "visibility_example" // string | Filter agents by access visibility (public, private, organization) (optional)
	limit := int32(56) // int32 | Limit the number of results (optional) (default to 1000)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)
	sortBy := "sortBy_example" // string | Field to sort by (name, createdat, updatedat) (optional) (default to "\"name\"")
	sortOrder := "sortOrder_example" // string | Sort order (asc or desc) (optional) (default to "\"asc\"")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ListAgents(context.Background(), orgId).ModelIds(modelIds).TagIds(tagIds).Q(q).Action(action).Types(types).AddPredefinedAgents(addPredefinedAgents).AdminMode(adminMode).Visibility(visibility).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ListAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgents`: AgentResults
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ListAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelIds** | **[]string** | Model ID to filter by (comma separated) | 
 **tagIds** | **[]string** | Tag IDs to filter by (comma separated) | 
 **q** | **string** | Search term for name or description | 
 **action** | **string** | Filter agents by model action (chat, image, etc.) | 
 **types** | **[]string** | Filter agents by types (basic, background, service - comma separated) | 
 **addPredefinedAgents** | **bool** | Include default agents to the list of org owned agents | 
 **adminMode** | **bool** | Admin mode to bypass certain permission checks | 
 **visibility** | **string** | Filter agents by access visibility (public, private, organization) | 
 **limit** | **int32** | Limit the number of results | [default to 1000]
 **offset** | **int32** | Offset for pagination | [default to 0]
 **sortBy** | **string** | Field to sort by (name, createdat, updatedat) | [default to &quot;\&quot;name\&quot;&quot;]
 **sortOrder** | **string** | Sort order (asc or desc) | [default to &quot;\&quot;asc\&quot;&quot;]

### Return type

[**AgentResults**](AgentResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAssignedTool

> Agent RemoveAssignedTool(ctx, orgId, id, toolId).Execute()

Remove an assigned tool from an agent



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
	id := "id_example" // string | Agent ID
	toolId := "toolId_example" // string | Tool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.RemoveAssignedTool(context.Background(), orgId, id, toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.RemoveAssignedTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAssignedTool`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.RemoveAssignedTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 
**toolId** | **string** | Tool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAssignedToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAttachedFileID

> Agent RemoveAttachedFileID(ctx, orgId, id, fileId).Execute()

Remove an attached file ID from an agent



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
	id := "id_example" // string | Agent ID
	fileId := "fileId_example" // string | File ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.RemoveAttachedFileID(context.Background(), orgId, id, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.RemoveAttachedFileID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAttachedFileID`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.RemoveAttachedFileID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 
**fileId** | **string** | File ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAttachedFileIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInitialUserMessage

> Agent RemoveInitialUserMessage(ctx, orgId, id).Message(message).Execute()

Remove an initial user message from an agent



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
	id := "id_example" // string | Agent ID
	message := "message_example" // string | Initial User Message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.RemoveInitialUserMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.RemoveInitialUserMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveInitialUserMessage`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.RemoveInitialUserMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInitialUserMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | **string** | Initial User Message | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSystemMessage

> Agent RemoveSystemMessage(ctx, orgId, id).Message(message).Execute()

Remove a system message from an agent



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
	id := "id_example" // string | Agent ID
	message := "message_example" // string | System Message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.RemoveSystemMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.RemoveSystemMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSystemMessage`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.RemoveSystemMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSystemMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | **string** | System Message | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgent

> Agent UpdateAgent(ctx, orgId, id).Agent(agent).Execute()

Update an agent



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
	id := "id_example" // string | Agent ID
	agent := *openapiclient.NewAgentWriteDto() // AgentWriteDto | Agent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.UpdateAgent(context.Background(), orgId, id).Agent(agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.UpdateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.UpdateAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**id** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agent** | [**AgentWriteDto**](AgentWriteDto.md) | Agent | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

