/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
)

// checks if the ErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponse{}

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	Code *ErrorCode `json:"code,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
	Message *string `json:"message,omitempty"`
	Status *int32 `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorResponse ErrorResponse

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorResponse) GetCode() ErrorCode {
	if o == nil || IsNil(o.Code) {
		var ret ErrorCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetCodeOk() (*ErrorCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given ErrorCode and assigns it to the Code field.
func (o *ErrorResponse) SetCode(v ErrorCode) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ErrorResponse) GetDetails() []ErrorDetail {
	if o == nil || IsNil(o.Details) {
		var ret []ErrorDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetDetailsOk() ([]ErrorDetail, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ErrorResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ErrorDetail and assigns it to the Details field.
func (o *ErrorResponse) SetDetails(v []ErrorDetail) {
	o.Details = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ErrorResponse) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ErrorResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ErrorResponse) SetStatus(v int32) {
	o.Status = &v
}

func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorResponse) UnmarshalJSON(data []byte) (err error) {
	varErrorResponse := _ErrorResponse{}

	err = json.Unmarshal(data, &varErrorResponse)

	if err != nil {
		return err
	}

	*o = ErrorResponse(varErrorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "details")
		delete(additionalProperties, "message")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


