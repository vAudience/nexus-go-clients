/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.18.0
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the AIModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIModel{}

// AIModel struct for AIModel
type AIModel struct {
	Abilities []Ability `json:"abilities,omitempty"`
	AcceptedFileMimetypes []string `json:"accepted_file_mimetypes,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	Features []AIModelFeature `json:"features,omitempty"`
	I18n *map[string]AIModelI18n `json:"i18n,omitempty"`
	Id string `json:"id"`
	InternalId *string `json:"internal_id,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	ModelId string `json:"model_id"`
	Name string `json:"name"`
	OwnerId string `json:"owner_id"`
	OwnerOrganizationId string `json:"owner_organization_id"`
	ParameterDefinitions map[string]interface{} `json:"parameter_definitions,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	ServiceHostLocations []HostingLocation `json:"service_host_locations,omitempty"`
	ServiceId string `json:"service_id"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIModel AIModel

// NewAIModel instantiates a new AIModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIModel(id string, modelId string, name string, ownerId string, ownerOrganizationId string, serviceId string) *AIModel {
	this := AIModel{}
	this.Id = id
	this.ModelId = modelId
	this.Name = name
	this.OwnerId = ownerId
	this.OwnerOrganizationId = ownerOrganizationId
	this.ServiceId = serviceId
	return &this
}

// NewAIModelWithDefaults instantiates a new AIModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIModelWithDefaults() *AIModel {
	this := AIModel{}
	return &this
}

// GetAbilities returns the Abilities field value if set, zero value otherwise.
func (o *AIModel) GetAbilities() []Ability {
	if o == nil || IsNil(o.Abilities) {
		var ret []Ability
		return ret
	}
	return o.Abilities
}

// GetAbilitiesOk returns a tuple with the Abilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetAbilitiesOk() ([]Ability, bool) {
	if o == nil || IsNil(o.Abilities) {
		return nil, false
	}
	return o.Abilities, true
}

// HasAbilities returns a boolean if a field has been set.
func (o *AIModel) HasAbilities() bool {
	if o != nil && !IsNil(o.Abilities) {
		return true
	}

	return false
}

// SetAbilities gets a reference to the given []Ability and assigns it to the Abilities field.
func (o *AIModel) SetAbilities(v []Ability) {
	o.Abilities = v
}

// GetAcceptedFileMimetypes returns the AcceptedFileMimetypes field value if set, zero value otherwise.
func (o *AIModel) GetAcceptedFileMimetypes() []string {
	if o == nil || IsNil(o.AcceptedFileMimetypes) {
		var ret []string
		return ret
	}
	return o.AcceptedFileMimetypes
}

// GetAcceptedFileMimetypesOk returns a tuple with the AcceptedFileMimetypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetAcceptedFileMimetypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AcceptedFileMimetypes) {
		return nil, false
	}
	return o.AcceptedFileMimetypes, true
}

// HasAcceptedFileMimetypes returns a boolean if a field has been set.
func (o *AIModel) HasAcceptedFileMimetypes() bool {
	if o != nil && !IsNil(o.AcceptedFileMimetypes) {
		return true
	}

	return false
}

// SetAcceptedFileMimetypes gets a reference to the given []string and assigns it to the AcceptedFileMimetypes field.
func (o *AIModel) SetAcceptedFileMimetypes(v []string) {
	o.AcceptedFileMimetypes = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AIModel) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AIModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *AIModel) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AIModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AIModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AIModel) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *AIModel) GetDocumentationUrl() string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentationUrl) {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *AIModel) HasDocumentationUrl() bool {
	if o != nil && !IsNil(o.DocumentationUrl) {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *AIModel) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *AIModel) GetFeatures() []AIModelFeature {
	if o == nil || IsNil(o.Features) {
		var ret []AIModelFeature
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetFeaturesOk() ([]AIModelFeature, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *AIModel) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []AIModelFeature and assigns it to the Features field.
func (o *AIModel) SetFeatures(v []AIModelFeature) {
	o.Features = v
}

// GetI18n returns the I18n field value if set, zero value otherwise.
func (o *AIModel) GetI18n() map[string]AIModelI18n {
	if o == nil || IsNil(o.I18n) {
		var ret map[string]AIModelI18n
		return ret
	}
	return *o.I18n
}

// GetI18nOk returns a tuple with the I18n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetI18nOk() (*map[string]AIModelI18n, bool) {
	if o == nil || IsNil(o.I18n) {
		return nil, false
	}
	return o.I18n, true
}

// HasI18n returns a boolean if a field has been set.
func (o *AIModel) HasI18n() bool {
	if o != nil && !IsNil(o.I18n) {
		return true
	}

	return false
}

// SetI18n gets a reference to the given map[string]AIModelI18n and assigns it to the I18n field.
func (o *AIModel) SetI18n(v map[string]AIModelI18n) {
	o.I18n = &v
}

// GetId returns the Id field value
func (o *AIModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AIModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AIModel) SetId(v string) {
	o.Id = v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *AIModel) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// HasInternalId returns a boolean if a field has been set.
func (o *AIModel) HasInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *AIModel) SetInternalId(v string) {
	o.InternalId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *AIModel) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *AIModel) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *AIModel) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetModelId returns the ModelId field value
func (o *AIModel) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *AIModel) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *AIModel) SetModelId(v string) {
	o.ModelId = v
}

// GetName returns the Name field value
func (o *AIModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AIModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AIModel) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *AIModel) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *AIModel) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *AIModel) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetOwnerOrganizationId returns the OwnerOrganizationId field value
func (o *AIModel) GetOwnerOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerOrganizationId
}

// GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field value
// and a boolean to check if the value has been set.
func (o *AIModel) GetOwnerOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerOrganizationId, true
}

// SetOwnerOrganizationId sets field value
func (o *AIModel) SetOwnerOrganizationId(v string) {
	o.OwnerOrganizationId = v
}

// GetParameterDefinitions returns the ParameterDefinitions field value if set, zero value otherwise.
func (o *AIModel) GetParameterDefinitions() map[string]interface{} {
	if o == nil || IsNil(o.ParameterDefinitions) {
		var ret map[string]interface{}
		return ret
	}
	return o.ParameterDefinitions
}

// GetParameterDefinitionsOk returns a tuple with the ParameterDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetParameterDefinitionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ParameterDefinitions) {
		return map[string]interface{}{}, false
	}
	return o.ParameterDefinitions, true
}

// HasParameterDefinitions returns a boolean if a field has been set.
func (o *AIModel) HasParameterDefinitions() bool {
	if o != nil && !IsNil(o.ParameterDefinitions) {
		return true
	}

	return false
}

// SetParameterDefinitions gets a reference to the given map[string]interface{} and assigns it to the ParameterDefinitions field.
func (o *AIModel) SetParameterDefinitions(v map[string]interface{}) {
	o.ParameterDefinitions = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *AIModel) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AIModel) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *AIModel) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetServiceHostLocations returns the ServiceHostLocations field value if set, zero value otherwise.
func (o *AIModel) GetServiceHostLocations() []HostingLocation {
	if o == nil || IsNil(o.ServiceHostLocations) {
		var ret []HostingLocation
		return ret
	}
	return o.ServiceHostLocations
}

// GetServiceHostLocationsOk returns a tuple with the ServiceHostLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetServiceHostLocationsOk() ([]HostingLocation, bool) {
	if o == nil || IsNil(o.ServiceHostLocations) {
		return nil, false
	}
	return o.ServiceHostLocations, true
}

// HasServiceHostLocations returns a boolean if a field has been set.
func (o *AIModel) HasServiceHostLocations() bool {
	if o != nil && !IsNil(o.ServiceHostLocations) {
		return true
	}

	return false
}

// SetServiceHostLocations gets a reference to the given []HostingLocation and assigns it to the ServiceHostLocations field.
func (o *AIModel) SetServiceHostLocations(v []HostingLocation) {
	o.ServiceHostLocations = v
}

// GetServiceId returns the ServiceId field value
func (o *AIModel) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *AIModel) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *AIModel) SetServiceId(v string) {
	o.ServiceId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AIModel) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AIModel) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *AIModel) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *AIModel) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModel) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AIModel) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *AIModel) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o AIModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Abilities) {
		toSerialize["abilities"] = o.Abilities
	}
	if !IsNil(o.AcceptedFileMimetypes) {
		toSerialize["accepted_file_mimetypes"] = o.AcceptedFileMimetypes
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
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
	toSerialize["id"] = o.Id
	if !IsNil(o.InternalId) {
		toSerialize["internal_id"] = o.InternalId
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	toSerialize["model_id"] = o.ModelId
	toSerialize["name"] = o.Name
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["owner_organization_id"] = o.OwnerOrganizationId
	if !IsNil(o.ParameterDefinitions) {
		toSerialize["parameter_definitions"] = o.ParameterDefinitions
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.ServiceHostLocations) {
		toSerialize["service_host_locations"] = o.ServiceHostLocations
	}
	toSerialize["service_id"] = o.ServiceId
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"model_id",
		"name",
		"owner_id",
		"owner_organization_id",
		"service_id",
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

	varAIModel := _AIModel{}

	err = json.Unmarshal(data, &varAIModel)

	if err != nil {
		return err
	}

	*o = AIModel(varAIModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "abilities")
		delete(additionalProperties, "accepted_file_mimetypes")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "documentation_url")
		delete(additionalProperties, "features")
		delete(additionalProperties, "i18n")
		delete(additionalProperties, "id")
		delete(additionalProperties, "internal_id")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "model_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "owner_organization_id")
		delete(additionalProperties, "parameter_definitions")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "service_host_locations")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIModel struct {
	value *AIModel
	isSet bool
}

func (v NullableAIModel) Get() *AIModel {
	return v.value
}

func (v *NullableAIModel) Set(val *AIModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModel(val *AIModel) *NullableAIModel {
	return &NullableAIModel{value: val, isSet: true}
}

func (v NullableAIModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


