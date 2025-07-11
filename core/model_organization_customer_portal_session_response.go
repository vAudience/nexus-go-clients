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

// checks if the OrganizationCustomerPortalSessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCustomerPortalSessionResponse{}

// OrganizationCustomerPortalSessionResponse struct for OrganizationCustomerPortalSessionResponse
type OrganizationCustomerPortalSessionResponse struct {
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationCustomerPortalSessionResponse OrganizationCustomerPortalSessionResponse

// NewOrganizationCustomerPortalSessionResponse instantiates a new OrganizationCustomerPortalSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCustomerPortalSessionResponse(url string) *OrganizationCustomerPortalSessionResponse {
	this := OrganizationCustomerPortalSessionResponse{}
	this.Url = url
	return &this
}

// NewOrganizationCustomerPortalSessionResponseWithDefaults instantiates a new OrganizationCustomerPortalSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCustomerPortalSessionResponseWithDefaults() *OrganizationCustomerPortalSessionResponse {
	this := OrganizationCustomerPortalSessionResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *OrganizationCustomerPortalSessionResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OrganizationCustomerPortalSessionResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OrganizationCustomerPortalSessionResponse) SetUrl(v string) {
	o.Url = v
}

func (o OrganizationCustomerPortalSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCustomerPortalSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationCustomerPortalSessionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varOrganizationCustomerPortalSessionResponse := _OrganizationCustomerPortalSessionResponse{}

	err = json.Unmarshal(data, &varOrganizationCustomerPortalSessionResponse)

	if err != nil {
		return err
	}

	*o = OrganizationCustomerPortalSessionResponse(varOrganizationCustomerPortalSessionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationCustomerPortalSessionResponse struct {
	value *OrganizationCustomerPortalSessionResponse
	isSet bool
}

func (v NullableOrganizationCustomerPortalSessionResponse) Get() *OrganizationCustomerPortalSessionResponse {
	return v.value
}

func (v *NullableOrganizationCustomerPortalSessionResponse) Set(val *OrganizationCustomerPortalSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCustomerPortalSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCustomerPortalSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCustomerPortalSessionResponse(val *OrganizationCustomerPortalSessionResponse) *NullableOrganizationCustomerPortalSessionResponse {
	return &NullableOrganizationCustomerPortalSessionResponse{value: val, isSet: true}
}

func (v NullableOrganizationCustomerPortalSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCustomerPortalSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


