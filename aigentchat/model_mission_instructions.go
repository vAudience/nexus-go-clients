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

// checks if the MissionInstructions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionInstructions{}

// MissionInstructions struct for MissionInstructions
type MissionInstructions struct {
	MaxTokens *int32 `json:"max_tokens,omitempty"`
	Temperature *float32 `json:"temperature,omitempty"`
	Text string `json:"text"`
	VarReplacements *map[string]string `json:"var_replacements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionInstructions MissionInstructions

// NewMissionInstructions instantiates a new MissionInstructions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionInstructions(text string) *MissionInstructions {
	this := MissionInstructions{}
	this.Text = text
	return &this
}

// NewMissionInstructionsWithDefaults instantiates a new MissionInstructions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionInstructionsWithDefaults() *MissionInstructions {
	this := MissionInstructions{}
	return &this
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise.
func (o *MissionInstructions) GetMaxTokens() int32 {
	if o == nil || IsNil(o.MaxTokens) {
		var ret int32
		return ret
	}
	return *o.MaxTokens
}

// GetMaxTokensOk returns a tuple with the MaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionInstructions) GetMaxTokensOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTokens) {
		return nil, false
	}
	return o.MaxTokens, true
}

// HasMaxTokens returns a boolean if a field has been set.
func (o *MissionInstructions) HasMaxTokens() bool {
	if o != nil && !IsNil(o.MaxTokens) {
		return true
	}

	return false
}

// SetMaxTokens gets a reference to the given int32 and assigns it to the MaxTokens field.
func (o *MissionInstructions) SetMaxTokens(v int32) {
	o.MaxTokens = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *MissionInstructions) GetTemperature() float32 {
	if o == nil || IsNil(o.Temperature) {
		var ret float32
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionInstructions) GetTemperatureOk() (*float32, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *MissionInstructions) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given float32 and assigns it to the Temperature field.
func (o *MissionInstructions) SetTemperature(v float32) {
	o.Temperature = &v
}

// GetText returns the Text field value
func (o *MissionInstructions) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MissionInstructions) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *MissionInstructions) SetText(v string) {
	o.Text = v
}

// GetVarReplacements returns the VarReplacements field value if set, zero value otherwise.
func (o *MissionInstructions) GetVarReplacements() map[string]string {
	if o == nil || IsNil(o.VarReplacements) {
		var ret map[string]string
		return ret
	}
	return *o.VarReplacements
}

// GetVarReplacementsOk returns a tuple with the VarReplacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionInstructions) GetVarReplacementsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.VarReplacements) {
		return nil, false
	}
	return o.VarReplacements, true
}

// HasVarReplacements returns a boolean if a field has been set.
func (o *MissionInstructions) HasVarReplacements() bool {
	if o != nil && !IsNil(o.VarReplacements) {
		return true
	}

	return false
}

// SetVarReplacements gets a reference to the given map[string]string and assigns it to the VarReplacements field.
func (o *MissionInstructions) SetVarReplacements(v map[string]string) {
	o.VarReplacements = &v
}

func (o MissionInstructions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionInstructions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxTokens) {
		toSerialize["max_tokens"] = o.MaxTokens
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	toSerialize["text"] = o.Text
	if !IsNil(o.VarReplacements) {
		toSerialize["var_replacements"] = o.VarReplacements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionInstructions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
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

	varMissionInstructions := _MissionInstructions{}

	err = json.Unmarshal(data, &varMissionInstructions)

	if err != nil {
		return err
	}

	*o = MissionInstructions(varMissionInstructions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "max_tokens")
		delete(additionalProperties, "temperature")
		delete(additionalProperties, "text")
		delete(additionalProperties, "var_replacements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionInstructions struct {
	value *MissionInstructions
	isSet bool
}

func (v NullableMissionInstructions) Get() *MissionInstructions {
	return v.value
}

func (v *NullableMissionInstructions) Set(val *MissionInstructions) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionInstructions) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionInstructions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionInstructions(val *MissionInstructions) *NullableMissionInstructions {
	return &NullableMissionInstructions{value: val, isSet: true}
}

func (v NullableMissionInstructions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionInstructions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


