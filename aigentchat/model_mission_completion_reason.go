/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.18.0
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// MissionCompletionReason the model 'MissionCompletionReason'
type MissionCompletionReason string

// List of MissionCompletionReason
const (
	MissionCompletionReasonNil MissionCompletionReason = ""
	MissionCompletionReasonCanceled MissionCompletionReason = "canceled"
	MissionCompletionReasonCompleted MissionCompletionReason = "completed"
	MissionCompletionReasonError MissionCompletionReason = "error"
)

// All allowed values of MissionCompletionReason enum
var AllowedMissionCompletionReasonEnumValues = []MissionCompletionReason{
	"",
	"canceled",
	"completed",
	"error",
}

func (v *MissionCompletionReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MissionCompletionReason(value)
	for _, existing := range AllowedMissionCompletionReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MissionCompletionReason", value)
}

// NewMissionCompletionReasonFromValue returns a pointer to a valid MissionCompletionReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMissionCompletionReasonFromValue(v string) (*MissionCompletionReason, error) {
	ev := MissionCompletionReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MissionCompletionReason: valid values are %v", v, AllowedMissionCompletionReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MissionCompletionReason) IsValid() bool {
	for _, existing := range AllowedMissionCompletionReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MissionCompletionReason value
func (v MissionCompletionReason) Ptr() *MissionCompletionReason {
	return &v
}

type NullableMissionCompletionReason struct {
	value *MissionCompletionReason
	isSet bool
}

func (v NullableMissionCompletionReason) Get() *MissionCompletionReason {
	return v.value
}

func (v *NullableMissionCompletionReason) Set(val *MissionCompletionReason) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionCompletionReason) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionCompletionReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionCompletionReason(val *MissionCompletionReason) *NullableMissionCompletionReason {
	return &NullableMissionCompletionReason{value: val, isSet: true}
}

func (v NullableMissionCompletionReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionCompletionReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

