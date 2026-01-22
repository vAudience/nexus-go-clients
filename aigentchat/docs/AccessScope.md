# AccessScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasTeams** | Pointer to **bool** |  | [optional] 
**IncludeOrganization** | Pointer to **bool** | Organization-wide access - all organization members When true, all other settings need to be false or empty | [optional] 
**IsPrivate** | Pointer to **bool** | Private access - only owner When true, all other settings need to be false or empty | [optional] 
**IsPublic** | Pointer to **bool** | Public access - all authenticated users can access When true, all other settings need to be false or empty | [optional] 
**TeamIds** | Pointer to **[]string** | Team-based access - members of these teams Can be combined with UserIDs, but not with IsPrivate, IsPublic and IncludeOrganization | [optional] 
**UserIds** | Pointer to **[]string** | User-specific access - these specific users Can be combined with TeamIDs, but not with IsPrivate, IsPublic and IncludeOrganization | [optional] 

## Methods

### NewAccessScope

`func NewAccessScope() *AccessScope`

NewAccessScope instantiates a new AccessScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessScopeWithDefaults

`func NewAccessScopeWithDefaults() *AccessScope`

NewAccessScopeWithDefaults instantiates a new AccessScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasTeams

`func (o *AccessScope) GetHasTeams() bool`

GetHasTeams returns the HasTeams field if non-nil, zero value otherwise.

### GetHasTeamsOk

`func (o *AccessScope) GetHasTeamsOk() (*bool, bool)`

GetHasTeamsOk returns a tuple with the HasTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTeams

`func (o *AccessScope) SetHasTeams(v bool)`

SetHasTeams sets HasTeams field to given value.

### HasHasTeams

`func (o *AccessScope) HasHasTeams() bool`

HasHasTeams returns a boolean if a field has been set.

### GetIncludeOrganization

`func (o *AccessScope) GetIncludeOrganization() bool`

GetIncludeOrganization returns the IncludeOrganization field if non-nil, zero value otherwise.

### GetIncludeOrganizationOk

`func (o *AccessScope) GetIncludeOrganizationOk() (*bool, bool)`

GetIncludeOrganizationOk returns a tuple with the IncludeOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOrganization

`func (o *AccessScope) SetIncludeOrganization(v bool)`

SetIncludeOrganization sets IncludeOrganization field to given value.

### HasIncludeOrganization

`func (o *AccessScope) HasIncludeOrganization() bool`

HasIncludeOrganization returns a boolean if a field has been set.

### GetIsPrivate

`func (o *AccessScope) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *AccessScope) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *AccessScope) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *AccessScope) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsPublic

`func (o *AccessScope) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AccessScope) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AccessScope) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AccessScope) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetTeamIds

`func (o *AccessScope) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *AccessScope) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *AccessScope) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *AccessScope) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUserIds

`func (o *AccessScope) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *AccessScope) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *AccessScope) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *AccessScope) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


