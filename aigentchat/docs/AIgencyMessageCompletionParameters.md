# AIgencyMessageCompletionParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**AssignedCollectionIds** | Pointer to **[]string** | AssignedCollectionIDs is the merged collection set the completion tried to resolve (ChatCompletionRequest.MergedCollectionIDs); the corpora they resolved to are in ToolConfigs. | [optional] 
**ContinueInstructionOnMaxTokens** | Pointer to **string** |  | [optional] 
**ContinueOnMaxTokens** | Pointer to **bool** |  | [optional] 
**DeeprToolAutoActivated** | Pointer to **bool** |  | [optional] 
**ModelHostLocation** | Pointer to [**HostingLocation**](HostingLocation.md) |  | [optional] 
**ModelParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**SelectedTools** | Pointer to **[]string** | SelectedTools is the effective selection, after resolveAgentToolDefaults and deepr auto-activation. | [optional] 
**SetMessageHistoryIds** | Pointer to **[]string** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 
**ToolConfigs** | Pointer to **map[string]interface{}** | ToolConfigs is the effective runtime config per delivered tool id (ExecutionContextBase.BuildToolConfig), e.g. the deepr tool&#39;s resolved corpus_ids. Tools without any config are omitted. | [optional] 
**ToolFunctionIds** | Pointer to **[]string** | ToolFunctionIDs are the functions actually delivered to the model (ExecutionContextBase.ToolFunctions). A selected tool that is unannounced, invisible to the org or outside its hosting locations is absent. | [optional] 
**UseChannelMessagesAsHistory** | Pointer to **bool** |  | [optional] 
**UseTools** | Pointer to **bool** |  | [optional] 
**VarReplacements** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAIgencyMessageCompletionParameters

`func NewAIgencyMessageCompletionParameters() *AIgencyMessageCompletionParameters`

NewAIgencyMessageCompletionParameters instantiates a new AIgencyMessageCompletionParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyMessageCompletionParametersWithDefaults

`func NewAIgencyMessageCompletionParametersWithDefaults() *AIgencyMessageCompletionParameters`

NewAIgencyMessageCompletionParametersWithDefaults instantiates a new AIgencyMessageCompletionParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AIgencyMessageCompletionParameters) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AIgencyMessageCompletionParameters) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AIgencyMessageCompletionParameters) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AIgencyMessageCompletionParameters) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAssignedCollectionIds

`func (o *AIgencyMessageCompletionParameters) GetAssignedCollectionIds() []string`

GetAssignedCollectionIds returns the AssignedCollectionIds field if non-nil, zero value otherwise.

### GetAssignedCollectionIdsOk

`func (o *AIgencyMessageCompletionParameters) GetAssignedCollectionIdsOk() (*[]string, bool)`

GetAssignedCollectionIdsOk returns a tuple with the AssignedCollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCollectionIds

`func (o *AIgencyMessageCompletionParameters) SetAssignedCollectionIds(v []string)`

SetAssignedCollectionIds sets AssignedCollectionIds field to given value.

### HasAssignedCollectionIds

`func (o *AIgencyMessageCompletionParameters) HasAssignedCollectionIds() bool`

HasAssignedCollectionIds returns a boolean if a field has been set.

### GetContinueInstructionOnMaxTokens

`func (o *AIgencyMessageCompletionParameters) GetContinueInstructionOnMaxTokens() string`

GetContinueInstructionOnMaxTokens returns the ContinueInstructionOnMaxTokens field if non-nil, zero value otherwise.

### GetContinueInstructionOnMaxTokensOk

`func (o *AIgencyMessageCompletionParameters) GetContinueInstructionOnMaxTokensOk() (*string, bool)`

GetContinueInstructionOnMaxTokensOk returns a tuple with the ContinueInstructionOnMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueInstructionOnMaxTokens

`func (o *AIgencyMessageCompletionParameters) SetContinueInstructionOnMaxTokens(v string)`

SetContinueInstructionOnMaxTokens sets ContinueInstructionOnMaxTokens field to given value.

### HasContinueInstructionOnMaxTokens

`func (o *AIgencyMessageCompletionParameters) HasContinueInstructionOnMaxTokens() bool`

HasContinueInstructionOnMaxTokens returns a boolean if a field has been set.

### GetContinueOnMaxTokens

`func (o *AIgencyMessageCompletionParameters) GetContinueOnMaxTokens() bool`

GetContinueOnMaxTokens returns the ContinueOnMaxTokens field if non-nil, zero value otherwise.

### GetContinueOnMaxTokensOk

`func (o *AIgencyMessageCompletionParameters) GetContinueOnMaxTokensOk() (*bool, bool)`

GetContinueOnMaxTokensOk returns a tuple with the ContinueOnMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnMaxTokens

`func (o *AIgencyMessageCompletionParameters) SetContinueOnMaxTokens(v bool)`

SetContinueOnMaxTokens sets ContinueOnMaxTokens field to given value.

### HasContinueOnMaxTokens

`func (o *AIgencyMessageCompletionParameters) HasContinueOnMaxTokens() bool`

HasContinueOnMaxTokens returns a boolean if a field has been set.

### GetDeeprToolAutoActivated

`func (o *AIgencyMessageCompletionParameters) GetDeeprToolAutoActivated() bool`

GetDeeprToolAutoActivated returns the DeeprToolAutoActivated field if non-nil, zero value otherwise.

### GetDeeprToolAutoActivatedOk

`func (o *AIgencyMessageCompletionParameters) GetDeeprToolAutoActivatedOk() (*bool, bool)`

GetDeeprToolAutoActivatedOk returns a tuple with the DeeprToolAutoActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeeprToolAutoActivated

`func (o *AIgencyMessageCompletionParameters) SetDeeprToolAutoActivated(v bool)`

SetDeeprToolAutoActivated sets DeeprToolAutoActivated field to given value.

### HasDeeprToolAutoActivated

`func (o *AIgencyMessageCompletionParameters) HasDeeprToolAutoActivated() bool`

HasDeeprToolAutoActivated returns a boolean if a field has been set.

### GetModelHostLocation

`func (o *AIgencyMessageCompletionParameters) GetModelHostLocation() HostingLocation`

GetModelHostLocation returns the ModelHostLocation field if non-nil, zero value otherwise.

### GetModelHostLocationOk

`func (o *AIgencyMessageCompletionParameters) GetModelHostLocationOk() (*HostingLocation, bool)`

GetModelHostLocationOk returns a tuple with the ModelHostLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelHostLocation

`func (o *AIgencyMessageCompletionParameters) SetModelHostLocation(v HostingLocation)`

SetModelHostLocation sets ModelHostLocation field to given value.

### HasModelHostLocation

`func (o *AIgencyMessageCompletionParameters) HasModelHostLocation() bool`

HasModelHostLocation returns a boolean if a field has been set.

### GetModelParameters

`func (o *AIgencyMessageCompletionParameters) GetModelParameters() map[string]interface{}`

GetModelParameters returns the ModelParameters field if non-nil, zero value otherwise.

### GetModelParametersOk

`func (o *AIgencyMessageCompletionParameters) GetModelParametersOk() (*map[string]interface{}, bool)`

GetModelParametersOk returns a tuple with the ModelParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelParameters

`func (o *AIgencyMessageCompletionParameters) SetModelParameters(v map[string]interface{})`

SetModelParameters sets ModelParameters field to given value.

### HasModelParameters

`func (o *AIgencyMessageCompletionParameters) HasModelParameters() bool`

HasModelParameters returns a boolean if a field has been set.

### GetSelectedTools

`func (o *AIgencyMessageCompletionParameters) GetSelectedTools() []string`

GetSelectedTools returns the SelectedTools field if non-nil, zero value otherwise.

### GetSelectedToolsOk

`func (o *AIgencyMessageCompletionParameters) GetSelectedToolsOk() (*[]string, bool)`

GetSelectedToolsOk returns a tuple with the SelectedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedTools

`func (o *AIgencyMessageCompletionParameters) SetSelectedTools(v []string)`

SetSelectedTools sets SelectedTools field to given value.

### HasSelectedTools

`func (o *AIgencyMessageCompletionParameters) HasSelectedTools() bool`

HasSelectedTools returns a boolean if a field has been set.

### GetSetMessageHistoryIds

`func (o *AIgencyMessageCompletionParameters) GetSetMessageHistoryIds() []string`

GetSetMessageHistoryIds returns the SetMessageHistoryIds field if non-nil, zero value otherwise.

### GetSetMessageHistoryIdsOk

`func (o *AIgencyMessageCompletionParameters) GetSetMessageHistoryIdsOk() (*[]string, bool)`

GetSetMessageHistoryIdsOk returns a tuple with the SetMessageHistoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetMessageHistoryIds

`func (o *AIgencyMessageCompletionParameters) SetSetMessageHistoryIds(v []string)`

SetSetMessageHistoryIds sets SetMessageHistoryIds field to given value.

### HasSetMessageHistoryIds

`func (o *AIgencyMessageCompletionParameters) HasSetMessageHistoryIds() bool`

HasSetMessageHistoryIds returns a boolean if a field has been set.

### GetStream

`func (o *AIgencyMessageCompletionParameters) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AIgencyMessageCompletionParameters) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AIgencyMessageCompletionParameters) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AIgencyMessageCompletionParameters) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetToolConfigs

`func (o *AIgencyMessageCompletionParameters) GetToolConfigs() map[string]interface{}`

GetToolConfigs returns the ToolConfigs field if non-nil, zero value otherwise.

### GetToolConfigsOk

`func (o *AIgencyMessageCompletionParameters) GetToolConfigsOk() (*map[string]interface{}, bool)`

GetToolConfigsOk returns a tuple with the ToolConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfigs

`func (o *AIgencyMessageCompletionParameters) SetToolConfigs(v map[string]interface{})`

SetToolConfigs sets ToolConfigs field to given value.

### HasToolConfigs

`func (o *AIgencyMessageCompletionParameters) HasToolConfigs() bool`

HasToolConfigs returns a boolean if a field has been set.

### GetToolFunctionIds

`func (o *AIgencyMessageCompletionParameters) GetToolFunctionIds() []string`

GetToolFunctionIds returns the ToolFunctionIds field if non-nil, zero value otherwise.

### GetToolFunctionIdsOk

`func (o *AIgencyMessageCompletionParameters) GetToolFunctionIdsOk() (*[]string, bool)`

GetToolFunctionIdsOk returns a tuple with the ToolFunctionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolFunctionIds

`func (o *AIgencyMessageCompletionParameters) SetToolFunctionIds(v []string)`

SetToolFunctionIds sets ToolFunctionIds field to given value.

### HasToolFunctionIds

`func (o *AIgencyMessageCompletionParameters) HasToolFunctionIds() bool`

HasToolFunctionIds returns a boolean if a field has been set.

### GetUseChannelMessagesAsHistory

`func (o *AIgencyMessageCompletionParameters) GetUseChannelMessagesAsHistory() bool`

GetUseChannelMessagesAsHistory returns the UseChannelMessagesAsHistory field if non-nil, zero value otherwise.

### GetUseChannelMessagesAsHistoryOk

`func (o *AIgencyMessageCompletionParameters) GetUseChannelMessagesAsHistoryOk() (*bool, bool)`

GetUseChannelMessagesAsHistoryOk returns a tuple with the UseChannelMessagesAsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseChannelMessagesAsHistory

`func (o *AIgencyMessageCompletionParameters) SetUseChannelMessagesAsHistory(v bool)`

SetUseChannelMessagesAsHistory sets UseChannelMessagesAsHistory field to given value.

### HasUseChannelMessagesAsHistory

`func (o *AIgencyMessageCompletionParameters) HasUseChannelMessagesAsHistory() bool`

HasUseChannelMessagesAsHistory returns a boolean if a field has been set.

### GetUseTools

`func (o *AIgencyMessageCompletionParameters) GetUseTools() bool`

GetUseTools returns the UseTools field if non-nil, zero value otherwise.

### GetUseToolsOk

`func (o *AIgencyMessageCompletionParameters) GetUseToolsOk() (*bool, bool)`

GetUseToolsOk returns a tuple with the UseTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTools

`func (o *AIgencyMessageCompletionParameters) SetUseTools(v bool)`

SetUseTools sets UseTools field to given value.

### HasUseTools

`func (o *AIgencyMessageCompletionParameters) HasUseTools() bool`

HasUseTools returns a boolean if a field has been set.

### GetVarReplacements

`func (o *AIgencyMessageCompletionParameters) GetVarReplacements() map[string]string`

GetVarReplacements returns the VarReplacements field if non-nil, zero value otherwise.

### GetVarReplacementsOk

`func (o *AIgencyMessageCompletionParameters) GetVarReplacementsOk() (*map[string]string, bool)`

GetVarReplacementsOk returns a tuple with the VarReplacements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarReplacements

`func (o *AIgencyMessageCompletionParameters) SetVarReplacements(v map[string]string)`

SetVarReplacements sets VarReplacements field to given value.

### HasVarReplacements

`func (o *AIgencyMessageCompletionParameters) HasVarReplacements() bool`

HasVarReplacements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


