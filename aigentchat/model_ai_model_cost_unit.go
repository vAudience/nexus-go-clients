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

// AIModelCostUnit the model 'AIModelCostUnit'
type AIModelCostUnit string

// List of AIModelCostUnit
const (
	AIModelCostUnitInputPerMillionTokens AIModelCostUnit = "input-tokens-per-million"
	AIModelCostUnitOutputPerMillionTokens AIModelCostUnit = "output-tokens-per-million"
	AIModelCostUnitInputPerMillionCharacters AIModelCostUnit = "input-characters-per-million"
	AIModelCostUnitImageInputPerFile AIModelCostUnit = "image-input-file"
	AIModelCostUnitAudioInputPerSecond AIModelCostUnit = "audio-input-per-second"
	AIModelCostUnitVideoInputPerSecond AIModelCostUnit = "video-input-per-second"
	AIModelCostUnitImageGenerationPerImage AIModelCostUnit = "image-generation-per-image"
	AIModelCostUnitImageGenerationPerPixel AIModelCostUnit = "image-generation-per-pixel"
	AIModelCostUnitAudioGenerationPerSecond AIModelCostUnit = "audio-generation-per-second"
	AIModelCostUnitVideoGenerationPerSecond AIModelCostUnit = "video-generation-per-second"
	AIModelCostUnitPerFunctionCall AIModelCostUnit = "per-function-call"
)

// All allowed values of AIModelCostUnit enum
var AllowedAIModelCostUnitEnumValues = []AIModelCostUnit{
	"input-tokens-per-million",
	"output-tokens-per-million",
	"input-characters-per-million",
	"image-input-file",
	"audio-input-per-second",
	"video-input-per-second",
	"image-generation-per-image",
	"image-generation-per-pixel",
	"audio-generation-per-second",
	"video-generation-per-second",
	"per-function-call",
}

func (v *AIModelCostUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AIModelCostUnit(value)
	for _, existing := range AllowedAIModelCostUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AIModelCostUnit", value)
}

// NewAIModelCostUnitFromValue returns a pointer to a valid AIModelCostUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAIModelCostUnitFromValue(v string) (*AIModelCostUnit, error) {
	ev := AIModelCostUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AIModelCostUnit: valid values are %v", v, AllowedAIModelCostUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AIModelCostUnit) IsValid() bool {
	for _, existing := range AllowedAIModelCostUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIModelCostUnit value
func (v AIModelCostUnit) Ptr() *AIModelCostUnit {
	return &v
}

type NullableAIModelCostUnit struct {
	value *AIModelCostUnit
	isSet bool
}

func (v NullableAIModelCostUnit) Get() *AIModelCostUnit {
	return v.value
}

func (v *NullableAIModelCostUnit) Set(val *AIModelCostUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelCostUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelCostUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelCostUnit(val *AIModelCostUnit) *NullableAIModelCostUnit {
	return &NullableAIModelCostUnit{value: val, isSet: true}
}

func (v NullableAIModelCostUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelCostUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

