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

// AiServiceId the model 'AiServiceId'
type AiServiceId string

// List of AiServiceId
const (
	AiServiceIdAnthropic AiServiceId = "anthropic"
	AiServiceIdAnthropicGcp AiServiceId = "anthropicgooglecloud"
	AiServiceIdAssemblyAI AiServiceId = "assemblyai"
	AiServiceIdAzureAI AiServiceId = "azureai"
	AiServiceIdBlackForest AiServiceId = "blackforest"
	AiServiceIdElevenLabs AiServiceId = "elevenlabs"
	AiServiceIdGoogleVertexAi AiServiceId = "googlevertexai"
	AiServiceIdLeonardoAI AiServiceId = "leonardoai"
	AiServiceIdMistralAI AiServiceId = "mistralai"
	AiServiceIdOpenAI AiServiceId = "openai"
	AiServiceIdReplicate AiServiceId = "replicate"
	AiServiceIdScaleway AiServiceId = "scaleway"
	AiServiceIdVaiEmbed AiServiceId = "vaiembed"
)

// All allowed values of AiServiceId enum
var AllowedAiServiceIdEnumValues = []AiServiceId{
	"anthropic",
	"anthropicgooglecloud",
	"assemblyai",
	"azureai",
	"blackforest",
	"elevenlabs",
	"googlevertexai",
	"leonardoai",
	"mistralai",
	"openai",
	"replicate",
	"scaleway",
	"vaiembed",
}

func (v *AiServiceId) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiServiceId(value)
	for _, existing := range AllowedAiServiceIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiServiceId", value)
}

// NewAiServiceIdFromValue returns a pointer to a valid AiServiceId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiServiceIdFromValue(v string) (*AiServiceId, error) {
	ev := AiServiceId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiServiceId: valid values are %v", v, AllowedAiServiceIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiServiceId) IsValid() bool {
	for _, existing := range AllowedAiServiceIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiServiceId value
func (v AiServiceId) Ptr() *AiServiceId {
	return &v
}

type NullableAiServiceId struct {
	value *AiServiceId
	isSet bool
}

func (v NullableAiServiceId) Get() *AiServiceId {
	return v.value
}

func (v *NullableAiServiceId) Set(val *AiServiceId) {
	v.value = val
	v.isSet = true
}

func (v NullableAiServiceId) IsSet() bool {
	return v.isSet
}

func (v *NullableAiServiceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiServiceId(val *AiServiceId) *NullableAiServiceId {
	return &NullableAiServiceId{value: val, isSet: true}
}

func (v NullableAiServiceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiServiceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

