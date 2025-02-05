# \MissionsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelMission**](MissionsAPI.md#CancelMission) | **Post** /v1/organizations/{org_id}/missions/{mission_id}/cancel | Cancel a mission
[**CreateMission**](MissionsAPI.md#CreateMission) | **Post** /v1/organizations/{org_id}/missions/{mission_executor_id} | Create a new mission
[**DeleteMission**](MissionsAPI.md#DeleteMission) | **Delete** /v1/organizations/{org_id}/missions/{mission_id} | Delete a mission
[**GetMission**](MissionsAPI.md#GetMission) | **Get** /v1/organizations/{org_id}/missions/{mission_id} | Get a mission
[**ListMissionsByExecutorID**](MissionsAPI.md#ListMissionsByExecutorID) | **Get** /v1/organizations/{org_id}/missions/executor/{mission_executor_id} | List all missions by a executorID (Agent Team or Agent)
[**ListMissionsByOrg**](MissionsAPI.md#ListMissionsByOrg) | **Get** /v1/organizations/{org_id}/missions | List all missions of an organization and owned by the current user



## CancelMission

> CancelMission(ctx, orgId, missionId).Execute()

Cancel a mission



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
	missionId := "missionId_example" // string | Mission ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MissionsAPI.CancelMission(context.Background(), orgId, missionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.CancelMission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**missionId** | **string** | Mission ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelMissionRequest struct via the builder pattern


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


## CreateMission

> Mission CreateMission(ctx, orgId, missionExecutorId).Mission(mission).Execute()

Create a new mission



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
	missionExecutorId := "missionExecutorId_example" // string | Agent Team ID OR Agent ID to run the mission with
	mission := *openapiclient.NewMissionWriteDto(*openapiclient.NewMissionInstructions("Text_example")) // MissionWriteDto | Mission object that needs to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.CreateMission(context.Background(), orgId, missionExecutorId).Mission(mission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.CreateMission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMission`: Mission
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.CreateMission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**missionExecutorId** | **string** | Agent Team ID OR Agent ID to run the mission with | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mission** | [**MissionWriteDto**](MissionWriteDto.md) | Mission object that needs to be created | 

### Return type

[**Mission**](Mission.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMission

> DeleteMission(ctx, orgId, missionId).Execute()

Delete a mission



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
	missionId := "missionId_example" // string | Mission ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MissionsAPI.DeleteMission(context.Background(), orgId, missionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.DeleteMission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**missionId** | **string** | Mission ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMissionRequest struct via the builder pattern


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


## GetMission

> Mission GetMission(ctx, orgId, missionId).Execute()

Get a mission



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
	missionId := "missionId_example" // string | Mission ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.GetMission(context.Background(), orgId, missionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.GetMission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMission`: Mission
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.GetMission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**missionId** | **string** | Mission ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Mission**](Mission.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMissionsByExecutorID

> []Mission ListMissionsByExecutorID(ctx, orgId, missionExecutorId).Offset(offset).Limit(limit).Execute()

List all missions by a executorID (Agent Team or Agent)



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
	missionExecutorId := "missionExecutorId_example" // string | Agent or Agent-Team ID
	offset := int32(56) // int32 | Offset the number of missions returned (optional)
	limit := int32(56) // int32 | Limit the number of missions returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.ListMissionsByExecutorID(context.Background(), orgId, missionExecutorId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.ListMissionsByExecutorID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMissionsByExecutorID`: []Mission
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.ListMissionsByExecutorID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**missionExecutorId** | **string** | Agent or Agent-Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMissionsByExecutorIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | Offset the number of missions returned | 
 **limit** | **int32** | Limit the number of missions returned | 

### Return type

[**[]Mission**](Mission.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMissionsByOrg

> []Mission ListMissionsByOrg(ctx, orgId).Offset(offset).Limit(limit).Execute()

List all missions of an organization and owned by the current user



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
	offset := int32(56) // int32 | Offset the number of missions returned (optional)
	limit := int32(56) // int32 | Limit the number of missions returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.ListMissionsByOrg(context.Background(), orgId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.ListMissionsByOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMissionsByOrg`: []Mission
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.ListMissionsByOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMissionsByOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Offset the number of missions returned | 
 **limit** | **int32** | Limit the number of missions returned | 

### Return type

[**[]Mission**](Mission.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

