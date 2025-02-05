/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.0
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the ApiError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiError{}

// ApiError struct for ApiError
type ApiError struct {
	Code *ApiErrorCode `json:"code,omitempty"`
	Error *string `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiError ApiError

// NewApiError instantiates a new ApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError() *ApiError {
	this := ApiError{}
	return &this
}

// NewApiErrorWithDefaults instantiates a new ApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorWithDefaults() *ApiError {
	this := ApiError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiError) GetCode() ApiErrorCode {
	if o == nil || IsNil(o.Code) {
		var ret ApiErrorCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetCodeOk() (*ApiErrorCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given ApiErrorCode and assigns it to the Code field.
func (o *ApiError) SetCode(v ApiErrorCode) {
	o.Code = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApiError) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApiError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ApiError) SetError(v string) {
	o.Error = &v
}

func (o ApiError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiError) UnmarshalJSON(data []byte) (err error) {
	varApiError := _ApiError{}

	err = json.Unmarshal(data, &varApiError)

	if err != nil {
		return err
	}

	*o = ApiError(varApiError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiError struct {
	value *ApiError
	isSet bool
}

func (v NullableApiError) Get() *ApiError {
	return v.value
}

func (v *NullableApiError) Set(val *ApiError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiError(val *ApiError) *NullableApiError {
	return &NullableApiError{value: val, isSet: true}
}

func (v NullableApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


