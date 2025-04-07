/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.18.2
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// PromptVisibilityStates the model 'PromptVisibilityStates'
type PromptVisibilityStates string

// List of PromptVisibilityStates
const (
	PromptVisibilityPrivate PromptVisibilityStates = "private"
	PromptVisibilityOrganization PromptVisibilityStates = "org"
	PromptVisibilityPublic PromptVisibilityStates = "public"
	PromptVisibilityCurated PromptVisibilityStates = "curated"
)

// All allowed values of PromptVisibilityStates enum
var AllowedPromptVisibilityStatesEnumValues = []PromptVisibilityStates{
	"private",
	"org",
	"public",
	"curated",
}

func (v *PromptVisibilityStates) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PromptVisibilityStates(value)
	for _, existing := range AllowedPromptVisibilityStatesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PromptVisibilityStates", value)
}

// NewPromptVisibilityStatesFromValue returns a pointer to a valid PromptVisibilityStates
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromptVisibilityStatesFromValue(v string) (*PromptVisibilityStates, error) {
	ev := PromptVisibilityStates(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromptVisibilityStates: valid values are %v", v, AllowedPromptVisibilityStatesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromptVisibilityStates) IsValid() bool {
	for _, existing := range AllowedPromptVisibilityStatesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PromptVisibilityStates value
func (v PromptVisibilityStates) Ptr() *PromptVisibilityStates {
	return &v
}

type NullablePromptVisibilityStates struct {
	value *PromptVisibilityStates
	isSet bool
}

func (v NullablePromptVisibilityStates) Get() *PromptVisibilityStates {
	return v.value
}

func (v *NullablePromptVisibilityStates) Set(val *PromptVisibilityStates) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptVisibilityStates) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptVisibilityStates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptVisibilityStates(val *PromptVisibilityStates) *NullablePromptVisibilityStates {
	return &NullablePromptVisibilityStates{value: val, isSet: true}
}

func (v NullablePromptVisibilityStates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptVisibilityStates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

