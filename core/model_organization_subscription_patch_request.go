/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the OrganizationSubscriptionPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSubscriptionPatchRequest{}

// OrganizationSubscriptionPatchRequest struct for OrganizationSubscriptionPatchRequest
type OrganizationSubscriptionPatchRequest struct {
	PeriodEnd *string `json:"periodEnd,omitempty"`
	Seats int32 `json:"seats"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationSubscriptionPatchRequest OrganizationSubscriptionPatchRequest

// NewOrganizationSubscriptionPatchRequest instantiates a new OrganizationSubscriptionPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSubscriptionPatchRequest(seats int32) *OrganizationSubscriptionPatchRequest {
	this := OrganizationSubscriptionPatchRequest{}
	this.Seats = seats
	return &this
}

// NewOrganizationSubscriptionPatchRequestWithDefaults instantiates a new OrganizationSubscriptionPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSubscriptionPatchRequestWithDefaults() *OrganizationSubscriptionPatchRequest {
	this := OrganizationSubscriptionPatchRequest{}
	return &this
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *OrganizationSubscriptionPatchRequest) GetPeriodEnd() string {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret string
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionPatchRequest) GetPeriodEndOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *OrganizationSubscriptionPatchRequest) HasPeriodEnd() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given string and assigns it to the PeriodEnd field.
func (o *OrganizationSubscriptionPatchRequest) SetPeriodEnd(v string) {
	o.PeriodEnd = &v
}

// GetSeats returns the Seats field value
func (o *OrganizationSubscriptionPatchRequest) GetSeats() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Seats
}

// GetSeatsOk returns a tuple with the Seats field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionPatchRequest) GetSeatsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seats, true
}

// SetSeats sets field value
func (o *OrganizationSubscriptionPatchRequest) SetSeats(v int32) {
	o.Seats = v
}

func (o OrganizationSubscriptionPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSubscriptionPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PeriodEnd) {
		toSerialize["periodEnd"] = o.PeriodEnd
	}
	toSerialize["seats"] = o.Seats

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationSubscriptionPatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seats",
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

	varOrganizationSubscriptionPatchRequest := _OrganizationSubscriptionPatchRequest{}

	err = json.Unmarshal(data, &varOrganizationSubscriptionPatchRequest)

	if err != nil {
		return err
	}

	*o = OrganizationSubscriptionPatchRequest(varOrganizationSubscriptionPatchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "periodEnd")
		delete(additionalProperties, "seats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationSubscriptionPatchRequest struct {
	value *OrganizationSubscriptionPatchRequest
	isSet bool
}

func (v NullableOrganizationSubscriptionPatchRequest) Get() *OrganizationSubscriptionPatchRequest {
	return v.value
}

func (v *NullableOrganizationSubscriptionPatchRequest) Set(val *OrganizationSubscriptionPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSubscriptionPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSubscriptionPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSubscriptionPatchRequest(val *OrganizationSubscriptionPatchRequest) *NullableOrganizationSubscriptionPatchRequest {
	return &NullableOrganizationSubscriptionPatchRequest{value: val, isSet: true}
}

func (v NullableOrganizationSubscriptionPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSubscriptionPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


