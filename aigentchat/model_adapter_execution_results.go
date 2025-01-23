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

// checks if the AdapterExecutionResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdapterExecutionResults{}

// AdapterExecutionResults struct for AdapterExecutionResults
type AdapterExecutionResults struct {
	AdapterName *string `json:"adapter_name,omitempty"`
	Err map[string]interface{} `json:"err,omitempty"`
	FinalState *AdapterExecutionState `json:"final_state,omitempty"`
	JobId *string `json:"job_id,omitempty"`
	ResultFiles []AdapterFileInfo `json:"result_files,omitempty"`
	ResultTexts []string `json:"result_texts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterExecutionResults AdapterExecutionResults

// NewAdapterExecutionResults instantiates a new AdapterExecutionResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterExecutionResults() *AdapterExecutionResults {
	this := AdapterExecutionResults{}
	return &this
}

// NewAdapterExecutionResultsWithDefaults instantiates a new AdapterExecutionResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterExecutionResultsWithDefaults() *AdapterExecutionResults {
	this := AdapterExecutionResults{}
	return &this
}

// GetAdapterName returns the AdapterName field value if set, zero value otherwise.
func (o *AdapterExecutionResults) GetAdapterName() string {
	if o == nil || IsNil(o.AdapterName) {
		var ret string
		return ret
	}
	return *o.AdapterName
}

// GetAdapterNameOk returns a tuple with the AdapterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExecutionResults) GetAdapterNameOk() (*string, bool) {
	if o == nil || IsNil(o.AdapterName) {
		return nil, false
	}
	return o.AdapterName, true
}

// HasAdapterName returns a boolean if a field has been set.
func (o *AdapterExecutionResults) HasAdapterName() bool {
	if o != nil && !IsNil(o.AdapterName) {
		return true
	}

	return false
}

// SetAdapterName gets a reference to the given string and assigns it to the AdapterName field.
func (o *AdapterExecutionResults) SetAdapterName(v string) {
	o.AdapterName = &v
}

// GetErr returns the Err field value if set, zero value otherwise.
func (o *AdapterExecutionResults) GetErr() map[string]interface{} {
	if o == nil || IsNil(o.Err) {
		var ret map[string]interface{}
		return ret
	}
	return o.Err
}

// GetErrOk returns a tuple with the Err field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExecutionResults) GetErrOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Err) {
		return map[string]interface{}{}, false
	}
	return o.Err, true
}

// HasErr returns a boolean if a field has been set.
func (o *AdapterExecutionResults) HasErr() bool {
	if o != nil && !IsNil(o.Err) {
		return true
	}

	return false
}

// SetErr gets a reference to the given map[string]interface{} and assigns it to the Err field.
func (o *AdapterExecutionResults) SetErr(v map[string]interface{}) {
	o.Err = v
}

// GetFinalState returns the FinalState field value if set, zero value otherwise.
func (o *AdapterExecutionResults) GetFinalState() AdapterExecutionState {
	if o == nil || IsNil(o.FinalState) {
		var ret AdapterExecutionState
		return ret
	}
	return *o.FinalState
}

// GetFinalStateOk returns a tuple with the FinalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExecutionResults) GetFinalStateOk() (*AdapterExecutionState, bool) {
	if o == nil || IsNil(o.FinalState) {
		return nil, false
	}
	return o.FinalState, true
}

// HasFinalState returns a boolean if a field has been set.
func (o *AdapterExecutionResults) HasFinalState() bool {
	if o != nil && !IsNil(o.FinalState) {
		return true
	}

	return false
}

// SetFinalState gets a reference to the given AdapterExecutionState and assigns it to the FinalState field.
func (o *AdapterExecutionResults) SetFinalState(v AdapterExecutionState) {
	o.FinalState = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *AdapterExecutionResults) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExecutionResults) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *AdapterExecutionResults) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *AdapterExecutionResults) SetJobId(v string) {
	o.JobId = &v
}

// GetResultFiles returns the ResultFiles field value if set, zero value otherwise.
func (o *AdapterExecutionResults) GetResultFiles() []AdapterFileInfo {
	if o == nil || IsNil(o.ResultFiles) {
		var ret []AdapterFileInfo
		return ret
	}
	return o.ResultFiles
}

// GetResultFilesOk returns a tuple with the ResultFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExecutionResults) GetResultFilesOk() ([]AdapterFileInfo, bool) {
	if o == nil || IsNil(o.ResultFiles) {
		return nil, false
	}
	return o.ResultFiles, true
}

// HasResultFiles returns a boolean if a field has been set.
func (o *AdapterExecutionResults) HasResultFiles() bool {
	if o != nil && !IsNil(o.ResultFiles) {
		return true
	}

	return false
}

// SetResultFiles gets a reference to the given []AdapterFileInfo and assigns it to the ResultFiles field.
func (o *AdapterExecutionResults) SetResultFiles(v []AdapterFileInfo) {
	o.ResultFiles = v
}

// GetResultTexts returns the ResultTexts field value if set, zero value otherwise.
func (o *AdapterExecutionResults) GetResultTexts() []string {
	if o == nil || IsNil(o.ResultTexts) {
		var ret []string
		return ret
	}
	return o.ResultTexts
}

// GetResultTextsOk returns a tuple with the ResultTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterExecutionResults) GetResultTextsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResultTexts) {
		return nil, false
	}
	return o.ResultTexts, true
}

// HasResultTexts returns a boolean if a field has been set.
func (o *AdapterExecutionResults) HasResultTexts() bool {
	if o != nil && !IsNil(o.ResultTexts) {
		return true
	}

	return false
}

// SetResultTexts gets a reference to the given []string and assigns it to the ResultTexts field.
func (o *AdapterExecutionResults) SetResultTexts(v []string) {
	o.ResultTexts = v
}

func (o AdapterExecutionResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdapterExecutionResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdapterName) {
		toSerialize["adapter_name"] = o.AdapterName
	}
	if !IsNil(o.Err) {
		toSerialize["err"] = o.Err
	}
	if !IsNil(o.FinalState) {
		toSerialize["final_state"] = o.FinalState
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.ResultFiles) {
		toSerialize["result_files"] = o.ResultFiles
	}
	if !IsNil(o.ResultTexts) {
		toSerialize["result_texts"] = o.ResultTexts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdapterExecutionResults) UnmarshalJSON(data []byte) (err error) {
	varAdapterExecutionResults := _AdapterExecutionResults{}

	err = json.Unmarshal(data, &varAdapterExecutionResults)

	if err != nil {
		return err
	}

	*o = AdapterExecutionResults(varAdapterExecutionResults)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "adapter_name")
		delete(additionalProperties, "err")
		delete(additionalProperties, "final_state")
		delete(additionalProperties, "job_id")
		delete(additionalProperties, "result_files")
		delete(additionalProperties, "result_texts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdapterExecutionResults struct {
	value *AdapterExecutionResults
	isSet bool
}

func (v NullableAdapterExecutionResults) Get() *AdapterExecutionResults {
	return v.value
}

func (v *NullableAdapterExecutionResults) Set(val *AdapterExecutionResults) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterExecutionResults) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterExecutionResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterExecutionResults(val *AdapterExecutionResults) *NullableAdapterExecutionResults {
	return &NullableAdapterExecutionResults{value: val, isSet: true}
}

func (v NullableAdapterExecutionResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterExecutionResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


