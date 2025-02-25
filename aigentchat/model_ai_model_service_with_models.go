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
)

// checks if the AIModelServiceWithModels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIModelServiceWithModels{}

// AIModelServiceWithModels struct for AIModelServiceWithModels
type AIModelServiceWithModels struct {
	Models []AIModel `json:"models,omitempty"`
	Service *AIModelServiceObject `json:"service,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIModelServiceWithModels AIModelServiceWithModels

// NewAIModelServiceWithModels instantiates a new AIModelServiceWithModels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIModelServiceWithModels() *AIModelServiceWithModels {
	this := AIModelServiceWithModels{}
	return &this
}

// NewAIModelServiceWithModelsWithDefaults instantiates a new AIModelServiceWithModels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIModelServiceWithModelsWithDefaults() *AIModelServiceWithModels {
	this := AIModelServiceWithModels{}
	return &this
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *AIModelServiceWithModels) GetModels() []AIModel {
	if o == nil || IsNil(o.Models) {
		var ret []AIModel
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceWithModels) GetModelsOk() ([]AIModel, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *AIModelServiceWithModels) HasModels() bool {
	if o != nil && !IsNil(o.Models) {
		return true
	}

	return false
}

// SetModels gets a reference to the given []AIModel and assigns it to the Models field.
func (o *AIModelServiceWithModels) SetModels(v []AIModel) {
	o.Models = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *AIModelServiceWithModels) GetService() AIModelServiceObject {
	if o == nil || IsNil(o.Service) {
		var ret AIModelServiceObject
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceWithModels) GetServiceOk() (*AIModelServiceObject, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *AIModelServiceWithModels) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given AIModelServiceObject and assigns it to the Service field.
func (o *AIModelServiceWithModels) SetService(v AIModelServiceObject) {
	o.Service = &v
}

func (o AIModelServiceWithModels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIModelServiceWithModels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Models) {
		toSerialize["models"] = o.Models
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIModelServiceWithModels) UnmarshalJSON(data []byte) (err error) {
	varAIModelServiceWithModels := _AIModelServiceWithModels{}

	err = json.Unmarshal(data, &varAIModelServiceWithModels)

	if err != nil {
		return err
	}

	*o = AIModelServiceWithModels(varAIModelServiceWithModels)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "models")
		delete(additionalProperties, "service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIModelServiceWithModels struct {
	value *AIModelServiceWithModels
	isSet bool
}

func (v NullableAIModelServiceWithModels) Get() *AIModelServiceWithModels {
	return v.value
}

func (v *NullableAIModelServiceWithModels) Set(val *AIModelServiceWithModels) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelServiceWithModels) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelServiceWithModels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelServiceWithModels(val *AIModelServiceWithModels) *NullableAIModelServiceWithModels {
	return &NullableAIModelServiceWithModels{value: val, isSet: true}
}

func (v NullableAIModelServiceWithModels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelServiceWithModels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


