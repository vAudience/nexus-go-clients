/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.18.0
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the Mission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mission{}

// Mission struct for Mission
type Mission struct {
	CompletedAt *int64 `json:"completed_at,omitempty"`
	CompletionReason *MissionCompletionReason `json:"completion_reason,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatorName *string `json:"creator_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Instructions *MissionInstructions `json:"instructions,omitempty"`
	MissionExecutorId *string `json:"mission_executor_id,omitempty"`
	OwnerId *string `json:"owner_id,omitempty"`
	OwnerOrganizationId *string `json:"owner_organization_id,omitempty"`
	StatusUpdates *MissionStatusUpdateList `json:"status_updates,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Mission Mission

// NewMission instantiates a new Mission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMission() *Mission {
	this := Mission{}
	return &this
}

// NewMissionWithDefaults instantiates a new Mission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionWithDefaults() *Mission {
	this := Mission{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *Mission) GetCompletedAt() int64 {
	if o == nil || IsNil(o.CompletedAt) {
		var ret int64
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetCompletedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *Mission) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given int64 and assigns it to the CompletedAt field.
func (o *Mission) SetCompletedAt(v int64) {
	o.CompletedAt = &v
}

// GetCompletionReason returns the CompletionReason field value if set, zero value otherwise.
func (o *Mission) GetCompletionReason() MissionCompletionReason {
	if o == nil || IsNil(o.CompletionReason) {
		var ret MissionCompletionReason
		return ret
	}
	return *o.CompletionReason
}

// GetCompletionReasonOk returns a tuple with the CompletionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetCompletionReasonOk() (*MissionCompletionReason, bool) {
	if o == nil || IsNil(o.CompletionReason) {
		return nil, false
	}
	return o.CompletionReason, true
}

// HasCompletionReason returns a boolean if a field has been set.
func (o *Mission) HasCompletionReason() bool {
	if o != nil && !IsNil(o.CompletionReason) {
		return true
	}

	return false
}

// SetCompletionReason gets a reference to the given MissionCompletionReason and assigns it to the CompletionReason field.
func (o *Mission) SetCompletionReason(v MissionCompletionReason) {
	o.CompletionReason = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Mission) GetContent() map[string]interface{} {
	if o == nil || IsNil(o.Content) {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetContentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Content) {
		return map[string]interface{}{}, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Mission) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given map[string]interface{} and assigns it to the Content field.
func (o *Mission) SetContent(v map[string]interface{}) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Mission) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Mission) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Mission) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Mission) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Mission) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Mission) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatorName returns the CreatorName field value if set, zero value otherwise.
func (o *Mission) GetCreatorName() string {
	if o == nil || IsNil(o.CreatorName) {
		var ret string
		return ret
	}
	return *o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetCreatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorName) {
		return nil, false
	}
	return o.CreatorName, true
}

// HasCreatorName returns a boolean if a field has been set.
func (o *Mission) HasCreatorName() bool {
	if o != nil && !IsNil(o.CreatorName) {
		return true
	}

	return false
}

// SetCreatorName gets a reference to the given string and assigns it to the CreatorName field.
func (o *Mission) SetCreatorName(v string) {
	o.CreatorName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Mission) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Mission) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Mission) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Mission) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Mission) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Mission) SetId(v string) {
	o.Id = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *Mission) GetInstructions() MissionInstructions {
	if o == nil || IsNil(o.Instructions) {
		var ret MissionInstructions
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetInstructionsOk() (*MissionInstructions, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *Mission) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given MissionInstructions and assigns it to the Instructions field.
func (o *Mission) SetInstructions(v MissionInstructions) {
	o.Instructions = &v
}

// GetMissionExecutorId returns the MissionExecutorId field value if set, zero value otherwise.
func (o *Mission) GetMissionExecutorId() string {
	if o == nil || IsNil(o.MissionExecutorId) {
		var ret string
		return ret
	}
	return *o.MissionExecutorId
}

// GetMissionExecutorIdOk returns a tuple with the MissionExecutorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetMissionExecutorIdOk() (*string, bool) {
	if o == nil || IsNil(o.MissionExecutorId) {
		return nil, false
	}
	return o.MissionExecutorId, true
}

// HasMissionExecutorId returns a boolean if a field has been set.
func (o *Mission) HasMissionExecutorId() bool {
	if o != nil && !IsNil(o.MissionExecutorId) {
		return true
	}

	return false
}

// SetMissionExecutorId gets a reference to the given string and assigns it to the MissionExecutorId field.
func (o *Mission) SetMissionExecutorId(v string) {
	o.MissionExecutorId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *Mission) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *Mission) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *Mission) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerOrganizationId returns the OwnerOrganizationId field value if set, zero value otherwise.
func (o *Mission) GetOwnerOrganizationId() string {
	if o == nil || IsNil(o.OwnerOrganizationId) {
		var ret string
		return ret
	}
	return *o.OwnerOrganizationId
}

// GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetOwnerOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerOrganizationId) {
		return nil, false
	}
	return o.OwnerOrganizationId, true
}

// HasOwnerOrganizationId returns a boolean if a field has been set.
func (o *Mission) HasOwnerOrganizationId() bool {
	if o != nil && !IsNil(o.OwnerOrganizationId) {
		return true
	}

	return false
}

// SetOwnerOrganizationId gets a reference to the given string and assigns it to the OwnerOrganizationId field.
func (o *Mission) SetOwnerOrganizationId(v string) {
	o.OwnerOrganizationId = &v
}

// GetStatusUpdates returns the StatusUpdates field value if set, zero value otherwise.
func (o *Mission) GetStatusUpdates() MissionStatusUpdateList {
	if o == nil || IsNil(o.StatusUpdates) {
		var ret MissionStatusUpdateList
		return ret
	}
	return *o.StatusUpdates
}

// GetStatusUpdatesOk returns a tuple with the StatusUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetStatusUpdatesOk() (*MissionStatusUpdateList, bool) {
	if o == nil || IsNil(o.StatusUpdates) {
		return nil, false
	}
	return o.StatusUpdates, true
}

// HasStatusUpdates returns a boolean if a field has been set.
func (o *Mission) HasStatusUpdates() bool {
	if o != nil && !IsNil(o.StatusUpdates) {
		return true
	}

	return false
}

// SetStatusUpdates gets a reference to the given MissionStatusUpdateList and assigns it to the StatusUpdates field.
func (o *Mission) SetStatusUpdates(v MissionStatusUpdateList) {
	o.StatusUpdates = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Mission) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mission) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Mission) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Mission) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o Mission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompletedAt) {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if !IsNil(o.CompletionReason) {
		toSerialize["completion_reason"] = o.CompletionReason
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.CreatorName) {
		toSerialize["creator_name"] = o.CreatorName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if !IsNil(o.MissionExecutorId) {
		toSerialize["mission_executor_id"] = o.MissionExecutorId
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.OwnerOrganizationId) {
		toSerialize["owner_organization_id"] = o.OwnerOrganizationId
	}
	if !IsNil(o.StatusUpdates) {
		toSerialize["status_updates"] = o.StatusUpdates
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Mission) UnmarshalJSON(data []byte) (err error) {
	varMission := _Mission{}

	err = json.Unmarshal(data, &varMission)

	if err != nil {
		return err
	}

	*o = Mission(varMission)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "completion_reason")
		delete(additionalProperties, "content")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "creator_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "instructions")
		delete(additionalProperties, "mission_executor_id")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "owner_organization_id")
		delete(additionalProperties, "status_updates")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMission struct {
	value *Mission
	isSet bool
}

func (v NullableMission) Get() *Mission {
	return v.value
}

func (v *NullableMission) Set(val *Mission) {
	v.value = val
	v.isSet = true
}

func (v NullableMission) IsSet() bool {
	return v.isSet
}

func (v *NullableMission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMission(val *Mission) *NullableMission {
	return &NullableMission{value: val, isSet: true}
}

func (v NullableMission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


