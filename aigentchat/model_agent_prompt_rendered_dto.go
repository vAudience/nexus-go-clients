/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.13.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the AgentPromptRenderedDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPromptRenderedDto{}

// AgentPromptRenderedDto struct for AgentPromptRenderedDto
type AgentPromptRenderedDto struct {
	RenderedPrompt *string `json:"rendered_prompt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentPromptRenderedDto AgentPromptRenderedDto

// NewAgentPromptRenderedDto instantiates a new AgentPromptRenderedDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPromptRenderedDto() *AgentPromptRenderedDto {
	this := AgentPromptRenderedDto{}
	return &this
}

// NewAgentPromptRenderedDtoWithDefaults instantiates a new AgentPromptRenderedDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPromptRenderedDtoWithDefaults() *AgentPromptRenderedDto {
	this := AgentPromptRenderedDto{}
	return &this
}

// GetRenderedPrompt returns the RenderedPrompt field value if set, zero value otherwise.
func (o *AgentPromptRenderedDto) GetRenderedPrompt() string {
	if o == nil || IsNil(o.RenderedPrompt) {
		var ret string
		return ret
	}
	return *o.RenderedPrompt
}

// GetRenderedPromptOk returns a tuple with the RenderedPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPromptRenderedDto) GetRenderedPromptOk() (*string, bool) {
	if o == nil || IsNil(o.RenderedPrompt) {
		return nil, false
	}
	return o.RenderedPrompt, true
}

// HasRenderedPrompt returns a boolean if a field has been set.
func (o *AgentPromptRenderedDto) HasRenderedPrompt() bool {
	if o != nil && !IsNil(o.RenderedPrompt) {
		return true
	}

	return false
}

// SetRenderedPrompt gets a reference to the given string and assigns it to the RenderedPrompt field.
func (o *AgentPromptRenderedDto) SetRenderedPrompt(v string) {
	o.RenderedPrompt = &v
}

func (o AgentPromptRenderedDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPromptRenderedDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RenderedPrompt) {
		toSerialize["rendered_prompt"] = o.RenderedPrompt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentPromptRenderedDto) UnmarshalJSON(data []byte) (err error) {
	varAgentPromptRenderedDto := _AgentPromptRenderedDto{}

	err = json.Unmarshal(data, &varAgentPromptRenderedDto)

	if err != nil {
		return err
	}

	*o = AgentPromptRenderedDto(varAgentPromptRenderedDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rendered_prompt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentPromptRenderedDto struct {
	value *AgentPromptRenderedDto
	isSet bool
}

func (v NullableAgentPromptRenderedDto) Get() *AgentPromptRenderedDto {
	return v.value
}

func (v *NullableAgentPromptRenderedDto) Set(val *AgentPromptRenderedDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPromptRenderedDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPromptRenderedDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPromptRenderedDto(val *AgentPromptRenderedDto) *NullableAgentPromptRenderedDto {
	return &NullableAgentPromptRenderedDto{value: val, isSet: true}
}

func (v NullableAgentPromptRenderedDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPromptRenderedDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


