/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.18.2
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the OrgCostBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCostBudget{}

// OrgCostBudget struct for OrgCostBudget
type OrgCostBudget struct {
	OrgId string `json:"org_id"`
	RemainingBudget *float64 `json:"remaining_budget,omitempty"`
	TotalBudget *float64 `json:"total_budget,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	UsedBudget *float64 `json:"used_budget,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgCostBudget OrgCostBudget

// NewOrgCostBudget instantiates a new OrgCostBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCostBudget(orgId string) *OrgCostBudget {
	this := OrgCostBudget{}
	this.OrgId = orgId
	return &this
}

// NewOrgCostBudgetWithDefaults instantiates a new OrgCostBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCostBudgetWithDefaults() *OrgCostBudget {
	this := OrgCostBudget{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *OrgCostBudget) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *OrgCostBudget) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *OrgCostBudget) SetOrgId(v string) {
	o.OrgId = v
}

// GetRemainingBudget returns the RemainingBudget field value if set, zero value otherwise.
func (o *OrgCostBudget) GetRemainingBudget() float64 {
	if o == nil || IsNil(o.RemainingBudget) {
		var ret float64
		return ret
	}
	return *o.RemainingBudget
}

// GetRemainingBudgetOk returns a tuple with the RemainingBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCostBudget) GetRemainingBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.RemainingBudget) {
		return nil, false
	}
	return o.RemainingBudget, true
}

// HasRemainingBudget returns a boolean if a field has been set.
func (o *OrgCostBudget) HasRemainingBudget() bool {
	if o != nil && !IsNil(o.RemainingBudget) {
		return true
	}

	return false
}

// SetRemainingBudget gets a reference to the given float64 and assigns it to the RemainingBudget field.
func (o *OrgCostBudget) SetRemainingBudget(v float64) {
	o.RemainingBudget = &v
}

// GetTotalBudget returns the TotalBudget field value if set, zero value otherwise.
func (o *OrgCostBudget) GetTotalBudget() float64 {
	if o == nil || IsNil(o.TotalBudget) {
		var ret float64
		return ret
	}
	return *o.TotalBudget
}

// GetTotalBudgetOk returns a tuple with the TotalBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCostBudget) GetTotalBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalBudget) {
		return nil, false
	}
	return o.TotalBudget, true
}

// HasTotalBudget returns a boolean if a field has been set.
func (o *OrgCostBudget) HasTotalBudget() bool {
	if o != nil && !IsNil(o.TotalBudget) {
		return true
	}

	return false
}

// SetTotalBudget gets a reference to the given float64 and assigns it to the TotalBudget field.
func (o *OrgCostBudget) SetTotalBudget(v float64) {
	o.TotalBudget = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrgCostBudget) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCostBudget) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrgCostBudget) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *OrgCostBudget) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *OrgCostBudget) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCostBudget) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *OrgCostBudget) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *OrgCostBudget) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUsedBudget returns the UsedBudget field value if set, zero value otherwise.
func (o *OrgCostBudget) GetUsedBudget() float64 {
	if o == nil || IsNil(o.UsedBudget) {
		var ret float64
		return ret
	}
	return *o.UsedBudget
}

// GetUsedBudgetOk returns a tuple with the UsedBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCostBudget) GetUsedBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.UsedBudget) {
		return nil, false
	}
	return o.UsedBudget, true
}

// HasUsedBudget returns a boolean if a field has been set.
func (o *OrgCostBudget) HasUsedBudget() bool {
	if o != nil && !IsNil(o.UsedBudget) {
		return true
	}

	return false
}

// SetUsedBudget gets a reference to the given float64 and assigns it to the UsedBudget field.
func (o *OrgCostBudget) SetUsedBudget(v float64) {
	o.UsedBudget = &v
}

func (o OrgCostBudget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCostBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	if !IsNil(o.RemainingBudget) {
		toSerialize["remaining_budget"] = o.RemainingBudget
	}
	if !IsNil(o.TotalBudget) {
		toSerialize["total_budget"] = o.TotalBudget
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if !IsNil(o.UsedBudget) {
		toSerialize["used_budget"] = o.UsedBudget
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCostBudget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
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

	varOrgCostBudget := _OrgCostBudget{}

	err = json.Unmarshal(data, &varOrgCostBudget)

	if err != nil {
		return err
	}

	*o = OrgCostBudget(varOrgCostBudget)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "remaining_budget")
		delete(additionalProperties, "total_budget")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		delete(additionalProperties, "used_budget")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCostBudget struct {
	value *OrgCostBudget
	isSet bool
}

func (v NullableOrgCostBudget) Get() *OrgCostBudget {
	return v.value
}

func (v *NullableOrgCostBudget) Set(val *OrgCostBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCostBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCostBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCostBudget(val *OrgCostBudget) *NullableOrgCostBudget {
	return &NullableOrgCostBudget{value: val, isSet: true}
}

func (v NullableOrgCostBudget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCostBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


