# \CollectionsAPI

All URIs are relative to *https://file-manager.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCollection**](CollectionsAPI.md#CreateCollection) | **Post** /v1/organizations/{org_id}/collections | Create a collection
[**DeleteCollection**](CollectionsAPI.md#DeleteCollection) | **Delete** /v1/organizations/{org_id}/collections/{collection_id} | Delete a collection
[**GetCollection**](CollectionsAPI.md#GetCollection) | **Get** /v1/organizations/{org_id}/collections/{collection_id} | Get a collection
[**GetCollectionFile**](CollectionsAPI.md#GetCollectionFile) | **Get** /v1/organizations/{org_id}/collections/{collection_id}/files/{file_id} | Get a collection file
[**GetCollectionSettings**](CollectionsAPI.md#GetCollectionSettings) | **Get** /v1/organizations/{org_id}/collections/settings | Get collection settings
[**ListCollectionFiles**](CollectionsAPI.md#ListCollectionFiles) | **Get** /v1/organizations/{org_id}/collections/{collection_id}/files | List a collection&#39;s files
[**ListMyCollections**](CollectionsAPI.md#ListMyCollections) | **Get** /v1/organizations/{org_id}/collections/me | List my collections
[**RetryCollectionFile**](CollectionsAPI.md#RetryCollectionFile) | **Post** /v1/organizations/{org_id}/collections/{collection_id}/files/{file_id}/retry | Retry processing a failed collection file
[**UpdateCollection**](CollectionsAPI.md#UpdateCollection) | **Patch** /v1/organizations/{org_id}/collections/{collection_id} | Update a collection



## CreateCollection

> CollectionResponse CreateCollection(ctx, orgId).CollectionCreateRequest(collectionCreateRequest).Execute()

Create a collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	collectionCreateRequest := *openapiclient.NewCollectionCreateRequest("Name_example") // CollectionCreateRequest | Collection create request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CreateCollection(context.Background(), orgId).CollectionCreateRequest(collectionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CreateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCollection`: CollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CreateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionCreateRequest** | [**CollectionCreateRequest**](CollectionCreateRequest.md) | Collection create request | 

### Return type

[**CollectionResponse**](CollectionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> DeleteCollection(ctx, orgId, collectionId).Execute()

Delete a collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	collectionId := "collectionId_example" // string | collection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.DeleteCollection(context.Background(), orgId, collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.DeleteCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**collectionId** | **string** | collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionRequest struct via the builder pattern


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


## GetCollection

> CollectionResponse GetCollection(ctx, orgId, collectionId).Execute()

Get a collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	collectionId := "collectionId_example" // string | collection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.GetCollection(context.Background(), orgId, collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.GetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollection`: CollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.GetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**collectionId** | **string** | collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponse**](CollectionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionFile

> CollectionFileResponse GetCollectionFile(ctx, orgId, collectionId, fileId).Execute()

Get a collection file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	collectionId := "collectionId_example" // string | collection ID
	fileId := "fileId_example" // string | file ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.GetCollectionFile(context.Background(), orgId, collectionId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.GetCollectionFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionFile`: CollectionFileResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.GetCollectionFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**collectionId** | **string** | collection ID | 
**fileId** | **string** | file ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CollectionFileResponse**](CollectionFileResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionSettings

> CollectionSettings GetCollectionSettings(ctx, orgId).Execute()

Get collection settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.GetCollectionSettings(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.GetCollectionSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionSettings`: CollectionSettings
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.GetCollectionSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionSettings**](CollectionSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionFiles

> CollectionFileListResponse ListCollectionFiles(ctx, orgId, collectionId).Offset(offset).Limit(limit).Sort(sort).Order(order).Q(q).Execute()

List a collection's files



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	collectionId := "collectionId_example" // string | collection ID
	offset := int32(56) // int32 | pagination offset (optional)
	limit := int32(56) // int32 | pagination limit (optional)
	sort := "sort_example" // string | sort column (optional)
	order := "order_example" // string | sort order (optional)
	q := "q_example" // string | case-insensitive partial match on original file name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ListCollectionFiles(context.Background(), orgId, collectionId).Offset(offset).Limit(limit).Sort(sort).Order(order).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ListCollectionFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionFiles`: CollectionFileListResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ListCollectionFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**collectionId** | **string** | collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | pagination offset | 
 **limit** | **int32** | pagination limit | 
 **sort** | **string** | sort column | 
 **order** | **string** | sort order | 
 **q** | **string** | case-insensitive partial match on original file name | 

### Return type

[**CollectionFileListResponse**](CollectionFileListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyCollections

> CollectionListResponse ListMyCollections(ctx, orgId).Offset(offset).Limit(limit).Sort(sort).Order(order).Q(q).Ids(ids).Execute()

List my collections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	offset := int32(56) // int32 | pagination offset (optional)
	limit := int32(56) // int32 | pagination limit (optional)
	sort := "sort_example" // string | sort column (optional)
	order := "order_example" // string | sort order (optional)
	q := "q_example" // string | case-insensitive partial match on collection name (optional)
	ids := []string{"Inner_example"} // []string | collection ids to restrict the listing to, comma-separated (?ids=a,b) or repeated (?ids=a&ids=b) -- both are accepted (max 100 distinct, AND-ed with q); unknown or non-owned ids are simply absent, and offset/limit still apply (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ListMyCollections(context.Background(), orgId).Offset(offset).Limit(limit).Sort(sort).Order(order).Q(q).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ListMyCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyCollections`: CollectionListResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ListMyCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMyCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | pagination offset | 
 **limit** | **int32** | pagination limit | 
 **sort** | **string** | sort column | 
 **order** | **string** | sort order | 
 **q** | **string** | case-insensitive partial match on collection name | 
 **ids** | **[]string** | collection ids to restrict the listing to, comma-separated (?ids&#x3D;a,b) or repeated (?ids&#x3D;a&amp;ids&#x3D;b) -- both are accepted (max 100 distinct, AND-ed with q); unknown or non-owned ids are simply absent, and offset/limit still apply | 

### Return type

[**CollectionListResponse**](CollectionListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryCollectionFile

> RetryCollectionFile(ctx, orgId, collectionId, fileId).Execute()

Retry processing a failed collection file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	collectionId := "collectionId_example" // string | collection ID
	fileId := "fileId_example" // string | file ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.RetryCollectionFile(context.Background(), orgId, collectionId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.RetryCollectionFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**collectionId** | **string** | collection ID | 
**fileId** | **string** | file ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryCollectionFileRequest struct via the builder pattern


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


## UpdateCollection

> CollectionResponse UpdateCollection(ctx, orgId, collectionId).CollectionUpdateRequest(collectionUpdateRequest).Execute()

Update a collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vaudience/nexus-go-clients/filemanager"
)

func main() {
	orgId := "orgId_example" // string | organization ID
	collectionId := "collectionId_example" // string | collection ID
	collectionUpdateRequest := *openapiclient.NewCollectionUpdateRequest() // CollectionUpdateRequest | Collection update request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.UpdateCollection(context.Background(), orgId, collectionId).CollectionUpdateRequest(collectionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.UpdateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCollection`: CollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.UpdateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization ID | 
**collectionId** | **string** | collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionUpdateRequest** | [**CollectionUpdateRequest**](CollectionUpdateRequest.md) | Collection update request | 

### Return type

[**CollectionResponse**](CollectionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

