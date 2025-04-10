/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the UserSettingsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSettingsPostRequest{}

// UserSettingsPostRequest struct for UserSettingsPostRequest
type UserSettingsPostRequest struct {
	ColorScheme string `json:"colorScheme"`
	Language string `json:"language"`
	UserId *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSettingsPostRequest UserSettingsPostRequest

// NewUserSettingsPostRequest instantiates a new UserSettingsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettingsPostRequest(colorScheme string, language string) *UserSettingsPostRequest {
	this := UserSettingsPostRequest{}
	this.ColorScheme = colorScheme
	this.Language = language
	return &this
}

// NewUserSettingsPostRequestWithDefaults instantiates a new UserSettingsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsPostRequestWithDefaults() *UserSettingsPostRequest {
	this := UserSettingsPostRequest{}
	return &this
}

// GetColorScheme returns the ColorScheme field value
func (o *UserSettingsPostRequest) GetColorScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ColorScheme
}

// GetColorSchemeOk returns a tuple with the ColorScheme field value
// and a boolean to check if the value has been set.
func (o *UserSettingsPostRequest) GetColorSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColorScheme, true
}

// SetColorScheme sets field value
func (o *UserSettingsPostRequest) SetColorScheme(v string) {
	o.ColorScheme = v
}

// GetLanguage returns the Language field value
func (o *UserSettingsPostRequest) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *UserSettingsPostRequest) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *UserSettingsPostRequest) SetLanguage(v string) {
	o.Language = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserSettingsPostRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsPostRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserSettingsPostRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserSettingsPostRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o UserSettingsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSettingsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["colorScheme"] = o.ColorScheme
	toSerialize["language"] = o.Language
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSettingsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"colorScheme",
		"language",
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

	varUserSettingsPostRequest := _UserSettingsPostRequest{}

	err = json.Unmarshal(data, &varUserSettingsPostRequest)

	if err != nil {
		return err
	}

	*o = UserSettingsPostRequest(varUserSettingsPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "colorScheme")
		delete(additionalProperties, "language")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSettingsPostRequest struct {
	value *UserSettingsPostRequest
	isSet bool
}

func (v NullableUserSettingsPostRequest) Get() *UserSettingsPostRequest {
	return v.value
}

func (v *NullableUserSettingsPostRequest) Set(val *UserSettingsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettingsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettingsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettingsPostRequest(val *UserSettingsPostRequest) *NullableUserSettingsPostRequest {
	return &NullableUserSettingsPostRequest{value: val, isSet: true}
}

func (v NullableUserSettingsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettingsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


