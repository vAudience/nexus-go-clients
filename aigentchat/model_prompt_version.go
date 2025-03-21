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

// checks if the PromptVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptVersion{}

// PromptVersion struct for PromptVersion
type PromptVersion struct {
	Content *string `json:"content,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Version *int32 `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PromptVersion PromptVersion

// NewPromptVersion instantiates a new PromptVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptVersion() *PromptVersion {
	this := PromptVersion{}
	return &this
}

// NewPromptVersionWithDefaults instantiates a new PromptVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptVersionWithDefaults() *PromptVersion {
	this := PromptVersion{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PromptVersion) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptVersion) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PromptVersion) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *PromptVersion) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PromptVersion) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptVersion) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PromptVersion) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *PromptVersion) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *PromptVersion) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptVersion) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PromptVersion) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *PromptVersion) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PromptVersion) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptVersion) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PromptVersion) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *PromptVersion) SetVersion(v int32) {
	o.Version = &v
}

func (o PromptVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PromptVersion) UnmarshalJSON(data []byte) (err error) {
	varPromptVersion := _PromptVersion{}

	err = json.Unmarshal(data, &varPromptVersion)

	if err != nil {
		return err
	}

	*o = PromptVersion(varPromptVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePromptVersion struct {
	value *PromptVersion
	isSet bool
}

func (v NullablePromptVersion) Get() *PromptVersion {
	return v.value
}

func (v *NullablePromptVersion) Set(val *PromptVersion) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptVersion) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptVersion(val *PromptVersion) *NullablePromptVersion {
	return &NullablePromptVersion{value: val, isSet: true}
}

func (v NullablePromptVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


