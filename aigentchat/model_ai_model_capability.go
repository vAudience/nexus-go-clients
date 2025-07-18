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

// AIModelCapability the model 'AIModelCapability'
type AIModelCapability string

// List of AIModelCapability
const (
	AIModelCapabilityNil AIModelCapability = ""
	AIModelCapabilityFunctionCalling AIModelCapability = "function-calling"
	AIModelCapabilityFunctionCallingStreaming AIModelCapability = "function-calling_streaming"
	AIModelCapabilityTextToEmbedding AIModelCapability = "text-to-embedding"
	AIModelCapabilityTextToText AIModelCapability = "text-to-text"
	AIModelCapabilityTextToTextStreaming AIModelCapability = "text-to-text_streaming"
	AIModelCapabilityTextToImage AIModelCapability = "text-to-image"
	AIModelCapabilityTextToSpeech AIModelCapability = "text-to-speech"
	AIModelCapabilityTextToSpeechStreaming AIModelCapability = "text-to-speech_streaming"
	AIModelCapabilityTextToMusic AIModelCapability = "text-to-music"
	AIModelCapabilityTextToMusicStreaming AIModelCapability = "text-to-music_streaming"
	AIModelCapabilityTextToVideo AIModelCapability = "text-to-video"
	AIModelCapabilityTextToVideoStreaming AIModelCapability = "text-to-video_streaming"
	AIModelCapabilitySpeechToEmbedding AIModelCapability = "speech-to-embedding"
	AIModelCapabilitySpeechToText AIModelCapability = "speech-to-text"
	AIModelCapabilitySpeechToTextStreaming AIModelCapability = "speech-to-text_streaming"
	AIModelCapabilityImageToEmbedding AIModelCapability = "image-to-embedding"
	AIModelCapabilityImageToText AIModelCapability = "image-to-text"
	AIModelCapabilityVideoToEmbedding AIModelCapability = "video-to-embedding"
	AIModelCapabilityVideoToText AIModelCapability = "video-to-text"
	AIModelCapabilityVideoToTextStreaming AIModelCapability = "video-to-text_streaming"
)

// All allowed values of AIModelCapability enum
var AllowedAIModelCapabilityEnumValues = []AIModelCapability{
	"",
	"function-calling",
	"function-calling_streaming",
	"text-to-embedding",
	"text-to-text",
	"text-to-text_streaming",
	"text-to-image",
	"text-to-speech",
	"text-to-speech_streaming",
	"text-to-music",
	"text-to-music_streaming",
	"text-to-video",
	"text-to-video_streaming",
	"speech-to-embedding",
	"speech-to-text",
	"speech-to-text_streaming",
	"image-to-embedding",
	"image-to-text",
	"video-to-embedding",
	"video-to-text",
	"video-to-text_streaming",
}

func (v *AIModelCapability) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AIModelCapability(value)
	for _, existing := range AllowedAIModelCapabilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AIModelCapability", value)
}

// NewAIModelCapabilityFromValue returns a pointer to a valid AIModelCapability
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAIModelCapabilityFromValue(v string) (*AIModelCapability, error) {
	ev := AIModelCapability(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AIModelCapability: valid values are %v", v, AllowedAIModelCapabilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AIModelCapability) IsValid() bool {
	for _, existing := range AllowedAIModelCapabilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIModelCapability value
func (v AIModelCapability) Ptr() *AIModelCapability {
	return &v
}

type NullableAIModelCapability struct {
	value *AIModelCapability
	isSet bool
}

func (v NullableAIModelCapability) Get() *AIModelCapability {
	return v.value
}

func (v *NullableAIModelCapability) Set(val *AIModelCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelCapability(val *AIModelCapability) *NullableAIModelCapability {
	return &NullableAIModelCapability{value: val, isSet: true}
}

func (v NullableAIModelCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

