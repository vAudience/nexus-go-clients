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

// checks if the TextEmbeddingRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextEmbeddingRequestDto{}

// TextEmbeddingRequestDto struct for TextEmbeddingRequestDto
type TextEmbeddingRequestDto struct {
	AgentId string `json:"agent_id"`
	Items []TextEmbeddingItemDto `json:"items"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TextEmbeddingRequestDto TextEmbeddingRequestDto

// NewTextEmbeddingRequestDto instantiates a new TextEmbeddingRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextEmbeddingRequestDto(agentId string, items []TextEmbeddingItemDto) *TextEmbeddingRequestDto {
	this := TextEmbeddingRequestDto{}
	this.AgentId = agentId
	this.Items = items
	return &this
}

// NewTextEmbeddingRequestDtoWithDefaults instantiates a new TextEmbeddingRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextEmbeddingRequestDtoWithDefaults() *TextEmbeddingRequestDto {
	this := TextEmbeddingRequestDto{}
	return &this
}

// GetAgentId returns the AgentId field value
func (o *TextEmbeddingRequestDto) GetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *TextEmbeddingRequestDto) GetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *TextEmbeddingRequestDto) SetAgentId(v string) {
	o.AgentId = v
}

// GetItems returns the Items field value
func (o *TextEmbeddingRequestDto) GetItems() []TextEmbeddingItemDto {
	if o == nil {
		var ret []TextEmbeddingItemDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TextEmbeddingRequestDto) GetItemsOk() ([]TextEmbeddingItemDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *TextEmbeddingRequestDto) SetItems(v []TextEmbeddingItemDto) {
	o.Items = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *TextEmbeddingRequestDto) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextEmbeddingRequestDto) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *TextEmbeddingRequestDto) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *TextEmbeddingRequestDto) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

func (o TextEmbeddingRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextEmbeddingRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent_id"] = o.AgentId
	toSerialize["items"] = o.Items
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TextEmbeddingRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent_id",
		"items",
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

	varTextEmbeddingRequestDto := _TextEmbeddingRequestDto{}

	err = json.Unmarshal(data, &varTextEmbeddingRequestDto)

	if err != nil {
		return err
	}

	*o = TextEmbeddingRequestDto(varTextEmbeddingRequestDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agent_id")
		delete(additionalProperties, "items")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTextEmbeddingRequestDto struct {
	value *TextEmbeddingRequestDto
	isSet bool
}

func (v NullableTextEmbeddingRequestDto) Get() *TextEmbeddingRequestDto {
	return v.value
}

func (v *NullableTextEmbeddingRequestDto) Set(val *TextEmbeddingRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTextEmbeddingRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTextEmbeddingRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextEmbeddingRequestDto(val *TextEmbeddingRequestDto) *NullableTextEmbeddingRequestDto {
	return &NullableTextEmbeddingRequestDto{value: val, isSet: true}
}

func (v NullableTextEmbeddingRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextEmbeddingRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


