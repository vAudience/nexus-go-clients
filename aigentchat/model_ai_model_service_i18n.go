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

// checks if the AIModelServiceI18n type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIModelServiceI18n{}

// AIModelServiceI18n struct for AIModelServiceI18n
type AIModelServiceI18n struct {
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AIModelServiceI18n AIModelServiceI18n

// NewAIModelServiceI18n instantiates a new AIModelServiceI18n object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIModelServiceI18n(name string) *AIModelServiceI18n {
	this := AIModelServiceI18n{}
	this.Name = name
	return &this
}

// NewAIModelServiceI18nWithDefaults instantiates a new AIModelServiceI18n object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIModelServiceI18nWithDefaults() *AIModelServiceI18n {
	this := AIModelServiceI18n{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AIModelServiceI18n) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceI18n) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AIModelServiceI18n) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AIModelServiceI18n) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *AIModelServiceI18n) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AIModelServiceI18n) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AIModelServiceI18n) SetName(v string) {
	o.Name = v
}

func (o AIModelServiceI18n) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIModelServiceI18n) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIModelServiceI18n) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAIModelServiceI18n := _AIModelServiceI18n{}

	err = json.Unmarshal(data, &varAIModelServiceI18n)

	if err != nil {
		return err
	}

	*o = AIModelServiceI18n(varAIModelServiceI18n)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIModelServiceI18n struct {
	value *AIModelServiceI18n
	isSet bool
}

func (v NullableAIModelServiceI18n) Get() *AIModelServiceI18n {
	return v.value
}

func (v *NullableAIModelServiceI18n) Set(val *AIModelServiceI18n) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelServiceI18n) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelServiceI18n) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelServiceI18n(val *AIModelServiceI18n) *NullableAIModelServiceI18n {
	return &NullableAIModelServiceI18n{value: val, isSet: true}
}

func (v NullableAIModelServiceI18n) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelServiceI18n) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


