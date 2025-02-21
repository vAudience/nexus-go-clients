/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the EntitlementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementResponse{}

// EntitlementResponse struct for EntitlementResponse
type EntitlementResponse struct {
	Enabled bool `json:"enabled"`
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementResponse EntitlementResponse

// NewEntitlementResponse instantiates a new EntitlementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementResponse(enabled bool, status string) *EntitlementResponse {
	this := EntitlementResponse{}
	this.Enabled = enabled
	this.Status = status
	return &this
}

// NewEntitlementResponseWithDefaults instantiates a new EntitlementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementResponseWithDefaults() *EntitlementResponse {
	this := EntitlementResponse{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *EntitlementResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EntitlementResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EntitlementResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetStatus returns the Status field value
func (o *EntitlementResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EntitlementResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EntitlementResponse) SetStatus(v string) {
	o.Status = v
}

func (o EntitlementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"status",
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

	varEntitlementResponse := _EntitlementResponse{}

	err = json.Unmarshal(data, &varEntitlementResponse)

	if err != nil {
		return err
	}

	*o = EntitlementResponse(varEntitlementResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementResponse struct {
	value *EntitlementResponse
	isSet bool
}

func (v NullableEntitlementResponse) Get() *EntitlementResponse {
	return v.value
}

func (v *NullableEntitlementResponse) Set(val *EntitlementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementResponse(val *EntitlementResponse) *NullableEntitlementResponse {
	return &NullableEntitlementResponse{value: val, isSet: true}
}

func (v NullableEntitlementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


