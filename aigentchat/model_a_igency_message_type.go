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

// AIgencyMessageType the model 'AIgencyMessageType'
type AIgencyMessageType string

// List of AIgencyMessageType
const (
	AIgencyMessageTypeMessage AIgencyMessageType = "message"
	AIgencyMessageTypeStateUpdate AIgencyMessageType = "stateUpdate"
	AIgencyMessageTypeDelta AIgencyMessageType = "delta"
	AIgencyMessageTypeError AIgencyMessageType = "error"
)

// All allowed values of AIgencyMessageType enum
var AllowedAIgencyMessageTypeEnumValues = []AIgencyMessageType{
	"message",
	"stateUpdate",
	"delta",
	"error",
}

func (v *AIgencyMessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AIgencyMessageType(value)
	for _, existing := range AllowedAIgencyMessageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AIgencyMessageType", value)
}

// NewAIgencyMessageTypeFromValue returns a pointer to a valid AIgencyMessageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAIgencyMessageTypeFromValue(v string) (*AIgencyMessageType, error) {
	ev := AIgencyMessageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AIgencyMessageType: valid values are %v", v, AllowedAIgencyMessageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AIgencyMessageType) IsValid() bool {
	for _, existing := range AllowedAIgencyMessageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIgencyMessageType value
func (v AIgencyMessageType) Ptr() *AIgencyMessageType {
	return &v
}

type NullableAIgencyMessageType struct {
	value *AIgencyMessageType
	isSet bool
}

func (v NullableAIgencyMessageType) Get() *AIgencyMessageType {
	return v.value
}

func (v *NullableAIgencyMessageType) Set(val *AIgencyMessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableAIgencyMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableAIgencyMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIgencyMessageType(val *AIgencyMessageType) *NullableAIgencyMessageType {
	return &NullableAIgencyMessageType{value: val, isSet: true}
}

func (v NullableAIgencyMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIgencyMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

