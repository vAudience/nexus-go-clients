# \ExecutionLogsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChatCompletionCosts**](ExecutionLogsAPI.md#CreateChatCompletionCosts) | **Post** /v1/organizations/{org_id}/execution-logs/chat-completion-costs | Track chat completion costs
[**CreateDocumentConversionCosts**](ExecutionLogsAPI.md#CreateDocumentConversionCosts) | **Post** /v1/organizations/{org_id}/execution-logs/document-conversion-costs | Track document conversion costs
[**CreateFileUploadCosts**](ExecutionLogsAPI.md#CreateFileUploadCosts) | **Post** /v1/organizations/{org_id}/execution-logs/file-upload-costs | Track file upload costs
[**GetExecutionLogsCosts**](ExecutionLogsAPI.md#GetExecutionLogsCosts) | **Get** /v1/organizations/{org_id}/execution-logs/costs | Get execution logs costs
[**SearchExecutionLogs**](ExecutionLogsAPI.md#SearchExecutionLogs) | **Get** /v1/organizations/{org_id}/execution-logs/search | Search execution logs



## CreateChatCompletionCosts

> CostTrackingResponse CreateChatCompletionCosts(ctx, orgId).Request(request).Execute()

Track chat completion costs



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
	request := *openapiclient.NewChatCompletionCostTrackingRequest("ModelInternalId_example", "ServiceInternalId_example") // ChatCompletionCostTrackingRequest | Chat completion cost tracking request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionLogsAPI.CreateChatCompletionCosts(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionLogsAPI.CreateChatCompletionCosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChatCompletionCosts`: CostTrackingResponse
	fmt.Fprintf(os.Stdout, "Response from `ExecutionLogsAPI.CreateChatCompletionCosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatCompletionCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ChatCompletionCostTrackingRequest**](ChatCompletionCostTrackingRequest.md) | Chat completion cost tracking request | 

### Return type

[**CostTrackingResponse**](CostTrackingResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocumentConversionCosts

> CostTrackingResponse CreateDocumentConversionCosts(ctx, orgId).Request(request).Execute()

Track document conversion costs



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
	request := *openapiclient.NewDocumentConversionCostTrackingRequest() // DocumentConversionCostTrackingRequest | Document conversion cost tracking request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionLogsAPI.CreateDocumentConversionCosts(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionLogsAPI.CreateDocumentConversionCosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDocumentConversionCosts`: CostTrackingResponse
	fmt.Fprintf(os.Stdout, "Response from `ExecutionLogsAPI.CreateDocumentConversionCosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentConversionCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DocumentConversionCostTrackingRequest**](DocumentConversionCostTrackingRequest.md) | Document conversion cost tracking request | 

### Return type

[**CostTrackingResponse**](CostTrackingResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFileUploadCosts

> CostTrackingResponse CreateFileUploadCosts(ctx, orgId).Request(request).Execute()

Track file upload costs



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
	request := *openapiclient.NewFileUploadCostTrackingRequest(int32(123)) // FileUploadCostTrackingRequest | File upload cost tracking request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionLogsAPI.CreateFileUploadCosts(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionLogsAPI.CreateFileUploadCosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileUploadCosts`: CostTrackingResponse
	fmt.Fprintf(os.Stdout, "Response from `ExecutionLogsAPI.CreateFileUploadCosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileUploadCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**FileUploadCostTrackingRequest**](FileUploadCostTrackingRequest.md) | File upload cost tracking request | 

### Return type

[**CostTrackingResponse**](CostTrackingResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	userId := "userId_example" // string | User ID or me (optional)
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

 **userId** | **string** | User ID or me | 
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

> ExecutionLogResults SearchExecutionLogs(ctx, orgId).UserId(userId).ExecType(execType).RequestId(requestId).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

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
	userId := "userId_example" // string | User ID or me (optional)
	execType := "execType_example" // string | Execution log type (optional)
	requestId := "requestId_example" // string | Request ID (optional)
	startDate := "startDate_example" // string | Start date in Unix milliseconds (optional)
	endDate := "endDate_example" // string | End date in Unix milliseconds (optional)
	offset := int32(56) // int32 | Offset (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionLogsAPI.SearchExecutionLogs(context.Background(), orgId).UserId(userId).ExecType(execType).RequestId(requestId).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
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

 **userId** | **string** | User ID or me | 
 **execType** | **string** | Execution log type | 
 **requestId** | **string** | Request ID | 
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

