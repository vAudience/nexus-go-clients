/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.17.8
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the OrgCostBudgetCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCostBudgetCheck{}

// OrgCostBudgetCheck struct for OrgCostBudgetCheck
type OrgCostBudgetCheck struct {
	OrgId string `json:"org_id"`
	SufficientBudget bool `json:"sufficient_budget"`
	AdditionalProperties map[string]interface{}
}

type _OrgCostBudgetCheck OrgCostBudgetCheck

// NewOrgCostBudgetCheck instantiates a new OrgCostBudgetCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCostBudgetCheck(orgId string, sufficientBudget bool) *OrgCostBudgetCheck {
	this := OrgCostBudgetCheck{}
	this.OrgId = orgId
	this.SufficientBudget = sufficientBudget
	return &this
}

// NewOrgCostBudgetCheckWithDefaults instantiates a new OrgCostBudgetCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCostBudgetCheckWithDefaults() *OrgCostBudgetCheck {
	this := OrgCostBudgetCheck{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *OrgCostBudgetCheck) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *OrgCostBudgetCheck) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *OrgCostBudgetCheck) SetOrgId(v string) {
	o.OrgId = v
}

// GetSufficientBudget returns the SufficientBudget field value
func (o *OrgCostBudgetCheck) GetSufficientBudget() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SufficientBudget
}

// GetSufficientBudgetOk returns a tuple with the SufficientBudget field value
// and a boolean to check if the value has been set.
func (o *OrgCostBudgetCheck) GetSufficientBudgetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SufficientBudget, true
}

// SetSufficientBudget sets field value
func (o *OrgCostBudgetCheck) SetSufficientBudget(v bool) {
	o.SufficientBudget = v
}

func (o OrgCostBudgetCheck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCostBudgetCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	toSerialize["sufficient_budget"] = o.SufficientBudget

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCostBudgetCheck) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
		"sufficient_budget",
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

	varOrgCostBudgetCheck := _OrgCostBudgetCheck{}

	err = json.Unmarshal(data, &varOrgCostBudgetCheck)

	if err != nil {
		return err
	}

	*o = OrgCostBudgetCheck(varOrgCostBudgetCheck)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "sufficient_budget")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCostBudgetCheck struct {
	value *OrgCostBudgetCheck
	isSet bool
}

func (v NullableOrgCostBudgetCheck) Get() *OrgCostBudgetCheck {
	return v.value
}

func (v *NullableOrgCostBudgetCheck) Set(val *OrgCostBudgetCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCostBudgetCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCostBudgetCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCostBudgetCheck(val *OrgCostBudgetCheck) *NullableOrgCostBudgetCheck {
	return &NullableOrgCostBudgetCheck{value: val, isSet: true}
}

func (v NullableOrgCostBudgetCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCostBudgetCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


