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

// checks if the AIgencyMessageFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIgencyMessageFile{}

// AIgencyMessageFile struct for AIgencyMessageFile
type AIgencyMessageFile struct {
	EmbeddedContent string `json:"embedded_content"`
	FileName string `json:"file_name"`
	FileSize int32 `json:"file_size"`
	Id string `json:"id"`
	MetaData map[string]interface{} `json:"meta_data"`
	MimeType string `json:"mime_type"`
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _AIgencyMessageFile AIgencyMessageFile

// NewAIgencyMessageFile instantiates a new AIgencyMessageFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIgencyMessageFile(embeddedContent string, fileName string, fileSize int32, id string, metaData map[string]interface{}, mimeType string, url string) *AIgencyMessageFile {
	this := AIgencyMessageFile{}
	this.EmbeddedContent = embeddedContent
	this.FileName = fileName
	this.FileSize = fileSize
	this.Id = id
	this.MetaData = metaData
	this.MimeType = mimeType
	this.Url = url
	return &this
}

// NewAIgencyMessageFileWithDefaults instantiates a new AIgencyMessageFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIgencyMessageFileWithDefaults() *AIgencyMessageFile {
	this := AIgencyMessageFile{}
	return &this
}

// GetEmbeddedContent returns the EmbeddedContent field value
func (o *AIgencyMessageFile) GetEmbeddedContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmbeddedContent
}

// GetEmbeddedContentOk returns a tuple with the EmbeddedContent field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetEmbeddedContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmbeddedContent, true
}

// SetEmbeddedContent sets field value
func (o *AIgencyMessageFile) SetEmbeddedContent(v string) {
	o.EmbeddedContent = v
}

// GetFileName returns the FileName field value
func (o *AIgencyMessageFile) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *AIgencyMessageFile) SetFileName(v string) {
	o.FileName = v
}

// GetFileSize returns the FileSize field value
func (o *AIgencyMessageFile) GetFileSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetFileSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *AIgencyMessageFile) SetFileSize(v int32) {
	o.FileSize = v
}

// GetId returns the Id field value
func (o *AIgencyMessageFile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AIgencyMessageFile) SetId(v string) {
	o.Id = v
}

// GetMetaData returns the MetaData field value
func (o *AIgencyMessageFile) GetMetaData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.MetaData, true
}

// SetMetaData sets field value
func (o *AIgencyMessageFile) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetMimeType returns the MimeType field value
func (o *AIgencyMessageFile) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *AIgencyMessageFile) SetMimeType(v string) {
	o.MimeType = v
}

// GetUrl returns the Url field value
func (o *AIgencyMessageFile) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessageFile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AIgencyMessageFile) SetUrl(v string) {
	o.Url = v
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
	toSerialize["embedded_content"] = o.EmbeddedContent
	toSerialize["file_name"] = o.FileName
	toSerialize["file_size"] = o.FileSize
	toSerialize["id"] = o.Id
	toSerialize["meta_data"] = o.MetaData
	toSerialize["mime_type"] = o.MimeType
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIgencyMessageFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"embedded_content",
		"file_name",
		"file_size",
		"id",
		"meta_data",
		"mime_type",
		"url",
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


