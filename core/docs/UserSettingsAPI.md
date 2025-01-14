# \UserSettingsAPI

All URIs are relative to *https://core.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserSettings**](UserSettingsAPI.md#CreateUserSettings) | **Post** /v1/user-settings | Create a UserSettings
[**DeleteUserSettings**](UserSettingsAPI.md#DeleteUserSettings) | **Delete** /v1/user-settings/{id} | Delete a UserSettings by ID
[**GetForUserSettings**](UserSettingsAPI.md#GetForUserSettings) | **Get** /v1/user-settings/me | Get UserSettings for the authenticated user
[**GetUserSettings**](UserSettingsAPI.md#GetUserSettings) | **Get** /v1/user-settings/{id} | Get a UserSettings by ID
[**PatchUserSettings**](UserSettingsAPI.md#PatchUserSettings) | **Patch** /v1/user-settings/{id} | Patch a UserSettings



## CreateUserSettings

> UserSettingsResponse CreateUserSettings(ctx).UserSettings(userSettings).Execute()

Create a UserSettings



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
	userSettings := *openapiclient.NewUserSettingsPostRequest("ColorScheme_example", "Language_example") // UserSettingsPostRequest | userSettings object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingsAPI.CreateUserSettings(context.Background()).UserSettings(userSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsAPI.CreateUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserSettings`: UserSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingsAPI.CreateUserSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userSettings** | [**UserSettingsPostRequest**](UserSettingsPostRequest.md) | userSettings object | 

### Return type

[**UserSettingsResponse**](UserSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserSettings

> DeleteUserSettings(ctx, id).Execute()

Delete a UserSettings by ID



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
	id := "id_example" // string | UserSettings ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserSettingsAPI.DeleteUserSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsAPI.DeleteUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UserSettings ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForUserSettings

> UserSettingsResponse GetForUserSettings(ctx).Execute()

Get UserSettings for the authenticated user



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
	resp, r, err := apiClient.UserSettingsAPI.GetForUserSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsAPI.GetForUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForUserSettings`: UserSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingsAPI.GetForUserSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetForUserSettingsRequest struct via the builder pattern


### Return type

[**UserSettingsResponse**](UserSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettings

> UserSettingsResponse GetUserSettings(ctx, id).Execute()

Get a UserSettings by ID



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
	id := "id_example" // string | UserSettings ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingsAPI.GetUserSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsAPI.GetUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSettings`: UserSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingsAPI.GetUserSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UserSettings ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSettingsResponse**](UserSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUserSettings

> UserSettingsResponse PatchUserSettings(ctx, id).UserSettings(userSettings).Execute()

Patch a UserSettings



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
	id := "id_example" // string | UserSettings ID
	userSettings := *openapiclient.NewUserSettingsPatchRequest() // UserSettingsPatchRequest | userSettings object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingsAPI.PatchUserSettings(context.Background(), id).UserSettings(userSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsAPI.PatchUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchUserSettings`: UserSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingsAPI.PatchUserSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UserSettings ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSettings** | [**UserSettingsPatchRequest**](UserSettingsPatchRequest.md) | userSettings object | 

### Return type

[**UserSettingsResponse**](UserSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

