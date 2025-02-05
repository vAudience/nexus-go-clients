# \AgentTeamsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAgentToTeam**](AgentTeamsAPI.md#AddAgentToTeam) | **Patch** /v1/organizations/{org_id}/agent-teams/{team_id}/add-agent/{agent_id} | Add an agent to a team
[**CreateAgentTeam**](AgentTeamsAPI.md#CreateAgentTeam) | **Post** /v1/organizations/{org_id}/agent-teams | Create a new agent team
[**DeleteAgentTeam**](AgentTeamsAPI.md#DeleteAgentTeam) | **Delete** /v1/organizations/{org_id}/agent-teams/{team_id} | Delete an agent team
[**GetAgentTeam**](AgentTeamsAPI.md#GetAgentTeam) | **Get** /v1/organizations/{org_id}/agent-teams/{team_id} | Get an agent team
[**GetFullAgentTeam**](AgentTeamsAPI.md#GetFullAgentTeam) | **Get** /v1/organizations/{org_id}/agent-teams/{team_id}/full | Get a full agent team
[**ListAgentTeams**](AgentTeamsAPI.md#ListAgentTeams) | **Get** /v1/organizations/{org_id}/agent-teams | List agent teams
[**RemoveAgentFromTeam**](AgentTeamsAPI.md#RemoveAgentFromTeam) | **Patch** /v1/organizations/{org_id}/agent-teams/{team_id}/remove-agent/{agent_id} | Remove an agent from a team
[**UpdateAgentTeam**](AgentTeamsAPI.md#UpdateAgentTeam) | **Put** /v1/organizations/{org_id}/agent-teams/{team_id} | Update an agent team
[**UpdateSystemMessages**](AgentTeamsAPI.md#UpdateSystemMessages) | **Patch** /v1/organizations/{org_id}/agent-teams/{team_id}/messages/system | Update system messages
[**UpdateUserMessages**](AgentTeamsAPI.md#UpdateUserMessages) | **Patch** /v1/organizations/{org_id}/agent-teams/{team_id}/messages/user | Update user messages



## AddAgentToTeam

> AgentTeam AddAgentToTeam(ctx, orgId, teamId, agentId).Execute()

Add an agent to a team



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
	teamId := "teamId_example" // string | Agent Team ID
	agentId := "agentId_example" // string | Agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.AddAgentToTeam(context.Background(), orgId, teamId, agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.AddAgentToTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAgentToTeam`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.AddAgentToTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**teamId** | **string** | Agent Team ID | 
**agentId** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAgentToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgentTeam

> AgentTeam CreateAgentTeam(ctx, orgId).AgentTeam(agentTeam).Execute()

Create a new agent team



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
	agentTeam := *openapiclient.NewAgentTeamWriteDto() // AgentTeamWriteDto | Agent Team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.CreateAgentTeam(context.Background(), orgId).AgentTeam(agentTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.CreateAgentTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgentTeam`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.CreateAgentTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentTeam** | [**AgentTeamWriteDto**](AgentTeamWriteDto.md) | Agent Team | 

### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentTeam

> AgentTeam DeleteAgentTeam(ctx, orgId, teamId).Execute()

Delete an agent team



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
	teamId := "teamId_example" // string | Agent Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.DeleteAgentTeam(context.Background(), orgId, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.DeleteAgentTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgentTeam`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.DeleteAgentTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**teamId** | **string** | Agent Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentTeam

> AgentTeam GetAgentTeam(ctx, orgId, teamId).Execute()

Get an agent team



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
	teamId := "teamId_example" // string | Agent Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.GetAgentTeam(context.Background(), orgId, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.GetAgentTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentTeam`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.GetAgentTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**teamId** | **string** | Agent Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullAgentTeam

> AgentTeam GetFullAgentTeam(ctx, orgId, teamId).Execute()

Get a full agent team



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
	teamId := "teamId_example" // string | Agent Team ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.GetFullAgentTeam(context.Background(), orgId, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.GetFullAgentTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFullAgentTeam`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.GetFullAgentTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**teamId** | **string** | Agent Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFullAgentTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentTeams

> []AgentTeam ListAgentTeams(ctx, orgId).Execute()

List agent teams



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
	resp, r, err := apiClient.AgentTeamsAPI.ListAgentTeams(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.ListAgentTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentTeams`: []AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.ListAgentTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAgentFromTeam

> AgentTeam RemoveAgentFromTeam(ctx, orgId, teamId, agentId).Execute()

Remove an agent from a team



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
	teamId := "teamId_example" // string | Agent Team ID
	agentId := "agentId_example" // string | Agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.RemoveAgentFromTeam(context.Background(), orgId, teamId, agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.RemoveAgentFromTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAgentFromTeam`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.RemoveAgentFromTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**teamId** | **string** | Agent Team ID | 
**agentId** | **string** | Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAgentFromTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentTeam

> AgentTeam UpdateAgentTeam(ctx, orgId, teamId).AgentTeam(agentTeam).Execute()

Update an agent team



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
	teamId := "teamId_example" // string | Agent Team ID
	agentTeam := *openapiclient.NewAgentTeamWriteDto() // AgentTeamWriteDto | Agent Team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.UpdateAgentTeam(context.Background(), orgId, teamId).AgentTeam(agentTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.UpdateAgentTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgentTeam`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.UpdateAgentTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**teamId** | **string** | Agent Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agentTeam** | [**AgentTeamWriteDto**](AgentTeamWriteDto.md) | Agent Team | 

### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSystemMessages

> AgentTeam UpdateSystemMessages(ctx, orgId, teamId).Messages(messages).Execute()

Update system messages



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
	teamId := "teamId_example" // string | Agent Team ID
	messages := []string{"Property_example"} // []string | System Messages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.UpdateSystemMessages(context.Background(), orgId, teamId).Messages(messages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.UpdateSystemMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSystemMessages`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.UpdateSystemMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**teamId** | **string** | Agent Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **messages** | **[]string** | System Messages | 

### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserMessages

> AgentTeam UpdateUserMessages(ctx, orgId, teamId).Messages(messages).Execute()

Update user messages



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
	teamId := "teamId_example" // string | Agent Team ID
	messages := []string{"Property_example"} // []string | User Messages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTeamsAPI.UpdateUserMessages(context.Background(), orgId, teamId).Messages(messages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTeamsAPI.UpdateUserMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserMessages`: AgentTeam
	fmt.Fprintf(os.Stdout, "Response from `AgentTeamsAPI.UpdateUserMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**teamId** | **string** | Agent Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **messages** | **[]string** | User Messages | 

### Return type

[**AgentTeam**](AgentTeam.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

