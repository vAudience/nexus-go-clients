# \HealthAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthGet**](HealthAPI.md#HealthGet) | **Get** /info/health | checks for service health



## HealthGet

> HealthResponse HealthGet(ctx).Execute()

checks for service health



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.HealthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.HealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthGet`: HealthResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.HealthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthGetRequest struct via the builder pattern


### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

