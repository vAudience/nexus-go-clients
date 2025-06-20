# OrganizationApiKeyPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OrganizationRoleId** | Pointer to **string** |  | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**TeamIds** | Pointer to **[]string** |  | [optional] 
**UserEmail** | **string** |  | 
**UserFullName** | **string** |  | 
**UserName** | **string** |  | 

## Methods

### NewOrganizationApiKeyPostRequest

`func NewOrganizationApiKeyPostRequest(name string, userEmail string, userFullName string, userName string, ) *OrganizationApiKeyPostRequest`

NewOrganizationApiKeyPostRequest instantiates a new OrganizationApiKeyPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiKeyPostRequestWithDefaults

`func NewOrganizationApiKeyPostRequestWithDefaults() *OrganizationApiKeyPostRequest`

NewOrganizationApiKeyPostRequestWithDefaults instantiates a new OrganizationApiKeyPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationApiKeyPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationApiKeyPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationApiKeyPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationRoleId

`func (o *OrganizationApiKeyPostRequest) GetOrganizationRoleId() string`

GetOrganizationRoleId returns the OrganizationRoleId field if non-nil, zero value otherwise.

### GetOrganizationRoleIdOk

`func (o *OrganizationApiKeyPostRequest) GetOrganizationRoleIdOk() (*string, bool)`

GetOrganizationRoleIdOk returns a tuple with the OrganizationRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleId

`func (o *OrganizationApiKeyPostRequest) SetOrganizationRoleId(v string)`

SetOrganizationRoleId sets OrganizationRoleId field to given value.

### HasOrganizationRoleId

`func (o *OrganizationApiKeyPostRequest) HasOrganizationRoleId() bool`

HasOrganizationRoleId returns a boolean if a field has been set.

### GetRoleIds

`func (o *OrganizationApiKeyPostRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *OrganizationApiKeyPostRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *OrganizationApiKeyPostRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *OrganizationApiKeyPostRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetTeamIds

`func (o *OrganizationApiKeyPostRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *OrganizationApiKeyPostRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *OrganizationApiKeyPostRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *OrganizationApiKeyPostRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUserEmail

`func (o *OrganizationApiKeyPostRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *OrganizationApiKeyPostRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *OrganizationApiKeyPostRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserFullName

`func (o *OrganizationApiKeyPostRequest) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *OrganizationApiKeyPostRequest) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *OrganizationApiKeyPostRequest) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.


### GetUserName

`func (o *OrganizationApiKeyPostRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OrganizationApiKeyPostRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OrganizationApiKeyPostRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


