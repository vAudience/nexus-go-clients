# AgentWriteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToolGuidelines** | Pointer to **bool** |  | [optional] 
**AssignedTools** | Pointer to **[]string** |  | [optional] 
**AttachedFileIds** | Pointer to **[]string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**DefaultFileUploadCategory** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**I18n** | Pointer to [**map[string]AgentI18n**](AgentI18n.md) |  | [optional] 
**IgnoreIncomingOverwrite** | Pointer to **bool** |  | [optional] 
**InitialUserMessages** | Pointer to **[]string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**ModelActions** | Pointer to **[]string** |  | [optional] 
**ModelHostLocation** | Pointer to [**HostingLocation**](HostingLocation.md) |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ReadAccess** | Pointer to [**AccessScope**](AccessScope.md) |  | [optional] 
**SystemMessages** | Pointer to **[]string** |  | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 
**ToolConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to [**AgentType**](AgentType.md) |  | [optional] 
**UseTools** | Pointer to **bool** |  | [optional] 
**WriteAccess** | Pointer to [**AccessScope**](AccessScope.md) |  | [optional] 

## Methods

### NewAgentWriteDto

`func NewAgentWriteDto() *AgentWriteDto`

NewAgentWriteDto instantiates a new AgentWriteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWriteDtoWithDefaults

`func NewAgentWriteDtoWithDefaults() *AgentWriteDto`

NewAgentWriteDtoWithDefaults instantiates a new AgentWriteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddToolGuidelines

`func (o *AgentWriteDto) GetAddToolGuidelines() bool`

GetAddToolGuidelines returns the AddToolGuidelines field if non-nil, zero value otherwise.

### GetAddToolGuidelinesOk

`func (o *AgentWriteDto) GetAddToolGuidelinesOk() (*bool, bool)`

GetAddToolGuidelinesOk returns a tuple with the AddToolGuidelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToolGuidelines

`func (o *AgentWriteDto) SetAddToolGuidelines(v bool)`

SetAddToolGuidelines sets AddToolGuidelines field to given value.

### HasAddToolGuidelines

`func (o *AgentWriteDto) HasAddToolGuidelines() bool`

HasAddToolGuidelines returns a boolean if a field has been set.

### GetAssignedTools

`func (o *AgentWriteDto) GetAssignedTools() []string`

GetAssignedTools returns the AssignedTools field if non-nil, zero value otherwise.

### GetAssignedToolsOk

`func (o *AgentWriteDto) GetAssignedToolsOk() (*[]string, bool)`

GetAssignedToolsOk returns a tuple with the AssignedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTools

`func (o *AgentWriteDto) SetAssignedTools(v []string)`

SetAssignedTools sets AssignedTools field to given value.

### HasAssignedTools

`func (o *AgentWriteDto) HasAssignedTools() bool`

HasAssignedTools returns a boolean if a field has been set.

### GetAttachedFileIds

`func (o *AgentWriteDto) GetAttachedFileIds() []string`

GetAttachedFileIds returns the AttachedFileIds field if non-nil, zero value otherwise.

### GetAttachedFileIdsOk

`func (o *AgentWriteDto) GetAttachedFileIdsOk() (*[]string, bool)`

GetAttachedFileIdsOk returns a tuple with the AttachedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFileIds

`func (o *AgentWriteDto) SetAttachedFileIds(v []string)`

SetAttachedFileIds sets AttachedFileIds field to given value.

### HasAttachedFileIds

`func (o *AgentWriteDto) HasAttachedFileIds() bool`

HasAttachedFileIds returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *AgentWriteDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *AgentWriteDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *AgentWriteDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *AgentWriteDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetDefaultFileUploadCategory

`func (o *AgentWriteDto) GetDefaultFileUploadCategory() string`

GetDefaultFileUploadCategory returns the DefaultFileUploadCategory field if non-nil, zero value otherwise.

### GetDefaultFileUploadCategoryOk

`func (o *AgentWriteDto) GetDefaultFileUploadCategoryOk() (*string, bool)`

GetDefaultFileUploadCategoryOk returns a tuple with the DefaultFileUploadCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFileUploadCategory

`func (o *AgentWriteDto) SetDefaultFileUploadCategory(v string)`

SetDefaultFileUploadCategory sets DefaultFileUploadCategory field to given value.

### HasDefaultFileUploadCategory

`func (o *AgentWriteDto) HasDefaultFileUploadCategory() bool`

HasDefaultFileUploadCategory returns a boolean if a field has been set.

### GetDescription

`func (o *AgentWriteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentWriteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentWriteDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentWriteDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetI18n

`func (o *AgentWriteDto) GetI18n() map[string]AgentI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *AgentWriteDto) GetI18nOk() (*map[string]AgentI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *AgentWriteDto) SetI18n(v map[string]AgentI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *AgentWriteDto) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetIgnoreIncomingOverwrite

`func (o *AgentWriteDto) GetIgnoreIncomingOverwrite() bool`

GetIgnoreIncomingOverwrite returns the IgnoreIncomingOverwrite field if non-nil, zero value otherwise.

### GetIgnoreIncomingOverwriteOk

`func (o *AgentWriteDto) GetIgnoreIncomingOverwriteOk() (*bool, bool)`

GetIgnoreIncomingOverwriteOk returns a tuple with the IgnoreIncomingOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIncomingOverwrite

`func (o *AgentWriteDto) SetIgnoreIncomingOverwrite(v bool)`

SetIgnoreIncomingOverwrite sets IgnoreIncomingOverwrite field to given value.

### HasIgnoreIncomingOverwrite

`func (o *AgentWriteDto) HasIgnoreIncomingOverwrite() bool`

HasIgnoreIncomingOverwrite returns a boolean if a field has been set.

### GetInitialUserMessages

`func (o *AgentWriteDto) GetInitialUserMessages() []string`

GetInitialUserMessages returns the InitialUserMessages field if non-nil, zero value otherwise.

### GetInitialUserMessagesOk

`func (o *AgentWriteDto) GetInitialUserMessagesOk() (*[]string, bool)`

GetInitialUserMessagesOk returns a tuple with the InitialUserMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUserMessages

`func (o *AgentWriteDto) SetInitialUserMessages(v []string)`

SetInitialUserMessages sets InitialUserMessages field to given value.

### HasInitialUserMessages

`func (o *AgentWriteDto) HasInitialUserMessages() bool`

HasInitialUserMessages returns a boolean if a field has been set.

### GetInternalId

`func (o *AgentWriteDto) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AgentWriteDto) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AgentWriteDto) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AgentWriteDto) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetMetaData

`func (o *AgentWriteDto) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AgentWriteDto) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AgentWriteDto) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AgentWriteDto) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetModelActions

`func (o *AgentWriteDto) GetModelActions() []string`

GetModelActions returns the ModelActions field if non-nil, zero value otherwise.

### GetModelActionsOk

`func (o *AgentWriteDto) GetModelActionsOk() (*[]string, bool)`

GetModelActionsOk returns a tuple with the ModelActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelActions

`func (o *AgentWriteDto) SetModelActions(v []string)`

SetModelActions sets ModelActions field to given value.

### HasModelActions

`func (o *AgentWriteDto) HasModelActions() bool`

HasModelActions returns a boolean if a field has been set.

### GetModelHostLocation

`func (o *AgentWriteDto) GetModelHostLocation() HostingLocation`

GetModelHostLocation returns the ModelHostLocation field if non-nil, zero value otherwise.

### GetModelHostLocationOk

`func (o *AgentWriteDto) GetModelHostLocationOk() (*HostingLocation, bool)`

GetModelHostLocationOk returns a tuple with the ModelHostLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelHostLocation

`func (o *AgentWriteDto) SetModelHostLocation(v HostingLocation)`

SetModelHostLocation sets ModelHostLocation field to given value.

### HasModelHostLocation

`func (o *AgentWriteDto) HasModelHostLocation() bool`

HasModelHostLocation returns a boolean if a field has been set.

### GetModelId

`func (o *AgentWriteDto) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AgentWriteDto) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AgentWriteDto) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *AgentWriteDto) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetName

`func (o *AgentWriteDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentWriteDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentWriteDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentWriteDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *AgentWriteDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AgentWriteDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AgentWriteDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AgentWriteDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReadAccess

`func (o *AgentWriteDto) GetReadAccess() AccessScope`

GetReadAccess returns the ReadAccess field if non-nil, zero value otherwise.

### GetReadAccessOk

`func (o *AgentWriteDto) GetReadAccessOk() (*AccessScope, bool)`

GetReadAccessOk returns a tuple with the ReadAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAccess

`func (o *AgentWriteDto) SetReadAccess(v AccessScope)`

SetReadAccess sets ReadAccess field to given value.

### HasReadAccess

`func (o *AgentWriteDto) HasReadAccess() bool`

HasReadAccess returns a boolean if a field has been set.

### GetSystemMessages

`func (o *AgentWriteDto) GetSystemMessages() []string`

GetSystemMessages returns the SystemMessages field if non-nil, zero value otherwise.

### GetSystemMessagesOk

`func (o *AgentWriteDto) GetSystemMessagesOk() (*[]string, bool)`

GetSystemMessagesOk returns a tuple with the SystemMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMessages

`func (o *AgentWriteDto) SetSystemMessages(v []string)`

SetSystemMessages sets SystemMessages field to given value.

### HasSystemMessages

`func (o *AgentWriteDto) HasSystemMessages() bool`

HasSystemMessages returns a boolean if a field has been set.

### GetTagIds

`func (o *AgentWriteDto) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *AgentWriteDto) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *AgentWriteDto) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *AgentWriteDto) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetToolConfig

`func (o *AgentWriteDto) GetToolConfig() map[string]interface{}`

GetToolConfig returns the ToolConfig field if non-nil, zero value otherwise.

### GetToolConfigOk

`func (o *AgentWriteDto) GetToolConfigOk() (*map[string]interface{}, bool)`

GetToolConfigOk returns a tuple with the ToolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfig

`func (o *AgentWriteDto) SetToolConfig(v map[string]interface{})`

SetToolConfig sets ToolConfig field to given value.

### HasToolConfig

`func (o *AgentWriteDto) HasToolConfig() bool`

HasToolConfig returns a boolean if a field has been set.

### GetType

`func (o *AgentWriteDto) GetType() AgentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentWriteDto) GetTypeOk() (*AgentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentWriteDto) SetType(v AgentType)`

SetType sets Type field to given value.

### HasType

`func (o *AgentWriteDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseTools

`func (o *AgentWriteDto) GetUseTools() bool`

GetUseTools returns the UseTools field if non-nil, zero value otherwise.

### GetUseToolsOk

`func (o *AgentWriteDto) GetUseToolsOk() (*bool, bool)`

GetUseToolsOk returns a tuple with the UseTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTools

`func (o *AgentWriteDto) SetUseTools(v bool)`

SetUseTools sets UseTools field to given value.

### HasUseTools

`func (o *AgentWriteDto) HasUseTools() bool`

HasUseTools returns a boolean if a field has been set.

### GetWriteAccess

`func (o *AgentWriteDto) GetWriteAccess() AccessScope`

GetWriteAccess returns the WriteAccess field if non-nil, zero value otherwise.

### GetWriteAccessOk

`func (o *AgentWriteDto) GetWriteAccessOk() (*AccessScope, bool)`

GetWriteAccessOk returns a tuple with the WriteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAccess

`func (o *AgentWriteDto) SetWriteAccess(v AccessScope)`

SetWriteAccess sets WriteAccess field to given value.

### HasWriteAccess

`func (o *AgentWriteDto) HasWriteAccess() bool`

HasWriteAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


