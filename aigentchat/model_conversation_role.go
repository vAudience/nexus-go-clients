/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.17
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// ConversationRole the model 'ConversationRole'
type ConversationRole string

// List of ConversationRole
const (
	ConversationRoleUnknown ConversationRole = "unknown"
	ConversationRoleUser ConversationRole = "user"
	ConversationRoleSystem ConversationRole = "system"
	ConversationRoleAssistant ConversationRole = "assistant"
	ConversationRoleAIgent ConversationRole = "aigent"
)

// All allowed values of ConversationRole enum
var AllowedConversationRoleEnumValues = []ConversationRole{
	"unknown",
	"user",
	"system",
	"assistant",
	"aigent",
}

func (v *ConversationRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConversationRole(value)
	for _, existing := range AllowedConversationRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConversationRole", value)
}

// NewConversationRoleFromValue returns a pointer to a valid ConversationRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConversationRoleFromValue(v string) (*ConversationRole, error) {
	ev := ConversationRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConversationRole: valid values are %v", v, AllowedConversationRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConversationRole) IsValid() bool {
	for _, existing := range AllowedConversationRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConversationRole value
func (v ConversationRole) Ptr() *ConversationRole {
	return &v
}

type NullableConversationRole struct {
	value *ConversationRole
	isSet bool
}

func (v NullableConversationRole) Get() *ConversationRole {
	return v.value
}

func (v *NullableConversationRole) Set(val *ConversationRole) {
	v.value = val
	v.isSet = true
}

func (v NullableConversationRole) IsSet() bool {
	return v.isSet
}

func (v *NullableConversationRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversationRole(val *ConversationRole) *NullableConversationRole {
	return &NullableConversationRole{value: val, isSet: true}
}

func (v NullableConversationRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversationRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

