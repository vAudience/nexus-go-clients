# ChatCompletionRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**AttachedFiles** | Pointer to **[]string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**ContinueInstructionOnMaxTokens** | Pointer to **string** |  | [optional] 
**ContinueOnMaxTokens** | Pointer to **bool** |  | [optional] 
**ExpireMessages** | Pointer to **bool** |  | [optional] 
**Message** | **string** |  | 
**MessageReferenceId** | Pointer to **string** |  | [optional] 
**MessageResponseToId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**SelectedTools** | Pointer to **[]string** |  | [optional] 
**ServiceChannel** | Pointer to **bool** |  | [optional] 
**SetMessageHistoryIds** | Pointer to **[]string** | If UseChannelMessagesAsHistory is false, this list of message IDs will be used as history, if empty, the history will be empty, ignored if UseChannelMessagesAsHistory is true | [optional] 
**UseChannelMessagesAsHistory** | Pointer to **bool** | If true, the channel messages will be used as history and SetMessageHistoryIds will be ignored | [optional] 
**UseSummaryService** | Pointer to **bool** |  | [optional] 
**UseTools** | Pointer to **bool** |  | [optional] 
**VarReplacements** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewChatCompletionRequestDto

`func NewChatCompletionRequestDto(agentId string, message string, ) *ChatCompletionRequestDto`

NewChatCompletionRequestDto instantiates a new ChatCompletionRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestDtoWithDefaults

`func NewChatCompletionRequestDtoWithDefaults() *ChatCompletionRequestDto`

NewChatCompletionRequestDtoWithDefaults instantiates a new ChatCompletionRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ChatCompletionRequestDto) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ChatCompletionRequestDto) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ChatCompletionRequestDto) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetAttachedFiles

`func (o *ChatCompletionRequestDto) GetAttachedFiles() []string`

GetAttachedFiles returns the AttachedFiles field if non-nil, zero value otherwise.

### GetAttachedFilesOk

`func (o *ChatCompletionRequestDto) GetAttachedFilesOk() (*[]string, bool)`

GetAttachedFilesOk returns a tuple with the AttachedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFiles

`func (o *ChatCompletionRequestDto) SetAttachedFiles(v []string)`

SetAttachedFiles sets AttachedFiles field to given value.

### HasAttachedFiles

`func (o *ChatCompletionRequestDto) HasAttachedFiles() bool`

HasAttachedFiles returns a boolean if a field has been set.

### GetChannelId

`func (o *ChatCompletionRequestDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ChatCompletionRequestDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ChatCompletionRequestDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ChatCompletionRequestDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetContinueInstructionOnMaxTokens

`func (o *ChatCompletionRequestDto) GetContinueInstructionOnMaxTokens() string`

GetContinueInstructionOnMaxTokens returns the ContinueInstructionOnMaxTokens field if non-nil, zero value otherwise.

### GetContinueInstructionOnMaxTokensOk

`func (o *ChatCompletionRequestDto) GetContinueInstructionOnMaxTokensOk() (*string, bool)`

GetContinueInstructionOnMaxTokensOk returns a tuple with the ContinueInstructionOnMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueInstructionOnMaxTokens

`func (o *ChatCompletionRequestDto) SetContinueInstructionOnMaxTokens(v string)`

SetContinueInstructionOnMaxTokens sets ContinueInstructionOnMaxTokens field to given value.

### HasContinueInstructionOnMaxTokens

`func (o *ChatCompletionRequestDto) HasContinueInstructionOnMaxTokens() bool`

HasContinueInstructionOnMaxTokens returns a boolean if a field has been set.

### GetContinueOnMaxTokens

`func (o *ChatCompletionRequestDto) GetContinueOnMaxTokens() bool`

GetContinueOnMaxTokens returns the ContinueOnMaxTokens field if non-nil, zero value otherwise.

### GetContinueOnMaxTokensOk

`func (o *ChatCompletionRequestDto) GetContinueOnMaxTokensOk() (*bool, bool)`

GetContinueOnMaxTokensOk returns a tuple with the ContinueOnMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnMaxTokens

`func (o *ChatCompletionRequestDto) SetContinueOnMaxTokens(v bool)`

SetContinueOnMaxTokens sets ContinueOnMaxTokens field to given value.

### HasContinueOnMaxTokens

`func (o *ChatCompletionRequestDto) HasContinueOnMaxTokens() bool`

HasContinueOnMaxTokens returns a boolean if a field has been set.

### GetExpireMessages

`func (o *ChatCompletionRequestDto) GetExpireMessages() bool`

GetExpireMessages returns the ExpireMessages field if non-nil, zero value otherwise.

### GetExpireMessagesOk

`func (o *ChatCompletionRequestDto) GetExpireMessagesOk() (*bool, bool)`

GetExpireMessagesOk returns a tuple with the ExpireMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireMessages

`func (o *ChatCompletionRequestDto) SetExpireMessages(v bool)`

SetExpireMessages sets ExpireMessages field to given value.

### HasExpireMessages

`func (o *ChatCompletionRequestDto) HasExpireMessages() bool`

HasExpireMessages returns a boolean if a field has been set.

### GetMessage

`func (o *ChatCompletionRequestDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatCompletionRequestDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatCompletionRequestDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageReferenceId

`func (o *ChatCompletionRequestDto) GetMessageReferenceId() string`

GetMessageReferenceId returns the MessageReferenceId field if non-nil, zero value otherwise.

### GetMessageReferenceIdOk

`func (o *ChatCompletionRequestDto) GetMessageReferenceIdOk() (*string, bool)`

GetMessageReferenceIdOk returns a tuple with the MessageReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReferenceId

`func (o *ChatCompletionRequestDto) SetMessageReferenceId(v string)`

SetMessageReferenceId sets MessageReferenceId field to given value.

### HasMessageReferenceId

`func (o *ChatCompletionRequestDto) HasMessageReferenceId() bool`

HasMessageReferenceId returns a boolean if a field has been set.

### GetMessageResponseToId

`func (o *ChatCompletionRequestDto) GetMessageResponseToId() string`

GetMessageResponseToId returns the MessageResponseToId field if non-nil, zero value otherwise.

### GetMessageResponseToIdOk

`func (o *ChatCompletionRequestDto) GetMessageResponseToIdOk() (*string, bool)`

GetMessageResponseToIdOk returns a tuple with the MessageResponseToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageResponseToId

`func (o *ChatCompletionRequestDto) SetMessageResponseToId(v string)`

SetMessageResponseToId sets MessageResponseToId field to given value.

### HasMessageResponseToId

`func (o *ChatCompletionRequestDto) HasMessageResponseToId() bool`

HasMessageResponseToId returns a boolean if a field has been set.

### GetParameters

`func (o *ChatCompletionRequestDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ChatCompletionRequestDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ChatCompletionRequestDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ChatCompletionRequestDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSelectedTools

`func (o *ChatCompletionRequestDto) GetSelectedTools() []string`

GetSelectedTools returns the SelectedTools field if non-nil, zero value otherwise.

### GetSelectedToolsOk

`func (o *ChatCompletionRequestDto) GetSelectedToolsOk() (*[]string, bool)`

GetSelectedToolsOk returns a tuple with the SelectedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedTools

`func (o *ChatCompletionRequestDto) SetSelectedTools(v []string)`

SetSelectedTools sets SelectedTools field to given value.

### HasSelectedTools

`func (o *ChatCompletionRequestDto) HasSelectedTools() bool`

HasSelectedTools returns a boolean if a field has been set.

### GetServiceChannel

`func (o *ChatCompletionRequestDto) GetServiceChannel() bool`

GetServiceChannel returns the ServiceChannel field if non-nil, zero value otherwise.

### GetServiceChannelOk

`func (o *ChatCompletionRequestDto) GetServiceChannelOk() (*bool, bool)`

GetServiceChannelOk returns a tuple with the ServiceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceChannel

`func (o *ChatCompletionRequestDto) SetServiceChannel(v bool)`

SetServiceChannel sets ServiceChannel field to given value.

### HasServiceChannel

`func (o *ChatCompletionRequestDto) HasServiceChannel() bool`

HasServiceChannel returns a boolean if a field has been set.

### GetSetMessageHistoryIds

`func (o *ChatCompletionRequestDto) GetSetMessageHistoryIds() []string`

GetSetMessageHistoryIds returns the SetMessageHistoryIds field if non-nil, zero value otherwise.

### GetSetMessageHistoryIdsOk

`func (o *ChatCompletionRequestDto) GetSetMessageHistoryIdsOk() (*[]string, bool)`

GetSetMessageHistoryIdsOk returns a tuple with the SetMessageHistoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetMessageHistoryIds

`func (o *ChatCompletionRequestDto) SetSetMessageHistoryIds(v []string)`

SetSetMessageHistoryIds sets SetMessageHistoryIds field to given value.

### HasSetMessageHistoryIds

`func (o *ChatCompletionRequestDto) HasSetMessageHistoryIds() bool`

HasSetMessageHistoryIds returns a boolean if a field has been set.

### GetUseChannelMessagesAsHistory

`func (o *ChatCompletionRequestDto) GetUseChannelMessagesAsHistory() bool`

GetUseChannelMessagesAsHistory returns the UseChannelMessagesAsHistory field if non-nil, zero value otherwise.

### GetUseChannelMessagesAsHistoryOk

`func (o *ChatCompletionRequestDto) GetUseChannelMessagesAsHistoryOk() (*bool, bool)`

GetUseChannelMessagesAsHistoryOk returns a tuple with the UseChannelMessagesAsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseChannelMessagesAsHistory

`func (o *ChatCompletionRequestDto) SetUseChannelMessagesAsHistory(v bool)`

SetUseChannelMessagesAsHistory sets UseChannelMessagesAsHistory field to given value.

### HasUseChannelMessagesAsHistory

`func (o *ChatCompletionRequestDto) HasUseChannelMessagesAsHistory() bool`

HasUseChannelMessagesAsHistory returns a boolean if a field has been set.

### GetUseSummaryService

`func (o *ChatCompletionRequestDto) GetUseSummaryService() bool`

GetUseSummaryService returns the UseSummaryService field if non-nil, zero value otherwise.

### GetUseSummaryServiceOk

`func (o *ChatCompletionRequestDto) GetUseSummaryServiceOk() (*bool, bool)`

GetUseSummaryServiceOk returns a tuple with the UseSummaryService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSummaryService

`func (o *ChatCompletionRequestDto) SetUseSummaryService(v bool)`

SetUseSummaryService sets UseSummaryService field to given value.

### HasUseSummaryService

`func (o *ChatCompletionRequestDto) HasUseSummaryService() bool`

HasUseSummaryService returns a boolean if a field has been set.

### GetUseTools

`func (o *ChatCompletionRequestDto) GetUseTools() bool`

GetUseTools returns the UseTools field if non-nil, zero value otherwise.

### GetUseToolsOk

`func (o *ChatCompletionRequestDto) GetUseToolsOk() (*bool, bool)`

GetUseToolsOk returns a tuple with the UseTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTools

`func (o *ChatCompletionRequestDto) SetUseTools(v bool)`

SetUseTools sets UseTools field to given value.

### HasUseTools

`func (o *ChatCompletionRequestDto) HasUseTools() bool`

HasUseTools returns a boolean if a field has been set.

### GetVarReplacements

`func (o *ChatCompletionRequestDto) GetVarReplacements() map[string]string`

GetVarReplacements returns the VarReplacements field if non-nil, zero value otherwise.

### GetVarReplacementsOk

`func (o *ChatCompletionRequestDto) GetVarReplacementsOk() (*map[string]string, bool)`

GetVarReplacementsOk returns a tuple with the VarReplacements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarReplacements

`func (o *ChatCompletionRequestDto) SetVarReplacements(v map[string]string)`

SetVarReplacements sets VarReplacements field to given value.

### HasVarReplacements

`func (o *ChatCompletionRequestDto) HasVarReplacements() bool`

HasVarReplacements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


