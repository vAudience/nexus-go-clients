# \AIModelServicesAPI

All URIs are relative to *https://aigentchat.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAIModelService**](AIModelServicesAPI.md#CreateAIModelService) | **Post** /v1/organizations/{org_id}/ai-model-services | Create a new AI model service
[**DeleteAIModelService**](AIModelServicesAPI.md#DeleteAIModelService) | **Delete** /v1/organizations/{org_id}/ai-model-services/{id} | Delete an AI model service
[**GetAIModelService**](AIModelServicesAPI.md#GetAIModelService) | **Get** /v1/organizations/{org_id}/ai-model-services/{id} | Get an AI model service by ID
[**GetAIModelServiceLegacy**](AIModelServicesAPI.md#GetAIModelServiceLegacy) | **Get** /v1/ai-model-services/{id} | Get an AI model service by ID
[**ListAIModelServices**](AIModelServicesAPI.md#ListAIModelServices) | **Get** /v1/organizations/{org_id}/ai-model-services | List AI model services
[**ListAIModelServicesLegacy**](AIModelServicesAPI.md#ListAIModelServicesLegacy) | **Get** /v1/ai-model-services | List AI model services
[**ListAIModelServicesWithModels**](AIModelServicesAPI.md#ListAIModelServicesWithModels) | **Get** /v1/organizations/{org_id}/ai-model-services-with-models | List AI services with models
[**ListAIModelServicesWithModelsLegacy**](AIModelServicesAPI.md#ListAIModelServicesWithModelsLegacy) | **Get** /v1/ai-model-services-with-models | List AI services with models
[**ListAIModelsForService**](AIModelServicesAPI.md#ListAIModelsForService) | **Get** /v1/organizations/{org_id}/ai-model-services/{id}/models | List AI models for a service
[**ListAIModelsForServiceLegacy**](AIModelServicesAPI.md#ListAIModelsForServiceLegacy) | **Get** /v1/ai-model-services/{id}/models | List AI models for a service
[**UpdateAIModelService**](AIModelServicesAPI.md#UpdateAIModelService) | **Put** /v1/organizations/{org_id}/ai-model-services/{id} | Update an AI model service



## CreateAIModelService

> AIModelServiceObject CreateAIModelService(ctx, orgId).Service(service).Execute()

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
	orgId := "orgId_example" // string | organization ID
	service := *openapiclient.NewAIModelServiceWriteDto("Name_example", openapiclient.AiServiceId("anthropic")) // AIModelServiceWriteDto | AI Model Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServicesAPI.CreateAIModelService(context.Background(), orgId).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.CreateAIModelService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAIModelService`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.CreateAIModelService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

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

> AIModelServiceObject DeleteAIModelService(ctx, orgId, id).Execute()

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
	orgId := "orgId_example" // string | organization ID
	id := "id_example" // string | AI Model Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServicesAPI.DeleteAIModelService(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.DeleteAIModelService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAIModelService`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.DeleteAIModelService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
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

> AIModelServiceObject GetAIModelService(ctx, orgId, id).Execute()

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
	orgId := "orgId_example" // string | organization ID
	id := "id_example" // string | AI Model Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServicesAPI.GetAIModelService(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.GetAIModelService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIModelService`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.GetAIModelService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
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


## GetAIModelServiceLegacy

> AIModelServiceObject GetAIModelServiceLegacy(ctx, id).Execute()

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
	resp, r, err := apiClient.AIModelServicesAPI.GetAIModelServiceLegacy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.GetAIModelServiceLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIModelServiceLegacy`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.GetAIModelServiceLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIModelServiceLegacyRequest struct via the builder pattern


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

> []AIModelServiceObject ListAIModelServices(ctx, orgId).Execute()

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
	orgId := "orgId_example" // string | organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServicesAPI.ListAIModelServices(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.ListAIModelServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelServices`: []AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.ListAIModelServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListAIModelServicesLegacy

> []AIModelServiceObject ListAIModelServicesLegacy(ctx).Execute()

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
	resp, r, err := apiClient.AIModelServicesAPI.ListAIModelServicesLegacy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.ListAIModelServicesLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelServicesLegacy`: []AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.ListAIModelServicesLegacy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelServicesLegacyRequest struct via the builder pattern


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

> []AIModelServiceWithModels ListAIModelServicesWithModels(ctx, orgId).OrgId2(orgId2).Execute()

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
	orgId := "orgId_example" // string | organization ID
	orgId2 := "orgId_example" // string | return only available AI model services for this organization (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServicesAPI.ListAIModelServicesWithModels(context.Background(), orgId).OrgId2(orgId2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.ListAIModelServicesWithModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelServicesWithModels`: []AIModelServiceWithModels
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.ListAIModelServicesWithModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelServicesWithModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgId2** | **string** | return only available AI model services for this organization | 

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


## ListAIModelServicesWithModelsLegacy

> []AIModelServiceWithModels ListAIModelServicesWithModelsLegacy(ctx).Execute()

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
	resp, r, err := apiClient.AIModelServicesAPI.ListAIModelServicesWithModelsLegacy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.ListAIModelServicesWithModelsLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelServicesWithModelsLegacy`: []AIModelServiceWithModels
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.ListAIModelServicesWithModelsLegacy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelServicesWithModelsLegacyRequest struct via the builder pattern


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

> []AIModel ListAIModelsForService(ctx, orgId, id).Execute()

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
	orgId := "orgId_example" // string | organization ID
	id := "id_example" // string | AI Model Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServicesAPI.ListAIModelsForService(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.ListAIModelsForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelsForService`: []AIModel
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.ListAIModelsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
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


## ListAIModelsForServiceLegacy

> []AIModel ListAIModelsForServiceLegacy(ctx, id).Execute()

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
	resp, r, err := apiClient.AIModelServicesAPI.ListAIModelsForServiceLegacy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.ListAIModelsForServiceLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModelsForServiceLegacy`: []AIModel
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.ListAIModelsForServiceLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AI Model Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelsForServiceLegacyRequest struct via the builder pattern


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

> AIModelServiceObject UpdateAIModelService(ctx, orgId, id).Service(service).Execute()

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
	orgId := "orgId_example" // string | organization ID
	id := "id_example" // string | AI Model Service ID
	service := *openapiclient.NewAIModelServiceWriteDto("Name_example", openapiclient.AiServiceId("anthropic")) // AIModelServiceWriteDto | AI Model Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelServicesAPI.UpdateAIModelService(context.Background(), orgId, id).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelServicesAPI.UpdateAIModelService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAIModelService`: AIModelServiceObject
	fmt.Fprintf(os.Stdout, "Response from `AIModelServicesAPI.UpdateAIModelService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
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

