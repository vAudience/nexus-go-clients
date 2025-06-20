# OrganizationApiKeyPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**OrganizationRoleId** | Pointer to **string** |  | [optional] 
**RequiresRotation** | Pointer to **bool** |  | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**TeamIds** | Pointer to **[]string** |  | [optional] 
**UserEmail** | Pointer to **string** |  | [optional] 
**UserFullName** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationApiKeyPatchRequest

`func NewOrganizationApiKeyPatchRequest() *OrganizationApiKeyPatchRequest`

NewOrganizationApiKeyPatchRequest instantiates a new OrganizationApiKeyPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiKeyPatchRequestWithDefaults

`func NewOrganizationApiKeyPatchRequestWithDefaults() *OrganizationApiKeyPatchRequest`

NewOrganizationApiKeyPatchRequestWithDefaults instantiates a new OrganizationApiKeyPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationApiKeyPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationApiKeyPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationApiKeyPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationApiKeyPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationRoleId

`func (o *OrganizationApiKeyPatchRequest) GetOrganizationRoleId() string`

GetOrganizationRoleId returns the OrganizationRoleId field if non-nil, zero value otherwise.

### GetOrganizationRoleIdOk

`func (o *OrganizationApiKeyPatchRequest) GetOrganizationRoleIdOk() (*string, bool)`

GetOrganizationRoleIdOk returns a tuple with the OrganizationRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleId

`func (o *OrganizationApiKeyPatchRequest) SetOrganizationRoleId(v string)`

SetOrganizationRoleId sets OrganizationRoleId field to given value.

### HasOrganizationRoleId

`func (o *OrganizationApiKeyPatchRequest) HasOrganizationRoleId() bool`

HasOrganizationRoleId returns a boolean if a field has been set.

### GetRequiresRotation

`func (o *OrganizationApiKeyPatchRequest) GetRequiresRotation() bool`

GetRequiresRotation returns the RequiresRotation field if non-nil, zero value otherwise.

### GetRequiresRotationOk

`func (o *OrganizationApiKeyPatchRequest) GetRequiresRotationOk() (*bool, bool)`

GetRequiresRotationOk returns a tuple with the RequiresRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresRotation

`func (o *OrganizationApiKeyPatchRequest) SetRequiresRotation(v bool)`

SetRequiresRotation sets RequiresRotation field to given value.

### HasRequiresRotation

`func (o *OrganizationApiKeyPatchRequest) HasRequiresRotation() bool`

HasRequiresRotation returns a boolean if a field has been set.

### GetRoleIds

`func (o *OrganizationApiKeyPatchRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *OrganizationApiKeyPatchRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *OrganizationApiKeyPatchRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *OrganizationApiKeyPatchRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetTeamIds

`func (o *OrganizationApiKeyPatchRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *OrganizationApiKeyPatchRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *OrganizationApiKeyPatchRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *OrganizationApiKeyPatchRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUserEmail

`func (o *OrganizationApiKeyPatchRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *OrganizationApiKeyPatchRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *OrganizationApiKeyPatchRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *OrganizationApiKeyPatchRequest) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserFullName

`func (o *OrganizationApiKeyPatchRequest) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *OrganizationApiKeyPatchRequest) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *OrganizationApiKeyPatchRequest) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.

### HasUserFullName

`func (o *OrganizationApiKeyPatchRequest) HasUserFullName() bool`

HasUserFullName returns a boolean if a field has been set.

### GetUserName

`func (o *OrganizationApiKeyPatchRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OrganizationApiKeyPatchRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OrganizationApiKeyPatchRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OrganizationApiKeyPatchRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


