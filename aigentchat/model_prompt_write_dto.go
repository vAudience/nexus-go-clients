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

// checks if the PromptWriteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptWriteDto{}

// PromptWriteDto struct for PromptWriteDto
type PromptWriteDto struct {
	Content *string `json:"content,omitempty"`
	DefaultAgentId *string `json:"default_agent_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []string `json:"tags,omitempty"`
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	Title *string `json:"title,omitempty"`
	Visibility *PromptVisibilityStates `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PromptWriteDto PromptWriteDto

// NewPromptWriteDto instantiates a new PromptWriteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptWriteDto() *PromptWriteDto {
	this := PromptWriteDto{}
	return &this
}

// NewPromptWriteDtoWithDefaults instantiates a new PromptWriteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptWriteDtoWithDefaults() *PromptWriteDto {
	this := PromptWriteDto{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PromptWriteDto) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptWriteDto) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PromptWriteDto) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *PromptWriteDto) SetContent(v string) {
	o.Content = &v
}

// GetDefaultAgentId returns the DefaultAgentId field value if set, zero value otherwise.
func (o *PromptWriteDto) GetDefaultAgentId() string {
	if o == nil || IsNil(o.DefaultAgentId) {
		var ret string
		return ret
	}
	return *o.DefaultAgentId
}

// GetDefaultAgentIdOk returns a tuple with the DefaultAgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptWriteDto) GetDefaultAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAgentId) {
		return nil, false
	}
	return o.DefaultAgentId, true
}

// HasDefaultAgentId returns a boolean if a field has been set.
func (o *PromptWriteDto) HasDefaultAgentId() bool {
	if o != nil && !IsNil(o.DefaultAgentId) {
		return true
	}

	return false
}

// SetDefaultAgentId gets a reference to the given string and assigns it to the DefaultAgentId field.
func (o *PromptWriteDto) SetDefaultAgentId(v string) {
	o.DefaultAgentId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PromptWriteDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptWriteDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PromptWriteDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PromptWriteDto) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PromptWriteDto) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptWriteDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PromptWriteDto) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PromptWriteDto) SetTags(v []string) {
	o.Tags = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *PromptWriteDto) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptWriteDto) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *PromptWriteDto) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *PromptWriteDto) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PromptWriteDto) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptWriteDto) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PromptWriteDto) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PromptWriteDto) SetTitle(v string) {
	o.Title = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *PromptWriteDto) GetVisibility() PromptVisibilityStates {
	if o == nil || IsNil(o.Visibility) {
		var ret PromptVisibilityStates
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptWriteDto) GetVisibilityOk() (*PromptVisibilityStates, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *PromptWriteDto) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given PromptVisibilityStates and assigns it to the Visibility field.
func (o *PromptWriteDto) SetVisibility(v PromptVisibilityStates) {
	o.Visibility = &v
}

func (o PromptWriteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptWriteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.DefaultAgentId) {
		toSerialize["default_agent_id"] = o.DefaultAgentId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ThumbnailUrl) {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PromptWriteDto) UnmarshalJSON(data []byte) (err error) {
	varPromptWriteDto := _PromptWriteDto{}

	err = json.Unmarshal(data, &varPromptWriteDto)

	if err != nil {
		return err
	}

	*o = PromptWriteDto(varPromptWriteDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "default_agent_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "thumbnail_url")
		delete(additionalProperties, "title")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePromptWriteDto struct {
	value *PromptWriteDto
	isSet bool
}

func (v NullablePromptWriteDto) Get() *PromptWriteDto {
	return v.value
}

func (v *NullablePromptWriteDto) Set(val *PromptWriteDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptWriteDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptWriteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptWriteDto(val *PromptWriteDto) *NullablePromptWriteDto {
	return &NullablePromptWriteDto{value: val, isSet: true}
}

func (v NullablePromptWriteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptWriteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


