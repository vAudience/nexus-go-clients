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
)

// checks if the AIModelWriteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIModelWriteDto{}

// AIModelWriteDto struct for AIModelWriteDto
type AIModelWriteDto struct {
	AcceptedFileMimetypes []string `json:"accepted_file_mimetypes,omitempty"`
	Description *string `json:"description,omitempty"`
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	Features []AIModelFeature `json:"features,omitempty"`
	I18n *map[string]AIModelI18n `json:"i18n,omitempty"`
	InternalId *string `json:"internal_id,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	ModelId *string `json:"model_id,omitempty"`
	Name *string `json:"name,omitempty"`
	ParameterDefinitions map[string]interface{} `json:"parameter_definitions,omitempty"`
	Parameters map[string]map[string]interface{} `json:"parameters,omitempty"`
	ServiceHostLocations []HostingLocation `json:"service_host_locations,omitempty"`
	ServiceId *string `json:"service_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIModelWriteDto AIModelWriteDto

// NewAIModelWriteDto instantiates a new AIModelWriteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIModelWriteDto() *AIModelWriteDto {
	this := AIModelWriteDto{}
	return &this
}

// NewAIModelWriteDtoWithDefaults instantiates a new AIModelWriteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIModelWriteDtoWithDefaults() *AIModelWriteDto {
	this := AIModelWriteDto{}
	return &this
}

// GetAcceptedFileMimetypes returns the AcceptedFileMimetypes field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetAcceptedFileMimetypes() []string {
	if o == nil || IsNil(o.AcceptedFileMimetypes) {
		var ret []string
		return ret
	}
	return o.AcceptedFileMimetypes
}

// GetAcceptedFileMimetypesOk returns a tuple with the AcceptedFileMimetypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetAcceptedFileMimetypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AcceptedFileMimetypes) {
		return nil, false
	}
	return o.AcceptedFileMimetypes, true
}

// HasAcceptedFileMimetypes returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasAcceptedFileMimetypes() bool {
	if o != nil && !IsNil(o.AcceptedFileMimetypes) {
		return true
	}

	return false
}

// SetAcceptedFileMimetypes gets a reference to the given []string and assigns it to the AcceptedFileMimetypes field.
func (o *AIModelWriteDto) SetAcceptedFileMimetypes(v []string) {
	o.AcceptedFileMimetypes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AIModelWriteDto) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetDocumentationUrl() string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentationUrl) {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasDocumentationUrl() bool {
	if o != nil && !IsNil(o.DocumentationUrl) {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *AIModelWriteDto) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetFeatures() []AIModelFeature {
	if o == nil || IsNil(o.Features) {
		var ret []AIModelFeature
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetFeaturesOk() ([]AIModelFeature, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []AIModelFeature and assigns it to the Features field.
func (o *AIModelWriteDto) SetFeatures(v []AIModelFeature) {
	o.Features = v
}

// GetI18n returns the I18n field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetI18n() map[string]AIModelI18n {
	if o == nil || IsNil(o.I18n) {
		var ret map[string]AIModelI18n
		return ret
	}
	return *o.I18n
}

// GetI18nOk returns a tuple with the I18n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetI18nOk() (*map[string]AIModelI18n, bool) {
	if o == nil || IsNil(o.I18n) {
		return nil, false
	}
	return o.I18n, true
}

// HasI18n returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasI18n() bool {
	if o != nil && !IsNil(o.I18n) {
		return true
	}

	return false
}

// SetI18n gets a reference to the given map[string]AIModelI18n and assigns it to the I18n field.
func (o *AIModelWriteDto) SetI18n(v map[string]AIModelI18n) {
	o.I18n = &v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// HasInternalId returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *AIModelWriteDto) SetInternalId(v string) {
	o.InternalId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *AIModelWriteDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetModelId() string {
	if o == nil || IsNil(o.ModelId) {
		var ret string
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given string and assigns it to the ModelId field.
func (o *AIModelWriteDto) SetModelId(v string) {
	o.ModelId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AIModelWriteDto) SetName(v string) {
	o.Name = &v
}

// GetParameterDefinitions returns the ParameterDefinitions field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetParameterDefinitions() map[string]interface{} {
	if o == nil || IsNil(o.ParameterDefinitions) {
		var ret map[string]interface{}
		return ret
	}
	return o.ParameterDefinitions
}

// GetParameterDefinitionsOk returns a tuple with the ParameterDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetParameterDefinitionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ParameterDefinitions) {
		return map[string]interface{}{}, false
	}
	return o.ParameterDefinitions, true
}

// HasParameterDefinitions returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasParameterDefinitions() bool {
	if o != nil && !IsNil(o.ParameterDefinitions) {
		return true
	}

	return false
}

// SetParameterDefinitions gets a reference to the given map[string]interface{} and assigns it to the ParameterDefinitions field.
func (o *AIModelWriteDto) SetParameterDefinitions(v map[string]interface{}) {
	o.ParameterDefinitions = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetParameters() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetParametersOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]map[string]interface{} and assigns it to the Parameters field.
func (o *AIModelWriteDto) SetParameters(v map[string]map[string]interface{}) {
	o.Parameters = v
}

// GetServiceHostLocations returns the ServiceHostLocations field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetServiceHostLocations() []HostingLocation {
	if o == nil || IsNil(o.ServiceHostLocations) {
		var ret []HostingLocation
		return ret
	}
	return o.ServiceHostLocations
}

// GetServiceHostLocationsOk returns a tuple with the ServiceHostLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetServiceHostLocationsOk() ([]HostingLocation, bool) {
	if o == nil || IsNil(o.ServiceHostLocations) {
		return nil, false
	}
	return o.ServiceHostLocations, true
}

// HasServiceHostLocations returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasServiceHostLocations() bool {
	if o != nil && !IsNil(o.ServiceHostLocations) {
		return true
	}

	return false
}

// SetServiceHostLocations gets a reference to the given []HostingLocation and assigns it to the ServiceHostLocations field.
func (o *AIModelWriteDto) SetServiceHostLocations(v []HostingLocation) {
	o.ServiceHostLocations = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *AIModelWriteDto) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelWriteDto) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *AIModelWriteDto) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *AIModelWriteDto) SetServiceId(v string) {
	o.ServiceId = &v
}

func (o AIModelWriteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIModelWriteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptedFileMimetypes) {
		toSerialize["accepted_file_mimetypes"] = o.AcceptedFileMimetypes
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DocumentationUrl) {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.I18n) {
		toSerialize["i18n"] = o.I18n
	}
	if !IsNil(o.InternalId) {
		toSerialize["internal_id"] = o.InternalId
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.ModelId) {
		toSerialize["model_id"] = o.ModelId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ParameterDefinitions) {
		toSerialize["parameter_definitions"] = o.ParameterDefinitions
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.ServiceHostLocations) {
		toSerialize["service_host_locations"] = o.ServiceHostLocations
	}
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIModelWriteDto) UnmarshalJSON(data []byte) (err error) {
	varAIModelWriteDto := _AIModelWriteDto{}

	err = json.Unmarshal(data, &varAIModelWriteDto)

	if err != nil {
		return err
	}

	*o = AIModelWriteDto(varAIModelWriteDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accepted_file_mimetypes")
		delete(additionalProperties, "description")
		delete(additionalProperties, "documentation_url")
		delete(additionalProperties, "features")
		delete(additionalProperties, "i18n")
		delete(additionalProperties, "internal_id")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "model_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "parameter_definitions")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "service_host_locations")
		delete(additionalProperties, "service_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIModelWriteDto struct {
	value *AIModelWriteDto
	isSet bool
}

func (v NullableAIModelWriteDto) Get() *AIModelWriteDto {
	return v.value
}

func (v *NullableAIModelWriteDto) Set(val *AIModelWriteDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelWriteDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelWriteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelWriteDto(val *AIModelWriteDto) *NullableAIModelWriteDto {
	return &NullableAIModelWriteDto{value: val, isSet: true}
}

func (v NullableAIModelWriteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelWriteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


