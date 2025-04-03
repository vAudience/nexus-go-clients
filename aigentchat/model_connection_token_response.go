/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.17.8
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the ConnectionTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionTokenResponse{}

// ConnectionTokenResponse struct for ConnectionTokenResponse
type ConnectionTokenResponse struct {
	Token *string `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionTokenResponse ConnectionTokenResponse

// NewConnectionTokenResponse instantiates a new ConnectionTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionTokenResponse() *ConnectionTokenResponse {
	this := ConnectionTokenResponse{}
	return &this
}

// NewConnectionTokenResponseWithDefaults instantiates a new ConnectionTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionTokenResponseWithDefaults() *ConnectionTokenResponse {
	this := ConnectionTokenResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ConnectionTokenResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ConnectionTokenResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ConnectionTokenResponse) SetToken(v string) {
	o.Token = &v
}

func (o ConnectionTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionTokenResponse) UnmarshalJSON(data []byte) (err error) {
	varConnectionTokenResponse := _ConnectionTokenResponse{}

	err = json.Unmarshal(data, &varConnectionTokenResponse)

	if err != nil {
		return err
	}

	*o = ConnectionTokenResponse(varConnectionTokenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionTokenResponse struct {
	value *ConnectionTokenResponse
	isSet bool
}

func (v NullableConnectionTokenResponse) Get() *ConnectionTokenResponse {
	return v.value
}

func (v *NullableConnectionTokenResponse) Set(val *ConnectionTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionTokenResponse(val *ConnectionTokenResponse) *NullableConnectionTokenResponse {
	return &NullableConnectionTokenResponse{value: val, isSet: true}
}

func (v NullableConnectionTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


