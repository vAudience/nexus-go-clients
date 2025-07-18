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

// checks if the Channel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Channel{}

// Channel struct for Channel
type Channel struct {
	CreatedAt *int64 `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	IsOrgPublic *bool `json:"is_org_public,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	LastMessageAt *int64 `json:"last_message_at,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	MissionId *string `json:"mission_id,omitempty"`
	Name string `json:"name"`
	OwnerId string `json:"owner_id"`
	OwnerOrganizationId string `json:"owner_organization_id"`
	Summary *string `json:"summary,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Channel Channel

// NewChannel instantiates a new Channel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannel(id string, name string, ownerId string, ownerOrganizationId string) *Channel {
	this := Channel{}
	this.Id = id
	this.Name = name
	this.OwnerId = ownerId
	this.OwnerOrganizationId = ownerOrganizationId
	return &this
}

// NewChannelWithDefaults instantiates a new Channel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelWithDefaults() *Channel {
	this := Channel{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Channel) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Channel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Channel) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Channel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Channel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Channel) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *Channel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Channel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Channel) SetId(v string) {
	o.Id = v
}

// GetIsOrgPublic returns the IsOrgPublic field value if set, zero value otherwise.
func (o *Channel) GetIsOrgPublic() bool {
	if o == nil || IsNil(o.IsOrgPublic) {
		var ret bool
		return ret
	}
	return *o.IsOrgPublic
}

// GetIsOrgPublicOk returns a tuple with the IsOrgPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetIsOrgPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrgPublic) {
		return nil, false
	}
	return o.IsOrgPublic, true
}

// HasIsOrgPublic returns a boolean if a field has been set.
func (o *Channel) HasIsOrgPublic() bool {
	if o != nil && !IsNil(o.IsOrgPublic) {
		return true
	}

	return false
}

// SetIsOrgPublic gets a reference to the given bool and assigns it to the IsOrgPublic field.
func (o *Channel) SetIsOrgPublic(v bool) {
	o.IsOrgPublic = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *Channel) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *Channel) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *Channel) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetLastMessageAt returns the LastMessageAt field value if set, zero value otherwise.
func (o *Channel) GetLastMessageAt() int64 {
	if o == nil || IsNil(o.LastMessageAt) {
		var ret int64
		return ret
	}
	return *o.LastMessageAt
}

// GetLastMessageAtOk returns a tuple with the LastMessageAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetLastMessageAtOk() (*int64, bool) {
	if o == nil || IsNil(o.LastMessageAt) {
		return nil, false
	}
	return o.LastMessageAt, true
}

// HasLastMessageAt returns a boolean if a field has been set.
func (o *Channel) HasLastMessageAt() bool {
	if o != nil && !IsNil(o.LastMessageAt) {
		return true
	}

	return false
}

// SetLastMessageAt gets a reference to the given int64 and assigns it to the LastMessageAt field.
func (o *Channel) SetLastMessageAt(v int64) {
	o.LastMessageAt = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Channel) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Channel) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Channel) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMissionId returns the MissionId field value if set, zero value otherwise.
func (o *Channel) GetMissionId() string {
	if o == nil || IsNil(o.MissionId) {
		var ret string
		return ret
	}
	return *o.MissionId
}

// GetMissionIdOk returns a tuple with the MissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetMissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.MissionId) {
		return nil, false
	}
	return o.MissionId, true
}

// HasMissionId returns a boolean if a field has been set.
func (o *Channel) HasMissionId() bool {
	if o != nil && !IsNil(o.MissionId) {
		return true
	}

	return false
}

// SetMissionId gets a reference to the given string and assigns it to the MissionId field.
func (o *Channel) SetMissionId(v string) {
	o.MissionId = &v
}

// GetName returns the Name field value
func (o *Channel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Channel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Channel) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *Channel) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *Channel) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *Channel) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetOwnerOrganizationId returns the OwnerOrganizationId field value
func (o *Channel) GetOwnerOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerOrganizationId
}

// GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field value
// and a boolean to check if the value has been set.
func (o *Channel) GetOwnerOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerOrganizationId, true
}

// SetOwnerOrganizationId sets field value
func (o *Channel) SetOwnerOrganizationId(v string) {
	o.OwnerOrganizationId = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Channel) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Channel) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Channel) SetSummary(v string) {
	o.Summary = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Channel) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Channel) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Channel) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o Channel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Channel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsOrgPublic) {
		toSerialize["is_org_public"] = o.IsOrgPublic
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.LastMessageAt) {
		toSerialize["last_message_at"] = o.LastMessageAt
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MissionId) {
		toSerialize["mission_id"] = o.MissionId
	}
	toSerialize["name"] = o.Name
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["owner_organization_id"] = o.OwnerOrganizationId
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Channel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varChannel := _Channel{}

	err = json.Unmarshal(data, &varChannel)

	if err != nil {
		return err
	}

	*o = Channel(varChannel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_org_public")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "last_message_at")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "mission_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "owner_organization_id")
		delete(additionalProperties, "summary")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannel struct {
	value *Channel
	isSet bool
}

func (v NullableChannel) Get() *Channel {
	return v.value
}

func (v *NullableChannel) Set(val *Channel) {
	v.value = val
	v.isSet = true
}

func (v NullableChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannel(val *Channel) *NullableChannel {
	return &NullableChannel{value: val, isSet: true}
}

func (v NullableChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


