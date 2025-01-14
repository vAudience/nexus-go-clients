# AgentTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | Pointer to **[]string** |  | [optional] 
**CoordinatingAgentId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**InitialUserMessages** | Pointer to **[]string** |  | [optional] 
**InitialUserMessagesMode** | Pointer to [**PromptInjectionMode**](PromptInjectionMode.md) |  | [optional] 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**OwnerOrganizationId** | **string** |  | 
**SystemMessages** | Pointer to **[]string** |  | [optional] 
**SystemMessagesMode** | Pointer to [**PromptInjectionMode**](PromptInjectionMode.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentTeam

`func NewAgentTeam(id string, name string, ownerId string, ownerOrganizationId string, ) *AgentTeam`

NewAgentTeam instantiates a new AgentTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTeamWithDefaults

`func NewAgentTeamWithDefaults() *AgentTeam`

NewAgentTeamWithDefaults instantiates a new AgentTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *AgentTeam) GetAgentIds() []string`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *AgentTeam) GetAgentIdsOk() (*[]string, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *AgentTeam) SetAgentIds(v []string)`

SetAgentIds sets AgentIds field to given value.

### HasAgentIds

`func (o *AgentTeam) HasAgentIds() bool`

HasAgentIds returns a boolean if a field has been set.

### GetCoordinatingAgentId

`func (o *AgentTeam) GetCoordinatingAgentId() string`

GetCoordinatingAgentId returns the CoordinatingAgentId field if non-nil, zero value otherwise.

### GetCoordinatingAgentIdOk

`func (o *AgentTeam) GetCoordinatingAgentIdOk() (*string, bool)`

GetCoordinatingAgentIdOk returns a tuple with the CoordinatingAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatingAgentId

`func (o *AgentTeam) SetCoordinatingAgentId(v string)`

SetCoordinatingAgentId sets CoordinatingAgentId field to given value.

### HasCoordinatingAgentId

`func (o *AgentTeam) HasCoordinatingAgentId() bool`

HasCoordinatingAgentId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentTeam) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentTeam) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentTeam) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentTeam) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AgentTeam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentTeam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentTeam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentTeam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AgentTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentTeam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentTeam) SetId(v string)`

SetId sets Id field to given value.


### GetInitialUserMessages

`func (o *AgentTeam) GetInitialUserMessages() []string`

GetInitialUserMessages returns the InitialUserMessages field if non-nil, zero value otherwise.

### GetInitialUserMessagesOk

`func (o *AgentTeam) GetInitialUserMessagesOk() (*[]string, bool)`

GetInitialUserMessagesOk returns a tuple with the InitialUserMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUserMessages

`func (o *AgentTeam) SetInitialUserMessages(v []string)`

SetInitialUserMessages sets InitialUserMessages field to given value.

### HasInitialUserMessages

`func (o *AgentTeam) HasInitialUserMessages() bool`

HasInitialUserMessages returns a boolean if a field has been set.

### GetInitialUserMessagesMode

`func (o *AgentTeam) GetInitialUserMessagesMode() PromptInjectionMode`

GetInitialUserMessagesMode returns the InitialUserMessagesMode field if non-nil, zero value otherwise.

### GetInitialUserMessagesModeOk

`func (o *AgentTeam) GetInitialUserMessagesModeOk() (*PromptInjectionMode, bool)`

GetInitialUserMessagesModeOk returns a tuple with the InitialUserMessagesMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUserMessagesMode

`func (o *AgentTeam) SetInitialUserMessagesMode(v PromptInjectionMode)`

SetInitialUserMessagesMode sets InitialUserMessagesMode field to given value.

### HasInitialUserMessagesMode

`func (o *AgentTeam) HasInitialUserMessagesMode() bool`

HasInitialUserMessagesMode returns a boolean if a field has been set.

### GetMetaData

`func (o *AgentTeam) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AgentTeam) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AgentTeam) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AgentTeam) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetName

`func (o *AgentTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentTeam) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *AgentTeam) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AgentTeam) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AgentTeam) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerOrganizationId

`func (o *AgentTeam) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *AgentTeam) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *AgentTeam) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.


### GetSystemMessages

`func (o *AgentTeam) GetSystemMessages() []string`

GetSystemMessages returns the SystemMessages field if non-nil, zero value otherwise.

### GetSystemMessagesOk

`func (o *AgentTeam) GetSystemMessagesOk() (*[]string, bool)`

GetSystemMessagesOk returns a tuple with the SystemMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMessages

`func (o *AgentTeam) SetSystemMessages(v []string)`

SetSystemMessages sets SystemMessages field to given value.

### HasSystemMessages

`func (o *AgentTeam) HasSystemMessages() bool`

HasSystemMessages returns a boolean if a field has been set.

### GetSystemMessagesMode

`func (o *AgentTeam) GetSystemMessagesMode() PromptInjectionMode`

GetSystemMessagesMode returns the SystemMessagesMode field if non-nil, zero value otherwise.

### GetSystemMessagesModeOk

`func (o *AgentTeam) GetSystemMessagesModeOk() (*PromptInjectionMode, bool)`

GetSystemMessagesModeOk returns a tuple with the SystemMessagesMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMessagesMode

`func (o *AgentTeam) SetSystemMessagesMode(v PromptInjectionMode)`

SetSystemMessagesMode sets SystemMessagesMode field to given value.

### HasSystemMessagesMode

`func (o *AgentTeam) HasSystemMessagesMode() bool`

HasSystemMessagesMode returns a boolean if a field has been set.

### GetTags

`func (o *AgentTeam) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AgentTeam) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AgentTeam) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AgentTeam) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentTeam) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentTeam) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentTeam) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentTeam) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AgentTeam) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AgentTeam) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AgentTeam) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AgentTeam) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


