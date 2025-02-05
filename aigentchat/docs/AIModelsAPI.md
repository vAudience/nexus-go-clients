# \AIModelsAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAIModel**](AIModelsAPI.md#CreateAIModel) | **Post** /v1/ai-models | Create a new AI model
[**DeleteAIModel**](AIModelsAPI.md#DeleteAIModel) | **Delete** /v1/ai-models/{id} | Delete an AI model
[**GetAIModel**](AIModelsAPI.md#GetAIModel) | **Get** /v1/ai-models/{id} | Get an AI model by ID
[**ListAIModels**](AIModelsAPI.md#ListAIModels) | **Get** /v1/ai-models | List AI models
[**UpdateAIModel**](AIModelsAPI.md#UpdateAIModel) | **Put** /v1/ai-models/{id} | Update an AI model



## CreateAIModel

> AIModel CreateAIModel(ctx).AiModel(aiModel).Execute()

Create a new AI model



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
	aiModel := *openapiclient.NewAIModelWriteDto() // AIModelWriteDto | AI Model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelsAPI.CreateAIModel(context.Background()).AiModel(aiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.CreateAIModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAIModel`: AIModel
	fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.CreateAIModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAIModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiModel** | [**AIModelWriteDto**](AIModelWriteDto.md) | AI Model | 

### Return type

[**AIModel**](AIModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAIModel

> AIModel DeleteAIModel(ctx, id).Execute()

Delete an AI model



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
	id := "id_example" // string | AI Model ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelsAPI.DeleteAIModel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.DeleteAIModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAIModel`: AIModel
	fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.DeleteAIModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAIModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AIModel**](AIModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIModel

> AIModel GetAIModel(ctx, id).Execute()

Get an AI model by ID



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
	id := "id_example" // string | AI Model ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelsAPI.GetAIModel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetAIModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIModel`: AIModel
	fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetAIModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AIModel**](AIModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIModels

> []AIModel ListAIModels(ctx).Execute()

List AI models



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
	resp, r, err := apiClient.AIModelsAPI.ListAIModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.ListAIModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModels`: []AIModel
	fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.ListAIModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelsRequest struct via the builder pattern


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


## UpdateAIModel

> AIModel UpdateAIModel(ctx, id).AiModel(aiModel).Execute()

Update an AI model



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
	id := "id_example" // string | AI Model ID
	aiModel := *openapiclient.NewAIModel("Id_example", "ModelId_example", "Name_example", "OwnerId_example", "OwnerOrganizationId_example", "ServiceId_example") // AIModel | AI Model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelsAPI.UpdateAIModel(context.Background(), id).AiModel(aiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.UpdateAIModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAIModel`: AIModel
	fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.UpdateAIModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAIModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aiModel** | [**AIModel**](AIModel.md) | AI Model | 

### Return type

[**AIModel**](AIModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

