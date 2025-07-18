/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.19.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// AIModelConstraintDirection the model 'AIModelConstraintDirection'
type AIModelConstraintDirection string

// List of AIModelConstraintDirection
const (
	AIModelConstraintDirectionInput AIModelConstraintDirection = "input"
	AIModelConstraintDirectionOutput AIModelConstraintDirection = "output"
)

// All allowed values of AIModelConstraintDirection enum
var AllowedAIModelConstraintDirectionEnumValues = []AIModelConstraintDirection{
	"input",
	"output",
}

func (v *AIModelConstraintDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AIModelConstraintDirection(value)
	for _, existing := range AllowedAIModelConstraintDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AIModelConstraintDirection", value)
}

// NewAIModelConstraintDirectionFromValue returns a pointer to a valid AIModelConstraintDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAIModelConstraintDirectionFromValue(v string) (*AIModelConstraintDirection, error) {
	ev := AIModelConstraintDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AIModelConstraintDirection: valid values are %v", v, AllowedAIModelConstraintDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AIModelConstraintDirection) IsValid() bool {
	for _, existing := range AllowedAIModelConstraintDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIModelConstraintDirection value
func (v AIModelConstraintDirection) Ptr() *AIModelConstraintDirection {
	return &v
}

type NullableAIModelConstraintDirection struct {
	value *AIModelConstraintDirection
	isSet bool
}

func (v NullableAIModelConstraintDirection) Get() *AIModelConstraintDirection {
	return v.value
}

func (v *NullableAIModelConstraintDirection) Set(val *AIModelConstraintDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelConstraintDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelConstraintDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelConstraintDirection(val *AIModelConstraintDirection) *NullableAIModelConstraintDirection {
	return &NullableAIModelConstraintDirection{value: val, isSet: true}
}

func (v NullableAIModelConstraintDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelConstraintDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

