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

// checks if the ChannelWriteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelWriteDto{}

// ChannelWriteDto struct for ChannelWriteDto
type ChannelWriteDto struct {
	Description *string `json:"description,omitempty"`
	IsOrgPublic *bool `json:"is_org_public,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	MissionId *string `json:"mission_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Summary *string `json:"summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelWriteDto ChannelWriteDto

// NewChannelWriteDto instantiates a new ChannelWriteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelWriteDto() *ChannelWriteDto {
	this := ChannelWriteDto{}
	return &this
}

// NewChannelWriteDtoWithDefaults instantiates a new ChannelWriteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelWriteDtoWithDefaults() *ChannelWriteDto {
	this := ChannelWriteDto{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChannelWriteDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelWriteDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChannelWriteDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChannelWriteDto) SetDescription(v string) {
	o.Description = &v
}

// GetIsOrgPublic returns the IsOrgPublic field value if set, zero value otherwise.
func (o *ChannelWriteDto) GetIsOrgPublic() bool {
	if o == nil || IsNil(o.IsOrgPublic) {
		var ret bool
		return ret
	}
	return *o.IsOrgPublic
}

// GetIsOrgPublicOk returns a tuple with the IsOrgPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelWriteDto) GetIsOrgPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrgPublic) {
		return nil, false
	}
	return o.IsOrgPublic, true
}

// HasIsOrgPublic returns a boolean if a field has been set.
func (o *ChannelWriteDto) HasIsOrgPublic() bool {
	if o != nil && !IsNil(o.IsOrgPublic) {
		return true
	}

	return false
}

// SetIsOrgPublic gets a reference to the given bool and assigns it to the IsOrgPublic field.
func (o *ChannelWriteDto) SetIsOrgPublic(v bool) {
	o.IsOrgPublic = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *ChannelWriteDto) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelWriteDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *ChannelWriteDto) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *ChannelWriteDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ChannelWriteDto) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelWriteDto) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ChannelWriteDto) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ChannelWriteDto) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMissionId returns the MissionId field value if set, zero value otherwise.
func (o *ChannelWriteDto) GetMissionId() string {
	if o == nil || IsNil(o.MissionId) {
		var ret string
		return ret
	}
	return *o.MissionId
}

// GetMissionIdOk returns a tuple with the MissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelWriteDto) GetMissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.MissionId) {
		return nil, false
	}
	return o.MissionId, true
}

// HasMissionId returns a boolean if a field has been set.
func (o *ChannelWriteDto) HasMissionId() bool {
	if o != nil && !IsNil(o.MissionId) {
		return true
	}

	return false
}

// SetMissionId gets a reference to the given string and assigns it to the MissionId field.
func (o *ChannelWriteDto) SetMissionId(v string) {
	o.MissionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChannelWriteDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelWriteDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChannelWriteDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChannelWriteDto) SetName(v string) {
	o.Name = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ChannelWriteDto) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelWriteDto) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ChannelWriteDto) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ChannelWriteDto) SetSummary(v string) {
	o.Summary = &v
}

func (o ChannelWriteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelWriteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsOrgPublic) {
		toSerialize["is_org_public"] = o.IsOrgPublic
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MissionId) {
		toSerialize["mission_id"] = o.MissionId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChannelWriteDto) UnmarshalJSON(data []byte) (err error) {
	varChannelWriteDto := _ChannelWriteDto{}

	err = json.Unmarshal(data, &varChannelWriteDto)

	if err != nil {
		return err
	}

	*o = ChannelWriteDto(varChannelWriteDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "is_org_public")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "mission_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelWriteDto struct {
	value *ChannelWriteDto
	isSet bool
}

func (v NullableChannelWriteDto) Get() *ChannelWriteDto {
	return v.value
}

func (v *NullableChannelWriteDto) Set(val *ChannelWriteDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelWriteDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelWriteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelWriteDto(val *ChannelWriteDto) *NullableChannelWriteDto {
	return &NullableChannelWriteDto{value: val, isSet: true}
}

func (v NullableChannelWriteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelWriteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


