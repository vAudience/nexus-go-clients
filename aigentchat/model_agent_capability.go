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

// AgentCapability the model 'AgentCapability'
type AgentCapability string

// List of AgentCapability
const (
	AgentCapabilityChat AgentCapability = "chat"
	AgentCapabilityEmbedding AgentCapability = "embedding"
	AgentCapabilityImage AgentCapability = "image"
)

// All allowed values of AgentCapability enum
var AllowedAgentCapabilityEnumValues = []AgentCapability{
	"chat",
	"embedding",
	"image",
}

func (v *AgentCapability) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentCapability(value)
	for _, existing := range AllowedAgentCapabilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentCapability", value)
}

// NewAgentCapabilityFromValue returns a pointer to a valid AgentCapability
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentCapabilityFromValue(v string) (*AgentCapability, error) {
	ev := AgentCapability(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentCapability: valid values are %v", v, AllowedAgentCapabilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentCapability) IsValid() bool {
	for _, existing := range AllowedAgentCapabilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentCapability value
func (v AgentCapability) Ptr() *AgentCapability {
	return &v
}

type NullableAgentCapability struct {
	value *AgentCapability
	isSet bool
}

func (v NullableAgentCapability) Get() *AgentCapability {
	return v.value
}

func (v *NullableAgentCapability) Set(val *AgentCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentCapability(val *AgentCapability) *NullableAgentCapability {
	return &NullableAgentCapability{value: val, isSet: true}
}

func (v NullableAgentCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

