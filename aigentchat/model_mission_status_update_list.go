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

// checks if the MissionStatusUpdateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionStatusUpdateList{}

// MissionStatusUpdateList struct for MissionStatusUpdateList
type MissionStatusUpdateList struct {
	StatusUpdates []MissionStatusUpdate `json:"status_updates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionStatusUpdateList MissionStatusUpdateList

// NewMissionStatusUpdateList instantiates a new MissionStatusUpdateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionStatusUpdateList() *MissionStatusUpdateList {
	this := MissionStatusUpdateList{}
	return &this
}

// NewMissionStatusUpdateListWithDefaults instantiates a new MissionStatusUpdateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionStatusUpdateListWithDefaults() *MissionStatusUpdateList {
	this := MissionStatusUpdateList{}
	return &this
}

// GetStatusUpdates returns the StatusUpdates field value if set, zero value otherwise.
func (o *MissionStatusUpdateList) GetStatusUpdates() []MissionStatusUpdate {
	if o == nil || IsNil(o.StatusUpdates) {
		var ret []MissionStatusUpdate
		return ret
	}
	return o.StatusUpdates
}

// GetStatusUpdatesOk returns a tuple with the StatusUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionStatusUpdateList) GetStatusUpdatesOk() ([]MissionStatusUpdate, bool) {
	if o == nil || IsNil(o.StatusUpdates) {
		return nil, false
	}
	return o.StatusUpdates, true
}

// HasStatusUpdates returns a boolean if a field has been set.
func (o *MissionStatusUpdateList) HasStatusUpdates() bool {
	if o != nil && !IsNil(o.StatusUpdates) {
		return true
	}

	return false
}

// SetStatusUpdates gets a reference to the given []MissionStatusUpdate and assigns it to the StatusUpdates field.
func (o *MissionStatusUpdateList) SetStatusUpdates(v []MissionStatusUpdate) {
	o.StatusUpdates = v
}

func (o MissionStatusUpdateList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionStatusUpdateList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusUpdates) {
		toSerialize["status_updates"] = o.StatusUpdates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionStatusUpdateList) UnmarshalJSON(data []byte) (err error) {
	varMissionStatusUpdateList := _MissionStatusUpdateList{}

	err = json.Unmarshal(data, &varMissionStatusUpdateList)

	if err != nil {
		return err
	}

	*o = MissionStatusUpdateList(varMissionStatusUpdateList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status_updates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionStatusUpdateList struct {
	value *MissionStatusUpdateList
	isSet bool
}

func (v NullableMissionStatusUpdateList) Get() *MissionStatusUpdateList {
	return v.value
}

func (v *NullableMissionStatusUpdateList) Set(val *MissionStatusUpdateList) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionStatusUpdateList) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionStatusUpdateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionStatusUpdateList(val *MissionStatusUpdateList) *NullableMissionStatusUpdateList {
	return &NullableMissionStatusUpdateList{value: val, isSet: true}
}

func (v NullableMissionStatusUpdateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionStatusUpdateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


