/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.14.0
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the ExecutionLogResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionLogResults{}

// ExecutionLogResults struct for ExecutionLogResults
type ExecutionLogResults struct {
	Results []ExecutionLog `json:"results,omitempty"`
	TotalResults *int64 `json:"total_results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionLogResults ExecutionLogResults

// NewExecutionLogResults instantiates a new ExecutionLogResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionLogResults() *ExecutionLogResults {
	this := ExecutionLogResults{}
	return &this
}

// NewExecutionLogResultsWithDefaults instantiates a new ExecutionLogResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionLogResultsWithDefaults() *ExecutionLogResults {
	this := ExecutionLogResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ExecutionLogResults) GetResults() []ExecutionLog {
	if o == nil || IsNil(o.Results) {
		var ret []ExecutionLog
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLogResults) GetResultsOk() ([]ExecutionLog, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ExecutionLogResults) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ExecutionLog and assigns it to the Results field.
func (o *ExecutionLogResults) SetResults(v []ExecutionLog) {
	o.Results = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ExecutionLogResults) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLogResults) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ExecutionLogResults) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *ExecutionLogResults) SetTotalResults(v int64) {
	o.TotalResults = &v
}

func (o ExecutionLogResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionLogResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.TotalResults) {
		toSerialize["total_results"] = o.TotalResults
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionLogResults) UnmarshalJSON(data []byte) (err error) {
	varExecutionLogResults := _ExecutionLogResults{}

	err = json.Unmarshal(data, &varExecutionLogResults)

	if err != nil {
		return err
	}

	*o = ExecutionLogResults(varExecutionLogResults)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		delete(additionalProperties, "total_results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionLogResults struct {
	value *ExecutionLogResults
	isSet bool
}

func (v NullableExecutionLogResults) Get() *ExecutionLogResults {
	return v.value
}

func (v *NullableExecutionLogResults) Set(val *ExecutionLogResults) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionLogResults) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionLogResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionLogResults(val *ExecutionLogResults) *NullableExecutionLogResults {
	return &NullableExecutionLogResults{value: val, isSet: true}
}

func (v NullableExecutionLogResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionLogResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


