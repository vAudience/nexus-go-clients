/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.13.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the EmbeddingItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddingItem{}

// EmbeddingItem struct for EmbeddingItem
type EmbeddingItem struct {
	Embedding []float32 `json:"embedding"`
	Index int32 `json:"index"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmbeddingItem EmbeddingItem

// NewEmbeddingItem instantiates a new EmbeddingItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddingItem(embedding []float32, index int32) *EmbeddingItem {
	this := EmbeddingItem{}
	this.Embedding = embedding
	this.Index = index
	return &this
}

// NewEmbeddingItemWithDefaults instantiates a new EmbeddingItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddingItemWithDefaults() *EmbeddingItem {
	this := EmbeddingItem{}
	return &this
}

// GetEmbedding returns the Embedding field value
func (o *EmbeddingItem) GetEmbedding() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Embedding
}

// GetEmbeddingOk returns a tuple with the Embedding field value
// and a boolean to check if the value has been set.
func (o *EmbeddingItem) GetEmbeddingOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embedding, true
}

// SetEmbedding sets field value
func (o *EmbeddingItem) SetEmbedding(v []float32) {
	o.Embedding = v
}

// GetIndex returns the Index field value
func (o *EmbeddingItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *EmbeddingItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *EmbeddingItem) SetIndex(v int32) {
	o.Index = v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *EmbeddingItem) GetMetaData() map[string]interface{} {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddingItem) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return map[string]interface{}{}, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *EmbeddingItem) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *EmbeddingItem) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

func (o EmbeddingItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddingItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["embedding"] = o.Embedding
	toSerialize["index"] = o.Index
	if !IsNil(o.MetaData) {
		toSerialize["meta_data"] = o.MetaData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmbeddingItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"embedding",
		"index",
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

	varEmbeddingItem := _EmbeddingItem{}

	err = json.Unmarshal(data, &varEmbeddingItem)

	if err != nil {
		return err
	}

	*o = EmbeddingItem(varEmbeddingItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "embedding")
		delete(additionalProperties, "index")
		delete(additionalProperties, "meta_data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmbeddingItem struct {
	value *EmbeddingItem
	isSet bool
}

func (v NullableEmbeddingItem) Get() *EmbeddingItem {
	return v.value
}

func (v *NullableEmbeddingItem) Set(val *EmbeddingItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddingItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddingItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddingItem(val *EmbeddingItem) *NullableEmbeddingItem {
	return &NullableEmbeddingItem{value: val, isSet: true}
}

func (v NullableEmbeddingItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddingItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


