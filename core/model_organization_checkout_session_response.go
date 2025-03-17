/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the OrganizationCheckoutSessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCheckoutSessionResponse{}

// OrganizationCheckoutSessionResponse struct for OrganizationCheckoutSessionResponse
type OrganizationCheckoutSessionResponse struct {
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationCheckoutSessionResponse OrganizationCheckoutSessionResponse

// NewOrganizationCheckoutSessionResponse instantiates a new OrganizationCheckoutSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCheckoutSessionResponse(url string) *OrganizationCheckoutSessionResponse {
	this := OrganizationCheckoutSessionResponse{}
	this.Url = url
	return &this
}

// NewOrganizationCheckoutSessionResponseWithDefaults instantiates a new OrganizationCheckoutSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCheckoutSessionResponseWithDefaults() *OrganizationCheckoutSessionResponse {
	this := OrganizationCheckoutSessionResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *OrganizationCheckoutSessionResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OrganizationCheckoutSessionResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OrganizationCheckoutSessionResponse) SetUrl(v string) {
	o.Url = v
}

func (o OrganizationCheckoutSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCheckoutSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationCheckoutSessionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varOrganizationCheckoutSessionResponse := _OrganizationCheckoutSessionResponse{}

	err = json.Unmarshal(data, &varOrganizationCheckoutSessionResponse)

	if err != nil {
		return err
	}

	*o = OrganizationCheckoutSessionResponse(varOrganizationCheckoutSessionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationCheckoutSessionResponse struct {
	value *OrganizationCheckoutSessionResponse
	isSet bool
}

func (v NullableOrganizationCheckoutSessionResponse) Get() *OrganizationCheckoutSessionResponse {
	return v.value
}

func (v *NullableOrganizationCheckoutSessionResponse) Set(val *OrganizationCheckoutSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCheckoutSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCheckoutSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCheckoutSessionResponse(val *OrganizationCheckoutSessionResponse) *NullableOrganizationCheckoutSessionResponse {
	return &NullableOrganizationCheckoutSessionResponse{value: val, isSet: true}
}

func (v NullableOrganizationCheckoutSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCheckoutSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


