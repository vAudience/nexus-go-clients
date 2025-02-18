/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.9
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the Ability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ability{}

// Ability struct for Ability
type Ability struct {
	Constraints []AIModelConstraint `json:"constraints,omitempty"`
	Type *AbilityType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ability Ability

// NewAbility instantiates a new Ability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbility() *Ability {
	this := Ability{}
	return &this
}

// NewAbilityWithDefaults instantiates a new Ability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbilityWithDefaults() *Ability {
	this := Ability{}
	return &this
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *Ability) GetConstraints() []AIModelConstraint {
	if o == nil || IsNil(o.Constraints) {
		var ret []AIModelConstraint
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ability) GetConstraintsOk() ([]AIModelConstraint, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *Ability) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []AIModelConstraint and assigns it to the Constraints field.
func (o *Ability) SetConstraints(v []AIModelConstraint) {
	o.Constraints = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Ability) GetType() AbilityType {
	if o == nil || IsNil(o.Type) {
		var ret AbilityType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ability) GetTypeOk() (*AbilityType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Ability) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AbilityType and assigns it to the Type field.
func (o *Ability) SetType(v AbilityType) {
	o.Type = &v
}

func (o Ability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ability) UnmarshalJSON(data []byte) (err error) {
	varAbility := _Ability{}

	err = json.Unmarshal(data, &varAbility)

	if err != nil {
		return err
	}

	*o = Ability(varAbility)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAbility struct {
	value *Ability
	isSet bool
}

func (v NullableAbility) Get() *Ability {
	return v.value
}

func (v *NullableAbility) Set(val *Ability) {
	v.value = val
	v.isSet = true
}

func (v NullableAbility) IsSet() bool {
	return v.isSet
}

func (v *NullableAbility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbility(val *Ability) *NullableAbility {
	return &NullableAbility{value: val, isSet: true}
}

func (v NullableAbility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


