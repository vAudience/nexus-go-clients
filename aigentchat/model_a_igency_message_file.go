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

// checks if the AIgencyMessageFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIgencyMessageFile{}

// AIgencyMessageFile struct for AIgencyMessageFile
type AIgencyMessageFile struct {
	EmbeddedContent *string `json:"embedded_content,omitempty"`
	FileName *string `json:"file_name,omitempty"`
	FileSize *int32 `json:"file_size,omitempty"`
	Id *string `json:"id,omitempty"`
	MetaData map[string]map[string]interface{} `json:"meta_data,omitempty"`
	MimeType *string `json:"mime_type,omitempty"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIgencyMessageFile AIgencyMessageFile

// NewAIgencyMessageFile instantiates a new AIgencyMessageFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIgencyMessageFile() *AIgencyMessageFile {
	this := AIgencyMessageFile{}
	return &this
}

// NewAIgencyMessageFileWithDefaults instantiates a new AIgencyMessageFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIgencyMessageFileWithDefaults() *AIgencyMessageFile {
	this := AIgencyMessageFile{}
	return &this
}

// GetEmbeddedContent returns the EmbeddedContent field value if set, zero value otherwise.
func (o *AIgencyMessageFile) GetEmbeddedContent() string {
	if o == nil || IsNil(o.EmbeddedContent) {
		var ret string
		return ret
	}
	return *o.EmbeddedContent
}

// GetEmbeddedContentOk returns a tuple with the EmbeddedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetEmbeddedContentOk() (*string, bool) {
	if o == nil || IsNil(o.EmbeddedContent) {
		return nil, false
	}
	return o.EmbeddedContent, true
}

// HasEmbeddedContent returns a boolean if a field has been set.
func (o *AIgencyMessageFile) HasEmbeddedContent() bool {
	if o != nil && !IsNil(o.EmbeddedContent) {
		return true
	}

	return false
}

// SetEmbeddedContent gets a reference to the given string and assigns it to the EmbeddedContent field.
func (o *AIgencyMessageFile) SetEmbeddedContent(v string) {
	o.EmbeddedContent = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AIgencyMessageFile) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AIgencyMessageFile) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AIgencyMessageFile) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *AIgencyMessageFile) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *AIgencyMessageFile) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *AIgencyMessageFile) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AIgencyMessageFile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AIgencyMessageFile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AIgencyMessageFile) SetId(v string) {
	o.Id = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *AIgencyMessageFile) GetMetaData() map[string]map[string]interface{} {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetMetaDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return map[string]map[string]interface{}{}, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *AIgencyMessageFile) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]map[string]interface{} and assigns it to the MetaData field.
func (o *AIgencyMessageFile) SetMetaData(v map[string]map[string]interface{}) {
	o.MetaData = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *AIgencyMessageFile) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *AIgencyMessageFile) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *AIgencyMessageFile) SetMimeType(v string) {
	o.MimeType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AIgencyMessageFile) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AIgencyMessageFile) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AIgencyMessageFile) SetUrl(v string) {
	o.Url = &v
}

func (o AIgencyMessageFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIgencyMessageFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmbeddedContent) {
		toSerialize["embedded_content"] = o.EmbeddedContent
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.FileSize) {
		toSerialize["file_size"] = o.FileSize
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MetaData) {
		toSerialize["meta_data"] = o.MetaData
	}
	if !IsNil(o.MimeType) {
		toSerialize["mime_type"] = o.MimeType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIgencyMessageFile) UnmarshalJSON(data []byte) (err error) {
	varAIgencyMessageFile := _AIgencyMessageFile{}

	err = json.Unmarshal(data, &varAIgencyMessageFile)

	if err != nil {
		return err
	}

	*o = AIgencyMessageFile(varAIgencyMessageFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "embedded_content")
		delete(additionalProperties, "file_name")
		delete(additionalProperties, "file_size")
		delete(additionalProperties, "id")
		delete(additionalProperties, "meta_data")
		delete(additionalProperties, "mime_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIgencyMessageFile struct {
	value *AIgencyMessageFile
	isSet bool
}

func (v NullableAIgencyMessageFile) Get() *AIgencyMessageFile {
	return v.value
}

func (v *NullableAIgencyMessageFile) Set(val *AIgencyMessageFile) {
	v.value = val
	v.isSet = true
}

func (v NullableAIgencyMessageFile) IsSet() bool {
	return v.isSet
}

func (v *NullableAIgencyMessageFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIgencyMessageFile(val *AIgencyMessageFile) *NullableAIgencyMessageFile {
	return &NullableAIgencyMessageFile{value: val, isSet: true}
}

func (v NullableAIgencyMessageFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIgencyMessageFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


