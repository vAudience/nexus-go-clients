# Prompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CurrentVersion** | Pointer to **int32** |  | [optional] 
**DefaultAgentId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**InternalId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerOrganizationId** | Pointer to **string** |  | [optional] 
**ReadAccess** | Pointer to [**AccessScope**](AccessScope.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UserAccess** | Pointer to [**UserAccessView**](UserAccessView.md) |  | [optional] 
**Versions** | Pointer to [**[]PromptVersion**](PromptVersion.md) |  | [optional] 
**WriteAccess** | Pointer to [**AccessScope**](AccessScope.md) |  | [optional] 

## Methods

### NewPrompt

`func NewPrompt(id string, title string, ) *Prompt`

NewPrompt instantiates a new Prompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptWithDefaults

`func NewPromptWithDefaults() *Prompt`

NewPromptWithDefaults instantiates a new Prompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Prompt) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Prompt) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Prompt) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Prompt) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *Prompt) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *Prompt) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *Prompt) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *Prompt) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDefaultAgentId

`func (o *Prompt) GetDefaultAgentId() string`

GetDefaultAgentId returns the DefaultAgentId field if non-nil, zero value otherwise.

### GetDefaultAgentIdOk

`func (o *Prompt) GetDefaultAgentIdOk() (*string, bool)`

GetDefaultAgentIdOk returns a tuple with the DefaultAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAgentId

`func (o *Prompt) SetDefaultAgentId(v string)`

SetDefaultAgentId sets DefaultAgentId field to given value.

### HasDefaultAgentId

`func (o *Prompt) HasDefaultAgentId() bool`

HasDefaultAgentId returns a boolean if a field has been set.

### GetDescription

`func (o *Prompt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Prompt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Prompt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Prompt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Prompt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Prompt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Prompt) SetId(v string)`

SetId sets Id field to given value.


### GetInternalId

`func (o *Prompt) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Prompt) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Prompt) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Prompt) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetOwnerId

`func (o *Prompt) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Prompt) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Prompt) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Prompt) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerOrganizationId

`func (o *Prompt) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *Prompt) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *Prompt) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.

### HasOwnerOrganizationId

`func (o *Prompt) HasOwnerOrganizationId() bool`

HasOwnerOrganizationId returns a boolean if a field has been set.

### GetReadAccess

`func (o *Prompt) GetReadAccess() AccessScope`

GetReadAccess returns the ReadAccess field if non-nil, zero value otherwise.

### GetReadAccessOk

`func (o *Prompt) GetReadAccessOk() (*AccessScope, bool)`

GetReadAccessOk returns a tuple with the ReadAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAccess

`func (o *Prompt) SetReadAccess(v AccessScope)`

SetReadAccess sets ReadAccess field to given value.

### HasReadAccess

`func (o *Prompt) HasReadAccess() bool`

HasReadAccess returns a boolean if a field has been set.

### GetTags

`func (o *Prompt) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Prompt) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Prompt) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Prompt) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *Prompt) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *Prompt) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *Prompt) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *Prompt) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetTitle

`func (o *Prompt) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Prompt) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Prompt) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *Prompt) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Prompt) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Prompt) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Prompt) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserAccess

`func (o *Prompt) GetUserAccess() UserAccessView`

GetUserAccess returns the UserAccess field if non-nil, zero value otherwise.

### GetUserAccessOk

`func (o *Prompt) GetUserAccessOk() (*UserAccessView, bool)`

GetUserAccessOk returns a tuple with the UserAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccess

`func (o *Prompt) SetUserAccess(v UserAccessView)`

SetUserAccess sets UserAccess field to given value.

### HasUserAccess

`func (o *Prompt) HasUserAccess() bool`

HasUserAccess returns a boolean if a field has been set.

### GetVersions

`func (o *Prompt) GetVersions() []PromptVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Prompt) GetVersionsOk() (*[]PromptVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Prompt) SetVersions(v []PromptVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Prompt) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetWriteAccess

`func (o *Prompt) GetWriteAccess() AccessScope`

GetWriteAccess returns the WriteAccess field if non-nil, zero value otherwise.

### GetWriteAccessOk

`func (o *Prompt) GetWriteAccessOk() (*AccessScope, bool)`

GetWriteAccessOk returns a tuple with the WriteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAccess

`func (o *Prompt) SetWriteAccess(v AccessScope)`

SetWriteAccess sets WriteAccess field to given value.

### HasWriteAccess

`func (o *Prompt) HasWriteAccess() bool`

HasWriteAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


