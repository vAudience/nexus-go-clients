# AgentTeamWriteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | Pointer to **[]string** |  | [optional] 
**CoordinatingAgentId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InitialUserMessages** | Pointer to **[]string** |  | [optional] 
**InitialUserMessagesMode** | Pointer to [**PromptInjectionMode**](PromptInjectionMode.md) |  | [optional] 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SystemMessages** | Pointer to **[]string** |  | [optional] 
**SystemMessagesMode** | Pointer to [**PromptInjectionMode**](PromptInjectionMode.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAgentTeamWriteDto

`func NewAgentTeamWriteDto() *AgentTeamWriteDto`

NewAgentTeamWriteDto instantiates a new AgentTeamWriteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTeamWriteDtoWithDefaults

`func NewAgentTeamWriteDtoWithDefaults() *AgentTeamWriteDto`

NewAgentTeamWriteDtoWithDefaults instantiates a new AgentTeamWriteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *AgentTeamWriteDto) GetAgentIds() []string`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *AgentTeamWriteDto) GetAgentIdsOk() (*[]string, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *AgentTeamWriteDto) SetAgentIds(v []string)`

SetAgentIds sets AgentIds field to given value.

### HasAgentIds

`func (o *AgentTeamWriteDto) HasAgentIds() bool`

HasAgentIds returns a boolean if a field has been set.

### GetCoordinatingAgentId

`func (o *AgentTeamWriteDto) GetCoordinatingAgentId() string`

GetCoordinatingAgentId returns the CoordinatingAgentId field if non-nil, zero value otherwise.

### GetCoordinatingAgentIdOk

`func (o *AgentTeamWriteDto) GetCoordinatingAgentIdOk() (*string, bool)`

GetCoordinatingAgentIdOk returns a tuple with the CoordinatingAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatingAgentId

`func (o *AgentTeamWriteDto) SetCoordinatingAgentId(v string)`

SetCoordinatingAgentId sets CoordinatingAgentId field to given value.

### HasCoordinatingAgentId

`func (o *AgentTeamWriteDto) HasCoordinatingAgentId() bool`

HasCoordinatingAgentId returns a boolean if a field has been set.

### GetDescription

`func (o *AgentTeamWriteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentTeamWriteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentTeamWriteDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentTeamWriteDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInitialUserMessages

`func (o *AgentTeamWriteDto) GetInitialUserMessages() []string`

GetInitialUserMessages returns the InitialUserMessages field if non-nil, zero value otherwise.

### GetInitialUserMessagesOk

`func (o *AgentTeamWriteDto) GetInitialUserMessagesOk() (*[]string, bool)`

GetInitialUserMessagesOk returns a tuple with the InitialUserMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUserMessages

`func (o *AgentTeamWriteDto) SetInitialUserMessages(v []string)`

SetInitialUserMessages sets InitialUserMessages field to given value.

### HasInitialUserMessages

`func (o *AgentTeamWriteDto) HasInitialUserMessages() bool`

HasInitialUserMessages returns a boolean if a field has been set.

### GetInitialUserMessagesMode

`func (o *AgentTeamWriteDto) GetInitialUserMessagesMode() PromptInjectionMode`

GetInitialUserMessagesMode returns the InitialUserMessagesMode field if non-nil, zero value otherwise.

### GetInitialUserMessagesModeOk

`func (o *AgentTeamWriteDto) GetInitialUserMessagesModeOk() (*PromptInjectionMode, bool)`

GetInitialUserMessagesModeOk returns a tuple with the InitialUserMessagesMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUserMessagesMode

`func (o *AgentTeamWriteDto) SetInitialUserMessagesMode(v PromptInjectionMode)`

SetInitialUserMessagesMode sets InitialUserMessagesMode field to given value.

### HasInitialUserMessagesMode

`func (o *AgentTeamWriteDto) HasInitialUserMessagesMode() bool`

HasInitialUserMessagesMode returns a boolean if a field has been set.

### GetMetaData

`func (o *AgentTeamWriteDto) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AgentTeamWriteDto) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AgentTeamWriteDto) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AgentTeamWriteDto) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetName

`func (o *AgentTeamWriteDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentTeamWriteDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentTeamWriteDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentTeamWriteDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSystemMessages

`func (o *AgentTeamWriteDto) GetSystemMessages() []string`

GetSystemMessages returns the SystemMessages field if non-nil, zero value otherwise.

### GetSystemMessagesOk

`func (o *AgentTeamWriteDto) GetSystemMessagesOk() (*[]string, bool)`

GetSystemMessagesOk returns a tuple with the SystemMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMessages

`func (o *AgentTeamWriteDto) SetSystemMessages(v []string)`

SetSystemMessages sets SystemMessages field to given value.

### HasSystemMessages

`func (o *AgentTeamWriteDto) HasSystemMessages() bool`

HasSystemMessages returns a boolean if a field has been set.

### GetSystemMessagesMode

`func (o *AgentTeamWriteDto) GetSystemMessagesMode() PromptInjectionMode`

GetSystemMessagesMode returns the SystemMessagesMode field if non-nil, zero value otherwise.

### GetSystemMessagesModeOk

`func (o *AgentTeamWriteDto) GetSystemMessagesModeOk() (*PromptInjectionMode, bool)`

GetSystemMessagesModeOk returns a tuple with the SystemMessagesMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMessagesMode

`func (o *AgentTeamWriteDto) SetSystemMessagesMode(v PromptInjectionMode)`

SetSystemMessagesMode sets SystemMessagesMode field to given value.

### HasSystemMessagesMode

`func (o *AgentTeamWriteDto) HasSystemMessagesMode() bool`

HasSystemMessagesMode returns a boolean if a field has been set.

### GetTags

`func (o *AgentTeamWriteDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AgentTeamWriteDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AgentTeamWriteDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AgentTeamWriteDto) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


