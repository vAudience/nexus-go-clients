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

// checks if the MissionWriteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionWriteDto{}

// MissionWriteDto struct for MissionWriteDto
type MissionWriteDto struct {
	Description *string `json:"description,omitempty"`
	Instructions MissionInstructions `json:"instructions"`
	AdditionalProperties map[string]interface{}
}

type _MissionWriteDto MissionWriteDto

// NewMissionWriteDto instantiates a new MissionWriteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionWriteDto(instructions MissionInstructions) *MissionWriteDto {
	this := MissionWriteDto{}
	this.Instructions = instructions
	return &this
}

// NewMissionWriteDtoWithDefaults instantiates a new MissionWriteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionWriteDtoWithDefaults() *MissionWriteDto {
	this := MissionWriteDto{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MissionWriteDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionWriteDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MissionWriteDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MissionWriteDto) SetDescription(v string) {
	o.Description = &v
}

// GetInstructions returns the Instructions field value
func (o *MissionWriteDto) GetInstructions() MissionInstructions {
	if o == nil {
		var ret MissionInstructions
		return ret
	}

	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value
// and a boolean to check if the value has been set.
func (o *MissionWriteDto) GetInstructionsOk() (*MissionInstructions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instructions, true
}

// SetInstructions sets field value
func (o *MissionWriteDto) SetInstructions(v MissionInstructions) {
	o.Instructions = v
}

func (o MissionWriteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionWriteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["instructions"] = o.Instructions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionWriteDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instructions",
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

	varMissionWriteDto := _MissionWriteDto{}

	err = json.Unmarshal(data, &varMissionWriteDto)

	if err != nil {
		return err
	}

	*o = MissionWriteDto(varMissionWriteDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "instructions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionWriteDto struct {
	value *MissionWriteDto
	isSet bool
}

func (v NullableMissionWriteDto) Get() *MissionWriteDto {
	return v.value
}

func (v *NullableMissionWriteDto) Set(val *MissionWriteDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionWriteDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionWriteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionWriteDto(val *MissionWriteDto) *NullableMissionWriteDto {
	return &NullableMissionWriteDto{value: val, isSet: true}
}

func (v NullableMissionWriteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionWriteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


