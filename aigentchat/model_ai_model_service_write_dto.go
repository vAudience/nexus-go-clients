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

// checks if the AIModelServiceWriteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIModelServiceWriteDto{}

// AIModelServiceWriteDto struct for AIModelServiceWriteDto
type AIModelServiceWriteDto struct {
	CostMultiplier *float64 `json:"cost_multiplier,omitempty"`
	Description *string `json:"description,omitempty"`
	HostingLocations *map[string]HostingLocation `json:"hosting_locations,omitempty"`
	I18n *map[string]AIModelServiceI18n `json:"i18n,omitempty"`
	InternalId *string `json:"internal_id,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	Name string `json:"name"`
	ServiceImpl AiServiceId `json:"service_impl"`
	AdditionalProperties map[string]interface{}
}

type _AIModelServiceWriteDto AIModelServiceWriteDto

// NewAIModelServiceWriteDto instantiates a new AIModelServiceWriteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIModelServiceWriteDto(name string, serviceImpl AiServiceId) *AIModelServiceWriteDto {
	this := AIModelServiceWriteDto{}
	this.Name = name
	this.ServiceImpl = serviceImpl
	return &this
}

// NewAIModelServiceWriteDtoWithDefaults instantiates a new AIModelServiceWriteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIModelServiceWriteDtoWithDefaults() *AIModelServiceWriteDto {
	this := AIModelServiceWriteDto{}
	return &this
}

// GetCostMultiplier returns the CostMultiplier field value if set, zero value otherwise.
func (o *AIModelServiceWriteDto) GetCostMultiplier() float64 {
	if o == nil || IsNil(o.CostMultiplier) {
		var ret float64
		return ret
	}
	return *o.CostMultiplier
}

// GetCostMultiplierOk returns a tuple with the CostMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceWriteDto) GetCostMultiplierOk() (*float64, bool) {
	if o == nil || IsNil(o.CostMultiplier) {
		return nil, false
	}
	return o.CostMultiplier, true
}

// HasCostMultiplier returns a boolean if a field has been set.
func (o *AIModelServiceWriteDto) HasCostMultiplier() bool {
	if o != nil && !IsNil(o.CostMultiplier) {
		return true
	}

	return false
}

// SetCostMultiplier gets a reference to the given float64 and assigns it to the CostMultiplier field.
func (o *AIModelServiceWriteDto) SetCostMultiplier(v float64) {
	o.CostMultiplier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AIModelServiceWriteDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceWriteDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AIModelServiceWriteDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AIModelServiceWriteDto) SetDescription(v string) {
	o.Description = &v
}

// GetHostingLocations returns the HostingLocations field value if set, zero value otherwise.
func (o *AIModelServiceWriteDto) GetHostingLocations() map[string]HostingLocation {
	if o == nil || IsNil(o.HostingLocations) {
		var ret map[string]HostingLocation
		return ret
	}
	return *o.HostingLocations
}

// GetHostingLocationsOk returns a tuple with the HostingLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceWriteDto) GetHostingLocationsOk() (*map[string]HostingLocation, bool) {
	if o == nil || IsNil(o.HostingLocations) {
		return nil, false
	}
	return o.HostingLocations, true
}

// HasHostingLocations returns a boolean if a field has been set.
func (o *AIModelServiceWriteDto) HasHostingLocations() bool {
	if o != nil && !IsNil(o.HostingLocations) {
		return true
	}

	return false
}

// SetHostingLocations gets a reference to the given map[string]HostingLocation and assigns it to the HostingLocations field.
func (o *AIModelServiceWriteDto) SetHostingLocations(v map[string]HostingLocation) {
	o.HostingLocations = &v
}

// GetI18n returns the I18n field value if set, zero value otherwise.
func (o *AIModelServiceWriteDto) GetI18n() map[string]AIModelServiceI18n {
	if o == nil || IsNil(o.I18n) {
		var ret map[string]AIModelServiceI18n
		return ret
	}
	return *o.I18n
}

// GetI18nOk returns a tuple with the I18n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceWriteDto) GetI18nOk() (*map[string]AIModelServiceI18n, bool) {
	if o == nil || IsNil(o.I18n) {
		return nil, false
	}
	return o.I18n, true
}

// HasI18n returns a boolean if a field has been set.
func (o *AIModelServiceWriteDto) HasI18n() bool {
	if o != nil && !IsNil(o.I18n) {
		return true
	}

	return false
}

// SetI18n gets a reference to the given map[string]AIModelServiceI18n and assigns it to the I18n field.
func (o *AIModelServiceWriteDto) SetI18n(v map[string]AIModelServiceI18n) {
	o.I18n = &v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *AIModelServiceWriteDto) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceWriteDto) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// HasInternalId returns a boolean if a field has been set.
func (o *AIModelServiceWriteDto) HasInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *AIModelServiceWriteDto) SetInternalId(v string) {
	o.InternalId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *AIModelServiceWriteDto) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceWriteDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *AIModelServiceWriteDto) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *AIModelServiceWriteDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetName returns the Name field value
func (o *AIModelServiceWriteDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AIModelServiceWriteDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AIModelServiceWriteDto) SetName(v string) {
	o.Name = v
}

// GetServiceImpl returns the ServiceImpl field value
func (o *AIModelServiceWriteDto) GetServiceImpl() AiServiceId {
	if o == nil {
		var ret AiServiceId
		return ret
	}

	return o.ServiceImpl
}

// GetServiceImplOk returns a tuple with the ServiceImpl field value
// and a boolean to check if the value has been set.
func (o *AIModelServiceWriteDto) GetServiceImplOk() (*AiServiceId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceImpl, true
}

// SetServiceImpl sets field value
func (o *AIModelServiceWriteDto) SetServiceImpl(v AiServiceId) {
	o.ServiceImpl = v
}

func (o AIModelServiceWriteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIModelServiceWriteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CostMultiplier) {
		toSerialize["cost_multiplier"] = o.CostMultiplier
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.HostingLocations) {
		toSerialize["hosting_locations"] = o.HostingLocations
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
	toSerialize["name"] = o.Name
	toSerialize["service_impl"] = o.ServiceImpl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIModelServiceWriteDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"service_impl",
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

	varAIModelServiceWriteDto := _AIModelServiceWriteDto{}

	err = json.Unmarshal(data, &varAIModelServiceWriteDto)

	if err != nil {
		return err
	}

	*o = AIModelServiceWriteDto(varAIModelServiceWriteDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cost_multiplier")
		delete(additionalProperties, "description")
		delete(additionalProperties, "hosting_locations")
		delete(additionalProperties, "i18n")
		delete(additionalProperties, "internal_id")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "name")
		delete(additionalProperties, "service_impl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIModelServiceWriteDto struct {
	value *AIModelServiceWriteDto
	isSet bool
}

func (v NullableAIModelServiceWriteDto) Get() *AIModelServiceWriteDto {
	return v.value
}

func (v *NullableAIModelServiceWriteDto) Set(val *AIModelServiceWriteDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelServiceWriteDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelServiceWriteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelServiceWriteDto(val *AIModelServiceWriteDto) *NullableAIModelServiceWriteDto {
	return &NullableAIModelServiceWriteDto{value: val, isSet: true}
}

func (v NullableAIModelServiceWriteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelServiceWriteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


