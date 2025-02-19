# \OrganizationsAPI

All URIs are relative to *https://core.dev.ai.vaud.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvite**](OrganizationsAPI.md#AcceptInvite) | **Post** /v1/organizations/{id}/invites/{inviteId}/accept | Accept Organization Invite
[**CreateMember**](OrganizationsAPI.md#CreateMember) | **Post** /v1/organizations/{id}/members | Create a member for an Organization
[**CreateOrganization**](OrganizationsAPI.md#CreateOrganization) | **Post** /v1/organizations | Create an Organization
[**CreateOrganizationApiKey**](OrganizationsAPI.md#CreateOrganizationApiKey) | **Post** /v1/organizations/{id}/keys | Create an api key for an Organization
[**CreateOrganizationRole**](OrganizationsAPI.md#CreateOrganizationRole) | **Post** /v1/organizations/{id}/roles | Create a role for an Organization
[**CreateOrganizationSettings**](OrganizationsAPI.md#CreateOrganizationSettings) | **Post** /v1/organizations/{id}/settings | Create settings for an Organization
[**CreateSubscription**](OrganizationsAPI.md#CreateSubscription) | **Post** /v1/organizations/{id}/subscription | Create a subscription for an Organization
[**CreateTeam**](OrganizationsAPI.md#CreateTeam) | **Post** /v1/organizations/{id}/teams | Create a team for an Organization
[**CreateTeamMember**](OrganizationsAPI.md#CreateTeamMember) | **Post** /v1/organizations/{id}/teams/{teamId}/members | Add a member to a team
[**DeleteInvite**](OrganizationsAPI.md#DeleteInvite) | **Delete** /v1/organizations/{id}/invites/{inviteId} | Delete Organization Invite by ID
[**DeleteMember**](OrganizationsAPI.md#DeleteMember) | **Delete** /v1/organizations/{id}/members/{memberId} | Delete a member for an Organization
[**DeleteOrganization**](OrganizationsAPI.md#DeleteOrganization) | **Delete** /v1/organizations/{id} | Delete Organization by ID
[**DeleteOrganizationApiKey**](OrganizationsAPI.md#DeleteOrganizationApiKey) | **Delete** /v1/organizations/{id}/keys/{keyId} | Delete an api key for an Organization
[**DeleteOrganizationRole**](OrganizationsAPI.md#DeleteOrganizationRole) | **Delete** /v1/organizations/{id}/roles/{roleId} | Delete a role for an Organization
[**DeleteOrganizationSettings**](OrganizationsAPI.md#DeleteOrganizationSettings) | **Delete** /v1/organizations/{id}/settings | Delete settings for an Organization
[**DeleteSubscription**](OrganizationsAPI.md#DeleteSubscription) | **Delete** /v1/organizations/{id}/subscription | Delete a subscription for an Organization
[**DeleteTeam**](OrganizationsAPI.md#DeleteTeam) | **Delete** /v1/organizations/{id}/teams/{teamId} | Delete a team for an Organization
[**DeleteTeamMember**](OrganizationsAPI.md#DeleteTeamMember) | **Delete** /v1/organizations/{id}/teams/{teamId}/members/{memberId} | Remove a member from a team
[**GetAllMyOrganizations**](OrganizationsAPI.md#GetAllMyOrganizations) | **Get** /v1/organizations/me | Get all my organizations
[**GetAllMyTeams**](OrganizationsAPI.md#GetAllMyTeams) | **Get** /v1/organizations/{id}/teams/me | Get all teams for an Organization of the current user
[**GetCheckoutSession**](OrganizationsAPI.md#GetCheckoutSession) | **Get** /v1/organizations/{id}/checkout-sessions/{productId} | Get a checkout session for an Organization
[**GetCreditsPayments**](OrganizationsAPI.md#GetCreditsPayments) | **Get** /v1/organizations/{id}/credits-payments | Get all credits payments for an Organization
[**GetCustomerPortalSession**](OrganizationsAPI.md#GetCustomerPortalSession) | **Get** /v1/organizations/{id}/customer-portal-sessions/{typeId} | Get a stripe customer portal session for an Organization
[**GetInvite**](OrganizationsAPI.md#GetInvite) | **Get** /v1/organizations/{id}/invites/{inviteId} | Get an invite for an Organization
[**GetInvites**](OrganizationsAPI.md#GetInvites) | **Get** /v1/organizations/{id}/invites | Get all invites for an Organization
[**GetMember**](OrganizationsAPI.md#GetMember) | **Get** /v1/organizations/{id}/members/{memberId} | Get a member for an Organization
[**GetMembers**](OrganizationsAPI.md#GetMembers) | **Get** /v1/organizations/{id}/members | Get all members for an Organization
[**GetOrganization**](OrganizationsAPI.md#GetOrganization) | **Get** /v1/organizations/{id} | Get an Organization by id
[**GetOrganizationApiKeys**](OrganizationsAPI.md#GetOrganizationApiKeys) | **Get** /v1/organizations/{id}/keys | Get all api keys for an Organization
[**GetOrganizationRole**](OrganizationsAPI.md#GetOrganizationRole) | **Get** /v1/organizations/{id}/roles/{roleId} | Get a role for an Organization
[**GetOrganizationRoles**](OrganizationsAPI.md#GetOrganizationRoles) | **Get** /v1/organizations/{id}/roles | Get all roles for an Organization
[**GetOrganizationSettings**](OrganizationsAPI.md#GetOrganizationSettings) | **Get** /v1/organizations/{id}/settings | Get settings for an Organization
[**GetSubscription**](OrganizationsAPI.md#GetSubscription) | **Get** /v1/organizations/{id}/subscription | Get a subscription for an Organization
[**GetTeam**](OrganizationsAPI.md#GetTeam) | **Get** /v1/organizations/{id}/teams/{teamId} | Get a team for an Organization
[**GetTeamMember**](OrganizationsAPI.md#GetTeamMember) | **Get** /v1/organizations/{id}/teams/{teamId}/members/{memberId} | Get a member for a team
[**GetTeamMembers**](OrganizationsAPI.md#GetTeamMembers) | **Get** /v1/organizations/{id}/teams/{teamId}/members | Get all members for a team
[**GetTeams**](OrganizationsAPI.md#GetTeams) | **Get** /v1/organizations/{id}/teams | Get all teams for an Organization
[**Invite**](OrganizationsAPI.md#Invite) | **Post** /v1/organizations/{id}/invites | Invite a user to an Organization
[**PatchMember**](OrganizationsAPI.md#PatchMember) | **Patch** /v1/organizations/{id}/members/{memberId} | Patch a member for an Organization
[**PatchOrganization**](OrganizationsAPI.md#PatchOrganization) | **Patch** /v1/organizations/{id} | Patch an Organization by ID
[**PatchOrganizationApiKey**](OrganizationsAPI.md#PatchOrganizationApiKey) | **Patch** /v1/organizations/{id}/keys/{keyId} | Patch an api key for an Organization
[**PatchOrganizationRole**](OrganizationsAPI.md#PatchOrganizationRole) | **Patch** /v1/organizations/{id}/roles/{roleId} | Patch a role for an Organization
[**PatchOrganizationSettings**](OrganizationsAPI.md#PatchOrganizationSettings) | **Patch** /v1/organizations/{id}/settings | Patch settings for an Organization
[**PatchTeam**](OrganizationsAPI.md#PatchTeam) | **Patch** /v1/organizations/{id}/teams/{teamId} | Patch a team for an Organization
[**ReactivateSubscription**](OrganizationsAPI.md#ReactivateSubscription) | **Patch** /v1/organizations/{id}/subscription/reactivate | Reactivate a subscription for an Organization
[**ResendInvite**](OrganizationsAPI.md#ResendInvite) | **Post** /v1/organizations/{id}/invites/{inviteId}/resend | Resend Organization Invite by ID
[**UpdateSubscription**](OrganizationsAPI.md#UpdateSubscription) | **Patch** /v1/organizations/{id}/subscription | Update a subscription for an Organization



## AcceptInvite

> OrganizationMemberResponse AcceptInvite(ctx, id, inviteId).Execute()

Accept Organization Invite



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
	id := "id_example" // string | id of the organization
	inviteId := "inviteId_example" // string | id of the invite

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AcceptInvite(context.Background(), id, inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AcceptInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptInvite`: OrganizationMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AcceptInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**inviteId** | **string** | id of the invite | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationMemberResponse**](OrganizationMemberResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> OrganizationMemberResponse CreateMember(ctx, id).Member(member).Execute()

Create a member for an Organization



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
	id := "id_example" // string | id of the organization
	member := *openapiclient.NewMemberPostRequest("UserId_example") // MemberPostRequest | member object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateMember(context.Background(), id).Member(member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMember`: OrganizationMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **member** | [**MemberPostRequest**](MemberPostRequest.md) | member object | 

### Return type

[**OrganizationMemberResponse**](OrganizationMemberResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> OrganizationResponse CreateOrganization(ctx).Organization(organization).Execute()

Create an Organization



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
	organization := *openapiclient.NewOrganizationPostRequest("Name_example") // OrganizationPostRequest | organization object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganization(context.Background()).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganization`: OrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | [**OrganizationPostRequest**](OrganizationPostRequest.md) | organization object | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApiKey

> OrganizationApiKeyResponse CreateOrganizationApiKey(ctx, id).Key(key).Execute()

Create an api key for an Organization



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
	id := "id_example" // string | id of the organization
	key := *openapiclient.NewOrganizationApiKeyPostRequest("Name_example") // OrganizationApiKeyPostRequest | api key object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganizationApiKey(context.Background(), id).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganizationApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationApiKey`: OrganizationApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganizationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | [**OrganizationApiKeyPostRequest**](OrganizationApiKeyPostRequest.md) | api key object | 

### Return type

[**OrganizationApiKeyResponse**](OrganizationApiKeyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationRole

> OrganizationRoleResponse CreateOrganizationRole(ctx, id).Role(role).Execute()

Create a role for an Organization



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
	id := "id_example" // string | id of the organization
	role := *openapiclient.NewRolePostRequest("Name_example", []string{"Permissions_example"}) // RolePostRequest | role object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganizationRole(context.Background(), id).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationRole`: OrganizationRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganizationRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | [**RolePostRequest**](RolePostRequest.md) | role object | 

### Return type

[**OrganizationRoleResponse**](OrganizationRoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationSettings

> OrganizationSettingsResponse CreateOrganizationSettings(ctx, id).Settings(settings).Execute()

Create settings for an Organization



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
	id := "id_example" // string | id of the organization
	settings := *openapiclient.NewOrganizationSettingsPostRequest() // OrganizationSettingsPostRequest | settings object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganizationSettings(context.Background(), id).Settings(settings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganizationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSettings`: OrganizationSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganizationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settings** | [**OrganizationSettingsPostRequest**](OrganizationSettingsPostRequest.md) | settings object | 

### Return type

[**OrganizationSettingsResponse**](OrganizationSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> OrganizationSubscriptionResponse CreateSubscription(ctx, id).Organization(organization).Execute()

Create a subscription for an Organization



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
	id := "id_example" // string | id of the organization
	organization := *openapiclient.NewOrganizationSubscriptionPostRequest("ProductId_example", int32(123)) // OrganizationSubscriptionPostRequest | organization subscription object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateSubscription(context.Background(), id).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscription`: OrganizationSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**OrganizationSubscriptionPostRequest**](OrganizationSubscriptionPostRequest.md) | organization subscription object | 

### Return type

[**OrganizationSubscriptionResponse**](OrganizationSubscriptionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeam

> TeamResponse CreateTeam(ctx, id).Team(team).Execute()

Create a team for an Organization



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
	id := "id_example" // string | id of the organization
	team := *openapiclient.NewTeamPostRequest("Name_example") // TeamPostRequest | team object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateTeam(context.Background(), id).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeam`: TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **team** | [**TeamPostRequest**](TeamPostRequest.md) | team object | 

### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeamMember

> TeamMemberResponse CreateTeamMember(ctx, id, teamId).Member(member).Execute()

Add a member to a team



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
	id := "id_example" // string | id of the organization
	teamId := "teamId_example" // string | id of the team
	member := *openapiclient.NewMemberPostRequest("UserId_example") // MemberPostRequest | member object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateTeamMember(context.Background(), id, teamId).Member(member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeamMember`: TeamMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateTeamMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**teamId** | **string** | id of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **member** | [**MemberPostRequest**](MemberPostRequest.md) | member object | 

### Return type

[**TeamMemberResponse**](TeamMemberResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvite

> DeleteInvite(ctx, id, inviteId).Execute()

Delete Organization Invite by ID



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
	id := "id_example" // string | id of the organization
	inviteId := "inviteId_example" // string | id of the invite

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteInvite(context.Background(), id, inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**inviteId** | **string** | id of the invite | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMember

> MemberDeleteResponse DeleteMember(ctx, id, memberId).ResourcesAction(resourcesAction).Execute()

Delete a member for an Organization



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
	id := "id_example" // string | id of the organization
	memberId := "memberId_example" // string | id of the member
	resourcesAction := "resourcesAction_example" // string | action to take on resources owned by the member (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DeleteMember(context.Background(), id, memberId).ResourcesAction(resourcesAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMember`: MemberDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DeleteMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**memberId** | **string** | id of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourcesAction** | **string** | action to take on resources owned by the member | 

### Return type

[**MemberDeleteResponse**](MemberDeleteResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganization(ctx, id).Execute()

Delete Organization by ID



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationApiKey

> DeleteOrganizationApiKey(ctx, id, keyId).Execute()

Delete an api key for an Organization



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
	id := "id_example" // string | id of the organization
	keyId := "keyId_example" // string | id of the key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationApiKey(context.Background(), id, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**keyId** | **string** | id of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationRole

> DeleteOrganizationRole(ctx, id, roleId).Execute()

Delete a role for an Organization



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
	id := "id_example" // string | id of the organization
	roleId := "roleId_example" // string | id of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationRole(context.Background(), id, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**roleId** | **string** | id of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSettings

> DeleteOrganizationSettings(ctx, id).Execute()

Delete settings for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, id).Execute()

Delete a subscription for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DeleteTeam(ctx, id, teamId).Execute()

Delete a team for an Organization



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
	id := "id_example" // string | id of the organization
	teamId := "teamId_example" // string | id of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteTeam(context.Background(), id, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**teamId** | **string** | id of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamMember

> DeleteTeamMember(ctx, id, teamId, memberId).Execute()

Remove a member from a team



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
	id := "id_example" // string | id of the organization
	teamId := "teamId_example" // string | id of the team
	memberId := "memberId_example" // string | id of the member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteTeamMember(context.Background(), id, teamId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**teamId** | **string** | id of the team | 
**memberId** | **string** | id of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllMyOrganizations

> []OrganizationResponse GetAllMyOrganizations(ctx).Execute()

Get all my organizations



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
	resp, r, err := apiClient.OrganizationsAPI.GetAllMyOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetAllMyOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMyOrganizations`: []OrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetAllMyOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMyOrganizationsRequest struct via the builder pattern


### Return type

[**[]OrganizationResponse**](OrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllMyTeams

> []TeamResponse GetAllMyTeams(ctx, id).Execute()

Get all teams for an Organization of the current user



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetAllMyTeams(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetAllMyTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMyTeams`: []TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetAllMyTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMyTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TeamResponse**](TeamResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckoutSession

> OrganizationCheckoutSessionResponse GetCheckoutSession(ctx, id, productId).SuccessUrl(successUrl).CancelUrl(cancelUrl).Execute()

Get a checkout session for an Organization



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
	id := "id_example" // string | id of the organization
	productId := "productId_example" // string | id of the product
	successUrl := "successUrl_example" // string | The URL to which Stripe should send customers when payment is complete
	cancelUrl := "cancelUrl_example" // string | If set, Checkout displays a back button and customers will be directed to this URL if they decide to cancel payment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetCheckoutSession(context.Background(), id, productId).SuccessUrl(successUrl).CancelUrl(cancelUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckoutSession`: OrganizationCheckoutSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**productId** | **string** | id of the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **successUrl** | **string** | The URL to which Stripe should send customers when payment is complete | 
 **cancelUrl** | **string** | If set, Checkout displays a back button and customers will be directed to this URL if they decide to cancel payment | 

### Return type

[**OrganizationCheckoutSessionResponse**](OrganizationCheckoutSessionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditsPayments

> []OrganizationCreditsPaymentResponse GetCreditsPayments(ctx, id).Execute()

Get all credits payments for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetCreditsPayments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetCreditsPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditsPayments`: []OrganizationCreditsPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetCreditsPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditsPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganizationCreditsPaymentResponse**](OrganizationCreditsPaymentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerPortalSession

> OrganizationCustomerPortalSessionResponse GetCustomerPortalSession(ctx, id, typeId).ReturnUrl(returnUrl).SuccessUrl(successUrl).ProductId(productId).Seats(seats).Execute()

Get a stripe customer portal session for an Organization



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
	id := "id_example" // string | id of the organization
	typeId := "typeId_example" // string | type of the customer portal session
	returnUrl := "returnUrl_example" // string | The URL to redirect customers to when they click on the portal’s link to return to the website.
	successUrl := "successUrl_example" // string | The URL to redirect customers to when they successfully complete the flow. (optional)
	productId := "productId_example" // string | The product id to switch to. (optional)
	seats := int32(56) // int32 | The number of seats. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetCustomerPortalSession(context.Background(), id, typeId).ReturnUrl(returnUrl).SuccessUrl(successUrl).ProductId(productId).Seats(seats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetCustomerPortalSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerPortalSession`: OrganizationCustomerPortalSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetCustomerPortalSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**typeId** | **string** | type of the customer portal session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerPortalSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **returnUrl** | **string** | The URL to redirect customers to when they click on the portal’s link to return to the website. | 
 **successUrl** | **string** | The URL to redirect customers to when they successfully complete the flow. | 
 **productId** | **string** | The product id to switch to. | 
 **seats** | **int32** | The number of seats. | 

### Return type

[**OrganizationCustomerPortalSessionResponse**](OrganizationCustomerPortalSessionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvite

> OrganizationInviteResponse GetInvite(ctx, id, inviteId).Execute()

Get an invite for an Organization



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
	id := "id_example" // string | id of the organization
	inviteId := "inviteId_example" // string | id of the invite

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetInvite(context.Background(), id, inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvite`: OrganizationInviteResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**inviteId** | **string** | id of the invite | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationInviteResponse**](OrganizationInviteResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvites

> []OrganizationInviteResponse GetInvites(ctx, id).Execute()

Get all invites for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetInvites(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvites`: []OrganizationInviteResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganizationInviteResponse**](OrganizationInviteResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMember

> OrganizationMemberResponse GetMember(ctx, id, memberId).Execute()

Get a member for an Organization



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
	id := "id_example" // string | id of the organization
	memberId := "memberId_example" // string | id of the member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetMember(context.Background(), id, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMember`: OrganizationMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**memberId** | **string** | id of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationMemberResponse**](OrganizationMemberResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembers

> []OrganizationMemberResponse GetMembers(ctx, id).Execute()

Get all members for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetMembers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMembers`: []OrganizationMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganizationMemberResponse**](OrganizationMemberResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> OrganizationResponse GetOrganization(ctx, id).Execute()

Get an Organization by id



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: OrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiKeys

> []OrganizationApiKeyResponse GetOrganizationApiKeys(ctx, id).Execute()

Get all api keys for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationApiKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApiKeys`: []OrganizationApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganizationApiKeyResponse**](OrganizationApiKeyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationRole

> OrganizationRoleResponse GetOrganizationRole(ctx, id, roleId).Execute()

Get a role for an Organization



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
	id := "id_example" // string | id of the organization
	roleId := "roleId_example" // string | id of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationRole(context.Background(), id, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationRole`: OrganizationRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**roleId** | **string** | id of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationRoleResponse**](OrganizationRoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationRoles

> []OrganizationRoleResponse GetOrganizationRoles(ctx, id).Execute()

Get all roles for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationRoles(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationRoles`: []OrganizationRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganizationRoleResponse**](OrganizationRoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSettings

> OrganizationSettingsResponse GetOrganizationSettings(ctx, id).Execute()

Get settings for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSettings`: OrganizationSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationSettingsResponse**](OrganizationSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> OrganizationSubscriptionResponse GetSubscription(ctx, id).Execute()

Get a subscription for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscription`: OrganizationSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationSubscriptionResponse**](OrganizationSubscriptionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> TeamResponse GetTeam(ctx, id, teamId).Execute()

Get a team for an Organization



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
	id := "id_example" // string | id of the organization
	teamId := "teamId_example" // string | id of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetTeam(context.Background(), id, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeam`: TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**teamId** | **string** | id of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamMember

> TeamMemberResponse GetTeamMember(ctx, id, teamId, memberId).Execute()

Get a member for a team



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
	id := "id_example" // string | id of the organization
	teamId := "teamId_example" // string | id of the team
	memberId := "memberId_example" // string | id of the member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetTeamMember(context.Background(), id, teamId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamMember`: TeamMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetTeamMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**teamId** | **string** | id of the team | 
**memberId** | **string** | id of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TeamMemberResponse**](TeamMemberResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamMembers

> []TeamMemberResponse GetTeamMembers(ctx, id, teamId).Execute()

Get all members for a team



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
	id := "id_example" // string | id of the organization
	teamId := "teamId_example" // string | id of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetTeamMembers(context.Background(), id, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamMembers`: []TeamMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetTeamMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**teamId** | **string** | id of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TeamMemberResponse**](TeamMemberResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeams

> []TeamResponse GetTeams(ctx, id).Execute()

Get all teams for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetTeams(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeams`: []TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TeamResponse**](TeamResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Invite

> OrganizationInviteResponse Invite(ctx, id).OrganizationInvite(organizationInvite).Execute()

Invite a user to an Organization



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
	id := "id_example" // string | id of the organization
	organizationInvite := *openapiclient.NewOrganizationInviteRequest("Email_example") // OrganizationInviteRequest | organization invite object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Invite(context.Background(), id).OrganizationInvite(organizationInvite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Invite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Invite`: OrganizationInviteResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Invite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationInvite** | [**OrganizationInviteRequest**](OrganizationInviteRequest.md) | organization invite object | 

### Return type

[**OrganizationInviteResponse**](OrganizationInviteResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMember

> OrganizationMemberResponse PatchMember(ctx, id, memberId).Member(member).Execute()

Patch a member for an Organization



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
	id := "id_example" // string | id of the organization
	memberId := "memberId_example" // string | id of the member
	member := *openapiclient.NewMemberPatchRequest("RoleId_example") // MemberPatchRequest | member object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PatchMember(context.Background(), id, memberId).Member(member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PatchMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMember`: OrganizationMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PatchMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**memberId** | **string** | id of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **member** | [**MemberPatchRequest**](MemberPatchRequest.md) | member object | 

### Return type

[**OrganizationMemberResponse**](OrganizationMemberResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrganization

> OrganizationResponse PatchOrganization(ctx, id).Organization(organization).Execute()

Patch an Organization by ID



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
	id := "id_example" // string | id of the organization
	organization := *openapiclient.NewOrganizationPatchRequest("Name_example") // OrganizationPatchRequest | organization object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PatchOrganization(context.Background(), id).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PatchOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOrganization`: OrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PatchOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**OrganizationPatchRequest**](OrganizationPatchRequest.md) | organization object | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrganizationApiKey

> OrganizationApiKeyResponse PatchOrganizationApiKey(ctx, id, keyId).Key(key).Execute()

Patch an api key for an Organization



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
	id := "id_example" // string | id of the organization
	keyId := "keyId_example" // string | id of the key
	key := *openapiclient.NewOrganizationApiKeyPatchRequest() // OrganizationApiKeyPatchRequest | api key object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PatchOrganizationApiKey(context.Background(), id, keyId).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PatchOrganizationApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOrganizationApiKey`: OrganizationApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PatchOrganizationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**keyId** | **string** | id of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**OrganizationApiKeyPatchRequest**](OrganizationApiKeyPatchRequest.md) | api key object | 

### Return type

[**OrganizationApiKeyResponse**](OrganizationApiKeyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrganizationRole

> OrganizationRoleResponse PatchOrganizationRole(ctx, id, roleId).Role(role).Execute()

Patch a role for an Organization



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
	id := "id_example" // string | id of the organization
	roleId := "roleId_example" // string | id of the role
	role := *openapiclient.NewRolePatchRequest() // RolePatchRequest | role object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PatchOrganizationRole(context.Background(), id, roleId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PatchOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOrganizationRole`: OrganizationRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PatchOrganizationRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**roleId** | **string** | id of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | [**RolePatchRequest**](RolePatchRequest.md) | role object | 

### Return type

[**OrganizationRoleResponse**](OrganizationRoleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrganizationSettings

> OrganizationSettingsResponse PatchOrganizationSettings(ctx, id).Settings(settings).Execute()

Patch settings for an Organization



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
	id := "id_example" // string | id of the organization
	settings := *openapiclient.NewOrganizationSettingsPatchRequest() // OrganizationSettingsPatchRequest | settings object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PatchOrganizationSettings(context.Background(), id).Settings(settings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PatchOrganizationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOrganizationSettings`: OrganizationSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PatchOrganizationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrganizationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settings** | [**OrganizationSettingsPatchRequest**](OrganizationSettingsPatchRequest.md) | settings object | 

### Return type

[**OrganizationSettingsResponse**](OrganizationSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTeam

> TeamResponse PatchTeam(ctx, id, teamId).Team(team).Execute()

Patch a team for an Organization



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
	id := "id_example" // string | id of the organization
	teamId := "teamId_example" // string | id of the team
	team := *openapiclient.NewTeamPatchRequest() // TeamPatchRequest | team object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PatchTeam(context.Background(), id, teamId).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PatchTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTeam`: TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PatchTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**teamId** | **string** | id of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **team** | [**TeamPatchRequest**](TeamPatchRequest.md) | team object | 

### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactivateSubscription

> OrganizationSubscriptionResponse ReactivateSubscription(ctx, id).Execute()

Reactivate a subscription for an Organization



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
	id := "id_example" // string | id of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ReactivateSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ReactivateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactivateSubscription`: OrganizationSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ReactivateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationSubscriptionResponse**](OrganizationSubscriptionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendInvite

> OrganizationInviteResponse ResendInvite(ctx, id, inviteId).Execute()

Resend Organization Invite by ID



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
	id := "id_example" // string | id of the organization
	inviteId := "inviteId_example" // string | id of the invite

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ResendInvite(context.Background(), id, inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ResendInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendInvite`: OrganizationInviteResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ResendInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 
**inviteId** | **string** | id of the invite | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationInviteResponse**](OrganizationInviteResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> OrganizationSubscriptionResponse UpdateSubscription(ctx, id).Organization(organization).Execute()

Update a subscription for an Organization



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
	id := "id_example" // string | id of the organization
	organization := *openapiclient.NewOrganizationSubscriptionPatchRequest(int32(123)) // OrganizationSubscriptionPatchRequest | organization subscription object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateSubscription(context.Background(), id).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscription`: OrganizationSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**OrganizationSubscriptionPatchRequest**](OrganizationSubscriptionPatchRequest.md) | organization subscription object | 

### Return type

[**OrganizationSubscriptionResponse**](OrganizationSubscriptionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

