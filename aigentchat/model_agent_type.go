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
	"fmt"
)

// AgentType the model 'AgentType'
type AgentType string

// List of AgentType
const (
	AgentTypeNil AgentType = ""
	AgentTypeBasic AgentType = "basic"
	AgentTypeService AgentType = "service"
)

// All allowed values of AgentType enum
var AllowedAgentTypeEnumValues = []AgentType{
	"",
	"basic",
	"service",
}

func (v *AgentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentType(value)
	for _, existing := range AllowedAgentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentType", value)
}

// NewAgentTypeFromValue returns a pointer to a valid AgentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentTypeFromValue(v string) (*AgentType, error) {
	ev := AgentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentType: valid values are %v", v, AllowedAgentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentType) IsValid() bool {
	for _, existing := range AllowedAgentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentType value
func (v AgentType) Ptr() *AgentType {
	return &v
}

type NullableAgentType struct {
	value *AgentType
	isSet bool
}

func (v NullableAgentType) Get() *AgentType {
	return v.value
}

func (v *NullableAgentType) Set(val *AgentType) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentType) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentType(val *AgentType) *NullableAgentType {
	return &NullableAgentType{value: val, isSet: true}
}

func (v NullableAgentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

