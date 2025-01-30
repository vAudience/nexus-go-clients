/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.14.1
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the AgentWriteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentWriteDto{}

// AgentWriteDto struct for AgentWriteDto
type AgentWriteDto struct {
	AssignedTools []string `json:"assigned_tools,omitempty"`
	AttachedFileIds []string `json:"attached_file_ids,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Description *string `json:"description,omitempty"`
	InitialUserMessages []string `json:"initial_user_messages,omitempty"`
	InternalId *string `json:"internal_id,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	ModelHostLocation *HostingLocation `json:"model_host_location,omitempty"`
	ModelId *string `json:"model_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	SystemMessages []string `json:"system_messages,omitempty"`
	UseTools *bool `json:"use_tools,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentWriteDto AgentWriteDto

// NewAgentWriteDto instantiates a new AgentWriteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentWriteDto() *AgentWriteDto {
	this := AgentWriteDto{}
	return &this
}

// NewAgentWriteDtoWithDefaults instantiates a new AgentWriteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWriteDtoWithDefaults() *AgentWriteDto {
	this := AgentWriteDto{}
	return &this
}

// GetAssignedTools returns the AssignedTools field value if set, zero value otherwise.
func (o *AgentWriteDto) GetAssignedTools() []string {
	if o == nil || IsNil(o.AssignedTools) {
		var ret []string
		return ret
	}
	return o.AssignedTools
}

// GetAssignedToolsOk returns a tuple with the AssignedTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetAssignedToolsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignedTools) {
		return nil, false
	}
	return o.AssignedTools, true
}

// HasAssignedTools returns a boolean if a field has been set.
func (o *AgentWriteDto) HasAssignedTools() bool {
	if o != nil && !IsNil(o.AssignedTools) {
		return true
	}

	return false
}

// SetAssignedTools gets a reference to the given []string and assigns it to the AssignedTools field.
func (o *AgentWriteDto) SetAssignedTools(v []string) {
	o.AssignedTools = v
}

// GetAttachedFileIds returns the AttachedFileIds field value if set, zero value otherwise.
func (o *AgentWriteDto) GetAttachedFileIds() []string {
	if o == nil || IsNil(o.AttachedFileIds) {
		var ret []string
		return ret
	}
	return o.AttachedFileIds
}

// GetAttachedFileIdsOk returns a tuple with the AttachedFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetAttachedFileIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AttachedFileIds) {
		return nil, false
	}
	return o.AttachedFileIds, true
}

// HasAttachedFileIds returns a boolean if a field has been set.
func (o *AgentWriteDto) HasAttachedFileIds() bool {
	if o != nil && !IsNil(o.AttachedFileIds) {
		return true
	}

	return false
}

// SetAttachedFileIds gets a reference to the given []string and assigns it to the AttachedFileIds field.
func (o *AgentWriteDto) SetAttachedFileIds(v []string) {
	o.AttachedFileIds = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *AgentWriteDto) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *AgentWriteDto) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *AgentWriteDto) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AgentWriteDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AgentWriteDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AgentWriteDto) SetDescription(v string) {
	o.Description = &v
}

// GetInitialUserMessages returns the InitialUserMessages field value if set, zero value otherwise.
func (o *AgentWriteDto) GetInitialUserMessages() []string {
	if o == nil || IsNil(o.InitialUserMessages) {
		var ret []string
		return ret
	}
	return o.InitialUserMessages
}

// GetInitialUserMessagesOk returns a tuple with the InitialUserMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetInitialUserMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.InitialUserMessages) {
		return nil, false
	}
	return o.InitialUserMessages, true
}

// HasInitialUserMessages returns a boolean if a field has been set.
func (o *AgentWriteDto) HasInitialUserMessages() bool {
	if o != nil && !IsNil(o.InitialUserMessages) {
		return true
	}

	return false
}

// SetInitialUserMessages gets a reference to the given []string and assigns it to the InitialUserMessages field.
func (o *AgentWriteDto) SetInitialUserMessages(v []string) {
	o.InitialUserMessages = v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *AgentWriteDto) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// HasInternalId returns a boolean if a field has been set.
func (o *AgentWriteDto) HasInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *AgentWriteDto) SetInternalId(v string) {
	o.InternalId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *AgentWriteDto) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *AgentWriteDto) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *AgentWriteDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *AgentWriteDto) GetMetaData() map[string]interface{} {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return map[string]interface{}{}, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *AgentWriteDto) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *AgentWriteDto) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetModelHostLocation returns the ModelHostLocation field value if set, zero value otherwise.
func (o *AgentWriteDto) GetModelHostLocation() HostingLocation {
	if o == nil || IsNil(o.ModelHostLocation) {
		var ret HostingLocation
		return ret
	}
	return *o.ModelHostLocation
}

// GetModelHostLocationOk returns a tuple with the ModelHostLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetModelHostLocationOk() (*HostingLocation, bool) {
	if o == nil || IsNil(o.ModelHostLocation) {
		return nil, false
	}
	return o.ModelHostLocation, true
}

// HasModelHostLocation returns a boolean if a field has been set.
func (o *AgentWriteDto) HasModelHostLocation() bool {
	if o != nil && !IsNil(o.ModelHostLocation) {
		return true
	}

	return false
}

// SetModelHostLocation gets a reference to the given HostingLocation and assigns it to the ModelHostLocation field.
func (o *AgentWriteDto) SetModelHostLocation(v HostingLocation) {
	o.ModelHostLocation = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *AgentWriteDto) GetModelId() string {
	if o == nil || IsNil(o.ModelId) {
		var ret string
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *AgentWriteDto) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given string and assigns it to the ModelId field.
func (o *AgentWriteDto) SetModelId(v string) {
	o.ModelId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentWriteDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentWriteDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgentWriteDto) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *AgentWriteDto) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AgentWriteDto) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *AgentWriteDto) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetSystemMessages returns the SystemMessages field value if set, zero value otherwise.
func (o *AgentWriteDto) GetSystemMessages() []string {
	if o == nil || IsNil(o.SystemMessages) {
		var ret []string
		return ret
	}
	return o.SystemMessages
}

// GetSystemMessagesOk returns a tuple with the SystemMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetSystemMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.SystemMessages) {
		return nil, false
	}
	return o.SystemMessages, true
}

// HasSystemMessages returns a boolean if a field has been set.
func (o *AgentWriteDto) HasSystemMessages() bool {
	if o != nil && !IsNil(o.SystemMessages) {
		return true
	}

	return false
}

// SetSystemMessages gets a reference to the given []string and assigns it to the SystemMessages field.
func (o *AgentWriteDto) SetSystemMessages(v []string) {
	o.SystemMessages = v
}

// GetUseTools returns the UseTools field value if set, zero value otherwise.
func (o *AgentWriteDto) GetUseTools() bool {
	if o == nil || IsNil(o.UseTools) {
		var ret bool
		return ret
	}
	return *o.UseTools
}

// GetUseToolsOk returns a tuple with the UseTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWriteDto) GetUseToolsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTools) {
		return nil, false
	}
	return o.UseTools, true
}

// HasUseTools returns a boolean if a field has been set.
func (o *AgentWriteDto) HasUseTools() bool {
	if o != nil && !IsNil(o.UseTools) {
		return true
	}

	return false
}

// SetUseTools gets a reference to the given bool and assigns it to the UseTools field.
func (o *AgentWriteDto) SetUseTools(v bool) {
	o.UseTools = &v
}

func (o AgentWriteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentWriteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignedTools) {
		toSerialize["assigned_tools"] = o.AssignedTools
	}
	if !IsNil(o.AttachedFileIds) {
		toSerialize["attached_file_ids"] = o.AttachedFileIds
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InitialUserMessages) {
		toSerialize["initial_user_messages"] = o.InitialUserMessages
	}
	if !IsNil(o.InternalId) {
		toSerialize["internal_id"] = o.InternalId
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.MetaData) {
		toSerialize["meta_data"] = o.MetaData
	}
	if !IsNil(o.ModelHostLocation) {
		toSerialize["model_host_location"] = o.ModelHostLocation
	}
	if !IsNil(o.ModelId) {
		toSerialize["model_id"] = o.ModelId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.SystemMessages) {
		toSerialize["system_messages"] = o.SystemMessages
	}
	if !IsNil(o.UseTools) {
		toSerialize["use_tools"] = o.UseTools
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentWriteDto) UnmarshalJSON(data []byte) (err error) {
	varAgentWriteDto := _AgentWriteDto{}

	err = json.Unmarshal(data, &varAgentWriteDto)

	if err != nil {
		return err
	}

	*o = AgentWriteDto(varAgentWriteDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assigned_tools")
		delete(additionalProperties, "attached_file_ids")
		delete(additionalProperties, "avatar_url")
		delete(additionalProperties, "description")
		delete(additionalProperties, "initial_user_messages")
		delete(additionalProperties, "internal_id")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "meta_data")
		delete(additionalProperties, "model_host_location")
		delete(additionalProperties, "model_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "system_messages")
		delete(additionalProperties, "use_tools")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentWriteDto struct {
	value *AgentWriteDto
	isSet bool
}

func (v NullableAgentWriteDto) Get() *AgentWriteDto {
	return v.value
}

func (v *NullableAgentWriteDto) Set(val *AgentWriteDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentWriteDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentWriteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentWriteDto(val *AgentWriteDto) *NullableAgentWriteDto {
	return &NullableAgentWriteDto{value: val, isSet: true}
}

func (v NullableAgentWriteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentWriteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


