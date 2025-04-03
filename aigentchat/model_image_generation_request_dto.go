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

// checks if the ImageGenerationRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageGenerationRequestDto{}

// ImageGenerationRequestDto struct for ImageGenerationRequestDto
type ImageGenerationRequestDto struct {
	AgentId string `json:"agent_id"`
	AttachedTemporaryFiles []string `json:"attached_temporary_files,omitempty"`
	Message string `json:"message"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImageGenerationRequestDto ImageGenerationRequestDto

// NewImageGenerationRequestDto instantiates a new ImageGenerationRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageGenerationRequestDto(agentId string, message string) *ImageGenerationRequestDto {
	this := ImageGenerationRequestDto{}
	this.AgentId = agentId
	this.Message = message
	return &this
}

// NewImageGenerationRequestDtoWithDefaults instantiates a new ImageGenerationRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageGenerationRequestDtoWithDefaults() *ImageGenerationRequestDto {
	this := ImageGenerationRequestDto{}
	return &this
}

// GetAgentId returns the AgentId field value
func (o *ImageGenerationRequestDto) GetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *ImageGenerationRequestDto) GetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *ImageGenerationRequestDto) SetAgentId(v string) {
	o.AgentId = v
}

// GetAttachedTemporaryFiles returns the AttachedTemporaryFiles field value if set, zero value otherwise.
func (o *ImageGenerationRequestDto) GetAttachedTemporaryFiles() []string {
	if o == nil || IsNil(o.AttachedTemporaryFiles) {
		var ret []string
		return ret
	}
	return o.AttachedTemporaryFiles
}

// GetAttachedTemporaryFilesOk returns a tuple with the AttachedTemporaryFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageGenerationRequestDto) GetAttachedTemporaryFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.AttachedTemporaryFiles) {
		return nil, false
	}
	return o.AttachedTemporaryFiles, true
}

// HasAttachedTemporaryFiles returns a boolean if a field has been set.
func (o *ImageGenerationRequestDto) HasAttachedTemporaryFiles() bool {
	if o != nil && !IsNil(o.AttachedTemporaryFiles) {
		return true
	}

	return false
}

// SetAttachedTemporaryFiles gets a reference to the given []string and assigns it to the AttachedTemporaryFiles field.
func (o *ImageGenerationRequestDto) SetAttachedTemporaryFiles(v []string) {
	o.AttachedTemporaryFiles = v
}

// GetMessage returns the Message field value
func (o *ImageGenerationRequestDto) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ImageGenerationRequestDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ImageGenerationRequestDto) SetMessage(v string) {
	o.Message = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ImageGenerationRequestDto) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageGenerationRequestDto) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ImageGenerationRequestDto) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *ImageGenerationRequestDto) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

func (o ImageGenerationRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageGenerationRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent_id"] = o.AgentId
	if !IsNil(o.AttachedTemporaryFiles) {
		toSerialize["attached_temporary_files"] = o.AttachedTemporaryFiles
	}
	toSerialize["message"] = o.Message
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImageGenerationRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent_id",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varImageGenerationRequestDto := _ImageGenerationRequestDto{}

	err = json.Unmarshal(data, &varImageGenerationRequestDto)

	if err != nil {
		return err
	}

	*o = ImageGenerationRequestDto(varImageGenerationRequestDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agent_id")
		delete(additionalProperties, "attached_temporary_files")
		delete(additionalProperties, "message")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImageGenerationRequestDto struct {
	value *ImageGenerationRequestDto
	isSet bool
}

func (v NullableImageGenerationRequestDto) Get() *ImageGenerationRequestDto {
	return v.value
}

func (v *NullableImageGenerationRequestDto) Set(val *ImageGenerationRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableImageGenerationRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableImageGenerationRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageGenerationRequestDto(val *ImageGenerationRequestDto) *NullableImageGenerationRequestDto {
	return &NullableImageGenerationRequestDto{value: val, isSet: true}
}

func (v NullableImageGenerationRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageGenerationRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


