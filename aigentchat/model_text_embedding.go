/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.17.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the TextEmbedding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextEmbedding{}

// TextEmbedding struct for TextEmbedding
type TextEmbedding struct {
	Embeddings []EmbeddingItem `json:"embeddings"`
	ExecutionId string `json:"execution_id"`
	ModelId string `json:"model_id"`
	Parameters map[string]interface{} `json:"parameters"`
	ServiceId string `json:"service_id"`
	Tokens int32 `json:"tokens"`
	AdditionalProperties map[string]interface{}
}

type _TextEmbedding TextEmbedding

// NewTextEmbedding instantiates a new TextEmbedding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextEmbedding(embeddings []EmbeddingItem, executionId string, modelId string, parameters map[string]interface{}, serviceId string, tokens int32) *TextEmbedding {
	this := TextEmbedding{}
	this.Embeddings = embeddings
	this.ExecutionId = executionId
	this.ModelId = modelId
	this.Parameters = parameters
	this.ServiceId = serviceId
	this.Tokens = tokens
	return &this
}

// NewTextEmbeddingWithDefaults instantiates a new TextEmbedding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextEmbeddingWithDefaults() *TextEmbedding {
	this := TextEmbedding{}
	return &this
}

// GetEmbeddings returns the Embeddings field value
func (o *TextEmbedding) GetEmbeddings() []EmbeddingItem {
	if o == nil {
		var ret []EmbeddingItem
		return ret
	}

	return o.Embeddings
}

// GetEmbeddingsOk returns a tuple with the Embeddings field value
// and a boolean to check if the value has been set.
func (o *TextEmbedding) GetEmbeddingsOk() ([]EmbeddingItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeddings, true
}

// SetEmbeddings sets field value
func (o *TextEmbedding) SetEmbeddings(v []EmbeddingItem) {
	o.Embeddings = v
}

// GetExecutionId returns the ExecutionId field value
func (o *TextEmbedding) GetExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value
// and a boolean to check if the value has been set.
func (o *TextEmbedding) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionId, true
}

// SetExecutionId sets field value
func (o *TextEmbedding) SetExecutionId(v string) {
	o.ExecutionId = v
}

// GetModelId returns the ModelId field value
func (o *TextEmbedding) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *TextEmbedding) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *TextEmbedding) SetModelId(v string) {
	o.ModelId = v
}

// GetParameters returns the Parameters field value
func (o *TextEmbedding) GetParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *TextEmbedding) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *TextEmbedding) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetServiceId returns the ServiceId field value
func (o *TextEmbedding) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *TextEmbedding) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *TextEmbedding) SetServiceId(v string) {
	o.ServiceId = v
}

// GetTokens returns the Tokens field value
func (o *TextEmbedding) GetTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *TextEmbedding) GetTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokens, true
}

// SetTokens sets field value
func (o *TextEmbedding) SetTokens(v int32) {
	o.Tokens = v
}

func (o TextEmbedding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextEmbedding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["embeddings"] = o.Embeddings
	toSerialize["execution_id"] = o.ExecutionId
	toSerialize["model_id"] = o.ModelId
	toSerialize["parameters"] = o.Parameters
	toSerialize["service_id"] = o.ServiceId
	toSerialize["tokens"] = o.Tokens

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TextEmbedding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"embeddings",
		"execution_id",
		"model_id",
		"parameters",
		"service_id",
		"tokens",
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

	varTextEmbedding := _TextEmbedding{}

	err = json.Unmarshal(data, &varTextEmbedding)

	if err != nil {
		return err
	}

	*o = TextEmbedding(varTextEmbedding)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "embeddings")
		delete(additionalProperties, "execution_id")
		delete(additionalProperties, "model_id")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "tokens")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTextEmbedding struct {
	value *TextEmbedding
	isSet bool
}

func (v NullableTextEmbedding) Get() *TextEmbedding {
	return v.value
}

func (v *NullableTextEmbedding) Set(val *TextEmbedding) {
	v.value = val
	v.isSet = true
}

func (v NullableTextEmbedding) IsSet() bool {
	return v.isSet
}

func (v *NullableTextEmbedding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextEmbedding(val *TextEmbedding) *NullableTextEmbedding {
	return &NullableTextEmbedding{value: val, isSet: true}
}

func (v NullableTextEmbedding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextEmbedding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


