# \AdminAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransferResources**](AdminAPI.md#TransferResources) | **Post** /v1/organizations/{org_id}/admin/resource-transfers | Transfer a user&#39;s resources between organizations



## TransferResources

> ResourceTransferResult TransferResources(ctx, orgId).Body(body).Execute()

Transfer a user's resources between organizations



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
	orgId := "orgId_example" // string | source organization ID
	body := *openapiclient.NewResourceTransferRequest([]openapiclient.TransferableResourceType{openapiclient.TransferableResourceType("agents")}, "SourceOwnerId_example", "TargetOrganizationId_example") // ResourceTransferRequest | Transfer request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.TransferResources(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.TransferResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferResources`: ResourceTransferResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.TransferResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | source organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ResourceTransferRequest**](ResourceTransferRequest.md) | Transfer request | 

### Return type

[**ResourceTransferResult**](ResourceTransferResult.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

