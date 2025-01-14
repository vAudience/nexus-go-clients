# OrganizationMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**CreatedBy** | **string** |  | 
**CreatedByName** | **string** |  | 
**Id** | **string** |  | 
**InvitedAt** | **string** |  | 
**OrganizationId** | **string** |  | 
**OrganizationOwner** | **bool** |  | 
**OrganizationProductRoles** | [**[]RoleReducedResponse**](RoleReducedResponse.md) |  | 
**OrganizationRole** | [**OrganizationRoleReducedResponse**](OrganizationRoleReducedResponse.md) |  | 
**UpdatedAt** | **string** |  | 
**User** | [**UserResponse**](UserResponse.md) |  | 

## Methods

### NewOrganizationMemberResponse

`func NewOrganizationMemberResponse(createdAt string, createdBy string, createdByName string, id string, invitedAt string, organizationId string, organizationOwner bool, organizationProductRoles []RoleReducedResponse, organizationRole OrganizationRoleReducedResponse, updatedAt string, user UserResponse, ) *OrganizationMemberResponse`

NewOrganizationMemberResponse instantiates a new OrganizationMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberResponseWithDefaults

`func NewOrganizationMemberResponseWithDefaults() *OrganizationMemberResponse`

NewOrganizationMemberResponseWithDefaults instantiates a new OrganizationMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganizationMemberResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationMemberResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationMemberResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *OrganizationMemberResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationMemberResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationMemberResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedByName

`func (o *OrganizationMemberResponse) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *OrganizationMemberResponse) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *OrganizationMemberResponse) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.


### GetId

`func (o *OrganizationMemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationMemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationMemberResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInvitedAt

`func (o *OrganizationMemberResponse) GetInvitedAt() string`

GetInvitedAt returns the InvitedAt field if non-nil, zero value otherwise.

### GetInvitedAtOk

`func (o *OrganizationMemberResponse) GetInvitedAtOk() (*string, bool)`

GetInvitedAtOk returns a tuple with the InvitedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedAt

`func (o *OrganizationMemberResponse) SetInvitedAt(v string)`

SetInvitedAt sets InvitedAt field to given value.


### GetOrganizationId

`func (o *OrganizationMemberResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationMemberResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationMemberResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationOwner

`func (o *OrganizationMemberResponse) GetOrganizationOwner() bool`

GetOrganizationOwner returns the OrganizationOwner field if non-nil, zero value otherwise.

### GetOrganizationOwnerOk

`func (o *OrganizationMemberResponse) GetOrganizationOwnerOk() (*bool, bool)`

GetOrganizationOwnerOk returns a tuple with the OrganizationOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationOwner

`func (o *OrganizationMemberResponse) SetOrganizationOwner(v bool)`

SetOrganizationOwner sets OrganizationOwner field to given value.


### GetOrganizationProductRoles

`func (o *OrganizationMemberResponse) GetOrganizationProductRoles() []RoleReducedResponse`

GetOrganizationProductRoles returns the OrganizationProductRoles field if non-nil, zero value otherwise.

### GetOrganizationProductRolesOk

`func (o *OrganizationMemberResponse) GetOrganizationProductRolesOk() (*[]RoleReducedResponse, bool)`

GetOrganizationProductRolesOk returns a tuple with the OrganizationProductRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProductRoles

`func (o *OrganizationMemberResponse) SetOrganizationProductRoles(v []RoleReducedResponse)`

SetOrganizationProductRoles sets OrganizationProductRoles field to given value.


### GetOrganizationRole

`func (o *OrganizationMemberResponse) GetOrganizationRole() OrganizationRoleReducedResponse`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *OrganizationMemberResponse) GetOrganizationRoleOk() (*OrganizationRoleReducedResponse, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *OrganizationMemberResponse) SetOrganizationRole(v OrganizationRoleReducedResponse)`

SetOrganizationRole sets OrganizationRole field to given value.


### GetUpdatedAt

`func (o *OrganizationMemberResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationMemberResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationMemberResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUser

`func (o *OrganizationMemberResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationMemberResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationMemberResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


