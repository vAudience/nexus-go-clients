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
	"fmt"
)

// TokenDirection the model 'TokenDirection'
type TokenDirection string

// List of TokenDirection
const (
	InputToken TokenDirection = "input"
	OutputToken TokenDirection = "output"
)

// All allowed values of TokenDirection enum
var AllowedTokenDirectionEnumValues = []TokenDirection{
	"input",
	"output",
}

func (v *TokenDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TokenDirection(value)
	for _, existing := range AllowedTokenDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TokenDirection", value)
}

// NewTokenDirectionFromValue returns a pointer to a valid TokenDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTokenDirectionFromValue(v string) (*TokenDirection, error) {
	ev := TokenDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TokenDirection: valid values are %v", v, AllowedTokenDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TokenDirection) IsValid() bool {
	for _, existing := range AllowedTokenDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TokenDirection value
func (v TokenDirection) Ptr() *TokenDirection {
	return &v
}

type NullableTokenDirection struct {
	value *TokenDirection
	isSet bool
}

func (v NullableTokenDirection) Get() *TokenDirection {
	return v.value
}

func (v *NullableTokenDirection) Set(val *TokenDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenDirection(val *TokenDirection) *NullableTokenDirection {
	return &NullableTokenDirection{value: val, isSet: true}
}

func (v NullableTokenDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

