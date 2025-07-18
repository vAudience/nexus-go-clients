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

// checks if the Agent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Agent{}

// Agent struct for Agent
type Agent struct {
	Abilities []Ability `json:"abilities,omitempty"`
	AbilitiesV2 []AbilityV2 `json:"abilities_v2,omitempty"`
	AssignedTools []string `json:"assigned_tools,omitempty"`
	AttachedFileIds []string `json:"attached_file_ids,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	I18n *map[string]AgentI18n `json:"i18n,omitempty"`
	Id string `json:"id"`
	InitialUserMessages []string `json:"initial_user_messages,omitempty"`
	InternalId *string `json:"internal_id,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	ModelCategory *string `json:"model_category,omitempty"`
	ModelHostLocation *HostingLocation `json:"model_host_location,omitempty"`
	ModelId string `json:"model_id"`
	ModelReleaseDate *int64 `json:"model_release_date,omitempty"`
	Name string `json:"name"`
	OwnerId string `json:"owner_id"`
	OwnerOrganizationId string `json:"owner_organization_id"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	SystemMessages []string `json:"system_messages,omitempty"`
	TeamIds []string `json:"team_ids,omitempty"`
	Type *AgentType `json:"type,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	UseTools *bool `json:"use_tools,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Agent Agent

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent(id string, modelId string, name string, ownerId string, ownerOrganizationId string) *Agent {
	this := Agent{}
	this.Id = id
	this.ModelId = modelId
	this.Name = name
	this.OwnerId = ownerId
	this.OwnerOrganizationId = ownerOrganizationId
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	return &this
}

// GetAbilities returns the Abilities field value if set, zero value otherwise.
func (o *Agent) GetAbilities() []Ability {
	if o == nil || IsNil(o.Abilities) {
		var ret []Ability
		return ret
	}
	return o.Abilities
}

// GetAbilitiesOk returns a tuple with the Abilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetAbilitiesOk() ([]Ability, bool) {
	if o == nil || IsNil(o.Abilities) {
		return nil, false
	}
	return o.Abilities, true
}

// HasAbilities returns a boolean if a field has been set.
func (o *Agent) HasAbilities() bool {
	if o != nil && !IsNil(o.Abilities) {
		return true
	}

	return false
}

// SetAbilities gets a reference to the given []Ability and assigns it to the Abilities field.
func (o *Agent) SetAbilities(v []Ability) {
	o.Abilities = v
}

// GetAbilitiesV2 returns the AbilitiesV2 field value if set, zero value otherwise.
func (o *Agent) GetAbilitiesV2() []AbilityV2 {
	if o == nil || IsNil(o.AbilitiesV2) {
		var ret []AbilityV2
		return ret
	}
	return o.AbilitiesV2
}

// GetAbilitiesV2Ok returns a tuple with the AbilitiesV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetAbilitiesV2Ok() ([]AbilityV2, bool) {
	if o == nil || IsNil(o.AbilitiesV2) {
		return nil, false
	}
	return o.AbilitiesV2, true
}

// HasAbilitiesV2 returns a boolean if a field has been set.
func (o *Agent) HasAbilitiesV2() bool {
	if o != nil && !IsNil(o.AbilitiesV2) {
		return true
	}

	return false
}

// SetAbilitiesV2 gets a reference to the given []AbilityV2 and assigns it to the AbilitiesV2 field.
func (o *Agent) SetAbilitiesV2(v []AbilityV2) {
	o.AbilitiesV2 = v
}

// GetAssignedTools returns the AssignedTools field value if set, zero value otherwise.
func (o *Agent) GetAssignedTools() []string {
	if o == nil || IsNil(o.AssignedTools) {
		var ret []string
		return ret
	}
	return o.AssignedTools
}

// GetAssignedToolsOk returns a tuple with the AssignedTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetAssignedToolsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignedTools) {
		return nil, false
	}
	return o.AssignedTools, true
}

// HasAssignedTools returns a boolean if a field has been set.
func (o *Agent) HasAssignedTools() bool {
	if o != nil && !IsNil(o.AssignedTools) {
		return true
	}

	return false
}

// SetAssignedTools gets a reference to the given []string and assigns it to the AssignedTools field.
func (o *Agent) SetAssignedTools(v []string) {
	o.AssignedTools = v
}

// GetAttachedFileIds returns the AttachedFileIds field value if set, zero value otherwise.
func (o *Agent) GetAttachedFileIds() []string {
	if o == nil || IsNil(o.AttachedFileIds) {
		var ret []string
		return ret
	}
	return o.AttachedFileIds
}

// GetAttachedFileIdsOk returns a tuple with the AttachedFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetAttachedFileIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AttachedFileIds) {
		return nil, false
	}
	return o.AttachedFileIds, true
}

// HasAttachedFileIds returns a boolean if a field has been set.
func (o *Agent) HasAttachedFileIds() bool {
	if o != nil && !IsNil(o.AttachedFileIds) {
		return true
	}

	return false
}

// SetAttachedFileIds gets a reference to the given []string and assigns it to the AttachedFileIds field.
func (o *Agent) SetAttachedFileIds(v []string) {
	o.AttachedFileIds = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *Agent) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *Agent) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *Agent) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Agent) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Agent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Agent) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Agent) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Agent) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Agent) SetDescription(v string) {
	o.Description = &v
}

// GetI18n returns the I18n field value if set, zero value otherwise.
func (o *Agent) GetI18n() map[string]AgentI18n {
	if o == nil || IsNil(o.I18n) {
		var ret map[string]AgentI18n
		return ret
	}
	return *o.I18n
}

// GetI18nOk returns a tuple with the I18n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetI18nOk() (*map[string]AgentI18n, bool) {
	if o == nil || IsNil(o.I18n) {
		return nil, false
	}
	return o.I18n, true
}

// HasI18n returns a boolean if a field has been set.
func (o *Agent) HasI18n() bool {
	if o != nil && !IsNil(o.I18n) {
		return true
	}

	return false
}

// SetI18n gets a reference to the given map[string]AgentI18n and assigns it to the I18n field.
func (o *Agent) SetI18n(v map[string]AgentI18n) {
	o.I18n = &v
}

// GetId returns the Id field value
func (o *Agent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Agent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Agent) SetId(v string) {
	o.Id = v
}

// GetInitialUserMessages returns the InitialUserMessages field value if set, zero value otherwise.
func (o *Agent) GetInitialUserMessages() []string {
	if o == nil || IsNil(o.InitialUserMessages) {
		var ret []string
		return ret
	}
	return o.InitialUserMessages
}

// GetInitialUserMessagesOk returns a tuple with the InitialUserMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetInitialUserMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.InitialUserMessages) {
		return nil, false
	}
	return o.InitialUserMessages, true
}

// HasInitialUserMessages returns a boolean if a field has been set.
func (o *Agent) HasInitialUserMessages() bool {
	if o != nil && !IsNil(o.InitialUserMessages) {
		return true
	}

	return false
}

// SetInitialUserMessages gets a reference to the given []string and assigns it to the InitialUserMessages field.
func (o *Agent) SetInitialUserMessages(v []string) {
	o.InitialUserMessages = v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *Agent) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// HasInternalId returns a boolean if a field has been set.
func (o *Agent) HasInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *Agent) SetInternalId(v string) {
	o.InternalId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *Agent) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *Agent) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *Agent) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *Agent) GetMetaData() map[string]interface{} {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return map[string]interface{}{}, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *Agent) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *Agent) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetModelCategory returns the ModelCategory field value if set, zero value otherwise.
func (o *Agent) GetModelCategory() string {
	if o == nil || IsNil(o.ModelCategory) {
		var ret string
		return ret
	}
	return *o.ModelCategory
}

// GetModelCategoryOk returns a tuple with the ModelCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetModelCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ModelCategory) {
		return nil, false
	}
	return o.ModelCategory, true
}

// HasModelCategory returns a boolean if a field has been set.
func (o *Agent) HasModelCategory() bool {
	if o != nil && !IsNil(o.ModelCategory) {
		return true
	}

	return false
}

// SetModelCategory gets a reference to the given string and assigns it to the ModelCategory field.
func (o *Agent) SetModelCategory(v string) {
	o.ModelCategory = &v
}

// GetModelHostLocation returns the ModelHostLocation field value if set, zero value otherwise.
func (o *Agent) GetModelHostLocation() HostingLocation {
	if o == nil || IsNil(o.ModelHostLocation) {
		var ret HostingLocation
		return ret
	}
	return *o.ModelHostLocation
}

// GetModelHostLocationOk returns a tuple with the ModelHostLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetModelHostLocationOk() (*HostingLocation, bool) {
	if o == nil || IsNil(o.ModelHostLocation) {
		return nil, false
	}
	return o.ModelHostLocation, true
}

// HasModelHostLocation returns a boolean if a field has been set.
func (o *Agent) HasModelHostLocation() bool {
	if o != nil && !IsNil(o.ModelHostLocation) {
		return true
	}

	return false
}

// SetModelHostLocation gets a reference to the given HostingLocation and assigns it to the ModelHostLocation field.
func (o *Agent) SetModelHostLocation(v HostingLocation) {
	o.ModelHostLocation = &v
}

// GetModelId returns the ModelId field value
func (o *Agent) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *Agent) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *Agent) SetModelId(v string) {
	o.ModelId = v
}

// GetModelReleaseDate returns the ModelReleaseDate field value if set, zero value otherwise.
func (o *Agent) GetModelReleaseDate() int64 {
	if o == nil || IsNil(o.ModelReleaseDate) {
		var ret int64
		return ret
	}
	return *o.ModelReleaseDate
}

// GetModelReleaseDateOk returns a tuple with the ModelReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetModelReleaseDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ModelReleaseDate) {
		return nil, false
	}
	return o.ModelReleaseDate, true
}

// HasModelReleaseDate returns a boolean if a field has been set.
func (o *Agent) HasModelReleaseDate() bool {
	if o != nil && !IsNil(o.ModelReleaseDate) {
		return true
	}

	return false
}

// SetModelReleaseDate gets a reference to the given int64 and assigns it to the ModelReleaseDate field.
func (o *Agent) SetModelReleaseDate(v int64) {
	o.ModelReleaseDate = &v
}

// GetName returns the Name field value
func (o *Agent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Agent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Agent) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *Agent) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *Agent) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *Agent) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetOwnerOrganizationId returns the OwnerOrganizationId field value
func (o *Agent) GetOwnerOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerOrganizationId
}

// GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field value
// and a boolean to check if the value has been set.
func (o *Agent) GetOwnerOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerOrganizationId, true
}

// SetOwnerOrganizationId sets field value
func (o *Agent) SetOwnerOrganizationId(v string) {
	o.OwnerOrganizationId = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Agent) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Agent) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *Agent) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetSystemMessages returns the SystemMessages field value if set, zero value otherwise.
func (o *Agent) GetSystemMessages() []string {
	if o == nil || IsNil(o.SystemMessages) {
		var ret []string
		return ret
	}
	return o.SystemMessages
}

// GetSystemMessagesOk returns a tuple with the SystemMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetSystemMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.SystemMessages) {
		return nil, false
	}
	return o.SystemMessages, true
}

// HasSystemMessages returns a boolean if a field has been set.
func (o *Agent) HasSystemMessages() bool {
	if o != nil && !IsNil(o.SystemMessages) {
		return true
	}

	return false
}

// SetSystemMessages gets a reference to the given []string and assigns it to the SystemMessages field.
func (o *Agent) SetSystemMessages(v []string) {
	o.SystemMessages = v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *Agent) GetTeamIds() []string {
	if o == nil || IsNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *Agent) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *Agent) SetTeamIds(v []string) {
	o.TeamIds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Agent) GetType() AgentType {
	if o == nil || IsNil(o.Type) {
		var ret AgentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetTypeOk() (*AgentType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Agent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AgentType and assigns it to the Type field.
func (o *Agent) SetType(v AgentType) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Agent) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Agent) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Agent) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Agent) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Agent) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *Agent) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUseTools returns the UseTools field value if set, zero value otherwise.
func (o *Agent) GetUseTools() bool {
	if o == nil || IsNil(o.UseTools) {
		var ret bool
		return ret
	}
	return *o.UseTools
}

// GetUseToolsOk returns a tuple with the UseTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetUseToolsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTools) {
		return nil, false
	}
	return o.UseTools, true
}

// HasUseTools returns a boolean if a field has been set.
func (o *Agent) HasUseTools() bool {
	if o != nil && !IsNil(o.UseTools) {
		return true
	}

	return false
}

// SetUseTools gets a reference to the given bool and assigns it to the UseTools field.
func (o *Agent) SetUseTools(v bool) {
	o.UseTools = &v
}

func (o Agent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Agent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Abilities) {
		toSerialize["abilities"] = o.Abilities
	}
	if !IsNil(o.AbilitiesV2) {
		toSerialize["abilities_v2"] = o.AbilitiesV2
	}
	if !IsNil(o.AssignedTools) {
		toSerialize["assigned_tools"] = o.AssignedTools
	}
	if !IsNil(o.AttachedFileIds) {
		toSerialize["attached_file_ids"] = o.AttachedFileIds
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.I18n) {
		toSerialize["i18n"] = o.I18n
	}
	toSerialize["id"] = o.Id
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
	if !IsNil(o.ModelCategory) {
		toSerialize["model_category"] = o.ModelCategory
	}
	if !IsNil(o.ModelHostLocation) {
		toSerialize["model_host_location"] = o.ModelHostLocation
	}
	toSerialize["model_id"] = o.ModelId
	if !IsNil(o.ModelReleaseDate) {
		toSerialize["model_release_date"] = o.ModelReleaseDate
	}
	toSerialize["name"] = o.Name
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["owner_organization_id"] = o.OwnerOrganizationId
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.SystemMessages) {
		toSerialize["system_messages"] = o.SystemMessages
	}
	if !IsNil(o.TeamIds) {
		toSerialize["team_ids"] = o.TeamIds
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if !IsNil(o.UseTools) {
		toSerialize["use_tools"] = o.UseTools
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Agent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"model_id",
		"name",
		"owner_id",
		"owner_organization_id",
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

	varAgent := _Agent{}

	err = json.Unmarshal(data, &varAgent)

	if err != nil {
		return err
	}

	*o = Agent(varAgent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "abilities")
		delete(additionalProperties, "abilities_v2")
		delete(additionalProperties, "assigned_tools")
		delete(additionalProperties, "attached_file_ids")
		delete(additionalProperties, "avatar_url")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "i18n")
		delete(additionalProperties, "id")
		delete(additionalProperties, "initial_user_messages")
		delete(additionalProperties, "internal_id")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "meta_data")
		delete(additionalProperties, "model_category")
		delete(additionalProperties, "model_host_location")
		delete(additionalProperties, "model_id")
		delete(additionalProperties, "model_release_date")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "owner_organization_id")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "system_messages")
		delete(additionalProperties, "team_ids")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		delete(additionalProperties, "use_tools")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


