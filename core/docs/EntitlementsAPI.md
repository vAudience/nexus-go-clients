# \EntitlementsAPI

All URIs are relative to *https://core.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyEntitlements**](EntitlementsAPI.md#GetMyEntitlements) | **Get** /v1/entitlements/me | Get my entitlements



## GetMyEntitlements

> map[string]EntitlementResponse GetMyEntitlements(ctx).OrganizationId(organizationId).Execute()

Get my entitlements



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/core"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetMyEntitlements(context.Background()).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetMyEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyEntitlements`: map[string]EntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetMyEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | Organization ID | 

### Return type

[**map[string]EntitlementResponse**](EntitlementResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

