/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.9
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the MissionStatusUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionStatusUpdate{}

// MissionStatusUpdate struct for MissionStatusUpdate
type MissionStatusUpdate struct {
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	MissionId *string `json:"mission_id,omitempty"`
	Status *MissionStatus `json:"status,omitempty"`
	Timestamp *int32 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionStatusUpdate MissionStatusUpdate

// NewMissionStatusUpdate instantiates a new MissionStatusUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionStatusUpdate() *MissionStatusUpdate {
	this := MissionStatusUpdate{}
	return &this
}

// NewMissionStatusUpdateWithDefaults instantiates a new MissionStatusUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionStatusUpdateWithDefaults() *MissionStatusUpdate {
	this := MissionStatusUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MissionStatusUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionStatusUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MissionStatusUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MissionStatusUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MissionStatusUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionStatusUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MissionStatusUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MissionStatusUpdate) SetId(v string) {
	o.Id = &v
}

// GetMissionId returns the MissionId field value if set, zero value otherwise.
func (o *MissionStatusUpdate) GetMissionId() string {
	if o == nil || IsNil(o.MissionId) {
		var ret string
		return ret
	}
	return *o.MissionId
}

// GetMissionIdOk returns a tuple with the MissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionStatusUpdate) GetMissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.MissionId) {
		return nil, false
	}
	return o.MissionId, true
}

// HasMissionId returns a boolean if a field has been set.
func (o *MissionStatusUpdate) HasMissionId() bool {
	if o != nil && !IsNil(o.MissionId) {
		return true
	}

	return false
}

// SetMissionId gets a reference to the given string and assigns it to the MissionId field.
func (o *MissionStatusUpdate) SetMissionId(v string) {
	o.MissionId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MissionStatusUpdate) GetStatus() MissionStatus {
	if o == nil || IsNil(o.Status) {
		var ret MissionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionStatusUpdate) GetStatusOk() (*MissionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MissionStatusUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MissionStatus and assigns it to the Status field.
func (o *MissionStatusUpdate) SetStatus(v MissionStatus) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MissionStatusUpdate) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissionStatusUpdate) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MissionStatusUpdate) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *MissionStatusUpdate) SetTimestamp(v int32) {
	o.Timestamp = &v
}

func (o MissionStatusUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionStatusUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MissionId) {
		toSerialize["mission_id"] = o.MissionId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionStatusUpdate) UnmarshalJSON(data []byte) (err error) {
	varMissionStatusUpdate := _MissionStatusUpdate{}

	err = json.Unmarshal(data, &varMissionStatusUpdate)

	if err != nil {
		return err
	}

	*o = MissionStatusUpdate(varMissionStatusUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "mission_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionStatusUpdate struct {
	value *MissionStatusUpdate
	isSet bool
}

func (v NullableMissionStatusUpdate) Get() *MissionStatusUpdate {
	return v.value
}

func (v *NullableMissionStatusUpdate) Set(val *MissionStatusUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionStatusUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionStatusUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionStatusUpdate(val *MissionStatusUpdate) *NullableMissionStatusUpdate {
	return &NullableMissionStatusUpdate{value: val, isSet: true}
}

func (v NullableMissionStatusUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionStatusUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


