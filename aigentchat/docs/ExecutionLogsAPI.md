# \ExecutionLogsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecutionLogsCosts**](ExecutionLogsAPI.md#GetExecutionLogsCosts) | **Get** /v1/organizations/{org_id}/execution-logs/costs | Get execution logs costs
[**SearchExecutionLogs**](ExecutionLogsAPI.md#SearchExecutionLogs) | **Get** /v1/organizations/{org_id}/execution-logs/search | Search execution logs



## GetExecutionLogsCosts

> ExecutionLogCosts GetExecutionLogsCosts(ctx, orgId).UserId(userId).StartDate(startDate).EndDate(endDate).Execute()

Get execution logs costs



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
	userId := "userId_example" // string | User ID (optional)
	startDate := "startDate_example" // string | Start date in Unix milliseconds (optional)
	endDate := "endDate_example" // string | End date in Unix milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionLogsAPI.GetExecutionLogsCosts(context.Background(), orgId).UserId(userId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionLogsAPI.GetExecutionLogsCosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecutionLogsCosts`: ExecutionLogCosts
	fmt.Fprintf(os.Stdout, "Response from `ExecutionLogsAPI.GetExecutionLogsCosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionLogsCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User ID | 
 **startDate** | **string** | Start date in Unix milliseconds | 
 **endDate** | **string** | End date in Unix milliseconds | 

### Return type

[**ExecutionLogCosts**](ExecutionLogCosts.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchExecutionLogs

> ExecutionLogResults SearchExecutionLogs(ctx, orgId).UserId(userId).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

Search execution logs



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
	userId := "userId_example" // string | User ID (optional)
	startDate := "startDate_example" // string | Start date in Unix milliseconds (optional)
	endDate := "endDate_example" // string | End date in Unix milliseconds (optional)
	offset := int32(56) // int32 | Offset (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionLogsAPI.SearchExecutionLogs(context.Background(), orgId).UserId(userId).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionLogsAPI.SearchExecutionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchExecutionLogs`: ExecutionLogResults
	fmt.Fprintf(os.Stdout, "Response from `ExecutionLogsAPI.SearchExecutionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchExecutionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User ID | 
 **startDate** | **string** | Start date in Unix milliseconds | 
 **endDate** | **string** | End date in Unix milliseconds | 
 **offset** | **int32** | Offset | 
 **limit** | **int32** | Limit | 

### Return type

[**ExecutionLogResults**](ExecutionLogResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

