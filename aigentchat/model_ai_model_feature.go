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

// checks if the AIModelFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIModelFeature{}

// AIModelFeature struct for AIModelFeature
type AIModelFeature struct {
	Capability *AIModelCapability `json:"capability,omitempty"`
	Constraints []AIModelConstraint `json:"constraints,omitempty"`
	CostItemTemplates []ExecutionCostTemplate `json:"cost_item_templates,omitempty"`
	CostItems []ExecutionUsageCost `json:"cost_items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIModelFeature AIModelFeature

// NewAIModelFeature instantiates a new AIModelFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIModelFeature() *AIModelFeature {
	this := AIModelFeature{}
	return &this
}

// NewAIModelFeatureWithDefaults instantiates a new AIModelFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIModelFeatureWithDefaults() *AIModelFeature {
	this := AIModelFeature{}
	return &this
}

// GetCapability returns the Capability field value if set, zero value otherwise.
func (o *AIModelFeature) GetCapability() AIModelCapability {
	if o == nil || IsNil(o.Capability) {
		var ret AIModelCapability
		return ret
	}
	return *o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelFeature) GetCapabilityOk() (*AIModelCapability, bool) {
	if o == nil || IsNil(o.Capability) {
		return nil, false
	}
	return o.Capability, true
}

// HasCapability returns a boolean if a field has been set.
func (o *AIModelFeature) HasCapability() bool {
	if o != nil && !IsNil(o.Capability) {
		return true
	}

	return false
}

// SetCapability gets a reference to the given AIModelCapability and assigns it to the Capability field.
func (o *AIModelFeature) SetCapability(v AIModelCapability) {
	o.Capability = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *AIModelFeature) GetConstraints() []AIModelConstraint {
	if o == nil || IsNil(o.Constraints) {
		var ret []AIModelConstraint
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelFeature) GetConstraintsOk() ([]AIModelConstraint, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *AIModelFeature) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []AIModelConstraint and assigns it to the Constraints field.
func (o *AIModelFeature) SetConstraints(v []AIModelConstraint) {
	o.Constraints = v
}

// GetCostItemTemplates returns the CostItemTemplates field value if set, zero value otherwise.
func (o *AIModelFeature) GetCostItemTemplates() []ExecutionCostTemplate {
	if o == nil || IsNil(o.CostItemTemplates) {
		var ret []ExecutionCostTemplate
		return ret
	}
	return o.CostItemTemplates
}

// GetCostItemTemplatesOk returns a tuple with the CostItemTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelFeature) GetCostItemTemplatesOk() ([]ExecutionCostTemplate, bool) {
	if o == nil || IsNil(o.CostItemTemplates) {
		return nil, false
	}
	return o.CostItemTemplates, true
}

// HasCostItemTemplates returns a boolean if a field has been set.
func (o *AIModelFeature) HasCostItemTemplates() bool {
	if o != nil && !IsNil(o.CostItemTemplates) {
		return true
	}

	return false
}

// SetCostItemTemplates gets a reference to the given []ExecutionCostTemplate and assigns it to the CostItemTemplates field.
func (o *AIModelFeature) SetCostItemTemplates(v []ExecutionCostTemplate) {
	o.CostItemTemplates = v
}

// GetCostItems returns the CostItems field value if set, zero value otherwise.
func (o *AIModelFeature) GetCostItems() []ExecutionUsageCost {
	if o == nil || IsNil(o.CostItems) {
		var ret []ExecutionUsageCost
		return ret
	}
	return o.CostItems
}

// GetCostItemsOk returns a tuple with the CostItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelFeature) GetCostItemsOk() ([]ExecutionUsageCost, bool) {
	if o == nil || IsNil(o.CostItems) {
		return nil, false
	}
	return o.CostItems, true
}

// HasCostItems returns a boolean if a field has been set.
func (o *AIModelFeature) HasCostItems() bool {
	if o != nil && !IsNil(o.CostItems) {
		return true
	}

	return false
}

// SetCostItems gets a reference to the given []ExecutionUsageCost and assigns it to the CostItems field.
func (o *AIModelFeature) SetCostItems(v []ExecutionUsageCost) {
	o.CostItems = v
}

func (o AIModelFeature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIModelFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capability) {
		toSerialize["capability"] = o.Capability
	}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	if !IsNil(o.CostItemTemplates) {
		toSerialize["cost_item_templates"] = o.CostItemTemplates
	}
	if !IsNil(o.CostItems) {
		toSerialize["cost_items"] = o.CostItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIModelFeature) UnmarshalJSON(data []byte) (err error) {
	varAIModelFeature := _AIModelFeature{}

	err = json.Unmarshal(data, &varAIModelFeature)

	if err != nil {
		return err
	}

	*o = AIModelFeature(varAIModelFeature)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capability")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "cost_item_templates")
		delete(additionalProperties, "cost_items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIModelFeature struct {
	value *AIModelFeature
	isSet bool
}

func (v NullableAIModelFeature) Get() *AIModelFeature {
	return v.value
}

func (v *NullableAIModelFeature) Set(val *AIModelFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelFeature(val *AIModelFeature) *NullableAIModelFeature {
	return &NullableAIModelFeature{value: val, isSet: true}
}

func (v NullableAIModelFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


