# \FunctionCallsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteFunctionCall**](FunctionCallsAPI.md#ExecuteFunctionCall) | **Post** /organizations/{org_id}/functioncalls/{name} | Execute a function call
[**GetFunctionCall**](FunctionCallsAPI.md#GetFunctionCall) | **Get** /organizations/{org_id}/functioncalls/{name} | Get function call definition
[**GetFunctionCalls**](FunctionCallsAPI.md#GetFunctionCalls) | **Get** /organizations/{org_id}/functioncalls | List accessible function calls



## ExecuteFunctionCall

> AdapterExecutionResults ExecuteFunctionCall(ctx, orgId, name).Arguments(arguments).Execute()

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
	name := "name_example" // string | Function call name
	arguments := map[string]interface{}{ ... } // map[string]interface{} | Function call arguments

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionCallsAPI.ExecuteFunctionCall(context.Background(), orgId, name).Arguments(arguments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionCallsAPI.ExecuteFunctionCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteFunctionCall`: AdapterExecutionResults
	fmt.Fprintf(os.Stdout, "Response from `FunctionCallsAPI.ExecuteFunctionCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**name** | **string** | Function call name | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteFunctionCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **arguments** | **map[string]interface{}** | Function call arguments | 

### Return type

[**AdapterExecutionResults**](AdapterExecutionResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionCall

> FunctionCall GetFunctionCall(ctx, orgId, name).Execute()

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
	name := "name_example" // string | Function call name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionCallsAPI.GetFunctionCall(context.Background(), orgId, name).Execute()
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
**name** | **string** | Function call name | 

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

