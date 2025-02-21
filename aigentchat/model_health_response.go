/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.10
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the HealthResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthResponse{}

// HealthResponse struct for HealthResponse
type HealthResponse struct {
	Health *string `json:"health,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HealthResponse HealthResponse

// NewHealthResponse instantiates a new HealthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthResponse() *HealthResponse {
	this := HealthResponse{}
	return &this
}

// NewHealthResponseWithDefaults instantiates a new HealthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthResponseWithDefaults() *HealthResponse {
	this := HealthResponse{}
	return &this
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *HealthResponse) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthResponse) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *HealthResponse) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *HealthResponse) SetHealth(v string) {
	o.Health = &v
}

func (o HealthResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HealthResponse) UnmarshalJSON(data []byte) (err error) {
	varHealthResponse := _HealthResponse{}

	err = json.Unmarshal(data, &varHealthResponse)

	if err != nil {
		return err
	}

	*o = HealthResponse(varHealthResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "health")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHealthResponse struct {
	value *HealthResponse
	isSet bool
}

func (v NullableHealthResponse) Get() *HealthResponse {
	return v.value
}

func (v *NullableHealthResponse) Set(val *HealthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthResponse(val *HealthResponse) *NullableHealthResponse {
	return &NullableHealthResponse{value: val, isSet: true}
}

func (v NullableHealthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


