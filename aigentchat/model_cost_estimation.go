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

// checks if the CostEstimation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CostEstimation{}

// CostEstimation struct for CostEstimation
type CostEstimation struct {
	CostInEuro float64 `json:"cost_in_euro"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CostEstimation CostEstimation

// NewCostEstimation instantiates a new CostEstimation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostEstimation(costInEuro float64) *CostEstimation {
	this := CostEstimation{}
	this.CostInEuro = costInEuro
	return &this
}

// NewCostEstimationWithDefaults instantiates a new CostEstimation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostEstimationWithDefaults() *CostEstimation {
	this := CostEstimation{}
	return &this
}

// GetCostInEuro returns the CostInEuro field value
func (o *CostEstimation) GetCostInEuro() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CostInEuro
}

// GetCostInEuroOk returns a tuple with the CostInEuro field value
// and a boolean to check if the value has been set.
func (o *CostEstimation) GetCostInEuroOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostInEuro, true
}

// SetCostInEuro sets field value
func (o *CostEstimation) SetCostInEuro(v float64) {
	o.CostInEuro = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CostEstimation) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimation) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CostEstimation) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *CostEstimation) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

func (o CostEstimation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CostEstimation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cost_in_euro"] = o.CostInEuro
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CostEstimation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cost_in_euro",
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

	varCostEstimation := _CostEstimation{}

	err = json.Unmarshal(data, &varCostEstimation)

	if err != nil {
		return err
	}

	*o = CostEstimation(varCostEstimation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cost_in_euro")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCostEstimation struct {
	value *CostEstimation
	isSet bool
}

func (v NullableCostEstimation) Get() *CostEstimation {
	return v.value
}

func (v *NullableCostEstimation) Set(val *CostEstimation) {
	v.value = val
	v.isSet = true
}

func (v NullableCostEstimation) IsSet() bool {
	return v.isSet
}

func (v *NullableCostEstimation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostEstimation(val *CostEstimation) *NullableCostEstimation {
	return &NullableCostEstimation{value: val, isSet: true}
}

func (v NullableCostEstimation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostEstimation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


