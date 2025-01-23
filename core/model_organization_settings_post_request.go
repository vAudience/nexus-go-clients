/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
)

// checks if the OrganizationSettingsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSettingsPostRequest{}

// OrganizationSettingsPostRequest struct for OrganizationSettingsPostRequest
type OrganizationSettingsPostRequest struct {
	DefaultAgentId *string `json:"defaultAgentId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationSettingsPostRequest OrganizationSettingsPostRequest

// NewOrganizationSettingsPostRequest instantiates a new OrganizationSettingsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSettingsPostRequest() *OrganizationSettingsPostRequest {
	this := OrganizationSettingsPostRequest{}
	return &this
}

// NewOrganizationSettingsPostRequestWithDefaults instantiates a new OrganizationSettingsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSettingsPostRequestWithDefaults() *OrganizationSettingsPostRequest {
	this := OrganizationSettingsPostRequest{}
	return &this
}

// GetDefaultAgentId returns the DefaultAgentId field value if set, zero value otherwise.
func (o *OrganizationSettingsPostRequest) GetDefaultAgentId() string {
	if o == nil || IsNil(o.DefaultAgentId) {
		var ret string
		return ret
	}
	return *o.DefaultAgentId
}

// GetDefaultAgentIdOk returns a tuple with the DefaultAgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettingsPostRequest) GetDefaultAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAgentId) {
		return nil, false
	}
	return o.DefaultAgentId, true
}

// HasDefaultAgentId returns a boolean if a field has been set.
func (o *OrganizationSettingsPostRequest) HasDefaultAgentId() bool {
	if o != nil && !IsNil(o.DefaultAgentId) {
		return true
	}

	return false
}

// SetDefaultAgentId gets a reference to the given string and assigns it to the DefaultAgentId field.
func (o *OrganizationSettingsPostRequest) SetDefaultAgentId(v string) {
	o.DefaultAgentId = &v
}

func (o OrganizationSettingsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSettingsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultAgentId) {
		toSerialize["defaultAgentId"] = o.DefaultAgentId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationSettingsPostRequest) UnmarshalJSON(data []byte) (err error) {
	varOrganizationSettingsPostRequest := _OrganizationSettingsPostRequest{}

	err = json.Unmarshal(data, &varOrganizationSettingsPostRequest)

	if err != nil {
		return err
	}

	*o = OrganizationSettingsPostRequest(varOrganizationSettingsPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultAgentId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationSettingsPostRequest struct {
	value *OrganizationSettingsPostRequest
	isSet bool
}

func (v NullableOrganizationSettingsPostRequest) Get() *OrganizationSettingsPostRequest {
	return v.value
}

func (v *NullableOrganizationSettingsPostRequest) Set(val *OrganizationSettingsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSettingsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSettingsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSettingsPostRequest(val *OrganizationSettingsPostRequest) *NullableOrganizationSettingsPostRequest {
	return &NullableOrganizationSettingsPostRequest{value: val, isSet: true}
}

func (v NullableOrganizationSettingsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSettingsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


