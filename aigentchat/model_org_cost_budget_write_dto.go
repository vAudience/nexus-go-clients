/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.0
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the OrgCostBudgetWriteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCostBudgetWriteDto{}

// OrgCostBudgetWriteDto struct for OrgCostBudgetWriteDto
type OrgCostBudgetWriteDto struct {
	TotalBudget *float64 `json:"total_budget,omitempty"`
	UsedBudget *float64 `json:"used_budget,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgCostBudgetWriteDto OrgCostBudgetWriteDto

// NewOrgCostBudgetWriteDto instantiates a new OrgCostBudgetWriteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCostBudgetWriteDto() *OrgCostBudgetWriteDto {
	this := OrgCostBudgetWriteDto{}
	return &this
}

// NewOrgCostBudgetWriteDtoWithDefaults instantiates a new OrgCostBudgetWriteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCostBudgetWriteDtoWithDefaults() *OrgCostBudgetWriteDto {
	this := OrgCostBudgetWriteDto{}
	return &this
}

// GetTotalBudget returns the TotalBudget field value if set, zero value otherwise.
func (o *OrgCostBudgetWriteDto) GetTotalBudget() float64 {
	if o == nil || IsNil(o.TotalBudget) {
		var ret float64
		return ret
	}
	return *o.TotalBudget
}

// GetTotalBudgetOk returns a tuple with the TotalBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCostBudgetWriteDto) GetTotalBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalBudget) {
		return nil, false
	}
	return o.TotalBudget, true
}

// HasTotalBudget returns a boolean if a field has been set.
func (o *OrgCostBudgetWriteDto) HasTotalBudget() bool {
	if o != nil && !IsNil(o.TotalBudget) {
		return true
	}

	return false
}

// SetTotalBudget gets a reference to the given float64 and assigns it to the TotalBudget field.
func (o *OrgCostBudgetWriteDto) SetTotalBudget(v float64) {
	o.TotalBudget = &v
}

// GetUsedBudget returns the UsedBudget field value if set, zero value otherwise.
func (o *OrgCostBudgetWriteDto) GetUsedBudget() float64 {
	if o == nil || IsNil(o.UsedBudget) {
		var ret float64
		return ret
	}
	return *o.UsedBudget
}

// GetUsedBudgetOk returns a tuple with the UsedBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCostBudgetWriteDto) GetUsedBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.UsedBudget) {
		return nil, false
	}
	return o.UsedBudget, true
}

// HasUsedBudget returns a boolean if a field has been set.
func (o *OrgCostBudgetWriteDto) HasUsedBudget() bool {
	if o != nil && !IsNil(o.UsedBudget) {
		return true
	}

	return false
}

// SetUsedBudget gets a reference to the given float64 and assigns it to the UsedBudget field.
func (o *OrgCostBudgetWriteDto) SetUsedBudget(v float64) {
	o.UsedBudget = &v
}

func (o OrgCostBudgetWriteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCostBudgetWriteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalBudget) {
		toSerialize["total_budget"] = o.TotalBudget
	}
	if !IsNil(o.UsedBudget) {
		toSerialize["used_budget"] = o.UsedBudget
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCostBudgetWriteDto) UnmarshalJSON(data []byte) (err error) {
	varOrgCostBudgetWriteDto := _OrgCostBudgetWriteDto{}

	err = json.Unmarshal(data, &varOrgCostBudgetWriteDto)

	if err != nil {
		return err
	}

	*o = OrgCostBudgetWriteDto(varOrgCostBudgetWriteDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total_budget")
		delete(additionalProperties, "used_budget")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCostBudgetWriteDto struct {
	value *OrgCostBudgetWriteDto
	isSet bool
}

func (v NullableOrgCostBudgetWriteDto) Get() *OrgCostBudgetWriteDto {
	return v.value
}

func (v *NullableOrgCostBudgetWriteDto) Set(val *OrgCostBudgetWriteDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCostBudgetWriteDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCostBudgetWriteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCostBudgetWriteDto(val *OrgCostBudgetWriteDto) *NullableOrgCostBudgetWriteDto {
	return &NullableOrgCostBudgetWriteDto{value: val, isSet: true}
}

func (v NullableOrgCostBudgetWriteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCostBudgetWriteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


