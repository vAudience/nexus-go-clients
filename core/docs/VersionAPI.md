# \VersionAPI

All URIs are relative to *https://core.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionGet**](VersionAPI.md#VersionGet) | **Get** /info/version | returns the service version



## VersionGet

> VersionResponse VersionGet(ctx).Execute()

returns the service version



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.VersionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.VersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionGet`: VersionResponse
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.VersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionGetRequest struct via the builder pattern


### Return type

[**VersionResponse**](VersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

