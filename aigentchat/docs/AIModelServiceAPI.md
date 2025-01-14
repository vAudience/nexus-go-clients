# \AIModelServiceAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAIModelService**](AIModelServiceAPI.md#CreateAIModelService) | **Post** /ai_model_services | Create a new AI model service
[**DeleteAIModelService**](AIModelServiceAPI.md#DeleteAIModelService) | **Delete** /ai_model_services/{id} | Delete an AI model service
[**GetAIModelService**](AIModelServiceAPI.md#GetAIModelService) | **Get** /ai_model_services/{id} | Get an AI model service by ID
[**ListAIModelServices**](AIModelServiceAPI.md#ListAIModelServices) | **Get** /ai_model_services | List AI model services
[**ListAIModelServicesWithModels**](AIModelServiceAPI.md#ListAIModelServicesWithModels) | **Get** /ai_model_services_with_models | List AI services with models
[**ListAIModelsForService**](AIModelServiceAPI.md#ListAIModelsForService) | **Get** /ai_model_services/{id}/models | List AI models for a service
[**UpdateAIModelService**](AIModelServiceAPI.md#UpdateAIModelService) | **Put** /ai_model_services/{id} | Update an AI model service



## CreateAIModelService

> AIModelServiceObject CreateAIModelService(ctx).Service(service).Execute()

Create a new AI model service



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
	service := *openapiclient.NewAIModelServiceWriteDto("Name_example", "ServiceImpl_example") // AIModelServiceWriteDto | AI Model Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServiceAPI.CreateAIModelService(context.Background()).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServiceAPI.CreateAIModelService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAIModelService`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServiceAPI.CreateAIModelService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAIModelServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | [**AIModelServiceWriteDto**](AIModelServiceWriteDto.md) | AI Model Service | 

### Return type

[**AIModelServiceObject**](AIModelServiceObject.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAIModelService

> AIModelServiceObject DeleteAIModelService(ctx, id).Execute()

Delete an AI model service



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
	id := "id_example" // string | AI Model Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServiceAPI.DeleteAIModelService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServiceAPI.DeleteAIModelService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAIModelService`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServiceAPI.DeleteAIModelService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAIModelServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AIModelServiceObject**](AIModelServiceObject.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIModelService

> AIModelServiceObject GetAIModelService(ctx, id).Execute()

Get an AI model service by ID



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
	id := "id_example" // string | AI Model Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServiceAPI.GetAIModelService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServiceAPI.GetAIModelService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIModelService`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServiceAPI.GetAIModelService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIModelServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AIModelServiceObject**](AIModelServiceObject.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIModelServices

> []AIModelServiceObject ListAIModelServices(ctx).Execute()

List AI model services



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
	resp, r, err := apiClient.AIModelServiceAPI.ListAIModelServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServiceAPI.ListAIModelServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelServices`: []AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServiceAPI.ListAIModelServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelServicesRequest struct via the builder pattern


### Return type

[**[]AIModelServiceObject**](AIModelServiceObject.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIModelServicesWithModels

> []AIModelServiceWithModels ListAIModelServicesWithModels(ctx).Execute()

List AI services with models



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
	resp, r, err := apiClient.AIModelServiceAPI.ListAIModelServicesWithModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServiceAPI.ListAIModelServicesWithModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelServicesWithModels`: []AIModelServiceWithModels
	fmt.Fprintf(os.Stdout, "Response from `AIModelServiceAPI.ListAIModelServicesWithModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelServicesWithModelsRequest struct via the builder pattern


### Return type

[**[]AIModelServiceWithModels**](AIModelServiceWithModels.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIModelsForService

> []AIModel ListAIModelsForService(ctx, id).Execute()

List AI models for a service



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
	id := "id_example" // string | AI Model Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServiceAPI.ListAIModelsForService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServiceAPI.ListAIModelsForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelsForService`: []AIModel
	fmt.Fprintf(os.Stdout, "Response from `AIModelServiceAPI.ListAIModelsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelsForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AIModel**](AIModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAIModelService

> AIModelServiceObject UpdateAIModelService(ctx, id).Service(service).Execute()

Update an AI model service



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
	id := "id_example" // string | AI Model Service ID
	service := *openapiclient.NewAIModelServiceWriteDto("Name_example", "ServiceImpl_example") // AIModelServiceWriteDto | AI Model Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServiceAPI.UpdateAIModelService(context.Background(), id).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServiceAPI.UpdateAIModelService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAIModelService`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServiceAPI.UpdateAIModelService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAIModelServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **service** | [**AIModelServiceWriteDto**](AIModelServiceWriteDto.md) | AI Model Service | 

### Return type

[**AIModelServiceObject**](AIModelServiceObject.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

