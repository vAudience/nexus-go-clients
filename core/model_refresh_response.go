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

// checks if the RefreshResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshResponse{}

// RefreshResponse struct for RefreshResponse
type RefreshResponse struct {
	TokenExpiresInSec int32 `json:"tokenExpiresInSec"`
	AdditionalProperties map[string]interface{}
}

type _RefreshResponse RefreshResponse

// NewRefreshResponse instantiates a new RefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshResponse(tokenExpiresInSec int32) *RefreshResponse {
	this := RefreshResponse{}
	this.TokenExpiresInSec = tokenExpiresInSec
	return &this
}

// NewRefreshResponseWithDefaults instantiates a new RefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshResponseWithDefaults() *RefreshResponse {
	this := RefreshResponse{}
	return &this
}

// GetTokenExpiresInSec returns the TokenExpiresInSec field value
func (o *RefreshResponse) GetTokenExpiresInSec() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenExpiresInSec
}

// GetTokenExpiresInSecOk returns a tuple with the TokenExpiresInSec field value
// and a boolean to check if the value has been set.
func (o *RefreshResponse) GetTokenExpiresInSecOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenExpiresInSec, true
}

// SetTokenExpiresInSec sets field value
func (o *RefreshResponse) SetTokenExpiresInSec(v int32) {
	o.TokenExpiresInSec = v
}

func (o RefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tokenExpiresInSec"] = o.TokenExpiresInSec

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RefreshResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tokenExpiresInSec",
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

	varRefreshResponse := _RefreshResponse{}

	err = json.Unmarshal(data, &varRefreshResponse)

	if err != nil {
		return err
	}

	*o = RefreshResponse(varRefreshResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tokenExpiresInSec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRefreshResponse struct {
	value *RefreshResponse
	isSet bool
}

func (v NullableRefreshResponse) Get() *RefreshResponse {
	return v.value
}

func (v *NullableRefreshResponse) Set(val *RefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshResponse(val *RefreshResponse) *NullableRefreshResponse {
	return &NullableRefreshResponse{value: val, isSet: true}
}

func (v NullableRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


