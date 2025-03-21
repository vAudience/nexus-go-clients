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
	"fmt"
)

// checks if the FileUploadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileUploadResponse{}

// FileUploadResponse struct for FileUploadResponse
type FileUploadResponse struct {
	Id string `json:"id"`
	ResultingFiles []ResultFile `json:"resulting_files"`
	AdditionalProperties map[string]interface{}
}

type _FileUploadResponse FileUploadResponse

// NewFileUploadResponse instantiates a new FileUploadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileUploadResponse(id string, resultingFiles []ResultFile) *FileUploadResponse {
	this := FileUploadResponse{}
	this.Id = id
	this.ResultingFiles = resultingFiles
	return &this
}

// NewFileUploadResponseWithDefaults instantiates a new FileUploadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileUploadResponseWithDefaults() *FileUploadResponse {
	this := FileUploadResponse{}
	return &this
}

// GetId returns the Id field value
func (o *FileUploadResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileUploadResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileUploadResponse) SetId(v string) {
	o.Id = v
}

// GetResultingFiles returns the ResultingFiles field value
func (o *FileUploadResponse) GetResultingFiles() []ResultFile {
	if o == nil {
		var ret []ResultFile
		return ret
	}

	return o.ResultingFiles
}

// GetResultingFilesOk returns a tuple with the ResultingFiles field value
// and a boolean to check if the value has been set.
func (o *FileUploadResponse) GetResultingFilesOk() ([]ResultFile, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultingFiles, true
}

// SetResultingFiles sets field value
func (o *FileUploadResponse) SetResultingFiles(v []ResultFile) {
	o.ResultingFiles = v
}

func (o FileUploadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileUploadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["resulting_files"] = o.ResultingFiles

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileUploadResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"resulting_files",
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

	varFileUploadResponse := _FileUploadResponse{}

	err = json.Unmarshal(data, &varFileUploadResponse)

	if err != nil {
		return err
	}

	*o = FileUploadResponse(varFileUploadResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "resulting_files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileUploadResponse struct {
	value *FileUploadResponse
	isSet bool
}

func (v NullableFileUploadResponse) Get() *FileUploadResponse {
	return v.value
}

func (v *NullableFileUploadResponse) Set(val *FileUploadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileUploadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileUploadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileUploadResponse(val *FileUploadResponse) *NullableFileUploadResponse {
	return &NullableFileUploadResponse{value: val, isSet: true}
}

func (v NullableFileUploadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileUploadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


