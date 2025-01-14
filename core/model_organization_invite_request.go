/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the OrganizationInviteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationInviteRequest{}

// OrganizationInviteRequest struct for OrganizationInviteRequest
type OrganizationInviteRequest struct {
	Email string `json:"email"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationInviteRequest OrganizationInviteRequest

// NewOrganizationInviteRequest instantiates a new OrganizationInviteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInviteRequest(email string) *OrganizationInviteRequest {
	this := OrganizationInviteRequest{}
	this.Email = email
	return &this
}

// NewOrganizationInviteRequestWithDefaults instantiates a new OrganizationInviteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInviteRequestWithDefaults() *OrganizationInviteRequest {
	this := OrganizationInviteRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *OrganizationInviteRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OrganizationInviteRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *OrganizationInviteRequest) SetEmail(v string) {
	o.Email = v
}

func (o OrganizationInviteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationInviteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationInviteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varOrganizationInviteRequest := _OrganizationInviteRequest{}

	err = json.Unmarshal(data, &varOrganizationInviteRequest)

	if err != nil {
		return err
	}

	*o = OrganizationInviteRequest(varOrganizationInviteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationInviteRequest struct {
	value *OrganizationInviteRequest
	isSet bool
}

func (v NullableOrganizationInviteRequest) Get() *OrganizationInviteRequest {
	return v.value
}

func (v *NullableOrganizationInviteRequest) Set(val *OrganizationInviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInviteRequest(val *OrganizationInviteRequest) *NullableOrganizationInviteRequest {
	return &NullableOrganizationInviteRequest{value: val, isSet: true}
}

func (v NullableOrganizationInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


