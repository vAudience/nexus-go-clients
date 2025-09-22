# \AuditTrailsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditTrailsByChannelId**](AuditTrailsAPI.md#GetAuditTrailsByChannelId) | **Get** /v1/organizations/{org_id}/audit-trails/channels/{channel_id} | Get the audit trails by channel ID



## GetAuditTrailsByChannelId

> AuditTrail GetAuditTrailsByChannelId(ctx, orgId, channelId).Execute()

Get the audit trails by channel ID



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
	channelId := "channelId_example" // string | channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditTrailsAPI.GetAuditTrailsByChannelId(context.Background(), orgId, channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditTrailsAPI.GetAuditTrailsByChannelId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditTrailsByChannelId`: AuditTrail
	fmt.Fprintf(os.Stdout, "Response from `AuditTrailsAPI.GetAuditTrailsByChannelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**channelId** | **string** | channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditTrailsByChannelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuditTrail**](AuditTrail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

