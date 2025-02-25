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

// MissionStatus the model 'MissionStatus'
type MissionStatus string

// List of MissionStatus
const (
	MissionStatusCreated MissionStatus = "created"
	MissionStatusStarted MissionStatus = "started"
	MissionStatusPaused MissionStatus = "paused"
	MissionStatusResumed MissionStatus = "resumed"
	MissionStatusCanceled MissionStatus = "canceled"
	MissionStatusFailed MissionStatus = "failed"
	MissionStatusCompleted MissionStatus = "completed"
)

// All allowed values of MissionStatus enum
var AllowedMissionStatusEnumValues = []MissionStatus{
	"created",
	"started",
	"paused",
	"resumed",
	"canceled",
	"failed",
	"completed",
}

func (v *MissionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MissionStatus(value)
	for _, existing := range AllowedMissionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MissionStatus", value)
}

// NewMissionStatusFromValue returns a pointer to a valid MissionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMissionStatusFromValue(v string) (*MissionStatus, error) {
	ev := MissionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MissionStatus: valid values are %v", v, AllowedMissionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MissionStatus) IsValid() bool {
	for _, existing := range AllowedMissionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MissionStatus value
func (v MissionStatus) Ptr() *MissionStatus {
	return &v
}

type NullableMissionStatus struct {
	value *MissionStatus
	isSet bool
}

func (v NullableMissionStatus) Get() *MissionStatus {
	return v.value
}

func (v *NullableMissionStatus) Set(val *MissionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionStatus(val *MissionStatus) *NullableMissionStatus {
	return &NullableMissionStatus{value: val, isSet: true}
}

func (v NullableMissionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

