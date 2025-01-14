# OrganizationApiKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** |  | 
**ApiKeyHint** | **string** |  | 
**CreatedAt** | **string** |  | 
**CreatedBy** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**OrganizationRole** | [**OrganizationRoleReducedResponse**](OrganizationRoleReducedResponse.md) |  | 
**Roles** | [**[]RoleReducedResponse**](RoleReducedResponse.md) |  | 
**UpdatedAt** | **string** |  | 
**UserEmail** | **string** |  | 
**UserFullName** | **string** |  | 
**UserId** | **string** |  | 
**UserName** | **string** |  | 

## Methods

### NewOrganizationApiKeyResponse

`func NewOrganizationApiKeyResponse(apiKey string, apiKeyHint string, createdAt string, createdBy string, id string, name string, organizationId string, organizationRole OrganizationRoleReducedResponse, roles []RoleReducedResponse, updatedAt string, userEmail string, userFullName string, userId string, userName string, ) *OrganizationApiKeyResponse`

NewOrganizationApiKeyResponse instantiates a new OrganizationApiKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiKeyResponseWithDefaults

`func NewOrganizationApiKeyResponseWithDefaults() *OrganizationApiKeyResponse`

NewOrganizationApiKeyResponseWithDefaults instantiates a new OrganizationApiKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *OrganizationApiKeyResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OrganizationApiKeyResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OrganizationApiKeyResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiKeyHint

`func (o *OrganizationApiKeyResponse) GetApiKeyHint() string`

GetApiKeyHint returns the ApiKeyHint field if non-nil, zero value otherwise.

### GetApiKeyHintOk

`func (o *OrganizationApiKeyResponse) GetApiKeyHintOk() (*string, bool)`

GetApiKeyHintOk returns a tuple with the ApiKeyHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyHint

`func (o *OrganizationApiKeyResponse) SetApiKeyHint(v string)`

SetApiKeyHint sets ApiKeyHint field to given value.


### GetCreatedAt

`func (o *OrganizationApiKeyResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationApiKeyResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationApiKeyResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *OrganizationApiKeyResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationApiKeyResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationApiKeyResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *OrganizationApiKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationApiKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationApiKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationApiKeyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationApiKeyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationApiKeyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *OrganizationApiKeyResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationApiKeyResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationApiKeyResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationRole

`func (o *OrganizationApiKeyResponse) GetOrganizationRole() OrganizationRoleReducedResponse`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *OrganizationApiKeyResponse) GetOrganizationRoleOk() (*OrganizationRoleReducedResponse, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *OrganizationApiKeyResponse) SetOrganizationRole(v OrganizationRoleReducedResponse)`

SetOrganizationRole sets OrganizationRole field to given value.


### GetRoles

`func (o *OrganizationApiKeyResponse) GetRoles() []RoleReducedResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrganizationApiKeyResponse) GetRolesOk() (*[]RoleReducedResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrganizationApiKeyResponse) SetRoles(v []RoleReducedResponse)`

SetRoles sets Roles field to given value.


### GetUpdatedAt

`func (o *OrganizationApiKeyResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationApiKeyResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationApiKeyResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserEmail

`func (o *OrganizationApiKeyResponse) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *OrganizationApiKeyResponse) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *OrganizationApiKeyResponse) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserFullName

`func (o *OrganizationApiKeyResponse) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *OrganizationApiKeyResponse) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *OrganizationApiKeyResponse) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.


### GetUserId

`func (o *OrganizationApiKeyResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationApiKeyResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationApiKeyResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *OrganizationApiKeyResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OrganizationApiKeyResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OrganizationApiKeyResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


