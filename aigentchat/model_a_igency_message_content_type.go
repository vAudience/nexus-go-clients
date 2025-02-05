/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.0
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// AIgencyMessageContentType the model 'AIgencyMessageContentType'
type AIgencyMessageContentType string

// List of AIgencyMessageContentType
const (
	AIgencyMessageContentTypeText AIgencyMessageContentType = "text"
	AIgencyMessageContentTypeFile AIgencyMessageContentType = "file"
)

// All allowed values of AIgencyMessageContentType enum
var AllowedAIgencyMessageContentTypeEnumValues = []AIgencyMessageContentType{
	"text",
	"file",
}

func (v *AIgencyMessageContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AIgencyMessageContentType(value)
	for _, existing := range AllowedAIgencyMessageContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AIgencyMessageContentType", value)
}

// NewAIgencyMessageContentTypeFromValue returns a pointer to a valid AIgencyMessageContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAIgencyMessageContentTypeFromValue(v string) (*AIgencyMessageContentType, error) {
	ev := AIgencyMessageContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AIgencyMessageContentType: valid values are %v", v, AllowedAIgencyMessageContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AIgencyMessageContentType) IsValid() bool {
	for _, existing := range AllowedAIgencyMessageContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIgencyMessageContentType value
func (v AIgencyMessageContentType) Ptr() *AIgencyMessageContentType {
	return &v
}

type NullableAIgencyMessageContentType struct {
	value *AIgencyMessageContentType
	isSet bool
}

func (v NullableAIgencyMessageContentType) Get() *AIgencyMessageContentType {
	return v.value
}

func (v *NullableAIgencyMessageContentType) Set(val *AIgencyMessageContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableAIgencyMessageContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableAIgencyMessageContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIgencyMessageContentType(val *AIgencyMessageContentType) *NullableAIgencyMessageContentType {
	return &NullableAIgencyMessageContentType{value: val, isSet: true}
}

func (v NullableAIgencyMessageContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIgencyMessageContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

