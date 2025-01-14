# \AgentAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAssignedTool**](AgentAPI.md#AddAssignedTool) | **Patch** /organizations/{org_id}/agents/{id}/addtool/{tool_id} | Add an assigned tool to an agent
[**AddAttachedFileID**](AgentAPI.md#AddAttachedFileID) | **Patch** /organizations/{org_id}/agents/{id}/addfile/{file_id} | Add an attached file ID to an agent
[**AddInitialUserMessage**](AgentAPI.md#AddInitialUserMessage) | **Patch** /organizations/{org_id}/agents/{id}/messages/addusermessage | Add an initial user message to an agent
[**AddSystemMessage**](AgentAPI.md#AddSystemMessage) | **Patch** /organizations/{org_id}/agents/{id}/messages/addsystemmessage | Add a system message to an agent
[**CreateAgent**](AgentAPI.md#CreateAgent) | **Post** /organizations/{org_id}/agents | Create a new agent
[**DeleteAgent**](AgentAPI.md#DeleteAgent) | **Delete** /organizations/{org_id}/agents/{id} | Delete an agent
[**GetAgent**](AgentAPI.md#GetAgent) | **Get** /organizations/{org_id}/agents/{id} | Get an agent
[**ListAgents**](AgentAPI.md#ListAgents) | **Get** /organizations/{org_id}/agents | List agents
[**RemoveAssignedTool**](AgentAPI.md#RemoveAssignedTool) | **Patch** /organizations/{org_id}/agents/{id}/removetool/{tool_id} | Remove an assigned tool from an agent
[**RemoveAttachedFileID**](AgentAPI.md#RemoveAttachedFileID) | **Patch** /organizations/{org_id}/agents/{id}/removefile/{file_id} | Remove an attached file ID from an agent
[**RemoveInitialUserMessage**](AgentAPI.md#RemoveInitialUserMessage) | **Patch** /organizations/{org_id}/agents/{id}/messages/removeusermessage | Remove an initial user message from an agent
[**RemoveSystemMessage**](AgentAPI.md#RemoveSystemMessage) | **Patch** /organizations/{org_id}/agents/{id}/messages/removesystemmessage | Remove a system message from an agent
[**SearchAgents**](AgentAPI.md#SearchAgents) | **Get** /organizations/{org_id}/agents/search | Search agents
[**UpdateAgent**](AgentAPI.md#UpdateAgent) | **Put** /organizations/{org_id}/agents/{id} | Update an agent



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
	resp, r, err := apiClient.AgentAPI.AddAssignedTool(context.Background(), orgId, id, toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AddAssignedTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAssignedTool`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.AddAssignedTool`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.AddAttachedFileID(context.Background(), orgId, id, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AddAttachedFileID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAttachedFileID`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.AddAttachedFileID`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.AddInitialUserMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AddInitialUserMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddInitialUserMessage`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.AddInitialUserMessage`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.AddSystemMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AddSystemMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSystemMessage`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.AddSystemMessage`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.CreateAgent(context.Background(), orgId).Agent(agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CreateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CreateAgent`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.DeleteAgent(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.DeleteAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.DeleteAgent`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.GetAgent(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgent`: %v\n", resp)
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

> []Agent ListAgents(ctx, orgId).AddDefaultAgents(addDefaultAgents).Capability(capability).Execute()

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
	addDefaultAgents := "addDefaultAgents_example" // string | Include default agents to the list of org owned agents (optional)
	capability := "capability_example" // string | Filter by agent capability (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.ListAgents(context.Background(), orgId).AddDefaultAgents(addDefaultAgents).Capability(capability).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.ListAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgents`: []Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.ListAgents`: %v\n", resp)
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

 **addDefaultAgents** | **string** | Include default agents to the list of org owned agents | 
 **capability** | **string** | Filter by agent capability | 

### Return type

[**[]Agent**](Agent.md)

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
	resp, r, err := apiClient.AgentAPI.RemoveAssignedTool(context.Background(), orgId, id, toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.RemoveAssignedTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAssignedTool`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.RemoveAssignedTool`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.RemoveAttachedFileID(context.Background(), orgId, id, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.RemoveAttachedFileID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAttachedFileID`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.RemoveAttachedFileID`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.RemoveInitialUserMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.RemoveInitialUserMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveInitialUserMessage`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.RemoveInitialUserMessage`: %v\n", resp)
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
	resp, r, err := apiClient.AgentAPI.RemoveSystemMessage(context.Background(), orgId, id).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.RemoveSystemMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSystemMessage`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.RemoveSystemMessage`: %v\n", resp)
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


## SearchAgents

> []Agent SearchAgents(ctx, orgId).Name(name).ModelID(modelID).Execute()

Search agents



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
	name := "name_example" // string | Agent Name (optional)
	modelID := "modelID_example" // string | Model ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.SearchAgents(context.Background(), orgId).Name(name).ModelID(modelID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.SearchAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAgents`: []Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.SearchAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Agent Name | 
 **modelID** | **string** | Model ID | 

### Return type

[**[]Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.AgentAPI.UpdateAgent(context.Background(), orgId, id).Agent(agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.UpdateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.UpdateAgent`: %v\n", resp)
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

