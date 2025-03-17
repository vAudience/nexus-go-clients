/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.17.2
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the MissionResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionResults{}

// MissionResults struct for MissionResults
type MissionResults struct {
	Results []Mission `json:"results,omitempty"`
	TotalResults *int64 `json:"total_results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionResults MissionResults

// NewMissionResults instantiates a new MissionResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionResults() *MissionResults {
	this := MissionResults{}
	return &this
}

// NewMissionResultsWithDefaults instantiates a new MissionResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionResultsWithDefaults() *MissionResults {
	this := MissionResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MissionResults) GetResults() []Mission {
	if o == nil || IsNil(o.Results) {
		var ret []Mission
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionResults) GetResultsOk() ([]Mission, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MissionResults) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Mission and assigns it to the Results field.
func (o *MissionResults) SetResults(v []Mission) {
	o.Results = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *MissionResults) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionResults) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *MissionResults) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *MissionResults) SetTotalResults(v int64) {
	o.TotalResults = &v
}

func (o MissionResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionResults) ToMap() (map[string]interface{}, error) {
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

func (o *MissionResults) UnmarshalJSON(data []byte) (err error) {
	varMissionResults := _MissionResults{}

	err = json.Unmarshal(data, &varMissionResults)

	if err != nil {
		return err
	}

	*o = MissionResults(varMissionResults)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		delete(additionalProperties, "total_results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionResults struct {
	value *MissionResults
	isSet bool
}

func (v NullableMissionResults) Get() *MissionResults {
	return v.value
}

func (v *NullableMissionResults) Set(val *MissionResults) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionResults) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionResults(val *MissionResults) *NullableMissionResults {
	return &NullableMissionResults{value: val, isSet: true}
}

func (v NullableMissionResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


