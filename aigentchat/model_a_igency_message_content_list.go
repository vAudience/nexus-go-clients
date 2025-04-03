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

// checks if the AIgencyMessageContentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIgencyMessageContentList{}

// AIgencyMessageContentList struct for AIgencyMessageContentList
type AIgencyMessageContentList struct {
	Data []AIgencyMessageContent `json:"data"`
	FullText string `json:"full_text"`
	AdditionalProperties map[string]interface{}
}

type _AIgencyMessageContentList AIgencyMessageContentList

// NewAIgencyMessageContentList instantiates a new AIgencyMessageContentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIgencyMessageContentList(data []AIgencyMessageContent, fullText string) *AIgencyMessageContentList {
	this := AIgencyMessageContentList{}
	this.Data = data
	this.FullText = fullText
	return &this
}

// NewAIgencyMessageContentListWithDefaults instantiates a new AIgencyMessageContentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIgencyMessageContentListWithDefaults() *AIgencyMessageContentList {
	this := AIgencyMessageContentList{}
	return &this
}

// GetData returns the Data field value
func (o *AIgencyMessageContentList) GetData() []AIgencyMessageContent {
	if o == nil {
		var ret []AIgencyMessageContent
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageContentList) GetDataOk() ([]AIgencyMessageContent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AIgencyMessageContentList) SetData(v []AIgencyMessageContent) {
	o.Data = v
}

// GetFullText returns the FullText field value
func (o *AIgencyMessageContentList) GetFullText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullText
}

// GetFullTextOk returns a tuple with the FullText field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageContentList) GetFullTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullText, true
}

// SetFullText sets field value
func (o *AIgencyMessageContentList) SetFullText(v string) {
	o.FullText = v
}

func (o AIgencyMessageContentList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIgencyMessageContentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["full_text"] = o.FullText

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIgencyMessageContentList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"full_text",
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

	varAIgencyMessageContentList := _AIgencyMessageContentList{}

	err = json.Unmarshal(data, &varAIgencyMessageContentList)

	if err != nil {
		return err
	}

	*o = AIgencyMessageContentList(varAIgencyMessageContentList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "full_text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIgencyMessageContentList struct {
	value *AIgencyMessageContentList
	isSet bool
}

func (v NullableAIgencyMessageContentList) Get() *AIgencyMessageContentList {
	return v.value
}

func (v *NullableAIgencyMessageContentList) Set(val *AIgencyMessageContentList) {
	v.value = val
	v.isSet = true
}

func (v NullableAIgencyMessageContentList) IsSet() bool {
	return v.isSet
}

func (v *NullableAIgencyMessageContentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIgencyMessageContentList(val *AIgencyMessageContentList) *NullableAIgencyMessageContentList {
	return &NullableAIgencyMessageContentList{value: val, isSet: true}
}

func (v NullableAIgencyMessageContentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIgencyMessageContentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


