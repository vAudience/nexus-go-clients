/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.12
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the AgentPromptRenderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPromptRenderDto{}

// AgentPromptRenderDto struct for AgentPromptRenderDto
type AgentPromptRenderDto struct {
	Content *string `json:"content,omitempty"`
	PromptId *string `json:"prompt_id,omitempty"`
	VarReplacements *map[string]string `json:"var_replacements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentPromptRenderDto AgentPromptRenderDto

// NewAgentPromptRenderDto instantiates a new AgentPromptRenderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPromptRenderDto() *AgentPromptRenderDto {
	this := AgentPromptRenderDto{}
	return &this
}

// NewAgentPromptRenderDtoWithDefaults instantiates a new AgentPromptRenderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPromptRenderDtoWithDefaults() *AgentPromptRenderDto {
	this := AgentPromptRenderDto{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AgentPromptRenderDto) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPromptRenderDto) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AgentPromptRenderDto) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AgentPromptRenderDto) SetContent(v string) {
	o.Content = &v
}

// GetPromptId returns the PromptId field value if set, zero value otherwise.
func (o *AgentPromptRenderDto) GetPromptId() string {
	if o == nil || IsNil(o.PromptId) {
		var ret string
		return ret
	}
	return *o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPromptRenderDto) GetPromptIdOk() (*string, bool) {
	if o == nil || IsNil(o.PromptId) {
		return nil, false
	}
	return o.PromptId, true
}

// HasPromptId returns a boolean if a field has been set.
func (o *AgentPromptRenderDto) HasPromptId() bool {
	if o != nil && !IsNil(o.PromptId) {
		return true
	}

	return false
}

// SetPromptId gets a reference to the given string and assigns it to the PromptId field.
func (o *AgentPromptRenderDto) SetPromptId(v string) {
	o.PromptId = &v
}

// GetVarReplacements returns the VarReplacements field value if set, zero value otherwise.
func (o *AgentPromptRenderDto) GetVarReplacements() map[string]string {
	if o == nil || IsNil(o.VarReplacements) {
		var ret map[string]string
		return ret
	}
	return *o.VarReplacements
}

// GetVarReplacementsOk returns a tuple with the VarReplacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPromptRenderDto) GetVarReplacementsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.VarReplacements) {
		return nil, false
	}
	return o.VarReplacements, true
}

// HasVarReplacements returns a boolean if a field has been set.
func (o *AgentPromptRenderDto) HasVarReplacements() bool {
	if o != nil && !IsNil(o.VarReplacements) {
		return true
	}

	return false
}

// SetVarReplacements gets a reference to the given map[string]string and assigns it to the VarReplacements field.
func (o *AgentPromptRenderDto) SetVarReplacements(v map[string]string) {
	o.VarReplacements = &v
}

func (o AgentPromptRenderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPromptRenderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.PromptId) {
		toSerialize["prompt_id"] = o.PromptId
	}
	if !IsNil(o.VarReplacements) {
		toSerialize["var_replacements"] = o.VarReplacements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentPromptRenderDto) UnmarshalJSON(data []byte) (err error) {
	varAgentPromptRenderDto := _AgentPromptRenderDto{}

	err = json.Unmarshal(data, &varAgentPromptRenderDto)

	if err != nil {
		return err
	}

	*o = AgentPromptRenderDto(varAgentPromptRenderDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "prompt_id")
		delete(additionalProperties, "var_replacements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentPromptRenderDto struct {
	value *AgentPromptRenderDto
	isSet bool
}

func (v NullableAgentPromptRenderDto) Get() *AgentPromptRenderDto {
	return v.value
}

func (v *NullableAgentPromptRenderDto) Set(val *AgentPromptRenderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPromptRenderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPromptRenderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPromptRenderDto(val *AgentPromptRenderDto) *NullableAgentPromptRenderDto {
	return &NullableAgentPromptRenderDto{value: val, isSet: true}
}

func (v NullableAgentPromptRenderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPromptRenderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


