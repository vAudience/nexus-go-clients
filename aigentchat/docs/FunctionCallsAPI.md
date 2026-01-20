# \FunctionCallsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteFunctionCall**](FunctionCallsAPI.md#ExecuteFunctionCall) | **Post** /v1/organizations/{org_id}/function-calls/{ref} | Execute a function call
[**GetFunctionCall**](FunctionCallsAPI.md#GetFunctionCall) | **Get** /v1/organizations/{org_id}/function-calls/{ref} | Get function call definition
[**GetFunctionCalls**](FunctionCallsAPI.md#GetFunctionCalls) | **Get** /v1/organizations/{org_id}/function-calls | List accessible function calls



## ExecuteFunctionCall

> FunctionCallResults ExecuteFunctionCall(ctx, orgId, ref).Request(request).Execute()

Execute a function call



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
	ref := "ref_example" // string | Function call name or ID
	request := *openapiclient.NewFunctionCallRequestDto() // FunctionCallRequestDto | Function call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionCallsAPI.ExecuteFunctionCall(context.Background(), orgId, ref).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionCallsAPI.ExecuteFunctionCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteFunctionCall`: FunctionCallResults
	fmt.Fprintf(os.Stdout, "Response from `FunctionCallsAPI.ExecuteFunctionCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**ref** | **string** | Function call name or ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteFunctionCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**FunctionCallRequestDto**](FunctionCallRequestDto.md) | Function call request | 

### Return type

[**FunctionCallResults**](FunctionCallResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionCall

> FunctionCall GetFunctionCall(ctx, orgId, ref).Execute()

Get function call definition



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
	ref := "ref_example" // string | Function call name or ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionCallsAPI.GetFunctionCall(context.Background(), orgId, ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionCallsAPI.GetFunctionCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionCall`: FunctionCall
	fmt.Fprintf(os.Stdout, "Response from `FunctionCallsAPI.GetFunctionCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**ref** | **string** | Function call name or ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FunctionCall**](FunctionCall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionCalls

> []FunctionCall GetFunctionCalls(ctx, orgId).Execute()

List accessible function calls



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
	resp, r, err := apiClient.FunctionCallsAPI.GetFunctionCalls(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionCallsAPI.GetFunctionCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionCalls`: []FunctionCall
	fmt.Fprintf(os.Stdout, "Response from `FunctionCallsAPI.GetFunctionCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FunctionCall**](FunctionCall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

