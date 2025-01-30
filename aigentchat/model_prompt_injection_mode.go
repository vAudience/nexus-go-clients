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

// PromptInjectionMode the model 'PromptInjectionMode'
type PromptInjectionMode string

// List of PromptInjectionMode
const (
	PromptInjectionModeAppend PromptInjectionMode = "append"
	PromptInjectionModePrepend PromptInjectionMode = "prepend"
	PromptInjectionModeReplace PromptInjectionMode = "replace"
)

// All allowed values of PromptInjectionMode enum
var AllowedPromptInjectionModeEnumValues = []PromptInjectionMode{
	"append",
	"prepend",
	"replace",
}

func (v *PromptInjectionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PromptInjectionMode(value)
	for _, existing := range AllowedPromptInjectionModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PromptInjectionMode", value)
}

// NewPromptInjectionModeFromValue returns a pointer to a valid PromptInjectionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromptInjectionModeFromValue(v string) (*PromptInjectionMode, error) {
	ev := PromptInjectionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromptInjectionMode: valid values are %v", v, AllowedPromptInjectionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromptInjectionMode) IsValid() bool {
	for _, existing := range AllowedPromptInjectionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PromptInjectionMode value
func (v PromptInjectionMode) Ptr() *PromptInjectionMode {
	return &v
}

type NullablePromptInjectionMode struct {
	value *PromptInjectionMode
	isSet bool
}

func (v NullablePromptInjectionMode) Get() *PromptInjectionMode {
	return v.value
}

func (v *NullablePromptInjectionMode) Set(val *PromptInjectionMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptInjectionMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptInjectionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptInjectionMode(val *PromptInjectionMode) *NullablePromptInjectionMode {
	return &NullablePromptInjectionMode{value: val, isSet: true}
}

func (v NullablePromptInjectionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptInjectionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

