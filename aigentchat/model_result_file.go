/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.18.2
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the ResultFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultFile{}

// ResultFile struct for ResultFile
type ResultFile struct {
	FileName string `json:"file_name"`
	FileSize int32 `json:"file_size"`
	MimeType string `json:"mime_type"`
	OriginalFileName string `json:"original_file_name"`
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _ResultFile ResultFile

// NewResultFile instantiates a new ResultFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultFile(fileName string, fileSize int32, mimeType string, originalFileName string, url string) *ResultFile {
	this := ResultFile{}
	this.FileName = fileName
	this.FileSize = fileSize
	this.MimeType = mimeType
	this.OriginalFileName = originalFileName
	this.Url = url
	return &this
}

// NewResultFileWithDefaults instantiates a new ResultFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultFileWithDefaults() *ResultFile {
	this := ResultFile{}
	return &this
}

// GetFileName returns the FileName field value
func (o *ResultFile) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *ResultFile) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *ResultFile) SetFileName(v string) {
	o.FileName = v
}

// GetFileSize returns the FileSize field value
func (o *ResultFile) GetFileSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *ResultFile) GetFileSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *ResultFile) SetFileSize(v int32) {
	o.FileSize = v
}

// GetMimeType returns the MimeType field value
func (o *ResultFile) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *ResultFile) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *ResultFile) SetMimeType(v string) {
	o.MimeType = v
}

// GetOriginalFileName returns the OriginalFileName field value
func (o *ResultFile) GetOriginalFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalFileName
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value
// and a boolean to check if the value has been set.
func (o *ResultFile) GetOriginalFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalFileName, true
}

// SetOriginalFileName sets field value
func (o *ResultFile) SetOriginalFileName(v string) {
	o.OriginalFileName = v
}

// GetUrl returns the Url field value
func (o *ResultFile) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ResultFile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ResultFile) SetUrl(v string) {
	o.Url = v
}

func (o ResultFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_name"] = o.FileName
	toSerialize["file_size"] = o.FileSize
	toSerialize["mime_type"] = o.MimeType
	toSerialize["original_file_name"] = o.OriginalFileName
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResultFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_name",
		"file_size",
		"mime_type",
		"original_file_name",
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

	varResultFile := _ResultFile{}

	err = json.Unmarshal(data, &varResultFile)

	if err != nil {
		return err
	}

	*o = ResultFile(varResultFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file_name")
		delete(additionalProperties, "file_size")
		delete(additionalProperties, "mime_type")
		delete(additionalProperties, "original_file_name")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResultFile struct {
	value *ResultFile
	isSet bool
}

func (v NullableResultFile) Get() *ResultFile {
	return v.value
}

func (v *NullableResultFile) Set(val *ResultFile) {
	v.value = val
	v.isSet = true
}

func (v NullableResultFile) IsSet() bool {
	return v.isSet
}

func (v *NullableResultFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultFile(val *ResultFile) *NullableResultFile {
	return &NullableResultFile{value: val, isSet: true}
}

func (v NullableResultFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


