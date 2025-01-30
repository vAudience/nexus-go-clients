/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.14.1
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the ExecutionUsageCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionUsageCost{}

// ExecutionUsageCost struct for ExecutionUsageCost
type ExecutionUsageCost struct {
	CostPerUnitInEuro *float64 `json:"cost_per_unit_in_euro,omitempty"`
	CostUnit *AIModelCostUnit `json:"cost_unit,omitempty"`
	Description *string `json:"description,omitempty"`
	ResultingCostInEuro *float64 `json:"resulting_cost_in_euro,omitempty"`
	UsedUnits *float64 `json:"used_units,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionUsageCost ExecutionUsageCost

// NewExecutionUsageCost instantiates a new ExecutionUsageCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionUsageCost() *ExecutionUsageCost {
	this := ExecutionUsageCost{}
	return &this
}

// NewExecutionUsageCostWithDefaults instantiates a new ExecutionUsageCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionUsageCostWithDefaults() *ExecutionUsageCost {
	this := ExecutionUsageCost{}
	return &this
}

// GetCostPerUnitInEuro returns the CostPerUnitInEuro field value if set, zero value otherwise.
func (o *ExecutionUsageCost) GetCostPerUnitInEuro() float64 {
	if o == nil || IsNil(o.CostPerUnitInEuro) {
		var ret float64
		return ret
	}
	return *o.CostPerUnitInEuro
}

// GetCostPerUnitInEuroOk returns a tuple with the CostPerUnitInEuro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionUsageCost) GetCostPerUnitInEuroOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerUnitInEuro) {
		return nil, false
	}
	return o.CostPerUnitInEuro, true
}

// HasCostPerUnitInEuro returns a boolean if a field has been set.
func (o *ExecutionUsageCost) HasCostPerUnitInEuro() bool {
	if o != nil && !IsNil(o.CostPerUnitInEuro) {
		return true
	}

	return false
}

// SetCostPerUnitInEuro gets a reference to the given float64 and assigns it to the CostPerUnitInEuro field.
func (o *ExecutionUsageCost) SetCostPerUnitInEuro(v float64) {
	o.CostPerUnitInEuro = &v
}

// GetCostUnit returns the CostUnit field value if set, zero value otherwise.
func (o *ExecutionUsageCost) GetCostUnit() AIModelCostUnit {
	if o == nil || IsNil(o.CostUnit) {
		var ret AIModelCostUnit
		return ret
	}
	return *o.CostUnit
}

// GetCostUnitOk returns a tuple with the CostUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionUsageCost) GetCostUnitOk() (*AIModelCostUnit, bool) {
	if o == nil || IsNil(o.CostUnit) {
		return nil, false
	}
	return o.CostUnit, true
}

// HasCostUnit returns a boolean if a field has been set.
func (o *ExecutionUsageCost) HasCostUnit() bool {
	if o != nil && !IsNil(o.CostUnit) {
		return true
	}

	return false
}

// SetCostUnit gets a reference to the given AIModelCostUnit and assigns it to the CostUnit field.
func (o *ExecutionUsageCost) SetCostUnit(v AIModelCostUnit) {
	o.CostUnit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExecutionUsageCost) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionUsageCost) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExecutionUsageCost) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExecutionUsageCost) SetDescription(v string) {
	o.Description = &v
}

// GetResultingCostInEuro returns the ResultingCostInEuro field value if set, zero value otherwise.
func (o *ExecutionUsageCost) GetResultingCostInEuro() float64 {
	if o == nil || IsNil(o.ResultingCostInEuro) {
		var ret float64
		return ret
	}
	return *o.ResultingCostInEuro
}

// GetResultingCostInEuroOk returns a tuple with the ResultingCostInEuro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionUsageCost) GetResultingCostInEuroOk() (*float64, bool) {
	if o == nil || IsNil(o.ResultingCostInEuro) {
		return nil, false
	}
	return o.ResultingCostInEuro, true
}

// HasResultingCostInEuro returns a boolean if a field has been set.
func (o *ExecutionUsageCost) HasResultingCostInEuro() bool {
	if o != nil && !IsNil(o.ResultingCostInEuro) {
		return true
	}

	return false
}

// SetResultingCostInEuro gets a reference to the given float64 and assigns it to the ResultingCostInEuro field.
func (o *ExecutionUsageCost) SetResultingCostInEuro(v float64) {
	o.ResultingCostInEuro = &v
}

// GetUsedUnits returns the UsedUnits field value if set, zero value otherwise.
func (o *ExecutionUsageCost) GetUsedUnits() float64 {
	if o == nil || IsNil(o.UsedUnits) {
		var ret float64
		return ret
	}
	return *o.UsedUnits
}

// GetUsedUnitsOk returns a tuple with the UsedUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionUsageCost) GetUsedUnitsOk() (*float64, bool) {
	if o == nil || IsNil(o.UsedUnits) {
		return nil, false
	}
	return o.UsedUnits, true
}

// HasUsedUnits returns a boolean if a field has been set.
func (o *ExecutionUsageCost) HasUsedUnits() bool {
	if o != nil && !IsNil(o.UsedUnits) {
		return true
	}

	return false
}

// SetUsedUnits gets a reference to the given float64 and assigns it to the UsedUnits field.
func (o *ExecutionUsageCost) SetUsedUnits(v float64) {
	o.UsedUnits = &v
}

func (o ExecutionUsageCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionUsageCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CostPerUnitInEuro) {
		toSerialize["cost_per_unit_in_euro"] = o.CostPerUnitInEuro
	}
	if !IsNil(o.CostUnit) {
		toSerialize["cost_unit"] = o.CostUnit
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ResultingCostInEuro) {
		toSerialize["resulting_cost_in_euro"] = o.ResultingCostInEuro
	}
	if !IsNil(o.UsedUnits) {
		toSerialize["used_units"] = o.UsedUnits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionUsageCost) UnmarshalJSON(data []byte) (err error) {
	varExecutionUsageCost := _ExecutionUsageCost{}

	err = json.Unmarshal(data, &varExecutionUsageCost)

	if err != nil {
		return err
	}

	*o = ExecutionUsageCost(varExecutionUsageCost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cost_per_unit_in_euro")
		delete(additionalProperties, "cost_unit")
		delete(additionalProperties, "description")
		delete(additionalProperties, "resulting_cost_in_euro")
		delete(additionalProperties, "used_units")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionUsageCost struct {
	value *ExecutionUsageCost
	isSet bool
}

func (v NullableExecutionUsageCost) Get() *ExecutionUsageCost {
	return v.value
}

func (v *NullableExecutionUsageCost) Set(val *ExecutionUsageCost) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionUsageCost) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionUsageCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionUsageCost(val *ExecutionUsageCost) *NullableExecutionUsageCost {
	return &NullableExecutionUsageCost{value: val, isSet: true}
}

func (v NullableExecutionUsageCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionUsageCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


