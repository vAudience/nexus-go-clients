# TeamMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**Id** | **string** |  | 
**TeamId** | **string** |  | 
**TeamOwner** | **bool** |  | 
**TeamRole** | [**TeamRoleReducedResponse**](TeamRoleReducedResponse.md) |  | 
**UpdatedAt** | **string** |  | 
**User** | [**UserResponse**](UserResponse.md) |  | 

## Methods

### NewTeamMemberResponse

`func NewTeamMemberResponse(createdAt string, id string, teamId string, teamOwner bool, teamRole TeamRoleReducedResponse, updatedAt string, user UserResponse, ) *TeamMemberResponse`

NewTeamMemberResponse instantiates a new TeamMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberResponseWithDefaults

`func NewTeamMemberResponseWithDefaults() *TeamMemberResponse`

NewTeamMemberResponseWithDefaults instantiates a new TeamMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TeamMemberResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamMemberResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamMemberResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *TeamMemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamMemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamMemberResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTeamId

`func (o *TeamMemberResponse) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamMemberResponse) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamMemberResponse) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetTeamOwner

`func (o *TeamMemberResponse) GetTeamOwner() bool`

GetTeamOwner returns the TeamOwner field if non-nil, zero value otherwise.

### GetTeamOwnerOk

`func (o *TeamMemberResponse) GetTeamOwnerOk() (*bool, bool)`

GetTeamOwnerOk returns a tuple with the TeamOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOwner

`func (o *TeamMemberResponse) SetTeamOwner(v bool)`

SetTeamOwner sets TeamOwner field to given value.


### GetTeamRole

`func (o *TeamMemberResponse) GetTeamRole() TeamRoleReducedResponse`

GetTeamRole returns the TeamRole field if non-nil, zero value otherwise.

### GetTeamRoleOk

`func (o *TeamMemberResponse) GetTeamRoleOk() (*TeamRoleReducedResponse, bool)`

GetTeamRoleOk returns a tuple with the TeamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRole

`func (o *TeamMemberResponse) SetTeamRole(v TeamRoleReducedResponse)`

SetTeamRole sets TeamRole field to given value.


### GetUpdatedAt

`func (o *TeamMemberResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamMemberResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamMemberResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUser

`func (o *TeamMemberResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeamMemberResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeamMemberResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


