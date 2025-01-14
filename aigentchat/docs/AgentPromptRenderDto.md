# AgentPromptRenderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**PromptId** | Pointer to **string** |  | [optional] 
**VarReplacements** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgentPromptRenderDto

`func NewAgentPromptRenderDto() *AgentPromptRenderDto`

NewAgentPromptRenderDto instantiates a new AgentPromptRenderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPromptRenderDtoWithDefaults

`func NewAgentPromptRenderDtoWithDefaults() *AgentPromptRenderDto`

NewAgentPromptRenderDtoWithDefaults instantiates a new AgentPromptRenderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *AgentPromptRenderDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AgentPromptRenderDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AgentPromptRenderDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AgentPromptRenderDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPromptId

`func (o *AgentPromptRenderDto) GetPromptId() string`

GetPromptId returns the PromptId field if non-nil, zero value otherwise.

### GetPromptIdOk

`func (o *AgentPromptRenderDto) GetPromptIdOk() (*string, bool)`

GetPromptIdOk returns a tuple with the PromptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptId

`func (o *AgentPromptRenderDto) SetPromptId(v string)`

SetPromptId sets PromptId field to given value.

### HasPromptId

`func (o *AgentPromptRenderDto) HasPromptId() bool`

HasPromptId returns a boolean if a field has been set.

### GetVarReplacements

`func (o *AgentPromptRenderDto) GetVarReplacements() map[string]string`

GetVarReplacements returns the VarReplacements field if non-nil, zero value otherwise.

### GetVarReplacementsOk

`func (o *AgentPromptRenderDto) GetVarReplacementsOk() (*map[string]string, bool)`

GetVarReplacementsOk returns a tuple with the VarReplacements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarReplacements

`func (o *AgentPromptRenderDto) SetVarReplacements(v map[string]string)`

SetVarReplacements sets VarReplacements field to given value.

### HasVarReplacements

`func (o *AgentPromptRenderDto) HasVarReplacements() bool`

HasVarReplacements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


