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

// checks if the AIModelConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIModelConstraint{}

// AIModelConstraint struct for AIModelConstraint
type AIModelConstraint struct {
	Direction *AIModelConstraintDirection `json:"direction,omitempty"`
	Max *float32 `json:"max,omitempty"`
	Min *float32 `json:"min,omitempty"`
	Unit *AIModelMinMaxUnit `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIModelConstraint AIModelConstraint

// NewAIModelConstraint instantiates a new AIModelConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIModelConstraint() *AIModelConstraint {
	this := AIModelConstraint{}
	return &this
}

// NewAIModelConstraintWithDefaults instantiates a new AIModelConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIModelConstraintWithDefaults() *AIModelConstraint {
	this := AIModelConstraint{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *AIModelConstraint) GetDirection() AIModelConstraintDirection {
	if o == nil || IsNil(o.Direction) {
		var ret AIModelConstraintDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelConstraint) GetDirectionOk() (*AIModelConstraintDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *AIModelConstraint) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given AIModelConstraintDirection and assigns it to the Direction field.
func (o *AIModelConstraint) SetDirection(v AIModelConstraintDirection) {
	o.Direction = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *AIModelConstraint) GetMax() float32 {
	if o == nil || IsNil(o.Max) {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelConstraint) GetMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *AIModelConstraint) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *AIModelConstraint) SetMax(v float32) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *AIModelConstraint) GetMin() float32 {
	if o == nil || IsNil(o.Min) {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelConstraint) GetMinOk() (*float32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *AIModelConstraint) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *AIModelConstraint) SetMin(v float32) {
	o.Min = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *AIModelConstraint) GetUnit() AIModelMinMaxUnit {
	if o == nil || IsNil(o.Unit) {
		var ret AIModelMinMaxUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelConstraint) GetUnitOk() (*AIModelMinMaxUnit, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *AIModelConstraint) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given AIModelMinMaxUnit and assigns it to the Unit field.
func (o *AIModelConstraint) SetUnit(v AIModelMinMaxUnit) {
	o.Unit = &v
}

func (o AIModelConstraint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIModelConstraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIModelConstraint) UnmarshalJSON(data []byte) (err error) {
	varAIModelConstraint := _AIModelConstraint{}

	err = json.Unmarshal(data, &varAIModelConstraint)

	if err != nil {
		return err
	}

	*o = AIModelConstraint(varAIModelConstraint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "max")
		delete(additionalProperties, "min")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIModelConstraint struct {
	value *AIModelConstraint
	isSet bool
}

func (v NullableAIModelConstraint) Get() *AIModelConstraint {
	return v.value
}

func (v *NullableAIModelConstraint) Set(val *AIModelConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelConstraint(val *AIModelConstraint) *NullableAIModelConstraint {
	return &NullableAIModelConstraint{value: val, isSet: true}
}

func (v NullableAIModelConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


