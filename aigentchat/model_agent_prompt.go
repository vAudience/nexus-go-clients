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
	"fmt"
)

// checks if the AgentPrompt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPrompt{}

// AgentPrompt struct for AgentPrompt
type AgentPrompt struct {
	CreatedAt *int32 `json:"created_at,omitempty"`
	CurrentVersion *int32 `json:"current_version,omitempty"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	OwnerId string `json:"owner_id"`
	OwnerOrganizationId string `json:"owner_organization_id"`
	Space *string `json:"space,omitempty"`
	Tags []string `json:"tags,omitempty"`
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	Title string `json:"title"`
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	Versions []PromptVersion `json:"versions,omitempty"`
	Visibility *AgentPromptVisibilityStates `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentPrompt AgentPrompt

// NewAgentPrompt instantiates a new AgentPrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPrompt(id string, ownerId string, ownerOrganizationId string, title string) *AgentPrompt {
	this := AgentPrompt{}
	this.Id = id
	this.OwnerId = ownerId
	this.OwnerOrganizationId = ownerOrganizationId
	this.Title = title
	return &this
}

// NewAgentPromptWithDefaults instantiates a new AgentPrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPromptWithDefaults() *AgentPrompt {
	this := AgentPrompt{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AgentPrompt) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AgentPrompt) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *AgentPrompt) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *AgentPrompt) GetCurrentVersion() int32 {
	if o == nil || IsNil(o.CurrentVersion) {
		var ret int32
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetCurrentVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *AgentPrompt) HasCurrentVersion() bool {
	if o != nil && !IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given int32 and assigns it to the CurrentVersion field.
func (o *AgentPrompt) SetCurrentVersion(v int32) {
	o.CurrentVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AgentPrompt) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AgentPrompt) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AgentPrompt) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *AgentPrompt) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentPrompt) SetId(v string) {
	o.Id = v
}

// GetOwnerId returns the OwnerId field value
func (o *AgentPrompt) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *AgentPrompt) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetOwnerOrganizationId returns the OwnerOrganizationId field value
func (o *AgentPrompt) GetOwnerOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerOrganizationId
}

// GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field value
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetOwnerOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerOrganizationId, true
}

// SetOwnerOrganizationId sets field value
func (o *AgentPrompt) SetOwnerOrganizationId(v string) {
	o.OwnerOrganizationId = v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *AgentPrompt) GetSpace() string {
	if o == nil || IsNil(o.Space) {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *AgentPrompt) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *AgentPrompt) SetSpace(v string) {
	o.Space = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AgentPrompt) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AgentPrompt) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AgentPrompt) SetTags(v []string) {
	o.Tags = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *AgentPrompt) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *AgentPrompt) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *AgentPrompt) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

// GetTitle returns the Title field value
func (o *AgentPrompt) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AgentPrompt) SetTitle(v string) {
	o.Title = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AgentPrompt) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AgentPrompt) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *AgentPrompt) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AgentPrompt) GetVersions() []PromptVersion {
	if o == nil || IsNil(o.Versions) {
		var ret []PromptVersion
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetVersionsOk() ([]PromptVersion, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AgentPrompt) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []PromptVersion and assigns it to the Versions field.
func (o *AgentPrompt) SetVersions(v []PromptVersion) {
	o.Versions = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *AgentPrompt) GetVisibility() AgentPromptVisibilityStates {
	if o == nil || IsNil(o.Visibility) {
		var ret AgentPromptVisibilityStates
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPrompt) GetVisibilityOk() (*AgentPromptVisibilityStates, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *AgentPrompt) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given AgentPromptVisibilityStates and assigns it to the Visibility field.
func (o *AgentPrompt) SetVisibility(v AgentPromptVisibilityStates) {
	o.Visibility = &v
}

func (o AgentPrompt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPrompt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CurrentVersion) {
		toSerialize["current_version"] = o.CurrentVersion
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["owner_organization_id"] = o.OwnerOrganizationId
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ThumbnailUrl) {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	toSerialize["title"] = o.Title
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentPrompt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"owner_id",
		"owner_organization_id",
		"title",
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

	varAgentPrompt := _AgentPrompt{}

	err = json.Unmarshal(data, &varAgentPrompt)

	if err != nil {
		return err
	}

	*o = AgentPrompt(varAgentPrompt)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "current_version")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "owner_organization_id")
		delete(additionalProperties, "space")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "thumbnail_url")
		delete(additionalProperties, "title")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "versions")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentPrompt struct {
	value *AgentPrompt
	isSet bool
}

func (v NullableAgentPrompt) Get() *AgentPrompt {
	return v.value
}

func (v *NullableAgentPrompt) Set(val *AgentPrompt) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPrompt) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPrompt(val *AgentPrompt) *NullableAgentPrompt {
	return &NullableAgentPrompt{value: val, isSet: true}
}

func (v NullableAgentPrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


