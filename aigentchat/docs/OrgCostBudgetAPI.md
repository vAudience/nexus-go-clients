# \OrgCostBudgetAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNewOrgCostBudget**](OrgCostBudgetAPI.md#AddNewOrgCostBudget) | **Post** /organizations/{org_id}/budget | Add new organization cost budget
[**CheckOrgCostBudget**](OrgCostBudgetAPI.md#CheckOrgCostBudget) | **Get** /organizations/{org_id}/budget/check | Check organization cost budget
[**GetOrgCostBudget**](OrgCostBudgetAPI.md#GetOrgCostBudget) | **Get** /organizations/{org_id}/budget | Get organization cost budget
[**UpdateOrgCostBudget**](OrgCostBudgetAPI.md#UpdateOrgCostBudget) | **Put** /organizations/{org_id}/budget | Update organization cost budget



## AddNewOrgCostBudget

> OrgCostBudget AddNewOrgCostBudget(ctx, orgId).Budget(budget).Execute()

Add new organization cost budget



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
	budget := *openapiclient.NewOrgCostBudgetWriteDto() // OrgCostBudgetWriteDto | New budget

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgCostBudgetAPI.AddNewOrgCostBudget(context.Background(), orgId).Budget(budget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCostBudgetAPI.AddNewOrgCostBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNewOrgCostBudget`: OrgCostBudget
	fmt.Fprintf(os.Stdout, "Response from `OrgCostBudgetAPI.AddNewOrgCostBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNewOrgCostBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budget** | [**OrgCostBudgetWriteDto**](OrgCostBudgetWriteDto.md) | New budget | 

### Return type

[**OrgCostBudget**](OrgCostBudget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckOrgCostBudget

> OrgCostBudgetCheck CheckOrgCostBudget(ctx, orgId).Execute()

Check organization cost budget



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgCostBudgetAPI.CheckOrgCostBudget(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCostBudgetAPI.CheckOrgCostBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckOrgCostBudget`: OrgCostBudgetCheck
	fmt.Fprintf(os.Stdout, "Response from `OrgCostBudgetAPI.CheckOrgCostBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckOrgCostBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgCostBudgetCheck**](OrgCostBudgetCheck.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgCostBudget

> OrgCostBudget GetOrgCostBudget(ctx, orgId).Execute()

Get organization cost budget



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgCostBudgetAPI.GetOrgCostBudget(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCostBudgetAPI.GetOrgCostBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCostBudget`: OrgCostBudget
	fmt.Fprintf(os.Stdout, "Response from `OrgCostBudgetAPI.GetOrgCostBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgCostBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgCostBudget**](OrgCostBudget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgCostBudget

> OrgCostBudget UpdateOrgCostBudget(ctx, orgId).Budget(budget).Execute()

Update organization cost budget



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
	budget := *openapiclient.NewOrgCostBudgetWriteDto() // OrgCostBudgetWriteDto | Updated budget

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgCostBudgetAPI.UpdateOrgCostBudget(context.Background(), orgId).Budget(budget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCostBudgetAPI.UpdateOrgCostBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgCostBudget`: OrgCostBudget
	fmt.Fprintf(os.Stdout, "Response from `OrgCostBudgetAPI.UpdateOrgCostBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgCostBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budget** | [**OrgCostBudgetWriteDto**](OrgCostBudgetWriteDto.md) | Updated budget | 

### Return type

[**OrgCostBudget**](OrgCostBudget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

