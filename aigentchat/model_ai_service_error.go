/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.19.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the AiServiceError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiServiceError{}

// AiServiceError struct for AiServiceError
type AiServiceError struct {
	ErrorCode ApiErrorCode `json:"error_code"`
	Message string `json:"message"`
	OriginalStatusCode int32 `json:"original_status_code"`
	StatusCode int32 `json:"status_code"`
	AdditionalProperties map[string]interface{}
}

type _AiServiceError AiServiceError

// NewAiServiceError instantiates a new AiServiceError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiServiceError(errorCode ApiErrorCode, message string, originalStatusCode int32, statusCode int32) *AiServiceError {
	this := AiServiceError{}
	this.ErrorCode = errorCode
	this.Message = message
	this.OriginalStatusCode = originalStatusCode
	this.StatusCode = statusCode
	return &this
}

// NewAiServiceErrorWithDefaults instantiates a new AiServiceError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiServiceErrorWithDefaults() *AiServiceError {
	this := AiServiceError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *AiServiceError) GetErrorCode() ApiErrorCode {
	if o == nil {
		var ret ApiErrorCode
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *AiServiceError) GetErrorCodeOk() (*ApiErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *AiServiceError) SetErrorCode(v ApiErrorCode) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *AiServiceError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiServiceError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiServiceError) SetMessage(v string) {
	o.Message = v
}

// GetOriginalStatusCode returns the OriginalStatusCode field value
func (o *AiServiceError) GetOriginalStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OriginalStatusCode
}

// GetOriginalStatusCodeOk returns a tuple with the OriginalStatusCode field value
// and a boolean to check if the value has been set.
func (o *AiServiceError) GetOriginalStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalStatusCode, true
}

// SetOriginalStatusCode sets field value
func (o *AiServiceError) SetOriginalStatusCode(v int32) {
	o.OriginalStatusCode = v
}

// GetStatusCode returns the StatusCode field value
func (o *AiServiceError) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *AiServiceError) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *AiServiceError) SetStatusCode(v int32) {
	o.StatusCode = v
}

func (o AiServiceError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiServiceError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["message"] = o.Message
	toSerialize["original_status_code"] = o.OriginalStatusCode
	toSerialize["status_code"] = o.StatusCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiServiceError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error_code",
		"message",
		"original_status_code",
		"status_code",
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

	varAiServiceError := _AiServiceError{}

	err = json.Unmarshal(data, &varAiServiceError)

	if err != nil {
		return err
	}

	*o = AiServiceError(varAiServiceError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error_code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "original_status_code")
		delete(additionalProperties, "status_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiServiceError struct {
	value *AiServiceError
	isSet bool
}

func (v NullableAiServiceError) Get() *AiServiceError {
	return v.value
}

func (v *NullableAiServiceError) Set(val *AiServiceError) {
	v.value = val
	v.isSet = true
}

func (v NullableAiServiceError) IsSet() bool {
	return v.isSet
}

func (v *NullableAiServiceError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiServiceError(val *AiServiceError) *NullableAiServiceError {
	return &NullableAiServiceError{value: val, isSet: true}
}

func (v NullableAiServiceError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiServiceError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


