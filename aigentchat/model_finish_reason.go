/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.17.2
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// FinishReason the model 'FinishReason'
type FinishReason string

// List of FinishReason
const (
	FinishReasonNil FinishReason = ""
	FinishReasonError FinishReason = "error"
	FinishReasonContentFilter FinishReason = "content_filter"
	FinishReasonMaxTokens FinishReason = "max_tokens"
	FinishReasonStop FinishReason = "stop"
	FinishReasonToolCalls FinishReason = "tool_calls"
	FinishReasonUnknown FinishReason = "unknown"
)

// All allowed values of FinishReason enum
var AllowedFinishReasonEnumValues = []FinishReason{
	"",
	"error",
	"content_filter",
	"max_tokens",
	"stop",
	"tool_calls",
	"unknown",
}

func (v *FinishReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FinishReason(value)
	for _, existing := range AllowedFinishReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FinishReason", value)
}

// NewFinishReasonFromValue returns a pointer to a valid FinishReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFinishReasonFromValue(v string) (*FinishReason, error) {
	ev := FinishReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FinishReason: valid values are %v", v, AllowedFinishReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FinishReason) IsValid() bool {
	for _, existing := range AllowedFinishReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FinishReason value
func (v FinishReason) Ptr() *FinishReason {
	return &v
}

type NullableFinishReason struct {
	value *FinishReason
	isSet bool
}

func (v NullableFinishReason) Get() *FinishReason {
	return v.value
}

func (v *NullableFinishReason) Set(val *FinishReason) {
	v.value = val
	v.isSet = true
}

func (v NullableFinishReason) IsSet() bool {
	return v.isSet
}

func (v *NullableFinishReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinishReason(val *FinishReason) *NullableFinishReason {
	return &NullableFinishReason{value: val, isSet: true}
}

func (v NullableFinishReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinishReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

