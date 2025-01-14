# AgentPrompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** |  | [optional] 
**CurrentVersion** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**OwnerId** | **string** |  | 
**OwnerOrganizationId** | **string** |  | 
**Space** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**UpdatedAt** | Pointer to **int32** |  | [optional] 
**Versions** | Pointer to [**[]PromptVersion**](PromptVersion.md) |  | [optional] 
**Visibility** | Pointer to [**AgentPromptVisibilityStates**](AgentPromptVisibilityStates.md) |  | [optional] 

## Methods

### NewAgentPrompt

`func NewAgentPrompt(id string, ownerId string, ownerOrganizationId string, title string, ) *AgentPrompt`

NewAgentPrompt instantiates a new AgentPrompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPromptWithDefaults

`func NewAgentPromptWithDefaults() *AgentPrompt`

NewAgentPromptWithDefaults instantiates a new AgentPrompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AgentPrompt) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentPrompt) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentPrompt) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentPrompt) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *AgentPrompt) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *AgentPrompt) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *AgentPrompt) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *AgentPrompt) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDescription

`func (o *AgentPrompt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentPrompt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentPrompt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentPrompt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AgentPrompt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentPrompt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentPrompt) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerId

`func (o *AgentPrompt) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AgentPrompt) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AgentPrompt) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerOrganizationId

`func (o *AgentPrompt) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *AgentPrompt) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *AgentPrompt) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.


### GetSpace

`func (o *AgentPrompt) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *AgentPrompt) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *AgentPrompt) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *AgentPrompt) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTags

`func (o *AgentPrompt) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AgentPrompt) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AgentPrompt) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AgentPrompt) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *AgentPrompt) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *AgentPrompt) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *AgentPrompt) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *AgentPrompt) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetTitle

`func (o *AgentPrompt) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentPrompt) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentPrompt) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *AgentPrompt) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentPrompt) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentPrompt) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentPrompt) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersions

`func (o *AgentPrompt) GetVersions() []PromptVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AgentPrompt) GetVersionsOk() (*[]PromptVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AgentPrompt) SetVersions(v []PromptVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AgentPrompt) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetVisibility

`func (o *AgentPrompt) GetVisibility() AgentPromptVisibilityStates`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AgentPrompt) GetVisibilityOk() (*AgentPromptVisibilityStates, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AgentPrompt) SetVisibility(v AgentPromptVisibilityStates)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AgentPrompt) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


