/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.14.1
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the ChatCompletionRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionRequestDto{}

// ChatCompletionRequestDto struct for ChatCompletionRequestDto
type ChatCompletionRequestDto struct {
	AgentId string `json:"agent_id"`
	AttachedTemporaryFiles []string `json:"attached_temporary_files,omitempty"`
	ChannelId *string `json:"channel_id,omitempty"`
	ContinueInstructionOnMaxTokens *string `json:"continue_instruction_on_max_tokens,omitempty"`
	ContinueOnMaxTokens *bool `json:"continue_on_max_tokens,omitempty"`
	ExpireMessages *bool `json:"expire_messages,omitempty"`
	Message string `json:"message"`
	MessageReferenceId *string `json:"message_reference_id,omitempty"`
	MessageResponseToId *string `json:"message_response_to_id,omitempty"`
	MissionId *string `json:"mission_id,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// If UseChannelMessagesAsHistory is false, this list of message IDs will be used as history, if empty, the history will be empty, ignored if UseChannelMessagesAsHistory is true
	SetMessageHistoryIds []string `json:"set_message_history_ids,omitempty"`
	// If true, the channel messages will be used as history and SetMessageHistoryIds will be ignored
	UseChannelMessagesAsHistory *bool `json:"use_channel_messages_as_history,omitempty"`
	UseTools *bool `json:"use_tools,omitempty"`
	VarReplacements *map[string]string `json:"var_replacements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatCompletionRequestDto ChatCompletionRequestDto

// NewChatCompletionRequestDto instantiates a new ChatCompletionRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionRequestDto(agentId string, message string) *ChatCompletionRequestDto {
	this := ChatCompletionRequestDto{}
	this.AgentId = agentId
	this.Message = message
	return &this
}

// NewChatCompletionRequestDtoWithDefaults instantiates a new ChatCompletionRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionRequestDtoWithDefaults() *ChatCompletionRequestDto {
	this := ChatCompletionRequestDto{}
	return &this
}

// GetAgentId returns the AgentId field value
func (o *ChatCompletionRequestDto) GetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *ChatCompletionRequestDto) SetAgentId(v string) {
	o.AgentId = v
}

// GetAttachedTemporaryFiles returns the AttachedTemporaryFiles field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetAttachedTemporaryFiles() []string {
	if o == nil || IsNil(o.AttachedTemporaryFiles) {
		var ret []string
		return ret
	}
	return o.AttachedTemporaryFiles
}

// GetAttachedTemporaryFilesOk returns a tuple with the AttachedTemporaryFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetAttachedTemporaryFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.AttachedTemporaryFiles) {
		return nil, false
	}
	return o.AttachedTemporaryFiles, true
}

// HasAttachedTemporaryFiles returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasAttachedTemporaryFiles() bool {
	if o != nil && !IsNil(o.AttachedTemporaryFiles) {
		return true
	}

	return false
}

// SetAttachedTemporaryFiles gets a reference to the given []string and assigns it to the AttachedTemporaryFiles field.
func (o *ChatCompletionRequestDto) SetAttachedTemporaryFiles(v []string) {
	o.AttachedTemporaryFiles = v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *ChatCompletionRequestDto) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetContinueInstructionOnMaxTokens returns the ContinueInstructionOnMaxTokens field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetContinueInstructionOnMaxTokens() string {
	if o == nil || IsNil(o.ContinueInstructionOnMaxTokens) {
		var ret string
		return ret
	}
	return *o.ContinueInstructionOnMaxTokens
}

// GetContinueInstructionOnMaxTokensOk returns a tuple with the ContinueInstructionOnMaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetContinueInstructionOnMaxTokensOk() (*string, bool) {
	if o == nil || IsNil(o.ContinueInstructionOnMaxTokens) {
		return nil, false
	}
	return o.ContinueInstructionOnMaxTokens, true
}

// HasContinueInstructionOnMaxTokens returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasContinueInstructionOnMaxTokens() bool {
	if o != nil && !IsNil(o.ContinueInstructionOnMaxTokens) {
		return true
	}

	return false
}

// SetContinueInstructionOnMaxTokens gets a reference to the given string and assigns it to the ContinueInstructionOnMaxTokens field.
func (o *ChatCompletionRequestDto) SetContinueInstructionOnMaxTokens(v string) {
	o.ContinueInstructionOnMaxTokens = &v
}

// GetContinueOnMaxTokens returns the ContinueOnMaxTokens field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetContinueOnMaxTokens() bool {
	if o == nil || IsNil(o.ContinueOnMaxTokens) {
		var ret bool
		return ret
	}
	return *o.ContinueOnMaxTokens
}

// GetContinueOnMaxTokensOk returns a tuple with the ContinueOnMaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetContinueOnMaxTokensOk() (*bool, bool) {
	if o == nil || IsNil(o.ContinueOnMaxTokens) {
		return nil, false
	}
	return o.ContinueOnMaxTokens, true
}

// HasContinueOnMaxTokens returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasContinueOnMaxTokens() bool {
	if o != nil && !IsNil(o.ContinueOnMaxTokens) {
		return true
	}

	return false
}

// SetContinueOnMaxTokens gets a reference to the given bool and assigns it to the ContinueOnMaxTokens field.
func (o *ChatCompletionRequestDto) SetContinueOnMaxTokens(v bool) {
	o.ContinueOnMaxTokens = &v
}

// GetExpireMessages returns the ExpireMessages field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetExpireMessages() bool {
	if o == nil || IsNil(o.ExpireMessages) {
		var ret bool
		return ret
	}
	return *o.ExpireMessages
}

// GetExpireMessagesOk returns a tuple with the ExpireMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetExpireMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpireMessages) {
		return nil, false
	}
	return o.ExpireMessages, true
}

// HasExpireMessages returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasExpireMessages() bool {
	if o != nil && !IsNil(o.ExpireMessages) {
		return true
	}

	return false
}

// SetExpireMessages gets a reference to the given bool and assigns it to the ExpireMessages field.
func (o *ChatCompletionRequestDto) SetExpireMessages(v bool) {
	o.ExpireMessages = &v
}

// GetMessage returns the Message field value
func (o *ChatCompletionRequestDto) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ChatCompletionRequestDto) SetMessage(v string) {
	o.Message = v
}

// GetMessageReferenceId returns the MessageReferenceId field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetMessageReferenceId() string {
	if o == nil || IsNil(o.MessageReferenceId) {
		var ret string
		return ret
	}
	return *o.MessageReferenceId
}

// GetMessageReferenceIdOk returns a tuple with the MessageReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetMessageReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageReferenceId) {
		return nil, false
	}
	return o.MessageReferenceId, true
}

// HasMessageReferenceId returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasMessageReferenceId() bool {
	if o != nil && !IsNil(o.MessageReferenceId) {
		return true
	}

	return false
}

// SetMessageReferenceId gets a reference to the given string and assigns it to the MessageReferenceId field.
func (o *ChatCompletionRequestDto) SetMessageReferenceId(v string) {
	o.MessageReferenceId = &v
}

// GetMessageResponseToId returns the MessageResponseToId field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetMessageResponseToId() string {
	if o == nil || IsNil(o.MessageResponseToId) {
		var ret string
		return ret
	}
	return *o.MessageResponseToId
}

// GetMessageResponseToIdOk returns a tuple with the MessageResponseToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetMessageResponseToIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageResponseToId) {
		return nil, false
	}
	return o.MessageResponseToId, true
}

// HasMessageResponseToId returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasMessageResponseToId() bool {
	if o != nil && !IsNil(o.MessageResponseToId) {
		return true
	}

	return false
}

// SetMessageResponseToId gets a reference to the given string and assigns it to the MessageResponseToId field.
func (o *ChatCompletionRequestDto) SetMessageResponseToId(v string) {
	o.MessageResponseToId = &v
}

// GetMissionId returns the MissionId field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetMissionId() string {
	if o == nil || IsNil(o.MissionId) {
		var ret string
		return ret
	}
	return *o.MissionId
}

// GetMissionIdOk returns a tuple with the MissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetMissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.MissionId) {
		return nil, false
	}
	return o.MissionId, true
}

// HasMissionId returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasMissionId() bool {
	if o != nil && !IsNil(o.MissionId) {
		return true
	}

	return false
}

// SetMissionId gets a reference to the given string and assigns it to the MissionId field.
func (o *ChatCompletionRequestDto) SetMissionId(v string) {
	o.MissionId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *ChatCompletionRequestDto) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetSetMessageHistoryIds returns the SetMessageHistoryIds field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetSetMessageHistoryIds() []string {
	if o == nil || IsNil(o.SetMessageHistoryIds) {
		var ret []string
		return ret
	}
	return o.SetMessageHistoryIds
}

// GetSetMessageHistoryIdsOk returns a tuple with the SetMessageHistoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetSetMessageHistoryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SetMessageHistoryIds) {
		return nil, false
	}
	return o.SetMessageHistoryIds, true
}

// HasSetMessageHistoryIds returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasSetMessageHistoryIds() bool {
	if o != nil && !IsNil(o.SetMessageHistoryIds) {
		return true
	}

	return false
}

// SetSetMessageHistoryIds gets a reference to the given []string and assigns it to the SetMessageHistoryIds field.
func (o *ChatCompletionRequestDto) SetSetMessageHistoryIds(v []string) {
	o.SetMessageHistoryIds = v
}

// GetUseChannelMessagesAsHistory returns the UseChannelMessagesAsHistory field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetUseChannelMessagesAsHistory() bool {
	if o == nil || IsNil(o.UseChannelMessagesAsHistory) {
		var ret bool
		return ret
	}
	return *o.UseChannelMessagesAsHistory
}

// GetUseChannelMessagesAsHistoryOk returns a tuple with the UseChannelMessagesAsHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetUseChannelMessagesAsHistoryOk() (*bool, bool) {
	if o == nil || IsNil(o.UseChannelMessagesAsHistory) {
		return nil, false
	}
	return o.UseChannelMessagesAsHistory, true
}

// HasUseChannelMessagesAsHistory returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasUseChannelMessagesAsHistory() bool {
	if o != nil && !IsNil(o.UseChannelMessagesAsHistory) {
		return true
	}

	return false
}

// SetUseChannelMessagesAsHistory gets a reference to the given bool and assigns it to the UseChannelMessagesAsHistory field.
func (o *ChatCompletionRequestDto) SetUseChannelMessagesAsHistory(v bool) {
	o.UseChannelMessagesAsHistory = &v
}

// GetUseTools returns the UseTools field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetUseTools() bool {
	if o == nil || IsNil(o.UseTools) {
		var ret bool
		return ret
	}
	return *o.UseTools
}

// GetUseToolsOk returns a tuple with the UseTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetUseToolsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTools) {
		return nil, false
	}
	return o.UseTools, true
}

// HasUseTools returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasUseTools() bool {
	if o != nil && !IsNil(o.UseTools) {
		return true
	}

	return false
}

// SetUseTools gets a reference to the given bool and assigns it to the UseTools field.
func (o *ChatCompletionRequestDto) SetUseTools(v bool) {
	o.UseTools = &v
}

// GetVarReplacements returns the VarReplacements field value if set, zero value otherwise.
func (o *ChatCompletionRequestDto) GetVarReplacements() map[string]string {
	if o == nil || IsNil(o.VarReplacements) {
		var ret map[string]string
		return ret
	}
	return *o.VarReplacements
}

// GetVarReplacementsOk returns a tuple with the VarReplacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestDto) GetVarReplacementsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.VarReplacements) {
		return nil, false
	}
	return o.VarReplacements, true
}

// HasVarReplacements returns a boolean if a field has been set.
func (o *ChatCompletionRequestDto) HasVarReplacements() bool {
	if o != nil && !IsNil(o.VarReplacements) {
		return true
	}

	return false
}

// SetVarReplacements gets a reference to the given map[string]string and assigns it to the VarReplacements field.
func (o *ChatCompletionRequestDto) SetVarReplacements(v map[string]string) {
	o.VarReplacements = &v
}

func (o ChatCompletionRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent_id"] = o.AgentId
	if !IsNil(o.AttachedTemporaryFiles) {
		toSerialize["attached_temporary_files"] = o.AttachedTemporaryFiles
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.ContinueInstructionOnMaxTokens) {
		toSerialize["continue_instruction_on_max_tokens"] = o.ContinueInstructionOnMaxTokens
	}
	if !IsNil(o.ContinueOnMaxTokens) {
		toSerialize["continue_on_max_tokens"] = o.ContinueOnMaxTokens
	}
	if !IsNil(o.ExpireMessages) {
		toSerialize["expire_messages"] = o.ExpireMessages
	}
	toSerialize["message"] = o.Message
	if !IsNil(o.MessageReferenceId) {
		toSerialize["message_reference_id"] = o.MessageReferenceId
	}
	if !IsNil(o.MessageResponseToId) {
		toSerialize["message_response_to_id"] = o.MessageResponseToId
	}
	if !IsNil(o.MissionId) {
		toSerialize["mission_id"] = o.MissionId
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.SetMessageHistoryIds) {
		toSerialize["set_message_history_ids"] = o.SetMessageHistoryIds
	}
	if !IsNil(o.UseChannelMessagesAsHistory) {
		toSerialize["use_channel_messages_as_history"] = o.UseChannelMessagesAsHistory
	}
	if !IsNil(o.UseTools) {
		toSerialize["use_tools"] = o.UseTools
	}
	if !IsNil(o.VarReplacements) {
		toSerialize["var_replacements"] = o.VarReplacements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatCompletionRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent_id",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varChatCompletionRequestDto := _ChatCompletionRequestDto{}

	err = json.Unmarshal(data, &varChatCompletionRequestDto)

	if err != nil {
		return err
	}

	*o = ChatCompletionRequestDto(varChatCompletionRequestDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agent_id")
		delete(additionalProperties, "attached_temporary_files")
		delete(additionalProperties, "channel_id")
		delete(additionalProperties, "continue_instruction_on_max_tokens")
		delete(additionalProperties, "continue_on_max_tokens")
		delete(additionalProperties, "expire_messages")
		delete(additionalProperties, "message")
		delete(additionalProperties, "message_reference_id")
		delete(additionalProperties, "message_response_to_id")
		delete(additionalProperties, "mission_id")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "set_message_history_ids")
		delete(additionalProperties, "use_channel_messages_as_history")
		delete(additionalProperties, "use_tools")
		delete(additionalProperties, "var_replacements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatCompletionRequestDto struct {
	value *ChatCompletionRequestDto
	isSet bool
}

func (v NullableChatCompletionRequestDto) Get() *ChatCompletionRequestDto {
	return v.value
}

func (v *NullableChatCompletionRequestDto) Set(val *ChatCompletionRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequestDto(val *ChatCompletionRequestDto) *NullableChatCompletionRequestDto {
	return &NullableChatCompletionRequestDto{value: val, isSet: true}
}

func (v NullableChatCompletionRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


