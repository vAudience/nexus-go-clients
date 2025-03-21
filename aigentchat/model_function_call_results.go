/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.17.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the FunctionCallResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionCallResults{}

// FunctionCallResults struct for FunctionCallResults
type FunctionCallResults struct {
	Error *string `json:"error,omitempty"`
	FinalState *string `json:"final_state,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ResultFiles []ResultFile `json:"result_files,omitempty"`
	ResultTexts []string `json:"result_texts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FunctionCallResults FunctionCallResults

// NewFunctionCallResults instantiates a new FunctionCallResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionCallResults() *FunctionCallResults {
	this := FunctionCallResults{}
	return &this
}

// NewFunctionCallResultsWithDefaults instantiates a new FunctionCallResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionCallResultsWithDefaults() *FunctionCallResults {
	this := FunctionCallResults{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FunctionCallResults) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallResults) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FunctionCallResults) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *FunctionCallResults) SetError(v string) {
	o.Error = &v
}

// GetFinalState returns the FinalState field value if set, zero value otherwise.
func (o *FunctionCallResults) GetFinalState() string {
	if o == nil || IsNil(o.FinalState) {
		var ret string
		return ret
	}
	return *o.FinalState
}

// GetFinalStateOk returns a tuple with the FinalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallResults) GetFinalStateOk() (*string, bool) {
	if o == nil || IsNil(o.FinalState) {
		return nil, false
	}
	return o.FinalState, true
}

// HasFinalState returns a boolean if a field has been set.
func (o *FunctionCallResults) HasFinalState() bool {
	if o != nil && !IsNil(o.FinalState) {
		return true
	}

	return false
}

// SetFinalState gets a reference to the given string and assigns it to the FinalState field.
func (o *FunctionCallResults) SetFinalState(v string) {
	o.FinalState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FunctionCallResults) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallResults) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FunctionCallResults) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FunctionCallResults) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FunctionCallResults) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallResults) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FunctionCallResults) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FunctionCallResults) SetName(v string) {
	o.Name = &v
}

// GetResultFiles returns the ResultFiles field value if set, zero value otherwise.
func (o *FunctionCallResults) GetResultFiles() []ResultFile {
	if o == nil || IsNil(o.ResultFiles) {
		var ret []ResultFile
		return ret
	}
	return o.ResultFiles
}

// GetResultFilesOk returns a tuple with the ResultFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallResults) GetResultFilesOk() ([]ResultFile, bool) {
	if o == nil || IsNil(o.ResultFiles) {
		return nil, false
	}
	return o.ResultFiles, true
}

// HasResultFiles returns a boolean if a field has been set.
func (o *FunctionCallResults) HasResultFiles() bool {
	if o != nil && !IsNil(o.ResultFiles) {
		return true
	}

	return false
}

// SetResultFiles gets a reference to the given []ResultFile and assigns it to the ResultFiles field.
func (o *FunctionCallResults) SetResultFiles(v []ResultFile) {
	o.ResultFiles = v
}

// GetResultTexts returns the ResultTexts field value if set, zero value otherwise.
func (o *FunctionCallResults) GetResultTexts() []string {
	if o == nil || IsNil(o.ResultTexts) {
		var ret []string
		return ret
	}
	return o.ResultTexts
}

// GetResultTextsOk returns a tuple with the ResultTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCallResults) GetResultTextsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResultTexts) {
		return nil, false
	}
	return o.ResultTexts, true
}

// HasResultTexts returns a boolean if a field has been set.
func (o *FunctionCallResults) HasResultTexts() bool {
	if o != nil && !IsNil(o.ResultTexts) {
		return true
	}

	return false
}

// SetResultTexts gets a reference to the given []string and assigns it to the ResultTexts field.
func (o *FunctionCallResults) SetResultTexts(v []string) {
	o.ResultTexts = v
}

func (o FunctionCallResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionCallResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.FinalState) {
		toSerialize["final_state"] = o.FinalState
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

func (o *FunctionCallResults) UnmarshalJSON(data []byte) (err error) {
	varFunctionCallResults := _FunctionCallResults{}

	err = json.Unmarshal(data, &varFunctionCallResults)

	if err != nil {
		return err
	}

	*o = FunctionCallResults(varFunctionCallResults)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "final_state")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "result_files")
		delete(additionalProperties, "result_texts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFunctionCallResults struct {
	value *FunctionCallResults
	isSet bool
}

func (v NullableFunctionCallResults) Get() *FunctionCallResults {
	return v.value
}

func (v *NullableFunctionCallResults) Set(val *FunctionCallResults) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionCallResults) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionCallResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionCallResults(val *FunctionCallResults) *NullableFunctionCallResults {
	return &NullableFunctionCallResults{value: val, isSet: true}
}

func (v NullableFunctionCallResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionCallResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


