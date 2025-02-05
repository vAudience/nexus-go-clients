# \OrgCostBudgetsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNewOrgCostBudget**](OrgCostBudgetsAPI.md#AddNewOrgCostBudget) | **Post** /v1/organizations/{org_id}/budgets | Add new organization cost budget
[**CheckOrgCostBudget**](OrgCostBudgetsAPI.md#CheckOrgCostBudget) | **Get** /v1/organizations/{org_id}/budgets/check | Check organization cost budget
[**GetOrgCostBudget**](OrgCostBudgetsAPI.md#GetOrgCostBudget) | **Get** /v1/organizations/{org_id}/budgets | Get organization cost budget
[**UpdateOrgCostBudget**](OrgCostBudgetsAPI.md#UpdateOrgCostBudget) | **Put** /v1/organizations/{org_id}/budgets | Update organization cost budget



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
	resp, r, err := apiClient.OrgCostBudgetsAPI.AddNewOrgCostBudget(context.Background(), orgId).Budget(budget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCostBudgetsAPI.AddNewOrgCostBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNewOrgCostBudget`: OrgCostBudget
	fmt.Fprintf(os.Stdout, "Response from `OrgCostBudgetsAPI.AddNewOrgCostBudget`: %v\n", resp)
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
	resp, r, err := apiClient.OrgCostBudgetsAPI.CheckOrgCostBudget(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCostBudgetsAPI.CheckOrgCostBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckOrgCostBudget`: OrgCostBudgetCheck
	fmt.Fprintf(os.Stdout, "Response from `OrgCostBudgetsAPI.CheckOrgCostBudget`: %v\n", resp)
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
	resp, r, err := apiClient.OrgCostBudgetsAPI.GetOrgCostBudget(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCostBudgetsAPI.GetOrgCostBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCostBudget`: OrgCostBudget
	fmt.Fprintf(os.Stdout, "Response from `OrgCostBudgetsAPI.GetOrgCostBudget`: %v\n", resp)
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
	resp, r, err := apiClient.OrgCostBudgetsAPI.UpdateOrgCostBudget(context.Background(), orgId).Budget(budget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCostBudgetsAPI.UpdateOrgCostBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgCostBudget`: OrgCostBudget
	fmt.Fprintf(os.Stdout, "Response from `OrgCostBudgetsAPI.UpdateOrgCostBudget`: %v\n", resp)
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

