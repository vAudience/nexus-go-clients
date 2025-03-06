/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the HealthResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthResponse{}

// HealthResponse struct for HealthResponse
type HealthResponse struct {
	Health string `json:"health"`
	AdditionalProperties map[string]interface{}
}

type _HealthResponse HealthResponse

// NewHealthResponse instantiates a new HealthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthResponse(health string) *HealthResponse {
	this := HealthResponse{}
	this.Health = health
	return &this
}

// NewHealthResponseWithDefaults instantiates a new HealthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthResponseWithDefaults() *HealthResponse {
	this := HealthResponse{}
	return &this
}

// GetHealth returns the Health field value
func (o *HealthResponse) GetHealth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *HealthResponse) GetHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *HealthResponse) SetHealth(v string) {
	o.Health = v
}

func (o HealthResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["health"] = o.Health

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HealthResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"health",
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

	varHealthResponse := _HealthResponse{}

	err = json.Unmarshal(data, &varHealthResponse)

	if err != nil {
		return err
	}

	*o = HealthResponse(varHealthResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "health")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHealthResponse struct {
	value *HealthResponse
	isSet bool
}

func (v NullableHealthResponse) Get() *HealthResponse {
	return v.value
}

func (v *NullableHealthResponse) Set(val *HealthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthResponse(val *HealthResponse) *NullableHealthResponse {
	return &NullableHealthResponse{value: val, isSet: true}
}

func (v NullableHealthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


